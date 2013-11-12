a = 4
b = 0
c = 7
d = 7

def f(n):
	global a
	a = n
	return a

print (a - f(b) - c * d)