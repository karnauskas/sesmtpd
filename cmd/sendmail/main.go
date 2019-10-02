package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func main() {
	var (
		content  = []string{}
		mailFrom = ""
	)

	log.SetFlags(0)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	// TODO: make this generic, too cron specific
	for _, line := range content {
		parts := strings.SplitN(strings.TrimRight(line, ">"), "=", 2)
		if parts[0] == "X-Cron-Env: <MAILFROM" {
			mailFrom = parts[1]
			break
		}
		if line == "" {
			break
		}
	}

	for x, line := range content {
		// fmt.Printf("[%d %s]\n", x, line)
		if strings.HasPrefix(line, "From:") {
			content[x] = fmt.Sprintf("From: %s", mailFrom)
		}
		if line == "" {
			break
		}
	}

	// TODO: make this flexible
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.EuWest1RegionID),
	}))

	svc := ses.New(sess)

	out, err := svc.SendRawEmail(&ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{
			Data: []byte(strings.Join(content, "\n")),
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Sent email using SES, ID %s", *out.MessageId)
}
