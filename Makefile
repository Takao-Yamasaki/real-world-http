.PHONY: help client server
.DEFAULT_GOAL := help
ups: ##  Goサーバーのコンテナを起動
	@docker compose up server
upc: ##  Goサーバーのコンテナを起動
	@docker compose up client
down: ## コンテナを停止
	@docker compose down
execc: ## Goクライアントのコンテナに入る（開発用）
	@docker compose exec -it client sh
execs: ## Goサーバーのコンテナに入る（開発用）
	@docker compose exec -it server sh
ps: ## コンテナの稼働状況確認
	@docker compose ps
runs: ## Goサーバーを起動（開発用）
	@docker compose exec -it server go run server.go
runc: ## Goクライアントを起動（開発用）
	@docker compose exec -it client go run client.go
rebuild: ## コンテナをリビルド（開発用）
	@docker compose up --build
help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
