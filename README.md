# Go Porter Stemmer

A native Go clean room implementation of the Porter Stemming Algorithm.

This is NOT a port. This is a native Go implementation from the human-readable
description of the algorithm.

I've tried to make it (more) efficient by NOT internally using string's, but
instead internally using []rune's and using the same (array) buffer used by
the []rune slice (and sub-slices) at all steps of the algorithm.

For Porter Stemmer algorithm, see:

http://tartarus.org/martin/PorterStemmer/def.txt

http://tartarus.org/martin/PorterStemmer/

Also, since when I initially implemented it, it failed the tests at...

http://tartarus.org/martin/PorterStemmer/voc.txt

http://tartarus.org/martin/PorterStemmer/output.txt

... I then looked at the source code here...

http://tartarus.org/martin/PorterStemmer/c.txt

... and noticed that there are some items marked as "DEPARTURE",
which differ from the original algorithm.

Implemented these departures too.

## Usage

To use this Golang library, use with something like:

    package main
    
    import (
      "fmt"
      "github.com/reiver/go-porterstemmer"
    )
    
    func main() {
    
      word := "waxes"
      
      stem := porterstemmer.StemString(word)
      
      fmt.Printf("The word [%s] has the stem [%s].", word, stem)
    }

