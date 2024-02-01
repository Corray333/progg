.SILENT:
build:
	cd cmd && go build main.go
front:
	cd frontend && npm run dev
run: build
	cd cmd && ./main