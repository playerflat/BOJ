max = 0
c = 0
for i in range(10):
    a, b = map(int, input().split())
    c -= a
    c += b
    if c > max:
        max = c

print(max)
