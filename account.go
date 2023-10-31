package account

import (
	"github.com/google/uuid"
	"github.com/hojin-kr/datastore"
)

type Account struct {
	UUID  string `json:"uuid"`
	APPID string `json:"appid"`
	Token string `json:"token"`
}

func NewAccount(APPID string, Token string) *Account {
	account := Account{}
	account.UUID = uuid.New().String()
	account.APPID = APPID
	account.Token = Token
	// put datastore
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	datastore.Put(APPID+":"+account.Token, account.UUID)
	return &account
}
