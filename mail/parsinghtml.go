package mail

import (
	"fmt"

	. "github.com/julvo/htmlgo"
	a "github.com/julvo/htmlgo/attributes"
)

func main() {
    var numberDivs HTML
    for i := 0; i < 3; i++ {
        numberDivs += Div(Attr(a.Style_("font-family:monospace;")),
                          Text(i))
    }

    page :=
        Html5_(
            Head_(),
            Body_(
                H1_(Text("Welcome <script>")),
                numberDivs,
                Div(Attr(a.Dataset("hello", "htmlgo"))),
                Script_(JavaScript("alert('This is escaped');")),
                Script_(JavaScript("This is escaped", "alert({{.}});"))))

    fmt.Println(page)
}