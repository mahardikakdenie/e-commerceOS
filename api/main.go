package main

import (
	"api/conn"
	dotenv "api/dot_env"
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// ========================================== //
	router := gin.Default()
	// ========================================== //
	// ===== 1. Load .env file in GoDotEnv =====//
	DB_HOST := dotenv.GoDotEnvVariable("DATABASE_HOST")
	DB_PORT := dotenv.GoDotEnvVariable("DATABASE_PORT")
	DB_USER := dotenv.GoDotEnvVariable("DATABASE_USER")
	DB_PASSWORD := dotenv.GoDotEnvVariable("DATABASE_PASSWORD")
	DB_SSL_MODE := dotenv.GoDotEnvVariable("DB_SSL_MODE")
	DB_NAME := dotenv.GoDotEnvVariable("DATABASE_NAME")
	// ==========================================//

	// ===== 2. Connect to PostgreSQL =====//
	conn.ConnectionPgsqlDB(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSL_MODE)
	// ===================================== //

	// ==== 3. Create a router with the default middleware ==== //

	routes.Router(conn.DBPgsql, router)

	// ========================================================//

	router.Run()
	// ==========================================//
}
