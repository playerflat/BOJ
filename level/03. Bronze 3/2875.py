n, m, k = map(int, input().split())
nm = n + m
i = min(n // 2, m)
while True:
    if nm - k >= 3*i:
        print(i)
        break
    else:
        i -= 1
