package main

import (
	dotenv "api/dot_env"

	"github.com/gin-gonic/gin"
)

func main() {
	dotenv.GoDotEnvVariable("DOT_ENV_VARIABLE")
	router := gin.Default()

	router.Run()
}
