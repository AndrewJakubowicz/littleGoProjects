# Flash messages using cookies!

The general concept is to send a cookie with an `Expires` or `MaxAge` that is in the past.
This causes the cookie to be removed with a page refresh or change.

Flash message cookies should have the following 2 properties:

```
exampleCookie := http.Cookie{
    ...
    MaxAge: -1,
    Expires: time.Unix(1, 0),
    ...
}
```
