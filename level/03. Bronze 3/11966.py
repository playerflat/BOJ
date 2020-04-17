a = []
for i in range(31):
    a.append(2 ** i)

if int(input()) in a:
    print(1)
else:
    print(0)
