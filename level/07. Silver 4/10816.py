from collections import defaultdict

a = int(input())
b = list(map(int, input().split()))
c = defaultdict(int)

for i in b:
    c[i] += 1

d = int(input())
e = list(map(int, input().split()))
for i in e:
    print(c.get(i, 0), end=' ')
