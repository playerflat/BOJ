a = list(map(int, input().split()))
b = list(map(int, input().split()))
sa = 0
sb = 0
f = 0

for i in range(len(a)):
    if a[i] > b[i]:
        sa += 3
        f = -1
    elif a[i] < b[i]:
        sb += 3
        f = 1
    else:
        sa += 1
        sb += 1

print(sa, sb)
if sa > sb or (sa == sb and f == -1):
    print("A")
elif sa < sb or (sa == sb and f == 1):
    print("B")
else:
    print("D")