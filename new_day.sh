#!/bin/bash

# Parse command-line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --selected_year)
            selected_year=$2
            shift
            ;;
        --selected_day)
            selected_day=$2
            shift
            ;;
        *)
            # Unknown option
            ;;
    esac
    shift
done

# Set default values if not provided
selected_year=${selected_year:-$(date +'%Y')}
selected_day=${selected_day:-$(date +'%d')}

function new_day {
    local year=$1
    local day=$2
    local input_file_name="input.txt"
    local domain="adventofcode.com"
    local input_url="https://www.$domain/$year/day/$day/input"
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
    local input_url=$2
    local cookie_val=$3
    local output_dir=$4

    if ! curl "$input_url" --compressed -H "Cookie: session=${cookie_val}" -o "$output_dir"; then
    echo "Failed to fetch Advent of Code input! Could not create $output_dir"
    rm -r "$output_dir"  # Clean up by removing the created directory
    exit 1
    else
        echo "Successfully downloaded Advent of Code input to $output_dir"
    fi
}

function get_cookie {
    dot_env=".env"
    if [ ! -e $dot_env ]; then
        echo "No $dot_env file found! Please create one and retry!"
        exit 1
    fi

    while IFS='=' read -r name value; do
        if [ -n "$name" ] && [ "${name:0:1}" != "#" ]; then
            # Ignore blank and comment lines
            echo "$value"
        fi
    done < "$dot_env"
}

# Entry point
new_day $selected_year $selected_day
