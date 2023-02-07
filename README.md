footprint-go
=====

数据库版本管控工具footprint的golang版本

# 一、用途
`footprint-go`用于数据库版本控制，适用于各种使用数据库且有数据库版本(主要是表结构与表初期数据的变化)升级需求的go应用。
> 目前只支持mysql数据库驱动，后续会逐步添加对其他数据库驱动的支持。

在每次应用升级时，如果有相应表结构或表初期数据的升级，那么就可以直接将对应的SQL脚本放到对应目录即可。`footprint-go`将为您自动执行本次添加的增量的SQL脚本。
> SQL脚本可以放到工程相应目录直接作为嵌入资源打包到二进制文件里，也可以另外指定应用所在平台的文件目录。

# 二、首次使用说明
如果你已经有了一个应用，且已经创建过数据库，那么可以按照下面的说明开始使用`footprint-go`。

## 2.1.添加依赖
在工程根目录执行:
```shell
go get gitee.com/zhaochuninhefei/footprint-go
```

如果无法下载，请配置本地go环境变量`GOPRIVATE`，允许下载`gitee.com/zhaochuninhefei/footprint-go`，如:
```shell
go env -w GOPRIVATE=gitee.com/zhaochuninhefei
```

## 2.2.添加SQL脚本
在工程创建相应目录(例如`resources/db/xxx`, xxx通常是database名)，并在该目录下添加SQL脚本，SQL脚本命名约定:
```text
<业务空间>_V<主版本号>.<次版本号>.<补丁版本号>[.扩展版本号]_<脚本自定义名称>.sql
```

说明:
- 业务空间 : 必填，用于同一database下的表集合划分，通常根据业务功能划分; 业务空间命名只支持大小写字母与数字。
- 主版本号 : 必填，一个业务空间对应的主版本号，对应"x.y.z.t"中的x，只支持非负整数。
- 次版本号 : 必填，一个业务空间对应的次版本号，对应"x.y.z.t"中的y，只支持非负整数。
- 补丁版本号 : 必填，一个业务空间对应的补丁版本号，对应"x.y.z.t"中的z，只支持非负整数。
- 扩展版本号 : 可选，一个业务空间对应的扩展版本号，对应"x.y.z.t"中的4，只支持非负整数。
- 脚本自定义名称 : 必填，该sql脚本的自定义名称，支持大小写字母，数字与下划线。

示例:
```text
resources/
└── db
    └── xxx
        ├── smtp_V3.1.0.1_add_smtp07.sql
        └── template_V3.12.1_add_template11.sql
```

### 关于业务空间的建议
一般的应用只需要一个业务空间，即应用所使用的数据库database的名称。

如果需要多个业务空间，那么务必注意，**相同表的SQL必须属于同一个业务空间**，
不同业务空间的SQL脚本的执行顺序没有保证，
同一个业务空间的SQL则按照`主版本号 > 次版本号 > 补丁版本号 > 扩展版本号`的顺序从小到大执行。

### 关于版本号的建议
版本号建议与应用版本号保持一致，每次应用升级时，如果有数据库版本升级，则将对应的相同版本号的SQL放入相应目录即可。
> 通常应用版本只有`[主版本号].[次版本号].[补丁版本号]`，如果同一个应用版本需要多个业务空间相同的SQL脚本，则可以使用`扩展版本号`加以区分。

如果某次应用升级没有数据库版本升级，则不需要添加SQL脚本，因此最终目录下的SQL脚本的版本号很可能并不连续，而是有很多跳号。

### SQL脚本的内容
SQL脚本的内容一般是DDL，如建表，修改字段等等，也可以是对某张表的数据(一般是初期数据)的增删改等等。

## 2.3.添加embed资源嵌入定义
推荐将SQL脚本直接作为嵌入资源打包到二进制文件里，这里使用`go1.16`添加的`embed.FS`包。
> 如果不打算将SQL脚本打入包中，那么这里可以省略。

示例: 

在SQL脚本根目录下创建`resources.go`，目录结构:
```text
resources/
├── db
│   └── xxx
│       ├── smtp_V3.1.0.1_add_smtp07.sql
│       └── template_V3.12.1_add_template11.sql
└── resources.go
```

`resources.go`内容如下:
```
package resources

import "embed"

// DBFilesXxx xxx的数据库SQL文件
//go:embed db
var DBFilesXxx embed.FS

```

## 2.4.添加数据库版本控制代码
在工程的初始化代码中，添加数据库版本控制代码，执行增量SQL。

