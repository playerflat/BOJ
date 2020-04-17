a = list(map(int, input().split()))
if a[0] == a[1] == a[2]:
    print(2)
elif min(a) ** 2 + (sum(a) - min(a) - max(a)) ** 2 == max(a) ** 2:
    print(1)
else:
    print(0)
