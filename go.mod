module github.com/karnauskas/sesmtpd

go 1.17

require (
	github.com/aws/aws-sdk-go v1.43.38
	github.com/chrj/smtpd v0.3.1
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
)

replace github.com/eaigner/dkim => github.com/eaigner/opendkim v0.0.0-20140108165311-5359c018d55a
