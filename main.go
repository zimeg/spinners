package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/briandowns/spinner"
)

// cleanup displays any signals before exiting
func cleanup(c chan os.Signal) {
	select {
	case sig := <-c:
		fmt.Printf("\nTEST: signal received (%+v)\n", sig)
		if s, ok := sig.(syscall.Signal); ok {
			time.Sleep(100 * time.Millisecond)
			os.Exit(int(s))
			_ = s
		} else {
			os.Exit(1)
		}
	}
}

// main loads a spinner that can be interrupted
func main() {
	signalFlag := flag.Bool("signal", false, "set true to catch signals in the test program")
	flag.Parse()

	fmt.Printf("SPIN TEST: preparing test program #%d\n", os.Getpid())
	time.Sleep(1 * time.Second)
	fmt.Printf("SPIN TEST: creating a new spinner...\n")

	s := spinner.New(spinner.CharSets[35], 140*time.Millisecond, spinner.WithWriter(os.Stderr))
	s.Suffix = " something is happening..."

	if *signalFlag {
		stopChan := make(chan os.Signal, 1)
		signals := []os.Signal{syscall.Signal(0x0), syscall.SIGINT, syscall.SIGTERM}
		signal.Notify(stopChan, signals...)

		fmt.Printf("SPIN TEST: press CTRL-C to begin the test... ")
		<-stopChan
		fmt.Printf("\n")
		go cleanup(stopChan)
	} else {
		time.Sleep(1 * time.Second)
		fmt.Printf("SPIN TEST: beginning the test program!\n")
	}

	fmt.Printf("SPIN TEST: starting spinner...")
	s.Start()
	fmt.Printf(" started!\n")

	time.Sleep(1 * time.Second)
	s.Prefix = "SPIN TEST: spinning the spinner! (x1)\n"
	time.Sleep(1 * time.Second)
	s.Prefix += "SPIN TEST: spinning the spinner! (x2)\n"
	time.Sleep(4 * time.Second)

	s.FinalMSG = "SPIN TEST: the spinner has spun!"
	s.Suffix = ""
	s.Restart()
	time.Sleep(1 * time.Millisecond)
	s.Stop()
	fmt.Printf("\n")
}
