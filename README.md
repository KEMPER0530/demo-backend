[![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://circleci.com/gh/KEMPER0530/go-lambda-api-demo)

# go-lambda-api-demo

![go-lambda-api-demo](https://user-images.githubusercontent.com/43329853/142201779-72293879-e582-4ddb-8f99-daa04bdb58a4.png)

サンプルは「API Gateway」 + 「Lambda」で作成したAPIとなります。
APIGatewayで作成したREST APIにリクエストを投げてdynamoDBへ接続できるかどうかを確認します。

Go言語で作成し、コンテナイメージとしてLambda上で動作します。

# Requirement

- golang 1.17.2
- Gin

# Usage

 1. dynamoDBの作成(data配下のサンプルを利用してテーブルを作成)
 2. Lambdaへ本プログラムをデプロイ
 3. APIGatewayの作成
 4. Curlの実行

```
# 例
curl -X GET https://XXXXXXXXXXXXXX.execute-api.ap-northeast-1.amazonaws.com/dev/V1/actuator-health | jq

% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   101  100   101    0     0    878      0 --:--:-- --:--:-- --:--:--   885
{
  "Code": 200,
  "DBSts": "CONNECTED",
  "Time": "2021-11-17T21:40:38.442062149+09:00",
  "Host": "XXX.XXX.XX.XXXX"
}

```
