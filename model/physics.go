package model

func getSpacialPoints(me Thing, them Thing) (SpacialPoint, SpacialPoint) {
	return me.Point.PointToSpacialPoint(), them.Point.PointToSpacialPoint()
}

//GetRotationBetween finds out which way to rotate to meet another tank
func GetRotationBetween(me Thing, them Thing) int {

	viewer, target := getSpacialPoints(me, them)
	gradient := GetGradient(viewer, target)

	//Impossible
	if gradient == 2 {
		return 0
	}

	//Todo: find quickest route between two tanks
	return DirectionClockwise
}

func IsAimingAt(viewer SpacialPoint, target SpacialPoint) bool {
	gradient := GetGradient(viewer, target)
	return Aligns(viewer, target, gradient)
}

func Aligns(viewer SpacialPoint, target SpacialPoint, gradient int) bool {

	return true
}

func GetGradient(viewer SpacialPoint, target SpacialPoint) int {
	deltaX := viewer.X - target.X
	deltaZ := viewer.Z - target.Z
	gradient := deltaZ / deltaX
	return int(gradient)
}
