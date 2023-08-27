$scriptPath = $MyInvocation.MyCommand.Path
$scriptDirectory = Split-Path -Path $scriptPath -Parent
Set-Location $scriptDirectory
Remove-Item *_srv/*.go
protoc  --go_out=./email_srv --go_opt=paths=source_relative --go-grpc_out=./email_srv --go-grpc_opt=paths=source_relative email.proto
protoc  --go_out=./user_srv --go_opt=paths=source_relative --go-grpc_out=./user_srv --go-grpc_opt=paths=source_relative user.proto
Write-Output Finish!
exit