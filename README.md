# elasticsearch-curator-go
Elasticsearch Curator (in Go)

Very basic implementation of some [Elasticsearch Curator](https://github.com/elastic/curator) functionality.

Currently supported (http://localhost:9200, unauthenticated)
* Create,Delete,Open,Close Index API
```
./elasticsearch-curator-go index create mynewindex -v
./elasticsearch-curator-go index close mynewindex -v
./elasticsearch-curator-go index open mynewindex -v
./elasticsearch-curator-go index delete mynewindex -v
```
