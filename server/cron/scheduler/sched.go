package scheduler

import (
	"context"
	"fmt"
	"server/cron/mail"
	"time"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

const (
	DefaultThankyouSubject   = "Thank you for purchasing from mystore.com"
	DefaultThankyouBodyPlain = "Thank you for purchasing from our store. Here's your order details:"
	DefaultThankyouBodyHtml  = "<strong>Thank you for purchasing from our store. Here's your order details:</strong>"
	DefaultFromName          = "duc hieu"
	DefaultFromEmail         = "duchieu.ctk41@gmail.com"
)

type Scheduler struct {
	db      *gorm.DB
	c       *cron.Cron
	outChan chan<- *mail.EmailContent
	ctx     context.Context
}

func NewScheduler(ctx context.Context, db *gorm.DB, ch chan<- *mail.EmailContent) *Scheduler {
	return &Scheduler{
		ctx:     ctx,
		db:      db,
		c:       cron.New(cron.WithSeconds()),
		outChan: ch,
	}
}

func (sched *Scheduler) Start() {
	// runs this function every minute
	sched.c.AddFunc("1 8 * * * *", sched.scheduleJob)
	sched.c.Start()
}

func (sched *Scheduler) Stop() {
	fmt.Println("Stopping scheduler")
	sched.c.Stop()
}

func (sched *Scheduler) scheduleJob() {
	fmt.Printf("Scanning for new order(s) at %v\n", time.Now().Format("2006-Jan-02 15:04:05"))
	resp, err := sched.getEmailForSending()
	if err != nil {
		return
	}
	fmt.Printf("Scheduling %v email(s) at %v\n", len(resp), time.Now().Format("2006-Jan-02 15:04:05"))
	for _, em := range resp {
		sched.outChan <- em
	}
}

// getEmailForSending get email and fill up enough information ready for sending
func (sched *Scheduler) getEmailForSending() ([]*mail.EmailContent, error) {
	resp, err := sched.scanFromDB()
	if err != nil {
		return resp, err
	}
	// fill FromUser
	// why we can set FromUser here?
	for _, emailContent := range resp {
		emailContent.FromUser = &mail.EmailUser{
			Name:  DefaultFromName,
			Email: DefaultFromEmail,
		}
	}

	return resp, err
}

// scanFromDB get all orders that match the predefined condition (created_at < now - 1 min && thankyou_email_sent == falses)
func (sched *Scheduler) scanFromDB() ([]*mail.EmailContent, error) {
	var resp []*mail.EmailContent
	// fromTime := time.Now().Add(-time.Minute * 2) // subtract by 2 minutes - why not one?

	type UUU struct {
		ID           uint   `json:"id"`
		CustomerName string `json:"customer_name"`
		Email        string `json:"email"`
	}
	var uuu UUU
	rows, err := sched.db.Raw("SELECT id, customer_name, email FROM ecommerce.orders WHERE thank_you_email_sent = ?;", false).Rows()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// MUST to call this function at the end to free connection to mysql
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&uuu.ID, &uuu.CustomerName, &uuu.Email)
		if err != nil {
			fmt.Println("Cannot scan row due to error: ", err)
			continue
		}
		fmt.Println("vao: ", uuu.Email)
		resp = append(resp, &mail.EmailContent{
			ID:               int64(uuu.ID),
			Subject:          DefaultThankyouSubject,
			PlainTextContent: DefaultThankyouBodyPlain,
			HtmlContent:      DefaultThankyouBodyHtml,
			ToUser: &mail.EmailUser{
				Name:  uuu.CustomerName,
				Email: uuu.Email,
			},
		})
	}
	return resp, nil
}
