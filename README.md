[![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://circleci.com/gh/KEMPER0530/mailform-demo-backend)

#

![mailform-demo-backend]()

Golang から AWS SES を呼びメール送信する機能です。
バックエンド部分は AppSync(GraphQL)から連携されたパラメータを Lambda で受け取り
SES へ送信内容を連携します。

# 開発環境

- golang 1.17.2
