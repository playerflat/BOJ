def base(num, b):
    s = 0
    while num > 0:
        s += num % b
        num //= b
    return s


for i in range(2992, 10000):
    if base(i, 10) == base(i, 12) == base(i, 16): print(i)
