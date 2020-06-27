c = 0
b = 1
n, m, l = map(int, input().split())
d = {i: 0 for i in range(1, n + 1)}
while 1:
    d[b] += 1
    if d[b] == m:
        break
    if d[b] % 2 == 0:
        if b - l <= 0:
            b = n - (l - b)
        else:
            b -= l
    else:
        if b + l > n:
            b = l - (n - b)
        else:
            b += l
    c += 1
print(c)
