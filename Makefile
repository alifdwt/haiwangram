test:
	go test -v -cover ./...

server:
	go run main.go

frontend:
	cd web/haiwangram && npm install && npm run dev

.PHONY: test server frontend