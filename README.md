# issues-crawler
Crawler for Github issues

To start, make a copy of the file `settings.example.yml` to `settings.yml` and include your credentials.

```
cp settings.example.yml settings.yml
```


To start the local database, you must install Docker and Docker Compose. Then, you just execute it with:

```
docker-compose up -d
```

Include credentials for the root user
```
grant all on *.* to 'usernameall'@'%';
```

To start the crawler:
```
go run *.go
```