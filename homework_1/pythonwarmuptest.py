from pythonwarmup import (change, strip_vowels, scramble, powers_of_two, powers,
    interleave, stutter)
import unittest, math

def anagram(s, t):
    return sorted(s) == sorted(t)

class TestGettingStartedFunctions(unittest.TestCase):

    def test_change(self):
        self.assertEqual((3, 2, 0, 2), change(97))
        self.assertEqual((0, 0, 1, 3), change(8))
        self.assertEqual((56514938, 0, 1, 1), change(1412873456))
        self.assertEqual((1, 1, 1, 1), change(41))
        self.assertEqual((35, 1, 1, 0), change(890))
        self.assertEqual((0, 0, 0, 4), change(4))
        self.assertEqual((4938271, 1, 0, 4), change(123456789))
        self.assertEqual((10, 0, 0, 0), change(250))
        self.assertEqual((20, 1, 1, 4), change(519))
        self.assertEqual((44, 1, 0, 1), change(1111))
        with self.assertRaises(ValueError):
            change(-2)
            change(-30000)

    def test_strip_vowels(self):
        self.assertEqual('Hll, wrld', strip_vowels("Hello, world"))
        self.assertEqual(' lk trtls', strip_vowels("I like turtles"))
        self.assertEqual('', strip_vowels("aeiou"))
        self.assertEqual('bcdfghjklmnpqrstvwxyz', strip_vowels("bcdfghjklmnpqrstvwxyz"))        
        self.assertEqual('bcdf', strip_vowels("abecidoeuf"))
        self.assertEqual('Hppnss dpnds pn rslvs', strip_vowels("Happiness depends upon ourselves"))
        self.assertEqual('Chrnnhtnthlgs', strip_vowels("Chrononhotonthologos"))
        self.assertEqual('HNRFCBLTDNTTBS', strip_vowels("HONORIFICABILITUDINITATIBUS"))
        self.assertEqual('Phlp M. Drn, Jhn Dvd N. Dns, Ry Tl, Ms Ktmr', strip_vowels("Philip M. Dorin, John David N. Dionisio, Ray Toal, Masao Kitamura"))
        self.assertEqual('`~!@#$%^&*()_+-=[]{};./>?', strip_vowels("`~!@#$%^&*()_+-=[]{};./>?"))

    def test_scramble(self):
        data = ["a", "rat", "BSOD", "BDFL", "Python testing", "I like turtles", "Happiness depends upon ourselves."]
        for s in data:
            self.assertTrue(anagram(s, scramble(s)))


    def test_powers_of_two(self):
        self.assertEqual(tuple(powers_of_two(-1)), ())
        self.assertEqual(tuple(powers_of_two(0)), ())
        self.assertEqual(tuple(powers_of_two(2)), (1,))
        self.assertEqual(tuple(powers_of_two(3)), (1, 2))
        self.assertEqual(tuple(powers_of_two(5)), (1, 2, 4))
        self.assertEqual(tuple(powers_of_two(9)), (1, 2, 4, 8))
        self.assertEqual(tuple(powers_of_two(300)), (1, 2, 4, 8, 16, 32, 64, 128, 256))
        self.assertEqual(tuple(powers_of_two(2500)), (1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048))
        self.assertEqual(tuple(powers_of_two(4194305)), (1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304))
        self.assertEqual(tuple(powers_of_two(134217729)), (1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728))

    def test_powers(self):
        self.assertEqual(tuple(powers(0, 0)), ())
        self.assertEqual(tuple(powers(0, -1)), ())
        self.assertEqual(tuple(powers(-1, -4)), ())
        self.assertEqual(tuple(powers(-2, 17)), (1, -2, 4, -8, 16, -32))
        self.assertEqual(tuple(powers(3, 400)), (1, 3, 9, 27, 81, 243))
        self.assertEqual(tuple(powers(57, 900000000000000)), (1, 57, 3249, 185193, 10556001, 601692057, 34296447249, 1954897493193, 111429157112001))
        self.assertEqual(tuple(powers(6, 666)), (1, 6, 36, 216))
        self.assertEqual(tuple(powers(-300, 28000000)), (1, -300, 90000, -27000000))

    def test_interleave(self):
        self.assertEqual(["a", 1, "b", 2, True, None], interleave(["a", "b"], [1, 2, True, None]))
        self.assertEqual([1, "a", 2, "b", True, None], interleave([1, 2, True, None], ["a", "b"]))
        self.assertEqual([None], interleave([None], []))
        self.assertEqual(["H", 4, "E", 5, "L", 6, "L", 7, "O"], interleave(["H", "E", "L", "L", "O"], [4, 5, 6, 7]))
        self.assertEqual([1, 0, 1, 0, 1, 0], interleave([1, 1, 1], [0, 0, 0]))
        self.assertEqual([None, True, False, len, True, [], {}, ()], interleave([None, False, True, [], {}, ()], [True, len]))
        self.assertEqual([0, 4, 1, 5, 2, 6, 3, 7], interleave([0, 1, 2, 3], [4, 5, 6, 7]))
        self.assertEqual([], interleave([], []))

    def test_stutter(self):
        self.assertEqual(['a', 'a', True, True, 32, 32, {'name': 'spike'}, {'name': 'spike'}, (6+5j), (6+5j)], stutter(['a', True, 32, {"name" : "spike"}, 6 + 5j]))
        self.assertEqual([1, 1], stutter([1]))
        self.assertEqual(["Hello", "Hello", None, None, False, False], stutter(["Hello", None, False]))
        self.assertEqual([None, None, True, True, False, False, len, len, [], [], {}, {}, (), ()], stutter([None, True, False, len, [], {}, ()]))
        self.assertEqual([], stutter([]))
        self.assertEqual([math.trunc(9), math.trunc(9), 5 >> 6, 5 >> 6, 1.592834756, 1.592834756], stutter([math.trunc(9), 5 >> 6, 1.592834756]))

if __name__ == '__main__':
    unittest.main()