package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.RegisterEvent(int64(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registered"})

}

func cancelRegisteration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId
	err = event.CancelRegisteration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel register "})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "canceled"})

}
