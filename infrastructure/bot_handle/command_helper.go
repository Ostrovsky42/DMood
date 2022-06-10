package bot_handle

import (
	lib "DMood/library"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

func (b *moodBot) SwitchCommand(command string, chatId int64, update *tgbotapi.Update) stateFn {

	switch command {
	case "start":
		b.SendTextMessage(lib.HelloText+lib.HelloTextRUS, chatId)
		return b.HandleMessage

	case "reg":
		e := b.Storage.CreateUser(update.Message)
		if e != nil {
			log.Error(e)
			b.SendTextMessage("not registered", chatId)
		} else {
			b.SetTimeKeyboard("select notification time", "open", chatId)
			return b.HandleSetNotificationTime
		}
		return b.HandleMessage

	case "help":
		b.SendTextMessage(lib.HelpText, chatId)
		return b.HandleMessage

	case "set_rating":
		b.SetRatingKeyboard(lib.RatingText, "open", chatId)
		return b.HandleRating

	case "change_rating":
		b.SetRatingKeyboard(lib.RatingText, "open", chatId)
		return b.HandleChangeRating

	case "get_rating":
		return b.HandleGetRating(update)

	case "change_notification":
		b.SetTimeKeyboard("set new notification time","open",chatId)
		return b.HandleSetNotificationTime(update)
	}



	b.SendTextMessage("What is the command?", chatId)
	return b.HandleMessage
}
