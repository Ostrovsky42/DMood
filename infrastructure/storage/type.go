package storage

import (
	"DMood/domain"
 	"time"
)
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

func (u *User)ToDomain()domain.User{
return domain.User{
	 UserId: 	u.UserId,
	 UserName: 	u.UserName,
	}
}

func (m *Mood)ToDomain()domain.Mood{
	return domain.Mood{
		 UserId: 		m.UserId,
		 Date: 			m.Date,
		 MoodRating: 	m.MoodRating,
		 Description: 	m.Description,
		 DayIdea: 		m.DayIdea,
		}
	}