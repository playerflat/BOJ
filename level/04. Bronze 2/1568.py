a = int(input())
c = 0
i = 1
while 1:
    a -= i
    c += 1
    if a == 0:
        print(c)
        break
    i += 1
    if i > a:
        i = 1
