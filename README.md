### Sample Go Application using MongoDB and Redis

This is a sample go application using mongoDB as database and redis as cache. The CRUD API is written for both mongo and redis.

#### Requirements to run the application in local:
1. **Go** - tested on go version 1.15.3, the latest stable go version at the time of the writing. Ideally, the go version should be at least 1.11+.
2. **Mongo** - tested on version 4.0.0. Older mongo version shall work as well, though not tested.
3. **Redis** - tested on version 4.0.11. Older redis version shall work as well, though not tested.

Apart from the above 3, we can also use IDE for ease of development. GoLand from IntelliJ was used for this project. The more comprehensive list of IDE's can be found at https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins.

#### Steps to run:
1. Start mongo in local with command
```mongod```.
2. Start redis in local with command
```redis-server```.
3. Go into the project directory and run either of the following: 
```make dev``` or ```go run main/main.go```.

#### Note
API doc of all the API's exposed is present in **API.md** file in doc folder.

