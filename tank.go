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

func (tank *Tank) OnTimePasses(elapsedTime float64) {
	for elapsedTime > 0.0 {
		time := math.Min(0.1, elapsedTime)
		elapsedTime -= time

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
		tank.location.X += time * speed * math.Cos(tank.bodyAngle*math.Pi/180)
		tank.location.Y += time * speed * math.Sin(tank.bodyAngle*math.Pi/180)

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
		tank.bodyAngle += time * turn
	}
}

/*
  void startRotatingLeft();
  void startRotatingRight();
  void stopRotating();

  void startRotatingTurretLeft();
  void startRotatingTurretRight();
  void stopRotatingTurret();
*/
func (tank *Tank) Location() Vector2D {
	return tank.location
}

func (tank *Tank) BodyAngle() float64 {
	return tank.bodyAngle
}

func (tank *Tank) TurretAngle() float64 {
	return tank.turretAngle
}

/*
  sf::Vector2<float> location() const { return bodyTransform_.getPosition(); }
  float bodyRotation() const { return bodyTransform_.getRotation(); }
  float turretRotation() const { return turretTransform_.getRotation(); }

  float speed() const;
  float bodyRotationRate();
  float turretRotationRate();

  void onTimePasses(float elapsedTime);
  void onDraw();

 private:
  MotionDirection motion_direction_;
  RotationDirection body_rotation_direction_;
  RotationDirection turret_rotation_direction_;

  sf::Transformable bodyTransform_;
  sf::Transformable turretTransform_;
};

::std::ostream& operator<<(::std::ostream& os, const Tank& t);

} // namespace tankbattle

namespace sf {
::std::ostream& operator<<(::std::ostream& os, const Vector2<float>& v);
} // namespace "sf" 
#endif
*/
