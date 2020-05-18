# PyPy3

def gcd(x, y):
    if x == 0:
        return y
    return gcd(y % x, x)


a = int(input())
b = list(map(int, input().split()))
c = gcd(b[0], gcd(b[1], b[-1]))

for i in range(1, (c // 2) + 1):
    if c % i == 0:
        print(i)
print(c)
