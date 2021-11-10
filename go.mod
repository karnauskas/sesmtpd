module github.com/karnauskas/sesmtpd

go 1.17

require (
	github.com/aws/aws-sdk-go v1.42.1
	github.com/chrj/smtpd v0.3.0
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

replace github.com/eaigner/dkim => github.com/eaigner/opendkim v0.0.0-20140108165311-5359c018d55a
