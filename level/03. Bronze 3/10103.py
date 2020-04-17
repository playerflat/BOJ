import sys

ra, rb = 100, 100
for _ in range(int(sys.stdin.readline())):
    a, b = map(int, sys.stdin.readline().split())
    if a > b:
        rb -= a
    elif b > a:
        ra -= b
print(ra)
print(rb)
