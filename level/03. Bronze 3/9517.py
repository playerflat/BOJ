while 1:
    s = int(input())
    t = 210
    f = False
    for i in range(int(input())):
        a, b = input().split()
        t -= int(a)
        if t < 0:
            print(s)
            f = True
            break
        if b == "T":
            s = s % 8 + 1
        else:
            continue
    if f:
        break
