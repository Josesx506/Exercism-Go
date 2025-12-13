package main

import "fmt"

type ElectionResult struct {
	Name  string // Name of the candidate
	Votes int    // Votes of votes the candidate had
}

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	// panic("Please implement the NewVoteCounter() function")
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	// panic("Please implement the VoteCount() function")
	if counter == nil {
		return 0
	} else {
		return *counter
	}
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	// panic("Please implement the IncrementVoteCount() function")
	*counter = *counter + increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	// panic("Please implement the NewElectionResult() function")
	result := &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
	return result

}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	// panic("Please implement the DisplayResult() function")
	display := fmt.Sprintf("%s (%d)", result.Name, result.Votes)
	return display
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	// panic("Please implement the DecrementVotesOfCandidate() function")
	results[candidate]--
}

func main() {
	twoVotes := 2
	fmt.Println(VoteCount(&twoVotes))
	IncrementVoteCount(&twoVotes, 3)
	fmt.Println(twoVotes)
	results := NewElectionResult("Peter", 33)
	fmt.Println(*results)
	fmt.Println(DisplayResult(results))
}
