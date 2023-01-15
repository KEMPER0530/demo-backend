package chat

type GPTRepository struct {
	GPT GPT
}

func (gpt *GPTRepository) GetKeyid() string {
	return gpt.GPT.GetKeyid()
}
