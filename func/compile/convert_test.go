package compile

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func randS(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var str string
	length := len(alphanum)
    for i := 0; i < n; i++ {
    	a := alphanum[r.Intn(len(alphanum))% length]
    	str += string(a)
	}
    return str
    }
func genData(size int) [][]byte {
	ret := make([][]byte, size)
	for i := 0; i < size; i++ {
		ret[i] = []byte(randS(i + 1))
	}
	return ret
}
const size = 8192

var d [][]byte

func init() {
	d = genData(size)
}

func BenchmarkEqual1(b *testing.B) {
	for s := 2; s <= size; s *= 2 {
		b.Run(fmt.Sprintf("size-%v", s), func(b *testing.B){
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = comp1(d[i % s], d[(i + 1) % s])
			}
		})
	}
}

func BenchmarkEqual2(b *testing.B) {
	for s := 2; s <= size; s *= 2 {
		b.Run(fmt.Sprintf("size-%v", s), func(b *testing.B){
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = comp2(d[i % s], d[(i + 1) % s])
			}
		})
	}
}