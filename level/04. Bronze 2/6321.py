for i in range(1, int(input()) + 1):
    print(f'String #{i}')
    for j in input():
        if j == 'Z':
            print('A', end='')
        else:
            print(chr(ord(j) + 1), end='')
    print()
    print()
