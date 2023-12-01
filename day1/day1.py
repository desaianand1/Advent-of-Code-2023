import re

def extract_digits_p1(line:str):
    first = None
    lastIdx = -1
    for i,el in enumerate(line):
        if el.isnumeric():
            if first == None:
                first = el
            else:
                lastIdx = i
    if first is None:
        return ""
    if lastIdx == -1:
        return first+first
    return first + line[lastIdx]


def extract_digits_p2(line:str):
    numMap = {"one" : 1,"two":2,"three":3,"four":4,"five":5,"six":6,"seven":7,"eight":8,"nine":9}
    pattern = r'(?=(one|two|three|four|five|six|seven|eight|nine|[1-9]))'
    matches = [match.group(1) for match in re.finditer(pattern, line)]
    first, second = matches[0], matches[-1]
    if first in numMap:
        first = numMap[first]
    if second in numMap:
        second = numMap[second]
    return f"{first}{second}"
        
    
def main():
    # f_name = "input/test.txt"
    f_name = "input/input.txt"
    with open(f_name) as f:
        lines = f.readlines()
        sum = 0
        for line in lines:
            digits = extract_digits_p2(line)
            # print("digit: "+digits)
            if digits !="":
                sum = sum + int(digits)
        print(f"sum: {sum}")


if __name__ == "__main__":
    main()
