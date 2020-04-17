a = int(input())
if a % 10 != 0:
    print(-1)
else:
    print(int(a / 300), int(a % 300 / 60), int(a % 300 % 60 / 10))
