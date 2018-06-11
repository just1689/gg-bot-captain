package personality

type Personality int

const (
	Hunter Personality = iota
	Wimp
)

func RandomPersonality() Personality {
	//Todo: actual random impl
	return Hunter
}
