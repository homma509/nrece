package command

import (
	"fmt"
	"os"
	"strings"
)

// BatchCommand コマンドの定義
type BatchCommand struct{}

// Synopsis コマンドの簡単な説明
func (c *BatchCommand) Synopsis() string {
	return "run a batch"
}

// Help コマンドの詳細なヘルプメッセージ
func (c *BatchCommand) Help() string {
	helpText := `
Usage: nrece batch

	run a batch.
`

	return strings.TrimSpace(helpText)
}

// Run コマンド処理の実行
func (c *BatchCommand) Run(args []string) int {
	// TODO SIGTERM
	fmt.Fprintln(os.Stdout, "run a batch?!")
	return 0
}
