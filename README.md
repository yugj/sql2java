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

### 3.2 Config template for your project
[examples](https://github.com/yugj/sql2java/master/examples)

### 3.3 Generate java file
```
./sql2java table.sql

or
./sql2java table.sql config.json
```

## 4. Templates Specification
### 4.1 naming
* lower case
* only support .tpl file
* support underline naming; service_impl.tpl -> XxServiceImpl.java
* template placeholder definition： {{.YourConfigKey}} 

### 4.2 config.json

[example](https://github.com/yugj/sql2java/master/config.json)

```json
{
  "BasePackage": "project base package, example com.yunx.base.abc", 
  "BasePath": "The path to place the java2sql file, templates / output, default java2sql file current dir",
  "Author": "author name, default your system username",
  "ParentEntityName": "Entity parent class name, like BaseEntity",
  "ParentEntityFields": "Entity parent class fields, like id, createBy and so on, this fields will not be generated",
  "TablePrefix": "the table prefix that you dont want to act on your EntityName, t_ -> t_user -> UserEntity "
}
```

the next stage will distinguish between built-in configurations and support for user-defined configurations