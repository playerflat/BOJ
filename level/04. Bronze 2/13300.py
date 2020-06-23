import collections

n, k = map(int, input().split())
r = collections.defaultdict(int)
t = 0
for _ in range(n):
    a = input().replace(' ', '')
    r[a] += 1

for v in r.values():
    if v % k == 0:
        t += v // k
    else:
        t += v // k + 1

print(t)
