# 项目介绍
利用百度云提供的人脸识别api，基于beego框架的人脸识别系统，由于对于前端的知识不太熟，界面写的比较乱，基本实现功能。
目前仅实现了人脸检测、人脸对比、以及人脸检测。后续将人脸查找功能(人脸注册、人脸更新、人脸删除、用户信息查询、组列表查询、组内用户列表查询、组间复制用户、组内删除用户)添加进来。
- web框架使用beego
- 前端页面使用easyUI

## 运行方式
- 从百度云中创建，将AppID、API Key、Secret Key填到配置文件conf/app.conf中
- 数据库为mysql，将host、port、user、passwd填到配置文件conf/app.conf中
- 数据库初始化脚本为database.sql，目前只有一张表存放目录结构，后续会将用户表(t_user)、用户组表(t_usergroup)添加进来
- 执行bee run beego-FaceRecognition启动

## 功能截图
- 人脸检测
![image]https://github.com/zhaojwaln/beego-FaceRecognition/blob/master/image/20171219104227.png
- 人脸识别
![image]https://github.com/zhaojwaln/beego-FaceRecognition/blob/master/image/20171219104528.png
- 人脸比对
![image]https://github.com/zhaojwaln/beego-FaceRecognition/blob/master/image/20171219104319.png

## 后续功能
- 用户管理
    - 人脸添加
    - 人脸删除
    - 人脸更新
    - 用户信息查询
- 用户组管理
    - 组列表查询
    - 组内用户列表查询
    - 组件复制用户
    - 组内删除用户