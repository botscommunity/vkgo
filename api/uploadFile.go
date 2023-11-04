package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

func get(link string) (data []byte, err error) {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, link, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if err := response.Body.Close(); err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	data = body

	return
}

type UploadFile []byte

func (bot *Bot) UploadFile(serverURL, fileURL, field, name string) (uploadedFile UploadFile, err error) {
	body := new(bytes.Buffer)
	write := multipart.NewWriter(body)

	part, err := write.CreateFormFile(field, name)
	if err != nil {
		return nil, err
	}

	file, err := get(fileURL)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, bytes.NewReader(file))
	if err != nil {
		return nil, err
	}

	contentType := write.FormDataContentType()
	_ = write.Close()

	request, err := http.NewRequestWithContext(context.Background(), "POST", serverURL, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", contentType)

	response, err := bot.Bot.HTTPClient.Do(request)
	if err != nil {
		err := response.Body.Close()
		if err != nil {
			return nil, err
		}

		return nil, err
	}

	content, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return nil, err
	}

	return content, nil
}

func (file UploadFile) JSON(data any) error {
	return json.Unmarshal(file, data)
}
