package snow_flake

import (
	"errors"
	"sync"
	"time"
)

const (
	workerIDBits = uint(10)
	sequenceBits = uint(12)

	maxWorkerID    = -1 ^ (-1 << workerIDBits)
	maxSequence    = -1 ^ (-1 << sequenceBits)
	workerIDShift  = sequenceBits
	timestampShift = sequenceBits + workerIDBits
)

type ID int64

type Snowflake struct {
	mu        sync.Mutex
	timestamp int64
	workerID  int64
	sequence  int64
}

func NewSnowflake(workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker ID excess of quantity")
	}

	return &Snowflake{workerID: workerID}, nil
}

func (s *Snowflake) Generate() ID {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / int64(time.Millisecond)

	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & maxSequence

		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now

	id := ID((now-1420041600000)<<timestampShift |
		(s.workerID << workerIDShift) |
		(s.sequence))

	return id
}
