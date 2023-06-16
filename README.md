# tofu_in_hamburger_be
ハックツハッカソンモサカップ　チーム豆腐入りハンバーグ　バックエンド

## migration
```
migrate create -ext mysql -dir app/db/migrations -seq create_materials
migrate -database 'mysql://<user>:<password>@tcp(db:3306)/tofu_in_hamburge' -path db/migrations up
migrate -database 'mysql://<user>:<password>@tcp(db:3306)/tofu_in_hamburge' -path db/migrations down
migrate -database 'mysql://<user>:<password>@tcp(db:3306)/tofu_in_hamburge' -path db/migrations force
migrate -database 'mysql://<user>:<password>@tcp(db:3306)/tofu_in_hamburge' -path db/migrations version
```
