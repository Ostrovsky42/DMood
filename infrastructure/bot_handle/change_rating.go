package bot_handle

import (
	"DMood/library"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *moodBot) HandleChangeRating(update *tgbotapi.Update) stateFn {
	var callbackRating string = ""
	if update.CallbackQuery != nil {
		callbackRating = GetRating(update, &b.API)
		if !IsRating(callbackRating) {
			b.SendTextMessage(library.IsNotRating,int64(update.CallbackQuery.From.ID))
			return b.HandleChangeRating
		}
		e := b.Storage.ChangeDayRating(update.CallbackQuery.From.ID, callbackRating)
		if e != nil {
			b.SendTextMessage(library.SomethingWentWrong,int64(update.CallbackQuery.From.ID))
			return b.HandleMessage
		}
		b.SendTextMessage("Send day description", int64(update.CallbackQuery.From.ID))
		return b.HandleDescription
	} else {
		if IsRating(update.Message.Text) {
			b.SendTextMessage(library.IsNotRating,int64(update.CallbackQuery.From.ID))
			return b.HandleChangeRating
		}
		if strings.Contains(update.Message.Text, "No") || strings.Contains(update.Message.Text, "no") {
			b.SendTextMessage("ok:(", update.Message.Chat.ID)
			return b.HandleMessage
		}
	}
	return b.HandleMessage
}

