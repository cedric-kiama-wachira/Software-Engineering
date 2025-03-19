This chapter is not only about variables, as they don't work alone. It will include the basics of other blocks of GO code plus the logical componet as well, and they are:

1. packages.
2. namespaces.
3. functions.
4. variables, constants and declaration.
   4.1 global variables.
   4.2 local variables.
   4.3 constants.
   4.4 variable declarations.
5. main.go file
6. more on variables and types as listed in sections{4.1..4.4} above.
7. idiomacy

Before writting any code we need to start with asking the what? and why? questions.

1. What is a package? and why do we need them?
   According to the core principal of Go, it is used to put together different pieces of code that work towards a common goal. 
   These different pieces include: variables, functions etc
   Packages are required to organize the code pieces for future reusability, encapsulation, namespace management and future collaboration.

2. What is a function and why do we need them?
   It is a block of code that is not mandatory and according to Go principal performs a specific task. 
   We need them as they are essential and will tell us what the program intent is.

3. What is a namespace? and why do we need them?
   They are not pieces of code, rather are labels logically created when we create our packages after which they are used as labels for 
   mitigating against future conflicts by helping us organize different pieces of code when running our program.

4. What is a variable? Why do we need a variable?
   A variable is a container in a program, that is defined within the code, and it's made up of a name, data type, plus a value. 
   It is needed because it references a location in computer memory where the value is stored, allowing the program to access or modify it later when the code runs.

5. what is the main.go file and why is it required?
   It is a collection of code blocks and specifically it contains the main() function and wich serves as an entry point for our program.
   without it our code will never be executable.

6. Expanding further on variables: 
   6.1 What is a global variable? and, why is it required?
       They are containers that give data access to other blocks of code and are usually declared or defined at the top of the go file.
       They are non mandatory and are mostly used by functions within the file.
   6.2 What is a local variable? and, why is it required?
       They are defined within a function, unlike a global variable, and it stores data in computer memory for use only within that function. 
       Its existence and purpose end when the function completes execution in the program.
       They are not mandatory but are good for Isolation of what specific blocks of code are supposed to do.
   6.3 what is a constant? and, why is it required?
       They are not variables, but serve the purpose of explicitly defining data that will never change in the computers memory.
       They are non mandatory and code can run without them. They may be required if we are looking at ways to incorporate mathematics functions or mandatory configuration settings for other blocks of code. They can be outside a block of code or at the top of the code file.
   6.4 What is a variable declaration? and why is it required?
       This is an act of defining any type of variable. It has two versions, one is explicit and the other short. 
       Explicit declaration contains keyword of the variable, the name, to be assigned, the data type to be used,
       an equals(=)sign and the value to be stored in the computer memory.
       Short hand version of variable declartion will include the name, the := sign and value.
7. Make sure your code is following what is accepted by the community via idiomacy.



