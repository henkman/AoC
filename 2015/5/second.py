import re
import sys

reFirst = re.compile(r"([a-z]{2}).*?\1")
reSecond = re.compile(r"([a-z])[a-z]\1")
sum = 0
for line in sys.stdin.readlines():
	line = line.rstrip()
	if reFirst.search(line) and reSecond.search(line):
		sum += 1
print(sum)
