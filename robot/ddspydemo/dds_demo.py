#!/usr/bin/env python3
"""Simple DDS pub/sub demo using Cyclone DDS Python API.

Usage:
  # Terminal 1
  python robot/dds_demo.py sub --topic DemoTopic

  # Terminal 2
  python robot/dds_demo.py pub --topic DemoTopic --count 10 --interval 1.0

Dependency:
  pip install cyclonedds
"""

from __future__ import annotations

import argparse
import signal
import sys
import time
from typing import Optional

from cyclonedds.domain import DomainParticipant
from cyclonedds.pub import DataWriter, Publisher
from cyclonedds.sub import DataReader, Subscriber
from cyclonedds.topic import Topic
from cyclonedds.util import duration

from dds_types import ChatMessage


RUNNING = True


def handle_signal(signum: int, _frame: Optional[object]) -> None:
    del signum
    global RUNNING
    RUNNING = False


def run_publisher(participant: DomainParticipant, topic_name: str, count: int, interval: float) -> None:
    topic = Topic(participant, topic_name, ChatMessage)
    pub = Publisher(participant)
    writer = DataWriter(pub, topic)

    print(f"[PUB] domain={participant.domain_id}, topic={topic_name}, count={count}, interval={interval}s")
    sent = 0
    while RUNNING and (count < 0 or sent < count):
        msg = ChatMessage(index=sent, message=sent % 10)
        writer.write(msg)
        print(f"[PUB] sent: index={msg.index}, message={msg.message}")
        sent += 1
        time.sleep(interval)

    print("[PUB] stopped")


def run_subscriber(participant: DomainParticipant, topic_name: str) -> None:
    topic = Topic(participant, topic_name, ChatMessage)
    sub = Subscriber(participant)
    reader = DataReader(sub, topic)

    print(f"[SUB] domain={participant.domain_id}, topic={topic_name}, waiting messages...")

    while RUNNING:
        got_data = False
        for msg in reader.take_iter(timeout=duration(milliseconds=500)):
            got_data = True
            print(f"[SUB] recv: index={msg.index}, message={msg.message}")

        if not got_data:
            print("[SUB] ...")

    print("[SUB] stopped")


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description="DDS pub/sub demo")
    parser.add_argument("mode", choices=["pub", "sub"], help="Run as publisher or subscriber")
    parser.add_argument("--domain", type=int, default=0, help="DDS domain id")
    parser.add_argument("--topic", default="DemoTopic", help="Topic name")
    parser.add_argument("--count", type=int, default=10, help="Publish message count, -1 for infinite")
    parser.add_argument("--interval", type=float, default=1.0, help="Publish interval in seconds")
    return parser


def main() -> int:
    signal.signal(signal.SIGINT, handle_signal)
    signal.signal(signal.SIGTERM, handle_signal)

    args = build_parser().parse_args()

    participant = DomainParticipant(args.domain)

    if args.mode == "pub":
        run_publisher(participant, args.topic, args.count, args.interval)
    else:
        run_subscriber(participant, args.topic)

    return 0


if __name__ == "__main__":
    sys.exit(main())
