package bot_handle

import (
	"DMood/library"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func (b *moodBot) HandleSetNotificationTime(update *tgbotapi.Update) stateFn {
	var hour string
	if update.CallbackQuery != nil {
		hour = GetHour(update, &b.API)

		e := b.Storage.SaveNotificationTime(update.CallbackQuery.From.ID, hour)
		if e != nil {
			b.SendTextMessage(library.SomethingWentWrong,int64(update.CallbackQuery.From.ID))
			return b.HandleMessage
		}
		msg:=fmt.Sprintf("I will ask about your day in %s:00",hour)
		b.SetCommandKeyboard(msg,"open",int64(update.CallbackQuery.From.ID))
	}else{
		if IsHour(update.Message.Text){
			e := b.Storage.SaveNotificationTime(update.CallbackQuery.From.ID, hour)
			if e != nil {
				b.SendTextMessage(library.SomethingWentWrong,int64(update.CallbackQuery.From.ID))
				return b.HandleMessage
			}
			msg:=fmt.Sprintf("I will ask about your day in %s:00",hour)
			b.SetCommandKeyboard(msg,"open",int64(update.CallbackQuery.From.ID))
			return b.HandleMessage
		}
		return b.HandleSetNotificationTime
	}

	return b.HandleMessage
}


func GetHour(update *tgbotapi.Update, API *tgbotapi.BotAPI) string {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	resp, err := API.AnswerCallbackQuery(callback)
	if err != nil {
		log.Error(err)
	}
	callbackRating := update.CallbackQuery.Data
	log.Debug(resp)
	return callbackRating
}

func IsHour(text string) bool {
	hour, e := strconv.Atoi(text)
	if e != nil || hour >= 24 || hour < 0 {
		return false
	}
	return true
}