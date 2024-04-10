SHELL := cmd.exe
.PHONY: all build clean

all: build

build:
	cd ./frontend && pnpm install
	cd ./frontend && pnpm build
	@if exist "./internal/webapp/frontend-resources" (rmdir /s /q "./internal/webapp/frontend-resources") else (mkdir "./internal/webapp/frontend-resources")
	robocopy "./frontend/dist" "./internal/webapp/frontend-resources" /E /NFL /NDL /NJH /NJS /np

clean:
	if exist "./internal/webapp/frontend-resources" (rmdir /s /q "./internal/webapp/frontend-resources")
