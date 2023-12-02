# 🎄 Advent of Code 2023 📅

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)    ![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54)

My attempt at Advent of Code 2023, learning Go lang while at it
> Note: On some days I had to fallback to Python to bypass Go's RE2 Regex implementation choice which prevents lookahead/lookbehind assertions

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

| Day | Name        | Stars |
| --- | ----------- | ----- |
| 01  | Trebuchet?! | ⭐⭐ |
| 02  |             |       |
| 03  |             |       |
| 04  |             |       |
| 05  |             |       |
| 06  |             |       |
| 07  |             |       |
| 08  |             |       |
| 09  |             |       |
| 10  |             |       |
| 11  |             |       |
| 12  |             |       |
| 13  |             |       |
| 14  |             |       |
| 15  |             |       |
| 16  |             |       |
| 17  |             |       |
| 18  |             |       |
| 19  |             |       |
| 20  |             |       |
| 21  |             |       |
| 22  |             |       |
| 23  |             |       |
| 24  |             |       |
| 25  |             |       |

## Setup <a name="setup"></a>

### Prerequisites <a name="prereq"></a>

I solved most of this year's problems in Go, and a few in Python. To get started, you'll need to install both by following their respective installation instructions:

- Python (> 3.8): [Download and Install](https://www.python.org/downloads/)

- Go (> 1.2): [Download and Install](https://go.dev/doc/install)

### 🍪 Session Cookies <a name="cookie"></a>

The ```new_day.ps1``` script uses Advent of Code's authentication [session cookie](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies) to auto-fetch the day's input etc.

To get your own cookie, visit [Advent of Code](https://adventofcode.com/). Once logged in:

#### Firefox <a name="firefox"></a>

- Right-click and select "Inspect". In the "Storage" tab, expand "Cookies" and select `https://adventofcode.com`. Copy the cookie titled "session"

#### Chrome <a name="chrome"></a>

- Right-click and select "Inspect Element". In the "Application" tab, under "Storage", expand "Cookies" and select `https://adventofcode.com`. Copy the cookie titled "session"

Paste your session cookie data into your ```.env``` file. (```.env.example``` provides a structural example)

### 📆 Generating a New Day <a name="new-day"></a>

- Add your [Advent of Code session cookie](#cookie) to the ```.env``` file.

- Run ```new_day.ps1``` to create the current day's directory

> Note: This script was only intended to be run during the duration of Advent of Code (i.e. Dec 1 through 25 of a given year)

- Navigate to the generated day's directory, equipped with the day's input and some boilerplate Go and Python files

## Running Code <a name="run-code">></a>

To run the code for day ```dd```, execute the following, replacing ```<dd>``` with the specific day (e.g. 01 - 25)

- Go: ```go run day<dd>/day<dd>.go```
- Python: ```python day<dd>/day<dd>.py```
