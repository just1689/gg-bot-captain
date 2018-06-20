package model

type Personality int

const (
	Hunter Personality = iota
	Wimp
)

func RandomPersonality() Personality {
	//Todo: actual random impl
	return Hunter
}

func (p Personality) IsStillRelavant() bool {
	//Todo: decide how to choose whether or not a personality is relevant
	return true
}
