// Copyright (C) 2013, Tim Boldt.  All rights reserved.

package tankbattle

import (
	"math"
)

//
// Stuff that could be generalized
//
type Vector2D struct {
	X, Y float64
}

type MotionDirection byte

const (
	NotMoving MotionDirection = iota
	MovingForward
	MovingBackward
)

type RotationDirection byte

const (
	NotTurning RotationDirection = iota
	TurningLeft
	TurningRight
)

//
// Tank-specific stuff
//

const (
	SpeedMax           = 100.0
	SpeedWhileRotating = 80.0
)

const (
	// Degrees per second
	BodyRotationRateMax          = 45.0
	BodyRotationRateWhileDriving = 22.5
	TurretRotationRateMax        = 90.0
)

type Tank struct {
	location    Vector2D
	bodyAngle   float64
	turretAngle float64

	motion        MotionDirection
	bodyTurning   RotationDirection
	turretTurning RotationDirection
}

func (tank *Tank) StartDrivingForwards() {
	tank.motion = MovingForward
}

func (tank *Tank) StartDrivingBackwards() {
	tank.motion = MovingBackward
}

func (tank *Tank) StopDriving() {
	tank.motion = NotMoving
}

func (tank *Tank) StartTurningLeft() {
	tank.bodyTurning = TurningLeft
}

func (tank *Tank) StartTurningRight() {
	tank.bodyTurning = TurningRight
}

func (tank *Tank) StopTurning() {
	tank.bodyTurning = NotTurning
}

func (tank *Tank) StartTurningTurretLeft() {
	tank.turretTurning = TurningLeft
}

func (tank *Tank) StartTurningTurretRight() {
	tank.turretTurning = TurningRight
}

func (tank *Tank) StopTurretTurning() {
	tank.turretTurning = NotTurning
}

func (tank *Tank) Speed() float64 {
	var speed float64

	switch tank.motion {
	case NotMoving:
		speed = 0
	case MovingBackward:
		if tank.bodyTurning == NotTurning {
			speed = -SpeedMax
		} else {
			speed = -SpeedWhileRotating
		}
	case MovingForward:
		if tank.bodyTurning == NotTurning {
			speed = SpeedMax
		} else {
			speed = SpeedWhileRotating
		}
	}
	return speed
}

func (tank *Tank) BodyTurnRate() float64 {
	var turn float64

	switch tank.bodyTurning {
	case NotTurning:
		turn = 0
	case TurningLeft:
		if tank.motion == NotMoving {
			turn = -BodyRotationRateMax
		} else {
			turn = -BodyRotationRateWhileDriving
		}
	case TurningRight:
		if tank.motion == NotMoving {
			turn = BodyRotationRateMax
		} else {
			turn = BodyRotationRateWhileDriving
		}
	}
	return turn
}

func (tank *Tank) TurretTurnRate() float64 {
	var turretTurn float64

	switch tank.turretTurning {
	case NotTurning:
		turretTurn = 0
	case TurningLeft:
		turretTurn = -TurretRotationRateMax
	case TurningRight:
		turretTurn = TurretRotationRateMax
	}
	return turretTurn
}

func (tank *Tank) OnTimePasses(elapsedTime float64) {
	for elapsedTime > 0.0 {
		time := math.Min(0.01, elapsedTime)
		elapsedTime -= time

		tank.location.X += time * tank.Speed() * math.Cos(tank.bodyAngle*math.Pi/180)
		tank.location.Y += time * tank.Speed() * math.Sin(tank.bodyAngle*math.Pi/180)

		tank.bodyAngle += time * tank.BodyTurnRate()
		tank.turretAngle += time * tank.TurretTurnRate()
	}
}

func (tank *Tank) Location() Vector2D {
	return tank.location
}

func (tank *Tank) BodyAngle() float64 {
	return tank.bodyAngle
}

func (tank *Tank) TurretAngle() float64 {
	return tank.turretAngle
}
