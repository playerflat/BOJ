while 1:
    a = input().lower()
    if a == '#':
        break
    print(sum(1 for i in a if i in ['a', 'e', 'i', 'o', 'u']))