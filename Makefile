APP=go-issues-api
DB=go-issues-db

app.start:
	docker-compose up -d;

app.stop:
	docker-compose down;

app.restart: app.stop app.start

db.cli:
	docker exec -it $(DB) psql -U $(DB_USER)

test.unit:
	docker exec -it $(APP) go test ./tests/unit/... -v
