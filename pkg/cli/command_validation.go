package cli

type commandStringRequirement uint8

const (
	commandStringOptional commandStringRequirement = iota
	commandStringRequired
	commandStringRequiredNonEmpty
)

type commandRequirements struct {
	operation         string
	inFile            commandStringRequirement
	outFile           commandStringRequirement
	outDir            commandStringRequirement
	inFileJSON        commandStringRequirement
	outFileJSON       commandStringRequirement
	minInputFiles     int
	maxInputFiles     int
	missingInputFiles error
	minIntValues      int
	missingIntValues  error
}

func commandValidationError(operation string, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCommandString(value *string, requirement commandStringRequirement, missing error) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCommandInputFiles(inFiles []string, minCount, maxCount int, missing error) error {
	_ = "STUB: not implemented"
	return nil
}

func validateCommandRequirements(cmd *Command, requirements commandRequirements) error {
	_ = "STUB: not implemented"
	return nil
}

func validatedCommandInFile(cmd *Command, operation string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func optionalCommandString(value *string) string { _ = "STUB: not implemented"; return "" }
