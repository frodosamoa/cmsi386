"""This script reports the number of non-blank, non-commented
lines in the file named by the first argument.
"""

import sys

with (open(sys.argv[1], 'r')) as file:
	lines = 0
	for line in file:
		if line.strip() and not line.strip().startswith('#'):
			lines += 1

print lines