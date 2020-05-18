input()
a = list(map(int, input().split()))
r = []
f = False
for i in range(1, len(a)):
    if f:
        if a[i - 1] >= a[i]:
            r.append(a[i - 1])
            f = False
    else:
        if a[i - 1] < a[i]:
            r.append(a[i - 1])
            f = True
        else:
            f = False
    if i == len(a)-1:
        if a[i] > a[i - 1]:
            r.append(a[i])

ri = 0
for i in range(0, len(r), 2):
    ri = max(ri, r[i + 1] - r[i])
print(ri)
