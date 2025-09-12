reports: list[list[int]] = [] 

def is_safe(report: list[int]) -> bool: 
    sorted_report = []
    cloned_report = report 
    if report[0] < report[len(report) - 1]: 
        report.sort()
        sorted_report = report 
    else: 
        report.sort(reverse=True)
        sorted_report = report 

    if cloned_report != sorted_report: 
        return False 
    
    for i in range(len(sorted_report) - 2): 
        diff = abs(sorted_report[i] - sorted_report[i+1])
        if diff > 3:
            return False
    
    return True 

with open('input-two.txt', 'r') as file: 
    for line in file.read().splitlines(): 
        reports.append(list(map(int, line.split())))

total_safe_reports = 0 

for report in reports: 
    if is_safe(report): 
        total_safe_reports = total_safe_reports + 1
    else: 
        print(report)
        break 

print(total_safe_reports)