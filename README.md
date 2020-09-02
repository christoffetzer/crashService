你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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

