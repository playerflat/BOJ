r = 0
for i in range(int(input())):
    a = input()
    r += int(a[:-1]) ** int(a[-1])
print(r)
