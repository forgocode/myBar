package processbar

import (
	"fmt"
	"log"
	"time"
)

type Bar struct {
	Total   int
	Bar     int
	Timer   time.Duration
	Tag     string
	Output  string
	Percent int
	Done    chan int
	Stop    chan struct{}
}

func (b *Bar) Run() {
	go b.input()
	b.timer()
	b.finsh()
}

func NewBar(total int, timer time.Duration, Tag string) *Bar {
	if total <= 0 {
		log.Panic(fmt.Sprintf("total task is %d, can't product bar", total))
	}
	if Tag == "" {
		Tag = "#"
	}
	if timer == 0 {
		timer = time.Millisecond * 200
	}
	bar := &Bar{
		Total: total,
		Timer: timer,
		Tag:   Tag,
		Done:  make(chan int),
		Stop:  make(chan struct{}),
	}
	return bar
}

func (b *Bar) disPlay() {
	fmt.Printf("\r<%-50s>%8s%d%%%8d/%d", b.Output, "", b.Percent, b.Bar, b.Total)
	b.Output = ""
}

func (b *Bar) finsh() {
	close(b.Stop)
	fmt.Println()
	fmt.Printf("everything completed!\n")
}

func (b *Bar) timer() {
	timer := time.NewTicker(b.Timer)
	for {
		select {
		case <-timer.C:
			if b.Bar >= b.Total {
				b.Percent = 100
				for i := 0; i <= b.Percent/2; i++ {
					b.Output = b.Output + b.Tag
				}
				b.disPlay()
				return
			}
			b.Percent = b.Bar * 100 / b.Total
			for i := 0; i <= b.Percent/2; i++ {
				b.Output = b.Output + b.Tag
			}
			b.disPlay()
		}
	}
}

func (b *Bar) input() {
	for {
		select {
		case <-b.Done:
			b.Bar++
		case <-b.Stop:
			return
		}
	}
}
