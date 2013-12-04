def list_min(list):
	def min_finder(l, m):
		if not l:
			print m
		else:
			min_finder(l[1:], m if m < l[0] else l[0])
	min_finder(list, float("inf"))