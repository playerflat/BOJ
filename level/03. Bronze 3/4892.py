c = 1
while 1:
    n = int(input())
    if not n:
        break
    n *= 3
    if n % 2 == 0:
        print(f'{c}. even', end=" ")
        n = n // 2
        f = 1
    else:
        print(f'{c}. odd', end=" ")
        n = (n + 1) // 2
    n = 3 * n // 9
    print(n)
    c += 1
