package router

import (
	"server/internal/model/message"
	"server/internal/model/room"
	"server/internal/model/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(user *user.Handler, room *room.Handler, message *message.Handler) *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", user.CreateUser)
		auth.POST("/login", user.LoginUser)
	}

	api := router.Group("/api", user.IdentityUser)
	{
		rooms := api.Group("/room")
		{
			rooms.POST("", room.CreateRoom)
			rooms.GET("/user", room.GetRoomByUserId)

			messages := rooms.Group("/:room_id/message", room.GetRoomById)
			{
				//messages.GET("", message.GetMessageByRoomId)
				messages.POST("", message.CreateMessage)
			}
		}
	}
	return router
}

func Start(r *gin.Engine, address string) error {
	return r.Run(address)
}
