package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/AbdulAleem-git/abdkart/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddress() gin.HandlerFunc      {}
func EditAddress() gin.HandlerFunc     {}
func EditWorkAddress() gin.HandlerFunc {}
func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Query("id")

		if userId == "" {
			c.Header("Content-type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)
		user_id, err := primitive.ObjectIDFromHex(userId)

		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		filter := bson.D{primitive.E{Key: "_id", Value: user_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)

		if err != nil {
			c.IndentedJSON(404, "Wrong Command")
			return
		}
		defer cancel()
		ctx.Done()

		c.IndentedJSON(200, "Successfully Deleted")

	}
}
