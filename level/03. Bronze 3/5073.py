while 1:
    a = list(map(int, input().split()))
    if sum(a) == 0: break
    if sum(a) - max(a) <= max(a):
        print("Invalid")
    elif a[0] == a[1] == a[2]:
        print("Equilateral")
    elif len(set(a)) == 2:
        print("Isosceles")
    else:
        print("Scalene")
