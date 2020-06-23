import collections

a = input()
b = input()

d = collections.defaultdict(int)
r = 0

for i in a:
    d[i] += 1
for i in b:
    d[i] -= 1

for v in d.values():
    r += abs(v)

print(r)