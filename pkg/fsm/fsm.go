package fsm

import (
	"errors"
	"fmt"
	"sync"
)

type FSM[StateType any] interface {
	SetState(id uint64, state StateType)
	GetState(id uint64) (StateType, error)
	ClearState(id uint64) error
	saveState() error
	dropState() error
}

type RAM struct{}

func Create[StateType any](backend interface{}) (FSM[StateType], error) {
	switch backend.(type) {
	case RAM:
		return &fsmRAM[StateType]{make(map[uint64]StateType), sync.Mutex{}}, nil
	default:
		return nil, errors.New("Backend type: " + fmt.Sprint(backend) + " doesn't exist")
	}
}

type fsmRAM[StateType any] struct {
	stateMap map[uint64]StateType
	mu       sync.Mutex
}

func (fsm *fsmRAM[StateType]) SetState(id uint64, state StateType) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	fsm.stateMap[id] = state
}

func (fsm *fsmRAM[StateType]) GetState(id uint64) (StateType, error) {
	fsm.mu.Lock()
	defer fsm.mu.Unlock()
	state, ok := fsm.stateMap[id]

	var err error
	if !ok {
		err = errors.New("State with id: " + fmt.Sprint(id) + " doesn't exist")
	}

	return state, err
}

func (fsm *fsmRAM[StateType]) ClearState(id uint64) error {
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

func (fsm *fsmRAM[StateType]) saveState() error {

	return nil
}

func (fsm *fsmRAM[StateType]) dropState() error {

	return nil
}
