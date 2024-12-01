package fsm

import (
	"errors"
	"fmt"
	"sync"
)

type FSM interface {
	SetState(id uint64, state interface{})
	GetState(id uint64) (interface{}, error)
	ClearState(id uint64) error
	saveState() error
	dropState() error
}

type BackendType int

const (
	RAM = iota
)

func Create(backend BackendType) (FSM, error) {
	switch backend {
	case RAM:
		return &FSMRAM{make(map[uint64]interface{}), sync.Mutex{}}, nil
	default:
		return nil, errors.New("Backend type: " + fmt.Sprint(backend) + " doesn't exist")
	}
}

type FSMRAM struct {
	stateMap map[uint64]interface{}
	mu       sync.Mutex
}

func (fsm *FSMRAM) SetState(id uint64, state interface{}) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	fsm.stateMap[id] = state
}

func (fsm *FSMRAM) GetState(id uint64) (interface{}, error) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	state, ok := fsm.stateMap[id]

	var err error
	if !ok {
		err = errors.New("State with id: " + fmt.Sprint(id) + " doesn't exist")
	}

	return state, err
}

func (fsm *FSMRAM) ClearState(id uint64) error {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	_, ok := fsm.stateMap[id]

	var err error
	if !ok {
		err = errors.New("State with id: " + fmt.Sprint(id) + " doesn't exist")
		return err
	}

	delete(fsm.stateMap, id)
	return nil
}

func (fsm *FSMRAM) saveState() error {

	return nil
}

func (fsm *FSMRAM) dropState() error {

	return nil
}
