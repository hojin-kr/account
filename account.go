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

func (account *Account) NewAccount() {
	account.UUID = uuid.New().String()
	// put datastore
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	datastore.Put(account.APPID+":"+account.Token, account.UUID)
}

func (account *Account) GetAccount() {
	// get datastore
	datastore := datastore.GcpDatastore{}
	datastore.Init()
	account.UUID = datastore.Get(account.APPID + ":" + account.Token)
}