例如，在`resources.go`中添加函数`InitDBCtl`，并在应用的main函数中调用它:
```
package resources

import (
	"embed"
    "gitee.com/zhaochuninhefei/footprint-go/db/mysql"
    "gitee.com/zhaochuninhefei/footprint-go/versionctl"
    "gorm.io/gorm"
)

// DBFilesXxx xxx的数据库SQL文件
//go:embed db
var DBFilesXxx embed.FS

// InitDBCtl 应用启动初期执行数据库版本控制操作
//  dbClient是应用已经获取的db客户端，如果此时尚未获取，则直接传nil，footprint-go将根据配置的数据库连接信息创建一个单独的DB客户端。
func InitDBCtl(dbClient *gorm.DB) {
    // 此处建议改为通过配置文件读取，每个应用应该都有自己的配置库，这里直接设置数据模拟。
    myProps := &versionctl.DbVersionCtlProps{
        ScriptResourceMode:               versionctl.EMBEDFS,
        ScriptDirs:                       "embedfs:db/xxx",
		// 注意这里模拟的是数据库已经手动创建，但尚未使用footprint-go的场景下，首次加入footprint-go时，
		// 需要填入一个基线版本，确保本次加入的SQL脚本的版本大于这个基线版本即可。
		// 如你所见，如果有多个业务空间，那么这里也要配置多个业务空间。
		// 其他场景下的配置，可以参考测试用例`test/footprint_test.go`。
        BaselineBusinessSpaceAndVersions: "template_V1.0.0,smtp_V1.0.0",
        DbVersionTableName:               versionctl.DefaultDbVersionTableName,
        DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
        DriverClassName:                  "mysql",
        Host:                             "localhost",
        Port:                             "3307",
        Database:                         "db_footprint_test",
        Username:                         "zhaochun1",
        Password:                         "zhaochun@GITHUB",
        ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
        BaselineReset:                    "",
        BaselineResetConditionSql:        "",
        ModifyDbVersionTable:             "",
        ModifyDbVersionTableSqlPath:      "",
    }
    // 执行数据库版本控制操作
    err := versionctl.DoDBVersionControl(dbClient, myProps, &DBFilesXxx)
    if err != nil {
		panic(err)
    }
}
```

## 2.5.刷新go.mod
执行`go mod tidy`刷新`go.mod`。

## 2.6.启动应用
在这之后，你就可以正常启动应用，你会在日志中看到`footprint-go`的相关日志，并在数据库版本控制表(默认`brood_db_version_ctl`)中看到执行的SQL脚本记录。
> 数据库版本控制表的结构参考`resources/db/versionctl/create_brood_db_version_ctl.sql`。

## 2.7.后续应用版本升级
之后应用再次升级，且有数据库版本升级时，就可以直接将增量SQL脚本放入之前约定好的目录下即可，`footprint-go`在应用重新启动时会自动将增量SQL执行一遍。

# 三、更多的例子
`footprint-go`支持以下四种场景(操作模式):
- DEPLOY_INIT : 项目首次部署，数据库没有任何表，只有一个空的database(如果没有请先创建)。该操作会生成数据库版本控制表，执行数据库初始化脚本，更新数据库版本控制表数据。
- DEPLOY_INCREASE : 项目增量部署，之前已经导入业务表与数据库版本控制表。该操作根据已有的数据库版本控制表中的记录判断哪些脚本需要执行，然后执行脚本并插入新的数据库版本记录。
- BASELINE_INIT : 一个已经上线的项目初次使用`footprint-go`，之前已经导入业务表，但没有数据库版本控制表。该操作会创建数据库版本控制表，并写入一条版本基线记录，然后基于属性配置的基线版本确定哪些脚本需要执行。执行脚本后向数据库版本控制表插入新的版本记录。
- BASELINE_RESET : 对一个已经使用数据库版本控制的项目，重置其数据库版本的基线。该操作会删除既有的数据库版本控制表，然后重新做一次`BASELINE_INIT`操作。注意该操作需要特殊的属性控制，要慎用。

前面的`二、首次使用说明`就是BASELINE_INIT场景。

其他场景下，如何配置`versionctl.DbVersionCtlProps`，实现对应的效果，请参考测试用例`test/footprint_test.go`。

footprint_test里有6个测试用例，分别对应以下场景:
- Test01_deploy_init : 首次部署项目并使用`footprint-go`，对应操作模式`DEPLOY_INIT`。
- Test02_deploy_increase : 在完成首次部署并使用`footprint-go`后，正常的版本升级部署，对应操作模式`DEPLOY_INCREASE`。
- Test03_baseline_init : 既有项目首次部署`footprint-go`，对应操作模式`BASELINE_INIT`。
- Test04_deploy_increase : 在完成既有项目首次使用`footprint-go`后，正常的版本升级部署，对应操作模式`DEPLOY_INCREASE`。
- Test05_baseline_reset : 强制重置数据库基线版本，对应操作模式`BASELINE_RESET`。
- Test06_deploy_increase : 在强制重置数据库基线版本后，正常的版本升级部署，对应操作模式`DEPLOY_INCREASE`。

