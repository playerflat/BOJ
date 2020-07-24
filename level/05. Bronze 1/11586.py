N = int(input())
a = []
for i in range(N):
    a.append(input())

K = int(input())
if K == 1:
    for i in a:
        print(i)
elif K == 2:
    for i in a:
        print(i[::-1])
else:
    for i in range(N-1, -1, -1):
        print(a[i])
