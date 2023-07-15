package domain

type ChatGptResult struct {
	MessageID string `json:"messageID" dynamo:"MessageID,hash"` // hash key
	User      string `json:"user" dynamo:"User"`
	Input     string `json:"input" dynamo:"Input"`
	Output    string `json:"output" dynamo:"Output"`
	Createdat string `json:"createdat" dynamo:"Createdat,range"` // range key
}
