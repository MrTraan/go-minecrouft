package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type KeyState int

const (
	KeyUp      KeyState = 0
	KeyDown    KeyState = 1
	KeyPressed KeyState = 2
)

type KeyCode int

const (
	KeyEsc KeyCode = iota
	KeyW
	KeyA
	KeyS
	KeyD
)

const numKeys int = 5

type KeyEvent struct {
	State KeyState
	Code  KeyCode
}

type Keyboard struct {
	states []KeyState
}

var keyCallbackChan chan KeyEvent

func KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	var event KeyEvent

	if action == glfw.Press {
		event.State = KeyPressed
	} else if action == glfw.Release {
		event.State = KeyUp
	}

	switch key {
	case glfw.KeyW:
		event.Code = KeyW
		break
	case glfw.KeyA:
		event.Code = KeyA
		break
	case glfw.KeyS:
		event.Code = KeyS
		break
	case glfw.KeyD:
		event.Code = KeyD
		break
	case glfw.KeyEscape:
		event.Code = KeyEsc
	default:
		return
	}
	keyCallbackChan <- event
}

func KeyboardInitCallback(window *glfw.Window) *Keyboard {
	k := new(Keyboard)
	k.states = make([]KeyState, numKeys)
	for i := range k.states {
		k.states[i] = KeyUp
	}

	keyCallbackChan = make(chan KeyEvent, 10)
	window.SetKeyCallback(KeyCallback)

	return k
}

func (keyboard *Keyboard) PollEvents() {
	for i := range keyboard.states {
		if keyboard.states[i] == KeyPressed {
			keyboard.states[i] = KeyDown
		}
	}

	for {
		select {
		case e := <-keyCallbackChan:
			keyboard.states[e.Code] = e.State
		default:
			return
		}
	}
}

func (keyboard *Keyboard) Get(key KeyCode) KeyState {
	return keyboard.states[key]
}
