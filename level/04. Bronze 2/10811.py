n, m = map(int, input().split())
a = [i for i in range(n + 1)]
for _ in range(m):
    i, j = map(int, input().split())
    for k in range((j - i + 1) // 2):
        a[i + k], a[j - k] = a[j - k], a[i + k]
print(' '.join(map(str, a))[2:])
