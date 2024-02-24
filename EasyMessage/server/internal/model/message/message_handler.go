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

	var message_req CreateMessageReq
	user_id, _ := c.Get("userId")
	room_id, _ := strconv.Atoi(c.Param("room_id"))
	message_req.UserId = user_id.(int)
	message_req.RoomId = room_id
	message_req.Content = content.Content

	message_res, err := h.model.CreateMessage(c, message_req)
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
