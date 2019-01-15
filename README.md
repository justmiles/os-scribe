# os-scribe

Create a one-time snapshot of your filesystem's attributes. Useful for reverse enineering monolythic applications.

## Usage

This tool creates a CSV index of your filesystem. You can parse the CSV with your favorite tool. As shown in the examples below when parsing the os-scribe csv I usually lean on [textql](https://github.com/dinedal/textql).

Quick start

```sh-session
sudo os-scribe / > scribe.txt
textql -header -console scribe.txt
> .headers on
> .mode column
> SELECT * FROM scribe LIMIT 10;
```

Compress the index for exporting

```sh-session
sudo os-scribe / | gzip > scribe.txt.gz
```

Query stdout with textql

```sh-session
zcat scribe.txt.gz | textql -pretty -header -output-header -sql 'SELECT * LIMIT 10'

+----------------------------------+---------------+-------+-------+------------+----------+-----------+------+--------+---------+------------+------------+------------+--------+--------+
|               path               |   basename    | owner | group |    mode    |  inode   | hardlinks | size | blocks | blksize |   mtime    |   ctime    |   atime    | device | is_dir |
+----------------------------------+---------------+-------+-------+------------+----------+-----------+------+--------+---------+------------+------------+------------+--------+--------+
| /home/linuxuser                  | linuxuser     |  1000 |  1000 | 2147484141 | 12582914 |        50 | 4096 |      8 |    4096 | 1547486285 | 1547486285 | 1547486002 |  64769 | true   |
| /home/linuxuser/.2fa             | .2fa          |  1000 |  1000 |  134218239 | 12585715 |         1 |   35 |      0 |    4096 | 1545106200 | 1545236286 | 1547405785 |  64769 | false  |
| /home/linuxuser/.DS_Store        | .DS_Store     |  1000 |  1000 |  134218239 | 12585714 |         1 |   40 |      0 |    4096 | 1545106200 | 1545236290 | 1547406001 |  64769 | false  |
| /home/linuxuser/.ICEauthority    | .ICEauthority |  1000 |  1000 |        384 | 12582936 |         1 | 7410 |     16 |    4096 | 1547478817 | 1547478817 | 1547478819 |  64769 | false  |
| /home/linuxuser/.Xauthority      | .Xauthority   |  1000 |  1000 |        384 | 12582931 |         1 |   52 |      8 |    4096 | 1547478817 | 1547478817 | 1547478817 |  64769 | false  |
| /home/linuxuser/.anyconnect      | .anyconnect   |  1000 |  1000 |  134218239 | 12585716 |         1 |   42 |      0 |    4096 | 1545106200 | 1545236291 | 1547406001 |  64769 | false  |
| /home/linuxuser/.aptitude        | .aptitude     |  1000 |  1000 | 2147484096 | 21372385 |         2 | 4096 |      8 |    4096 | 1545276758 | 1545276758 | 1547414046 |  64769 | true   |
| /home/linuxuser/.aptitude/config | config        |  1000 |  1000 |        436 | 21372386 |         1 |   59 |      8 |    4096 | 1545276758 | 1545276758 | 1545276758 |  64769 | false  |
| /home/linuxuser/.atom            | .atom         |  1000 |  1000 |  134218239 | 12582916 |         1 |   36 |      0 |    4096 | 1545106249 | 1545236291 | 1547405453 |  64769 | false  |
| /home/linuxuser/.awless          | .awless       |  1000 |  1000 | 2147484096 | 17471414 |         4 | 4096 |      8 |    4096 | 1545237669 | 1545237669 | 1547414046 |  64769 | true   |
+----------------------------------+---------------+-------+-------+------------+----------+-----------+------+--------+---------+------------+------------+------------+--------+--------+
```

### Biggest Files

```sql
SELECT path, basename, owner, size
FROM scribe
WHERE is_dir = "false"
ORDER BY size DESC
LIMIT 10;
```

### First 10 Rows

```sql
SELECT * FROM scribe LIMIT 10;
```

### All Devices

```sql
SELECT distinct(device) FROM scribe;
```

### Last Modified

```sql
SELECT path, DATETIME(mtime, 'unixepoch') as modifed, owner, size
FROM scribe
WHERE is_dir = "false"
ORDER BY mtime DESC
LIMIT 10;
```

### Find specific files

```sql
SELECT path, DATETIME(mtime, 'unixepoch') as modifed, owner, size
FROM scribe
WHERE is_dir = "false"
AND basename LIKE "%gorelease%"
AND path NOT LIKE "%vendor%"
ORDER BY mtime DESC
LIMIT 10;
```
