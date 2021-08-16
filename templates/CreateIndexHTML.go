package templates

import (
	"io/ioutil"
	"log"
)

var indexHTMLTemplate = []byte(`
<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Document</title>
	<link rel="preconnect" href="https://fonts.googleapis.com">
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
	<link href="https://fonts.googleapis.com/css2?family=Roboto:wght@300;400&display=swap" rel="stylesheet">
	<link rel="preconnect" href="https://fonts.googleapis.com">
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
	<link href="https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;500&display=swap" rel="stylesheet">

	<style>
		.alchemist {
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
			height: 100vh;
			padding: 0 100px 0 100px;
		}

		.alchemist p {
			font-family: 'Roboto', sans-serif;
			text-align: justify;
			line-height: 20px;
			letter-spacing: 0.5px;
			width: 590px;
		}

		.alchemist h3 {
			font-family: 'Roboto', sans-serif;
			text-align: justify;
			line-height: 20px;
			letter-spacing: 0.5px;
			font-weight: bold;
		}

		.alchemist a {
			font-family: 'Roboto', sans-serif;
			text-align: justify;
			line-height: 20px;
			letter-spacing: 0.5px;
		}
	</style>

</head>

<body>
	<div class="alchemist">
		<img src="https://user-images.githubusercontent.com/26718123/94979651-96461b00-04e9-11eb-94d6-660af9663975.png"
			alt="alchemist_cli_logo">
		<h3>Hello, fellow alchemists of the OSS world!</h3>
		<p>I'm Joaquin Camara. Thank you so much for give a try to Alchemist Cli, hope you find this tool usefull!</p>

		<h3>A little about me:</h3>
		<p>I am a software engineer who has been working with a variety of applications. Specialized in Go, JavaScript,
			and some front-end frameworks and libraries such as React.js, Redux.js, Meteor.js, Node.js, express.js,
			media-soup, React-Native, and other technologies like CSS, HTML, and HTTPS</p>
		<p>I am also an adventurer at heart and software development has given me the freedom to live great adventures,
			I have been a pirate (like Jack Sparrow), a jungle scout, and a brewmaster, and now with the help of
			open-source, I will be able to contribute back to the world of software development.</p>

		<h3>I will be eternally grateful with your help to keep alchemist updated</h3>
		<a href="https://github.com/joaquincamara">Check my github sponsor account</a>

		<h3>What your supports mean to me:</h3>
		<p>I love OSS and develop software!</p>
		<p>The Alchemist project was born as a simple hobby for the pandemic age, with the pass of time and a lot of
			feedback from some of my mentors, Alchemist took the form of an awesome idea to help developers create
			amazing software.</p>
		<p>Now I have so many ideas for the "Alchemist" project that I either want to finish or build, and your support
			will help me dedicate the time to do that work.</p>
	</div>
</body>

</html>
`)

func CreateIndexHTML() {
	err := ioutil.WriteFile("public/index.html", indexHTMLTemplate, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
