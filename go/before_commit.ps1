$ErrorActionPreference = "Stop"

$tests = Get-ChildItem -Directory -Exclude '.*' | ForEach-Object { $_.FullName }

Write-Output 'Testing'
go test $tests
if ($LASTEXITCODE -ne 0)
{
    throw 'testing'
}

Write-Output 'Formatting'
go fmt $tests
if ($LASTEXITCODE -ne 0)
{
    throw 'formatting'
}

Write-Output 'Success!'
