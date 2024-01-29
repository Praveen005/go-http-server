## Why shouldn't we use DefaultServeMux?

The default
handler, `DefaultServeMux`. It serves as the default registry of how to
handle a request path. `DefaultServeMux` is an object of type ServeMux
defined in the http package. It is a global object, which means that
any other code you may be using in your application can also register
handler functions with your server. There is absolutely nothing
preventing a third-party rogue package from exposing an HTTP path
without you even knowing it. Additionally, as with
any other global object, this opens your code up to unforeseen
concurrency bugs and non-robust behavior. Hence, we are not going
to use it. Instead, we will create a new ServeMux object.

![DefaultServeMux Behaviour](./assets/DefaultServeMux.png)
*Figure 1: Any package can register a handler function with the `DefaultServeMux` object.*



## The Request Struct
An HTTP handler function accepts two parameters: a value of type
http.ResponseWriter and a pointer to a value of type http.Request.
The pointer object of http.Request type describes the incoming request.

Request is a struct type defined in
the net/http package. . Some key fields and methods in the struct
type that are relevant in the context of an incoming request are:

1. **Method:**
    This is a string and its value represents the HTTP method of the
request being processed. Like GET POST PUT DELETE

2. **URL:**
    This is a pointer to a value of type url.URL  representing the path of the request. 

    For:  `http://example.com/api/?
name=jane&age=25#page1.`
the URL object's fields are set as follows:
- Path: `/api/`
- RawQuery : `name=jane&age=25`
- Fragment : `page1` <br><br>
If your server accepts an HTTP basic authentication, then the request
URL would be of the form `http://user:pass@example.com/api/?
name=jane&age=25&name=john.`

    To obtain the username, call the `User() `method. To obtain the password, call the `Password()` method.<br>

    To access a specific individual query parameter and its value, you use the `Query()` method.

    - This method returns an object of type Values,
    which is defined as `map[string][]string`.
    
        - For `http://user:pass@example.com/api/?
name=jane&age=25&name=john.`, calling
the `Query()` method will return the following:<br><br>
    `url.Values{"age":[]string{"25"}, "name":[]string{"jane"}}`

        - If a query parameter is specified more than once, for example
`http://example.com/api/?name=jane&age=25&name=john#page1`, the
returned value of `Query()` would be:
<br><br>
`url.Values{"age":[]string{"25"}, "name":[]string{"jane",
"john"}}"`


<br>
<br>

3. **Proto, ProtoMajor, and ProtoMinor**
- These fields identify the HTTP protocol over which the client and
server are communicating. 
- `Proto` is a string identifying the protocol
and the version (for example "`HTTP /1.1`" ). 
- `ProtoMajor` and
`ProtoMinor` are integers identifying the major and minor protocol
versions, respectively. For `HTTP /1.1`, the major and minor protocol
versions are both 1.


4. **Header:** 
    This is a map of type `map[string][]string`, and it contains the
incoming headers.

5. **Host:** This is a string containing the hostname and port combination (
`example.com:8080` ) or the IP address and port combination (
`127.0.0.1:8080` ) that the client used to make a request to the server.

6. **Body:**
-  This is a value of type `io.ReadCloser`, which refers to the request
body
- We can use any function that understands the `io.Reader`
interface to read the body.
    - you could use the
    io.ReadAll() method to read the entire request body.
    - The function
    returns a byte slice containing the entire request body.
    - A related field is `ContentLength`, which is the maximum number of
bytes available to be read from the request body.


7. **Form, PostForm**
- If the handler function is processing an HTML form submission,
then instead of reading the body directly, we can call the
`ParseForm()` method of the request object.

- `ParseForm` populates `r.Form` and `r.PostForm`.

- For all requests, ParseForm parses the raw query from the URL and updates `r.Form`.

- If the request Body's size has not already been limited by `MaxBytesReader`, the size is capped at 10MB.

- For `POST`, `PUT`, and `PATCH` requests, it also reads the request body, parses it as a form and puts the results into both `r.PostForm` and `r.Form`

- Request body parameters take precedence over URL query string values in `r.Form`

