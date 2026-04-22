"""DDS message types for demo."""

from __future__ import annotations

from dataclasses import dataclass

from cyclonedds.idl import IdlStruct
from cyclonedds.idl.types import int32


@dataclass
class ChatMessage(IdlStruct, typename="robot.dds.ChatMessage"):
    """A simple chat message type."""

    index: int32
    message: int32  # 0-9 maps to message 0-9
