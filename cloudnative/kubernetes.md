

# kubernetes

+ 容器
    - linux namespace、linux cgroups、rootfs
    - 容器运行时、容器镜像
+ kubeadm
    - kubeadm init
    - kubeadm join
    - kubelet / kubectl / kubeadm
    - port: 10250/10251/10252
    - /etc/kubernetes/pki

+ pod
    - kubernetes 最小 API 对象，原子调度单位
    - 一组对等关系的容器，共享某些资源
    - infra 容器，k8s.gcr.io/pause
    - init Container
    - sidecar 模式
    - 日志收集 /var/log
    - <a href="./pdf/design_patterns_for_container_based_distributed_system.pdf" title="容器设计模式">容器设计模式论文</a>
    - 调度、网络、存储、安全等属性
    - image、command、workingDir、Ports、volumeMounts
    - status: Pending、Running、Succeeded、Failed、Unknown
    - Projected Volume: Secret、ConfigMap、Downward API、ServiceAccountToken
    - Secret


# 资料
- <a  href="https://www.yuque.com/baxiaoshi/tyado3/bl6lev">图解 Kubernetes Pod 创建流程</a>