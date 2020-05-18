import sys

a = int(sys.stdin.readline())
b = list(map(int, sys.stdin.readline().split()))
c = list(map(int, sys.stdin.readline().split()))
r = 0

for i in range(a):
    if b[i] <= c[0]:
        r += 1
    else:
        r += (b[i] - c[0]) // c[1] + (1 if (b[i] - c[0]) % c[1] else 0) + 1

print(r)
