import sys

for _ in range(3):
    r = 0
    for i in range(int(sys.stdin.readline())):
        r += int(sys.stdin.readline())
    if r > 0:
        print("+")
    elif r < 0:
        print("-")
    else:
        print(0)