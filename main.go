package main

import (
	"log"
	"os"

	"github.com/AbdulAleem-git/abdkart/database"
	"github.com/abdulaleem-git/abdkart/controllers"
	"github.com/abdulaleem-git/abdkart/database"
	"github.com/abdulaleem-git/abdkart/middleware"
	"github.com/abdulaleem-git/abdkart/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.Get("/addtocart", app.AddToCart())
	router.Get("/removeitem", app.RemoveItem())
	router.Get("/cartcheckout", app.BuyFromCart())
	router.Get("/instantbuy", app.InstantBuy())
	log.Fatal(router.Run(":" + port))

}
