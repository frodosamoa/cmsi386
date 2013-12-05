def odd_generator():
	odd = {'current' : -1}
	def next_odd(): odd['current'] += 2; print odd['current']
	return next_odd

a = odd_generator()

a()
a()
a()
a()