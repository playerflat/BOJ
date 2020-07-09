a, b = map(int, input().split())
c = [0 * i for i in range(b + 1)]
for i in range(a):
    d = int(input())
    if d == 1:
        c = [b]
        break
    for j in range(d, len(c), d):
        c[j] = 1

print(sum(c))
