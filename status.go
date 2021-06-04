package gohelp

type Status string
func (s Status)StatusName(a, b, c, d string) int64 {
  if s == "a" {
		return 1
	} else if s == "b" {
		return 2
	}else if s == "c" {
		return 3
	}
	return 4
}
