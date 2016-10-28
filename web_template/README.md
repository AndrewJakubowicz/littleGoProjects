# Templates

Ways to parse template files.

`t, err := template.ParseFiles("t.html")`

or for pattern matching:

`t, err := template.ParseGlob("*html")`


If we want Go to handle the errors for us we can use:

`t := template.Must(template.ParseFiles("t.html"))`

To execute the templates with data use:

`t.Execute(w, "Some Data")`

If you had a bunch of templates and you want to pick one use:

`t.ExecuteTemplate(w, "t2.html", "Other data")`


## Template language features

Template language features live inside `{{` and `}}`.

If statements:

```
{{ if arg }}
    Stuff in here happens. It could just be html or more template conditions...
{{ end }}
```

also

```
{{ if arg }}
    content!
{{ else }}
    other content!
{{ end }}
```

Templates use this crazy `{{ . }}` construct.
What it contains is context sensitive.

For example.
Passing an array into a template allows you do to this:

```
<body>
    <ul>
    {{ range . }}
        <li>{{ . }}</li>
    {{ else }}
        <li> Empty list .... </li>
    {{ end }}
    </ul>
</body>
```

Within the range, the `{{ . }}` refers to the current item.

With allows us to choose `{{ . }}` for blocks.
Can also have an else which runs if `with` points to a falsy value.

```
{{ with "blah" }}
 <p> Hello {{ . }}.</p>
{{ end }}
```

## Nesting templates

`{{ template "anotherTemplate.html" }}`

And you can pass arguments into those nested templates.

`{{ template "anotherTemplate.html" arg}}`

However the better way to do this is to declare the template sections.

```
{{ define "body" }}
    ...
    {{ template "content" }}
    ...
{{ end }}

{{ define "content" }}
    oooh content
{{ end }}
```

And this way, many layouts can be in the same file!

Things that can be done.

Have many files with the same `{{ define "content" }}`, but then pass in different files into the `template.ParseFiles(......)`, for the various outcomes.

What if that `{{ define "content" }}` is missing? What if it wasn't parsed?
A default can be used. Instead of `{{ template "content" }}` use `{{ block "content" .}} ... {{ end }}`.


## Variables in template land yay!

`$variable := someVal`

Examples:

```
{{ range $key, $value := . }}
    Do stuff with {{ $key }} and {{ $value }}
{{ end }}
```

## Pipes...

{{ a | b | c }}

Stuff pipes from the left to the right.
This allows us to modify stuff.

Display some number with only 1 decimal place.

```
{{ 3.152341 | printf "%.1d" }}
```

Where did that printf come from?

## Making our own functions

Attach your function to `FuncMap`.
Can take many inputs, but must return 1 or 2 (val, err) outputs.

