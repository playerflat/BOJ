a = [1]

for i in range(2, 1001):
    a.append(i + a[i - 2])

for _ in range(int(input())):
    f = False
    n = int(input())

    for i in a:
        if i * 3 < n:
            continue
        if i > n or f:
            break
        for j in a:
            if j > n or f:
                break
            for k in a:
                if k > n:
                    break
                if i + j + k == n:
                    print(1)
                    f = True
                    break
    if not f:
        print(0)
