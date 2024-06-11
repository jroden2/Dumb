package github.com/jroden2/dumb

func ToCamelCase(input string) string {
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
