a, b = map(int, input().split())
c = [str(i) for i in range(1, a + 1)]
for _ in range(b):
    x, y = map(int, input().split())
    c[x - 1], c[y - 1] = c[y - 1], c[x - 1]

print(" ".join(c))
