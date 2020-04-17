from datetime import datetime

a, b, c = map(int, input().split())
start = datetime(2011, 11, 11, 11, 11, 11)
end = datetime(2011, 11, a, b, c, 0)

if round((end-start).days*1440+(end-start).seconds/60) < 0:
    print(-1)
else:
    print(round((end-start).days*1440+(end-start).seconds/60))