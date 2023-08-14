# 简介
构建 flink 的 docker dmeo, 操作及应用实践. 

## 实践介绍
1. 通过 docker yml 构建 flink 服务
2. 通过 docker yml 构建 kafka 等数据服务
3. 在 flink 中下载 jar 依赖
4. 调整 flink 的 Slots 数量, 配置 askmanager.numberOfTaskSlots 调整为 10
6. 进入 flink 命令行使用 DDL 创建 user_behavior 表, 从 kafka 中读取数据, 语句如下:
```
    CREATE TABLE user_behavior (
        user_id BIGINT,
        item_id BIGINT,
        category_id BIGINT,
        behavior STRING,
        ts TIMESTAMP(3),
        proctime as PROCTIME(),   -- 通过计算列产生一个处理时间列
        WATERMARK FOR ts as ts - INTERVAL '5' SECOND  -- 在ts上定义watermark，ts成为事件时间列
    ) WITH (
        'connector.type' = 'kafka',  -- 使用 kafka connector
        'connector.version' = 'universal',  -- kafka 版本，universal 支持 0.11 以上的版本
        'connector.topic' = 'user_behavior',  -- kafka topic
        'connector.startup-mode' = 'earliest-offset',  -- 从起始 offset 开始读取
        'connector.properties.zookeeper.connect' = 'zookeeper:2181',  -- zookeeper 地址
        'connector.properties.bootstrap.servers' = 'kafka:9092',  -- kafka broker 地址
        'format.type' = 'json'  -- 数据源格式为 json
    );
```

# 参考
1. [Demo：基于 Flink SQL 构建流式应用](https://wuchong.me/blog/2020/02/25/demo-building-real-time-application-with-flink-sql/)
