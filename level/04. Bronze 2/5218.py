for _ in range(int(input())):
    a, b = input().split()
    print("Distances: ", end='')
    for i in range(len(a)):
        if ord(a[i]) <= ord(b[i]):
            print(ord(b[i]) - ord(a[i]), end=' ')
        else:
            print(abs(ord(b[i]) - ord(a[i]) + 26), end=' ')
    print()