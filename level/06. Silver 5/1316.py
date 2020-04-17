a = int(input())
count = 0

for i in range(a):
    count += 1
    z = {}
    b = input()
    for j in range(len(b)):
        if b[j] not in z:
            z[b[j]] = j
        else:
            if z[b[j]] != j - 1:
                count -= 1
                break
            else:
                z[b[j]] = j

print(count)
