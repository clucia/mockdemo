package tests

import (
	"fmt"
	"github.com/clucia/mockdemo/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test000(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mys := mocks.NewMockScanner(ctrl)

	mys.EXPECT().Scan(gomock.Any()).SetArg(0, 23).Do(func(in ...interface{}) {
		fmt.Println("len in = ", len(in))
		*(in[0].(*int)) = 25
		*(in[1].(*int)) = 33
	})

	var x, y int
	mys.Scan(&x, &y)
	fmt.Println("x = ", x, ", y = ", y)
}
