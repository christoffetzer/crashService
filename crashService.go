// (C) Christof Fetzer, 2016
//
// only for education use

// A simple crash service that starts and then exits after a random number of seconds.
// The minimum and the maximum uptime can be given via argument flags as well as the exit code.
//
// Example: to exit with error code -2 within 5..10 seconds after being started, execute:
//      ./crashService -min_uptime=5 -max_uptime=10 -exit_code=-2
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	minUp    = flag.Int("min_uptime", 1, "minimum uptime in seconds")
	maxUp    = flag.Int("max_uptime", 5, "maximum uptime in seconds")
	exitCode = flag.Int("exit_code", -1, "exit code")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			`%s exists with an error after a random number of seconds. 
The minimum and the maximum uptime can be given via argument flags
as well as the exit code.
`, os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	sleeptime := *minUp
	if *maxUp > *minUp {
		sleeptime = *minUp + rand.Intn(*maxUp-*minUp)
	}

	log.Printf("Started - exiting in %d seconds\n", sleeptime)
	time.Sleep(time.Duration(sleeptime) * time.Second)
	log.Println("Exiting.")
	os.Exit(*exitCode)
}
