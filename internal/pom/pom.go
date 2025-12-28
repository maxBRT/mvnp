package pom

import (
	"github.com/beevik/etree"
)

func ParsePOM(pomPath string) *etree.Document {
	doc := etree.NewDocument()
	doc.ReadFromFile(pomPath)
	return doc
}

func SetJavaVersion(root *etree.Element, version string) {
	props := root.SelectElement("properties")
	if props == nil {
		props = root.CreateElement("properties")
	}
	release := props.FindElement("maven.compiler.release")
	if release == nil {
		release = props.CreateElement("maven.compiler.release")
		release.SetText(version)
	} else {
		release.SetText(version)
	}
	source := props.FindElement("maven.compiler.source")
	if source == nil {
		source = props.CreateElement("maven.compiler.source")
		source.SetText(version)
	} else {
		source.SetText(version)
	}
	target := props.FindElement("maven.compiler.target")
	if target == nil {
		target = props.CreateElement("maven.compiler.target")
		target.SetText(version)
	} else {
		target.SetText(version)
	}
}
