a, b, c = map(int, input().split())
if a == b == c:
    print(10000 + a * 1000)
elif a == b or b == c or a == c:
    print(1000 + sorted([a, b, c])[1] * 100)
else:
    print(max(a, b, c) * 100)
