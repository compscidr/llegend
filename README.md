# llegend
WiP, but goal is:
  - shows segments you are local legend of on a map
  - color coding of the segment showing how close to being overtaken
  - color red = someone is nearly passed efforts
  - color green = many efforts before someone catches you
  - color gray = lost LL
  - later: show segments you are getting close to being LL for with similar color coding

## building and running
There are two necessary environment variables to run. The first is the access
token for the strava API, and the second is the athlete ID.

These can be set via `.env` file for docker and docker-compose, otherwise
they can be passed in via the command to run the program:

### locally
```
go build
ACCESS_TOKEN=<insert> ATHLETE=<insert> ./llegend 
```

### docker
```
docker build -f Dockerfile . -t llegend
docker run -p 8080:8080 --env-file .env llegend
```

### docker-compose
`docker-compose up --build`

Then browse to http://localhost:8080

## inspiration projects
http://www.raceleap.live/
https://stravanity.vercel.app/
