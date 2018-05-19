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
