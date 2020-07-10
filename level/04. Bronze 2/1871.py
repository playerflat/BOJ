for _ in range(int(input())):
    a, b = input().split('-')
    f = 0
    for i in range(3):
        f += (ord(a[i]) - 65) * 26 ** (2 - i)
    if abs(f - int(b)) <= 100:
        print('nice')
    else:
        print('not nice')
