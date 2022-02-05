package main

import (
	"fmt"
	"strings"
	"time"
)

var csett = "abcdefghijklmnopqrstuvwxyz0123456789!@#ABCDEFGHIJKLMNOPQRSTUVWXYZ$%^&*"
var cset = strings.Split(csett, "")
var csetLen = len(cset)
var stt = time.Now().UnixMilli()
var ch = make(chan struct{})

func main() {
	pass := "thftgr"
	for i := 1; i < 20; i++ {
		st := time.Now().UnixMilli()
		bpass := make([]string, i+1)
		for j := 0; j < csetLen; j++ {
			j := j
			go func() {
				find(i, pass, bpass, cset[j])
			}()

		}
		for j := 0; j < csetLen; j++ {
			_ = <-ch
		}
		fmt.Printf("len:%d %dms. next len:%d will %dms\n", i, time.Now().UnixMilli()-st, i+1, (time.Now().UnixMilli()-st)*int64(csetLen))

	}
}

func find(idx int, pass string, bass []string, tset string) {
	if tset != "" {
		bass = make([]string, idx+1)
		bass[idx] = tset
		if pass == strings.Join(bass, "") {
			fmt.Println("====================")
			fmt.Printf("%d '%s' %d %s\n", len(bass), strings.Join(bass, ""), time.Now().UnixMilli()-stt, "ms")
			panic("found") // throw new Exepction();
		}
		if idx > 0 {
			find(idx-1, pass, bass, "")
		}
		ch <- struct{}{}
		return
	}
	for i := 0; i < csetLen; i++ {

		bass[idx] = cset[i]
		if pass == strings.Join(bass, "") {
			fmt.Println("====================")
			fmt.Printf("%d '%s' %d %s\n", len(bass), strings.Join(bass, ""), time.Now().UnixMilli()-stt, "ms")
			panic("found") // throw new Exepction();
		}
		if idx > 0 {
			find(idx-1, pass, bass, "")
		}
	}
}
