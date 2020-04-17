a, b = map(int, input().split())
c = list(map(int, input().split()))

res = 0

for i in range(len(c) - 2):
    for j in range(i + 1, len(c) - 1):
        for k in range(j + 1, len(c)):
            if b >= c[i] + c[j] + c[k] > res:
                res = c[i] + c[j] + c[k]

print(res)