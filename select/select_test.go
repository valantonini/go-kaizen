package _select

import (
	"github.com/matryer/is"
	"testing"
	"time"
)

func writeToChanAfterDelay(c chan string, value string, delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	c <- value
}

func Test_Select(t *testing.T) {
	Is := is.New(t)

	t.Run("select will wait on multiple channels until the first response", func(t *testing.T) {
		chan1 := make(chan string)
		chan2 := make(chan string)

		go writeToChanAfterDelay(chan1, "chan1", 10)
		go writeToChanAfterDelay(chan2, "chan2", 5)

		got := ""

		// select will execute when the first channel in the case statement receives a value
		select {
		case str := <-chan1:
			got = str
		case str := <-chan2:
			got = str
		}

		// ensure we have given enough time for the first channel to finish proving the select executed once
		time.Sleep(20 * time.Millisecond)

		Is.Equal(got, "chan2")
	})

	t.Run("timeout can be used to ensure select doesn't block forever", func(t *testing.T) {
		chan1 := make(chan string)

		go writeToChanAfterDelay(chan1, "chan1", 250)

		got := ""
		select {
		case str := <-chan1:
			got = str
		case _ = <-time.After(5 * time.Millisecond):
			got = "timeout"
		}

		Is.Equal(got, "timeout")
	})

	t.Run("timeout can be used to ensure select doesn't block forever", func(t *testing.T) {
		chan1 := make(chan string)

		got := ""

		// https://gobyexample.com/non-blocking-channel-operations

		// Basic sends and receives on channels are blocking. However, we can use select with a default clause to
		// implement non-blocking sends, receives, and even non-blocking multi-way selects. If a value is available on
		// messages then select will take the <-messages case with that value. If not it will immediately take the
		// default case.
		select {
		case str := <-chan1:
			got = str
		default:
			got = "default"
		}

		Is.Equal(got, "default")
	})
}
