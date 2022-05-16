package bot_handle

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func (b *moodBot) HandleFuck(update *tgbotapi.Update) stateFn {
	count:=1
		for count<40 {
		b.SendTextMessage("👉🏿_____👌🏻", update.Message.Chat.ID)
		time.Sleep(time.Millisecond*5)		
		count++
		b.SendTextMessage("👉🏿__👌🏻",update.Message.Chat.ID)
		time.Sleep(time.Millisecond*5)
			count++
			b.SendTextMessage("👉🏿👌🏻",update.Message.Chat.ID)
		time.Sleep(time.Millisecond*5)
			count++
			b.SendTextMessage("👉🏿__👌🏻",update.Message.Chat.ID)
		time.Sleep(time.Millisecond*5)
			count++
	}
	id:=update.Message.MessageID+count
	for id!=update.Message.MessageID{
		b.API.DeleteMessage(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, id))
		id--
	}
	time.Sleep(time.Second)
			b.SendTextMessage("cum 💦💦💦",update.Message.Chat.ID)
	return b.HandleMessage
}