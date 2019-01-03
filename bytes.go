package bytes

func IndexQuoted(b []byte, esc, quote byte) (int, int, bool) {
	n := len(b)
	if n < 2 {
		return 0, 0, false
	}
	s := 0
	i := 0
	for j := 0; j < n; j++ {
		switch b[j] {
		case esc:
			j++
		case quote:
			switch s {
			case 0:
				s = 1
				i = j
			case 1:
				return i, j, true
			}
		}
	}
	return 0, 0, false
}

func IndexScoped(b []byte, esc, quote, open, close byte) (int, int, bool) {
	n := len(b)
	if n < 2 {
		return 0, 0, false
	}
	s := 0
	i := 0
	q := 0
	for j := 0; j < n; j++ {
		switch b[j] {
		case esc:
			j++
		case quote:
			switch s {
			case 0:
				s = 1
			case 1:
				s = 0
			case 2:
				s = 3
			case 3:
				s = 2
			}
		case open:
			switch s {
			case 0:
				s = 2
				i = j
				q = 1
			case 2:
				q++
			}
		case close:
			switch s {
			case 2:
				q--
				if q == 0 {
					return i, j, true
				}
			}
		}
	}
	return 0, 0, false
}

type IndexFunc func([]byte) (int, int, bool)

type EachFunc func([]byte) error

func IndexForEach(b []byte, index IndexFunc, each EachFunc) (p int, err error) {
	for {
		i, j, ok := index(b[p:])
		if !ok {
			break
		}
		err = each(b[p+i+1 : p+j])
		if err != nil {
			return
		}
		p += j + 1
	}
	return
}
