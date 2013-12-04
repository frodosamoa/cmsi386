#include <stdio.h>

#define MAX_INT (1.0 / 0)

int minFinder(int *a, int size, int min, int index) {
	if (index == size - 1) {
		return min;
	} else {
		min = min < a[index] ? min : a[index];
		index++;
	}
	return minFinder(a, size, min, index);
}

int arrayMin(int *a, int size) {
	return minFinder(a, size, MAX_INT, 0);
}

int main () {
	int a[10] = {5, 6, 7, 12, 3, 1, 234, 123, 67, 9000};
	int b[4] = {-432, 3, 1, 234};
	int c[7] = {-12, -11, -4, 1344234, -87653, 0, 75};
	printf("%d, %d, %d\n", arrayMin(a, 10), arrayMin(b, 4), arrayMin(c, 7));
}