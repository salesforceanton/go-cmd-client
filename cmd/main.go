package main

import (
	"fmt"

	"github.com/salesforceanton/go-cmd-client/internal/config"
	"github.com/salesforceanton/go-cmd-client/internal/logger"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		logger.LogError("Runtime", err)
		return
	}
	fmt.Println(cfg)
}
