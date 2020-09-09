package model

import "github.com/slack-go/slack"

type Upload struct {
	FileName string
	Title     string
}

type Client struct {
	User      		  string
	Service    		  string
	Client            *slack.Client
	VerificationToken string
	ChannelName         string
}

type Message struct {
	Type            string       `json:"type,omitempty"`
	Text            string       `json:"text,omitempty"`
}