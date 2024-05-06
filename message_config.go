package telegram_bot

// FileData 文件数据
type FileData struct {
	Name  string
	Bytes []byte
}

// PhotoConfig 图片配置
type PhotoConfig struct {
	FileData FileData
	Text     string
}

// NewPhotoConfig 创建图片配置
func NewPhotoConfig(fileData FileData, text string) PhotoConfig {
	return PhotoConfig{
		FileData: fileData,
		Text:     text,
	}
}

// DocumentConfig 文件配置
type DocumentConfig struct {
	FileData FileData
	Text     string
}

// NewDocumentConfig 创建文件配置
func NewDocumentConfig(fileData FileData, text string) DocumentConfig {
	return DocumentConfig{
		FileData: fileData,
		Text:     text,
	}
}
