a = int(input())

for i in range(a):
    b = float(input())
    print("$", "%0.2f" % round(b / 100 * 80, 2), sep='')