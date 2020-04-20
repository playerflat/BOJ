a, b = map(int, input().split())
c, d = map(int, input().split())
r = {b / a + d / c: 3, d / b + c / a: 2, c / d + a / b: 1, a / c + b / d: 0}
print(r[max(r)])
