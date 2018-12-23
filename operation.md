## monitor
grafana 是一个开源有丰富指标的可视化面板


### prometheus
prometheus 是一套开源的监控系统，将所有信息都存储为时间序列数据.
#### 功能
1. 多维数据模型（时序名由 metric 名字和 k/v 的 labels 构成）
2. 查询语句（ PromQL）
3. 无依赖存储，支持 local 和 remote 不同模型
4. http 协议，使用 pull 拉取数据
5. 监控目标可以采用服务发现或静态配置方式
6. 支持多种统计数据模型，图形化友好
#### 核心组件
1. Prometheus Server
2. Client libraries
3. push gateway
4. exporters
5. alertmanager

#### ref
1. https://mp.weixin.qq.com/s?src=11&timestamp=1545555403&ver=1317&signature=U4wF9oFqqGmz6Dt2vQlTtHBPKKb8nil8IGxzmFLp39YwinEWWWgoV-IUPZwaol-qunApkWr3DXDbQtNYpcNEmr-*i0qCN-BovpAzd9RzSNhJlaTnPDZwcDIHr512S1IA&new=1
