def s(n):
    if n - int(n) != 0:
        return int(n) + 1
    else:
        return int(n)


l = int(input())
a = int(input())
b = int(input())
c = int(input())
d = int(input())

a = s(a / c)
b = s(b / d)
print(l - a if a > b else l - b)