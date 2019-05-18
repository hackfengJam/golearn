# 流量统计系统

- 【run】
    - [run.go - 生成大量日志](./run/run.go)
- 【analysis】
    - channel + redis
    - [analysis.go - 基于日志流量统计](./analysis.go)

#### 流量统计系统架构
```
+--------------------------+                        +---------------------+
|                          |                        |                     |
|   read log line by line  |  Channel(logChannel)   |     logConsumer     |
|                          +----------------------->+                     |
+--------------------------+                        +---------+-----------+
                                                              |
                                                              |
                                                              |  Channel(pvChannel uvChannel ...)
                                                              |
                                                              |
                                                              v
                                           +------------------+-----------------------+
                                           |                                          |
                                           |     Counter(pvCounter uvCounter ...)     |
                                           |                                          |
                                           +------------------+-----------------------+
                                                              |
                                                              |
                                                              |  Channel(storageChannel ...)
                                                              |
                                                              |
                                                              v
                                          +-------------------+--------------------------+
                                          |                                              |
                                          |       dataSorage(into Redis/Hbase ...)       |
                                          |                                              |
                                          +----------------------------------------------+

```