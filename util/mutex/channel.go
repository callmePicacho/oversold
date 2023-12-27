package mutex

var OversoldChan *Mutex

func init() {
	OversoldChan = NewMutex()
}

type Mutex struct {
	ch chan struct{}
}

// init clock
func NewMutex() *Mutex {
	mutex := &Mutex{
		ch: make(chan struct{}, 1),
	}
	return mutex
}

// get lock
func (m *Mutex) Lock() {
	m.ch <- struct{}{}
}

// return lock
func (m *Mutex) Unlock() {
	select {
	case <-m.ch:
	default:
		panic("unlock the unlocked mutex")
	}
}
