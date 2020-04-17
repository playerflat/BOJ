a = list()
for i in range(3):
    b = int(input())
    a.append(b)

if sum(a) != 180:
    print("Error")
elif len(set(a)) == 1:
    print("Equilateral")
elif sum(a) == 180 and len(set(a)) == 2:
    print("Isosceles")
else:
    print("Scalene")