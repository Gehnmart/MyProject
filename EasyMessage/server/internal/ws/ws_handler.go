package ws

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	Hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		Hub: hub,
	}
}

type CreateRoomRequest struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.Hub.Rooms[req.Id] = &Room{
		Id:      req.Id,
		Name:    req.Name,
		Clients: make(map[uint64]*Client),
	}

	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	roomId, _ := strconv.Atoi(c.Param("roomId"))
	clientId, _ := strconv.Atoi(c.Query("userId"))
	username := c.Query("username")

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		Id:       uint64(clientId),
		RoomId:   uint64(roomId),
		Username: username,
	}

	m := &Message{
		Content:  "A new user join in this room",
		RoomId:   uint64(roomId),
		Username: username,
	}

	h.Hub.Register <- cl

	h.Hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.Hub)
}

type RoomRes struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.Hub.Rooms {
		rooms = append(rooms, RoomRes{
			Id:   r.Id,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

type ClientRes struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientRes
	roomId, _ := strconv.Atoi(c.Param("roomId"))

	if _, ok := h.Hub.Rooms[uint64(roomId)]; ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, cl := range h.Hub.Rooms[uint64(roomId)].Clients {
		clients = append(clients, ClientRes{
			Id:       cl.Id,
			Username: cl.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}
