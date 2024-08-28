package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insuficcient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, initialBalance)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
