# 简介
MIME (Multipurpose Internet Mail Extensions) 多用途互联网邮件扩展类型. 是用于设定某种扩展名的文件用一种应用程序打开的方式类型，当扩展名被访问时候，浏览器会自动使用指定应用程序来打开. 

在 HTTP 协议中也用到了 MIME 的框架，标准被扩展为互联网媒体类型.

golang 的 mime 标准库实现了 MIME 的部分规定.
 
## 源码
golang 的 mime 基础库实现了部分 MIME 规定, 总计 5000 多行代码. 除测试外的总代码数 2300 行. mime 库提供了四个基本的格式化、解析方法:
- func AddExtensionType(ext, typ string) error
- func FormatMediaType(t string, param map[string]string) string
- func ParseMediaType(v string) (mediatype string, params map[string]string, err error)
- func TypeByExtension(ext string) string

## 模块
+ mime/multipart
    - 实现了 MIME 的 multipart 解析，参见 [RFC 2046](http://tools.ietf.org/html/rfc2046)
    - 适用于 HTTP ([RFC 2388](http://tools.ietf.org/html/rfc2388)) 和常见浏览器生成的 multipart 主体
    - type File
    - type FileHeader
    - type Part
    - type Form
    - type Reader
    - type Writer
+ mime/quotedprintable
    - 标准库 quotedprintable 实现了 quoted-printable , 参见 [RFC 2045](http://tools.ietf.org/html/rfc2045)
    - type Reader
    - type Writer

## ref
1. [HTTP协议之multipart/form-data请求分析](https://blog.csdn.net/five3/article/details/7181521)