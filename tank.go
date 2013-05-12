// Copyright (C) 2013, Tim Boldt.  All rights reserved.

package tankbattle

import ( "math" )

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
	location        Vector2D
	bodyAngle	float64
	turretAngle	float64

	motion MotionDirection
	bodyRotation RotationDirection
	turretRotation RotationDirection
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

func (tank *Tank) OnTimePasses(elapsedTime float64) {
	var speed float64
	switch tank.motion {
		case NotMoving:
			speed = 0
		case MovingBackward:
			speed = -SpeedMax
		case MovingForward:
			speed = SpeedMax
	}
	tank.location.X += elapsedTime * speed * math.Cos(tank.bodyAngle*math.Pi/180)
	tank.location.Y += elapsedTime * speed * math.Sin(tank.bodyAngle*math.Pi/180)
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
