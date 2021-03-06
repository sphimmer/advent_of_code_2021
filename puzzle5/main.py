import pprint

def drawStraightLine(start: list, end: list, map: list) -> list:
 
    y = start[1] if start[1] == end[1] else False

    if y or type(y) == type(0):
        lowestX = start[0] if start[0] < end[0] else end[0]
        highestX = start[0] if start[0] > end[0] else end[0]
        if len(map) < y:
            map = map + [[0] * (highestX+1) for i in range(y+1)]
        if len(map[y]) < highestX+1:
            map[y] = map[y] + [0] * (highestX - len(map[y]) + 1)
        for x in range(lowestX, highestX+1):
            map[y][x] = map[y][x] + 1

    x = start[0] if start[0] == end[0] else False
    
    if x or type(x) == type(0):
        lowestY = start[1] if start[1] < end[1] else end[1]
        highestY = start[1] if start[1] > end[1] else end[1]
        for y in range(lowestY, highestY + 1):
            if len(map) < y+1:
                map = map + [[0] * (x+1) for i in range(y+1 - len(map))]
            if len(map[y]) < x+1:    
                map[y] = map[y] + [0] * (x - len(map[y]) + 1)
                
            map[y][x] = map[y][x] + 1

    return map

def drawDiagnolLine(start: list, end: list, map: list) -> list:
    minY = start[1] if start[1] < end[1] else end[1]
    maxX = start[0] if start[0] > end[0] else end[0]
    maxY = start[1] if start[1] > end[1] else end[1]
    diffY = maxY - minY
    if len(map) < maxY+1:
        map = map + [[0] * (maxX+1) for i in range(maxY)]

    x = start[0]
    y = start[1]
    for _ in range(diffY+1):
        if len(map[y]) < maxX+1:    
            map[y] = map[y] + [0] * (maxX - len(map[y]) + 1)
        map[y][x] = map[y][x] + 1
        if start[0] < end[0]:
            x = x+1
        else: 
            x = x-1
        if start[1] < end[1]:
            y = y+1
        else:
            y = y-1
       
    return map
            

def part1(map: list[list[int]], thermalLines: list[str]) -> int:
    for tl in thermalLines:
        coords = tl.split(' -> ')
        start = coords[0].split(',')
        end = coords[1].split(',')
        for i in range(len(start)):
            start[i] = int(start[i])
        for i in range(len(end)):
            end[i] = int(end[i])
        if start[0] == end[0] or start[1] == end[1]:
            map = drawStraightLine(start, end, map)
    count = 0
    for y in map:
        for x in y:
            if x >= 2:
                count = count+1
    return count

def part2(map: list[list[int]], thermalLines: list[str]) -> int:
    for tl in thermalLines:
        # print(tl)
        coords = tl.split(' -> ')
        start = coords[0].split(',')
        end = coords[1].split(',')
        for i in range(len(start)):
            start[i] = int(start[i])
        for i in range(len(end)):
            end[i] = int(end[i])
        if start[0] == end[0] or start[1] == end[1]:
            map = drawStraightLine(start, end, map)
        else:
            map = drawDiagnolLine(start, end, map)
    count = 0
    for y in map:
        for x in y:
            if x >= 2:
                count = count+1
    return count

thermalLines = ""
with open('input') as f:
    thermalLines = f.read().splitlines()
map = [[]]

print("Part1", part1(map, thermalLines))
print("Part2", part2(map, thermalLines))
