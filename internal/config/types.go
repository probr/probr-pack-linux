package config

type varOptions struct {
	VarsFile     string       // Required to initialize the sdk global config object
	Verbose      bool         // Recommended for flag handling
	ServicePacks servicePacks `yaml:"ServicePacks"` // Optional
}

// servicePacks is only required if this pack accepts custom vars
type servicePacks struct {
	Ubuntu ubuntu `yaml:"Ubuntu"`
}

// Ubuntu defines the custom vars for this service pack
type ubuntu struct {
	Password      string   `yaml:"Password"`
	Ip            string   `yaml:"Ip"`
	Port          string   `yaml:"Port"`
	Username      string   `yaml:"Username"`
	TagInclusions []string `yaml:"TagInclusions"`
	TagExclusions []string `yaml:"TagExclusions"`
	Pass          string   `yaml:"Pass"`
}
