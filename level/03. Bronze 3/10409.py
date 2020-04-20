a = list(map(int, input().split()))
b = list(map(int, input().split()))
for i in range(1, len(b)+1):
    if sum(b[:i]) > a[1]:
        print(i-1)
        break
if sum(b) <= a[1]:
    print(len(b))

