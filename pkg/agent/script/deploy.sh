function install_docker() {
    curl -sS https://get.docker.com/ | sh
}

function install_outline {
    bash -c "$(wget -qO- https://raw.githubusercontent.com/Jigsaw-Code/outline-server/master/src/server_manager/install_scripts/install_server.sh)"
}

function main() {
    install_docker
    install_outline
}

main