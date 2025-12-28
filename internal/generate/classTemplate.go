package generate

import (
	"bytes"
	"text/template"
)

type ClassData struct {
	Package   string
	ClassName string
	Type      string
}

func NewClassData(packageName, className string, typeName Type) *ClassData {
	return &ClassData{
		Package:   packageName,
		ClassName: className,
		Type:      typeName.String(),
	}
}

const classTemplateStr = `package {{.Package}};

public {{.Type}} {{.ClassName}} {}
`

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

type Type int

const (
	Class Type = iota
	Interface
	AbstractClass
	Enum
	Record
)

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
