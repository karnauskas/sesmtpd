sendmail:
	GOOS=linux go build -ldflags="-s -w" -o sendmail cmd/sendmail/main.go

sesmtpd:
	GOOS=linux CGO=0 go build -ldflags="-s -w" -o sesmtpd cmd/sesmtpd/main.go

clean:
	rm -f sendmail sesmtpd
	go clean
	go mod tidy
