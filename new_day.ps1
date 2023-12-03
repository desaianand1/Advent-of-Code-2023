param(
    [Parameter(Mandatory = $False)]
    [ValidateRange(1, 25)]
    [int]$SelectedDay = (Get-Date).ToString("%d")
)
function New-Day {
    param (
        [Parameter(Mandatory)]
        $Day,
        [String]$InputFileName = "input.txt"
    )
    $Year = (Get-Date).Year
    $Domain = "adventofcode.com"
    $InputURL = "https://www.$Domain/$Year/day/$Day/input"
    Write-Output  "url: $InputURL"

    $DayOutputDir = "day$Day"
    if (Test-Path $DayOutputDir) {
        # if directory already exists, stop
        Write-Error -Message "$DayOutputDir already exists! Cannot create new day's files" -ErrorAction Break        
    }
    New-Item -Path "/" -Name $DayOutputDir -ItemType Directory
    $DayInputFile = Join-Path -Path $DayOutputDir -ChildPath $InputFileName
    if (Test-Path $DayInputFile) {
        Write-Error -Message "Cannot create input file since $DayInputFile already exists!" -ErrorAction Break
    }

    $CookieVal = Get-Cookie
    # Create input.txt file in current day directory
    New-Day-Input-File -Domain $Domain -InputURL $InputURL -CookieVal $CookieVal -OutputDir $DayInputFile
    # Copy template files to current day directory
    $PyTemplate = Join-Path -Path "template" -ChildPath "template.py"
    $GoTemplate = Join-Path -Path "template" -ChildPath "template.go"
    $PyFile = Join-Path -Path $DayOutputDir -ChildPath "soln.py"
    $GoFile = Join-Path -Path $DayOutputDir -ChildPath "soln.go"

    Copy-Item $PyTemplate $PyFile
    Copy-Item $GoTemplate $GoFile
}

function New-Day-Input-File {
    param (
        [Parameter(Mandatory)]
        [String]$Domain,
        [Parameter(Mandatory)]
        [String]$InputURL,
        [Parameter(Mandatory)]
        [String]$CookieVal,
        [Parameter(Mandatory)]
        [String]$OutputDir
    )
    $Session = New-Object Microsoft.PowerShell.Commands.WebRequestSession
    $Cookie = New-Object System.Net.Cookie
    $Cookie.Name = "session"
    $Cookie.Value = $CookieVal
    $Cookie.Domain = $Domain
    $Session.Cookies.Add($Cookie)
    try {
        Invoke-WebRequest -Uri $InputURL -WebSession $Session -OutFile $OutputDir
        Write-Information "Succesfully downloaded Advent of Code input to $OutputDir"
    }
    catch {
        $StatusCode = $_.Exception.Response.StatusCode.Value__
        Write-Error -Message "HTTP $StatusCode : Failed to fetch Advent of Code input! Could not create $OutputDir" -ErrorAction Continue
    }
}

function Get-Cookie {
    $DotEnv = ".env"
    if ( !(Test-Path $DotEnv)) {
        Write-Error -Message "No $DotEnv file found! Please create one and retry!" -ErrorAction Break
    }
    switch -File $DotEnv {
        default {
            $name, $value = $_.Trim() -split '=', 2
            if ($name -and $name[0] -ne '#') {
                # ignore blank and comment lines.
                return $value
            }
        }
    }
}
# Entry point:
New-Day -Day $SelectedDay