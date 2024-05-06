package telegram_bot

type Message struct {
	BotID       int64           // 机器人ID
	MessageID   int             // 消息ID
	ChatInfo    *ChatInfo       // 聊天信息
	FromInfo    *FromInfo       // 发送者信息
	MessageType string          // 消息类型
	Container   *MessageContent // 消息内容
}

// ChatInfo 聊天信息
type ChatInfo struct {
	ChatID   int64  // 聊天ID
	Title    string // 聊天标题
	Type     string // 聊天类型
	UserName string // 用户名
}

// FromInfo 发送者信息
type FromInfo struct {
	ID        int64  // 发送者ID
	UserName  string // 用户名
	FirstName string // 名字
	LastName  string // 姓氏
}

// MessageContent 消息内容
type MessageContent struct {
	Text     string      // 文本消息
	Photo    []PhotoSize // 图片消息
	Document *Document   // 文件消息
}

// Document 文件消息
type Document struct {
	FileID       string     // 文件ID
	FileUniqueID string     // 文件唯一ID
	Thumbnail    *PhotoSize // 缩略图
	FileName     string     // 文件名
	MimeType     string     // MIME类型
	FileSize     int        // 文件大小
}

// PhotoSize 图片消息
type PhotoSize struct {
	FileID       string // 文件ID
	FileUniqueID string // 文件唯一ID
	Width        int    // 宽度
	Height       int    // 高度
	FileSize     int    // 文件大小
}
