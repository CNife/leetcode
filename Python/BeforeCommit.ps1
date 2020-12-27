#Requires -Version 7

$ErrorActionPreference = "Stop"

. $PSScriptRoot\venv\Scripts\Activate.ps1
$env:PYTHONPATH = $PSScriptRoot

Get-ChildItem $PSScriptRoot\code\*.py | ForEach-Object -Parallel {
    python $_.FullName
}

python -m black $PSScriptRoot\code $PSScriptRoot\leetcode
