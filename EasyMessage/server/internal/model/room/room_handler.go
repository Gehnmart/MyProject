package room

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	model *RoomModel
}

func InitHandler(db *sql.DB) *Handler {
	return &Handler{
		&RoomModel{
			db,
		},
	}
}

type CreateRoomOnUserReq struct {
	Name         string `json:"name"`
	SecondUserId int    `json:"second_user_id"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var room_req CreateRoomOnUserReq

	err := c.ShouldBindJSON(&room_req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	first_user_id, _ := c.Get("userId")

	room_res, err := h.model.CreateRoom(c, CreateRoomReq{
		Name:         room_req.Name,
		FirstUserId:  first_user_id.(int),
		SecondUserId: room_req.SecondUserId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, room_res)
}

func (h *Handler) GetRoomByUserId(c *gin.Context) {
	user_id, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token is bad"})
		return
	}
	room_res, err := h.model.GetRoomByUserId(c, user_id.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, room_res)
}

func (h *Handler) GetRoomById(c *gin.Context) {
	room_id, _ := strconv.Atoi(c.Param("room_id"))
	room_res, err := h.model.GetRoomById(c, room_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, room_res.Id)
}
