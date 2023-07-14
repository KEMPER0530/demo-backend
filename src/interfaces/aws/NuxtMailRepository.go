package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kemper0530/demo-backend/src/domain"
	"log"
)

type NuxtMailRepository struct{}

func (nm *NuxtMailRepository) Send(arg domain.NuxtMail, region, keyID, secretKey string) (*string, error) {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(keyID, secretKey, ""),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	svc := ses.New(session)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(arg.To)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(arg.Body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(arg.Subject),
			},
		},
		Source: aws.String(arg.From),
	}
	result, err := svc.SendEmail(input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Email Sent to address: " + arg.To)
	log.Println(result)
	return result.MessageId, nil
}
