for _ in range(int(input())):
    a = input()
    b = str(int(a) + int(a[::-1]))
    if b == b[::-1]:
        print("YES")
    else:
        print("NO")