在本地执行这6个用例需要事先创建一个空的database`db_footprint_test`，然后修改数据库配置(`footprint_test.go`的常量定义)，之后按顺序执行即可。

`footprint-go`会收集当前数据库的情报，包括已经存在的表的数量，是否存在数据库版本控制表，如果存在其具体记录了哪些业务空间的版本信息等等，然后配合每次应用启动时的配置，来判断应该执行哪种操作模式。

## 3.1.Test01_deploy_init
该测试案例用于说明在首次部署项目时，如何使用`footprint-go`，对应操作模式`DEPLOY_INIT`。
> 所谓首次部署项目，就是说数据库database还是空的，一张表都没有。

其效果是，在一个空的database里，创建数据库版本控制表`brood_db_version_ctl`，并执行`ScriptDirs`中定义的资源目录下满足命名规约的sql文。
> 注意空的database需要事先创建。

典型配置如下:
```
	myProps := &versionctl.DbVersionCtlProps{
		ScriptResourceMode:               versionctl.EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01",
		BaselineBusinessSpaceAndVersions: "",
		DbVersionTableName:               versionctl.DefaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql",
		Host:                             dbHost,
		Port:                             dbPort,
		Database:                         dbName,
		Username:                         dbUser,
		Password:                         dbPwd,
		ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}
```
这种场景下，只需要注意以下配置:
- ScriptDirs : 脚本目录访问路径，多个时用","连接
- DbVersionTableName : 数据库版本控制表，推荐用默认的`versionctl.DefaultDbVersionTableName`
- DbVersionTableCreateSqlPath : 数据库版本控制表的建表语句，推荐用默认的`versionctl.DefaultDbVersionTableCreateSqlPath`，如果这里自定义，请确保表结构与默认一致，只有表名可以改。
- DriverClassName : 目前只支持mysql，但如果数据库是mysql8的话，则可以设置为`mysql8`，在从数据库版本控制表中查询最新版本时会使用`ROW_NUMBER() OVER()`函数，但其实没啥意义。。。
- 其他数据库相关字段 : DoDBVersionControl的参数`existDB`传nil的话，`footprint-go`使用这些配置创建一个单独的数据库连接。
- ExistTblQuerySql : 查看数据库当前有哪些表，默认`show tables`。
- 其他字段 : 该场景用不上。

## 3.2.Test02_deploy_increase
该测试案例用于模拟在使用了`footprint-go`之后的某次正常的数据库版本升级的场景，对应操作模式`DEPLOY_INCREASE`。
> 该测试用例在Test01_deploy_init之后执行，此时数据库中已经存在业务表和数据库版本控制表，且版本控制表中存在已经导入的脚本及其版本信息的记录。

典型配置如下:
```
	myProps := &versionctl.DbVersionCtlProps{
		ScriptResourceMode:               versionctl.EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01,embedfs:db/test02",
		BaselineBusinessSpaceAndVersions: "",
		DbVersionTableName:               versionctl.DefaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql",
		Host:                             dbHost,
		Port:                             dbPort,
		Database:                         dbName,
		Username:                         dbUser,
		Password:                         dbPwd,
		ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}
```
这里的不同之处是`ScriptDirs`增加了一个目录，但这是为了方便测试用例使用。在实际项目开发中，是直接将新的SQL脚本放到原先的目录下即可，所以这个属性一般不会修改。

后续每个测试用例都会新增一个脚本目录，但实际项目中往往是一直不变的，不会增加新的目录。

## 3.3.Test03_baseline_init与Test04_deploy_increase
这两个测试用例分别对应以下场景：
- Test03_baseline_init : 既有项目首次使用`footprint-go`，即之前已经上线的项目，现在引入`footprint-go`。业务表已经存在，但没有数据库版本控制表，对应操作模式`BASELINE_INIT`。
- Test04_deploy_increase : 在既有项目使用`footprint-go`之后，任意一次版本升级部署，与`test02_deploy_increase`相同，对应操作模式`DEPLOY_INCREASE`。

Test03_baseline_init的典型配置如下:
```
	myProps := &versionctl.DbVersionCtlProps{
		ScriptResourceMode:               versionctl.EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01,embedfs:db/test02,embedfs:db/test03",
		BaselineBusinessSpaceAndVersions: "template_V2.11.0,smtp_V2.0.0",
		DbVersionTableName:               versionctl.DefaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql",
		Host:                             dbHost,
		Port:                             dbPort,
		Database:                         dbName,
		Username:                         dbUser,
		Password:                         dbPwd,
		ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
		BaselineReset:                    "",
		BaselineResetConditionSql:        "",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}
```
在`BASELINE_INIT`场景下，项目已经上线了，数据库已经有业务表了，但是还没有数据库版本控制表，所以此时开始使用`footprint-go`就需要添加`BaselineBusinessSpaceAndVersions`配置。

