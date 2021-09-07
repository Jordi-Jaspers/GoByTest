package wallet

// ========== IMPORTS ==========
import (
	"errors"
	"fmt"
)

// ========== GLOBAL VARIABLES ==========
// To check unfound errors import the following dependency
// go get -u github.com/kisielk/errcheck
var ErrInsufficientFunds = errors.New("you're not that rich, try withdrawing less")

// ========== Structures & interfaces ==========
type Stringer interface {
	String() string
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// ========== FUNCTIONS ==========
// The created object needs a reference to what object you are pointing to.
// Therefore, we use '*' it points to the reference address of the object that is saved in the memory.
// NOTE: Golang does not neet explicit derefencing to be able to use the object
// This differs the language from C or C++

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() string {
	return w.balance.String()
}
