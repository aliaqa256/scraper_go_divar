package main

import (
	"os"
	"strings"
	"github.com/anaskhan96/soup"
)


func main() {
	ch := make(chan string)


	go getpage(ch)
	res := <-ch
	doc := soup.HTMLParse(res)
	cards := doc.FindAll("div","class","waf972")
	// slice of maps
	sliceOfMaps := make([]map[string]string, 0)

	for _, card := range cards {
		title := card.Find("h2","class","kt-post-card__title").FullText()
		description := card.FindAll("div","class","kt-post-card__description")
		var allDescription string
		for _, desc := range description {
			allDescription += desc.FullText()+"\n"
		}
		n := card.Find("span","class","kt-post-card__bottom-description").FullText()
		nigborhood:= strings.Split(n, " ")
		var indexOfdar int
		for i, v := range nigborhood {
			if v == "در" {
				indexOfdar = i+1
			}
		}
		address := strings.Join(nigborhood[indexOfdar:], " ")
		link := "https://divar.ir"+card.Find("a").Attrs()["href"]

		mapOfdata:=map[string]string{
			"title": title,
			"description": allDescription,
			"address": address,
			"link": link,
		}
		sliceOfMaps = append(sliceOfMaps, mapOfdata)
	}



	f, err := os.Create("data.csv")
	must(err)
	defer f.Close()
	for _, row := range sliceOfMaps {
		sliceOfstrings := make([]string, 0)
		for _, value := range row {
			sliceOfstrings = append(sliceOfstrings, strings.Trim(value,"" ) )
		}
		f.WriteString(strings.Join(sliceOfstrings, ","))
	}	
	
}

func getpage(ch chan string)  {
	str,err:=soup.Get("https://divar.ir/s/tehran/real-estate?user_type=personal")
	must(err)
	ch <- str
}


func must(err error) {
	if err != nil {
		panic(err)
	}
}

