a = map(int, input().split())
a = list(a)
prev = -1000001
flag = True
for i in range(len(a)):
    if prev > a[i]:
        flag = False
        break
    prev = a[i]

if flag:
    print("Good")
else:
    print("Bad")