#include <iostream>
#include <vector>
using namespace std;

template <typename T>
vector <T> zip(const vector <T> & a, const vector <T> & b, size_t len) {
    vector <T> result;
    int aindex = 0;
    int bindex = 0;

    for (size_t i = 0; i < len; ) {
    	if (aindex < a.size()) {
        	result.push_back(a[aindex]);
        	aindex++;
        	i++;
    	}

    	if (bindex < b.size()) {
	        result.push_back(b[bindex]);
	        bindex++;
	        i++;
	    }
    }
    return result;
}

int main() {
    vector <int> a = {};
    vector <int> b = {};
    size_t lenc;
    int i, n, m, lena, lenb;

	cout << "Input number of elements in first array" << endl;
	scanf("%d", &lena);

	cout << "Input " << lena << " integers" << endl;
	for (i = 0; i < lena; i++) {
		scanf("%d", &n);
		a.push_back(n);
	}

	cout << "Input number of elements in second array" << endl;
	scanf("%d", &lenb);

	cout << "Input " << lenb << " integers" << endl;
	for (i = 0; i < lenb; i++) {
		scanf("%d", &m);
		b.push_back(m);
	}

	len = (a.size() + b.size());
    vector <int> c = zip(a, b, lenc);
    for(size_t i = 0; i < c.size(); i++) {
        cout << c[i] << " ";
    }
    cout << endl;
}