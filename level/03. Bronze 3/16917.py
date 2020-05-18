a, b, c, x, y = map(int, input().split())
r = 0

if a + b > c * 2:
    r += min(x, y) * c * 2
    if x > y:
        r += min((x - y) * a, (x - y) * c * 2)
    else:
        r += min((y - x) * b, (y - x) * c * 2)
else:
    r += a * x + b * y

print(r)
