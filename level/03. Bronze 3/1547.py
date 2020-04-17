c = [1, 2, 3]
for _ in range(int(input())):
    a, b = map(int, input().split())
    ac = c.index(a)
    bc = c.index(b)
    c[ac], c[bc] = c[bc], c[ac]
print(c[0])