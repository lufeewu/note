import librosa
import sys
import librosa.display
import matplotlib.pyplot as plt
import numpy as np
import soundfile as sf
from moviepy import VideoFileClip

def analyAudio(path, filename):
    # 加载音频文件
    fullname = filename + '.mp4'  # 替换为你的音频文件路径
    clip = VideoFileClip(path + fullname)
    audio = clip.audio.write_audiofile(path + filename + ".wav")

    y, sr = librosa.load(path + filename + ".wav", sr=None)

    # 1. 绘制波形图
    plt.figure(figsize=(14, 5))
    librosa.display.waveshow(y, sr=sr)
    plt.title(filename + ' waveform')
    plt.ylim(-0.3, 0.3)
    plt.savefig(path + filename + "_waveform.png")
    print("end waveform draw!")

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

    if len(sys.argv) > 2:
        path = sys.argv[1]
        filename = sys.argv[2]
        print("arg: ", path, filename)
        analyAudio(path, filename)
