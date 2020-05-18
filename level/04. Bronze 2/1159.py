import sys
import collections

d = collections.defaultdict(int)
r = []
for _ in range(int(input())):
    d[sys.stdin.readline().strip()[0]] += 1

for i, j in d.items():
    if j >= 5:
        r.append(i)

r.sort()
if len(r) == 0:
    print("PREDAJA")
else:
    print(''.join(r))
