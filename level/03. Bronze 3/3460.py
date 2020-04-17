for _ in range(int(input())):
    a = int(input())
    a = bin(a)
    for i in range(len(a) - 1, 1, -1):
        if a[i] == "1":
            print(len(a)-1 - i, end=' ')