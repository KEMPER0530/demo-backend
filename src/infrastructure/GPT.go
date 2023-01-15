package infrastructure

import (
	"os"
)

type GPT struct {
	Keyid string
}

func NewGPT() *GPT {
	return &GPT{
		Keyid: os.Getenv("CHATGPT_API_KEY"),
	}
}

func (gpt *GPT) GetKeyid() string {
	return gpt.Keyid
}
