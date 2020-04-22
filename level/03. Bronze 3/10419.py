import sys

for _ in range(int(sys.stdin.readline())):
    a = int(sys.stdin.readline())
    for i in range(999):
        if i**2+i>a:
            print(i-1)
            break