import sys

input()
a = list(map(int, sys.stdin.readline().split()))

print(max(a) - min(a))