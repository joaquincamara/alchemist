package nestContent

func Prettierrc() []byte {
	prettierrc := []byte(`
	{
		"singleQuote": true,
		"trailingComma": "all"
	  }
	`)

	return prettierrc
}
