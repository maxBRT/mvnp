package projectobjectmodel

import (
	"encoding/xml"
	"fmt"
)

type Plugin struct {
	GroupId       string        `xml:"groupId,omitempty"`
	ArtifactId    string        `xml:"artifactId,omitempty"`
	Version       string        `xml:"version,omitempty"`
	Configuration *PluginConfig `xml:"configuration,omitempty"`
}

type PluginConfig struct {
	ConfigTags []ConfigTag
}

type ConfigTag struct {
	Name  string
	Value string
}

func (pom *ProjectObjecModel) AddPlugin(plugin Plugin) error {
	// Check if plugin already exists
	for _, p := range pom.Build.Plugins {
		if p.GroupId == plugin.GroupId && p.ArtifactId == plugin.ArtifactId {
			return fmt.Errorf("plugin already exists")
		}
	}

	// Validate configuration tags
	if plugin.Configuration != nil {
		for _, c := range plugin.Configuration.ConfigTags {
			if c.Name == "" {
				return fmt.Errorf("plugin configuration name cannot be empty")
			}
		}
	}

	pom.Build.Plugins = append(pom.Build.Plugins, plugin)

	return nil
}

// MarshalXML implements custom XML marshaling for PluginConfig
// This nests each ConfigTag as a child element inside <configuration>
func (pc PluginConfig) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Start the configuration element
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Encode each ConfigTag as a nested element
	for _, tag := range pc.ConfigTags {
		if err := tag.MarshalXML(e, xml.StartElement{}); err != nil {
			return err
		}
	}

	// Close the configuration element
	return e.EncodeToken(start.End())
}

func (c ConfigTag) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Use the Name field as the actual XML tag name
	start.Name = xml.Name{Local: c.Name}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	if c.Value != "" {
		// Encode the value as text content
		if err := e.EncodeToken(xml.CharData(c.Value)); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}
