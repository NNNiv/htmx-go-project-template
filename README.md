# htmx-go-project-template

A basic GOTTH (Go, Tailwind CSS, Templ, HTMX) stack project template.

## Instructions

Install [air](https://github.com/cosmtrek/air) for live recompilation.

```
go install github.com/cosmtrek/air@latest
```

Install [templ](https://templ.guide/).

```
go install github.com/a-h/templ/cmd/templ@latest
```


`cd` into the project directory and install `tailwind`.

```
cd htmx-go-project-template
npm install -D tailwindcss
```

Run the app using `air`.
```
air
```

Navigate to http://localhost:3000 to view the app.

Edit `./internal/templates/index.templ` and reload the browser to view the changes.
