$ErrorActionPreference = "Stop"

$tests = Get-ChildItem -Directory -Exclude '.*' | ForEach-Object { $_.FullName }

go test $tests
if ($LASTEXITCODE -ne 0)
{
    throw 'testing'
}

go fmt $tests
if ($LASTEXITCODE -ne 0)
{
    throw 'formatting'
}

Write-Output 'Success!'
