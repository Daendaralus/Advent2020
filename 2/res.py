import re


regex = "(\d+)-(\d+) (\w): (\w+)"

f = open("data", "r")
lines = f.readlines()
count = 0
for x in lines:
    groups = re.match(regex, x)
    min, max, tar, stuff = groups.groups()[0:4]
    if (stuff[int(min)-1]==tar) != (stuff[int(max)-1]==tar):
        count+=1
print(count)
