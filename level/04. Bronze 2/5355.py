for _ in range(int(input())):
    f = 1
    for j in input().split():
        if f:
            r = float(j)
            f = 0
        elif j == '@':
            r *= 3
        elif j == '%':
            r += 5
        elif j == '#':
            r -= 7
    print(f'{r:.2f}')
