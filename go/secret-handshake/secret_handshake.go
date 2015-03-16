package secret

var decode = [4]string{"wink", "double blink", "close your eyes", "jump"}

func prepend(code []string, cmd string) []string {
	return append([]string{cmd}, code...)
}

// I guess append is a builtin, not a function
// So this is just a shim
func myappend(code []string, cmd string) []string {
	return append(code, cmd)
}

func Handshake(n int) (code []string) {
	if n <= 0 || n >= 32 {
		return nil
	}

	addfunc := myappend
	if n&16 != 0 {
		addfunc = prepend
	}
	for i, cmd := range decode {
		if n&(1<<uint(i)) != 0 {
			code = addfunc(code, cmd)
		}
	}
	return
}
