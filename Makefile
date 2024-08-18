dev/templ:
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false

dev/server:
	air \
		--build.cmd "go build -o tmp/bin/main ./cmd/web/" --build.bin "tmp/bin/main --dev" --build.delay "100" \
		--build.exclude_dir "node_modules" \
		--build.include_dir "assets/static" \
		--build.include_ext "go,css" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true
		--main_only "true"

dev/tailwind:
	pnpm run tailwind:watch

dev/assets:
	air \
		--build.cmd "templ generate --notify-proxy" \
		--build.bin "true" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_dir "assets/static" \
		--build.include_ext "css"
		--main_only "true"

dev:
	make -j4 dev/templ dev/server dev/assets
