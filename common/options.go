package libcommon

type WithFlag struct {
	Debug bool
}

type BasicAuthCredential struct {
	UseBasicAuth       bool
	Username, Password string
}

type Options struct {
	ServerAddr string
	Flag       WithFlag
	Basic      BasicAuthCredential
}
