import sys

a = int(sys.stdin.readline())
b = list(map(int, sys.stdin.readline().split()))
r = 0
for i in range(a):
    if i + 1 != b[i]:
        r += 1
print(r)