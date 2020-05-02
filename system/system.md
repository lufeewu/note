# 简介
操作系统的常用命令

## Shell
ls
cat/nl/tac
    cat -n
more/less
head/tail

df/du/fdisk/mount/umount

cut/grep
wc/sort/uniq
xargs

正则: +?|()()+

ifconfig/ifup/ifdown
route
ip
iwlist/iwconfig
dhclient

ping/traceroute/netstat/host/nslookup

tcpdump/wireshark/nc/netcat


## tcp 问题分析
1. 查看 time_wait 数量
netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'
netstat -an|awk '/tcp/ {print $6}'|sort|uniq -c


## 问题
1. 查找 /home/interview/ 目录下包含 "in" 的文件，显示上下行

## 文章
1. <a href="https://zhuanlan.zhihu.com/p/57473266">面试题：2018最全Redis面试题整理</a>
2. <a href="https://www.cnblogs.com/qccz123456/p/11385748.html">Linux性能优化从入门到实战：07 CPU篇：CPU性能优化方法</a>
3. <a href="https://www.cnblogs.com/qccz123456/p/11204172.html">Linux性能优化从入门到实战：13 内存篇：内存指标/工具总结、问题定位和调优</a>
4. <a href="https://mp.weixin.qq.com/s?src=11&timestamp=1587822345&ver=2300&signature=A8oSIuGrMdixEvmEVvZKwCM-KaSKcFskKZlJUTWKnJB16ridX66wP2mg0QcAnJ5o2ZKwbnosUcZJCbK4svSzn-N1ToQ*JHPSRvUtlQrN6VIS7ot0ouJy-1IiC2HpUv4z&new=1">什么是僵尸进程，如何找到并杀掉僵尸进程？ </a>
5. <a href="https://mp.weixin.qq.com/s/BSZ-mf6YAeHMlc_pc8CTPw">Linux、K8S、Go等是如何设计调度系统的？调度系统设计精要</a>
6. <a href="https://cloud.tencent.com/developer/article/1560422">腾讯、阿里、滴滴后台面试题汇总总结 — （含答案）</a>
7. <a href="https://michaelyou.github.io/2015/03/24/TCP%E6%98%AF%E4%B8%80%E7%A7%8D%E6%B5%81%E5%8D%8F%E8%AE%AE/">tcp 是一种流协议</a>
8. <a href="https://mp.weixin.qq.com/s?src=11&timestamp=1587871941&ver=2301&signature=s3kYqBFmKQC57vxhtcChhWsNSD15HEOZjiqX1G7bzzBP1H0y0dSo5OIz8nqrxUziTI2kAIPNEZPJUwOBLlfZgijXp84fXJ7*Pgb9nI2IITBBdBHkHZn631upnynGqSiz&new=1">TCP滑动窗口</a>
9. <a href="https://zhuanlan.zhihu.com/p/60382685">服务器TIME_WAIT和CLOSE_WAIT详解和解决办法</a>