# sql to java
a tool to generate java files from create table sql 

## 1.Features
* In view of Java language features, it supports configurable template files to generate corresponding java code

## 2.Architecture

## 3.Usage

### 3.1 Package
```
sh build.sh
```
Or download the [compiled file ](https://github.com/yugj/sql2java/releases)

### 3.2 Config template for your project
[examples](https://github.com/yugj/sql2java/tree/master/templates)

### 3.3 Generate java file
switch to sql2java folder like this:
```
├── config.json
├── sql2java
├── table.sql
└── templates
    ├── controller.tpl
    ├── dao.tpl
    ├── entity.tpl
    ├── mapper.tpl
    ├── request.tpl
    ├── service.tpl
    └── service_impl.tpl
```

run
```
./sql2java table.sql

or
./sql2java table.sql config.json
```

output
```shell
└── com
    └── yunx
        └── ldct
            ├── controller
            │   ├── ProductController.java
            │   └── UserController.java
            ├── dao
            │   ├── ProductDao.java
            │   └── UserDao.java
            ├── entity
            │   ├── ProductEntity.java
            │   └── UserEntity.java
            ├── mapper
            │   ├── ProductMapper.java
            │   └── UserMapper.java
            ├── model
            │   └── request
            │       ├── ProductRequest.java
            │       └── UserRequest.java
            └── service
                ├── ProductService.java
                ├── UserService.java
                └── impl
                    ├── ProductServiceImpl.java
                    └── UserServiceImpl.java
```

## 4. Templates Specification
### 4.1 naming
* lower case
* only support .tpl file
* support underline naming; service_impl.tpl -> XxServiceImpl.java
* template placeholder definition： {{.YourConfigKey}} 

### 4.2 config.json

[example](https://github.com/yugj/sql2java/blob/master/config.json)

```json
{
  "BasePackage": "project base package, example com.yunx.base.abc", 
  "BasePath": "The path to place the java2sql file, templates / output, default java2sql file current dir",
  "Author": "author name, default your system username",
  "ParentEntityName": "Entity parent class name, like BaseEntity",
  "ParentEntityFields": "Entity parent class fields, like id, createBy and so on, this fields will not be generated",
  "TablePrefix": "the table prefix that you dont want to act on your EntityName, t_ -> t_user -> UserEntity ",
  "OutputFormat": "true organizing directories by package path "
}
```

### 4.3 supported functions
| function |                                                   description |             example              | 
|:---------|--------------------------------------------------------------:|:--------------------------------:|
 | toLower  |   returns with all Unicode letters mapped to their lower case | {{.EntityName &verbar; toLower}} | 
 | toUpper  |   returns with all Unicode letters mapped to their upper case | {{.EntityName &verbar; toUpper}} | 

the next stage will distinguish between built-in configurations and support for user-defined configurations