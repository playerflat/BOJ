a = sorted(input(), reverse=True)

if a[-1] != '0':
    print(-1)
else:
    s = 0
    for i in a:
        s += int(i)

    if s % 3 != 0:
        print(-1)
    else:
        print(''.join(a))
