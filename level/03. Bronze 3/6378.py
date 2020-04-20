while 1:
    a = int(input())
    if not a:
        break
    while 1:
        if a < 10:
            print(a)
            break
        a = a % 10 + a // 10 % 10 + a // 100 % 10 + a // 1000
