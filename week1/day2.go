package week1

import (
	"fmt"
	"io"
	"os"
	"strings"
	"golang.org/x/tour/wc"
	"golang.org/x/tour/reader"

)

type IPAddr [4]byte

func Day2() {
	fmt.Println("================== Day 2 =================")

	// maps exercise
	wc.Test(wordCount)

	// function closures exercise
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	
	// ============= Exercise: Stringers =====================
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	// ============== Exercise: Errors =======================
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))

	// ================ Exercise: Readers ====================
	reader.Validate(MyReader{})

	// ================ Exercise: rot13Reader ================
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}

func fibonacci() func() int {
	n1 := 0
	n2 := 1
	return func() int {
		result := n1
		n1, n2 = n2, n1+n2
		return result
	}
}

func wordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		if _, ok := m[word]; !ok {
			m[word] = 1
		} else {
			m[word] += 1
		}
	}
	return m
}


// ============= Exercise: Stringers ====================

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// ============== Exercise: Errors =======================
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt2(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }

    z := 1.0
    for {
        next := z - (z*z-x)/(2*z)
        if next-z < 1e-10 && next-z > -1e-10 {
            break
        }
        z = next
    }

    return z, nil
}


// ================ Exercise: Readers ======================
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

// ================ Exercise: rot13Reader ==================
type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	if err != nil {
			return 0, err  
	}
	
	
	// 65 - 90 (A - Z) 97 - 122 (a - z)
	for i, v := range b[:n] {
		if v >= 'A' && v <= 'Z' {
			b[i] = (((v-'A')+13) % 26) + 'A'
		} else if v >= 'a' && v <= 'z' {
			b[i] = (((v-'a')+13) % 26) + 'a'
		}

	}
	return n, err
}

