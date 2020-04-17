b1, b2 = map(int, input().split())
d1, d2 = map(int, input().split())
j1, j2 = map(int, input().split())

b = max(abs(j1 - b1), abs(j2 - b2))
d = abs((j1 - d1)) + abs((j2 - d2))

if b > d:
    print("daisy")
elif d > b:
    print("bessie")
else:
    print("tie")
