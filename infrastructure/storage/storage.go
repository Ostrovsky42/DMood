package storage

import (
	"DMood/domain"
	"DMood/localservices"
	"database/sql"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lensesio/tableprinter"
)

type DMoodStorage interface {
	GetDayRating(user *tgbotapi.User) ([]domain.Mood, error)
	ChangeDayRating(user_id int, rating string) error
	SaveDayRating(userId int, rating string) error
	SaveDayDescription(message *tgbotapi.Message) error
	SaveDayIdea(message *tgbotapi.Message) error
	CreateUser(message *tgbotapi.Message) error
}

type moodStorage struct {
	printer *tableprinter.Printer
	db      sql.DB
}

func NewStorage(db sql.DB) *moodStorage {
	return &moodStorage{db: db, printer: localservices.NewPrinter()}
}

func (m *moodStorage) CreateUser(message *tgbotapi.Message) error {
	q := "insert into \"users\" (user_id,user_name) values($1,$2)"
	if _, err := m.db.Exec(q, message.From.ID, message.From.UserName); err != nil {
		return err
	}
	return nil
}

var date string = time.Now().Format("2006-01-02")

func (m *moodStorage) SaveDayRating(userId int, rating string) error {
	q := "insert into \"mood\" (user_id, date, mood_rating , description, day_idea) values($1,$2,$3,$4,$5)"
	if _, err := m.db.Exec(q, userId, date, rating, "", "no idea"); err != nil {
		return err
	}
	return nil
}

func (m *moodStorage) ChangeDayRating(user_id int, rating string) error {
	q := "UPDATE mood  SET mood_rating =$1 WHERE date=$2 AND user_id=$3 "
	if _, err := m.db.Exec(q, rating, date, user_id); err != nil {
		return err
	}
	return nil
}

func (m *moodStorage) SaveDayDescription(message *tgbotapi.Message) error {
	q := "UPDATE mood  SET description =$1 WHERE date=$2 AND user_id=$3 "
	if _, err := m.db.Exec(q, message.Text, date, message.From.ID); err != nil {
		return err
	}
	return nil
}

func (m *moodStorage) SaveDayIdea(message *tgbotapi.Message) error {
	q := "UPDATE mood  SET day_idea =$1 WHERE date=$2 AND user_id=$3 "
	if _, err := m.db.Exec(q, message.Text, date, message.From.ID); err != nil {
		return err
	}
	return nil
}

func (m *moodStorage) GetDayRating(user *tgbotapi.User) ([]domain.Mood, error) {
	q := "select * from mood where user_id=$1"
	rows, err := m.db.Query(q, user.ID)
	if err != nil {
		return []domain.Mood{}, err
	}
	mood, mo := make([]domain.Mood, 0, 100), domain.Mood{}

	for rows.Next() {
		err := rows.Scan(&mo.UserId, &mo.Date, &mo.MoodRating, &mo.Description, &mo.DayIdea)
		mood = append(mood, mo)
		if err != nil {
			fmt.Println(err)
		}
	}
	defer rows.Close()
	return mood, nil
}
