# elasticsearch-curator-go
Elasticsearch Curator (in Go)

Very basic implementation of some [Elasticsearch Curator](https://github.com/elastic/curator) functionality.

Currently supported
* Open/Close Index API
```
./elasticsearch-curator-go --host foobar --port 12345 --protocol https --index foobar --action close
./elasticsearch-curator-go --host localhost --port 9200 --protocol http --index foobar --action open
./elasticsearch-curator-go --host localhost --port 9200 --protocol http --index barfoo --action close
```
