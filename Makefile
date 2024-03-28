test:
	go test -v -cover ./...

server:
	go run main.go

frontend:
	cd web/frontend && npm install && npm run build

.PHONY: test server frontend