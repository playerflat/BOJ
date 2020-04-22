import sys

a, r = int(sys.stdin.readline()), int(sys.stdin.readline())
m = r
for i in range(a):
    x, y = map(int, sys.stdin.readline().split())
    r += x - y
    if r > m:
        m = r
    if r < 0:
        print(0)
        break
    if i == a - 1:
        print(m)
