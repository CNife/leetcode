$ErrorActionPreference = "Stop"

$tests = Get-ChildItem -Directory -Exclude '.*' | ForEach-Object { $_.FullName }

Set-Location $PSScriptRoot
go test $tests
go fmt $tests
