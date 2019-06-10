# go-todolist-server
Go言語によるTODOリスト管理のサーバーアプリケーション

## ビルド
```
go get ./...
go run todo_server.go
```
`localhost:8080`をListenする．
`go_test`という名前のdatabaseをMySQLで作成しておく必要がある．

## 概要
### 実装した機能
- todoの追加
- todo一覧の取得
- todoを一件取得
- todoを全削除
- todoを一件削除
- DBにtodoリストを保存

### 使った技術
- 言語: golang (v.1.12.5)
- Webフレームワーク: gin
- DB: MySQL (v.8.0.16)
- ORM: GORM
- CI: CircleCI (v.2.1)

## API一覧
```
# イベント登録 request
POST /api/v1/event
{"deadline": "2019-06-11T14:00:00+09:00", "title": "レポート提出", "memo": ""}

# イベント登録 response
200 OK
{"status": "success", "message": "registered", "id": 1}

400 Bad Request
{"status": "failure", "message": "invalid date format"}
```

```
# イベント全取得 request
GET /api/v1/event

# イベント全取得 response
200 OK
{"events": [
    {"id": 1, "deadline": "2019-06-11T14:00:00+09:00", "title": "レポート提出", "memo": ""},
    ...
]}
```

```
# イベント1件取得 request
GET /api/v1/event/${id}

# イベント1件取得 response
200 OK
{"id": 1, "deadline": "2019-06-11T14:00:00+09:00", "title": "レポート提出", "memo": ""}

404 Not Found
```

```
# イベント全削除 request
DELETE /api/v1/event

# イベント全削除 response
200 OK
{"status": "success", "message": "deleted"}
```

```
# イベント1件削除 request
DELETE /api/v1/event/${id}

# イベント1件削除 response
200 OK
{"status": "success", "message": "deleted", "id": 1}

404 Not Found
```

`deadline`はRFC3339形式の文字列のみ許容される．

## ローカルテスト
```
go run todo_server.go
go test todo_server_test.go
```
