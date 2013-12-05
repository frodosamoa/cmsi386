#include <iostream>
using namespace std;

class OddGenerator {
	private:
		int x = -1;
	public:
		int nextOdd();
};

int OddGenerator::nextOdd() {
	return x += 2;
}

int main () {
	OddGenerator odds;
	cout << odds.nextOdd() << endl;
	cout << odds.nextOdd() << endl;
	cout << odds.nextOdd() << endl;
	cout << odds.nextOdd() << endl;
}