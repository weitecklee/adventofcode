file1 = open('input202202.txt','r')
lines = file1.readlines()
score = 0

for line in lines:
    plays = line.strip().split(' ')
    if plays[1] == 'X':
        score += 1
        if plays[0] == 'C':
            score += 6
        elif plays[0] == 'A':
            score += 3
    elif plays[1] == 'Y':
        score += 2
        if plays[0] == 'A':
            score += 6
        elif plays[0] == 'B':
            score += 3
    elif plays[1] == 'Z':
        score += 3
        if plays[0] == 'B':
            score += 6
        elif plays[0] == 'C':
            score += 3

print(score)

score2 = 0

for line in lines:
    plays = line.strip().split(' ')
    if plays[1] == 'X':
        if plays[0] == 'A':
            score2 += 3
        elif plays[0] == 'B':
            score2 += 1
        elif plays[0] == 'C':
            score2 += 2
    elif plays[1] == 'Y':
        if plays[0] == 'A':
            score2 += 4
        elif plays[0] == 'B':
            score2 += 5
        elif plays[0] == 'C':
            score2 += 6
    elif plays[1] == 'Z':
        if plays[0] == 'A':
            score2 += 8
        elif plays[0] == 'B':
            score2 += 9
        elif plays[0] == 'C':
            score2 += 7

print(score2)