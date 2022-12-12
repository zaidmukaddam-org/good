package helpers

import (
	"io"
	"os"
	"strings"
)

type rot13 struct {
	r io.Reader
}

func Str2leet(t string) string {
	r := strings.NewReplacer("E","3","e","3","i","1","I","1","l","1","L","1","O","0","o","0","t","7","T","7","a","4","A","4","s","5","S","5")
	return r.Replace(t)
}

func BrokenCipher(key string, t string) string {
	smsg, skey := s(t), s(key)
	out := make([]rune, 0, len(t))
	for i, v := range smsg {
		out = append(out, e(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}

func BrokenDeCipher(key string, cipher string) string {
	smsg, skey := s(cipher), s(key)
	out := make([]rune, 0, len(cipher))
	for i, v := range smsg {
		out = append(out, d(v, rune(skey[i%len(skey)])))
	}
	return string(out)
}

func s(in string) string {
	out := []rune{}
	for _, v := range in {
		if 65 <= v && v <= 90 {
			out = append(out, v)
		} else if 97 <= v && v <= 122 {
			out = append(out, v-32)
		}
	}

	return string(out)
}

func e(x, y rune) rune {
	return (((x - 'A') + (y - 'A')) % 26) + 'A'
}

func d(x, y rune) rune {
	return (((((x - 'A') - (y - 'A')) + 26) % 26) + 'A')
}

func (r rot13) Read(y []byte) (int, error) {
	max, e := r.r.Read(y)
	for i := 0; i < max; i++ {
		switch {
		case 'a' <= y[i] && y[i] <= 'z':
			y[i] = 'a' + (y[i]-'a'+13)%26
		case 'A' <= y[i] && y[i] <= 'Z':
			y[i] = 'A' + (y[i]-'A'+13)%26
		}
	}
	return max, e
}

func RotReader(str string) {
	s := strings.NewReader(str)
	r := rot13{s}
	io.Copy(os.Stdout, &r)
}
