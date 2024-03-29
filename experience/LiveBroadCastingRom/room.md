# 简介
在 4G、5G 及宽带基础设施的不断提升下，线上直播应用快速进入生活，包括电商直播、教育直播、游戏直播等各类直播应用进入大众生活。各类头部应用也在 app 中嵌入了直播功能。
随着用户的爆发增长，也给直播间的技术带来了挑战。如何进行百万用户在线、数十万高并发、低延时的直播间的技术架构，满足直播间的各类互动需求？


## 直播间功能
直播间的功能根据场景不同，会略有差异，但主要都围绕直播及互动相关。
- 基础功能: 直播的基础功能包括连麦互动(多码率、多协议、多主播同框)、美颜特效、弹幕、IM 聊天、点赞、刷礼物、屏幕共享等。
- 非功能性需求: 防盗链、涉黄鉴别等。
- 个性化教育需求: 答题场景、答案公布、生生互动、游戏互动等。
- 个性化电商需求: 商品展示、一键下单等。


## 技术挑战
直播间面临的技术挑战整体上主要是高并发、高带宽、低延时等。
- 音视频处理及传输: 音视频上涉及音视频解码、实时美颜、视频推流、CDN 加速分发、终端适配、播放、流量统计等技术点。
- 高并发: 百万级别的直播间，同时答题、抢红包等情况，可能有百万级别的 QPS。
- 高带宽: 直播观看的码流至少为 1M bps, 百万用户在线，出口带宽则可能达到 1000 G bps ，弹幕推送的出口带宽可能达到 20G bps。
- 高计算: 互动中的瞬时流量。如主播向观众同时推送题目，公布答案瞬间的百万用户瞬时流量等。
- 低延时: 直播场景下视频流和数据流如何整合，做到音画、主播画面、题目同步等，从而保证用户体验。
- 资金流支付: 多人抢红包、答题奖金等的金额一致性。

## 消息推送
实时消息推送通常是 IM 即时通讯类应用的技术。但随着移动互联网的普及，实时推送的应用越来越广泛，如直播间中的弹幕、评论、互动等。
实时推送的到达率、平均端延迟、并发能力等都会影响到用户体验，这些也是消息推送面对的技术挑战，微信、美团、字节等各大公司都有自己的实时推送系统以应对业务需求。

## 参考
1. [带你认识直播平台的技术架构](https://juejin.cn/post/6844904104083324941)
2. [微信团队分享：微信直播聊天室单房间1500万在线的消息架构演进之路](https://chowdera.com/2021/03/20210308230350050s.html)
3. [消息推送技术干货：美团实时消息推送服务的技术演进之路](https://segmentfault.com/a/1190000040481008)
4. [追求极致，揭秘抖音背后的RTC技术](https://juejin.cn/post/7033698281293086733)