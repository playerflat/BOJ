n = int(input())
a = list(map(int, input().split()))
a.sort()

r = 0
for i in range(len(a)):
    r += a[i] * ((2 * i) - (2 * (n - i - 1)))
print(r)