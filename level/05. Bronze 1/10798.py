a = input()
b = input()
c = input()
d = input()
e = input()

for i in range(15):
    for j in [a, b, c, d, e]:
        try:
            print(j[i], end='')
        except IndexError:
            pass
