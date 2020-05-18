import sys

sys.stdin.readline()
b = list(map(int, sys.stdin.readline().split()))
c = 0
r = 0

for i in b:
    if i == c:
        r += 1
        c += 1
        c %= 3

print(r)