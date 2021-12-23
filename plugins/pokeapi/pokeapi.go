package main

import (
	"context"

	"github.com/merico-dev/lake/logger" // A pseudo type for Plugin Interface implementation
	"github.com/merico-dev/lake/plugins/core"
	"github.com/merico-dev/lake/plugins/pokeapi/api"
)

type PokeAPI string

func (plugin PokeAPI) Description() string {
	return "To collect data from PokeAPI"
}

func (plugin PokeAPI) Execute(options map[string]interface{}, progress chan<- float32, ctx context.Context) error {
	logger.Print("Starting PokeAPI execution...")
	return nil
}

func (plugin PokeAPI) RootPkgPath() string {
	return "github.com/merico-dev/lake/plugins/pokeapi"
}

func (plugin PokeAPI) ApiResources() map[string]map[string]core.ApiResourceHandler {
	return map[string]map[string]core.ApiResourceHandler{
		"test": {
			"GET": api.TestConnection,
		},
	}
}

// Export a variable named PluginEntry for Framework to search and load
var PluginEntry PokeAPI //nolint
