package nestContent

func AppService() []byte {
	appService := []byte(`
	import { Injectable } from '@nestjs/common';

	@Injectable()
	export class AppService {
	  getHello(): string {
		return 'Hello World!';
	  }
	}
	
	`)

	return appService
}
