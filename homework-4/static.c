#include <stdio.h>

int f(int x) {
	int a = 0;
	if (x == 0) {
		a = 1;
	} else {
		return f(x - x) + a;
	}
	return 0;
}

int main() {
	printf("%d\n", f(5));
	return 0;
}