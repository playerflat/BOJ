y, m = 0, 0
int(input())
c = input().split()
for i in range(len(c)):
    a = int(c[i])
    y += (a // 30 + 1) * 10
    m += (a // 60 + 1) * 15

if y > m:
    print(f"M {m}")
elif m > y:
    print(f"Y {y}")
else:
    print(f"Y M {y}")
