package bot_handle

import (
	"DMood/library"
	"strconv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

func GetRating(update *tgbotapi.Update, API *tgbotapi.BotAPI) string {
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	resp, err := API.AnswerCallbackQuery(callback)
	if err != nil {
		log.Error(err)
	}
	callbackRating := update.CallbackQuery.Data
	log.Debug(resp)
	return callbackRating
}

func IsSayNo(msg string)bool{
	if msg=="No" || msg=="no" {
	return true
	}
	return false
}

func IsRating(callbackRating string) bool {
	rating, e := strconv.Atoi(callbackRating)
	if e != nil || rating > 5 || rating < 0 {
		return false
	}
	return true
}

func (b *moodBot) HandleRating(update *tgbotapi.Update) stateFn {
	var callbackRating string
	if update.CallbackQuery != nil {
		callbackRating = GetRating(update, &b.API)
		if !IsRating(callbackRating) {
			b.SendTextMessage(library.IsNotRating,int64(update.CallbackQuery.From.ID))
			return b.HandleRating
		}
		e := b.Storage.SaveDayRating(update.CallbackQuery.From.ID, callbackRating)
		if e != nil {
			b.SendTextMessage(library.SomethingWentWrong,int64(update.CallbackQuery.From.ID))
			return b.HandleMessage
		}
		b.SendTextMessage("Send day description", int64(update.CallbackQuery.From.ID))
		return b.HandleDescription
	} else {
		if IsRating(update.Message.Text) {
			b.SendTextMessage(library.IsNotRating,int64(update.CallbackQuery.From.ID))
			return b.HandleRating
		}
		if IsSayNo(update.Message.Text){
			b.SendTextMessage("ok:(", update.Message.Chat.ID)
			return b.HandleMessage
		}
	}
	return b.HandleMessage
}

func (b *moodBot) HandleDescription(update *tgbotapi.Update) stateFn {
	if IsSayNo(update.Message.Text) {
		b.SendTextMessage("ok:(", update.Message.Chat.ID)
			return b.HandleMessage
	}
	b.Storage.SaveDayDescription(update.Message)
	b.SendTextMessage("Send day idea", update.Message.Chat.ID)
	return b.HandleDayIdea
}

func (b *moodBot) HandleDayIdea(update *tgbotapi.Update) stateFn {
	b.Storage.SaveDayIdea(update.Message)
	b.SendTextMessage("be happy:3", update.Message.Chat.ID)
	return b.HandleMessage
}
