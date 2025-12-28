package pom

import (
	"encoding/json"
	"fmt"
	"github.com/beevik/etree"
	"net/http"
	"time"
)

type Plugin struct {
	*etree.Element
	groupId, artifactId, version string
}

type MavenResponse struct {
	Response struct {
		Docs []struct {
			LatestVersion string `json:"latestVersion"`
		} `json:"docs"`
	} `json:"response"`
}

func GetLatestVersion(groupID, artifactID string) (string, error) {
	// 1. Construct the URL
	url := fmt.Sprintf(
		"https://search.maven.org/solrsearch/select?q=g:%s+AND+a:%s&rows=1&wt=json",
		groupID, artifactID,
	)

	// 2. Create client with timeout (good practice for CLIs)
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("maven central returned status: %d", resp.StatusCode)
	}

	// 3. Decode JSON
	var result MavenResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// 4. Extract version
	if len(result.Response.Docs) == 0 {
		return "", fmt.Errorf("artifact not found")
	}

	return result.Response.Docs[0].LatestVersion, nil
}

func (p *Plugin) AddConfiguration(name, value string) {
	config := p.SelectElement("configuration")
	if config == nil {
		config = p.CreateElement("configuration")
	}
	if existing := config.FindElement(name); existing != nil {
		existing.SetText(value)
		return
	}
	config.CreateElement(name).SetText(value)
}

func AddPlugin(root *etree.Element, groupId, artifactId, version string) *Plugin {
	// Find the plugins element
	plugins := root.FindElement(".//build/pluginManagement/plugins")

	// If it doesn't exist, create it
	if plugins == nil {
		build := root.FindElement(".//build")
		if build == nil {
			build = root.CreateElement("build")
		}
		pManagement := build.FindElement(".//pluginManagement")
		if pManagement == nil {
			pManagement = build.CreateElement("pluginManagement")
		}
		plugins = pManagement.CreateElement("plugins")
	}

	// Add the new plugin
	plugin := plugins.CreateElement("plugin")
	plugin.CreateElement("groupId").SetText(groupId)
	plugin.CreateElement("artifactId").SetText(artifactId)
	plugin.CreateElement("version").SetText(version)

	return &Plugin{
		Element: plugin,
	}
}
