a = list(map(int, input().split()))
a.sort()
b = input()
for i in b:
    if i == 'A':
        print(a[0], end=' ')
    elif i == 'B':
        print(a[1], end=' ')
    else:
        print(a[2], end=' ')
