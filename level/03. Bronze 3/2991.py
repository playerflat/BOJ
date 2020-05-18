def calc(v, x, y):
    if 0 < v % (x + y) <= x:
        return 1
    else:
        return 0


a, b, c, d = map(int, input().split())
p, m, n = map(int, input().split())

print(calc(p, a, b) + calc(p, c, d))
print(calc(m, a, b) + calc(m, c, d))
print(calc(n, a, b) + calc(n, c, d))
