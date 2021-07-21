package s3configsource

import{
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
}

type Config struct{
	//koanf has a nice YAML parser and is already being used in OTEL - no particular dependence on it
	//Will use this just to parse the YAML and find specific keys in retrieve
	InternalConfig koanf.Koanf 


	//AWS variables that likely won't be used but nice to have in case of future modification
	Region string 

	Bucket string 

	Key string
}