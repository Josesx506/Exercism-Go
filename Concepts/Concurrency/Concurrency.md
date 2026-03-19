
### Channels
[Explainer video](https://www.youtube.com/watch?v=LvgVSSpwND8) <br>

Channels are ways for go routines to communicate with each other. They behave like a pipe that allow you to send or receive a message. Channels have a type e.g. *string*, *int32* etc
```go
ch =  make(chan string);
```
The channel type impacts what type of messages can be transmitted e.g. a string channel can transmit strings
```go
func count(thing string, c chan string) {
    for i:=1; i<=5; i++ {
        c <- thing // transmit message across channel
        time.Sleep(time.Millisecond * 500)
    }

    // Close channel after all messages to prevent deadlock 
    // on the receiver routine
    close(c) 
}

func main() {
    c := make(chan string)
    go count("sheep", c)

    // Receive single message from channel and check 
    // if channel is open
    msg, open := <- c 
    fmt.Printf(msg)

    // Iterate over channel and receive all messages
    for msg := range c {
        fmt.Println(msg)
    }
}
```
Transmitting and receiving messages across channels is a blocking operation.

### Buffered channels
Channels can also be buffered to receive several mesages until they're full
```go
func main() {
    c := make(chan string, 2) // Channel that'll allow 2 entries
    c <- "hello"
    c <- "world"
    // Messages above are non-blocking but channel is full
    // If we add an additional message, the channel becomes blocked
    // and encounters deadlock
}
```

### `SELECT` statements
When receiving data from multiple channels, because each message you receive is 
a blocking operation, a routine would only run as fast as the slowest routine
```go
func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for { // loop infinitely
            c1 <- "Every 500ms"
            time.Sleep(time.Millisecond * 500)
        }
    }

    go func() {
        for { // loop infinitely
            c1 <- "Every 2secs"
            time.Sleep(time.Second * 2)
        }
    }

    for { // Receive infinitely in main routine
        fmt.Println(<- c1)
        fmt.Println(<- c2)
    }
}
```
In the example above, the routine would wait every 2 seconds to receive messages 
because channel 2 always blocks each time it's waiting to receive a message. 
`select` statements fix this issue
```go
func main() {
    ...
    for {
        select {
            case msg1 := <- c1:
                fmt.Println(msg1)
            case msg2 := <- c2:
                fmt.Println(msg2)
        }
    }
}
```
With the select statements, the receivers become none-blocking, and `c1` will 
transmit four messages since it runs in 500ms before `c2` receiver consumes its 
message.

> [!Note]
> Channels are bidirectional (send and receive) by default, but we can specify them 
as send only / receive only when defining them.

### Worker pools
To run multiple jobs in parallel from a queue, we can use worker pools
```go
func fib(n int) int {
    if (n <= 1) {
        return 1
    } else {
        return fib(n-1) + fib(n-2)
    }
}

// uni-directional channels
// jobs can only receive
// results can only send
func worker(jobs <-chan int, results chan<- int) {
    for n := range jobs {
        results <- fib(n)
    }
}

func main() {
    // Buffered channels
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Worker pools are fired into the background waiting for messages
    go worker(jobs, results) // worker 1
    go worker(jobs, results) // worker 2
    go worker(jobs, results) // worker 3
    go worker(jobs, results) // worker 4

    for i:= 0; i<100; i++ {
        jobs <- i // Send 100 messages into the jobs buffered channel
    }
    close(jobs)

    for j:=0; j<100; j++ {
        fmt.Println(<-results) // receive fibonacci numbers from results channel
    }
    close(results)
}
```

With the program above, the function will utilize four workers to calculate the 
fibonacci numbers from 1 to 100. Each additional routine is an independent worker 
and the results will return as they're concluded and not necessarily in the correct 
order as expected from ditributed processes