# 简介
音视频技术相关知识.

## 概念
h264/h265/266: 视频号压缩标准. 其中 h265 比 h264 压缩效率提升 50%, 而 h266 比 h265 压缩效率提升 50%.
硬/软编码: 硬编码是通过硬件设备提供的模块进行编解码，而软编码则是通过软件程序进行编解码.硬编的好处是编码快、不占用 CPU 资源.而软编的优点是不受硬件限制，但是需要占用大量的 CPU.
关键 I 帧: Intra-coded picture, 帧内编码图像帧, I 帧表示关键帧. 这一帧保留完整画面, 作为随机访问的参考点.
P 帧、B 帧、IDR 关键帧: P 帧表示这一帧和之前的一个关键帧的差别，解码时需要之前缓存的画面叠加上本帧定义的差别，生成最终画面.B 帧是双向差别帧，B 帧记录的是本帧与前后帧的差别.IDR 帧属于 I 帧，解码器收到 IDR frame 时，将所有的参考帧队列丢弃，是强制刷新帧.

## 视频质量评价
- PSNR: Peak signal-to-noise ratio 峰值信噪比, 是一种工程表示用来表示最高的信号量和影响其表示精度的噪声功率的比值.
- VMAF: Video Multi-method Assessment Fusion, 它借助人类视觉模型以及机器学习来评估视频的质量. [Netflix/vmaf](https://github.com/Netflix/vmaf)
- SSIM: Structural Similarity Index Measure 是一种预测数字电视、电影图像以及其他类型数字图像或视频的感知质量的方法.

## 音视频优化实践
视频内容已经成为各大互联网应用标配. 视频播放体验关系到用户停留时长. 下面是一些常见的优化实践方法.
- **首帧渲染时长**: 视频页面完全展示到首帧渲染时刻的耗时.
- **预加载**: 在合适的时机提前下载视频数据. 预加载包括预下载和预解码
    - 预下载: 预先将视频资源从 CDN 下载到本地
    - 预解码: 预先让播放器对下载到本地的视频文件进行解码
- **码率**: h266、h265、h264 等规格.
- **双播放器策略**: 预播下一个视频, 加载到首帧并暂停.
- **体验质量**: QoE(Quailty of Experience), 从终端用户角度出发, 以用户对 App 的主观感受衡量满意程度. 短视频的主要 QoE 指标有观看次数(video view)、人均观看时长(average played time)、完播率(play complete ratio)、评论率(comment rate)、人均评论停留时长(average stay duration in comment)、点赞率(like rate)、收藏率(favorite rate)、转发率(forward rate)、加粉率(follow rate)、负反馈率(negative feedback rate)
- **服务质量**: QoS(Quality of Service), 偏向于从客观角度出发, 通过各种参数衡量服务的整体性能. 对短视频 QoS 指标主要有传输延时(transport latency)、编码延时(encode latency)、解码延时(decode latency)、首帧时间(time to first frame)、帧率(frame per second)、秒开率(sec-opening rate)、缓存命中率(cache hit ratio)等
- **多码率**: 同一个视频会存在多个码率, 根据用户的网络环境、视频内容、设备性能自适应码率, 称为码率自适应技术(Adaptive Bitrate Streaming).
- **moov**: moov atom 定义了时间尺度、时长、显示特性以及用于在电影每个轨道的信息.
- **极速高清**: 
## ffmpeg
ffmpeg 是一个处理多媒体内容如视频、音频、字幕和相关元数据的库和工具集合.
- 库集合: libavcodec、libavformat、libavutil、libavfilter、libavdevice、libswresample、libswscale.
- 工具: ffmpeg、ffplay、ffprobe、aviocat、ismindex、qt-faststart

## srs
SRS 是一个简单、高效、实时的媒体服务, 支持 RTMP、WebRTC、HLS、HTTP-FLV、SRT、MPEG-DASH、GB28181 协议. 单进程支持约 9000 并发.

目前主流的流媒体服务器有开源的 NginxRTMP、Crtmpd、Red5、SRS 等.
- **NginxRTMP**: 支持音视频直播, 支持 flv/mp4 等格式, 支持 push、pull 模式, 可以录制 flv 等. 单进程支持约 3000 并发.
- **Crtmpd**: CrtmpServer 是一个由 c++ 编写的开源高性能 RTMP 流媒体服务器. [shiretu/crtmpserver](https://github.com/shiretu/crtmpserver)
- **Red5**: java 编写的开源流媒体服务 [red5-server](https://github.com/Red5/red5-server), 支持视频流、音频流、直播发布等功能, 支持 RTMP、RTMPT、RTMPS 和 RTMPE 等协议.

## 音视频格式
1. AAC: 高级音频编码(Advanced Audio Coding), 一种专门声音数据设计的文件压缩格式, 它采用了全新的算法进行编码, 更加高效, 有更高的性价比.
2. HDR: 高动态范围(High-Dynamic Range), 指动态范围特别高的应用, 一般的图像可能是 8 bit 的无符号字节表示, 而 HDR 图像则一般为 32 位浮点类型. HDR 可以呈现(模拟)出人色彩感知范围内相似的色彩, 尽可能地还原人眼所看到的景色.
3. SDR: 标准动态范围(Standard Dynamic Range), 是指一种很常见的色彩显示方式, 信息大小相比 HDR 更小, 普及度更高.
4. HLS: 动态自适应技术(HTTP Live Streaming) 是 Apple 的动态自适应技术, 主要用于 PC 和 Apple 终端的音视频服务.

## 音频处理
音频处理库
- [audioflux](https://github.com/libAudioFlux/audioFlux): 用于音频、音乐分析、特征提取
- [torchaudio](https://github.com/pytorch/audio): 用于音频信号处理的数据操作和转换，由 PyTorch 提供技术支持
- [librosa](https://github.com/librosa/librosa): 用于音频和音乐分析、描述及合成的 c++ 库, 提供 Python 绑定
- [essentia](https://github.com/MTG/essentia): 用于音频和音乐分析的 Python 库

## 参考文献
1. [I帧、P帧、B帧、GOP、IDR 和PTS, DTS之间的关系](https://www.cnblogs.com/yongdaimi/p/10676309.html)
2. [VMAF: The Journey Continues](https://netflixtechblog.com/vmaf-the-journey-continues-44b51ee9ed12)
3. [从 350ms 到 80ms，打造 iOS 短视频的极致丝滑体验](https://learnku.com/articles/60568)
4. [短视频预加载技术（一）](https://johnsonlee.io/2021/02/10/short-video-preloading-1/)
5. [github - FFmpeg/FFmpeg](https://github.com/FFmpeg/FFmpeg)
6. [github - ossrs/srs](https://github.com/ossrs/srs)
7. [crtmpserver 系列(二): 搭建简易流媒体直播系统](https://www.cnblogs.com/wangqiguo/p/6014519.html#_label0)
8. [OpenCV4学习笔记（59）——高动态范围（HDR）成像](https://blog.csdn.net/weixin_45224869/article/details/105895367)
9. [SDR - 百度百科](https://baike.baidu.com/item/SDR/22316143)
10. [音频处理效率测评：audioflux、torchaudio、librosa和essentia库哪个更快？](https://juejin.cn/post/7225856176131293243)