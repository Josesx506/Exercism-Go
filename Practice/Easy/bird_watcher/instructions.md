### Instructions
You are an avid bird watcher that keeps track of how many birds have visited your garden. Usually you use a tally in a notebook to count the birds, but to better work with your data, you've digitalized the bird counts for the past weeks. <br>

1. Let us start analyzing the data by getting a high level view. Find out how many birds you counted in total since you started your logs.<br>

    Implement a function `TotalBirdCount` that accepts a slice of ints that contains the bird count per day. It should return the total number of birds that you counted.
    ```go
    birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
    TotalBirdCount(birdsPerDay)
    // => 34
    ```

2. Now that you got a general feel for your bird count numbers, you want to make a more fine-grained analysis. <br>

    Implement a function `BirdsInWeek` that accepts a slice of bird counts per day and a week number.<br>

    It returns the total number of birds that you counted in that specific week. You can assume weeks are always tracked completely.
    ```go
    birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
    BirdsInWeek(birdsPerDay, 2)
    // => 12
    ```

3. You realized that all the time you were trying to keep track of the birds, there was one bird that was hiding in a far corner of the garden.<br>

    You figured out that this bird always spent every second day in your garden.<br>

    You do not know exactly where it was in between those days but definitely not in your garden.<br>

    Your bird watcher intuition also tells you that the bird was in your garden on the first day that you tracked in your list.<br>

    Given this new information, write a function `FixBirdCountLog` that takes a slice of birds counted per day as an argument and returns the slice after correcting the counting mistake.
    ```go
    birdsPerDay := []int{2, 5, 0, 7, 4, 1}
    FixBirdCountLog(birdsPerDay)
    // => [3 5 1 7 5 1]
    ```