package message

type StartGame struct {
	Time      uint
	Increment uint
}

type LoadBoard struct {
	Time      uint
	Increment uint
	GameID    string
}
