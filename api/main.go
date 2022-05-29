package main

import (
	"api/conn"
	dotenv "api/dot_env"
	"api/routes"
	"fmt"

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
	PORT := dotenv.GoDotEnvVariable("PORT")
	// ==========================================//

	fmt.Println("DB_HOST: ", DB_HOST)

	// ===== 2. Connect to PostgreSQL =====//
	conn.ConnectionPgsqlDB(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSL_MODE)
	// ===================================== //

	// ==== 3. Create a router with the default middleware ==== //

	routes.Router(conn.DBPgsql, router)

	// ========================================================//

	// ==== 4. Start the server ===== //

	if err := router.Run(fmt.Sprintf(":%s", PORT)); err != nil {
		panic(err)
	}
	//http.ListenAndServe(":"+PORT, nil)
	// http.ListenAndServe("127.0.0.1:8080", nil)
	// ==========================================//
}
