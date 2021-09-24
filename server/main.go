package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"server/cron/mail"
	"server/cron/scheduler"
	"server/cron/worker"
	"server/database"
	"server/routes"
	"sync"
)

func main() {
	database.Connect()
	apiKey := "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" //TODO Change that!
	msgExchange := make(chan *mail.EmailContent)

	wg := &sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())

	// sql
	mailer := mail.NewSendgrid(apiKey)
	sched := scheduler.NewScheduler(ctx, database.DB, msgExchange)
	worker := worker.NewWorker(ctx, wg, database.DB, mailer, msgExchange)

	/////////////////////////////////////////////////
	// graceful shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c // waits for the termination signal
		fmt.Printf("Got %s signal. Exiting...\n", sig)
		sched.Stop()
		cancelFunc()
		routes.StopRouter(ctx, wg)
	}()
	/////////////////////////////////////////////////

	wg.Add(2) // add 1 for worker only. don't need for scheduler
	// run worker (as a receiver of msgExchange channel first)
	go worker.Start()
	go routes.Init()

	sched.Start() // start to scan order

	// wait for the worker finishes its job
	wg.Wait()
	// prepare params
}
