package github.com/jroden2/dumb

func ToReverseCamelCase(input string) string {
	uppercase := false
	output := ""

	for _, x := range input {
		x := string(x)
		if uppercase {
			x = strings.ToUpper(x)
			uppercase = false
		} else {
			x = strings.ToLower(x)
			uppercase = true
		}

		output = output + x
	}

	return output
}

func ToCamelCase(input string) string {
	uppercase := true
	output := ""

	for _, x := range input {
		x := string(x)
		if !uppercase {
			x = strings.ToLower(x)
			uppercase = true
		} else {
			x = strings.ToUpper(x)
			uppercase = false
		}

		output = output + x
	}

	return output
}
