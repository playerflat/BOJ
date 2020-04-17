import sys

for _ in range(int(sys.stdin.readline())):
    a = list(map(int, sys.stdin.readline().split()))
    b = []
    for i in range(7):
        if a[i] % 2 == 0:
            b.append(a[i])
    print(sum(b), min(b))
