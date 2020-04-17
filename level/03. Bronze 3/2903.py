a = int(input())
r = 3
for i in range(a-1):
    r += r-1
print(r**2)