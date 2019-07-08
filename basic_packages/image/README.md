# 简介
标准库 image 提供了基本的 2D 图片库. 主要提供 png、gif、jpeg 类型的基本处理

## 源码
总计 15000 多行代码,除去测试代码仅 5400 多行.
+ image/color
+ image/draw
+ image/gif
+ image/internal
+ image/jpeg
+ image/png
+ type Config struct
+ type Image interface
+ type PalettedImage interface
+ type RGBA struct
+ type NRGBA strcut
+ type NRGBA64 struct
+ type Gray struct
+ type Paletted struct
+ type Point struct
+ type Rectangle struct
+ type format struct
+ RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error))
+ Decode(r io.Reader) (Image, string, error)
+ DecodeConfig(r io.Reader)(Config, string, error) 解析颜色模型和维度
+ image/color 
    - 实现了基本的色彩库
+ image/palette 
    - 提供标准的调色板
+ image/draw
    - 提供了图像合成函数
    - type Quantizer interface 
    - type Drawer interface
+ image/gif
    - gif包实现了gif文件的编码器和解码器
    - type GIF struct
    - func Decode(r io.Reader) (image.Image, error)
    - func DecodeConfig(r io.Reader) (image.Config, error)
+ image/jpeg
    - 实现 jpeg 格式编解码
+ image/png
    - 实现了 png 图像的编解码
+ image/internal/imageutil
    - DrawYCbCr(dst *image.RGBA, r image.Rectangle, src *image.YCbCr, sp image.Point) (ok bool)

## 应用
1. golang 的 GUI 库 github.com/andlabs/ui


## ref
1. [golang image/draw](http://golang.org/doc/articles/image_draw.html)
2. [gif](http://www.w3.org/Graphics/GIF/spec-gif89a.txt)
3. [Joint Photographic Experts Group (JPEG)](http://www.w3.org/Graphics/JPEG/itu-t81.pdf)
3. [Portable Network Graphics (PNG) Specification (Second Edition)](http://www.w3.org/TR/PNG/)