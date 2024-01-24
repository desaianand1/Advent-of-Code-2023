#!/bin/bash

# Parse command-line arguments
# Check if two arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <yyyy> <d>"
    exit 1
fi

# Parse and validate the arguments
selected_year=$1
if ! [[ $selected_year =~ ^[0-9]{4}$ ]]; then
    echo "Error: Invalid year format. Please provide a 4-digit year (yyyy)."
    exit 1
fi

# Parse and validate the second argument (day)
selected_day=$2
if ! [[ $selected_day =~ ^[1-9]$|^1[0-9]$|^2[0-5]$ ]]; then
    echo "Error: Invalid day format. Please provide a valid day (1-25)."
    exit 1
fi

# Set default values if not provided
selected_year=${selected_year:-$(date +'%Y')}
selected_day=${selected_day:-$(date +'%d')}

function new_day {
    local year=$1
    local day=$2
    local input_file_name="input.txt"
    local domain="adventofcode.com"
    local input_url="https://$domain/$year/day/$day/input"
    local day_output_dir="day$day"

    if [ -d $day_output_dir ]; then
        # if directory already exists, stop
        echo "$day_output_dir already exists! Cannot create new day's files"
        exit 1
    fi

    mkdir $day_output_dir
    day_input_file="$day_output_dir/$input_file_name"

    if [ -e $day_input_file ]; then
        echo "Cannot create input file since $day_input_file already exists!"
        exit 1
    fi

    cookie_val=$(get_cookie)
    # Create input.txt file in current day directory
    new_day_input_file $input_url $cookie_val $day_input_file
    # Copy template files to current day directory
    py_template="template/template.py"
    go_template="template/template.go"
    soln_json_template="template/soln.json"
    py_file="$day_output_dir/soln.py"
    go_file="$day_output_dir/soln.go"
    soln_json="$day_output_dir/soln.json"

    cp $py_template $py_file
    cp $go_template $go_file
    cp $soln_json_template $soln_json
}

function new_day_input_file {
    local input_url=$1
    local cookie_val=$2
    local output_dir=$3

    if ! curl "$input_url" --compressed -H "Cookie: session=${cookie_val}" -o "$output_dir"; then
        echo "Failed to fetch Advent of Code input! Could not create $output_dir"
        rm -r "$output_dir" # Clean up by removing the created directory
        exit 1
    else
        echo "Successfully downloaded Advent of Code input to $output_dir"
    fi
}

function get_cookie {

    dot_env=".env"
    echo $(grep 'SESSION_COOKIE=' $dot_env | cut -d= -f2)

}

# Entry point
new_day $selected_year $selected_day
