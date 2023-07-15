# Go-cache

设计一个分布式缓存系统，需要考虑资源控制、淘汰策略、并发、分布式节点通信等各个方面的问题。而且，针对不同的应用场景，还需要在不同的特性之间权衡，例如，是否需要支持缓存更新？还是假定缓存在淘汰之前是不允许改变的。不同的权衡对应着不同的实现。

groupcache 是 Go 语言版的 memcached，目的是在某些特定场合替代 memcached。groupcache 的作者也是 memcached 的作者。无论是了解单机缓存还是分布式缓存，深入学习这个库的实现都是非常有意义的。

GeeCache 基本上模仿了 groupcache 的实现，裁剪了部分功能。但总体实现上，还是与 groupcache 非常接近的。支持特性有：

    单机缓存和基于 HTTP 的分布式缓存
    最近最少访问(Least Recently Used, LRU) 缓存策略
    使用 Go 锁机制防止缓存击穿
    使用一致性哈希选择节点，实现负载均衡
    使用 protobuf 优化节点间二进制通信
    …


# 目录

LRU 缓存淘汰策略 | Code - Github

单机并发缓存 | Code - Github

HTTP 服务端 | Code - Github

一致性哈希(Hash) | Code - Github

分布式节点 | Code - Github

防止缓存击穿 | Code - Github

使用 Protobuf 通信 | Code - Github


