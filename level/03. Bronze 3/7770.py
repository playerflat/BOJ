a = int(input())
b = 0
for i in range(1, 999999):
    b += i ** 2 + (i - 1) ** 2
    if a / b < 1:
        print(i-1)
        break
