package store

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type IntentStatus int

const (
	IntentStatusCreated IntentStatus = iota
	IntentStatusFulfilled
	IntentStatusSettled
	IntentStatusCancelled
)

type Intent struct {
	ID              common.Hash
	Maker           common.Address
	InputToken      common.Address
	InputAmount     *big.Int
	OutputToken     common.Address
	Solver          common.Address
	MinOutputAmount *big.Int
	TargetChainId   uint32
	Recipient       common.Address
	Deadline        uint64
	Nonce           uint64
	Status          IntentStatus
}

type IntentStore struct {
	mu      sync.RWMutex
	intents map[common.Hash]Intent
}

func (s *IntentStore) Save(intent Intent) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.intents[intent.ID] = intent
}

func (s *IntentStore) Get(id common.Hash) (Intent, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	intent, ok := s.intents[id]
	return intent, ok
}

func (s *IntentStore) GetAll() []Intent {
	s.mu.RLock()
	defer s.mu.RUnlock()
	intents := make([]Intent, 0, len(s.intents))
	for _, intent := range s.intents {
		intents = append(intents, intent)
	}
	return intents
}

func NewIntentStore() *IntentStore {
	return &IntentStore{
		intents: make(map[common.Hash]Intent),
	}
}
