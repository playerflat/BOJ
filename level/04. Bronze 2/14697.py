a, b, c, d = map(int, input().split())
e = [0] * 350
e[0] = 1
for i in range(d):
    if e[i]:
        e[i + a] = e[i + b] = e[i + c] = 1
print(e[d])
