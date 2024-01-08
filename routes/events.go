package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {
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
	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {
	// token := context.Request.Header.Get("Authorization")
	// if token == "" {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
	// 	return
	// }
	// userId, err := util.VerifyToken(token)
	// if err != nil {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
	// 	return
	// }
	event := models.Event{}
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}
	//	event.ID = 1
	context.GetInt64("userId")
	event.UserID = context.GetInt64("userId")
	err = event.Save()
	//fmt.Println("userId==", userId)
	if err != nil {
		fmt.Println("error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"event": event})

}

func updateEvents(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	userId := context.GetInt64("userId")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "no authorized to update event"})
		return
	}
	var upadtedEvent models.Event
	err = context.ShouldBindJSON(&upadtedEvent)
	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request"})
		return
	}
	upadtedEvent.ID = eventId
	err = upadtedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event updated successfully!"})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	userId := context.GetInt64("userId")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "no authorized to update event"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delet event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})

}
