# MIT License
#
# Copyright (c) 2024 buuhuu
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

# ConvertFileToHex.ps1
# ====================
#
# Converts a given file to its hexadecimal representation on byte level
# Usage: ./ConvertFileToHex.ps1 -inputFile <input file path> -outputFile <output file path>

param (
    [string]$inputFile,
    [string]$outputFile
)

function ConvertToHex {
    param (
        [string]$filePath
    )

    try {
        $bytes = [System.IO.File]::ReadAllBytes($filePath)
        $hexData = [System.BitConverter]::ToString($bytes) -replace '-'
        return $hexData
    } catch {
        Write-Error "Error occurred while reading file: $_"
        exit 1
    }
}

function ShowUsage {
    Write-Host "Usage: ./ConvertFileToHex.ps1 -inputFile <input file path> -outputFile <output file path>"
}

if (-not $inputFile -or -not $outputFile) {
    ShowUsage
    exit 1
}

if (-not (Test-Path $inputFile)) {
    Write-Error "Input file does not exist: $inputFile"
    exit 1
}

try {
    $hexContent = ConvertToHex $inputFile
    $hexContent | Out-File -FilePath $outputFile -Encoding ASCII
    Write-Host "Success! Hex data: $outputFile"
} catch {
    Write-Error "An error occurred: $_"
    exit 1
}
