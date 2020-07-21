a = ''
while 1:
    try:
        a += input()
    except EOFError:
        print(sum(map(int, a.replace(',', ' ').split())))
        break
