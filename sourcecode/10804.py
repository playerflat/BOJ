import math

a = [str(i) for i in range(21)]

for _ in range(10):
    x, y = map(int, input().split())
    for i in range(math.ceil((y - x) / 2)):
        a[x + i], a[y - i] = a[y - i], a[x + i]

print(' '.join(a[1:]))
