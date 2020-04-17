a = int(input())
count = 1
min = 1
max = 1
m = []
while True:
    if a == 1:
        print("1/1")
        break
    if min <= a <= max:
        break
    min = max + 1
    max = min + count
    count += 1

for i in range(max-min+1):
    m.append(min+i)

if (max - min) % 2 == 0:
    count += 1
    print(count-m.index(a)-1, end='')
    print("/", end='')
    print(count-list(reversed(m)).index(a)-1)
else:
    count +=1
    print(count-list(reversed(m)).index(a)-1, end='')
    print("/", end='')
    print(count-m.index(a)-1)