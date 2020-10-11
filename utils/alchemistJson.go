package utils

func AlchemistJson(appName string, appType string) []byte {
	alchemistJson := []byte(`
	{
		"name": ` + `"` + appName + `"` + `,
		"version": "0.1.0",
         "appType": ` + `"` + appType + `"` + `
	  }
	`)

	return alchemistJson
}
