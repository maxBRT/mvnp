package projectobjectmodel

import (
	"encoding/xml"
	"os"
)

type ProjectObjecModel struct {
	XMLName              xml.Name     `xml:"project"`
	ModelVersion         string       `xml:"modelVersion"`
	GroupId              string       `xml:"groupId"`
	ArtifactId           string       `xml:"artifactId"`
	Version              string       `xml:"version"`
	JavaVersionSource    string       `xml:"properties>maven.compiler.source"`
	JavaVersionRelease   string       `xml:"properties>maven.compiler.release"`
	DependencyManagement []Dependency `xml:"dependencyManagement>dependencies>dependency"`
	Dependencies         []Dependency `xml:"dependencies>dependency"`
	Build                Build        `xml:"build"`
}

type Build struct {
	Plugins []Plugin `xml:"pluginManagement>plugins>plugin"`
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

func (pom *ProjectObjecModel) MarshalPOM() ([]byte, error) {
	return xml.MarshalIndent(pom, "", "\t")
}
