package heron

func Init () {
    ParseJsonFile("etc/config.json", &Config)
}