for _ in range(3):
    a = input()
    b = 1
    c = ''
    d = []
    for i in a:
        if i == c:
            b += 1
        else:
            d.append(b)
            c = i
            b = 1
    d.append(b)
    print(max(d))
