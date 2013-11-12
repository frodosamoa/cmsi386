#include <iostream>
using namespace std;

void interleave(int a[], int lena, int b[], int lenb, int c[]) {
	int aindex = 0;
	int bindex = 0;

	for (int i = 0; i < (lena + lenb); ) {
		if (aindex < lena) {
			c[i] = a[aindex];
			aindex++;
			i++;
		}

		if (bindex < lenb) {
			c[i] = b[bindex];
			bindex++;
			i++;
		}
    }
}

int main () {

	int a[50], b[50], c[100], lena, lenb, lenc, i;

	cout << "Input number of elements in first array\n";
	scanf("%d", &lena);

	cout << "Input " << lena << " integers\n";
	for (i = 0; i < lena; i++) {
		scanf("%d", &a[i]);
	}

	cout << "Input number of elements in second array\n";
	scanf("%d", &lenb);

	cout << "Input " << lenb << " integers\n";
	for (i = 0; i < lenb; i++) {
		scanf("%d", &b[i]);
	}

	interleave(a, lena, b, lenb, c);
	lenc = lena + lenb;

	for (int i = 0; i < lenc; i++) {
		cout << c[i] << " ";
	}
	cout << endl;

  	return 0;
}