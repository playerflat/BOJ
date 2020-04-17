n = int(input())

d = {}

a = input().split()
for i in range(len(a)):
    d[a[i]] = 1

n = int(input())

a = input().split()
for i in range(len(a)):
    if a[i] in d:
        print(1)
    else:
        print(0)