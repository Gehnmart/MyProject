package message

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	model *MessageModel
}

func InitHandler(db *sql.DB) *Handler {
	return &Handler{
		&MessageModel{
			db,
		},
	}
}

func (h *Handler) CreateMessage(c *gin.Context) {

	var content struct {
		Content string `json:"content"`
	}
	err := c.ShouldBindJSON(&content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	var messageReq CreateMessageReq
	userId, _ := c.Get("userId")
	room_id, _ := strconv.Atoi(c.Param("room_id"))
	messageReq.UserId = userId.(int)
	messageReq.RoomId = room_id
	messageReq.Content = content.Content

	message_res, err := h.model.CreateMessage(c, messageReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, message_res)
}

// func (h *Handler) GetMessagesByRoomId(c *gin.Context) {
// 	room_id, _ := strconv.Atoi(c.Param("room_id"))
// 	message_res, err := h.model.GetMessagesByRoomId(c, room_id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
// 	}
// 	c.JSON(http.StatusOK, message_res)
// }
