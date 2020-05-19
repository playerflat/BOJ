r = 0
input()
for i in input():
    if i == 'A':
        r += 1
    else:
        r -= 1

if r > 0:
    print('A')
elif r < 0:
    print('B')
else:
    print('Tie')
