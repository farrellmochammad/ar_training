package main

import (
  "fmt"
  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "database/sql"
  "log"
  _ "github.com/mattn/go-sqlite3"
  supplier_usecase "supplier-service/usecases"
  supplier_datasource "supplier-service/drivers/sqllite"
  supplier_repository "supplier-service/repositories"
  supplier_handler "supplier-service/handler"
  supplier_middlewares "supplier-service/middlewares"
)

func main() {

	db, err := sql.Open("sqlite3", "file:./db/generator/supplier_data.db")

	if err != nil {
		log.Panicln("Failed to Initialized sqllite DB:", err)
	}

	db.SetMaxOpenConns(1)

	supplierDb := supplier_datasource.NewSqlLiteDatasource(db)

	supplierRepo := supplier_repository.NewSupplierRepository(supplierDb)

	supplierUsecase := supplier_usecase.NewSupplierUsecase(supplierRepo)

	supplierHandler := supplier_handler.SupplierNewHandler(supplierUsecase)

	supplierMiddleware := supplier_middlewares.SupplierMiddlewareNewHandler(supplierUsecase)

	router := gin.Default()


	v1 := router.Group("api/v1")

	v1.Use(cors.Default())
	{
		v1.POST("/supplier", supplierHandler.RegisterSeller)

	}

	v1.Use(supplierMiddleware.AuthHandler(supplierUsecase))
	{
		v1.POST("/:id/store", supplierHandler.RegisterStore)
	}

	router.Run(fmt.Sprintf(":8083"))
}
