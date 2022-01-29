package methodReceivers

import (
	"github.com/matryer/is"
	"testing"
)

type account struct {
	balance int
}

// A method on a struct is called a receiver
// try to stick to 1 type

// value receiver
func (a account) depositByValueReceiver(amount int) {
	a.balance = amount
}

// pointer receiver
func (a *account) depositByPointerReceiver(amount int) {
	// Go automagically dereferences the pointer but could also be done similar to depositByPointerReceiverExplicitDereference
	a.balance = amount
}

func (a *account) depositByPointerReceiverExplicitDereference(amount int) {
	// unnecessary because Go dereferences
	(*a).balance = amount
}

func (a account) getValueReceiverAddress() *account {
	return &a
}

func (a *account) getPointerReceiverAddress() *account {
	return a
}

func Test_MethodReceivers(t *testing.T) {
	Is := is.New(t)

	// If you don’t need to edit the receiver value, use a value receiver.
	t.Run("value receiver does not change state of struct", func(t *testing.T) {
		acc := account{}

		// Value receivers are concurrency safe
		acc.depositByValueReceiver(10)

		// passed by value, remains unmodified
		Is.Equal(acc.balance, 0)
		Is.True(acc.getValueReceiverAddress() != &acc)
	})

	t.Run("pointer receiver changes state of struct", func(t *testing.T) {
		acc := account{}

		// Pointer receivers are not concurrency safe.
		acc.depositByPointerReceiver(10)

		// passed by pointer, remains unmodified
		Is.Equal(acc.balance, 10)
		Is.True(acc.getPointerReceiverAddress() == &acc)
	})

	t.Run("explicit pointer receiver changes state of struct", func(t *testing.T) {
		acc := account{}

		acc.depositByPointerReceiverExplicitDereference(10)

		// passed by pointer, remains unmodified
		Is.Equal(acc.balance, 10)
	})
}
