package lib

const (
	black   = "\033[30m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
	reset   = "\033[0m"
)

var (
	md_strongem = red     // ***strongem***
	md_strong   = yellow  // **strong**
	md_emph     = magenta // *emph*
)

type State struct {
	Pos      int
	Strongem bool
	Strong   bool
	Emph     bool
	Code     bool
	Dquote   bool
	Quote    bool
	Link     bool
}

func (s State) String() string { return toYAML(s) }
func (s State) JSON() string   { return toJSON(s) }
func (s State) YAML() string   { return toYAML(s) }

func Mark(in string) string {
	out := ""
	state := new(State)

	for state.Pos = 0; state.Pos < len(in); state.Pos++ {

		// ***strongem***

		if in[state.Pos:state.Pos+3] == "***" {
			println("<strongem>")
		}

	}

	return out
}
