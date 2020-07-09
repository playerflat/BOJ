for _ in range(int(input())):
    input()
    a = list(map(int, input().split()))
    r = 0
    a.sort()
    for i in range(1, len(a)):
        r += abs(a[i - 1] - a[i])
    r += abs(a[-1] - a[0])
    print(r)
