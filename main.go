package main

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	keyboard := KeyboardInitCallback(window)

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()

		keyboard.PollEvents()
		if keyboard.Get(KeyEsc) == KeyDown {
			window.SetShouldClose(true)
		}
	}
}