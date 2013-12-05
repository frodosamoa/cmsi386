class odd_generator:
	def __init__(self):
		self.__x = -1

	def next_odd(self):
		self.__x = self.__x + 2
		return self.__x

z = odd_generator()
print z.next_odd()
print z.next_odd()
print z.next_odd()
print z.next_odd()