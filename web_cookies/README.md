# Cookies

A simple cookie example in go.

This example only tackles **session cookies** and **persistent cookies**.



Other cookies to look into (variants of persistent cookies):

- super cookies
- third-party cookies
- zombie cookies


## What a cookie looks like in go

```
type Cookie struct {
    Name string
    Value string
    Path string
    Domain string
    Expires time.Time
    RawExpires string
    MaxAge int
    Secure bool

    HttpOnly bool
    Raw string
    Unparsed []string
}
```

Use both `MaxAge` and `Expires` to support all browsers.
`Expires` depreciated in favor of `MaxAge` in HTTP1.1

If you leave those age related fields blank, the cookie will only exist for the browser session.

## Setting cookies

```
aCookie := http.Cookie{
    Name: "COOKIES NAME",
    HttpOnly: true,
}
http.SetCookie(w, &aCookie)

// Equivalent of doing
// w.Header().Set("Set-Cookie", aCookie.String())
// Adding another cookie requires Add()
// w.Header().Add("Set-Cookie", anotherCookie.String())
```


## Getting Cookies

Get one cookie:

```
r.Cookie(<cookiesName: string>)
```

Get all the cookies as a slice:

```
r.Cookies()
```