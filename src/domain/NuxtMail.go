package domain

type NuxtMail struct {
  From string `json:"from"`
  To string `json:"to"`
  Subject string `json:"subject"`
  Body string `json:"body"`
  Createdat string `json:"createdat"`
  Updatedat string `json:"updatedat"`
  Responce int `json:"responce"`
  Result string `json:"result"`
}
