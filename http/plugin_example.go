package tcp

import (
	gohttp "net/http"

	"github.com/cyberark/secretless-broker/pkg/secretless/plugin/connector"
	"github.com/cyberark/secretless-broker/pkg/secretless/plugin/connector/http"
)

// PluginInfo is required in the shared object file
func PluginInfo() map[string]string {
	return map[string]string{
		"pluginAPIVersion": "0.1.0",
		"type":             "connector.http",
		"id":               "example", // plugin identifier used in configuration
		"description":      "useful description of http connector plugin",
	}
}

// GetHTTPPlugin is required in the shared object file
func GetHTTPPlugin() http.Plugin {
	return &Plugin{}
}

// Plugin definition

type Plugin struct{}

func (p *Plugin) NewConnector(_ connector.Resources) http.Connector {
	return Connect
}

// Connect is a standalone function that implements the http.Connector func
// signature
func Connect(
	request *gohttp.Request,
	credentialValuesByID connector.CredentialValuesByID,
) error {
	// inject credentials
	request.Header.Set("injected-auth", string(credentialValuesByID["auth"]))

	return nil
}
