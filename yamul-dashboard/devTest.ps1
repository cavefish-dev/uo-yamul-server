$imageTag = "uo-yamul-dashboard-test"
$pubCacheVolume = "uo-yamul-pub-cache"

Write-Host "Building test image..."
docker build -t $imageTag -f "$PSScriptRoot/Dockerfile.test" "$PSScriptRoot"
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "Running Flutter tests in Docker..."
docker run --rm `
    -v "${PSScriptRoot}:/app" `
    -v "${pubCacheVolume}:/root/.pub-cache" `
    $imageTag `
    sh -c "flutter pub get && flutter test --platform chrome"
