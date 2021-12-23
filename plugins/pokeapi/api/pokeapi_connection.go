package api

import (
	"github.com/merico-dev/lake/plugins/core"
)

/*
GET /plugins/pokeapi/test
*/
func TestConnection(input *core.ApiResourceInput) (*core.ApiResourceOutput, error) {
	// TODO: implement test connection
	return &core.ApiResourceOutput{Body: true}, nil
}
