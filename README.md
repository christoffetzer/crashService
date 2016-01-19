# crashService

A simple crash service that starts and then exits after a random number of seconds.
The minimum and the maximum uptime can be given via argument flags as well as the exit code.

Example: 

To exit with error code -2 within 5..10 seconds after being started, execute:

     ./crashService -min_uptime=5 -max_uptime=10 -exit_code=-2

## Building this service:

Install this service on your machine:

    go get github.com/christoffetzer/crashService

Build the service:

    go build github.com/christoffetzer/crashService

Run the service:

    ./crashService --help

