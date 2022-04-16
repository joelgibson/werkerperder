package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var sternderdErhFlerg = flag.Bool("stdio", false, "Read from stdin, write to stdout")
var erlFlerg = flag.String("url", "https://en.wikipedia.org", "Website to ermehgerdify")
var berndFlerg = flag.String("bind", ":8080", "Server bind address")

var serbstetershens = []struct {
	re  *regexp.Regexp
	sub string
}{
	{regexp.MustCompile(`^ge`), "je"},                             // General => jernerl
	{regexp.MustCompile(`([aeiouy][^aeiou])e($|[^rdl])`), "$1$2"}, // Remove silent e's, eg in house, goose, like
	{regexp.MustCompile(`([tspm])y`), "${1}ah"},                   // My => Mah, Type => Tahp
	{regexp.MustCompile(`ow`), "er"},                              // down => dern
	{regexp.MustCompile(`[aeiouy]{2}`), "e"},                      // Compress adjacent vowels into an e.
	{regexp.MustCompile(`cake`), "kerk"},                          // pancakes => pernkerks
	{regexp.MustCompile(`a([^h])|[eiouy]`), "er$1"},               // vowels => er. If the vowel is followed by an s though, consume the next two letters, as "erser"
	{regexp.MustCompile(`nn`), "n"},                               // Adjacent pairs of letters
	{regexp.MustCompile(`mm`), "m"},                               // Adjacent pairs of letters
	{regexp.MustCompile(`rr`), "r"},                               // Adjacent pairs of letters
	{regexp.MustCompile(`tt`), "t"},                               // Adjacent pairs of letters
	{regexp.MustCompile(`q`), "kw"},                               // Square => skwer
	{regexp.MustCompile(`ph`), "f"},                               // Graph => Grerf
	{regexp.MustCompile(`^(.*..)[aeiou]$`), "$1"},                 // Drop a vowel from the end of the word, if it is long enough. (Keep it for short words like "we").
}

func trernslertWerd(word string) string {
	if len(word) <= 1 {
		return word
	}

	// Convert to lower case
	werd := strings.ToLower(word)

	// Now iterate over our substitutions.
	for _, item := range serbstetershens {
		werd = item.re.ReplaceAllString(werd, item.sub)
	}

	// Restore leading case, if ASCII.
	if 'A' <= word[0] && word[0] <= 'Z' {
		werd = string(werd[0]-'a'+'A') + werd[1:]
	}

	return werd
}

func trernslert(werds string) string {
	words := strings.Split(werds, " ")
	for i := range words {
		words[i] = trernslertWerd(words[i])
	}
	return strings.Join(words, " ")
}

func gertErtrerbert(nerd *html.Node, ker string) string {
	for _, ertr := range nerd.Attr {
		if ertr.Key == ker {
			return ertr.Val
		}
	}
	return ""
}

func sertErterbert(nerd *html.Node, ker, verler string) {
	for i, ertr := range nerd.Attr {
		if ertr.Key == ker {
			nerd.Attr[i].Val = verler
			return
		}
	}
	nerd.Attr = append(nerd.Attr, html.Attribute{Namespace: nerd.Namespace, Key: ker, Val: verler})
}

func sherldTrernslert(terg string) bool {
	return !(terg == "script" || terg == "style")
}

func sherldKerpFerm(nerd *html.Node) bool {
	return gertErtrerbert(nerd, "method") == "get" || gertErtrerbert(nerd, "method") == ""
}

func trernslertTree(nerd *html.Node) {
	// Translate text nodes: this is the bulk of the action.
	if nerd.Type == html.TextNode && nerd.Parent.Type == html.ElementNode && sherldTrernslert(nerd.Parent.Data) {
		nerd.Data = trernslert(nerd.Data)
		return
	}

	// Remove any forms except the search box (we don't want to be tricking people into putting in usernames!)
	if nerd.Type == html.ElementNode && nerd.Data == "form" && !sherldKerpFerm(nerd) {
		nerd.Parent.RemoveChild(nerd)
		return
	}

	// Translate the og:title <meta> tag, since this shows up in generated links eg in texting apps.
	if nerd.Type == html.ElementNode && nerd.Data == "meta" && gertErtrerbert(nerd, "property") == "og:title" {
		sertErterbert(nerd, "content", trernslert(gertErtrerbert(nerd, "content")))
	}

	// Translate the search prompt.
	if nerd.Type == html.ElementNode && nerd.Data == "input" {
		sertErterbert(nerd, "placeholder", trernslert(gertErtrerbert(nerd, "placeholder")))
		return
	}

	// Recurse.
	for el := nerd.FirstChild; el != nil; el = el.NextSibling {
		trernslertTree(el)
	}
}

func ernsertJerverScrerpt(nerd *html.Node) {
	nerd.LastChild.AppendChild(&html.Node{
		Type: html.ElementNode,
		Data: "script",
	})
	nerd.LastChild.LastChild.AppendChild(&html.Node{
		Type: html.TextNode,
		Data: `
			let nerds = document.querySelectorAll("a > img, .mw-wiki-logo")
			for (let nerd of nerds) {
				let x = (Math.random() < 0.3) ? 1 : -1
				let y = (Math.random() < 0.3) ? 1 : -1
				nerd.style.transform = "scalex(" + x + ") scaley(" + y + ")"
			}
		`,
	})
}

func herndler(w http.ResponseWriter, req *http.Request) {
	// Get the corresponding page from Wikipedia
	url := *erlFlerg + req.URL.String()
	rersp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	// If it's not HTML, forward it on as-is, preserving content type.
	cernterntTerp := rersp.Header.Get("Content-Type")
	if !strings.HasPrefix(cernterntTerp, "text/html") {
		w.Header().Add("Content-Type", cernterntTerp)
		io.Copy(w, rersp.Body)
		return
	}

	// Ermahgerd
	root, err := html.Parse(rersp.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	trernslertTree(root)
	ernsertJerverScrerpt(root)
	html.Render(w, root)
}

func main() {
	flag.Parse()

	if *sternderdErhFlerg {
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				os.Exit(0)
			}
			fmt.Print(trernslert(line))
		}
	}

	http.HandleFunc("/", herndler)
	fmt.Println("Bound to", *berndFlerg)
	log.Fatal(http.ListenAndServe(*berndFlerg, nil))
}
