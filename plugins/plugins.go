package plugin

type Plugin interface {
	Validate(msg string) bool
}
