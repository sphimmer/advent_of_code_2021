def solve2a(input):
    x = 0
    y = 0
    for line in input:
        splitLine = line.split(" ")
        direction = splitLine[0]
        amount = int(splitLine[1])

        if direction == "forward":
            x = x + amount
        elif direction == "down":
            y = y + amount
        elif direction == "up":
            y = y - amount
    return x * y

def solve2b(input):
    aim = 0
    x = 0
    y = 0
    for line in input:
        splitLine = line.split(" ")
        direction = splitLine[0]
        amount = int(splitLine[1])

        if direction == "forward":
            x = x + amount
            y = y + (aim * amount)
        elif direction == "down":
            aim = aim + amount
        elif direction == "up":
            aim = aim - amount
    return x * y

with open('input') as f:
    lines = f.readlines()

print(solve2a(lines))
print(solve2b(lines))