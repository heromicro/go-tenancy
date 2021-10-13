package model

type InitDBFunc interface {
	Init() (err error)
}

// Seed 数据填充
func Seed(InitDBFunctions ...InitDBFunc) error {
	for _, v := range InitDBFunctions {
		err := v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