这个配置是给出一个基线版本，小于或等于这些版本的SQL脚本将不会被执行，它们是已经上线的项目之前已经导入的表和数据。只有版本大于基线版本的SQL才会被执行。

Test04_deploy_increase的配置与Test03_baseline_init基本相同，无需再述。
> 在Test03_baseline_init的`BASELINE_INIT`操作之后，数据库版本控制表被创建，`BaselineBusinessSpaceAndVersions`指定的基线版本与可能存在的版本更新的SQL脚本被执行，其版本记录也被写入数据库版本控制表，
于是在下一次版本升级时，无论`BaselineBusinessSpaceAndVersions`有没有被修改回空值，都不会影响正常的`DEPLOY_INCREASE`版本升级。


## 3.4.Test05_baseline_reset与Test06_deploy_increase
在使用`footprint-go`之后，可能有时我们的数据库会因为某些原因还是手动执行了一些DDL或DML，并没有准备SQL脚本并通过`footprint-go`自动执行。

那么在这之后的版本升级里，我们就需要重置数据库版本控制表的基线版本。

这两个案例就分别模拟了强制重置基线，以及之后再次升级版本的场景：
- Test05_baseline_reset : 强制重置数据库基线版本，对应操作模式`BASELINE_RESET`。
- Test06_deploy_increase : 在强制重置数据库基线版本后，正常的版本升级部署，对应操作模式`DEPLOY_INCREASE`。

Test05_baseline_reset的典型配置如下：
```
	myProps := &versionctl.DbVersionCtlProps{
		ScriptResourceMode:               versionctl.EMBEDFS,
		ScriptDirs:                       "embedfs:db/test01,embedfs:db/test02,embedfs:db/test03,embedfs:db/test04,embedfs:db/test05",
		BaselineBusinessSpaceAndVersions: "template_V3.11.999,smtp_V3.0.999",
		DbVersionTableName:               versionctl.DefaultDbVersionTableName,
		DbVersionTableCreateSqlPath:      versionctl.DefaultDbVersionTableCreateSqlPath,
		DriverClassName:                  "mysql",
		Host:                             dbHost,
		Port:                             dbPort,
		Database:                         dbName,
		Username:                         dbUser,
		Password:                         dbPwd,
		ExistTblQuerySql:                 versionctl.DefaultExistTblQuerySql,
		BaselineReset:                    "y",
		BaselineResetConditionSql:        "SELECT 1 FROM brood_db_version_ctl WHERE version = 'template_V3.10.11'",
		ModifyDbVersionTable:             "",
		ModifyDbVersionTableSqlPath:      "",
	}
```
注意，`BASELINE_RESET`是一种相对比较危险的操作，要注意以下配置:
- BaselineBusinessSpaceAndVersions : 准备重置的基线版本，当BaselineReset与BaselineResetConditionSql条件满足，最终执行`BASELINE_RESET`操作时，数据库版本控制表会被重建，并插入这里指定的新的基线版本，SQL脚本的版本小于或等于它的话，不会被执行。
- BaselineReset : 设置为`y`，指明本次操作是`BASELINE_RESET`，但还要结合BaselineResetConditionSql判断数据是否匹配，才会真的做`BASELINE_RESET`，否则只做`DEPLOY_INCREASE`。
- BaselineResetConditionSql ： 当BaselineReset被设置为`y`时，使用该条sql语句查询，如果查询到的记录数大于0，则认为数据匹配，将执行`BASELINE_RESET`操作。通常建议查询数据库版本控制表，并使用时间戳字段`install_time`作为查询SQL的条件，如`SELECT 1 FROM brood_db_version_ctl where install_time = '2023-02-07 10:20:28'`，因为数据库基线版本重置操作会删除数据库版本控制表，于是该条件SQL只会生效一次，以后升级版本时，即使忘记将`BaselineReset`属性清除或设置为"n"也不会导致数据库基线版本被错误地再次重置。


# 四、对应的footprint地址
footprint是java版本的数据库版本管控工具，参考:

<a href="https://gitee.com/zhaochuninhefei/footprint" target="_blank">footprint in gitee</a>

<a href="https://github.com/zhaochuninhefei/footprint" target="_blank">footprint in github</a>


# JetBrains support
Thanks to JetBrains for supporting open source projects.

https://jb.gg/OpenSourceSupport.