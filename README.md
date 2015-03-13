goneybadger
===========

[![circleci](https://circleci.com/gh/agonzalezro/goneybadger.png)](https://circleci.com/gh/agonzalezro/goneybadger)
[![godoc reference](https://godoc.org/github.com/agonzalezro/goneybadger?status.png)](https://godoc.org/github.com/agonzalezro/goneybadger)

It's a simple wrapper around the Honeybadger API that allows you to notify it
with the errors on your application.

It was mainly created to be used in the
[logrus](https://github.com/Sirupsen/logrus) hook, but it's an standalone app
as well.

How to use
----------

    c := goneybadger.New(
        "YOUR_API_KEY",
        "YOUR_CURRENT_ENVIRONMENT. Ex: live",
    )
    c.Notify("Your awesome message here!")

You can as well add a timeout for the http calls using the method
`NewWithTimeout` which will receive an extra parameter of type `time.Duration`.

TODO
----

- Honeybadger allows backtraces POSTing but it's not being used.
