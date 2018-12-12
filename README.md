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

现在在浏览器中输入 `http://127.0.0.1:8081/v1/hello_world?say=hi` 就能看到浏览器再向你打招呼。

----

这只是第一版，实现了我最开始的想法，框架看起来还是很简陋的，我会持续更新，做到足够轻量而又足够方便 ：）

PS:附上输出hello_world接口的压测
```
hey -n 100000 -c 10 http://127.0.0.1:8081/v1/hello_world\?say\=hi

Summary:
  Total:	4.0429 secs
  Slowest:	0.0953 secs
  Fastest:	0.0001 secs
  Average:	0.0004 secs
  Requests/sec:	24734.5539
  
  Total data:	2500000 bytes
  Size/request:	25 bytes

Response time histogram:
  0.000 [1]	    |
  0.010 [99933]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.019 [47]	|
  0.029 [8]	    |
  0.038 [5]   	|
  0.048 [2] 	|
  0.057 [2] 	|
  0.067 [0] 	|
  0.076 [0] 	|
  0.086 [0] 	|
  0.095 [2] 	|


Latency distribution:
  10% in 0.0001 secs
  25% in 0.0002 secs
  50% in 0.0003 secs
  75% in 0.0004 secs
  90% in 0.0006 secs
  95% in 0.0009 secs
  99% in 0.0020 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0001 secs, 0.0953 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0000 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0352 secs
  resp wait:	0.0003 secs, 0.0000 secs, 0.0901 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0509 secs

Status code distribution:
  [200]	100000 responses
```
