public class OddGenerator {
	
	private int x = -1;
	public int nextOdd() {
		return x += 2;
	}

	public static void main (String[] args) {
		OddGenerator odds = new OddGenerator();
		System.out.println(odds.nextOdd());
		System.out.println(odds.nextOdd());
		System.out.println(odds.nextOdd());
		System.out.println(odds.nextOdd());
	}
}

