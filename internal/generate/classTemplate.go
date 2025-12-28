// Package generate provides functions for generating Java files like classes, interfaces, enums, and records
package generate

import (
	"bytes"
	"text/template"
)

// Represents a Java class file template
type ClassData struct {
	Package   string
	ClassName string
	Type      string
}

// Constructor for ClassData
// Take in a string but make sure its a valid class type
func NewClassData(packageName, className, typeName string) *ClassData {
	parsedType := ParseType(typeName)
	return &ClassData{
		Package:   packageName,
		ClassName: className,
		Type:      parsedType.String(),
	}
}

// Basic template for generating files
const classTemplateStr = `package {{.Package}};

public {{.Type}} {{.ClassName}} {}
`

// ClassTemplate generates a class template
func ClassTemplate(data ClassData) (string, error) {
	tmpl, err := template.New("class").Parse(classTemplateStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Type Enum
type Type int

const (
	Class Type = iota
	Interface
	AbstractClass
	Enum
	Record
)

// String returns the string representation of the Type
func (t Type) String() string {
	switch t {
	case Class:
		return "class"
	case Interface:
		return "interface"
	case AbstractClass:
		return "abstract class"
	case Enum:
		return "enum"
	case Record:
		return "record"
	default:
		return ""
	}
}

// ParseType parses a string into a Type
func ParseType(s string) Type {
	switch s {
	case "Class":
		return Class
	case "Interface":
		return Interface
	case "Abstract Class":
		return AbstractClass
	case "Enum":
		return Enum
	case "Record":
		return Record
	default:
		return Class // default fallback
	}
}
