package core

// The idea is to iterate over all different bots, analyze the input
// and if the bot think it can handle the text request in the input, run the bot.

type KumaBot interface {
	// Match can run some NLP or regex and return true if it can handle the request
	// Maybe we can return something like 0-100% here and not bool
	Match(in string) bool
	// Run the actual logic, this can be gong running, like ecternal API calls
	Run(in string) string
	// return name of this bot
	Name() string
}
