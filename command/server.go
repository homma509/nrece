package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/homma509/nrece/handler"
)

const (
	defaultPort int = 80
)

// ServerCommand Serverコマンドの定義
type ServerCommand struct{}

// Synopsis Serverコマンドの簡単な説明
func (c *ServerCommand) Synopsis() string {
	return "run a rest api server"
}

// Help Serverコマンドの詳細なヘルプメッセージ
func (c *ServerCommand) Help() string {
	helpText := `
Usage: nrece server [options]

	run a rest api server.

Options:

	-p, -port	use port number (default 80)
`

	return strings.TrimSpace(helpText)
}

// Run Serverコマンド処理の実行
func (c *ServerCommand) Run(args []string) int {
	var port int
	flags := flag.NewFlagSet("server", flag.ContinueOnError)
	flags.IntVar(&port, "p", defaultPort, "port to use")
	flags.IntVar(&port, "port", defaultPort, "port to use")
	if err := flags.Parse(args); err != nil {
		return 1
	}

	// TODO SIGTERM
	router := handler.NewRouter()
	router.Logger.Fatal(router.Start(fmt.Sprintf(":%d", port)))
	return 0
}
