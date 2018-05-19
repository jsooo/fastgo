# fastgo
Fastgo 是一个基于 `fasthttp` 和 `fasthttprouter` 的API微框架，甚至微小到没有框架应有的那种全局掌控力，所有的东西都可以由开发者自行定义，当然了，在开发者什么都不写的情况下pull下来也是可以直接使用。

### 使用方法：
#### step.1 
`go get github.com/jsooo/fastgo`
#### step.2
`go build`
#### step.3
`./fastgo`
output:
```
➜  fastgo git:(master) ✗ ./fastgo
2018/05/19 20:00:21 Fastgo Running On Port: 8081 .....
```

好了，此时已经运行起来了。

现在在浏览器中输入 `http://127.0.0.1:8081/v1/hello_world?say=hi` 就能看到浏览器再向你打招呼。

----

这只是第一版，实现了我最开始的想法，框架看起来还是很简陋的，我会持续更新，做到足够轻量而又足够方便 ：）

PS:附上输出hello_world接口的压测
```
ab -n 10000 -c 100 http://127.0.0.1:8081/v1/hello_world\?say\=hi
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        fasthttp
Server Hostname:        127.0.0.1
Server Port:            8081

Document Path:          /v1/hello_world?say=hi
Document Length:        25 bytes

Concurrency Level:      100
Time taken for tests:   0.918 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1790000 bytes
HTML transferred:       250000 bytes
Requests per second:    10890.66 [#/sec] (mean)
Time per request:       9.182 [ms] (mean)
Time per request:       0.092 [ms] (mean, across all concurrent requests)
Transfer rate:          1903.74 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        2    4   1.3      4      10
Processing:     1    5   1.4      4      11
Waiting:        1    3   1.2      3      10
Total:          6    9   1.6      9      16

Percentage of the requests served within a certain time (ms)
  50%      9
  66%      9
  75%     10
  80%     10
  90%     12
  95%     12
  98%     13
  99%     13
 100%     16 (longest request)
```
