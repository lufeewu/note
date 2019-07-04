# golang IO
I/O 是计算机的基础之一，赋予计算机强大的扩展能力，使得其可以为众多业务提供支持. golang 作为一门编程语言，自然需要对软件提供 I/O 能力，使得软件能够更好的使用操作系统的能力.

## 源码结构
+ ioutil
    + ioutil.go
    + example_test.go    
    + ioutil_test.go     
    + tempfile.go   
    + tempfile_test.go
+ example_test.go 
+ io.go         
+ io_test.go  
+ multi.go 
+ multi_test.go 
+ pipe.go      
+ pipe_test.go

golang 的 io 源码库中总共 2800 多行代码，除去 go test 代码，总计 1100 多行代码. 主要代码是 io.go 中提供的一些 i/o interface 及方法. 

ioutil 中则实现了一些 i/o 方法

## 接口
io 中提供了 4 个基本 interface，以及它们的组合，此外还有一些接口提供诸如 WriteAt、ReadAt 等操作. interface 如下
+ **Reader**
+ **Writer**
+ **Closer**
+ **Seeker**
+ ReadWriter
+ ReadCloser
+ WriteCloser
+ ReadWriteCloser
+ ReadSeeker
+ WriteSeeker
+ ReadWriteSeeker
+ ReaderFrom
+ WriterTo
+ ReaderAt
+ WriterAt
+ ByteReader
+ ByteScanner
+ ByteWriter
+ RuneReader
+ RuneScanner
+ StringWriter

## io 的应用
bufio 实现了带缓存的 I/O 操作，它也是 golang 基础库之一. 它实现了 io 库的 Reader、Writer、ReaderWriter 接口, 其中方法 Read(p []byte)(n int,err error) 的实现如下

    func (b *Reader) Read(p []byte) (n int, err error) {
        n = len(p)
        if n == 0 {
            if b.Buffered() > 0 {
                return 0, nil
            }
            return 0, b.readErr()
        }
        if b.r == b.w {
            if b.err != nil {
                return 0, b.readErr()
            }
            if len(p) >= len(b.buf) {
                // Large read, empty buffer.
                // Read directly into p to avoid copy.
                n, b.err = b.rd.Read(p)
                if n < 0 {
                    panic(errNegativeRead)
                }
                if n > 0 {
                    b.lastByte = int(p[n-1])
                    b.lastRuneSize = -1
                }
                return n, b.readErr()
            }
            // One read.
            // Do not use b.fill, which will loop.
            b.r = 0
            b.w = 0
            n, b.err = b.rd.Read(b.buf)
            if n < 0 {
                panic(errNegativeRead)
            }
            if n == 0 {
                return 0, b.readErr()
            }
            b.w += n
        }

        // copy as much as we can
        n = copy(p, b.buf[b.r:b.w])
        b.r += n
        b.lastByte = int(b.buf[b.r-1])
        b.lastRuneSize = -1
        return n, nil
    }

该 Read 将从写入 buf 但还为读出的字节缓存读出到 byte 数组 p 中，并在 buf 都已经读取出时候，归零缓存空间（即 b.r = 0 、 b.w = 0), 并写入新数据到缓存 bufio 中(b.rd.Read(b.buf). 对于缓存 bufio 的写入，可以参看 Write 方法的实现.
