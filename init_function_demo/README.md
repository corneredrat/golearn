
Context: log package of golang defaultly logs to stderr

```bash
go run ./with_init.go 1>! stdout.txt 2>! stderr.txt
```

```bash
go run ./without_init.go 1>! stdout.txt 2>! stderr.txt
```