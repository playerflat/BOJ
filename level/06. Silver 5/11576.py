import sys

A, B = map(int, sys.stdin.readline().split())
m = int(sys.stdin.readline())
number = list(map(int, sys.stdin.readline().split()))
baseA = 0
baseB = []
digit = m - 1
for i in range(m):
    baseA += number[i] * (A ** digit)
    digit -= 1

while baseA > 0:
    baseB.append(baseA % B)
    baseA //= B

for i in range(len(baseB) - 1, -1, -1):
    print(baseB[i], end=' ')
