import sys

a = [2 ** i for i in range(0, 32)]

for i in range(int(sys.stdin.readline())):
    b = int(sys.stdin.readline())
    if b in a:
        print(1)
    else:
        print(0)
