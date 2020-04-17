a = input()
if len(a) == 4:
    print(20)
elif len(a) == 2:
    print(int(a[0]) + int(a[1]))
elif a[1] == "0":
    print(10 + int(a[2]))
elif a[2] == "0":
    print(10 + int(a[0]))
