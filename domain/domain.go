package domain

import "time"

type Mood struct {
	UserId      string `header:"user_id"` 
	Date        time.Time `header:"date"`
	MoodRating  string `header:"mood_rating"`
	Description string `header:"description"`
	DayIdea     string `header:"day_idea"`
}

type User struct {
	UserId string 	`header:"user_id"`
	UserName string	`header:"user_name"`
}