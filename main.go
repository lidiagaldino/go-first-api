package main

import (
	"github.com/lidiagaldino/go-first-api/config"
	"github.com/lidiagaldino/go-first-api/router"
)

var(
    logger config.Logger
)

func main() {
    logger = *config.GetLogger("main")
    err := config.Init()
    if err != nil {
        logger.Errorf("config initialization error: %v", err)
        return
    }

   router.Initialize()
}