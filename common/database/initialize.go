package database

import (
	"time"

	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	toolsConfig "github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	mycasbin "github.com/go-admin-team/go-admin-core/sdk/pkg/casbin"
	toolsDB "github.com/go-admin-team/go-admin-core/tools/database"
	. "github.com/go-admin-team/go-admin-core/tools/gorm/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"go-admin/common/global"
)

// Setup 配置数据库
func Setup() {
	for k := range toolsConfig.DatabasesConfig {
		setupSimpleDatabase(k, toolsConfig.DatabasesConfig[k])
	}
}

func setupSimpleDatabase(host string, c *toolsConfig.Database) {
	if global.Driver == "" {
		global.Driver = c.Driver
	}
	if global.TablePrefix == "" && c.TablePrefix != "" {
		global.TablePrefix = c.TablePrefix + "_"
	}
	schemaName := ""
	if c.Schema != "" && global.Driver == "postgres" {
		schemaName = c.Schema
		global.TablePrefix = schemaName + "." + global.TablePrefix
	}
	log.Infof("%s => %s, schema: %s", host, pkg.Green(c.Source), pkg.Green(schemaName))
	registers := make([]toolsDB.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = toolsDB.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := toolsDB.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
	db, err := resolverConfig.Init(&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   global.TablePrefix,
		},
		Logger: New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					log.DefaultLogger.Options().Level.LevelForGorm()),
			},
		),
	}, opens[c.Driver])

	if err != nil {
		log.Fatal(pkg.Red(c.Driver+" connect error :"), err)
	} else {
		log.Info(pkg.Green(c.Driver + " connect success !"))
	}

	e := mycasbin.Setup(db, global.TablePrefix)

	sdk.Runtime.SetDb(host, db)
	sdk.Runtime.SetCasbin(host, e)
}
