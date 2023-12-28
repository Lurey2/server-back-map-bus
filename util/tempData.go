package util

var arrayTempory []TemporyData

type TemporyData struct {
	Code string
	Data interface{}
}

func (t *TemporyData) PushData() {
	arrayTempory = append(arrayTempory, *t)
}

func (t *TemporyData) RemoveItem(code string) bool {
	for index, td := range arrayTempory {
		if td.Code == code {
			arrayTempory = append(arrayTempory[:index], arrayTempory[(index+1):]...)
			return true
		}
	}
	return false
}

func (t *TemporyData) FindData(code string) interface{} {

	for _, td := range arrayTempory {
		if td.Code == code {
			t.Data = &t.Data
			return true
		}
	}
	return nil
}
