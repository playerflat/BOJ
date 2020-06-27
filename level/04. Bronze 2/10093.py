a = list(map(int, input().split()))
if a[0] == a[1]:
    print(0)
else:
    b, c = max(a), min(a)
    print(b - c - 1)
    for i in range(c + 1, b):
        print(i, end=' ')
