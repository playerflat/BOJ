import sys

d = {}
for _ in range(int(sys.stdin.readline())):
    a, b = sys.stdin.readline().split()
    if b == 'enter':
        d[a] = 1
    else:
        d[a] = 0

for i, v in sorted(d.items(), key=lambda x: x[0], reverse=True):
    if v == 1:
        print(i)
