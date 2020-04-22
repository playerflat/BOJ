import sys

x, y, z = 1, 1, 0
for _ in range(int(sys.stdin.readline())):
    a = list(map(int, sys.stdin.readline().split()))
    x *= a[0]
    y *= a[1]
    z += a[2]

print(z % 2, y // x)
