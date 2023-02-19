package strings_lib

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RndStr struct {
	sync.Mutex
	Data string
}

type RndInt struct {
	sync.Mutex
	Data int
}

func NewRndStr() *RndStr {
	return &RndStr{Data: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"}
}

func NewRndInt() *RndInt {
	return &RndInt{Data: 1234567890}
}

func (r *RndStr) RandString(n int) string {
	r.Lock()
	defer r.Unlock()
	rand.Seed(time.Now().UnixNano())
	chars := []rune(r.Data)
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func (r *RndInt) RandInt(n int) string {
	r.Lock()
	defer r.Unlock()
	rand.Seed(time.Now().UnixNano())
	swap := strconv.Itoa(r.Data)
	chars := []rune(swap)
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
