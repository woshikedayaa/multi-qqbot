package utils

import (
	"strconv"
	"time"
)

const (
	ProjectName = "mulit-qqbot"
)

func CreateNewDailyTimer(afterSecond int64, action func()) {
	go func() {
		timer := time.NewTimer(time.Duration(afterSecond) * time.Second)

		for {
			select {
			case <-timer.C:
				action()
				timer.Reset(24 * time.Hour)
			}
		}

	}()
}

func CreateFullTime(s string) string {
	t := time.Now()
	year := strconv.Itoa(t.Year())

	month := strconv.Itoa(int(t.Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	day := strconv.Itoa(t.Day())
	if len(day) == 1 {
		day = "0" + day
	}
	ft := year + "-" + month + "-" + day + " " + s
	return ft
}
