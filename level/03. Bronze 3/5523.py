import sys

ra, rb = 0, 0
for i in range(int(sys.stdin.readline())):
    a, b = map(int, sys.stdin.readline().split())
    if a > b:
        ra += 1
    elif a < b:
        rb += 1

print(ra,rb)