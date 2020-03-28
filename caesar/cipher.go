package caesar

import "strings"

func Encrypt(unencrypted string, rotations int) string {
	rotate := func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+int32(rotations))%26
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+int32(rotations))%26
		default:
			return r
		}
	}

	return strings.Map(rotate, unencrypted)
}
