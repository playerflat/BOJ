import collections

n, p = map(int, input().split())

d = collections.defaultdict(int)
m = n

while 1:
    d[n] += 1
    n = n * m % p
    if d[n] == 3:
        break

r = 0
for v in d.values():
    if v >= 2:
        r += 1
print(r)
