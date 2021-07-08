package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {

	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		//time.Sleep(1 * time.Second)
		sleeper.Sleep()
		fmt.Fprintln(out, i) // print number and newline
	}
	//time.Sleep(1 * time.Second)
	sleeper.Sleep()
	fmt.Fprint(out, "Go!")

}

func main() {
	sleeper := &DefaultSleeper{}

	Countdown(os.Stdout, sleeper)

	csleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, csleeper)

}

/*
print 3
print 3,2,1 and Go!
wait a second between each line

*/
