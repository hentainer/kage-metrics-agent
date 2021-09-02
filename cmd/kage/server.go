package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/msales/kage"
	"github.com/msales/kage/server"
	"gopkg.in/urfave/cli.v1"
)

func runServer(c *cli.Context) {
	app, err := newApplication(c)
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	monitorTicker := time.NewTicker(30 * time.Second)
	defer monitorTicker.Stop()
	go func() {
		for rang