package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const coutdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := coutdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{
		1 * time.Second, time.Sleep,
	}
	Countdown(os.Stdout, sleeper)
}
