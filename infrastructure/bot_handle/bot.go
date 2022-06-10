package bot_handle

import (
	"DMood/infrastructure/bot_handle/keyboard"
	"DMood/infrastructure/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

type stateFn func(*tgbotapi.Update) stateFn

type moodBot struct {
	ChatID     int
	State      stateFn
	Dispatcher Dispatcher
	API        tgbotapi.BotAPI
	UpdateCh   chan tgbotapi.Update
	Storage    storage.DMoodStorage
	KBord      keyboard.KBoard
	GraphicPath string
}

func New(chatId int, storage storage.DMoodStorage, api tgbotapi.BotAPI, KBord keyboard.KBoard, graphicPath string) *moodBot {
	bot := moodBot{
		ChatID:   chatId,
		API:      api,
		Storage:  storage,
		UpdateCh: make(chan tgbotapi.Update, 100),
		KBord:    KBord,
		GraphicPath: graphicPath,
	}
	bot.State = bot.HandleMessage
	return &bot
}

func (b *moodBot) Update() {
	for update := range b.UpdateCh {
		b.State = b.State(&update)
	}
}

func (b *moodBot) HandleMessage(update *tgbotapi.Update) stateFn {
	if update.Message != nil {
		
		if command:=update.Message.Command(); command!=""{
			return b.SwitchCommand(command, update.Message.Chat.ID, update)
		}		
		//b.API.Send(tgbotapi.Animation{FileName: })
		b.SetCommandKeyboard("мя","open", update.Message.Chat.ID)

		return b.HandleMessage
	} else {
		if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if resp, err := b.API.AnswerCallbackQuery(callback); err != nil {
				log.Print(resp)
			}

			return b.SwitchCommand(update.CallbackQuery.Data, update.CallbackQuery.Message.Chat.ID, update)
		}
	}
	return b.HandleMessage
}

func (b *moodBot) SendTextMessage(text string, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.API.Send(msg)
	if err != nil {
		log.Error(err)
	}
	log.Info(msg)
}

func (b *moodBot) SetRatingKeyboard(text string, io string, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, text)
	switch io {
	case "open":
		msg.ReplyMarkup = b.KBord.RatingKeyboard
	case "close":
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	if _, err := b.API.Send(msg); err != nil {
		log.Error(err)
	}

}

func (b *moodBot) SetCommandKeyboard(text string, io string, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, text)
	switch io {
	case "open":
		msg.ReplyMarkup = b.KBord.CommandKeyboard
	case "close":
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	if _, err := b.API.Send(msg); err != nil {
		log.Error(err)
	}

}

func (b *moodBot) SetTimeKeyboard(text string, io string, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, text)
	switch io {
	case "open":
		msg.ReplyMarkup = b.KBord.TimeKeyboard
	case "close":
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	if _, err := b.API.Send(msg); err != nil {
		log.Error(err)
	}

}