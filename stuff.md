#we need binary code for machine to understand after compilation our code turns into that executable binary code 10101010
#main.go->go build->hello world.exe this why compiling is faster at run time 
#python is not a compiled language so if I want to give a python code to my friend I would have to give the .py file to him and must have python installed in his system this makes distribution of code difficult
#now in case of Go if i want to provide the code to my friend I would compile that go file and would just provide the .exe file to him and this reduces the need to install Go it will be completely independent and they wont need the access of the original code
# go is strongly typed and static typed
in languages like c and rust memory managment is manual , with java memory managemnt is automated and is done using the jvm that every time we run a code a virtual machine is createed that is it uses more memory as compared to c and rust go comes in between it does have a garbage collector but it inlcludes a runtime 
the purpose of go runtime is to cleanup unused memory