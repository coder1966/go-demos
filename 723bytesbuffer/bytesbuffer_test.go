package a0221bytesbuffer

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_bytesBuffer(t *testing.T) {
	t.Run("tt.name", func(t *testing.T) {
		bytesBuffer()
	})
}
func bytesBuffer() {

	var b bytes.Buffer
	fmt.Printf("step 1: len=%d cap=%d \n", b.Len(), b.Cap())

	b.WriteString("1234567890abcdefghijklmnopqrstuvwxyz")
	fmt.Printf("step 2: len=%d cap=%d \n", b.Len(), b.Cap())

	b = *bytes.NewBuffer([]byte("1234567890abcdefghijklmnopqrstuvwxyz"))
	fmt.Printf("step 3: len=%d cap=%d \n", b.Len(), b.Cap())

	buf := make([]byte, 5)
	for i := 0; i < 11; i++ {
		in, err := b.Read(buf)
		fmt.Printf("step 4: len=%d cap=%d in=%d  err=%V \n", b.Len(), b.Cap(), in, err)
	}

	for i := 0; i < 111; i++ {
		b.WriteString("1234567")
		fmt.Printf("step 5: len=%d cap=%d \n", b.Len(), b.Cap())
	}

	buf = make([]byte, 36)
	for i := 0; i < 111; i++ {
		in, err := b.Read(buf)
		fmt.Printf("step 6: len=%d cap=%d in=%d  err=%V \n", b.Len(), b.Cap(), in, err)
		b.WriteString("1234567890abcdefghijklmnopqrstuvwxyz")
		fmt.Printf("step 7: len=%d cap=%d \n", b.Len(), b.Cap())
	}

}
