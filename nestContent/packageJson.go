package nestContent

func PackageJson() []byte {
	packageJson := []byte(`
	{
		"name": "nest-typescript-starter",
		"private": true,
		"version": "1.0.0",
		"description": "Nest TypeScript starter repository",
		"license": "MIT",
		"scripts": {
		  "prebuild": "rimraf dist",
		  "build": "nest build",
		  "format": "prettier --write \"src/**/*.ts\" \"test/**/*.ts\"",
		  "start": "nest start",
		  "start:dev": "nest start --watch",
		  "start:debug": "nest start --debug --watch",
		  "start:prod": "node dist/main",
		  "lint": "eslint \"{src,apps,libs,test}/**/*.ts\" --fix",
		  "test": "jest",
		  "test:watch": "jest --watch",
		  "test:cov": "jest --coverage",
		  "test:debug": "node --inspect-brk -r tsconfig-paths/register -r ts-node/register node_modules/.bin/jest --runInBand",
		  "test:e2e": "jest --config ./test/jest-e2e.json"
		},
		"dependencies": {
		  "@nestjs/common": "^7.1.1",
		  "@nestjs/core": "^7.1.1",
		  "@nestjs/platform-express": "^7.1.1",
		  "reflect-metadata": "^0.1.13",
		  "rimraf": "^3.0.2",
		  "rxjs": "^6.5.5"
		},
		"devDependencies": {
		  "@nestjs/cli": "^7.2.0",
		  "@nestjs/schematics": "^7.0.0",
		  "@nestjs/testing": "^7.1.1",
		  "@types/express": "^4.17.6",
		  "@types/jest": "^25.2.3",
		  "@types/node": "^14.0.6",
		  "@types/supertest": "^2.0.9",
		  "@typescript-eslint/eslint-plugin": "^3.0.2",
		  "@typescript-eslint/parser": "^3.0.2",
		  "eslint": "^7.1.0",
		  "eslint-config-prettier": "^6.11.0",
		  "eslint-plugin-import": "^2.20.2",
		  "jest": "^26.0.1",
		  "prettier": "^2.0.5",
		  "supertest": "^4.0.2",
		  "ts-jest": "^26.1.0",
		  "ts-loader": "^7.0.5",
		  "ts-node": "^8.10.2",
		  "tsconfig-paths": "^3.9.0",
		  "typescript": "^3.9.3"
		},
		"jest": {
		  "moduleFileExtensions": [
			"js",
			"json",
			"ts"
		  ],
		  "rootDir": "src",
		  "testRegex": ".spec.ts$",
		  "transform": {
			"^.+\\.(t|j)s$": "ts-jest"
		  },
		  "collectCoverageFrom": [
			"**/*.(t|j)s"
		  ],
		  "coverageDirectory": "../coverage",
		  "testEnvironment": "node"
		}
	  }
	  
	`)

	return packageJson

}
