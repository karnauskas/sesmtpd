module github.com/karnauskas/sesmtpd

go 1.18

require (
	github.com/aws/aws-sdk-go v1.44.100
	github.com/chrj/smtpd v0.3.1
	github.com/sirupsen/logrus v1.9.0
)

require (
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace github.com/eaigner/dkim => github.com/eaigner/opendkim v0.0.0-20140108165311-5359c018d55a
