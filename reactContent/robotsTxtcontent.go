package reactContent

func RobotsTxtContentGenerator() []byte {
	robotsTxtContent := []byte(`# https://www.robotstxt.org/robotstxt.html
	User-agent: *
	Disallow:
	
	`)

	return robotsTxtContent
}
