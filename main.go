package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/coffemanfp/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	if apiKey == "" {
		log.Fatalln("no Big Hugh Thesaurus api key found")
	}

	thesaurusClient := &thesaurus.BigHugh{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := s.Text()
		syns, err := thesaurusClient.Synonyms(word)
		if err != nil {
			err = fmt.Errorf("failed when looking for synonyms for \"%s\":\n%s", word, err)
			log.Fatalln(err)
		}

		if len(syns) == 0 {
			err = fmt.Errorf("couldn't find any synonyms for \"%s\":\n%s", word, err)
			log.Fatalln(err)
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
