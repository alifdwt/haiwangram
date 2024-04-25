test:
	go test -v -cover ./...

swag:
	swag init

server:
	go run main.go

frontend:
	cd web/haiwangram && npm install && npm run dev

.PHONY: test server frontend