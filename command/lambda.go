package command

import (
	"context"
	"log"
	"strings"

	"github.com/homma509/nrece/domain/repository"
	"github.com/homma509/nrece/handler"
)

// LambdaCommand コマンドの定義
type LambdaCommand struct {
	ReceiptRepo repository.ReceiptRepository
}

// Synopsis コマンドの簡単な説明
func (c *LambdaCommand) Synopsis() string {
	return "run a lambda function"
}

// Help コマンドの詳細なヘルプメッセージ
func (c *LambdaCommand) Help() string {
	helpText := `
Usage: nrece lambda

	run a lambda function.
`

	return strings.TrimSpace(helpText)
}

// Run コマンド処理の実行
func (c *LambdaCommand) Run(args []string) int {
	if len(args) == 0 {
		log.Println(c.Help())
		return 1
	}
	log.Println("[info] command run", args)

	fileHandler, err := handler.NewFileHandler(c.ReceiptRepo)
	if err != nil {
		log.Println("[error] command run", err)
		return 1
	}
	for _, s3url := range args {
		log.Println("[info] command move file", s3url)
		err = fileHandler.MoveFile(context.Background(), s3url)
		if err != nil {
			log.Println("[error] command move file", s3url, err)
			return 1
		}
	}
	return 0
}
