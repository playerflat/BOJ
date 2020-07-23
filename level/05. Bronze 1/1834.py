r = 0
N = int(input())
for i in range(N + 1, N ** 2, N + 1):
    if i % N == 0:
        break
    r += i
print(r)
