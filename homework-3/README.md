##Homework 3

For this homework assignment, you'll be warming up to Go. 

Package everything up nicely in a github repo.

1. Write a Go package containing the following:
  1. A function that accepts a number of U.S. cents and returns a tuple containing, respectively, the smallest number of U.S. quarters, dimes, nickels, and pennies that equal the given amount.
  2. A function that takes in a string s and returns the string which is equivalent to s but with all ASCII vowels removed.
  3. A function that randomly permutes a string. By random we mean that each time you call the function for a given argument all possible permutations are equally likely (note that "random" is not the same as "arbitrary").
  4. A function that yields successive powers of two starting at 1 and going up to some limit, consumed with a callback.
  5. A Go function that yields powers of an arbitrary base starting at exponent 0 and going up to some limit, consumed with a callback.
  6. A function that interleaves two lists. If the lists do not have the same length, the elements of the longer list should end up at the end of the result list.
  7. A function that stutters a sequence, i.e. (4, 3) ⇒ (4, 4, 3, 3)
2. Write a unit test suite for the package in the previous problem.

3. Write a Go program (in the file prefixes.go) that writes successive prefixes of its first input argument, one per line, starting with the first prefix, which is zero characters long.

4. Write a Go program (in the file lines.go) that reports the number of non-blank, non-commented lines in the file named by the first argument. Blank lines are those that have either no characters or consist entirely of whitespace; commented lines are those whose first non-whitespace character is the "#" character.
