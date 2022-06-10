package local_service

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Scheduler struct {
	Time chan int
}

func NewScheduler() *Scheduler {	
	return &Scheduler{
		Time: make(chan int),
	}
}

func (s *Scheduler) Process() {
	for {
		hour, minute,seconds:=time.Now().Clock()
		minutesLeft, secondLeft:=60-minute,60-seconds
		log.Info("Scheduler sleep", minutesLeft,":",60-seconds)
	    s.Time<-hour
		time.Sleep(time.Duration(minutesLeft)*time.Minute+time.Duration(secondLeft)*time.Second)				
	}
}