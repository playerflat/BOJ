import sys

s1, s2 = map(int, sys.stdin.readline().split())
f = 1
for _ in range(s1):
    if len(set(sys.stdin.readline().split())) == 2:
        f = 0
        print("Wrong Answer")
        break
if f:
    for _ in range(s2):
        if len(set(sys.stdin.readline().split())) == 2:
            f = 0
            print("Why Wrong!!!")
            break
if f:
    print("Accepted")
