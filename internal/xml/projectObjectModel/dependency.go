package projectobjectmodel

import (
	"fmt"
)

type Dependency struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version,omitempty"`
	Type       string `xml:"type,omitempty"`
	Scope      string `xml:"scope"`
}

func (pom *ProjectObjecModel) AddDependency(dependency Dependency) error {
	// Check if dependency already exists
	for _, d := range pom.Dependencies {
		if d.GroupId == dependency.GroupId && d.ArtifactId == dependency.ArtifactId {
			return fmt.Errorf("dependency already exists")
		}
	}

	pom.Dependencies = append(pom.Dependencies, dependency)

	return nil
}
