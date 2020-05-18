import sys

print(sum(1 for i in map(int, sys.stdin.readline().split()) if i > 0))
