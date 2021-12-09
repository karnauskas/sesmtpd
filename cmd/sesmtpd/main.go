package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/chrj/smtpd"
	log "github.com/sirupsen/logrus"
)

var (
	bind   = flag.String("bind", "127.0.0.1:1025", "bind address")
	region = flag.String("region", "eu-west-1", "aws region")
)

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", *bind)
	if err != nil {
		log.Fatalln(err)
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}

	server := &smtpd.Server{
		WelcomeMessage: fmt.Sprintf("%s ESMTP ready", hostname),
		Handler:        handler,
	}

	log.Printf("Listening on %s ..\n", ln.Addr())
	log.Fatalln(server.Serve(ln))
}

func handler(peer smtpd.Peer, env smtpd.Envelope) error {
	config := &aws.Config{
		Region: aws.String(*region),
	}

	svc := ses.New(session.Must(session.NewSession(config)))

	_, err := svc.SendRawEmail(&ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{
			Data: env.Data,
		},
	})

	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}
