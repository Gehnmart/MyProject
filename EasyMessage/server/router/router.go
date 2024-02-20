package router

import (
	"MyProject/EasyMessage/server/internal/user"
	"MyProject/EasyMessage/server/internal/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(user_handler *user.Handler, ws_handler *ws.Handler) {
	r = gin.Default()

	r.POST("/signup", user_handler.CreateUser)
	r.POST("/login", user_handler.LoginUser)
	r.GET("/logout", user_handler.LogoutUser)

	r.POST("/ws/createRoom", ws_handler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", ws_handler.JoinRoom)
	r.GET("/ws/getRooms", ws_handler.GetRooms)
	r.GET("/ws/getClients/:roomId", ws_handler.GetClients)
}

func Start(addres string) error {
	return r.Run(addres)
}
