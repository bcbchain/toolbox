package db

import "fmt"

func keyFirstHeight() string {
	return "/firstheight"
}
func keyLastHeight() string {
	return "/lastheight"
}

func keyOfHeader(height int64) string {
	return fmt.Sprintf("/header/%d", height)
}

func keyOfTx(height, index int64) string {
	return fmt.Sprintf("/tx/%d/%d", height, index)
}

func keyOfTxOK(height int64) string {
	return fmt.Sprintf("/tx/%d/ok", height)
}
