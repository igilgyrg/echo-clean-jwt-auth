run-local:
	docker compose -f ./local/docker-compose.yml up -d \
	&& go build -o build/todo cmd/main.go && ./build/todo