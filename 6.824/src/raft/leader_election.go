package raft

import (
	"math/rand"
	"sync"
	"time"
)

func (rf *Raft) resetElectionTimer()  {
	t := time.Now()
	electionTimeout := time.Duration(150 + rand.Intn(150)) * time.Millisecond
	rf.electionTime = t.Add(electionTimeout)
}

func (l *Log) append(entries ...Entry)  {
	l.Entries = append(l.Entries, entries...)
}

func (rf *Raft) setNewTerm(term int) {
	if term > rf.currentTerm || rf.currentTerm == 0 {
		rf.state = Follower
		rf.currentTerm = term
		rf.voteFor = -1
		DPrintf("[%d]: set term %v\n", rf.me, rf.currentTerm)
		rf.persist()
	}
}

func (rf *Raft) leaderElection() {
	rf.currentTerm++
	rf.state = Candidate
	rf.voteFor = rf.me
	rf.persist()
	rf.resetElectionTimer()
	term := rf.currentTerm
	voteCounter := 1
	lastLog := rf.log.lastLog()
	DPrintf("[%v]: start leader election, term %d\n", rf.me, rf.currentTerm)
	args := RequestVoteArgs{
		Term: term,
		CandidateId: rf.me,
		LastLogIndex: lastLog.Index,
		LastLogTerm: lastLog.Term,
	}
	var becomeLeader sync.Once
	for serverId, _ := range rf.peers {
		if serverId != rf.me {
			go rf.candidateRequestVote(serverId, &args, &voteCounter, &becomeLeader)
		}
	}

}
