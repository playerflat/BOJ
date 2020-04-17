n = int(input())

for i in range(n):
    x = input()
    count = 0
    for j in range(len(x)):
        if x[j] == "(":
            count += 1
        elif x[j] == ")":
            count -= 1

        if count<0:
            break

    if count==0:
        print("YES")
    else:
        print("NO")