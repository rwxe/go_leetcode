package leetcode

import (
	"fmt"
)

type Logger struct {
	logHist  map[string]int //log历史，k:log信息，v:时间戳
	lruQueue []struct {     //lru缓存
		msg string
		ts  int
	}
	timeLimit int
	capacity  int //设定LRU容量

	maxMapLen   int //出现过的最长map
	maxSliceLen int //出现过的最长slice长度
}

// 调试用，观察字段出现的最长长度
func (l *Logger) PrintDebugMaxLen() {
	fmt.Println("map", l.maxMapLen)
	fmt.Println("slice", l.maxSliceLen)
}

// 调试用，记录字段出现的最长长度
func (l *Logger) debuglogMaxLen() {
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	l.maxMapLen = max(l.maxMapLen, len(l.logHist))
	l.maxSliceLen = max(l.maxSliceLen, len(l.lruQueue))
}

func Constructor359() Logger {
	const CAPACITY = 10
	const TIMELIMIT = 10
	return Logger{
		logHist: make(map[string]int, CAPACITY),
		lruQueue: make([]struct {
			msg string
			ts  int
		}, 0, CAPACITY),
		timeLimit: TIMELIMIT,
		capacity:  CAPACITY}
}

// 简单的实现，直接在map中记录时间戳，map会无限增长
func (l *Logger) ShouldPrintMessageDefault(ts int, msg string) bool {
	l.debuglogMaxLen()
	if lastTs, ok := l.logHist[msg]; ok && ts-lastTs < l.timeLimit {
		return false
	}
	l.logHist[msg] = ts
	return true
}

// 简单的LRU实现，但是map仍然可能突然增长
func (l *Logger) ShouldPrintMessageLRU(ts int, msg string) bool {
	l.debuglogMaxLen()
	// get
	if lastTs, ok := l.logHist[msg]; ok && ts-lastTs < l.timeLimit {
		return false
	}
	//clean
	if len(l.lruQueue) > l.capacity && ts-l.lruQueue[0].ts > l.timeLimit {
		delete(l.logHist, l.lruQueue[0].msg)
		l.lruQueue[0] = struct { //避免内存泄露
			msg string
			ts  int
		}{}
		//l.lruQueue = l.lruQueue[1:] //避免底层数组发生重新分配
		l.lruQueue = append(l.lruQueue[:0], l.lruQueue[1:]...)
	}
	//put
	newNode := struct {
		msg string
		ts  int
	}{msg, ts}
	if _, ok := l.logHist[msg]; ok {
		if len(l.lruQueue) > 1 {
			var oldIndex int
			for oldIndex = 0; l.lruQueue[oldIndex].msg == msg; oldIndex++ {
			}
			l.lruQueue = append(l.lruQueue[:oldIndex], l.lruQueue[oldIndex+1:]...)
			l.lruQueue[len(l.lruQueue)-1] = newNode
		}
	} else {
		l.lruQueue = append(l.lruQueue, newNode)
	}
	l.logHist[msg] = ts
	return true
}
