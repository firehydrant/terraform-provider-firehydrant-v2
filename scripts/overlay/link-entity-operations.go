package main

import "strings"

func mapCrudToEntityOperation(crudType, entityName string) string {
	switch crudType {
	case "create":
		return entityName + "#create"
	case "read":
		return entityName + "#read"
	case "update":
		return entityName + "#update"
	case "delete":
		return entityName + "#delete"
	case "list":
		// For list operations, pluralize the entity name and use #read
		pluralEntityName := pluralizeEntityName(entityName)
		return pluralEntityName + "#read"
	default:
		return entityName + "#" + crudType
	}
}

// Simplified pluralization logic
func pluralizeEntityName(entityName string) string {
	// Remove "Entity" suffix
	baseName := strings.TrimSuffix(entityName, "Entity")

	// Simple pluralization
	if strings.HasSuffix(baseName, "y") && len(baseName) > 1 && !isVowel(baseName[len(baseName)-2]) {
		baseName = baseName[:len(baseName)-1] + "ies"
	} else if strings.HasSuffix(baseName, "s") ||
		strings.HasSuffix(baseName, "ss") ||
		strings.HasSuffix(baseName, "sh") ||
		strings.HasSuffix(baseName, "ch") ||
		strings.HasSuffix(baseName, "x") ||
		strings.HasSuffix(baseName, "z") {
		baseName = baseName + "es"
	} else {
		baseName = baseName + "s"
	}

	return baseName + "Entities"
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
