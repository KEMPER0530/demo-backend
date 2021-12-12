package usecase

type NuxtMailRepository interface {
    Send(from string,to string,title string,body string) (*string, error)
}
