package luvcheckerlib

import "sync"

// All global variables are defined here and can be accessed from any package.
var (
	AccountRegex = "^[a-zA-Z0-9]{3,20}:[a-zA-Z0-9]{3,20}$"
	ProxyRegex   = `((http|https|socks4|socks5):\/\/)?([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?`
	HitsNumber   = 0
	FailsNumber  = 0
	CPM          = 0
	Retries      = 0
	AccountList  []Account
	HitList      []*Account
	ProxyList    []*Proxy
	OnCheck      = make(chan int)
	BotCount     = 50
	WaitGroup    = sync.WaitGroup{}
	Semaphore    chan int
)
