install:
	which templ || go install github.com/a-h/templ/cmd/templ@latest
	which modd || go install github.com/cortesi/modd/cmd/modd@latest
