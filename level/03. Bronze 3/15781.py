import sys
sys.stdin.readline().split()
b = list(map(int, sys.stdin.readline().split()))
c = list(map(int, sys.stdin.readline().split()))
print(max(b) + max(c))
