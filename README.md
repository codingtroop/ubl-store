# ubl-store

[![Go](https://github.com/codingtroop/ubl-store/actions/workflows/go.yml/badge.svg)](https://github.com/codingtroop/ubl-store/actions/workflows/go.yml)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/codingtroop/ubl-store/issues)
[![HitCount](http://hits.dwyl.com/codingtroop/ubl-store.svg)](http://hits.dwyl.com/codingtroop/ubl-store)
[![codecov](https://codecov.io/gh/codingtroop/ubl-store/branch/main/graph/badge.svg?token=6E72396ORB)](https://codecov.io/gh/codingtroop/ubl-store)


## Swaggo

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

# Storage Size Comparisation

10.000 Sample Invoice from 3 different issuer

|           | Size    | Rate |
|-----------|---------|------|
| **Original**  | 2.46 GB | 100% |
| **Zipped**   | 519 MB  | 21%  |
| **ubl-store** | 41 MB   | 1.6% |