package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rajatjindal/ttrpc/proxy"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ttrpc-proxy <dir-with-sock-files>")
		os.Exit(0)
	}

	err := proxy.Start(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)
	signal.Notify(sigterm, syscall.SIGINT)
	<-sigterm

	<-time.NewTicker(100 * time.Millisecond).C
}
