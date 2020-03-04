package main

import (
	"errors"
	"fmt"
	"testing"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var InsufficientFundsError = errors.New("cannot withdraw, insuffic funds")

func (w *Wallet) Withdraw(ammount Bitcoin) error {
	if ammount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= ammount
	return nil
}

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Errorf("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

	}

	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Fatal("got an error but didnot want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("withdraw infufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})
}

/*
nil
	指针可以是 nil
	但函数返回一个指针，你需要确保检查过它是否是nil，否则可能会抛出一个执行异常（运行时）
	nil 非常适合描述一个可能丢失的值
错误
	调用函数时表示失败
	在错误中检查字符串会导致测试不稳定，因此我们用有意义的值重构了，这样报错会显示出来，不然报错不会显示出来 和 react 那个有点像
*/
