package paths

func indent(amount int) string {
	if amount < 1 {
		return ""
	} else if amount == 1 {
		return "\t| "
	} else {
		out := "\t|"
		for i := 1; i < amount; i++ {
			out = out + "\t"
		}
		out = out + "> "
		return out
	}
}
