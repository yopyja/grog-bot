package remove 


func Remove(l []string, s string) []string {
	for i := 0; i < len(l); i++ {
		if l[i] == s {
			l[i] = l[len(l)-1]
			l = l[:len(l)-1]
			return l
		}
	}
	return l
}