file1 = open('input.txt','r')
text = file1.read()

letters = {}

for i in range(0, len(text)):
    if i >= 4:
        letters[text[i - 4]] -= 1
        if letters[text[i - 4]] == 0:
            letters.pop(text[i - 4])
    if text[i] in letters:
        letters[text[i]] += 1
    else:
        letters[text[i]] = 1
    if len(letters) == 4:
        print(i + 1)
        break

letters2 = {}
for i in range(0, len(text)):
    if i >= 14:
        letters2[text[i - 14]] -= 1
        if letters2[text[i - 14]] == 0:
            letters2.pop(text[i - 14])
    if text[i] in letters2:
        letters2[text[i]] += 1
    else:
        letters2[text[i]] = 1
    if len(letters2) == 14:
        print(i + 1)
        break
