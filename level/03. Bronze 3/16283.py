a, b, n, w = map(int, input().split())
r = []
for i in range(1, n):
    for j in range(1, n - i + 1):
        if i * a + j * b > w:
            break
        if i * a + j * b == w and i + j == n:
            r.append(i)
            r.append(j)
if len(r) == 2:
    print(r[0], r[1])
else:
    print(-1)
