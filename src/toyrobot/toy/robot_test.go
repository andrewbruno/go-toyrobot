package toy_test

import (
	"testing"
	"toyrobot/compass"
	"toyrobot/toy"
	"toyrobot/unit"
)

func TestNewRobot(t *testing.T) {
	var robot toy.Toyer
	robot = toy.NewRobot()
	if robot == nil {
		t.Errorf("Could not instantiate a new table object")
	}
}

func TestUpdate(t *testing.T) {
	robot := toy.NewRobot()
	robot.Update(unit.X(2), unit.Y(3), compass.SOUTH)
	if robot.X != unit.X(2) || robot.Y != unit.Y(3) || robot.F != compass.SOUTH {
		t.Errorf("Update failed, expected %s,%s,%s but got %s,%s,%s", unit.X(2), unit.Y(2), compass.SOUTH, robot.X, robot.Y, robot.F)
	}
}
