package app

type validateCmd struct {
	Source validateSourceCmd `default:"withargs" cmd:"" help:"Check a package manifest source for errors." group:"global"`
	Env    validateEnvCmd    `cmd:"" help:"Validate an environment." group:"global"`
	Script validateScriptCmd `cmd:"" help:"Validate a shell script uses only builtin or Hermit commands." group:"env"`
}
