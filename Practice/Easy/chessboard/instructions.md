### Instructions
1. Given a chessboard and a file(row), count how many squares(columns) are occupied <br>
    Implement the` CountInFile(board Chessboard, file string) int` function. It should 
    count the total number of occupied squares by ranging over a map. Return an integer. 
    Return a count of zero (0) if the given file cannot be found in the map.
    ```go
    CountInFile(board, "A")
    // => 3
    ```

2. Given a chessboard and a rank(column), count how many squares(rows) are occuppied <br>
    Implement the `CountInRank(board Chessboard, rank int) int` function. It should count 
    the total number of occupied squares by ranging over the given rank. Return an integer. Return a count of zero (0) if the given rank is not a valid one (not between 1 and 8, inclusive).
    ```go
    CountInRank(board, 2)
    // => 1
    ```

3. Count how many squares are present in a given chessboard <br>
    Implement the `CountAll(board Chessboard) int` function. It should count how many 
    squares are present in the chessboard and returns an integer. Since you don't need to 
    check the content of the squares, consider using range omitting both index and value.
    ```go
    CountAll(board)
    // => 64
    ```

4. Count how many squares are occuppied in a given chessboard <br>
    Implement the CountOccupied(board Chessboard) int function. It should count how many squares are occupied in the chessboard. Return an integer.
    ```go
    CountOccupied(board)
    // => 15
    ```