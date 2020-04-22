import math

a = list(map(int, input().split()))
b = list(map(int, input().split()))
r = 0
for i in b:
    r += math.ceil(i / 2)
if a[0] <= r:
    print("YES")
else:
    print("NO")