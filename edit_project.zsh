RED_TEXT='\e[31m'
YELLOW_TEXT='\e[33m'
GREEN_TEXT='\e[32m'
DEFAULT_TEXT='\e[0m'

ERROR="${RED_TEXT}[ERROR]${DEFAULT_TEXT}"
WARNING="${YELLOW_TEXT}[WARNING]${DEFAULT_TEXT}"

EDIT_PROJECT_SCRIPT="${HOME}/.local/share/edit_project.sh"
EDIT_PROJECT_CONFIG="${HOME}/.config/edit_project.conf"

function open_project() {
    local input="$1"
    local path="$(bash "$EDIT_PROJECT_SCRIPT" "$input")"
    if [ $? -ne 0 ]; then
        return 1
    fi
    cd "$path" || return 1
}

function edit_project() {
    # load config to get $editor and validate
    source "$EDIT_PROJECT_CONFIG" 2>/dev/null || true
    [[ -z "$editor" ]] && echo "${ERROR} could not load configuration file: $EDIT_PROJECT_CONFIG" >&2 && return 1

    local input="$1"

    # if no input, just open $editor
    if [[ -z "$input" ]]; then
        $editor
        return 0
    fi

    # if input is a relative path to a file or directory, open it directly
    if [[ -f "$input" || -d "$input" ]]; then
        $editor "$input"
        return 0
    fi

    # otherwise, resolve the project path and open it
    local path="$(bash "$EDIT_PROJECT_SCRIPT" "$input")"
    if [ $? -ne 0 ]; then
        return 1
    fi
    cd "$path" || return 1
    $editor .
    return 0
}

