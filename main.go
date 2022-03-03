package main

import (
	"github.com/shenoy-anurag/go-rest-example/routes"

	"github.com/shenoy-anurag/go-rest-example/configs"
)

func main() {
	// load environment variables
	configs.LoadEnvVariables()

	router := routes.SetupRouter()

	router.Run() // listen and serve on 0.0.0.0:8080
}
