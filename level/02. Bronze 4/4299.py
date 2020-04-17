a, b = map(int, input().split())
if a < b or (a + b) % 2 != 0:
    print("-1")
else:
    c, d = (a + b) / 2, (a + b) / 2 - b
    print(int(c), int(d))