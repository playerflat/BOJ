n = int(input())
for i in range(n):
    a = input().lower()
    if a == a[::-1]:
        print("Yes")
    else:
        print("No")
