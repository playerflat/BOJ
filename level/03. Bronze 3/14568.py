c = int(input())
s = 0
for y in range(1, c):
    for n in range(y + 2, c):
        t = c - y - n
        if t % 2 == 0 and t != 0 and t > 0 and c - t - y - n == 0:
            s += 1
print(s)
