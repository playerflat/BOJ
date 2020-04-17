import sys

a, b = map(int, sys.stdin.readline().split())
print(int(((max(a, b) - min(a, b) + 1) * (a + b)) / 2))
