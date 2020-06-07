package impl

import (
	"context"
	"fmt"

	"github.com/spadesk1991/skip-micro/app/common"

	"github.com/spadesk1991/skip-micro/app/dao"

	"github.com/spadesk1991/skip-micro/app/services/pb"

	"github.com/spadesk1991/skip-micro/app/models"
)

type UsersService struct {
}

func (us UsersService) Registry(ctx context.Context, in *pb.RegRequest, res *pb.RegResponse) (e error) {
	e = common.Validate(in)
	if e != nil {
		return
	}
	u := &models.Users{
		Name: in.Name,
		PWD:  in.Pwd,
	}
	e = dao.GetDB().Create(u).Error
	if e != nil {
		return
	}
	res.Data = fmt.Sprintf(`%s注册成功`, in.Name)
	return
}
