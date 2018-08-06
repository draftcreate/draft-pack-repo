$name = "pack-repo"
$version = "0.4.2"
$url = "https://azuredraft.blob.core.windows.net/draft/$name-v$version-windows-amd64.zip"

if ($env:TEMP -eq $null) {
  $env:TEMP = Join-Path $env:SystemDrive 'temp'
}
$tempDir = Join-Path $env:TEMP $name
if (![System.IO.Directory]::Exists($tempDir)) {[void][System.IO.Directory]::CreateDirectory($tempDir)}
$file = Join-Path $env:TEMP "$name-v$version-windows-amd64.zip"

$proxy = New-Object System.Net.WebClient
$Proxy.Proxy.Credentials = [System.Net.CredentialCache]::DefaultNetworkCredentials

Write-Output "Downloading $url"
($proxy).DownloadFile($url, $file)

$installPath = "$env:DRAFT_HOME\plugins\draft-pack-repo\bin"
if (![System.IO.Directory]::Exists($installPath)) {[void][System.IO.Directory]::CreateDirectory($installPath)}
Write-Output "Preparing to install into $installPath"

$binaryPath = "$installPath\$name.exe"
Expand-Archive -Path "$file" -DestinationPath "$tempDir" -Force
if ([System.IO.File]::Exists("$binaryPath")) {[void][System.IO.File]::Delete("$binaryPath")}
Move-Item -Path "$tempDir\windows-amd64\$name.exe" -Destination "$binaryPath"

Write-Output "$name installed into $installPath\$name.exe"
