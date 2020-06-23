n, m = map(int, input().split())
r = n
while 1:
    if n < m:
        break
    r += n // m
    n //= m
print(r)
