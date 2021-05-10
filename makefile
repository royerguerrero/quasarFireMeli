runserver:
	go run ./cmd/server/main.go

httpd:
	CompileDaemon -build="go build -o daemon ./cmd/server/main.go" -command="./daemon"

deploy:
	docker-compose -f ./deployments/docker-compose.yml up --build