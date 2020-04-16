func isIsomorphic(s string, t string) bool {
    l := len(s)
	s2t, t2s := make(map[byte]byte, l), make(map[byte]byte, l)
	for i := 0; i < l; i++ {
		if s2t[s[i]] == 0 && t2s[t[i]] == 0 {
			s2t[s[i]] = t[i]
			t2s[t[i]] = s[i]
		} else if s2t[s[i]] != t[i] || t2s[t[i]] != s[i]{
			return false
		}
	}
	return true
}
