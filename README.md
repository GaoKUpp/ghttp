# ghttp
http command line tool for humans

### Request

```
go version go1.13
```



### Installation

```
git clone https://github.com/ehwjh2010/ghttp.git
cd ghttp && ./make
```



### Comment

```
ghttp default use http protocol
if you want use https: ghttp -s=true

you also can create shortcuts for https: alias ghttps=ghttp -s=true
```





### Usage

+ GET 

  ```
  ghttp get httpbin.org/status/418
  ```

  ```
  ghttp get httpbin.org/get hello==world
  ```

  

+ POST

  > post request only support json now

  ```
  ghttp post example.org hello=World
  ```

  ```
  ghttp post example.org name=Tom numberList:=\[1,2,3\]
  ```



+ Header

  ```
  ghttp get httpbin.org/get testHeader:testHeader
  ```



+ Pipe line

  ```
  cat test.json | ghttp post example.org
  ```

  