package database

func Init() {
        initConfig()
        initConnections()
}

func Close() {
        closeConnections()
}
