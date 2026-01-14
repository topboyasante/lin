package middleware

type Handler func(string) string
type Middleware func(Handler) Handler

func Chain(middleware ...Middleware) Handler {
	result := func(s string) string { return s }

	/* so this is how middlewares work:
		somefunction = functionA(functionB(function c(someVal someType)))
		c must run first, returns the next function B, and then function A.
		so we are moving backwards in a loop.

		in go, the parameter of a variadic function(functions that look like this func funcName(params...SomeType)),
		is a slice of the type specified.

		so in the Chain function, middleware is []Middleware{}.

		With this in mind, we can then loop backwards in the slice.
		the last middleware in the slice will run, and then call the function before it.

		so from main.go, the slice of Middleware were (LoggingMiddleware,UppercaseMiddleware,TrimMiddleware)
		TrimMiddleware run, and returned the result of calling UppercaseMiddleware.
		the result of calling UppercaseMiddleware was a string.

		//now i'm lost in a rabbit hole and i cannot explain the rest

	 */

	for i := len(middleware) - 1; i >= 0; i-- {
		result = middleware[i](result)
	}
	return result
}
