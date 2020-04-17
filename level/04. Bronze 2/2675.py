a = int(input())
for i in range(a):
    b, c = input().split()
    b = int(b)
    for j in range(len(c)):
        print(c[j]*b, end='')
    print()