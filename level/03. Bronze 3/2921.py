a = int(input())
r = 0
for i in range(1, a + 1):
    r += i * (i + 1) + (i * (i + 1) / 2)
print(int(r))