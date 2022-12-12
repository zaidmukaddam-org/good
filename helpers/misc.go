package helpers

import (
	"net"
	"os"
	"regexp"
)

func GetHomedir() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return userHomeDir
}

func IsKanji(k string) bool {
	// see https://regex101.com/r/xhHFs2/1/codegen?language=golang
	var re = regexp.MustCompile(`[\x{3041}-\x{3096}\x{30A0}-\x{30FF}\x{3400}-\x{4DB5}\x{4E00}-\x{9FCB}\x{F900}-\x{FA6A}\x{2E80}-\x{2FD5}\x{FF5F}-\x{FF9F}\x{3000}-\x{303F}\x{31F0}-\x{31FF}\x{3220}-\x{3243}\x{3280}-\x{337F}]`)
	match := re.FindAllString(k, -1)
	if match != nil {
		return true
	}

	return false
}

func IsIp(host string) bool {
	if r := net.ParseIP(host); r == nil {
		return false
	}
	return true
}
