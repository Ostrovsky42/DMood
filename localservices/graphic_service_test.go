package localservices

import (
	"DMood/domain"
	"testing"
	"time"
)



func TestGetpng(t *testing.T) {
	t.Run("PRINT", func(t *testing.T) {
		slice:= []domain.Mood{

			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,1,0,0,0,0,time.UTC),MoodRating:"2",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,2,0,0,0,0,time.UTC),MoodRating:"3",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,3,0,0,0,0,time.UTC),MoodRating:"3",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,4,0,0,0,0,time.UTC),MoodRating:"3",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,5,0,0,0,0,time.UTC),MoodRating:"3",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,6,0,0,0,0,time.UTC),MoodRating:"4",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,7,0,0,0,0,time.UTC),MoodRating:"1",Description: "",DayIdea: "хочу есть"},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,8,0,0,0,0,time.UTC),MoodRating:"2",Description: "",DayIdea: "хочу спать"},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,9,0,0,0,0,time.UTC),MoodRating:"1",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,10,0,0,0,0,time.UTC),MoodRating:"1",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,11,0,0,0,0,time.UTC),MoodRating:"5",Description: "",DayIdea: "Спорт сила"},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,12,0,0,0,0,time.UTC),MoodRating:"3",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,13,0,0,0,0,time.UTC),MoodRating:"3",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,14,0,0,0,0,time.UTC),MoodRating:"3",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,15,0,0,0,0,time.UTC),MoodRating:"4",Description: "",DayIdea: "Погодка супер"},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,16,0,0,0,0,time.UTC),MoodRating:"3",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,17,0,0,0,0,time.UTC),MoodRating:"2",Description: "",DayIdea: "Опять дожди("},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,18,0,0,0,0,time.UTC),MoodRating:"2",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,19,0,0,0,0,time.UTC),MoodRating:"2",Description: "",DayIdea: ""},
			domain.Mood{UserId: "uidi",Date: time.Date(2022,5,20,0,0,0,0,time.UTC),MoodRating:"2",Description: "",DayIdea: ""},

		}
	//	actual
		Getpng(slice)
	//	assert
	})
}