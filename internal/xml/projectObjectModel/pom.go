package projectobjectmodel

import (
	"encoding/xml"
	"os"
)

type ProjectObjecModel struct {
	XMLName            xml.Name     `xml:"project"`
	GroupId            string       `xml:"groupId"`
	ArtifactId         string       `xml:"artifactId"`
	Version            string       `xml:"version"`
	JavaVersionSource  string       `xml:"properties>maven.compiler.source"`
	JavaVersionRelease string       `xml:"properties>maven.compiler.release"`
	Dependencies       []Dependency `xml:"dependencies>dependency"`
	Build              Build        `xml:"build"`
}

type Dependency struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
	Scope      string `xml:"scope"`
}

type Build struct {
	Plugins []Plugin `xml:"pluginManagement>plugins>plugin"`
}

type Plugin struct {
	GroupId    string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
}

func UnmarshalPOM(filePath string) (*ProjectObjecModel, error) {
	pom := &ProjectObjecModel{}
	f, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(f, pom)
	if err != nil {
		return nil, err
	}

	return pom, nil
}
