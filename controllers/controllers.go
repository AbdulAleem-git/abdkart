package controllers

import "github.com/gin-gonic/gin"

func HashPasword(password string) string {

}

func verifyPassword(userPassword string, givenPassword string) (bool, string) {}

func Login() gin.HandlerFunc {}

func Signup() gin.HandlerFunc {}

func ProductViewerAdmin() gin.HandlerFunc {}

func SearchProduct() gin.HandlerFunc {}

func SearchProductByQuery() gin.HandlerFunc {}
