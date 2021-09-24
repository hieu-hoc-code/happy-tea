package worker

import (
	"context"
	"fmt"
	"server/cron/mail"
	"sync"

	"gorm.io/gorm"
)

// Worker defines a worker
type Worker struct {
	wg     *sync.WaitGroup
	mailer mail.Mailer
	inChan <-chan *mail.EmailContent
	ctx    context.Context
	db     *gorm.DB
}

// NewWorker creates new worker
func NewWorker(ctx context.Context, wg *sync.WaitGroup, db *gorm.DB, mailer mail.Mailer, ch <-chan *mail.EmailContent) *Worker {
	return &Worker{
		ctx:    ctx,
		wg:     wg,
		mailer: mailer,
		inChan: ch,
		db:     db,
	}
}

func (w *Worker) Start() {
	if w.mailer == nil || w.db == nil {
		fmt.Println("cannot start worker since mailer is nil")
		return
	}
	for {
		select {
		case em := <-w.inChan: // sending email now
			err := w.mailer.Send(em)
			if err != nil {
				fmt.Println("Cannot send email due to error: ", err)
				continue
			}
			// update sql data
			w.db.Exec("UPDATE ecommerce.orders SET thank_you_email_sent = true WHERE id = ?", em.ID)

		case <-w.ctx.Done():
			fmt.Println("Exiting worker")
			w.wg.Done()
			return
		}
	}
}
