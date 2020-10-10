package nestContent

func JestE2e() []byte {
	jestE2e := []byte(`
	{
		"moduleFileExtensions": ["js", "json", "ts"],
		"rootDir": ".",
		"testEnvironment": "node",
		"testRegex": ".e2e-spec.ts$",
		"transform": {
		  "^.+\\.(t|j)s$": "ts-jest"
		}
	  }
	  
	
	`)

	return jestE2e
}
