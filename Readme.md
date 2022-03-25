
# Bitaksi-Driver

## For The General Project
```
├─ bin           //The folder where the binary files was created
├─ cmd           //The code that started it all
├─ config.yml    //Config file for backend
├─ go.mod        //3rd party libraries
├─ go.sum        //Sums and versions of 3rd party libraries
├─ makefile      //MakeFile for build,test and version control 
└─ internal
   ├─ api                    //Api Layer for project
   ├─ model                  //Models for every type of object
   ├─ server                 //Server Layer for all aplication.
   ├─ service                //Service Layer
   │  ├─ driver               //Service Layer for driver
      ├─ search               //Service Layer for search
   ├─ repository                //Service Layer
   │  ├─ driver               //Repository Layer for driver
   └─ version                //Version control&save for git
└─ docs         //Auto Generated Documentation
```

## ⚡️ Quick start

First of all, [download](https://golang.org/dl/) and install **Go**. :)

## For Documentation Creation
```bash
make swagger
```

## For build

```bash
make build
```
## For Test

```bash
make test
```
## For run
After Build

```bash
./bin/bitaksi-driver
```