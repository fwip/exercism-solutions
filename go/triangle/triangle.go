package triangle

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	// Each side must be > 0 and less than half the total length
	max_len := (a + b + c) / 2
	if !(a > 0 && b > 0 && c > 0 && a < max_len && b < max_len && c < max_len) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}
