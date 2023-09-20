package leetcode

type Logger struct {
	LogHist  map[string]int //log历史，k:log信息，v:时间戳
	LruQueue []struct {     //lru缓存
		msg string
		ts  int
	}
	timeLimit int
	capacity  int //设定LRU容量
}

func Constructor359() Logger {
	const CAPACITY = 10
	const TIMELIMIT = 10
	return Logger{
		LogHist: make(map[string]int, CAPACITY),
		LruQueue: make([]struct {
			msg string
			ts  int
		}, 0, CAPACITY),
		timeLimit: TIMELIMIT,
		capacity:  CAPACITY}
}

// 简单的实现，直接在map中记录时间戳，map会无限增长
func (l *Logger) ShouldPrintMessageDefault(ts int, msg string) bool {
	if lastTs, ok := l.LogHist[msg]; ok && ts-lastTs < l.timeLimit {
		return false
	}
	l.LogHist[msg] = ts
	return true
}

// 简单的LRU实现，但是map仍然可能突然增长
func (l *Logger) ShouldPrintMessageLRU(ts int, msg string) bool {
	// get
	if lastTs, ok := l.LogHist[msg]; ok && ts-lastTs < l.timeLimit {
		return false
	}
	//clean
	if len(l.LruQueue) > l.capacity && ts-l.LruQueue[0].ts > l.timeLimit {
		delete(l.LogHist, l.LruQueue[0].msg)
		//防止内存泄露
		l.LruQueue[0] = struct {
			msg string
			ts  int
		}{msg, ts}
		l.LruQueue = l.LruQueue[1:]
	}
	//put
	newNode := struct {
		msg string
		ts  int
	}{msg, ts}
	if _, ok := l.LogHist[msg]; ok && len(l.LruQueue) > 1 {
		var oldIndex int
		//不需要防止越界，因为LruQueue理应存在对应数据
		for oldIndex = 0; l.LruQueue[oldIndex].msg != msg; oldIndex++ {
		}
		l.LruQueue = append(l.LruQueue[:oldIndex], l.LruQueue[oldIndex+1:]...)
		l.LruQueue = append(l.LruQueue, newNode)
	} else {
		l.LruQueue = append(l.LruQueue, newNode)
	}
	l.LogHist[msg] = ts
	return true
}
