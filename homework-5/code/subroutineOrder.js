var a = function () {alert("first");};
var b = function () {alert("second");};
var both = function (a, b) {};

both(a(), b());