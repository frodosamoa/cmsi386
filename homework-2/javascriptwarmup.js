 //Code available at http://jsfiddle.net/cXbxj/

$(function () {
 
    /**
     * Accepts a number of U.S. cents and returns an array containing the  smallest
     * number of U.S. quarters, dimes, nickels, and pennies that equal the given amount.
     */

    function change(cents) {
        var remaining = cents,
            denoms = [25, 10, 5];
        
        if (cents < 0) {
             throw new RangeError("The amount of cents must be 0 or greater.");   
        }

        var quarters = Math.floor(remaining % denoms[0]);
        remaining %= denoms[0];
        var dimes = Math.floor(remaining % denoms[0]);
        remaining %= denoms[1];
        var nickels = Math.floor(remaining % denoms[0]);
        var pennies = remaining % denoms[2];

        return [quarters, nickels, dimes, pennies];
    };

    /**
     * Returns the string which is equivalent to the passed string but with
     * all ASCII vowels removed.
     */

    function stripVowels(s) {
        return s.replace(/[aeiou]/ig, "");
    };

    /**
     * Randomly permutes a string.
     */

    function scramble(s) {
        var a = s.split("");
        for (var i = a.length; i > 0; i--) {
          var j = Math.floor(Math.random() * i);
          var c = a[i];
          a[i] = a[j];
          a[j] = c;
        }
        return a.join("");
    };

    /**
     * Yields successive powers of two starting at 1 and going up to some
     * limit. Consumes the values with a callback.
     */

    function powersOfTwo(limit, f) {
        for (var i = 1; i <= limit; i *= 2) {
            f(i);
        }
    };

    /**
     * Yields powers of an arbitrary base starting at exponent 0 and going up
     * to some limit. Consumes the values with a callback.
     */

    function powers(base, limit, f) {
        if (base === 1) {
            f(base);
        } else {
            for (var i = 1; i <= limit; i *= base) {             
                f(i);
            }
        }
    };

    /**
     * Returns an array with all of the values that either powersOfTwo() or
     * powers() would return. Base is optional (if present, use powers(), else
     * use powersOfTwo(). Helper function for testing.
     */

    function powerArray(limit, base) {
        var powerArray = [];

        if (base) {
            powers(base, limit, function (p) { powerArray.push(p) }); 
        } else {
            powersOfTwo(limit, function (p) { powerArray.push(p) });
        }
        
        return powerArray;
    };

    /**
     * Interleaves two arrays. If the arrays do not have the same length, the
     * elements of the longer array end up at the end of the result array.
     */

    function interleave(a, b) {
        var c = [];
        for (var i = 0, max = Math.max(a.length, b.length); i < max; i++) {
            if (i < a.length) c.push(a[i]);
            if (i < b.length) c.push(b[i]);
        }
        return c;
    };

    /**
     * Doubles up each item in an array.
     */

    function stutter(array) {
        return interleave(array, array);
    };

    /**
     * Produces a word count object from a string. Apostrophes are included
     * in word count.
     */

    function wordCount(string) {
        var wordCounts = {},
            words = [];

        words = string.toLowerCase().split(/[^a-zA-Z\']+/);

        for (var i = 0; i < words.length; i++) {
            if (words[i] !== "") {
                if (words[i] in wordCounts) {
                    wordCounts[words[i]] = wordCounts[words[i]] + 1;
                } else {
                    wordCounts[words[i]] = 1;
                }
            }
        }

        return wordCounts;
    };

    /**
     * Unit tests.
     */

    var anagram = function (s, t) {
        if (s.length !== t.length) {
            return false;
        }
        var a = s.split(""), b = t.split("");
        return $(a).not(b).length === 0 && $(b).not(a).length === 0;
    };

    test("Change Tests", function () {
        deepEqual(change(97), [3, 2, 0, 2]);
        deepEqual(change(8), [0, 0, 1, 3]);
        deepEqual(change(1412873456), [56514938, 0, 1, 1]);
        deepEqual(change(41), [1, 1, 1, 1]);
        deepEqual(change(890), [35, 1, 1, 0]);
        deepEqual(change(0), [0, 0, 0, 0]);
        deepEqual(change(123456789), [4938271, 1, 0, 4]);
        deepEqual(change(250), [10, 0, 0, 0]);
        raises(function () {change(-519)}, RangeError);
        raises(function () {change(-1)}, RangeError);
    });

    test("Strip Vowels Tests", function () {
        deepEqual(stripVowels("Hello, world"), "Hll, wrld");
        deepEqual(stripVowels("I like turtles"), " lk trtls");
        deepEqual(stripVowels("aeiou"), "");
        deepEqual(stripVowels("bcdfghjklmnpqrstvwxyz"), "bcdfghjklmnpqrstvwxyz");      
        deepEqual(stripVowels("abecidoeuf"), "bcdf");
        deepEqual(stripVowels("Happiness depends upon ourselves"), "Hppnss dpnds pn rslvs");
        deepEqual(stripVowels("Chrononhotonthologos"), "Chrnnhtnthlgs");
        deepEqual(stripVowels("HONORIFICABILITUDINITATIBUS"), "HNRFCBLTDNTTBS");
        deepEqual(stripVowels("Philip M. Dorin, John David N. Dionisio, Ray Toal, Masao Kitamura"), "Phlp M. Drn, Jhn Dvd N. Dns, Ry Tl, Ms Ktmr");
        deepEqual(stripVowels("`~!@#$%^&*()_+-=[]{};./>?"), "`~!@#$%^&*()_+-=[]{};./>?");
    });

    test("Scramble Tests", function () {
        var data = ["", "a", "rat", "zzz", "^*&^*&^▱ÄÈËɡɳɷ", "I like turtles", "Happiness depends upon ourselves."]
        data.forEach(function (s) {
            ok(anagram(s, scramble(s)));
        });
    });

    test("Powers of Two Tests", function () {
        deepEqual(powerArray(-1), []);
        deepEqual(powerArray(0), []);
        deepEqual(powerArray(2), [1, 2]);
        deepEqual(powerArray(3), [1, 2]);
        deepEqual(powerArray(5), [1, 2, 4]);
        deepEqual(powerArray(8), [1, 2, 4, 8]);
        deepEqual(powerArray(300), [1, 2, 4, 8, 16, 32, 64, 128, 256]);
        deepEqual(powerArray(2500), [1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048]);
        deepEqual(powerArray(4194305), [1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304]);
        deepEqual(powerArray(134217729), [1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728]);
    });

    test("Powers test", function () {
        deepEqual(powerArray(0, 0), []);
        deepEqual(powerArray(-1, 0), []);
        deepEqual(powerArray(-4, -1), []);
        deepEqual(powerArray(17, -2), [1, -2, 4, -8, 16, -32]);
        deepEqual(powerArray(400, 3), [1, 3, 9, 27, 81, 243]);
        deepEqual(powerArray(900000000000000, 57), [1, 57, 3249, 185193, 10556001, 601692057, 34296447249, 1954897493193, 111429157112001]);
        deepEqual(powerArray(666, 6), [1, 6, 36, 216]);
        deepEqual(powerArray(28000000, -300), [1, -300, 90000, -27000000]);
    });

    test("Interleave test", function () {
        deepEqual(interleave(["a", "b"], [1, 2, true, false]), ["a", 1, "b", 2, true, false]);
        deepEqual(interleave([1, 2, true, false], ["a", "b"]), [1, "a", 2, "b", true, false]);
        deepEqual(interleave([null], []), [null]);
        deepEqual(interleave(["H", "E", "L", "L", "O"], [4, 5, 6, 7]), ["H", 4, "E", 5, "L", 6, "L", 7, "O"]);
        deepEqual(interleave([1, 1, 1], [0, 0, 0]), [1, 0, 1, 0, 1, 0]);
        deepEqual(interleave([null, false, true, [1, 2, 3], { x : 2 }], [true]), [null, true, false, true, [1, 2, 3], { x : 2 }]);
        deepEqual(interleave([0, 1, 2, 3], [4, 5, 6, 7]), [0, 4, 1, 5, 2, 6, 3, 7]);
        deepEqual(interleave([], []), []);
    });

    test("Stutter test", function () {
        deepEqual(stutter(['a', true, 32, {"name" : "spike"}]), ['a', 'a', true, true, 32, 32, {'name': 'spike'}, {'name': 'spike'}]);
        deepEqual(stutter([1]), [1, 1]);
        deepEqual(stutter(["Hello", null, false]), ["Hello", "Hello", null, null, false, false]);
        deepEqual(stutter([null, true, false, [], {}]), [null, null, true, true, false, false, [], [], {}, {}]);
        deepEqual(stutter([]), []);
        deepEqual(stutter([Math.PI, 5 >> 6, 1.592834756]), [Math.PI, Math.PI, 5 >> 6, 5 >> 6, 1.592834756, 1.592834756]); 
    });

    test("Word count test", function () {
        deepEqual(wordCount(""), {});
        deepEqual(wordCount("Hello, world!"), {"hello" : 1, "world": 1});
        deepEqual(wordCount("Hello, HELLO><.,><>< heLlo!"), {"hello" : 3});
        deepEqual(wordCount("I_i-found,a, pink LION"), {"i" : 2, "found" : 1, "a" : 1, "pink" : 1, "lion" : 1});
        deepEqual(wordCount("hello dog-tired"), {"hello" : 1, "dog": 1, "tired" : 1});
        deepEqual(wordCount("a.b,c<d>e?f/g;h:i!j@k\[l\]m{n}o=p+q_r-s)t(u*v&w^x%y$z "), {"a" : 1, "b" : 1, "c" : 1, "d" : 1, "e" : 1, "f" : 1, "g" : 1, "h" : 1, "i" : 1, "j" : 1, "k" : 1, "l" : 1, "m" : 1, "n" : 1, "o" : 1, "p" : 1, "q" : 1, "r" : 1, "s" : 1, "t" : 1, "u" : 1, "v" : 1, "w" : 1, "x" : 1, "y" : 1, "z" : 1,});
    });
    
});