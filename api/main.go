package main

import "github.com/homma509/9go/handler"

func main() {
	router := handler.NewRouter()
	router.Logger.Fatal(router.Start(":80"))
}
