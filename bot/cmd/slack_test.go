package tools

import (
	"testing"
)
type simpleMessage struct {
	kinds		string
	channel		string
	user		string
	text		string
	ts			string
}

func TestUpload()