import sys

N, M = map(int, sys.stdin.readline().split())
A = list(map(int, sys.stdin.readline().split()))
B = list(map(int, sys.stdin.readline().split()))

i, j = 0, 0

for _ in range(N + M):
    if len(A) == i:
        print(B[j])
        j += 1
    elif len(B) == j:
        print(A[i])
        i += 1
    elif A[i] < B[j]:
        print(A[i], end=' ')
        i += 1
    else:
        print(B[j], end=' ')
        j += 1
