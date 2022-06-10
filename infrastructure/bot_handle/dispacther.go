package bot_handle

import (
	"DMood/config"
	"DMood/domain"
	"DMood/infrastructure/bot_handle/keyboard"
	"DMood/infrastructure/storage"
	"DMood/local_service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

type IBot interface {
	Update(*tgbotapi.Update)
}

type NewBotFn func(chatId int,s storage.DMoodStorage, api tgbotapi.BotAPI, kBoard keyboard.KBoard, graphicPath string) *moodBot


type Dispatcher struct {
	api        *tgbotapi.BotAPI
	sessionMap map[int]moodBot
	newBot     func(chatId int,s storage.DMoodStorage, api tgbotapi.BotAPI,kBoard keyboard.KBoard, graphicPath string) *moodBot
	updates    *tgbotapi.UpdatesChannel
	scheduler  *local_service.Scheduler
	storage    storage.DMoodStorage
	graphicPath string
	kBoard 	   keyboard.KBoard
	mu         sync.Mutex
}

func NewDispatcher(conf config.BotConfig,storage storage.DMoodStorage) *Dispatcher {
	api, err := tgbotapi.NewBotAPI(conf.TgToken)
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
		newBot:     New,
		updates:    &updates,
		scheduler: local_service.NewScheduler(),
		storage: storage,
		graphicPath:conf.PngPath,
		kBoard : keyboard.NewKeyboard(),
	}
	go d.listen()
	d.StartScheduler()
	return d
}

func (d *Dispatcher) DelSession(chatID int) {
	d.mu.Lock()
	delete(d.sessionMap, chatID)
	d.mu.Unlock()
}

func (d *Dispatcher) AddSession(chatID int) {
	d.mu.Lock()
	if _, isIn := d.sessionMap[chatID]; !isIn {
	}
	d.mu.Unlock()
}

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

func (d *Dispatcher) instance(chatID int,update tgbotapi.Update) (moodBot, bool) {
	bot, ok := d.sessionMap[chatID]
		if !ok {		
		//	user,err:=d.storage.GetUser(update.Message.From.ID)// todo get
		newBot := d.newBot(chatID, d.storage, *d.api, d.kBoard, d.graphicPath)
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
		d.updateBot(int(chatID),update)
	}
}
}

func (d *Dispatcher) StartScheduler()  {
log.Info("start Scheduler")
d.scheduler.Process()
for hour:=range d.scheduler.Time{
	users,err:=d.storage.GetUsersByNotificationTime(hour)
	if err!=nil{
		log.Error("GetUsersByNotificationTime:",err)
	}
	for _,user:=range users{
	go	func(user domain.User){
		userId,_:=strconv.Atoi(user.UserId)
		update:=prepareUpdateForNotification(userId)
		d.updateBot(userId,update)
		}(user)		
	}
}
}

func prepareUpdateForNotification(userId int)tgbotapi.Update{
	chat:=tgbotapi.Chat{ID: int64(userId)}
	ent:=[]tgbotapi.MessageEntity{{Type: "bot_command",Offset: 0,Length:11}}
	msg:=tgbotapi.Message{Text:"/set_rating",Chat: &chat,Entities: &ent}
	return  tgbotapi.Update{Message:&msg}
}

func (d *Dispatcher)updateBot(userId int,update tgbotapi.Update)  {
	bot,ok:= d.instance(userId,update)
	if !ok{
		go bot.Update()
	}
	bot.UpdateCh <-update
}