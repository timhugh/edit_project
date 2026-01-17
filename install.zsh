open_project() {
	local project_dir="$(edit-cli projects search $1)" || return 1
	cd "$project_dir" || return 1
}

edit_project() {
	open_project "$1"
	local editor="$(edit-cli config editor-path)" || return 1
	"$editor" .
}
