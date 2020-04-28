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