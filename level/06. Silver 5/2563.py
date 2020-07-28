a = []

for i in range(100):
    a.append([0 for j in range(100)])

for _ in range(int(input())):
    x, y = map(int, input().split())
    for i in range(x, x + 10):
        for j in range(y, y + 10):
            a[i][j] = 1

r = 0
for i in a:
    r += sum(i)

print(r)