"""This script writes successive prefixes of its first input argument, 
one per line, starting with the first prefix, which is zero characters long.
"""

import sys

arg = sys.argv[1]

def prefixes (string):
	for x in range(0, len(string) + 1):
		print ''.join(string[:x])

prefixes(arg)