package bot_handle

import (
	"DMood/domain"
	"DMood/local_service"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func (b *moodBot) HandleGetRating(update *tgbotapi.Update) stateFn {
	mood, err := b.Storage.GetDayRating(update.Message.From)
	if err != nil {
		b.SendTextMessage(err.Error(), update.Message.Chat.ID)
	}
	err=local_service.Getpng(mood,b.GraphicPath)
	if err!=nil{
	b.SendTextMessage(err.Error(), update.Message.Chat.ID)//todo log it
	return b.HandleMessage
	}
	if err=SendRatingPng(b);err!=nil{
	b.SendTextMessage(err.Error(), update.Message.Chat.ID)//todo log it
	return b.HandleMessage
	}

	return b.HandleMessage
}


func PrintTable( mood []domain.Mood,b *moodBot,update *tgbotapi.Update ){
	table := ""
	for _, row := range mood {
		table = fmt.Sprintf("|%s|Rating:%s|%s|%s|\n",
			row.Date.Format("2006-01-02"), row.MoodRating, row.Description, row.DayIdea)
		b.SendTextMessage(table, update.Message.Chat.ID)
	}
}

func SendRatingPng(b *moodBot) error {
photoBytes, err := ioutil.ReadFile(b.GraphicPath)
if err != nil {
	log.Error(err)
}
file := tgbotapi.FileBytes{
	Name: "myMood.png",
	Bytes: photoBytes,
}

message, err := b.API.Send(tgbotapi.NewDocumentUpload(int64(b.ChatID), file))
if err!=nil{
	log.Error(err,  message)
return err
}
return nil
}