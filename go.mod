module github.com/karnauskas/sesmtpd

go 1.18

require (
	github.com/aws/aws-sdk-go v1.44.238
	github.com/chrj/smtpd v0.3.1
	github.com/sirupsen/logrus v1.9.0
)

require (
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
)

replace github.com/eaigner/dkim => github.com/eaigner/opendkim v0.0.0-20140108165311-5359c018d55a
