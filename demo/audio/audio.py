import librosa
import sys
import librosa.display
import matplotlib.pyplot as plt
import numpy as np
import soundfile as sf
from moviepy import VideoFileClip
import pandas as pd
import cv2
from pydub import AudioSegment
import subprocess
import json
import os
import re
import pyloudnorm as pyln


def calculate_audio_rms(filepath):
    # 加载音频文件（会自动转为单声道）
    y, sr = librosa.load(filepath, sr=None)  # sr=None 表示保留原始采样率

    # 计算 RMS
    rms = np.sqrt(np.mean(y**2))
    return rms

def calculate_audio_std(filepath):
    # 加载音频文件（会自动转为单声道）
    y, sr = librosa.load(filepath, sr=None)  # sr=None 表示保留原始采样率

    # 计算 RMS
    std_dev = np.std(y)
    return std_dev

def getaudioloudness_overtime(filepath):
    """
    调用 ffmpeg 的 loudnorm 滤镜，获取每帧的响度值
    返回 (times, loudness_values)
    """
    cmd = [
        'ffmpeg',
        '-i', filepath,
        '-af', 'loudnorm=I=-23:LRA=11:TP=-1.5:print_format=json',
        '-f', 'null',
        '/dev/null'
    ]

    print("Running command:", ' '.join(cmd))
    result = subprocess.run(cmd, stderr=subprocess.PIPE, text=True)

    # 提取所有 {input_i: ...} 的块
    stderr_output = result.stderr
    print("result: ", result)

    # 匹配所有的 JSON 响度块
    pattern = r'\{.*?"measured_s":.*?\}'
    matches = re.findall(pattern, stderr_output, re.DOTALL)

    times = []
    loudness = []

    print("matches: ", matches)
    for match in matches:
        try:
            data = json.loads(match)
            times.append(data["measured_s"])         # 时间戳（秒）
            loudness.append(data["input_i"])          # 瞬时响度（LUFS）
        except Exception as e:
            print("解析错误:", e)

    print("len:", len(times), len(loudness))
    return np.array(times), np.array(loudness)


def pltaudioloudness_overtime_tmp(path, filename):
    """
    绘制响度随时间变化的折线图
    """
    times, loudness = getaudioloudness_overtime(path+filename)

    if len(times) == 0 or len(loudness) == 0:
        print("无法获取响度数据，请检查音频文件或 ffmpeg 版本")
        return

    plt.figure(figsize=(12, 6))
    plt.plot(times, loudness, label="Loudness (LUFS)", color='blue')
    plt.axhline(y=-23, color='r', linestyle='--', label="Target -23 LUFS")
    plt.title("Audio Loudness Over Time")
    plt.xlabel("Time (seconds)")
    plt.ylabel("Loudness (LUFS)")
    plt.grid(True)
    plt.legend()
    plt.tight_layout()
    # plt.show()
    plt.savefig(path + filename.split('.')[0] + "_loudness_spectrogram.png")

def pltaudioloudness_overtime(path, filename):

    # 加载音频文件
    filepath = path + filename  # 替换为你的音频路径
    data, rate = librosa.load(filepath, sr=None, mono=False)

    # 如果是立体声，则转为双通道 NumPy 数组
    if data.ndim == 1:
        data = np.expand_dims(data, axis=0)  # 单声道 -> (1, samples)
    else:
        data = data.T  # 确保 shape 是 (samples, channels)

    # 创建 Loudness Meter（符合 ITU-R BS.1770 标准）
    meter = pyln.Meter(rate)

    # 计算整体响度（用于参考）
    overall_loudness = meter.integrated_loudness(data)
    print(f"整体响度（Integrated Loudness）: {overall_loudness:.2f} LUFS")

    # 分帧计算瞬时响度（Momentary Loudness，每400ms一帧）
    window_size = int(0.4 * rate)  # 400ms
    hop_size = window_size // 2     # 50% 重叠

    momentary_loudness = []

    for start in range(0, data.shape[0], hop_size):
        end = start + window_size
        if end > data.shape[0]:
            break
        segment = data[start:end]
        loudness = meter.integrated_loudness(segment)
        if loudness == -np.inf:
            continue
        momentary_loudness.append(loudness)

    # 时间轴生成
    times = np.arange(len(momentary_loudness)) * hop_size / rate

    loudness_std = np.std(momentary_loudness)
    print(f"响度标准差 std: {loudness_std:.2f}")

    # 绘制响度图
    plt.figure(figsize=(12, 6))
    plt.plot(times, momentary_loudness, label="Momentary Loudness (LUFS), std: %.3f" % loudness_std, color='blue')
    plt.axhline(y=overall_loudness, color='red', linestyle='--', label=f"Overall Loudness: {overall_loudness:.2f} LUFS")

    plt.xlabel("time (s)")
    plt.ylabel("loudness (LUFS)")
    plt.title(filename.split('.')[0] + ' loudness')
    plt.grid(True)
    plt.legend()
    plt.tight_layout()
    plt.savefig(path + filename.split('.')[0] + "_loudness_timeline.png")


