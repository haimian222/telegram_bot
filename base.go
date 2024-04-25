package telegram_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// BotBase 机器人基类
type BotBase struct {
	botID       int64            // 机器人ID
	botApi      *tgbotapi.BotAPI // 机器人API
	messageChan chan *Message    // 消息通道
	eventChan   chan *Event      // 事件通道
}

// NewBotBase 创建机器人基类
func NewBotBase(token string, messageChan chan *Message, eventChan chan *Event) (bot *BotBase, err error) {
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot = &BotBase{
		botID:       botApi.Self.ID,
		botApi:      botApi,
		messageChan: messageChan,
		eventChan:   eventChan,
	}
	return bot, nil
}

// GetBotUsername 获取机器人用户名
func (bot *BotBase) GetBotUsername() string {
	return bot.botApi.Self.UserName
}

// ReceiveMessage 接收消息
func (bot *BotBase) ReceiveMessage() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.botApi.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			// ignore any non-Message Updates
			continue
		}

		if update.Message.Text != "" {
			//处理文本消息
			message := Message{
				BotID:     bot.botID,
				MessageID: update.Message.MessageID,
				ChatInfo: &ChatInfo{
					ChatID:   update.Message.Chat.ID,
					Title:    update.Message.Chat.Title,
					Type:     update.Message.Chat.Type,
					UserName: update.Message.Chat.UserName,
				},
				FromInfo: &FromInfo{
					ID:        update.Message.From.ID,
					UserName:  update.Message.From.UserName,
					FirstName: update.Message.From.FirstName,
					LastName:  update.Message.From.LastName,
				},
				MessageType: "text",
				Container: &MessageContent{
					Text: update.Message.Text,
				},
			}
			bot.messageChan <- &message
		}

	}
}

// SendMessageText 发送消息文本
func (bot *BotBase) SendMessageText(chatID int64, text string) (messageID int, err error) {
	msg := tgbotapi.NewMessage(chatID, text)
	message, err := bot.botApi.Send(msg)
	if err != nil {
		return 0, err
	}
	return message.MessageID, nil
}
