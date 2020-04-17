a, b = map(int, input().split())
if a >= b * 7:
    print(b * 7000)
elif a * 2 >= b * 7:
    print(b * 3500)
elif a * 4 >= b * 7:
    print(b * 1750)
else:
    print(0)
