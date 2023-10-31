package account_test

import (
	"github.com/google/uuid"
	"github.com/hojin-kr/account"
	"testing"
)

func TestAccount(t *testing.T) {
	a := account.NewAccount("first_app", "hojin")
	if a == nil {
		t.Error("NewAccount() error")
	}
	// is valid uuid
	_, err := uuid.Parse(a.UUID)
	if err != nil {
		t.Error("NewAccount() error")
	}
}
