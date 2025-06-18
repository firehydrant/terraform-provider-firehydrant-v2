package common

import (
	"strings"
)

// ResolveRef follows a $ref to get the actual schema
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
	}

	return propMap, "inline"
}

// ComparePropertyStructures does a deep comparison of two property structures
func ComparePropertyStructures(entityProp, requestProp interface{}, schemas map[string]interface{}) bool {
	entityResolved, _ := GetResolvedPropertyType(entityProp, schemas)
	requestResolved, _ := GetResolvedPropertyType(requestProp, schemas)

	if entityResolved == nil || requestResolved == nil {
		return false
	}

	// Compare types
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
