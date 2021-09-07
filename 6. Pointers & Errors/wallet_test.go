package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Depositing bitcoins in a wallet is possible.", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10).String()

		AssertBalance(t, wallet, want)
	})
	t.Run("Depositing bitcoins in a wallet is possible.", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(5).String()

		AssertBalance(t, wallet, want)
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		AssertBalance(t, wallet, startingBalance.String())
		AssertError(t, err, ErrInsufficientFunds)
	})
	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		AssertNoError(t, err)
		AssertBalance(t, wallet, Bitcoin(10).String())
	})
}

func AssertError(t testing.TB, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if err.Error() != want.Error() {
		t.Errorf("got %q, want %q \n", err, want)
	}
}

func AssertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func AssertBalance(t testing.TB, got Wallet, want string) {
	t.Helper()
	balance := got.Balance()
	if balance != want {
		t.Errorf("The balance in the account is %s and not %s. \n", balance, want)
	}
}
