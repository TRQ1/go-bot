package tools

import (
	"fmt"
	"os"

	model "github.com/TRQ1/go-bot/core/types"
	"github.com/slack-go/slack"
)

func UploadFile(u model.Upload, c model.Client) error {
	t := os.Getenv("SlackToken")
	env := os.Getenv("Environment")
	a := slack.New(t)

	filePath := ""
	if env == "local" {
		filePath = u.FileName
	} else {
		filePath = "/tmp/" + u.FileName
	}

	params := slack.FileUploadParameters{
		Title:          c.User + " 님이 요청하신 " + u.FileName,
		File:           filePath,
		InitialComment: c.Service,
		Channels:       []string{c.ChannelName},
	}

	file, err := a.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	fmt.Printf("Name: %s is sent.\n", file.Name)

	return nil
}
