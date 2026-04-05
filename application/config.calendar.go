package application

import (
	"context"
	"log"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var cal_srv *calendar.Service

func gCalendarConnect() {
	ctxb := context.Background()
	srv, err := calendar.NewService(ctxb, option.WithCredentialsFile("./secret.GCPcredentials.json"))
	if err != nil {
		log.Printf("Unable to retrieve Calendar client: %v", err)
		return
	}
	cal_srv = srv
}
