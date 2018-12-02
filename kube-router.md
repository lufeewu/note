# kube-router 
kubernetes 1.9 推出的代替 kube-proxy 和 calico 等插件的网络方案。采用 lvs 实现 svc 网络，使用 bgp 实现 pod 网络，相比而言更易发布和扩展对接。

# bgp 
## next hop 
1. 从 EBGP 邻居学到一条路由，传递给自己的 EBGP 邻居时，下一跳改变
2. 从 IBGP 邻居学到一条路由，传递给自己的 EBGP 邻居，下一跳改变
3. 从 EBGP 邻居学到的一条路由，传递给自己的 IBGP 邻居时，下一跳不变，仍为 EBGP 邻居的更新源
4. 从 IBGP 邻居学到的一条路由。传递给自己的 IBGP 邻居时，这种情况不会发生，IBGP 邻居之间只传一跳，为IBGP的防环机制
## 


# ref
1. Kube-Router: kubernetes pod-to-pod networking with BGP. https://asciinema.org/a/120885
2. https://cloudnativelabs.github.io/post/2017-05-22-kube-pod-networking/
3. https://github.com/cloudnativelabs/kube-router	
