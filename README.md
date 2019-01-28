# HTML/JS Document Template IDE
is based on HTML and Javascript. The input output interface is described [here](#_Interface). 

## **Prerequisite**
> <a href="https://git-scm.com/downloads">install</a> and <a href="https://git.proxeus.com/snippets/1">configure</a> git 

## **Start** instruction
> ### **1.** Checkout
>
>> ```sh
>> $ git clone --recursive git@git.proxeus.com:ui/doctmpl.git
>> ```
>> OR
>> ```sh
>> $ git clone git@git.proxeus.com:ui/doctmpl.git
>> $ cd doctmpl
>> ```
>
> ### **2.** Run
>> ```sh
>># dev:~/goworkspace/src/git.proxeus.com/ui/doctmpl ./server --help
>># Usage of ./server:
>>#  -ds string
>>#        Document-Service URL (default "http://127.0.0.1:2115/")
>>#   -h string
>>#     	Host (default "127.0.0.1")
>>#   -p string
>>#     	Port (default "58084")
>> 
>> $ ./server
>> ```
>
> ### **3.** Access
>
>> <p>click <a href="http://localhost:58084/">here</a> to open the <b>IDE</b></p>
>> **main src:** [dist/view/document-tmpl-ide/main.html](./dist/view/document-tmpl-ide/main.html)

## Todo
>   - ...

## <a name="_Interface"></a>Interface
coming soon..

## **GO Build and run with all cmd's at once**
>```sh
>$ go run server_exe/embed/make/main.go && rm ./bindata.go && go build -o server server_exe/main.go && ./server
>```

## **GO Build**
>#### get/update dependencies
>```sh
>go get -u git.proxeus.com/...
>```
>#### Executable
>```sh
>$ go run server_exe/embed/make/main.go
>```
>now we should have `bindata.go` under `doctmpl/embed` with all the assets of `dist/*`
>Please note. bindata tool has a bug that causes an empty file to be created in the current working directory. To solve that we delete this file because Go does not accept empty source files.
>```sh
>$ rm ./bindata.go
>```
>**and then build it with**
>```sh
>$ go build -o server server_exe/main.go
>```
