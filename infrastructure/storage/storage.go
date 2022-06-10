package storage

import (
	"DMood/domain"
	"database/sql"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type DMoodStorage interface {
	GetDayRating(user *tgbotapi.User) ([]domain.Mood, error)
	ChangeDayRating(user_id int, rating string) error
	SaveDayRating(userId int, rating string) error
	SaveDayDescription(message *tgbotapi.Message) error
	SaveDayIdea(message *tgbotapi.Message) error
	CreateUser(message *tgbotapi.Message) error
	GetUser(userId int) error
	GetUsersByNotificationTime(hour int)([]domain.User, error)
	SaveNotificationTime(userId int, hour string) error
}

type moodStorage struct {
	db      sql.DB
}

func NewStorage(db sql.DB) *moodStorage {
	return &moodStorage{db: db}
}

func (m *moodStorage) CreateUser(message *tgbotapi.Message) error {
	q := "insert into \"users\" (user_id,user_name) values($1,$2)"
	if _, err := m.db.Exec(q, message.From.ID, message.From.UserName); err != nil {
		return err
	}
	return nil
}

func (m *moodStorage) GetUser(userId int)  error{
//todo get what!?
	return nil
}

func (m *moodStorage) GetUsersByNotificationTime(hour int)([]domain.User, error){
	q:="SELECT user_id,user_name FROM users WHERE request_time=$1 AND enabled_statistic=true"
	rows,err:=m.db.Query(q,hour)
	if err != nil {
		return nil, err
	}

	users, user := make([]domain.User, 0, 100), domain.User{}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.UserId, &user.UserName)
		users = append(users, user)
		if err != nil {
			fmt.Println(err)
		}
	}
	return users, nil
}

func (m *moodStorage) SaveNotificationTime(userId int, hour string) error {
	q := "UPDATE users  SET request_time =$1 WHERE  user_id=$2 "
	if _, err := m.db.Exec(q, hour, userId); err != nil {
		return err
	}
	return nil
}

func (m *moodStorage) SaveDayRating(userId int, rating string) error {
	q := "insert into \"mood\" (user_id, date, mood_rating , description, day_idea) values($1,$2,$3,$4,$5)"
	if _, err := m.db.Exec(q, userId, time.Now().Format("2006-01-02"), rating, "", "no idea"); err != nil {
		return err
	}
	return nil
}

func (m *moodStorage) ChangeDayRating(user_id int, rating string) error {
	q := "UPDATE mood  SET mood_rating =$1 WHERE date=$2 AND user_id=$3 "
	if _, err := m.db.Exec(q, rating, time.Now().Format("2006-01-02"), user_id); err != nil {
		return err
	}
	return nil
}

func (m *moodStorage) SaveDayDescription(message *tgbotapi.Message) error {
	q := "UPDATE mood  SET description =$1 WHERE date=$2 AND user_id=$3 "
	if _, err := m.db.Exec(q, message.Text, time.Now().Format("2006-01-02"), message.From.ID); err != nil {
		return err
	}
	return nil
}



func (m *moodStorage) SaveDayIdea(message *tgbotapi.Message) error {
	q := "UPDATE mood  SET day_idea =$1 WHERE date=$2 AND user_id=$3 "
	if _, err := m.db.Exec(q, message.Text, time.Now().Format("2006-01-02"), message.From.ID); err != nil {
		return err
	}
	return nil
}

func (m *moodStorage) GetDayRating(user *tgbotapi.User) ([]domain.Mood, error) {
	q := "SELECT * FROM mood WHERE user_id=$1 ORDER BY date ASC "
	rows, err := m.db.Query(q, user.ID)
	if err != nil {
		return []domain.Mood{}, err
	}
	moods, mood := make([]domain.Mood, 0, 100), domain.Mood{}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&mood.UserId, &mood.Date, &mood.MoodRating, &mood.Description, &mood.DayIdea)
		moods = append(moods, mood)
		if err != nil {
			fmt.Println(err)
		}
	}
	return moods, nil
}
