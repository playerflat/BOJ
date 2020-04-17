r = 0
a = int(input())
b = list(map(int, input().split()))
for i in range(a - 1):
    r += b[i] * b[i + 1]
    b[i + 1] = b[i] + b[i + 1]

print(r)
