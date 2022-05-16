package bot_handle

import (
	"DMood/infrastructure/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"sync"
)

// Bot is the interface that must be implemented by your definition of
// the struct thus it represent each open session with a user on Telegram.
type IBot interface {
	// Update will be called upon receiving any update from Telegram.
	Update(*tgbotapi.Update)
}


type NewBotFn func(chatId int,s storage.DMoodStorage, api tgbotapi.BotAPI, UpdateCh tgbotapi.UpdatesChannel) *moodBot


// The Dispatcher passes the updates from the Telegram Bot API to the Bot instance
// associated with each chatID. When a new chat ID is found, the provided function
// of type NewBotFn will be called.
type Dispatcher struct {
	api        *tgbotapi.BotAPI
	sessionMap map[int]moodBot
	newBot     func(chatId int,s storage.DMoodStorage, api tgbotapi.BotAPI,UpdateCh tgbotapi.UpdatesChannel) *moodBot
	updates    *tgbotapi.UpdatesChannel
	mu         sync.Mutex
	storage    storage.DMoodStorage
}

// NewDispatcher returns a new instance of the Dispatcher object.
// Calls the Update function of the bot associated with each chat ID.
// If a new chat ID is found, newBotFn will be called first.
func NewDispatcher(token string,s storage.DMoodStorage, newBotFn NewBotFn) *Dispatcher {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	//api.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates,err := api.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Authorized on account %s", api.Self)
	d := &Dispatcher{
		api:        api,
		sessionMap: make(map[int]moodBot),
		newBot:     newBotFn,
		updates:    &updates,
		storage: s,
	}
	go d.listen()

	return d
}

// DelSession deletes the Bot instance, seen as a session, from the
// map with all of them.
func (d *Dispatcher) DelSession(chatID int) {
	d.mu.Lock()
	delete(d.sessionMap, chatID)
	d.mu.Unlock()
}

// AddSession allows to arbitrarily create a new Bot instance.
func (d *Dispatcher) AddSession(chatID int) {
	d.mu.Lock()
	if _, isIn := d.sessionMap[chatID]; !isIn {
//		d.sessionMap[chatID] = d.newBot(chatID)
	}
	d.mu.Unlock()
}

// Poll is a wrapper function for PollOptions.
func (d *Dispatcher) Poll() error {
	return d.PollOptions(true, tgbotapi.UpdateConfig{Timeout: 120})
}

// PollOptions starts the polling loop so that the dispatcher calls the function Update
// upon receiving any update from Telegram.
func (d *Dispatcher) PollOptions(dropPendingUpdates bool, opts tgbotapi.UpdateConfig) error {
	var (
		timeout      = opts.Timeout
		isFirstRun   = true
		lastUpdateID = -1
	)
	
	
	for {
		if isFirstRun {
			opts.Timeout = 0
		}
	
		opts.Offset = lastUpdateID + 1
		response, err := d.api.GetUpdates(opts)
		if err != nil {
			return err
		}

		if l := len(response); l > 0 {
			lastUpdateID = response[0].UpdateID
		}

		if isFirstRun {
			isFirstRun = false
			opts.Timeout = timeout
		}
	}
	return nil
}

func (d *Dispatcher) instance(chatID int) (moodBot, bool) {
	bot, ok := d.sessionMap[chatID]
	if !ok {
		newBot := d.newBot(chatID, d.storage, *d.api, *d.updates)
		d.mu.Lock()
		d.sessionMap[chatID] = *newBot
		d.mu.Unlock()
	return *newBot, false
	}
	return bot, true
}

func (d *Dispatcher) listen() {
for {
	for update := range *d.updates {
		var chatID int64
		switch {
		case update.Message != nil:
			chatID = update.Message.Chat.ID
		case update.EditedMessage != nil:
			chatID = update.EditedMessage.Chat.ID
		case update.ChannelPost != nil:
			chatID = update.ChannelPost.Chat.ID
		case update.EditedChannelPost != nil:
			chatID = update.EditedChannelPost.Chat.ID
		case update.InlineQuery != nil:
			chatID = int64(update.InlineQuery.From.ID)			
		case update.ChosenInlineResult != nil:
			chatID = int64(update.ChosenInlineResult.From.ID)
		case update.CallbackQuery != nil:
			chatID = update.CallbackQuery.Message.Chat.ID
		case update.ShippingQuery != nil:
			chatID = int64(update.ShippingQuery.From.ID)
		case update.PreCheckoutQuery != nil:
			chatID = int64(update.PreCheckoutQuery.From.ID)
		// case update.MyChatMember != nil:
		// 	chatID = update.MyChatMember.Chat.ID
		// case update.ChatMember != nil:
		// 	chatID = update.ChatMember.Chat.ID
		// case update.ChatJoinRequest != nil:
		// 	chatID = update.ChatJoinRequest.Chat.ID
		default:
			continue
		}

		bot,ok:= d.instance(int(chatID))
		if !ok{
			go bot.Update()			
		}

		bot.UpdateCh<-update
	}
}
}

