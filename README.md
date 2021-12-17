## Micro Summator

Microservice accepts many parameters with integers and returns the status and sum of these parameters in JSON format.</br>
The number of parameters can be any. Parameter names are not important.

localhost:8020/?arg1=8&arg2=5...&argN=N
```
Request:
http://localhost:8020/?arg1=8&arg2=5&arg3=2

Response:
{"Status":"OK","Result":15}
```
If the parameter names are the same, the parameter values are not Int or other errors, the service returns Status: "Error", Result: 0
```
Request:
http://localhost:8020/?arg1=8.5&arg2=5&arg3=bla-bla

Response:
{"Status":"Error","Result":0}
```
HTTP Code is always 200 for any result.
```
Request from command line: 
curl -i localhost:8020

Responce:
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 17 Dec 2021 11:21:17 GMT
Content-Length: 27
{"Status":"OK","Result":0}
```

Microservice fast, reliable and small. It accurately adds up integers and processes a large number of requests per minute.

The service uses modern technologies, has a microservice architecture and returns the result in JSON format. 

Summator microservice can take a decent place in your production if you use microservice architecture.

Do not use your servers to add numbers, use the dedicated service for this operation! :-)
