$ErrorActionPreference = "Stop"

$tests = Get-ChildItem -Directory -Exclude '.*' | ForEach-Object { $_.FullName }

go test $tests
go fmt $tests
