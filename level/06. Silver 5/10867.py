a = input()
b = map(int, input().split())
b = set(b)
b = list(b)
b.sort()
for i in b:
    print(i, end=" ")
