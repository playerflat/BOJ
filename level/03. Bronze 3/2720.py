for _ in range(int(input())):
    a = int(input())
    print(a // 25, end=' ')
    a %= 25
    print(a // 10, end=' ')
    a %= 10
    print(a // 5, end=' ')
    a %= 5
    print(a)
