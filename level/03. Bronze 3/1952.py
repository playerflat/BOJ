m, n = map(int, input().split())
c = 0
while 1:
    c += 1
    m -= 1
    if n == 1:
        break
    c += 1
    n -= 1
    if m == 1:
        break
print(c)
