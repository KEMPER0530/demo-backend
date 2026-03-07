[![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://circleci.com/gh/KEMPER0530/demo-backend)

# demo-backend

このプロジェクトは、Go で作られたバックエンドです。  
主に次の 2 つを行います。

1. AppSync(Lambda) 経由で受けたメール情報を AWS SES で送信する
2. ChatGPT の入出力結果を DynamoDB に保存する

フロントエンドは [demo-front](https://github.com/KEMPER0530/demo-front) を参照してください。

## プロジェクト解説（全体像）

この API は、`operationName` に応じて処理を切り替えます。

1. `createNuxtMail`: メール送信処理
2. `createChatGptResult`: 結果保存処理

### 処理フロー（シンプル版）

```text
AppSync / Lambda Event
        |
        v
   src/main.go
        |
        v
infrastructure.RouteRequest
   |                         |
   | createNuxtMail          | createChatGptResult
   v                         v
NuxtMailController      PutChatGptController
   |                         |
   v                         v
SESRepository           ChatGptRepository
   |                         |
   v                         v
AWS SES                  DynamoDB
```

## プロジェクト構成（構成図）

```text
src/
├── main.go                    # エントリポイント (Lambda / ローカル起動の分岐)
├── domain/                    # エンティティ・値オブジェクト
├── usecase/                   # ユースケース（アプリケーション固有の業務ロジック）
├── interfaces/                # 入出力アダプタ層
│   ├── controllers/           # ユースケース呼び出し
│   ├── aws/                   # SES への具体実装
│   └── database/              # DynamoDB への具体実装
└── infrastructure/            # 実行基盤（ルーティング、環境変数、AWS初期化など）
```

### 各層の役割（初学者向け）

1. `domain`: データの形を定義します（メール、レスポンスなど）。
2. `usecase`: 「何をするか」を定義します（送信する、保存する）。
3. `interfaces`: 外部入出力とユースケースをつなぎます。
4. `infrastructure`: AWS や起動モードなど、実行環境依存の処理を持ちます。

この分離により、変更の影響範囲が小さくなり、テストもしやすくなります。

## 必要環境

- Go 1.26.1（`go.mod` の `go` ディレクティブは `1.26`）
- AWS Lambda
- AWS SES
- AWS AppSync
- AWS DynamoDB

## ローカル実行

事前に `src/infrastructure/development.env` を用意し、必要な AWS 関連環境変数を設定してください。
（`sample.env` をベースに作成）

```bash
GO_ENV=development go run src/main.go -i 0 -f from@test.com -t to@test.com -s "件名" -b "本文"
```

```bash
GO_ENV=development go run src/main.go -i 1 -us test@test.com -in "Input" -ou "Output" -cr "2022/01/01"
```

## テストコードと実行方法

### ユニットテスト

```bash
go test ./src/...
```

### カバレッジ測定（100%目標）

```bash
go test ./src/... -covermode=atomic -coverpkg=./src/... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## GitHub Actions（push 時に自動テスト）

`.github/workflows/test.yml` で、**すべてのブランチへの push** を対象にテストを実行します。

実行内容:

1. Go セットアップ
2. ユニットテスト実行
3. カバレッジ計測
4. 合計カバレッジが `100.0%` でない場合は失敗
