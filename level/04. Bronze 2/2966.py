A = ['A', 'B', 'C']
B = ['B', 'A', 'B', 'C']
C = ['C', 'C', 'A', 'A', 'B', 'B']
d = {'Adrian': 0, 'Bruno': 0, 'Goran': 0}
r = []
a = int(input())
w = input()
for i in range(a):
    if w[i] == A[i % 3]:
        d['Adrian'] += 1
    if w[i] == B[i % 4]:
        d['Bruno'] += 1
    if w[i] == C[i % 6]:
        d['Goran'] += 1

print(max(d.values()))
for i, v in d.items():
    if v == max(d.values()):
        print(i)
