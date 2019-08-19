# 爬虫

通过 channel - 爬取珍爱网

- 【engine】
    - [引擎](./engine)
- 【fetcher】
    - [抓取器](./fetcher)
- 【scheduler】
    - [调度器](./scheduler)
- 【parser】
    - [解析器](./zhenai/parser)
- 【爬虫架构】
    - [单任务版爬虫架构](./zhenai/parser)
    - [并发版爬虫架构](#并发版爬虫架构)
    - 【并发版爬虫架构】
        - [实现I](#实现i)
        - [实现II：并发分发Request](#实现ii并发分发request)
        - [实现III：Request队列和Worker队列](#实现iiirequest队列和worker队列)

#### 单任务版爬虫架构

[!Alt Text](./doc_images/simple.png)

#### 并发版爬虫架构
```
                                       +----------------------------------------------+
                                       |                                              |
                                       |                                     Worker   |
                                       |                                              |
                                       |                                              |
+--------------+            +--------------+      response      +------- ---------+   |
|              |  request   |          |   +------------------->+                 |   |
|     Seed     +----------->+    Engine|   |  requests,items    |     Parser      |   |
|              |            |          |   +<-------------------+                 |   |
+--------------+            +------------+-+                    +-----------------+   |
                     +------->         | > ^------+                                   |
                     |-------+         | |        |                                   |
                     ||                | |        |                                   |
                     ||                | |url     |response                           |
                     <>                | |        |                                   |
     +-----------------------+         | |        +-------------+------------------+  |
     |                       |         | |                      |                  |  |
     |     Task Queue        |         | ----------------------->     Fetcher      |  |
     |                       |         |                        |                  |  |
     +-----------------------+         |                        +------------------+  |
                                       |                                              |
                                       +----------------------------------------------+

```

#### 实现I:

[!Alt Text](./doc_images/concurrent_simple.png)


#### 实现II：并发分发Request

[!Alt Text](./doc_images/concurrent_simple_goroutine.png)

#### 实现III：Request队列和Worker队列

[!Alt Text](./doc_images/concurrent.png)
