a = input()

for i in a:
    if i.islower():
        b = ord(i) + 13
        if b >= 123:
            b = b % 123 + 97
        print(chr(b), end='')
    elif i.isupper():
        b = ord(i) + 13
        if b >= 91:
            b = b % 91 + 65
        print(chr(b), end='')
    else:
        print(i, end='')
