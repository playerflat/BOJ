import sys
a = int(sys.stdin.readline())
b = [0] * 10001

for i in range(a):
    c = int(sys.stdin.readline())
    b[c] += 1

for i in range(1, len(b)):
    if b[i] != 0:
        for j in range(b[i]):
            print(i)