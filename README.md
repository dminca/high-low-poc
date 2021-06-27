# High-low number return
> Experiment to return the highest and lowest number when
> given a string of space separated numbers

Taken from codewars https://www.codewars.com/kumite/new?group_id=5bd0b42721a99d54cc00305d&review_id=5bd029a0abc6619b260001e2

In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

## Example:

```go
highAndLow("1 2 3 4 5");  // return "5 1"
highAndLow("1 2 -3 4 5"); // return "5 -3"
highAndLow("1 9 3 4 -5"); // return "9 -5"
```

```go
Kata.HighAndLow("1 2 3 4 5");  // return "5 1"
Kata.HighAndLow("1 2 -3 4 5"); // return "5 -3"
Kata.HighAndLow("1 9 3 4 -5"); // return "9 -5"
```

```go
highAndLow "1 2 3 4 5"  // return "5 1"
highAndLow "1 2 -3 4 5" // return "5 -3"
highAndLow "1 9 3 4 -5" // return "9 -5"
```

```go
highAndLow("1 2 3 4 5");  // return "5 1"
highAndLow("1 2 -3 4 5"); // return "5 -3"
highAndLow("1 9 3 4 -5"); // return "9 -5"
```

```go
highAndLow("1 2 3 4 5");  // return "5 1"
highAndLow("1 2 -3 4 5"); // return "5 -3"
highAndLow("1 9 3 4 -5"); // return "9 -5"
```

```go
highAndLow("1 2 3 4 5");  // return "5 1"
highAndLow("1 2 -3 4 5"); // return "5 -3"
highAndLow("1 9 3 4 -5"); // return "9 -5"
```

```go
highAndLow("1 2 3 4 5")  # return "5 1"
highAndLow("1 2 -3 4 5") # return "5 -3"
highAndLow("1 9 3 4 -5") # return "9 -5"
```

```go
high_and_low("1 2 3 4 5")  # return "5 1"
high_and_low("1 2 -3 4 5") # return "5 -3"
high_and_low("1 9 3 4 -5") # return "9 -5"
```

```go
high_and_low("1 2 3 4 5")  # return "5 1"
high_and_low("1 2 -3 4 5") # return "5 -3"
high_and_low("1 9 3 4 -5") # return "9 -5"
```

```go
high_and_low("1 2 3 4 5")  # return "5 1"
high_and_low("1 2 -3 4 5") # return "5 -3"
high_and_low("1 9 3 4 -5") # return "9 -5"
```

```go
high_and_low("1 2 3 4 5")  # return "5 1"
high_and_low("1 2 -3 4 5") # return "5 -3"
high_and_low("1 9 3 4 -5") # return "9 -5"
```

```go
highAndLow("1 2 3 4 5")  // return "5 1"
highAndLow("1 2 -3 4 5") // return "5 -3"
highAndLow("1 9 3 4 -5") // return "9 -5"
```

```go
highAndLow "1 2 3 4 5")  # return "5 1"
highAndLow "1 2 -3 4 5") # return "5 -3"
highAndLow "1 9 3 4 -5") # return "9 -5"
```

```go
HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"
```

```go
highAndLow("1 2 3 4 5")  // return "5 1"
highAndLow("1 2 -3 4 5") // return "5 -3"
highAndLow("1 9 3 4 -5") // return "9 -5"
```

```go
Kata.high_and_low("1 2 3 4 5")  # return "5 1"
Kata.high_and_low("1 2 -3 4 5") # return "5 -3"
Kata.high_and_low("1 9 3 4 -5") # return "9 -5"
```

## Notes:

* All numbers are valid Int32, no need to validate them.
* There will always be at least one number in the input string.
* Output string must be two numbers separated by a single space, and highest number is first.
