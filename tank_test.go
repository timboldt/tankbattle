// Copyright (C) 2013, Tim Boldt.  All rights reserved.

package tankbattle

import (
	"math"
	"testing"
)

const (
	East        = 0.0
	West        = 180.0
	North       = 270.0
	South       = 90.0
	Triangle345 = 36.86 // A 3-4-5 triangle has this angle between the "4" and "5" sides
)

func checkLocation(t *testing.T, locationName string, actual, expected Vector2D) {
	if math.Abs(actual.X-expected.X) > 0.1 {
		t.Errorf("%v: X value was %v, expected %v", locationName, actual.X, expected.X)
	}
	if math.Abs(actual.Y-expected.Y) > 0.1 {
		t.Errorf("%v: Y value was %v, expected %v", locationName, actual.Y, expected.Y)
	}
}

func checkFloat(t *testing.T, locationName string, actual, expected float64) {
	if math.Abs(actual-expected) > 0.1 {
		t.Errorf("%v: actual %v, expected %v", locationName, actual, expected)
	}
}

func TestSimpleForwardBackward(t *testing.T) {
	tank := Tank{Vector2D{0.0, 0.0}, North, North, NotMoving, NotTurning, NotTurning}
	tank.StartDrivingForwards()
	tank.OnTimePasses(1.0)
	checkLocation(t, "After moving forwards", tank.Location(), Vector2D{0.0, -SpeedMax})
	tank.StopDriving()
	tank.OnTimePasses(1.0)
	checkLocation(t, "After being stopped", tank.Location(), Vector2D{0.0, -SpeedMax})
	tank.StartDrivingBackwards()
	tank.OnTimePasses(1.0)
	checkLocation(t, "After driving backwards", tank.Location(), Vector2D{0.0, 0.0})
}

func TestDriveAtAngle(t *testing.T) {
	tank := Tank{Vector2D{0.0, 0.0}, Triangle345, North, MovingForward, NotTurning, NotTurning}
	tank.OnTimePasses(5.0)
	checkLocation(t, "After moving forwards", tank.Location(), Vector2D{4.0 * SpeedMax, 3.0 * SpeedMax})
}

func TestDriveInASquarePattern(t *testing.T) {
	tank := Tank{Vector2D{0.0, 0.0}, East, East, NotMoving, NotTurning, NotTurning}

	tank.StartDrivingForwards()
	tank.OnTimePasses(1.0)
	tank.StopDriving()
	checkLocation(t, "After travelling east", tank.Location(), Vector2D{SpeedMax, 0.0})

	tank.StartTurningRight()
	tank.OnTimePasses(90.0 / BodyRotationRateMax)
	tank.StopTurning()

	tank.StartDrivingForwards()
	tank.OnTimePasses(1.0)
	tank.StopDriving()
	checkLocation(t, "After travelling south", tank.Location(), Vector2D{SpeedMax, SpeedMax})

	tank.StartTurningRight()
	tank.OnTimePasses(90.0 / BodyRotationRateMax)
	tank.StopTurning()

	tank.StartDrivingForwards()
	tank.OnTimePasses(1.0)
	tank.StopDriving()
	checkLocation(t, "After travelling west", tank.Location(), Vector2D{0.0, SpeedMax})

	tank.StartTurningRight()
	tank.OnTimePasses(90.0 / BodyRotationRateMax)
	tank.StopTurning()

	tank.StartDrivingForwards()
	tank.OnTimePasses(1.0)
	tank.StopDriving()
	checkLocation(t, "After travelling north", tank.Location(), Vector2D{0.0, 0.0})
}

func TestDriveInACircle(t *testing.T) {
	tank := Tank{Vector2D{1.0, 0.0}, South, South, MovingForward, TurningRight, NotTurning}
	tank.OnTimePasses(180.0 / BodyRotationRateWhileDriving)
	checkFloat(t, "Angle after travelling 180 degrees", tank.BodyAngle(), North)
	tank.OnTimePasses(180.0 / BodyRotationRateWhileDriving)
	checkLocation(t, "After travelling 360 degrees", tank.Location(), Vector2D{1.0, 0.0})
}

/*

TEST(DrivingTest, SquarePattern) { 
  Tank t(0.0, 0.0, kEast, 0.0);

  for (int i = 0; i < 4; ++i) {
    t.startDrivingForwards();
    t.onTimePasses(1.0);

    t.stopDriving();
    t.onTimePasses(1.0);
    EXPECT_NEAR(1.0, std::max(fabs(t.location().x), fabs(t.location().y)), 0.1);

    t.startRotatingRight();
    const float kBodyRotationPerTimeUnit = 90.0; //TODO
    t.onTimePasses(90.0 / kBodyRotationPerTimeUnit);
    t.stopRotating();
    EXPECT_NEAR(90.0 * i, t.bodyRotation(), 0.1);
  }
  EXPECT_PRED2(IsNearLocation, t, Vector(0.0, 0.0));
}

} // Namespace

int main(int argc, char **argv) {
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}*/
