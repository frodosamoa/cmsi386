import random

def change(cents):
	"""Returns a tuple containing the smallest number of U.S. quarters,
	dimes, nickels, and pennies that equal the passed amount of cents."""
	if cents < 0:
		raise ValueError('Amount of cents must be 0 or greater.')
	quarters, remaining = divmod(cents, 25)
	dimes, remaining = divmod(remaining, 10)
	nickels, pennies = divmod(remaining, 5)
	return (quarters, dimes, nickels, pennies)

def strip_vowels(string):
	"""Returns the passed string with all ASCII vowels removed."""
	return ''.join([char for char in string if char.lower() not in 'aeiou'])

def scramble (string):
	"""Returns the passed string randomly permuted."""
	return ''.join(random.sample(string, len(string)))

def powers_of_two(n):
	"""Generator of powers of two up to and including the passed limit."""
	num = 1
	while num <= n:
		yield num
		num *= 2

def powers(base, n):
	"Generator of powers of an arbitrary base up to and including the passed limit."
	num = 1
	while num <= n:
		yield num
		num *= base

def interleave(l1, l2):
	"""Returns the two passed lists interleaved."""
	l3 = [value for pair in zip(l1, l2) for value in pair]
	if len(l1) != len(l2):
		l3.extend(l1[len(l2):] if len(l1) > len(l2) else l2[len(l1):])
	return l3

def stutter (list):
	"""Returns the list with each item doubled up."""
	return interleave(list, list)