while 1:
    a, b, c = map(int, input().split())
    if a == b == c == 0:
        break
    if b - a == c - b:
        print(f'AP {c + (b - a)}')
    else:
        print(f'GP {c * (b // a)}')
