
linux:
	GOOS=linux CGO=0 go build -o sesmtpd.linux

mac:
	GOOS=darwin CGO=0 go build -o sesmtpd.osx
