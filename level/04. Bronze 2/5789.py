for i in range(int(input())):
    a = input()
    if a[len(a) // 2 - 1] == a[len(a) // 2]:
        print('Do-it')
    else:
        print('Do-it-Not')
