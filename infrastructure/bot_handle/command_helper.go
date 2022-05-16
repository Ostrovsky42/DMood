package bot_handle

import (
	lib "DMood/libiary"
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
			b.SetCommandKeyboard("Done", "open", chatId)
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
		b.UpdateCh <- *update
		return b.HandleGetRating

	case "fuck":
		b.UpdateCh <- *update
		return b.HandleFuck
	}

	b.SendTextMessage("What is the command?", chatId)
	return b.HandleMessage
}