def audioloudness(filepath):
    """
    使用 ffmpeg 的 loudnorm 滤镜获取音频文件的响度（单位 LUFS）
    """
    cmd = [
        'ffmpeg',
        '-i', filepath,
        '-af', 'loudnorm=I=-23:LRA=11:TP=-1.5:print_format=json',
        '-f', 'null',
        '/dev/null'  # Linux/Mac 上用 /dev/null，Windows 上可用 NUL
    ]

    print("Running command:", ' '.join(cmd))
    result = subprocess.run(cmd, stderr=subprocess.PIPE, text=True)

    # 提取 JSON 输出
    output = result.stderr
    start_idx = output.find("{")
    end_idx = output.rfind("}") + 1
    if start_idx != -1 and end_idx != -1:
        json_str = output[start_idx:end_idx]
        try:
            loudness_data = json.loads(json_str)
            integrated_loudness = loudness_data.get('input_i')
            print("audio loudness: ", integrated_loudness)
            return integrated_loudness
        except Exception as e:
            print("解析 JSON 出错:", e)
    else:
        print("未找到响度数据")
    return None



def subaudio(path, filename):
    # 输入视频路径
    audio_path = path + filename
    # 输出视频路径
    output_path = path + "sub_" + filename

    audio = AudioSegment.from_file(audio_path)

    # 截取前 2 分钟（单位是毫秒）
    two_minutes = 105 * 1000
    first_two_minutes = audio[:two_minutes]  # 前 120 秒

    # 导出新文件
    first_two_minutes.export(output_path, format="wav")

    print(f"已保存前 2 分钟音频到：{output_path}")

def subvideo(path, filename):
    # 输入视频路径
    video_path = path + filename
    # 输出视频路径
    output_path = path + "sub_" + filename

    # 打开视频
    cap = cv2.VideoCapture(video_path)

    # 获取基本信息
    fps = int(cap.get(cv2.CAP_PROP_FPS))
    width = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
    height = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))

    # 设置前 1 分钟（60秒）
    duration_seconds = 105
    total_frames_to_save = fps * duration_seconds

    # 视频编码器和输出对象
    fourcc = cv2.VideoWriter_fourcc(*'mp4v')  # 或者 'XVID' 等
    out = cv2.VideoWriter(output_path, fourcc, fps, (width, height))

    frame_index = 0

    while cap.isOpened() and frame_index < total_frames_to_save:
        ret, frame = cap.read()
        if not ret:
            break

        out.write(frame)
        frame_index += 1

    # 释放资源
    cap.release()
    out.release()

    print(f"已保存前 {duration_seconds} 秒视频到：{output_path}")


def saveCSV(time, y, path, filename):
    # 创建一个DataFrame来存储时间和幅度数据
    df = pd.DataFrame({
        'Time (s)': time,
        'Amplitude': y
    })

    # 将DataFrame保存为.csv文件
    df.to_csv(path + filename + '_data.csv', index=False)

