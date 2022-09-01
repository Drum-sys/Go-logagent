package raft

type Entry struct {
	Command interface{}
	Term    int
	Index   int
}

type Log struct {
	Entries []Entry
	Index0  int
}

func makeEmptyLog() Log {
	log := Log{
		Entries: make([]Entry, 0),
		Index0:  0,
	}
	return log
}

func (l *Log) len() int {
	return len(l.Entries)
}

func (l *Log) at(idx int) *Entry {
	return &l.Entries[idx]
}

func (l *Log) lastLog() *Entry {
	return l.at(l.len() - 1)
}

func (l *Log) slice(idx int) []Entry {
	return l.Entries[idx:]
}

func (l *Log) truncate(idx int)  {
	l.Entries = l.Entries[:idx]
}