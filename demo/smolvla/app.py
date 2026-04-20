import os
import sys
import importlib
from pathlib import Path

CURRENT_DIR = Path(__file__).resolve().parent
LOCAL_LEROBOT_SRC = CURRENT_DIR / "lerobot" / "src"

if LOCAL_LEROBOT_SRC.exists():
    sys.path.insert(0, str(LOCAL_LEROBOT_SRC))

MODEL_ID = os.getenv("SMOLVLA_MODEL_ID", "lerobot/smolvla_base")

_policy = None


def get_runtime_modules():
    try:
        gradio = importlib.import_module("gradio")
        torch = importlib.import_module("torch")
        image_module = importlib.import_module("PIL.Image")
        smolvla_module = importlib.import_module("lerobot.policies.smolvla.modeling_smolvla")
    except ImportError as exc:
        raise RuntimeError(
            "缺少运行依赖，请先安装 gradio、torch、Pillow，以及包含 SmolVLA 的 lerobot 环境。"
        ) from exc

    return gradio, torch, image_module, smolvla_module.SmolVLAPolicy


def get_device() -> str:
    _, torch, _, _ = get_runtime_modules()
    return "cuda" if torch.cuda.is_available() else "cpu"


def load_policy():
    global _policy
    if _policy is None:
        _, torch, _, smolvla_policy = get_runtime_modules()
        device = torch.device(get_device())
        _policy = smolvla_policy.from_pretrained(MODEL_ID).to(device).eval()
    return _policy


def build_batch(policy, image, instruction: str, device):
    """把 PIL 图像和指令字符串组装成 select_action 所需的 batch dict。"""
    _, torch, _, _ = get_runtime_modules()

    # ---- 图像 ----
    try:
        import torchvision.transforms.functional as TF
        img_tensor = TF.to_tensor(image.convert("RGB"))
    except ImportError:
        import numpy as np
        np_img = np.array(image.convert("RGB"), dtype=np.float32) / 255.0  # (H,W,3)
        img_tensor = torch.from_numpy(np_img).permute(2, 0, 1)  # (3,H,W)
    img_tensor = img_tensor.unsqueeze(0).to(device)  # (1,3,H,W) in [0,1]
    img_keys = list(policy.config.image_features.keys())
    first_img_key = img_keys[0] if img_keys else "observation.images.image"

    # ---- 语言 token ----
    processor = policy.model.vlm_with_expert.processor
    tokenizer = getattr(processor, "tokenizer", processor)
    token_out = tokenizer(
        instruction,
        return_tensors="pt",
        padding="max_length",
        truncation=True,
        max_length=policy.config.tokenizer_max_length,
    )
    lang_tokens = token_out["input_ids"].to(device)
    lang_mask = token_out["attention_mask"].bool().to(device)

    # ---- robot state（推理 Demo 用全零） ----
    state_feat = policy.config.robot_state_feature
    state_dim = state_feat.shape[0] if state_feat is not None else 1
    state = torch.zeros(1, state_dim, device=device)

    batch = {
        first_img_key: img_tensor,
        "observation.language.tokens": lang_tokens,
        "observation.language.attention_mask": lang_mask,
        "observation.state": state,
    }
    return batch


def action_to_description(action_tensor) -> str:
    """把归一化的动作向量转换成自然语言描述。
    
    smolvla_base 的 action 是 6D 向量，对应 SO-100 机械臂的 6 个关节。
    """
    action = action_tensor.squeeze().cpu().numpy() if hasattr(action_tensor, 'numpy') else action_tensor
    if action.ndim > 1:
        action = action[0]  # 取 batch 中的第一个

    joint_names = [
        "肩部旋转",
        "肩部升降",
        "肘部弯曲",
        "腕部屈伸",
        "腕部旋转",
        "夹爪开合",
    ]

    descriptions = []
    for i, (name, value) in enumerate(zip(joint_names, action)):
        # 根据值的正负和大小生成自然语言
        abs_value = abs(value)
        if abs_value < 0.1:
            direction = "保持"
        elif i == 5:  # 夹爪特殊处理
            direction = "关闭中" if value < 0 else "打开中"
        else:
            direction = "↑" if value > 0 else "↓"

        # 格式化数值（保留两位小数）
        descriptions.append(f"{name} {direction} ({value:+.4f})")

    nl_text = " | ".join(descriptions)
    return nl_text


def predict(image, instruction: str) -> str:
    if image is None:
        return "请先上传一张图片。"

    instruction = instruction.strip()
    if not instruction:
        return "请先输入控制指令。"

    try:
        _, torch, _, _ = get_runtime_modules()
        device = torch.device(get_device())
        policy = load_policy()
        policy.reset()

        batch = build_batch(policy, image, instruction, device)
        with torch.inference_mode():
            action = policy.select_action(batch)

        # 生成自然语言描述
        nl_description = action_to_description(action)
        
        # 返回数值 + 自然语言描述
        result = f"原始动作向量：{str(action.cpu().numpy())}\n\n"
        result += f"自然语言描述：\n{nl_description}"
        return result
    except Exception as exc:
        return f"推理失败: {exc}"


def build_demo():
    gr, _, _, _ = get_runtime_modules()

    with gr.Blocks(title="SmolVLA Web UI") as demo:
        gr.Markdown("# SmolVLA Web UI")
        gr.Markdown(
            f"当前模型: `{MODEL_ID}`  |  运行设备: `{get_device()}`\n\n"
            "上传机器人观测图像并输入自然语言指令，界面会返回模型预测动作。"
        )

        with gr.Row():
            image_input = gr.Image(type="pil", label="观测图像")
            with gr.Column():
                instruction_input = gr.Textbox(
                    label="控制指令",
                    placeholder="例如: pick up the red cube",
                    lines=3,
                )
                submit_button = gr.Button("执行推理", variant="primary")
                clear_button = gr.Button("清空")

        result_output = gr.Textbox(label="预测动作", lines=8)

        submit_button.click(
            fn=predict,
            inputs=[image_input, instruction_input],
            outputs=result_output,
        )
        instruction_input.submit(
            fn=predict,
            inputs=[image_input, instruction_input],
            outputs=result_output,
        )
        clear_button.click(
            fn=lambda: (None, "", ""),
            inputs=None,
            outputs=[image_input, instruction_input, result_output],
        )

        gr.Examples(
            examples=[
                [None, "pick up the red cube"],
                [None, "move the gripper to the cup"],
            ],
            inputs=[image_input, instruction_input],
        )

    return demo


if __name__ == "__main__":
    demo = build_demo()
    demo.launch(
        server_name=os.getenv("SMOLVLA_HOST", "0.0.0.0"),
        server_port=int(os.getenv("SMOLVLA_PORT", "7860")),
        debug=True,
        show_error=True,
        share=False,
    )