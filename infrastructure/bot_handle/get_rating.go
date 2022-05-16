package bot_handle

import (
	"DMood/domain"
	"DMood/localservices"
	"fmt"
	"io/ioutil"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

func (b *moodBot) HandleGetRating(update *tgbotapi.Update) stateFn {
	mood, err := b.Storage.GetDayRating(update.Message.From)
	if err != nil {
		b.SendTextMessage(err.Error(), update.Message.Chat.ID)
	}
	localservices.Getpng(mood)
	hehehhe(b)
	
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

func hehehhe(b *moodBot) {

photoBytes, err := ioutil.ReadFile("C:/Users/Serj/go/src/DMood/localservices/myMood.png")
if err != nil {
	log.Error(err)
}
file := tgbotapi.FileBytes{
	Name: "myMood.png",
	Bytes: photoBytes,
}

message, err := b.API.Send(tgbotapi.NewDocumentUpload(int64(b.ChatID), file))
if err!=nil{
	log.Fatal(err,  message)

}

}