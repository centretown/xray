package notes

type COMMAND int

const (
	NONE COMMAND = iota
	OPTIONS
	MORE
	NEXT
	PREVIOUS
	INCREMENT
	DECREMENT
	INCREMENT_MORE
	DECREMENT_MORE
	PAUSE_PLAY
	SHARE

	ACTION
	BACK
	CANCEL
	OUT

	COMMANDS
)