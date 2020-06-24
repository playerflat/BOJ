import sys

k, n = map(int, sys.stdin.readline().split())
a = []
for _ in range(k):
    a.append(int(sys.stdin.readline()))

right = sum(a) // n
left = 0
while 1:
    mid = (right + left) // 2
    if right == left == mid:
        print(mid)
        break
    elif (sum(i // right for i in a) < n and sum(i // mid for i in a) < n) or (
            left == mid and sum(i // right for i in a) < n):
        right = mid
    elif sum(i // right for i in a) >= n:
        left = right
    else:
        left = mid
