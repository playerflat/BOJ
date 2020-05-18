while 1:
    a, b, c, d = map(int, input().split())
    if a == b == c == d == 0:
        break
    print(abs(b - c), d - a)
