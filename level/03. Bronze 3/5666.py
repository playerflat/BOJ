while 1:
    try:
        h, p = map(int, input().split())
        print(f'{h / p:.2f}')
    except EOFError:
        break