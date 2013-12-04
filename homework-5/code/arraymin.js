var arrayMin = function (array) {
	var minFinder = function (a, m) {
		return a.length === 0 ? m : minFinder(a.slice(1), m < a[0] ? m : a[0]);
	}
	return minFinder(array, Infinity);
}