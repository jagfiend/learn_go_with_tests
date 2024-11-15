package money_problems

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got.Error() != want.Error() {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("check wallet for cash", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw money from wallet", func(t *testing.T) {
		wallet := Wallet{Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(50))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(50))
	})
	t.Run("cannot withdraw more than balance", func(t *testing.T) {
		wallet := Wallet{Bitcoin(50)}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrNotEnoughJuice)
		assertBalance(t, wallet, Bitcoin(50))
	})
}
