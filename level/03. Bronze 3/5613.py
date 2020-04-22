r = int(input())
a = 0
o = "+"
while 1:
    if o == "+":
        r += a
    elif o == "-":
        r -= a
    elif o == "*":
        r *= a
    elif o == "/":
        r //= a
    o = input()
    if o == "=":
        break
    a = int(input())
print(r)
