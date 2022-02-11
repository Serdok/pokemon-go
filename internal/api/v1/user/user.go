package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Uid       string `json:"uid" firestore:"uid" mapstructure:"uid"`
	FirstName string `json:"first_name" firestore:"first_name" mapstructure:"first_name"`
	LastName  string `json:"last_name" firestore:"last_name" mapstructure:"last_name"`
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": User{
			Uid:       "uid-test",
			FirstName: "Anass",
			LastName:  "Lahnin",
		},
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": User{
			Uid:       "uid-test",
			FirstName: "Anass",
			LastName:  "Lahnin",
		},
	})
}
