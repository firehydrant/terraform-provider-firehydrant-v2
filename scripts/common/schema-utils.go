package common

import (
	"fmt"
	"strings"
)

func ResolveRef(ref string, schemas map[string]interface{}) (map[string]interface{}, string) {
	if ref == "" {
		return nil, ""
	}

	parts := strings.Split(ref, "/")
	if len(parts) == 0 {
		return nil, ""
	}

	schemaName := parts[len(parts)-1]
	if schema, ok := schemas[schemaName].(map[string]interface{}); ok {
		// Handle nullable wrappers (e.g., NullableXXX schemas that use allOf)
		if allOf, hasAllOf := schema["allOf"].([]interface{}); hasAllOf && len(allOf) > 0 {
			if refMap, ok := allOf[0].(map[string]interface{}); ok {
				if innerRef, ok := refMap["$ref"].(string); ok {
					return ResolveRef(innerRef, schemas)
				}
			}
		}
		return schema, schemaName
	}

	return nil, ""
}

// GetResolvedPropertyType gets the actual type of a property, following refs if needed
func GetResolvedPropertyType(prop interface{}, schemas map[string]interface{}) (map[string]interface{}, string) {
	propMap, ok := prop.(map[string]interface{})
	if !ok {
		return nil, "unknown"
	}

	// If it's a ref, resolve it
	if ref, hasRef := propMap["$ref"].(string); hasRef {
		resolved, _ := ResolveRef(ref, schemas)
		if resolved != nil {
			return resolved, "ref"
		}
		return nil, "unknown"
	}

	return propMap, "inline"
}

func GetPropertyStructure(prop interface{}) string {
	propMap, ok := prop.(map[string]interface{})
	if !ok {
		return "unknown"
	}

	if ref, hasRef := propMap["$ref"].(string); hasRef {
		return fmt.Sprintf("$ref:%s", ref)
	}

	propType, _ := propMap["type"].(string)

	switch propType {
	case "array":
		items, hasItems := propMap["items"]
		if hasItems {
			itemStructure := GetPropertyStructure(items)
			return fmt.Sprintf("array[%s]", itemStructure)
		}
		return "array[unknown]"

	case "object":
		properties, hasProps := propMap["properties"]
		_, hasAdditional := propMap["additionalProperties"]

		if hasProps {
			propsMap, _ := properties.(map[string]interface{})
			if len(propsMap) == 0 {
				return "object{empty}"
			}
			return "object{defined}"
		}

		if hasAdditional {
			return "object{additional}"
		}

		return "object{}"

	case "string", "integer", "number", "boolean":
		return propType

	default:
		if propType == "" {
			if _, hasProps := propMap["properties"]; hasProps {
				return "implicit-object"
			}
			if _, hasItems := propMap["items"]; hasItems {
				return "implicit-array"
			}
		}
		return fmt.Sprintf("type:%s", propType)
	}
}

func GetArrayItemStructure(arrayProp map[string]interface{}) string {
	if arrayProp == nil {
		return "unknown"
	}

	items, hasItems := arrayProp["items"]
	if !hasItems {
		return "unknown"
	}

	return GetPropertyStructure(items)
}

func ComparePropertyStructures(entityProp, requestProp interface{}, schemas map[string]interface{}) bool {
	entityResolved, entityErr := GetResolvedPropertyType(entityProp, schemas)
	requestResolved, requestErr := GetResolvedPropertyType(requestProp, schemas)

	if entityErr == "unknown" || requestErr == "unknown" || entityResolved == nil || requestResolved == nil {
		return false
	}

	return compareResolvedProperties(entityResolved, requestResolved, schemas)
}

func HasTopLevelStructuralMismatch(entityProp, requestProp interface{}, schemas map[string]interface{}) bool {
	if entityProp == nil || requestProp == nil {
		return false
	}

	entityResolved, err1 := GetResolvedPropertyType(entityProp, schemas)
	requestResolved, err2 := GetResolvedPropertyType(requestProp, schemas)

	if err1 == "unknown" || err2 == "unknown" {
		// If we can't resolve, assume there might be a mismatch
		return true
	}

	entityStructure := GetPropertyStructure(entityResolved)
	requestStructure := GetPropertyStructure(requestResolved)

	if entityStructure != requestStructure {
		return true
	}

	// For arrays, check if the item types are fundamentally different
	if entityStructure == "array" || strings.HasPrefix(entityStructure, "array[") {
		entityItems := GetArrayItemStructure(entityResolved)
		requestItems := GetArrayItemStructure(requestResolved)

		// Different item structures indicate a mismatch
		// e.g., array[object] vs array[string]
		if entityItems != requestItems {
			return true
		}
	}

	// For objects, check if one is a ref and the other isn't
	if strings.HasPrefix(entityStructure, "$ref:") && !strings.HasPrefix(requestStructure, "$ref:") {
		return true
	}
	if !strings.HasPrefix(entityStructure, "$ref:") && strings.HasPrefix(requestStructure, "$ref:") {
		return true
	}

	return false
}

func compareResolvedProperties(entityResolved, requestResolved map[string]interface{}, schemas map[string]interface{}) bool {
	entityType, _ := entityResolved["type"].(string)
	requestType, _ := requestResolved["type"].(string)

	if entityType != requestType {
		return false
	}

	// For objects, recursively compare properties
	if entityType == "object" {
		entityProps, _ := entityResolved["properties"].(map[string]interface{})
		requestProps, _ := requestResolved["properties"].(map[string]interface{})

		// Check if request properties exist in entity (one-way check)
		for propName, requestSubProp := range requestProps {
			if entitySubProp, exists := entityProps[propName]; exists {
				if !ComparePropertyStructures(entitySubProp, requestSubProp, schemas) {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}

	// For arrays, compare item types
	if entityType == "array" {
		entityItems := entityResolved["items"]
		requestItems := requestResolved["items"]
		return ComparePropertyStructures(entityItems, requestItems, schemas)
	}

	return true
}
