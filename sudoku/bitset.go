package sudoku

import (
	"fmt"
	"math/big"
	"strings"
)

type bitSet struct {
	s big.Int
}

func (bs *bitSet) Set(i int, b uint) {
	bs.s.SetBit(&bs.s, i, b)
}
func (bs *bitSet) Reset() {
	bs.s.SetInt64(0)
}

func (bs *bitSet) Get(i int) uint {
	return bs.s.Bit(i)
}

func (bs *bitSet) Len() int {
	return bs.s.BitLen()
}

func (bs *bitSet) String() string {
	sb := strings.Builder{}
	for i := 0; i < bs.Len(); i++ {
		sb.WriteString(fmt.Sprintf("%d: %d  ", i, bs.Get(i)))
	}
	return sb.String()
}
