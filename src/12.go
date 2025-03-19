func getRandomNumber() int {
	return rand.Intn(10) + 1
}

func solveProblem(problem string) (string, error) {
	// Check if the problem is a valid math expression
	if !mathExpressionRegex.MatchString(problem) {
		return "", errors.New("invalid math expression")
	}

	// Solve the problem using the eval function from the math/eval package
	result, err := eval.Eval(problem)
	if err != nil {
		return "", err
	}

	// Return the result as a string
	return fmt.Sprintf("%f", result), nil
}
