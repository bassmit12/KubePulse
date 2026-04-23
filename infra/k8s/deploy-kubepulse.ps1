$ErrorActionPreference = 'Stop'

$root = Split-Path -Parent (Split-Path -Parent $PSScriptRoot)
$backendDir = Join-Path $root 'backend-encore'
$frontendDir = Join-Path $root 'frontend'
$template = Join-Path $PSScriptRoot 'kubepulse-deployment.yaml'
$rendered = Join-Path $PSScriptRoot 'kubepulse-deployment.rendered.yaml'

$rand = Get-Random -Minimum 100000 -Maximum 999999
$backendImage = "ttl.sh/kubepulse-backend-${rand}:1h"
$frontendImage = "ttl.sh/kubepulse-frontend-${rand}:1h"

Write-Host "Building and pushing backend image: $backendImage"
Push-Location $backendDir
& 'C:\Users\bassm\.encore\bin\encore.exe' build docker $backendImage --push
Pop-Location

Write-Host "Building and pushing frontend image: $frontendImage"
Push-Location $frontendDir
docker build -t $frontendImage .
docker push $frontendImage
Pop-Location

Write-Host 'Rendering manifest with image tags...'
$yaml = Get-Content -Raw $template
$yaml = $yaml.Replace('__BACKEND_IMAGE__', $backendImage).Replace('__FRONTEND_IMAGE__', $frontendImage)
Set-Content -Path $rendered -Value $yaml -Encoding UTF8

Write-Host 'Applying manifests...'
kubectl apply -f $rendered

Write-Host 'Waiting for rollout...'
kubectl rollout status deploy/kubepulse-backend -n kubepulse --timeout=180s
kubectl rollout status deploy/kubepulse-frontend -n kubepulse --timeout=180s

Write-Host 'Done.'
Write-Host 'Use this to test locally via port-forward:'
Write-Host 'kubectl port-forward -n kubepulse svc/kubepulse-frontend 3000:80'
Write-Host 'Then open http://localhost:3000'
