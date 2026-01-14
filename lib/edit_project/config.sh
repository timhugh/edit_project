EDIT_PROJECT_CONFIG_PATH="${EDIT_PROJECT_CONFIG_PATH:-"$HOME/.config/edit_project/config.sh"}"

load_config() {
    if [[ ! -f "$EDIT_PROJECT_CONFIG_PATH" ]]; then
        echo "No config file found at $EDIT_PROJECT_CONFIG_PATH" >&2
        return 1
    fi
    source "$EDIT_PROJECT_CONFIG_PATH"
}

# write_config() {
# }

prompt_config() {
    local reconfigure="${EDIT_PROJECT_OVERWRITE_CONFIG:-"false"}"

    # set defaults
    root_dir="${HOME}/git"
    auto_clone="true"
    use_ssh="true"
    github_username=""
    editor=""

    # try to load existing config
    load_config || true

    if  [ -z "${github_username}" ]; then
        read -r -p "What is your github username? [${github_username}] " github_username_input
        github_username="${github_username_input:-${github_username}}"
    fi

    if [ "${reconfigure}" = "true" ]; then
        read -r -p "Where do you want to store your repositories? [${root_dir}] " root_dir_input
        root_dir="${root_dir_input:-${root_dir}}"

        read -r -p "Do you want to auto clone repositories if they don't exist? ($(y_n_for_value "${auto_clone}")) " auto_clone_input
        if [ -n "${auto_clone_input}" ]; then
            auto_clone="$(value_for_y_n "${auto_clone_input}")"
        fi

        read -r -p "Do you want to use SSH for cloning repositories? ($(y_n_for_value "${use_ssh}")) " use_ssh_input
        if [ -n "${use_ssh_input}" ]; then
            use_ssh="$(value_for_y_n "${use_ssh_input}")"
        fi
    fi

    if [ "${reconfigure}" = "true" ] || [ -z "${editor}" ]; then
        read -r -p "What editor do you want to use? [${editor}] " editor_input
        editor="${editor_input:-${editor}}"
    fi

    echo <<EOF
github_username="${github_username}"
root_dir="${root_dir}"
auto_clone="${auto_clone}"
use_ssh="${use_ssh}"
editor="${editor}"
EOF
}
