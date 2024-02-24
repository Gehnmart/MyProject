package main

import (
	"log"
	"os"
	"server/internal/config"
	"server/internal/database"
	"server/internal/model/message"
	"server/internal/model/room"
	"server/internal/model/user"
	"server/router"
)

func main() {
	cfg := config.Init()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := database.New(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	user_handler := user.InitHandler(db)
	room_handler := room.InitHandler(db)
	message_handler := message.InitHandler(db)
	r := router.InitRouter(user_handler, room_handler, message_handler)
	err = router.Start(r, "localhost:4000")
	if err != nil {
		logger.Fatal(err)
	}
}
