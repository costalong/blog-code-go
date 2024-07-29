package demo_one

import "strconv"

type AdminInfoSrv interface {
	GetAdminInfo(id int) string
}

func (s *AdminService) GetAdminInfo(id int) string {
	return strconv.Itoa(id)
}
