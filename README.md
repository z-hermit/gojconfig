# gojconfig
a simple golang json config reader

Usage:

#json
```
{
	"server":{
		"Port" : 9000,
		"MaxFormDataSize" : 1000,
		"StorageDir" : "./package"
	}
}
```
```
	cfg, err := configs.NewConfig("config path")
	if err != nil {
		fmt.Printf("Fail to read config: %v", err)
		os.Exit(1)
	}

	//UnsafeValue do not return value and may cause a panic while Value return err, but not terse.
	cfgServer, err := cfg.Section("server")
	port := cfgServer.UnsafeValue("Port").(float64)
	maxFormDataSize, err := cfgServer.Value("MaxFormDataSize")
	maxFormDataSize.(float64)
```