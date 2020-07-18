# Background

Su Doku (Japanese meaning number place) is the name given to a popular puzzle concept. Its origin is unclear, but credit must be attributed to Leonhard Euler who invented a similar, and much more difficult, puzzle idea called Latin Squares. The objective of Su Doku puzzles, however, is to replace the blanks (or zeros) in a 9 by 9 grid in such that each row, column, and 3 by 3 box contains each of the digits 1 to 9. Below is an example of a typical starting puzzle grid and its solution grid.

```txt
# question
0 0 3|0 2 0|6 0 0
9 0 0|3 0 5|0 0 1
0 0 1|8 0 6|4 0 0
-----------------
0 0 8|1 0 2|9 0 0
7 0 0|0 0 0|0 0 8
0 0 6|7 0 8|2 0 0
-----------------
0 0 2|6 0 9|5 0 0
8 0 0|2 0 3|0 0 9
0 0 5|0 1 0|3 0 0

# solution
4 8 3|9 2 1|6 5 7
9 6 7|3 4 5|8 2 1
2 5 1|8 7 6|4 9 3
-----------------
5 4 8|1 3 2|9 7 6
7 2 9|5 6 4|1 3 8
1 3 6|7 9 8|2 4 5
-----------------
3 7 2|6 8 9|5 1 4
8 1 4|2 5 3|7 6 9
6 9 5|4 1 7|3 8 2
```

A well constructed Su Doku puzzle has a unique solution and can be solved by logic, although it may be necessary to employ "guess and test" methods in order to eliminate options (there is much contested opinion over this). The complexity of the search determines the difficulty of the puzzle; the example above is considered easy because it can be solved by straight forward direct deduction.

## Target

The 6K text file, [sudoku.txt][sudokuFile], contains fifty different Su Doku puzzles ranging in difficulty, but all with unique solutions (the first puzzle in the file is the example above).

By solving all fifty puzzles find the sum of the 3-digit numbers found in the top left corner of each solution grid; for example, 483 is the 3-digit number found in the top left corner of the solution grid above.

## How to run this program

No package dependent, just need golang environment.
`go run main.go` or `go build main.go` then run the build output.
You will see the output like bellow

```txt
Grid 1, The sum of top left 3-digit numbers: 15
Grid 2, The sum of top left 3-digit numbers: 11
Grid 3, The sum of top left 3-digit numbers: 12
Grid 4, The sum of top left 3-digit numbers: 11
Grid 5, The sum of top left 3-digit numbers: 10
Grid 6, The sum of top left 3-digit numbers: 14
Grid 7, The sum of top left 3-digit numbers: 8
Grid 8, The sum of top left 3-digit numbers: 19
Grid 9, The sum of top left 3-digit numbers: 13
Grid 10, The sum of top left 3-digit numbers: 14
Grid 11, The sum of top left 3-digit numbers: 22
Grid 12, The sum of top left 3-digit numbers: 17
Grid 13, The sum of top left 3-digit numbers: 19
Grid 14, The sum of top left 3-digit numbers: 18
Grid 15, The sum of top left 3-digit numbers: 22
Grid 16, The sum of top left 3-digit numbers: 10
Grid 17, The sum of top left 3-digit numbers: 17
Grid 18, The sum of top left 3-digit numbers: 21
Grid 19, The sum of top left 3-digit numbers: 14
Grid 20, The sum of top left 3-digit numbers: 17
Grid 21, The sum of top left 3-digit numbers: 14
Grid 22, The sum of top left 3-digit numbers: 11
Grid 23, The sum of top left 3-digit numbers: 15
Grid 24, The sum of top left 3-digit numbers: 7
Grid 25, The sum of top left 3-digit numbers: 10
Grid 26, The sum of top left 3-digit numbers: 14
Grid 27, The sum of top left 3-digit numbers: 18
Grid 28, The sum of top left 3-digit numbers: 12
Grid 29, The sum of top left 3-digit numbers: 10
Grid 30, The sum of top left 3-digit numbers: 19
Grid 31, The sum of top left 3-digit numbers: 14
Grid 32, The sum of top left 3-digit numbers: 6
Grid 33, The sum of top left 3-digit numbers: 23
Grid 34, The sum of top left 3-digit numbers: 15
Grid 35, The sum of top left 3-digit numbers: 12
Grid 36, The sum of top left 3-digit numbers: 12
Grid 37, The sum of top left 3-digit numbers: 18
Grid 38, The sum of top left 3-digit numbers: 14
Grid 39, The sum of top left 3-digit numbers: 8
Grid 40, The sum of top left 3-digit numbers: 13
Grid 41, The sum of top left 3-digit numbers: 13
Grid 42, The sum of top left 3-digit numbers: 15
Grid 43, The sum of top left 3-digit numbers: 19
Grid 44, The sum of top left 3-digit numbers: 10
Grid 45, The sum of top left 3-digit numbers: 19
Grid 46, The sum of top left 3-digit numbers: 18
Grid 47, The sum of top left 3-digit numbers: 15
Grid 48, The sum of top left 3-digit numbers: 15
Grid 49, The sum of top left 3-digit numbers: 15
Grid 50, The sum of top left 3-digit numbers: 9
```