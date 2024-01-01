.PHONY: help
.DEFAULT_GOAL := help
up: ## コンテナを起動
	@docker compose up -d
down: ## コンテナを停止
	@docker compose down
exec: ## Goのコンテナに入る（開発用）
	@docker compose exec -it app sh
ps: ## コンテナの稼働状況確認
	docker compose ps
run: ## Goサーバーを起動（開発用）
	@docker compose exec -it app go run server.go
rebuild: ## コンテナをリビルド（開発用）
	@docker compose up --build
help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
