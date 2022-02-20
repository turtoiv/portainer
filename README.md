I)Setup the application

1)Run the application: 

go build

validator.exe containerrepresentation

2)Run the tests:

cd validator

go test

II)Algorithm


In order to detect if a string is a valid representation we are using a stack which holds the elements from the string.
When a match is found the element is removed from the stack. A string is considered valid representation if at the end of 
the iteration the stack will be empty (all elements have been closed and remvoed from stack)