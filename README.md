[![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://circleci.com/gh/KEMPER0530/mailform-demo-backend)

# mailform-demo-backend

![MailDemo_ALL](https://user-images.githubusercontent.com/43329853/145986290-6506f0ad-6043-4e03-b13d-04553d92be1f.png)

このプロジェクトは Golang から AWS SES にメールを送信する機能です。バックエンド部分は AppSync(GraphQL)からパラメータを受け取り、Lambda で SES に送信します。
送信結果は「delivery」「bounce」を含め、SES→SNS→Lambda→dynamoDB に格納されます。
加えて、GPT-3 を使って、チャットボットシステムを構築しています。

フロントは[こちら](https://github.com/KEMPER0530/mailform-demo-front)を参照ください

# Requirement

- golang 1.17.2
- AWS Lambda
- AWS SES
- AWS AppSync
- AWS dynamoDB
- AWS SNS
- AWS ECR
- [AWS Route53](https://qiita.com/NaokiIshimura/items/89e104dd2d8dd950780e)
- [ChatGPT](https://openai.com/api/)

# Usage

ローカル環境で実行する場合、事前にドメイン取得、Route53 の設定、SES の設定は実施済を前提とします

```
$ GO_ENV=development go run src/main.go -i 0 -f from@test.com -t to@test.com -s [件名] -b [本文]
```
