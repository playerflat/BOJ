a = int(input())
if a % 2 != 0:
    a /= 2
    a = int(a)
    print((a + 1) * (a + 2))
else:
    a /= 2
    a = int(a)
    print((a + 1) * (a + 1))
