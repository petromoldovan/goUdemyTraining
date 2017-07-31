package main

import (
	"net/http"
	"log"
	"fmt"
	"bufio"
)

func main() {
	//get the book
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	//set the split function
	scanner.Split(bufio.ScanWords)

	//create slice of slices of strings to store words. Use 12 because we want to have only 12 buckets
	buckets := make([][]string, 12)

	//create slices to hold words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	//loop over words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	//print length of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, "--", len(buckets[i]))
	}

}

func HashBucket(word string, numBuckets int) int {
	var sum int

	for _, v := range word {
		sum += int(v)
	}
	
	return sum % numBuckets
}
