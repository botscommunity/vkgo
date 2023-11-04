package api

import (
	"github.com/botscommunity/vkgo/internal/pkg/responses"
)

func (bot *Bot) UploadMessagesPhoto(chatID int, fileURL string) (photos []responses.SaveUploadMessagesPhoto, err *responses.Error) {
	server := bot.GetUploadMessagesPhotoServer(chatID)
	if server.Error != nil {
		err = server.Error
		return
	}

	uploaded, uploadErr := bot.UploadFile(server.UploadURL, fileURL, "photo", "photo.png")
	if uploadErr != nil {
		err = responses.NewInternalError(uploadErr)
		return
	}

	photoFile := responses.UploadPhotoFile{}
	if serializeErr := responses.NewInternalError(uploaded.JSON(&photoFile)); serializeErr != nil {
		err = serializeErr
		return
	}

	return bot.SaveUploadMessagesPhoto(photoFile.Server, photoFile.Photo, photoFile.Hash)
}
