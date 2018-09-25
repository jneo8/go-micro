GVM_PROJ_PATH=/home/ubuntu/.gvm/pkgsets/go1.11/micro/src/go-micro

fetch_deps:
	govendor fetch -v +m
