@test "loads config from file" {
    source lib/edit_project/config.sh
    temp_config="$(mktemp)"
    cat > "$temp_config" <<EOF
github_username="testuser"
root_dir="/tmp/git"
auto_clone="false"
use_ssh="false"
editor="vim"
EOF
    EDIT_PROJECT_CONFIG_PATH="$temp_config"
    load_config

    [ "$github_username" = "testuser" ]
    [ "$root_dir" = "/tmp/git" ]
    [ "$auto_clone" = "false" ]
    [ "$use_ssh" = "false" ]
    [ "$editor" = "vim" ]
}
