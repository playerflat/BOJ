for _ in range(int(input())):
    a = int(input())
    b = []
    print(f'Pairs for {a}:', end=' ')
    for i in range(1, a // 2 if a % 2 == 0 else a // 2 + 1):
        b.append(f"{i} {a - i}")
    for i in range(len(b)):
        print(f'{b[i]}', end='')
        if b[i] != b[-1]:
            print(', ', end='')
    print()
