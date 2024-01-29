package logs
  
  import (
  	"go.uber.org/zap/zapcore"
  	"os"
  	"sync"
  	"time"
  
  	"go.uber.org/zap"
  )
  
  var once sync.Once
  
  var instance *zap.Logger
  
  func Connect() *zap.Logger {
  
  	once.Do(func() {
  
  		config := zap.NewProductionEncoderConfig()
  		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
  		config.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
  		consoleEncoder := zapcore.NewConsoleEncoder(config)
  		defaultLogLevel := zapcore.DebugLevel
  
  		deepLevel := zapcore.FatalLevel
  		core := zapcore.NewTee(
  			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
  		)
  
  		instance = zap.New(core, zap.AddCaller(), zap.AddStacktrace(deepLevel))
  
  	})
  
  	return instance
  }
