package main

import (
	"bytes"
	"io"
	"log"
	"time"
)

type Follower struct {
	w io.Writer
	t time.Duration
	c *Cantor
}

func NewFollower(w io.Writer, c *Cantor) *Follower {
	return &Follower{
		w: w,
		t: configPollPeriod(),
		c: c,
	}
}

func (f *Follower) poll() {
	ticker := time.Tick(f.t)

	cache := []byte{}

	for _ = range ticker {
		data := f.c.getDisplay()

		if bytes.Equal(cache, data) {
			continue
		}

		cache = data

		_, err := f.w.Write(data)
		if err != nil {
			log.Println("connection closed: ", err.Error())
			return
		}
	}
}
