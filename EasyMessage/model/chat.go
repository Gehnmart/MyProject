package model

type Chat struct {
	ID     int    	`json:"id"`
	Name   string 	`json:"name"`
	UserId []int	`json:"user_id"`
}