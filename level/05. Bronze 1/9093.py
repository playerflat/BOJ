import sys

for _ in range(int(sys.stdin.readline())):
    a = sys.stdin.readline().split()
    for i in a:
        print(i[::-1], end=' ')
    print()
