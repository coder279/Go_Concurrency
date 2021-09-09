// 伪代码 无法直接运行
// 1.去掉state标志
// 2.如果检测到没有加锁，直接去掉锁标志直接panic
// 3.第一种情况:goroutine新来的goroutine无需唤醒直接return
// 4.第二种情况:唤醒旧goroutine，释放锁

func (m *Mutex) Unlock() {
// Fast path: drop lock bit.
	new := atomic.AddInt32(&m.state, -mutexLocked) //去掉锁标志
	if (new+mutexLocked)&mutexLocked == 0 { //本来就没有加锁
		panic("sync: unlock of unlocked mutex")
	}
	// 唤醒休眠的goroutine
	old := new
	for {
		//这里应该是新goroutine,无需唤醒
		if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken) != 0 { // 没有等待者，或者有唤醒的waiter，或者锁原来已加锁
		return
	}
	//唤醒goroutine
	new = (old - 1<<mutexWaiterShift) | mutexWoken // 新状态，准备唤醒goroutine，并设置唤醒标志
	if atomic.CompareAndSwapInt32(&m.state, old, new) {
		runtime.Semrelease(&m.sema)
		return
	}
	old = m.state
	}
}
