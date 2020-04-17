a = int(input())
count = 1
m = 0
min = 2
max = 7

while True:
    if a == 1:
        break
    else:
        count += 1
        if min+m*6 <= a <= max+m*6:
            break
        min += m*6
        m += 1
        max += m*6
print(count)
