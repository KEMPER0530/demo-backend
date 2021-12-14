[![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://circleci.com/gh/KEMPER0530/mailform-demo-backend)

# mailform-demo-backend

![MailDemo_ALL](https://user-images.githubusercontent.com/43329853/145986290-6506f0ad-6043-4e03-b13d-04553d92be1f.png)

Golang から AWS SES を呼びメール送信する機能です。
バックエンド部分は AppSync(GraphQL)から連携されたパラメータを Lambda で受け取り
SES へ送信内容を連携します。

フロントは[こちら](https://github.com/KEMPER0530/mailform-demo-front)を参照ください

# 開発環境

- golang 1.17.2
