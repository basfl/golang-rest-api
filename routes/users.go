package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/util"
	"github.com/gin-gonic/gin"
)

func singup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})

}

func login(context *gin.Context) {
	var user models.User
	//eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := util.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})

}
