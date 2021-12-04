# ubl-store

[![Go](https://github.com/codingtroop/ubl-store/actions/workflows/go.yml/badge.svg)](https://github.com/codingtroop/ubl-store/actions/workflows/go.yml)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/codingtroop/ubl-store/issues)
[![HitCount](http://hits.dwyl.com/codingtroop/ubl-store.svg)](http://hits.dwyl.com/codingtroop/ubl-store)
[![codecov](https://codecov.io/gh/codingtroop/ubl-store/branch/main/graph/badge.svg?token=6E72396ORB)](https://codecov.io/gh/codingtroop/ubl-store)


ubl document storage with filesystem or AWS S3 compliant storage support.


# Run
To run on docker-compose with [minio](https://min.io/)
```
$ docker-compose up
````
Or run with filesystem
```
$ docker-compose -f docker-compose-standalone.yml up
```
<br/>


# Storage Size 

10.000 Sample Invoice from 3 issuer

|           | Size    | Compress Rate | Saving |
|-----------|---------:|------:|------:|
| **Original**  | 2,443,6 MB | 100.0% |  0.0% |
| **Zipped**   | 519.0 MB  | 21.0%  | 79.0% |
| **ubl-store** |  25.5 MB   | 1.0% | 99.0% |


<br/>

# Development Dependencies 

## Swaggo

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```