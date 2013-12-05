// Function in action at http://codepad.org/D4reNoZ5

int arrayMin(int *a, int size, int min, int index) {
	if (index == size - 1) {
		return min;
	} else {
		min = min < a[index] ? min : a[index];
		index++;
	}
	return arrayMin(a, size, min, index);
}