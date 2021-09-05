# 简介
golang 的 archive 标准库提供了压缩文件的存取. 包括 tar 和 zip.


## 源码
当前涉及 tar 及 zip 库. 总计 11000 多行代码，除去测试代码 4600 多行
+ tar
    - 总计 7200 多行，除去测试代码 3000 多行
    - type Header 文件头部信息
    - type Reader
        - 实现对 tar 档案文件的顺序读取
    - type Writer
        - 提供了对 POSIX.1 格式的 tar 档案文件的顺序写入
+ zip
    - 总计 4000 多行，除去测试代码 1600 多行
    - 提供 zip 档案文件的读写服务
    - type FileHeader
    - type File
    - type File
    - type ReadCloser
    - type Writer


## ref
1. [.ZIP File Format Specification](http://www.pkware.com/documents/casestudies/APPNOTE.TXT)
2. [FreeBSD File Formats Manual](https://www.freebsd.org/cgi/man.cgi?query=tar&sektion=5)
3. [Basic Tar Format](http://www.gnu.org/software/tar/manual/html_node/Standard.html)
4. [pax - portable archive interchange](http://pubs.opengroup.org/onlinepubs/9699919799/utilities/pax.html)