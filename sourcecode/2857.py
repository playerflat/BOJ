t = False
for i in range(1, 6):
    if 'FBI' in input():
        print(i, end=' ')
        t = True
if not t:
    print("HE GOT AWAY!")
