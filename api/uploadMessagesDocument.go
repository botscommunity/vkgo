package api

import (
	"github.com/botscommunity/vkgo/internal/pkg/responses"
)

func (bot *Bot) UploadMessagesDocument(chatID int, fileURL, extension string) (typ string, document responses.SaveUploadMessageDocument, err *responses.Error) {
	server := bot.GetUploadMessagesDocumentServer(chatID)
	if server.Error != nil {
		err = server.Error
		return
	}

	uploaded, uploadErr := bot.UploadFile(server.UploadURL, fileURL, "file", "document."+extension)
	if err != nil {
		err = responses.NewInternalError(uploadErr)
		return
	}

	documentFile := responses.UploadDocumentFile{}
	if serializeErr := responses.NewInternalError(uploaded.JSON(&documentFile)); serializeErr != nil {
		err = serializeErr
		return
	}

	return bot.SaveUploadMessagesDocument(documentFile.File)
}
