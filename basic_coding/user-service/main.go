package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	user_datasource "user-service/drivers/sqllite"
	user_repository "user-service/repositories"
	user_usecase "user-service/usecases"
	user_handler "user-service/handlers"

)

func main() {
	db, err := sql.Open("sqlite3", "file:./db/generator/user_data.db")

	if err != nil {
		log.Panicln("Failed to Initialized sqllite DB:", err)
	}

	db.SetMaxOpenConns(1)

	userDb := user_datasource.NewSqlLiteDatasource(db)

	
	userRepo := user_repository.NewUserRepository(userDb)

	userUsecase := user_usecase.NewUserUsecase(userRepo)

	userHandler := user_handler.UserNewHandler(userUsecase)


	router := gin.Default()

	v1 := router.Group("api/v1")

	v1.Use(cors.Default())
	{
		v1.POST("/user", userHandler.RegisterUser)

	}


	router.Run(fmt.Sprintf(":8084"))

}