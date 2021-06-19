# go-vue-demo
a gin + vue demo project.

## 1. 基本介绍

Gin-vue-demo 是一个基于vue和gin开发的全栈前后端分离的演示项目，集成jwt鉴权，动态路由，动态菜单等功能。

## 2. 使用说明

```
- node版本 > v6.9.0
- golang版本 >= v1.16
- IDE推荐：VSCode
- 初始化项目： 不同版本数据库初始化不通
```

### 2.1 web端

```bash
# enter the project directory
cd web

# install dependency
npm install

# develop
npm run serve
```

### 2.2 server端

```bash
# 使用 go.mod

# 安装go依赖包
go list (go mod tidy)

# 编译
go build
```

## 3. 目录结构

```
    ├─demo  	     （后端文件夹）
    │  ├─common         （公共功能）
    │  ├─config         （配置包）
    │  ├─controller  	（控制器）
    │  ├─global         （全局对象）
    │  ├─middleware     （中间件）
    │  ├─model          （结构体层）
    │  ├─route          （路由）
    └─frontend			 (前端入口)
    	├─ web（前端文件）
        	├─public        （发布模板）
        	└─src           （源码包）
           		├─assets	（静态文件）
            	├─components（组件）
            	├─router	（前端路由）
            	├─store     （vuex 状态管理仓）
            	├─service   （通用服务）
            	├─utils     （前端工具库）
            	└─view      （前端页面）
```

## 4. 主要功能

- 权限管理
- 文件上传。
- 登陆注册演示。

## 5. 注意事项

项目仅用于演示，请遵守Apache2.0协议并保留作者技术支持声明。