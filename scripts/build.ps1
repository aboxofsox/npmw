$package = 'npmw'
$platforms = 'amd64/windows', 'amd64/linux', '386/windows', '386/linux'

if (!(test-path './bin')) {
    new-item -itemtype 'folder' -path './bin'
}

foreach ($platform in $platforms) {
    $split = $platform -split '/'

    $arch = $split[0]
    $os = $split[1]

    $Env:GOARCH = $arch
    $Env:GOOS = $os

    $out = "$package-$os-$arch"

    if ($os -eq 'windows') {
        $out += '.exe'
    }

    cmd.exe /c "go build -o ./bin/$out"

    if ($os -eq 'windows') {
        $compress = @{
            Path = "./bin/$rout"
            CompressionLevel = 'Fastest'
            DestinationPath = "./bin/$($out.replace('.exe', '')).zip"
        }
        compress-archive @compress
        remove-item "./bin/$out"
    } elseif ($os -ne 'windows') {
        if ($null -eq (get-command tar)) {
            return
        }
        tar -cvzf "./bin/$out.tgz" "./bin/$out" | out-null
        remove-item "./bin/$out"
    }
}