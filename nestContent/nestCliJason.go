package nestContent

func NestCliJson() []byte {
	nestCliJson := []byte(`
	{
		"collection": "@nestjs/schematics",
		"sourceRoot": "src"
	  }
	  
	`)

	return nestCliJson
}
