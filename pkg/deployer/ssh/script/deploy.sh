function install_docker() {
    curl -sS https://get.docker.com/ | sh
}

function install_outline {
    bash -c "$(wget -qO- https://raw.githubusercontent.com/Jigsaw-Code/outline-server/master/src/server_manager/install_scripts/install_server.sh)"
}

function write_PID {
    echo $$ > /tmp/pid
}

function main() {
    write_PID
    install_docker
    install_outline
}

main