package notes

type Command int

const (
	NONE Command = iota
	OPTIONS
	MORE
	NEXT
	PREVIOUS

	SET
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

	COMMAND_COUNT
)

//go:generate stringer -type=Command
