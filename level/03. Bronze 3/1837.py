a, b = map(int, input().split())
p = [1] * (b + 1)
for i in range(2, int(b ** 0.5) + 1):
    if p[i] == 1:
        for j in range(i + i, b + 1, i):
            p[j] = 0
p = [i for i in range(2, b + 1) if p[i] == 1]

for i in p:
    if a % i == 0:
        if i < b:
            print(f"BAD {i}")
            break
        else:
            print("GOOD")
            break
    if i == p[-1]:
        print("GOOD")
