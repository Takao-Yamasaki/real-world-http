up:
	docker compose up -d
exec:
	docker compose exec -it app sh
ps:
	docker compose ps
run:
	docker compose exec -it app go run server.go
