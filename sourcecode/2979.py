import collections

d = collections.defaultdict(int)
r = 0
a, b, c = map(int, input().split())
for _ in range(3):
    x, y = map(int, input().split())
    for i in range(x, y):
        d[i] += 1

for v in d.values():
    if v == 1:
        r += a
    elif v == 2:
        r += b * 2
    else:
        r += c * 3

print(r)