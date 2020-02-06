package mypackage

var pop map[string]string
var PopList map[string]string

func init() {
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallin"
	pop["John Legent"] = "All of Me"

	PopList = pop
}

func GetMusic()
