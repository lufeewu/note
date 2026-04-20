import torch
from lerobot.datasets.lerobot_dataset import LeRobotDataset
from lerobot.policies.factory import make_pre_post_processors

# Swap this import per-policy
from lerobot.policies.smolvla.modeling_smolvla import SmolVLAPolicy

# load a policy
model_id = "lerobot/smolvla_base"  # <- swap checkpoint
device = torch.device("cuda" if torch.cuda.is_available() else "cpu")

policy = SmolVLAPolicy.from_pretrained(model_id).to(device).eval()

preprocess, postprocess = make_pre_post_processors(
    policy.config,
    model_id,
    preprocessor_overrides={"device_processor": {"device": str(device)}},
)
# load a lerobotdataset
dataset = LeRobotDataset("lerobot/libero")

# pick an episode
episode_index = 0

# each episode corresponds to a contiguous range of frame indices
from_idx = dataset.meta.episodes["dataset_from_index"][episode_index]
to_idx   = dataset.meta.episodes["dataset_to_index"][episode_index]

# get a single frame from that episode (e.g. the first frame)
frame_index = from_idx
frame = dict(dataset[frame_index])

batch = preprocess(frame)

# Some datasets expose images as observation.images.image/image2, while
# SmolVLA checkpoints can expect camera1/camera2/camera3 feature names.
expected_image_keys = [
    key
    for key, feature in policy.config.input_features.items()
    if getattr(getattr(feature, "type", None), "value", None) == "VISUAL"
]
available_image_keys = [key for key in batch.keys() if key.startswith("observation.images.")]

if expected_image_keys and not any(key in batch for key in expected_image_keys):
    for src_key, dst_key in zip(sorted(available_image_keys), sorted(expected_image_keys)):
        batch[dst_key] = batch[src_key]

with torch.inference_mode():
    pred_action = policy.select_action(batch)
    # use your policy postprocess, this post process the action
    # for instance unnormalize the actions, detokenize it etc..
    pred_action = postprocess(pred_action)
