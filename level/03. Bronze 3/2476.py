r = []
for _ in range(int(input())):
    a, b, c = map(int, input().split())
    if a == b == c:
        r.append(10000 + a * 1000)
    elif a == b or b == c or a == c:
        r.append(1000 + (sum([a, b, c]) - sum({a, b, c})) * 100)
    else:
        r.append(max(a, b, c) * 100)
print(max(r))
