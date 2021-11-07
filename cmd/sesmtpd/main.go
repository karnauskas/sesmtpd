package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/chrj/smtpd"
)

var (
	bind      = flag.String("bind", "127.0.0.1:1025", "bind address")
	debug     = flag.Bool("debug", false, "debug")
	region    = flag.String("region", "eu-west-1", "aws region")
	timestamp = flag.Bool("timestamp", false, "show timestamp")
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
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}

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

	sess, err := session.NewSession()
	if err != nil {
		return err
	}

	client := ec2metadata.New(sess)

	creds := credentials.NewCredentials(&ec2rolecreds.EC2RoleProvider{
		Client: client,
	})

	config := &aws.Config{
		Region:                        aws.String(*region),
		Credentials:                   creds,
		CredentialsChainVerboseErrors: aws.Bool(*debug),
	}

	sess = session.Must(session.NewSession(config))

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
