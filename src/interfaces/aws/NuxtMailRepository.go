package aws

import (
	"errors"
	"log"
	"mailform-demo-backend/src/domain"
	//"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type NuxtMailRepository struct {
}

func (nm *NuxtMailRepository) Send(arg domain.NuxtMail, region string, id string, secret string) (*string, error) {
	AWS_REGION := region
	AWS_ACCESS_KEY_ID := id
	AWS_SECRET_KEY := secret

	from := arg.From
	to := arg.To
	subject := arg.Subject
	body := arg.Body

	//AWS-SESの設定情報を格納する
	awsSession := session.New(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_KEY, ""),
	})

	// メール送信情報を設定する
	svc := ses.New(awsSession)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(from),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				log.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				log.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				log.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Println(err.Error())
		}
		log.Println(result)
		return result.MessageId, errors.New(err.Error())
	}

	log.Println("Email Sent to address: " + to)
	log.Println(result)

	return result.MessageId, nil
}
