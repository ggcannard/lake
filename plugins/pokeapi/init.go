package main

import (
	"github.com/merico-dev/lake/logger"
)

func (plugin PokeAPI) Init() {
	logger.Info("Init PokeAPI plugin", true)
}
