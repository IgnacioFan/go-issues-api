APP=go-issues-api

app.start:
	docker-compose up -d;

app.stop:
	docker-compose down;

app.restart: app.stop app.start

db.cli:
	docker exec -it $(APP)_db_1 psql -U $(DB)

test.unit:
	docker exec -it $(APP)-web-1 go test ./tests/unit/... -v
