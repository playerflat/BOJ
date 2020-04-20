a = 0
b = list(map(int, input().split()))
for i in range(1, b[0] + 1):
    a += (b[1] * i) + (b[2] * (i ** 2))
print(a)
