package goroutine

type GO struct {
	active      chan bool
	recoverFunc func(interface{})
}

// NewGO
// maxNum: 允许的最大线程数
func NewGO(maxNum int) *GO {
	return &GO{active: make(chan bool, maxNum)}
}

func (t *GO) RecoverFunc(f func(interface{})) *GO {
	t.recoverFunc = f
	return t
}

func (t *GO) Run(f func()) {
	t.active <- true
	go func() {
		defer func() {
			<-t.active
			if t.recoverFunc == nil {
				return
			}
			t.recoverFunc(recover())
		}()
		f()
	}()
}
