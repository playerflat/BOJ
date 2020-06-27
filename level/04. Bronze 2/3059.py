for _ in range(int(input())):
    a = {chr(i): 1 for i in range(65, 91)}
    b = 0
    for i in input():
        a[i] -= 1
    for i, v in a.items():
        if v == 1:
            b += ord(i)
    print(b)
