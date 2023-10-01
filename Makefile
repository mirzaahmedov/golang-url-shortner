include configs/config.toml

.DEFAULT_GOAL := build

cmd_dir = ./cmd
out_dir = ./bin

migrations_dir = ./migrations

build:
	@go build -o $(out_dir)/url-shortner $(cmd_dir)/url-shortner

run: build
	@exec $(out_dir)/url-shortner

migrate_create:
	@migrate create -ext sql -dir $(migrations_dir) -seq $(name)

migrate_up:
	@migrate -path $(migrations_dir) -database "postgres://$(postgres_host):$(postgres_port)"
