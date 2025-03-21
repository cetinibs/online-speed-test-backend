# Bu script, tüm Go dosyalarındaki import yollarını günceller
$oldPath = "github.com/hizim-ne/backend"
$newPath = "github.com/cetinibs/online-speed-test-backend"

# Tüm Go dosyalarını bul
$goFiles = Get-ChildItem -Path . -Filter "*.go" -Recurse

foreach ($file in $goFiles) {
    $content = Get-Content -Path $file.FullName -Raw
    $updatedContent = $content -replace $oldPath, $newPath
    
    # Eğer değişiklik yapıldıysa dosyayı güncelle
    if ($content -ne $updatedContent) {
        Set-Content -Path $file.FullName -Value $updatedContent
        Write-Host "Updated imports in: $($file.FullName)"
    }
}

Write-Host "Import paths updated successfully!"
