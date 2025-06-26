import numpy as np
from numpy.random import randn

def softmax(x):
    # print(x)
    e_x = np.exp(x - np.max(x))
    # result = e_x / np.sum(e_x, axis=-1, keepdims=True)
    result = e_x / e_x.sum(axis=0)
    return result

def attention(x):
    d, n = x.shape

    Wq = randn(d, d)
    Wk = randn(d, d)
    Wv = randn(d, d)

    q = Wq @ x
    k = Wk @ x
    v = Wv @ x

    A = k.T @ q
    A = A / np.sqrt(d)
    A_hat = softmax(A)
    output = v @ A_hat
    print(output.shape) # n, d


def multi_head_attention(x, head_n=16):
    n, d = x.shape
    assert d % head_n == 0
    Wq = np.random.rand(d, d)
    Wk = np.random.rand(d, d)
    Wv = np.random.rand(d, d)
    q = x @ Wq
    k = x @ Wk
    v = x @ Wv
    q = np.reshape(q, (n, head_n, d // head_n))
    k = np.reshape(k, (n, head_n, d // head_n))
    v = np.reshape(v, (n, head_n, d // head_n))
    q = np.transpose(q, (1, 0, 2))  # head_n, n, d // head_n
    k = np.transpose(k, (1, 0, 2))
    v = np.transpose(v, (1, 0, 2))
    A = q @ np.transpose(k, (0, 2, 1))
    A = A / np.sqrt(d // head_n)
    A_hat = softmax(A) # head_n, n, n
    output = A_hat @ v # head_n, n, d // head_n
    output = np.transpose(output, (1, 0, 2))    # n, head_n, d // head_n
    output = np.reshape(output, (n, d)) 
    print(output.shape) # n, d
    

if __name__ == "__main__":
    attention(randn(256, 32))
    multi_head_attention(np.random.rand(512, 768))