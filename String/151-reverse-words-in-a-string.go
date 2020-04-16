func reverseWords(s string) string {
	bs := make([]byte, 0)
	space := byte(' ')
	// build byte list and handle illegal spaces
	for i := 0; i < len(s); {
		if i < len(s)-1  && s[i] == space && s[i+1] == space {
			i++
			continue
		}
		if i == 0 && s[i] == space {
			i++
			continue
		}
		if i == len(s)-1 && s[i] == space {
			i++
			continue
		}
		bs = append(bs, s[i])
        i++
	}
    // this is f**king ugly but it works
    if len(bs) > 0 && bs[0] == space {
        bs = bs[1 : ]
    }
	// reverse the whole byte list
	reverseByteList(bs)
	// reverse each word to correct order
	start := 0
	for end, b := range bs {
		if b == space {
			reverseByteList(bs[start : end])
			start = end+1
		}
        if end == len(bs)-1 {
            reverseByteList(bs[start : end+1])
        }
	}
    return string(bs)
}

func reverseByteList(s []byte)  {
    l, r := 0, len(s)-1
    for l < r {
        s[l], s[r] = s[r], s[l]
        l++
        r--
    }
    return
}
