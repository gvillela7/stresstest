# stresstest
Desafio  fullcycle stress test

## Build
```
 docker build --no-cache -t stress:latest .
 ```

## RUN
```
 docker run stress --url=http://www.google.com --requests=20 --concurrency=3
```