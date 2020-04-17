r = 0
for i in range(int(input())):
    a = int(input())
    if a == 1:
        r += 1
    else:
        r -= 1

if r < 0:
    print("Junhee is not cute!")
else:
    print("Junhee is cute!")
