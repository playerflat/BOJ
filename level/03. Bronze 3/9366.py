for i in range(1, int(input()) + 1):
    a = list(map(int, input().split()))
    print(f'Case #{i}:', end=' ')
    if max(a) >= sum(a) - max(a):
        print('invalid!')
    elif len(set(a)) == 1:
        print('equilateral')
    elif len(set(a)) == 2:
        print('isosceles')
    else:
        print('scalene')
