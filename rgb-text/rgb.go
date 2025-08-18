package main

/*
	WOW!!! RGB TEXT!?!?!?!?!?
	AAYAYYAYAYAYYAYAY

	What am I saying.
	Anyways, this one is just pure boredom.
	but tis cool alr?

	k bai!
*/

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func hsvToRGB(h, s, v float64) (int, int, int) {
	if s == 0 {
		gray := int(v * 255)
		return gray, gray, gray
	}
	h = math.Mod(h, 360)
	if h < 0 {
		h += 360
	}
	c := v * s
	x := c * (1 - math.Abs(math.Mod(h/60.0, 2)-1))
	m := v - c

	var r1, g1, b1 float64
	switch {
	case h < 60:
		r1, g1, b1 = c, x, 0
	case h < 120:
		r1, g1, b1 = x, c, 0
	case h < 180:
		r1, g1, b1 = 0, c, x
	case h < 240:
		r1, g1, b1 = 0, x, c
	case h < 300:
		r1, g1, b1 = x, 0, c
	default:
		r1, g1, b1 = c, 0, x
	}

	r := int((r1 + m) * 255)
	g := int((g1 + m) * 255)
	b := int((b1 + m) * 255)
	return r, g, b
}

func main() {
	// Use Program's Args as string to display, if not exists it just displays THY END IS NOW!
	// Minos Prime.
	text := strings.Join(os.Args[1:], " ")
	if strings.TrimSpace(text) == "" {
		text = "THY END IS NOW!"
	}

	fmt.Print("\x1b[?25l")
	defer fmt.Print("\x1b[0m\x1b[?25h\n")

	// Makes sure the Terminal "Cursor" doesn't disappear when quitting
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigc

		fmt.Print("\x1b[0m\x1b[?25h\n\033[2J\033[H")
		os.Exit(0)
	}()

	frame := 0
	for {

		fmt.Print("\r\x1b[K")
		for i, r := range text {

			h := math.Mod(float64(i)*12+float64(frame)*4, 360)
			rr, gg, bb := hsvToRGB(h, 1.0, 1.0)

			fmt.Printf("\x1b[38;2;%d;%d;%dm%c", rr, gg, bb, r)
		}
		fmt.Print("\x1b[0m")
		time.Sleep(20 * time.Millisecond)
		frame++
	}
}
