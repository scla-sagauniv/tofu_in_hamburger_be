# tofu_in_hamburger_be
ハックツハッカソンモサカップ　チーム豆腐入りハンバーグ　バックエンド

## migration
```
migrate create -ext sql -dir app/db/migrations -seq create_materials
migrate -database 'mysql://<user>:<password>@tcp(db:3306)/tofu_in_hamburger' -path db/migrations up
migrate -database 'mysql://<user>:<password>@tcp(db:3306)/tofu_in_hamburger' -path db/migrations down
migrate -database 'mysql://<user>:<password>@tcp(db:3306)/tofu_in_hamburger' -path db/migrations force 1
migrate -database 'mysql://<user>:<password>@tcp(db:3306)/tofu_in_hamburger' -path db/migrations version
```
