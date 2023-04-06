# 简介
音视频技术相关知识

## 概念
h264/h265/266: 视频号压缩标准. 其中 h265 比 h264 压缩效率提升 50%, 而 h266 比 h265 压缩效率提升 50%.
硬/软编码: 硬编码是通过硬件设备提供的模块进行编解码，而软编码则是通过软件程序进行编解码.硬编的好处是编码快、不占用 CPU 资源.而软编的优点是不受硬件限制，但是需要占用大量的 CPU.
关键 I 帧: Intra-coded picture, 帧内编码图像帧, I 帧表示关键帧. 这一帧保留完整画面, 作为随机访问的参考点.
P 帧、B 帧、IDR 关键帧: P 帧表示这一帧和之前的一个关键帧的差别，解码时需要之前缓存的画面叠加上本帧定义的差别，生成最终画面.B 帧是双向差别帧，B 帧记录的是本帧与前后帧的差别.IDR 帧属于 I 帧，解码器收到 IDR frame 时，将所有的参考帧队列丢弃，是强制刷新帧.

## 视频质量评价
- PSNR: Peak signal-to-noise ratio 峰值信噪比, 是一种工程表示用来表示最高的信号量和影响其表示精度的噪声功率的比值.
- VMAF: Video Multi-method Assessment Fusion, 它借助人类视觉模型以及机器学习来评估视频的质量.
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
## 参考文献
1. [I帧、P帧、B帧、GOP、IDR 和PTS, DTS之间的关系](https://www.cnblogs.com/yongdaimi/p/10676309.html)
2. [VMAF: The Journey Continues](https://netflixtechblog.com/vmaf-the-journey-continues-44b51ee9ed12)
3. [从 350ms 到 80ms，打造 iOS 短视频的极致丝滑体验](https://learnku.com/articles/60568)
4. [短视频预加载技术（一）](https://johnsonlee.io/2021/02/10/short-video-preloading-1/)