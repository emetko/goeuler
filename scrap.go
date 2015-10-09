package main
import (
	"net/http"
	"fmt"
	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
	"bytes"
	"strings"
	"os"
	"text/template"
	"bufio"
)

func main() {
	const FROM,TO = 26,29

	for i:=FROM;i<=TO;i++ {
		genSolutionFile(i)
	}
}

type problem struct {
	Id          string
	URL         string
	Title       string
	Description string
	Answer      string
}

func genSolutionFile(num int) {
	templ, _ := template.ParseFiles("resources/solutionTemplate")
	answers := getAnswers()

	info := getProblem(num)
	info.Answer = answers[num-1]
	f, err := os.Create("solutions/" + strings.ToLower(info.Id) + "_test.go")
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	templ.Execute(&buf, info)
	f.WriteString(string(buf.Bytes()))
	fmt.Print(info)
}

func getProblem(num int) (p problem) {
	url := "https://projecteuler.net/problem=" + fmt.Sprintf("%v", num);
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	dom, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	content := cascadia.MustCompile("#content").MatchFirst(dom)

	p.Title = getText(cascadia.MustCompile("h2").MatchFirst(content))
	p.Description = getText(cascadia.MustCompile(".problem_content").MatchFirst(content))
	p.URL = url
	p.Id = fmt.Sprintf("Euler%03d", num)

	return p

}


/*
  getText returns the inner text of an HTML node.
  also try to replace some math/HTML representations with ascii notation.
*/
func getText(n *html.Node) string {
	var buf bytes.Buffer
	html.Render(&buf, n)
	nodeText := string(buf.Bytes())

	nodeText = strings.Replace(nodeText,"<sup>","^",-1)
	nodeText = strings.Replace(nodeText,"<br/>","\n",-1)
	ts := strings.Index(nodeText, "<")
	te := strings.Index(nodeText, ">")
	for ts > -1 {
		nodeText = nodeText[:ts] + nodeText[te + 1:]
		ts = strings.Index(nodeText, "<")
		te = strings.Index(nodeText, ">")
	}

	nodeText = strings.Replace(nodeText,"&gt;",">",-1)
	nodeText = strings.Replace(nodeText,"&lt;","<",-1)

	return strings.TrimRight(nodeText,"\n")
}

func getAnswers() []string {
	inFile, err := os.Open("resources/answers")
	defer inFile.Close()
	if err != nil {
		fmt.Println("[ERROR] Could not read answers file: " + err.Error())
	}
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	answers := make([]string, 0)
	for scanner.Scan() {
		solInfo := strings.Split(scanner.Text(), "|")
		answers = append(answers, solInfo[1])
	}
	return answers
}

