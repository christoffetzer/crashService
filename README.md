# crashService

A simple crash service that starts and then exits after a random number of seconds.
The minimum and the maximum uptime can be given via argument flags as well as the exit code.

Example: 

To exit with error code -2 within 5..10 seconds after being started, execute:

     ./crashService -min_uptime=5 -max_uptime=10 -exit_code=-2
