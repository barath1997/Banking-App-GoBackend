package main

import (
	"github.com/gin-gonic/gin"
	"google.com/banking-app/controller"
	"google.com/banking-app/database"
	// "google.com/banking-app/migrations"
)

func main() {
	database.DbConnection()

	g := gin.Default()

	var ctlr *controller.Controller

	group := g.Group("/v1")
	{
		group.POST("/login", ctlr.LoginController)
		group.POST("/register", ctlr.RegisterController)
		group.GET("/get-user/:id", ctlr.GetUserController)
		group.POST("/transaction", ctlr.TransactionController)
		group.GET("/get-transactions/:id", ctlr.GetMyTransactionsController)
	}

	g.Run("localhost:8080")

}
