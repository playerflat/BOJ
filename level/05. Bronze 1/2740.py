import sys


def init():
    X, Y = map(int, sys.stdin.readline().split())
    matrix = [list(map(int, sys.stdin.readline().split())) for _ in range(X)]
    return matrix, X, Y


A, N, _ = init()
B, M, K = init()

for n in range(N):
    for i in range(K):
        temp = 0
        for j in range(M):
            temp += A[n][j] * B[j][i]
        print(temp, end=' ')
    print()
