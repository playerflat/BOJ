stick = [64]

X = int(input())

while True:
    if X != sum(stick):
        a = stick.pop()
        a /= 2
        stick.append(a)
        stick.append(a)
        if X <= sum(stick[:-1]):
            stick.pop()
    else:
        print(len(stick))
        break
        