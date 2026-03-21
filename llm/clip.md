# 简介
学习关于 clip 相关的模型.

## CLIP
clip 模型是 OpenAI 在 2021 年提出的开创性多模态基础模型, 核心思想是通过对比学习, 将图像和文本映射到同一个语义向量空间中, 从而实现"图文匹配"、"零样本分类"等强大能力. Clip 由两个独立编码器图像编码器(ViT 或 ResNet) 或文本编码器(Text Encoder) Transformer-based 组成. 主要的训练方式是通过对比学习, 比如图文对(如 32768 对) 分别用图像编码器和文本编码器输出图像、文本特征.
-   ViT-L/14 @ 336px: OpenAI 官方发布的模型, 约 304M 参数量. 优势是高分辨率细节感知强, 适合检测模糊、雨滴、遮挡等像素级问题. HF ID 是 openai/clip-vit-large-patch14-336.
-   ViT-g/14: LAION 团队基于 open_clip 库训练的超大规模 CLIP 模型, 是目前最大的 CLIP 视觉主干. 参数量约 1.02B. 优势是语义理解能力强, 零样本分类性能 SOTA. HF ID 是 laion/CLIP-ViT-g-14-laion2B-s34B-b88K.

### Open Clip 
