package main

// Only get what I want
// https://en.wikipedia.org/wiki/Code_page_437
// Copy from https://github.com/BenLubar/df2014/blob/master/cp437/cp437.go
var (
	CP437Raws = []rune("░▒▓│┤╡╢╖╕╣║╗╝╜╛┐└┴┬├─┼╞╟╚╔╩╦╠═╬╧╨╤╥╙╘╒╓╫╪┘┌█▄▌▐▀")
)

const (
	cp437Empty                   = " "
	cp437Vertical                = "│"
	cp437VerticalLeft            = "┤"
	cp437BottomVerticalLeft      = "┐"
	cp437TopVerticalRight        = "└"
	cp437TopVerticalLeftRight    = "┴"
	cp437BottomVerticalLeftRight = "┬"
	cp437VerticalRight           = "├"
	cp437Horizontal              = "─"
	cp437VerticalLeftRight       = "┼"
	cp437TopVerticalLeft         = "┘"
	cp437BottomVerticalRight     = "┌"
)

var (
	// 0 exist 1 not exist
	// left right top bottom
	cp437EmptyB                   = 0b0000
	cp437VerticalB                = 0b0011
	cp437VerticalLeftB            = 0b1011
	cp437BottomVerticalLeftB      = 0b1001
	cp437TopVerticalRightB        = 0b0110
	cp437TopVerticalLeftRightB    = 0b1110
	cp437BottomVerticalLeftRightB = 0b1101
	cp437VerticalRightB           = 0b0111
	cp437HorizontalB              = 0b1100
	cp437VerticalLeftRightB       = 0b1111
	cp437TopVerticalLeftB         = 0b1010
	cp437BottomVerticalRightB     = 0b0101
)

var cp437Str2B = map[string]int{
	cp437Empty:                   cp437EmptyB,
	cp437Vertical:                cp437VerticalB,
	cp437VerticalLeft:            cp437VerticalLeftB,
	cp437BottomVerticalLeft:      cp437BottomVerticalLeftB,
	cp437TopVerticalRight:        cp437TopVerticalRightB,
	cp437TopVerticalLeftRight:    cp437TopVerticalLeftRightB,
	cp437BottomVerticalLeftRight: cp437BottomVerticalLeftRightB,
	cp437VerticalRight:           cp437VerticalRightB,
	cp437Horizontal:              cp437HorizontalB,
	cp437VerticalLeftRight:       cp437VerticalLeftRightB,
	cp437TopVerticalLeft:         cp437TopVerticalLeftB,
	cp437BottomVerticalRight:     cp437BottomVerticalRightB,
}

var cp437B2Str = map[int]string{
	cp437EmptyB:                   cp437Empty,
	cp437VerticalB:                cp437Vertical,
	cp437VerticalLeftB:            cp437VerticalLeft,
	cp437BottomVerticalLeftB:      cp437BottomVerticalLeft,
	cp437TopVerticalRightB:        cp437TopVerticalRight,
	cp437TopVerticalLeftRightB:    cp437TopVerticalLeftRight,
	cp437BottomVerticalLeftRightB: cp437BottomVerticalLeftRight,
	cp437VerticalRightB:           cp437VerticalRight,
	cp437HorizontalB:              cp437Horizontal,
	cp437VerticalLeftRightB:       cp437VerticalLeftRight,
	cp437TopVerticalLeftB:         cp437TopVerticalLeft,
	cp437BottomVerticalRightB:     cp437BottomVerticalRight,
}

// Return x + y with combine rule
func cp437Plus(x, y string) string {
	bx, ok := cp437Str2B[x]
	if !ok {
		return y
	}

	by, ok := cp437Str2B[y]
	if !ok {
		return x
	}

	// Use OR because if either one has vertical -> final one has vertical and so on
	result, ok := cp437B2Str[bx|by]
	if !ok {
		return cp437Empty
	}

	return result
}
