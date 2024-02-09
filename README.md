# 🎄 Advent of Code 2023 📅

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54)

My attempt at Advent of Code 2023, learning Go lang while at it

> Note: On day 1, I had to fallback to Python to bypass Go's RE2 Regex implementation choice which prevents lookahead/lookbehind assertions

**Disclaimer:** `input.txt`(s) have been omitted after the fact once I was made aware [they are not to be publicly shared](https://adventofcode.com/about#faq_copying).

## Table of Contents

- [Overview](#overview)
- [Setup](#setup)
  - [Prerequisites](#prereq)
  - [Session Cookies](#cookie)
    - [Firefox](#firefox)
    - [Chrome](#chrome)
  - [Generating a New Day](#new-day)
- [Running Code](#run-code)

## Overview <a name="overview"></a>

| Day | Name                            | Stars |
| --- | ------------------------------- | ----- |
| 01  | Trebuchet?!                     | ⭐⭐  |
| 02  | Cube Conundrum                  | ⭐⭐  |
| 03  | Gear Ratios                     | ⭐⭐  |
| 04  | Scratchcards                    | ⭐⭐  |
| 05  | If You Give A Seed A Fertilizer | ⭐⭐  |
| 06  | Wait For It                     | ⭐⭐  |
| 07  | Camel Cards                     | ⭐⭐  |
| 08  | Haunted Wasteland               | ⭐⭐  |
| 09  | Mirage Maintenance              | ⭐    |
| 10  |                                 |       |
| 11  |                                 |       |
| 12  |                                 |       |
| 13  |                                 |       |
| 14  |                                 |       |
| 15  |                                 |       |
| 16  |                                 |       |
| 17  |                                 |       |
| 18  |                                 |       |
| 19  |                                 |       |
| 20  |                                 |       |
| 21  |                                 |       |
| 22  |                                 |       |
| 23  |                                 |       |
| 24  |                                 |       |
| 25  |                                 |       |

## Setup <a name="setup"></a>

### Prerequisites <a name="prereq"></a>

I solved all of this year's problems in Go but also translated solutions to Python to benchmark them. To get started, you'll need to install both by following their respective installation instructions:

- Python (> 3.8): [Download and Install](https://www.python.org/downloads/)

- Go (> 1.2): [Download and Install](https://go.dev/doc/install)

### 🍪 Session Cookies <a name="cookie"></a>

The `new_day.ps1` (or `new_day.sh`, depending on your platform) script uses Advent of Code's authentication [session cookie](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies) to auto-fetch the day's input for the current year.

To get your own cookie, visit [Advent of Code](https://adventofcode.com/). Once logged in:

#### Firefox <a name="firefox"></a>

- Right-click and select "Inspect". In the "Storage" tab, expand "Cookies" and select `https://adventofcode.com`. Copy the cookie titled "session"

#### Chrome <a name="chrome"></a>

- Right-click and select "Inspect Element". In the "Application" tab, under "Storage", expand "Cookies" and select `https://adventofcode.com`. Copy the cookie titled "session"

Paste your session cookie data into a newly created `.env` file. (`.env.example` provides a structural example)

### 📆 Generating a New Day <a name="new-day"></a>

- Add your [Advent of Code session cookie](#cookie) to the `.env` file.

- Option 1: Run `new_day.ps1` or `new_day.sh` to create the current day's directory
- Option 2: Run `new_day.ps1 d` where `d` is a day between `1` - `25` to create that day's directory for the current year (if it doesn't already exist)
- Option 3: Run `new_day.ps1 yyyy d` where `yyyy` is a 4-digit year between `2015` and the current year, `d` is a day between `1` - `25` to create that date's directory (if it doesn't already exist)
- Navigate to the generated day's directory, equipped with the day's input and some boilerplate Go and Python files

> Note: This script was only intended to be run during the duration of Advent of Code (i.e. Dec 1 through 25 of a given year).
> Additionally, it is not currently equipped to handle mixing of years (e.g. day 3 of 2023 alongside day 8 of 2022).

## Running Code <a name="run-code"></a>

To run the code for day `d`, execute the following, replacing `<d>` with the specific day (e.g. 1 - 25)

- Go: `go run day<d>/soln.go`
- Python: `python day<d>/soln.py`
