// Copyright (C) 2013, Tim Boldt.  All rights reserved.

package tankbattle

import (
	"math"
	"testing"
)

const (
	East  = 0.0
	West  = 180.0
	North = -90.0
	South = 90.0
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
	checkLocation(t, "After moving forwards", tank.Location(), Vector2D{4.0*SpeedMax, 3.0*SpeedMax})
}

/*
TEST(DrivingTest, StartStop) {
  Tank t(0.0, 0.0, kEast, 0.0);
  EXPECT_PRED2(HasSpeed, t, 0.0);

  t.startDrivingForwards();
  EXPECT_PRED2(HasSpeed, t, kSpeedMax);
  t.onTimePasses(1.0);
  EXPECT_PRED2(HasSpeed, t, kSpeedMax);

  t.stopDriving();
  EXPECT_PRED2(HasSpeed, t, 0.0);
  t.onTimePasses(1.0);
  EXPECT_PRED2(HasSpeed, t, 0.0);

  t.startDrivingBackwards();
  EXPECT_PRED2(HasSpeed, t, -1.0 * kSpeedMax);
  t.onTimePasses(1.0);
  EXPECT_PRED2(HasSpeed, t, -1.0 * kSpeedMax);
}

TEST(DrivingTest, DriveForward) { 
  {
    Tank t(0.0, 0.0, kEast, 0.0);

    t.startDrivingForwards();
    t.onTimePasses(1.0);
    EXPECT_PRED2(IsNearLocation, t, Vector(kSpeedMax, 0));
  }

  {
    Tank t(0.0, 0.0, kWest, 0.0);

    t.startDrivingForwards();
    t.onTimePasses(1.0);
    EXPECT_PRED2(IsNearLocation, t, Vector(-1.0 * kSpeedMax, 0));
  }

  {
    Tank t(0.0, 0.0, kNorth, 0.0);

    t.startDrivingForwards();
    t.onTimePasses(1.0);
    EXPECT_PRED2(IsNearLocation, t, Vector(0, -1.0 * kSpeedMax));
  }

  {
    Tank t(0.0, 0.0, kSouth, 0.0);

    t.startDrivingForwards();
    t.onTimePasses(1.0);
    EXPECT_PRED2(IsNearLocation, t, Vector(0, kSpeedMax));
  }

  {
    // 3-4-5 triangle = 36.86 degrees
    Tank t(0.0, 0.0, 36.86, 0.0);

    t.startDrivingForwards();
    t.onTimePasses(5.0);
    EXPECT_PRED2(IsNearLocation, t, Vector(4.0 * kSpeedMax, 3.0 * kSpeedMax));
  }
}

TEST(DrivingTest, SimpleForwardBackwards) { 
  Tank t(0.0, 0.0, kSouth, 0.0);

  t.startDrivingForwards();
  t.onTimePasses(1.0);
  EXPECT_PRED2(IsNearLocation, t, Vector(0.0, 1.0 * kSpeedMax));

  t.stopDriving();
  t.onTimePasses(1.0);
  EXPECT_PRED2(IsNearLocation, t, Vector(0.0, 1.0 * kSpeedMax));

  // Assumes tank speed is the same forward and backwards
  t.startDrivingBackwards();
  t.onTimePasses(1.0);
  EXPECT_PRED2(IsNearLocation, t, Vector(0.0, 0.0));
}

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
