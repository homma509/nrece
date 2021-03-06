package main

import (
	"context"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/homma509/nrece/command"
	"github.com/homma509/nrece/handler"
	"github.com/homma509/nrece/infra"
	"github.com/homma509/nrece/infra/db"
	"github.com/homma509/nrece/infra/file"
	"github.com/mitchellh/cli"
)

func main() {
	if strings.HasPrefix(os.Getenv("AWS_EXECUTION_ENV"), "AWS_Lambda_go") {
		lambda.Start(lambdaHandler)
		return
	}

	exitCode, err := cmd()
	if err != nil {
		log.Println("[error] main", err)
	}
	os.Exit(exitCode)
}

func lambdaHandler(ctx context.Context, evt events.S3Event) error {
	log.Printf("[info] s3 events %#v", evt)
	for _, record := range evt.Records {
		log.Printf("[info] s3 event %#v", record.S3)
		u := url.URL{
			Scheme: "s3",
			Host:   record.S3.Bucket.Name,
			Path:   record.S3.Object.Key,
		}
		log.Println("[info] main lambda handler", u.String())

		s3Repo, err := file.NewS3Repository()
		if err != nil {
			log.Println("[error] main new s3 hander", err)
			return err
		}
		fileHandler, err := handler.NewFileHandler(s3Repo)
		if err != nil {
			log.Println("[error] main new file handler", err)
			return err
		}
		err = fileHandler.CopyFile(ctx, u.String())
		if err != nil {
			log.Println("[error] main copy file", err)
			return err
		}
	}
	return nil
}

func cmd() (int, error) {
	// TODO implements version
	c := cli.NewCLI("nrece", "1.0.0")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"server": func() (cli.Command, error) {
			sql, err := db.NewSQLHandler()
			if err != nil {
				log.Println("[error] main server command", err)
				// return nil, err
			}
			app := infra.NewAppRepository(sql)
			return &command.ServerCommand{AppRepo: app}, nil
		},
		"batch": func() (cli.Command, error) {
			return &command.BatchCommand{}, nil
		},
		"lambda": func() (cli.Command, error) {
			s3, err := file.NewS3Repository()
			if err != nil {
				log.Println("[error] main lambda", err)
				return nil, err
			}
			return &command.LambdaCommand{ReceiptRepo: s3}, nil
		},
	}

	return c.Run()
}
