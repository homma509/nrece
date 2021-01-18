package command

import (
	"fmt"
	"os"
	"strings"
)

// BatchCommand Batchコマンドの定義
type BatchCommand struct{}

// Synopsis Batchコマンドの簡単な説明
func (c *BatchCommand) Synopsis() string {
	return "run a batch"
}

// Help Batchコマンドの詳細なヘルプメッセージ
func (c *BatchCommand) Help() string {
	helpText := `
Usage: nrece batch

	run a batch.
`

	return strings.TrimSpace(helpText)
}

// Run Batchコマンド処理の実行
func (c *BatchCommand) Run(args []string) int {
	// TODO SIGTERM
	fmt.Fprintln(os.Stdout, "run a batch?!")
	return 0
}
