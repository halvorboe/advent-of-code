
with open("input/1/a.txt") as f:
    data = [[int(n) for n in l.splitlines()] for l in f.read().split("\n\n")]

print("part 1", (max(sum(d) for d in data)))

l = [sum(d) for d in data]
l.sort()
print("part 2", sum(l[-3:]))
