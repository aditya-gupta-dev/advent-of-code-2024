lines, left, right = [], [], [] 

with open('input.txt', 'r') as file: 
    lines = file.read().splitlines()
    
for line in lines: 
    l, r = list(map(int, line.split()))
    left.append(l)
    right.append(r)

left.sort()
right.sort()

result = [] 

for i in range(len(left)):
    result.append(abs(left[i] - right[i]))

print(sum(result))

## part 2 

lines, left, right = [], [], [] 

def similarity(num: int, right_list: list[int]) -> int:
    counter = 0 
    for right_item in right_list:
        if right_item == num: 
            counter = counter + 1
    return counter 

with open('input.txt', 'r') as file: 
    lines = file.read().splitlines()
    
for line in lines: 
    l, r = list(map(int, line.split()))
    left.append(l)
    right.append(r)

total_sum = 0 

for i in range(len(left)):
    total_sum = total_sum + left[i] * similarity(num=left[i], right_list=right)

print(total_sum)

