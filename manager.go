package telegram_bot

import (
	"errors"
	"strconv"
	"strings"
)

// BotManager 机器人管理类
type BotManager struct {
	botMap      map[int64]*BotBase // 机器人映射
	messageChan chan *Message      // 消息通道
	eventChan   chan *Event        // 事件通道
}

// NewBotManager 创建机器人管理类
func NewBotManager() *BotManager {
	return &BotManager{
		botMap:      make(map[int64]*BotBase),
		messageChan: make(chan *Message, 10000),
		eventChan:   make(chan *Event, 1000),
	}
}

// botExists 判断机器人是否存在
func (manager *BotManager) botExists(botID int64) bool {
	_, exists := manager.botMap[botID]
	return exists
}

// GetBotIDFromToken 切割token获取botID
func GetBotIDFromToken(token string) (botID int64, err error) {
	tempList := strings.Split(token, ":")
	if len(tempList) != 2 {
		return 0, errors.New("token format error")
	}
	botID, err = strconv.ParseInt(tempList[0], 10, 64)
	if err != nil {
		return 0, err
	}
	return botID, nil
}

// AddBot 添加机器人
func (manager *BotManager) AddBot(token string) (botID int64, err error) {
	bot, err := NewBotBase(token, manager.messageChan, manager.eventChan)
	if err != nil {
		return 0, err
	}
	//判断机器人是否存在
	if manager.botExists(bot.botID) {
		return 0, errors.New("bot already exists")
	}
	//添加机器人
	manager.botMap[bot.botID] = bot
	go manager.botMap[bot.botID].ReceiveMessage()

	return bot.botID, nil
}

// GetMessageChan 获取消息通道
func (manager *BotManager) GetMessageChan() chan *Message {
	return manager.messageChan
}

// GetEventChan 获取事件通道
func (manager *BotManager) GetEventChan() chan *Event {
	return manager.eventChan
}

// SendMessageText 发送消息文本
func (manager *BotManager) SendMessageText(botID int64, chatID int64, text string) (messageID int, err error) {
	if !manager.botExists(botID) {
		return 0, errors.New("bot not exists")
	}
	return manager.botMap[botID].SendMessageText(chatID, text)
}

// SendMessagePhoto 发送消息图片
func (manager *BotManager) SendMessagePhoto(botID int64, chatID int64, photoBytes []byte, phoneName string, text string) (messageID int, err error) {
	if !manager.botExists(botID) {
		return 0, errors.New("bot not exists")
	}
	return manager.botMap[botID].SendMessagePhoto(chatID, photoBytes, phoneName, text)
}

// GetBotUsername 获取机器人用户名
func (manager *BotManager) GetBotUsername(botID int64) (username string, err error) {
	if !manager.botExists(botID) {
		return "", errors.New("bot not exists")
	}
	return manager.botMap[botID].GetBotUsername(), nil
}
