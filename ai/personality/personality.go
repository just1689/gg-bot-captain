package personality

type Personality int

const (
	Hunter Personality = iota
	Wimp
)

func randomPersonality() Personality {
	//Todo: actual random impl
	return Hunter
}
