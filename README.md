To run this:

Run a local TiDB database on port 40000

```
tiup playground --tiflash 0 --without-monitor v7.2.0
```

Create the database schema and initial data

```
mysql --comments -h 127.0.0.1 -P 4000 -u root
...
source 0001_schema.sql
source 0002_data.sql
```

Now build and run the code

```
go build
./demoblog
```

And finally open the webpage in your browser.
```
xdg-open http://127.0.0.1:8080/
```