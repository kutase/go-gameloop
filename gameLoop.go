package gameLoop

import (
	"runtime"
	"time"
)

type GameLoop struct {
	onUpdate  func(float64)
	tickRate  time.Duration
	canUpdate bool
}

// Create new game loop
func New(tickRate time.Duration, onUpdate func(float64)) *GameLoop {
	return &GameLoop{
		onUpdate:  onUpdate,
		tickRate:  tickRate,
		canUpdate: false,
	}
}

func (gl *GameLoop) startLoop() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	tickInterval := time.Second / gl.tickRate
	timeStart := time.Now().UnixNano()

	tick := time.Tick(tickInterval)

	for {
		if !gl.canUpdate {
			break
		}

		select {
		case <-tick:
			now := time.Now().UnixNano()
			// DT in seconds
			delta := float64(now-timeStart) / 1000000000
			timeStart = now
			gl.onUpdate(delta)
		}
	}
}

func (gl *GameLoop) GetTickRate () time.Duration {
	return gl.tickRate
}

// Set tickRate and restart game loop
func (gl *GameLoop) SetTickRate (tickRate time.Duration) {
	gl.tickRate = tickRate
	gl.Restart()
}

// Start game loop
func (gl *GameLoop) Start() {
	gl.canUpdate = true
	go gl.startLoop()
}

// Stop game loop
func (gl *GameLoop) Stop() {
	gl.canUpdate = false
}

// Restart game loop
func (gl *GameLoop) Restart () {
	gl.Stop()
	gl.Start()
}
