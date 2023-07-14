package cryptograph

// is package inited
var inited bool

// private init method mostly used
// to set vars based on constant
// and not to do it in frequently
// called functions
func init() {
	hashVarSet()
}

// public init method that used
// for activate logic that can return error
func Init() error {
	if inited {
		return nil
	}

	err := initChecks()
	if err != nil {
		return err
	}

	inited = true
	return nil
}
