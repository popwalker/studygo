build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
	docker build -t superfat/say .
	rm -f app
