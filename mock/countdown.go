package mock

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

// type DefaultSleeper struct{}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

type ConfigurableSleeper struct {
	DurationConfig time.Duration
	SleepConfig    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepConfig(c.DurationConfig)

}
