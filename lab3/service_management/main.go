package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Println("Quitting program...")
				Exit()
			case syscall.SIGHUP:
				log.Println("Reloading...")
			default:
				log.Println("Received signal ", s)
			}
		}
	}()

	log.Println("Program started...")
	for {
		time.Sleep(time.Second)
		log.Println("Program running...")
	}
}

func Exit() {
	os.Exit(0)
}
