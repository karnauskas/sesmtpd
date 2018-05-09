package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/chrj/smtpd"
)

var (
	timestamp = flag.Bool("timestamp", false, "show timestamp")
	bind      = flag.String("bind", "127.0.0.1:25", "bind address")
	region    = flag.String("region", "eu-west-1", "aws region")
	debug     = flag.Bool("debug", false, "debug")
)

func main() {
	flag.Parse()

	if !*timestamp {
		log.SetFlags(0)
	}

	ln, err := net.Listen("tcp", *bind)
	if err != nil {
		log.Fatalln(err)
	}
	hostname, _ := os.Hostname()

	server := &smtpd.Server{
		WelcomeMessage: fmt.Sprintf("%s ESMTP ready", hostname),
		Handler:        handler,
	}

	if *debug {
		logger := log.New(os.Stdout, "", 0)
		server.ProtocolLogger = logger
	}

	log.Printf("Listening on %s ..\n", ln.Addr())

	log.Fatalln(server.Serve(ln))
}

func handler(peer smtpd.Peer, env smtpd.Envelope) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(*region),
	}))
	svc := ses.New(sess)

	if *debug {
		log.Println(string(env.Data))
	}

	output, err := svc.SendRawEmail(&ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{
			Data: env.Data,
		},
	})

	if err != nil {
		return err
	}

	if *debug {
		log.Println(output.String())
	}

	return nil
}
