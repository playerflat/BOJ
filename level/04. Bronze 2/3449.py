for _ in range(int(input())):
    r = 0
    a = input()
    b = input()
    for i in range(len(a)):
        if a[i] != b[i]:
            r+=1
    print(f'Hamming distance is {r}.')