def analyAudio(path, filename):
    # 加载音频文件
    # '''
    fullname = filename + '.mp4'  # 替换为你的音频文件路径
    clip = VideoFileClip(path + fullname)
    audio = clip.audio.write_audiofile(path + filename + ".wav")
    # '''

    y, sr = librosa.load(path + filename + ".wav", sr=None)

    # 1. 存储波形图数据
    # 提取时间轴数据
    time = np.arange(0, len(y)) / sr
    # 打印一些基本信息
    print(f"采样率: {sr} samples/s")
    print(f"音频长度: {len(y)/sr:.2f} seconds")
    saveCSV(time, y, path, filename + "_waveform")


    # 1. 绘制波形图
    plt.figure(figsize=(14, 5))
    librosa.display.waveshow(y, sr=sr)
    plt.title(filename + ' waveform')
    plt.ylim(-0.4, 0.4)
    plt.savefig(path + filename + "_waveform.png")
    print("end waveform draw 1!")

    plt.title(filename + ' wide waveform')
    plt.ylim(-1, 1)
    plt.savefig(path + filename + "_wide_waveform.png")
    print("end waveform draw 2!")

    # 2.绘制频谱图
    # 计算短时傅里叶变换（STFT）
    D = np.abs(librosa.stft(y))
    # 将幅度转换为分贝单位（dB）。参考值设置为最大幅度值。
    DB = librosa.amplitude_to_db(D, ref=np.max)
    plt.figure(figsize=(14, 5))
    librosa.display.specshow(DB, sr=sr, x_axis='time', y_axis='log')
    plt.colorbar(format='%+2.0f dB')
    plt.title(filename + ' spectrogram')
    plt.savefig(path + filename + "_spectrogram.png")
    print("draw spectrogram done!")


def genAudio():
    # 音频参数
    duration = 30  # 持续时间，单位为秒
    sample_rate = 44100  # 采样率，每秒样本数
    frequency = 440  # 正弦波的频率，单位为赫兹(Hz)

    # 生成时间轴
    t = np.linspace(0, duration, int(sample_rate * duration), endpoint=False)

    # 生成440Hz的正弦波
    waveform = np.sin(2 * np.pi * frequency * t)

    # 将波形数据标准化到-1.0到1.0之间
    waveform /= np.max(np.abs(waveform))

    # 导出为WAV文件
    sf.write('440hz_tone.wav', waveform, sample_rate)

    print("440Hz音调的WAV文件已生成")


if __name__ == "__main__":
    # genAudio()

    func = ""
    if len(sys.argv) > 1:
        func = sys.argv[1]

    if func == "analy":
        if len(sys.argv) <= 3:
            print("less argv, func '%s' need [path] [filename]\n" % func)
            sys.exit(2)
        path = sys.argv[2]
        filename = sys.argv[3]
        print("arg: ", path, filename)
        analyAudio(path, filename)
    elif func == "subaudio":
        if len(sys.argv) <= 3:
            print("less argv, func '%s' need [path] [filename]\n" % func)
            sys.exit(2)
        path = sys.argv[2]
        filename = sys.argv[3]
        subaudio(path, filename)
    elif func == "audioloudness":
        if len(sys.argv) <= 2:
            print("less argv, func '%s' need [fullpath]\n" % func)
            sys.exit(2)
        fullfilepath = sys.argv[2]
        audioloudness(fullfilepath)
    elif func == "pltloudness":
        if len(sys.argv) <= 3:
            print("less argv, func '%s' need [path] [filename]\n" % func)
            sys.exit(2)
        path = sys.argv[2]
        filename = sys.argv[3]
        print("arg: ", path, filename)
        pltaudioloudness_overtime(path, filename)
    elif func == "rms":
        if len(sys.argv) <= 2:
            print("less argv, func '%s' need [fullpath]\n" % func)
            sys.exit(2)
        fullfilepath = sys.argv[2]
        rms = calculate_audio_rms(fullfilepath)
        print("file '%s' rms is %f" % (fullfilepath, rms))
    elif func == "std":
        if len(sys.argv) <= 2:
            print("less argv, func '%s' need [fullpath]\n" % func)
            sys.exit(2)
        fullfilepath = sys.argv[2]
        std_dev = calculate_audio_std(fullfilepath)
        print("file '%s' std is %f" % (fullfilepath, std_dev))
    else:
        print("func '%s' not found \n" % func)


