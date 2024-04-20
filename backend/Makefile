run-server:
	@echo "Running server..."
	@go run cmd/main.go

run-migrate:
	go run cmd/migrate/main.go $(name)

apply-migrate:
	@chmod +x scripts/migrate.sh
	@scripts/migrate.sh

ent-create-model:
	go run -mod=mod entgo.io/ent/cmd/ent new $(model) --target ./db/ent

ent-gen:
	go generate ./db/ent

ent-describe:
	go run -mod=mod entgo.io/ent/cmd/ent describe ./db/ent/schema

con-db:
	psql -h localhost -p 5432 -U postgres -d app

new-network:
	docker network create kendai-navi-network

delete-network:
	docker network rm kendai-navi-network