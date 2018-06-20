package controller

import "github.com/just1689/gg-bot-captain/model"

func autoMoveNear(tag uint) {
	//Todo: implement
}

func autoAim(targetTag uint) {
	me, target := GetMeAndTarget(GetMyTag(), targetTag)
	//Todo: implement
	r := model.GetRotationBetween(me, target)
	if r > 0 {
		//Rotate left
	} else {
		//Rotate right
	}

}

func autoAttack(tag uint) {
	//Todo: implement
}

func autoRun() {
	//Todo: implement
}
