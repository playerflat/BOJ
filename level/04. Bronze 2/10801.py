a = list(map(int, input().split()))
b = list(map(int, input().split()))
s = 0
for i in range(10):
    if a[i] > b[i]:
        s += 1
    elif a[i] < b[i]:
        s -= 1
if s > 0:
    print("A")
elif s < 0:
    print("B")
else:
    print("D")
