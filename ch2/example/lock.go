// 伪代码，并非可以直接运行的代码

// 1.如果CAS检测state状态幸运直接加锁
// 2.获取不到锁，进入休眠状态
// 3.锁释放了，goroutine被唤醒，清除mutexWoken标志，休眠,加入等待队列
// 3.新来的goroutine，直接加入等待队列,休眠
// 4.旧的goroutine，清除mutexWoken，获取到锁
// 5.新的goroutine直接获取到锁

func (m *Mutex) Lock() {
	// Fast path: 幸运case，能够直接获取到锁
	// CAS检测state字段
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}
	//CAS有点像自旋锁
	awoke := false
	for {
		old := m.state
		new := old | mutexLocked // 新状态加锁
		if old&mutexLocked != 0 {
			new = old + 1<<mutexWaiterShift //等待者数量加一
		}
		if awoke {
			// goroutine是被唤醒的，
			// 新状态清除唤醒标志
			new &^= mutexWoken
		}
		if atomic.CompareAndSwapInt32(&m.state, old, new) {//设置新状态
			if old&mutexLocked == 0 { // 锁原状态未加锁
				break
			}
			runtime.Semacquire(&m.sema) // 请求信号量
			awoke = true
		}
	}
}