package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) sleep() {
	s.Calls++
}

type ConfigerableSleeper struct {
	duration time.Duration
}

func (o *ConfigerableSleeper) sleep() {
	time.Sleep(o.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigerableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)

}
