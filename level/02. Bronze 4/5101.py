while True:
    a, b, c = map(int, input().split())
    if a == b == c == 0:
        break
    else:
        d = (c - a) / b
        if d - int(d) != 0 or d < 0:
            print("X")
        elif a + b * int(d) == c:
            print(abs(int(d) + 1))