package fsm_test

import (
	"testing"

	"github.com/avi-gecko/fsm/pkg/fsm"
)

func TestCreateFSM(t *testing.T) {
	fsm, err := fsm.Create(fsm.RAM)

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%v", fsm)
}

func TestGetState(t *testing.T) {
	fsm_test, err := fsm.Create(fsm.RAM)

	if err != nil {
		t.Error(err)
		return
	}

	fsm_test.SetState(0, "Test")
	state, err := fsm_test.GetState(0)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(state)
}

func TestGetFailedState(t *testing.T) {
	fsm_test, err := fsm.Create(fsm.RAM)

	if err != nil {
		t.Error(err)
		return
	}

	fsm_test.SetState(0, "Test")
	state, err := fsm_test.GetState(1)

	if err != nil {
		t.Log(err)
	}

	t.Log(state)
}

func TestGetFailedStateEmptyString(t *testing.T) {
	fsm_test, err := fsm.Create(fsm.RAM)

	if err != nil {
		t.Error(err)
		return
	}

	fsm_test.SetState(0, "")
	state, err := fsm_test.GetState(0)

	if err != nil {
		t.Error(err)
	}

	t.Log(state)
}

func TestClearState(t *testing.T) {
	fsm_test, err := fsm.Create(fsm.RAM)

	if err != nil {
		t.Error(err)
		return
	}

	fsm_test.SetState(0, "Test")
	err = fsm_test.ClearState(0)

	if err != nil {
		t.Error(err)
		return
	}

	state, err := fsm_test.GetState(0)

	if err != nil {
		t.Log(err)
		return
	}

	t.Log(state)
}