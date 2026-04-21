# hyoka — developer convenience targets.
#
# The canonical build system for hyoka is `go build`. This Makefile wraps the
# multi-step workflows that can't be expressed as a single `go` invocation,
# starting with refreshing the embedded site bundle.
#
# Usage:
#   make site-embed     # rebuild site/ and copy dist/ into the embedded tree
#   make site-build     # rebuild site/ only (output stays in site/dist/)
#   make site-install   # install site/ node_modules (npm ci)
#   make verify-embed   # fail if the embedded bundle is out of date
#
# See .squad/skills/embedded-asset-freshness/SKILL.md for the rationale.

SITE_DIR  := site
EMBED_DIR := hyoka/internal/serve/site

.PHONY: help site-install site-build site-embed verify-embed

help:
	@echo "hyoka Makefile targets:"
	@echo "  site-install   Install site/ dependencies (npm ci)"
	@echo "  site-build     Run 'npm run build' in site/"
	@echo "  site-embed     Rebuild site/ and refresh $(EMBED_DIR)/"
	@echo "  verify-embed   Rebuild and verify $(EMBED_DIR)/ matches site/dist/"

site-install:
	cd $(SITE_DIR) && npm ci

site-build:
	cd $(SITE_DIR) && npm run build

# site-embed is idempotent: rerunning with no source changes leaves the
# embedded tree byte-identical (vite content-hashes filenames deterministically).
#
# Cleanup assumption: vite's output under site/dist/ is a flat tree of
# `assets/` plus root-level files (currently `index.html`; future builds may
# emit favicons, manifests, etc.). We wipe $(EMBED_DIR)/* wholesale before
# copy so stale root-level files are pruned alongside hashed asset files.
# There are no hand-maintained files (e.g. .gitkeep) inside $(EMBED_DIR);
# embed.go points at the directory itself, not at a sentinel.
site-embed: site-build
	rm -rf $(EMBED_DIR)/*
	mkdir -p $(EMBED_DIR)
	cp -R $(SITE_DIR)/dist/. $(EMBED_DIR)/
	@echo "[site-embed] refreshed $(EMBED_DIR)/ from $(SITE_DIR)/dist/"

# verify-embed is the CI gate: rebuild and diff. Non-zero exit if the
# committed embedded bundle drifted from what site/src/** currently produces.
#
# We use `git status --porcelain` rather than `git diff --quiet` so that
# NEW files inside $(EMBED_DIR)/ (untracked by git) also trip the gate.
# `git diff` only sees modifications to tracked files and would silently
# pass a build that introduced a new asset filename.
verify-embed: site-embed
	@if [ -n "$$(git status --porcelain -- $(EMBED_DIR))" ]; then \
		echo ""; \
		echo "ERROR: embedded site bundle is stale."; \
		echo "  site/src/** has changes that were not copied into $(EMBED_DIR)/."; \
		echo "  Run 'make site-embed' and commit the result."; \
		echo ""; \
		echo "Status:"; \
		git status --porcelain -- $(EMBED_DIR); \
		echo ""; \
		echo "Diff (tracked files only):"; \
		git --no-pager diff --stat -- $(EMBED_DIR); \
		exit 1; \
	fi
	@echo "[verify-embed] $(EMBED_DIR)/ is up to date."
