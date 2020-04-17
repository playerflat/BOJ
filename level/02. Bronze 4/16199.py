a = list(map(int, input().split()))
b = list(map(int, input().split()))

if b[0] == a[0]:
    print(0)
elif b[1] > a[1] or (b[1] == a[1] and b[2] >= a[2]):
    print(b[0] - a[0])
else:
    print(b[0] - a[0] - 1)
print(b[0] - a[0] + 1)
print(b[0] - a[0])
