for _ in range(int(input())):
    m, n = map(int, input().split())
    b = []
    c = 0
    for _ in range(m):
        b.append(list(map(int, input().split())))
    while 1:
        temp = 0
        for i in range(m - 2, -1, -1):
            for j in range(n - 1, -1, -1):
                if b[i][j] == 1 and b[i + 1][j] == 0:
                    temp += 1
                    b[i][j], b[i + 1][j] = 0, 1
        c += temp
        if not temp:
            break
    print(c)
