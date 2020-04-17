import sys

while True:
    r = 0
    a = int(sys.stdin.readline())
    if a == 0:
        break
    else:
        r += len(str(a))
        for i in range(r):
            if a % 10 == 1:
                r += 2
            elif a % 10 == 0:
                r += 4
            else:
                r += 3
            a //= 10

        print(r + 1)
