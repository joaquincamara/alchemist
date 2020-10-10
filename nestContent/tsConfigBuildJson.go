package nestContent

func TsConfigBuildJson() []byte {
	tsConfigBuildJson := []byte(`
	{
		"extends": "./tsconfig.json",
		"exclude": ["node_modules", "dist", "test", "**/*spec.ts"]
	  }
	  
	`)

	return tsConfigBuildJson
}
