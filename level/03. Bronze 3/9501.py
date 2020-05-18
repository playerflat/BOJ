import sys

a = lambda: sys.stdin.readline()
for _ in range(int(a())):
    r = 0
    b = list(map(int, a().split()))
    for _ in range(b[0]):
        c = list(map(int, a().split()))
        if c[1] / c[2] * c[0] >= b[1]:
            r += 1
    print(r)
