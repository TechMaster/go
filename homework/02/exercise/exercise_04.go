package exercise

import "sort"

/*
Bài 4 Một nhân viên trong công ty bao gồm các thuộc tính sau : Tên, Hệ số lương, Tiền trợ cấp
Tạo 1 mảng nhân viên (số lượng tuỳ ý) và thực hiện các chức năng sau:

	- Câu 1: Sắp xếp tên nhân viên tăng dần theo bảng chữ cái
	- Câu 2: Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)
	- Câu 3: Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
*/

// Khai báo kiểu dữ liệu
type Employee struct {
	Name string
	SalaryRatio int 
	BonusMoney  int
}

// Câu 1: Sắp xếp tên nhân viên tăng dần theo bảng chữ cái
func SortByName(list []Employee) (result []Employee) {
	sort.Slice(list, func (i int, j int) bool {
		return list[i].Name < list[j].Name
	})

	return list
}

// Câu 2: Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)
func SortBySalary(list []Employee) (result []Employee) {
	sort.Slice(list, func (i int, j int) bool {
		return list[i].getSalary() > list[j].getSalary()
	})

	return list
}

// Công thức tính tổng lương của từng nhân viên
func (e Employee) getSalary() (int) {
	return e.SalaryRatio * 1500000 + e.BonusMoney
}

// Câu 3: Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
func Max2Salary(list []Employee) (result []Employee) {
	// Tìm mức lương có giá trị lớn thứ 2
	max2 := find2Salary(list)

	// Lấy ra danh sách nhân viên có lương lớn thứ 2
	for _, e := range list {
		if e.getSalary() == max2 {
			result = append(result, e)
		}
	}

	return result
}

// Tìm Tìm mức lương có giá trị lớn thứ 2 trong danh sách nhân viên
func find2Salary(list []Employee) (result int) {

	// Sắp xếp danh sách nhân viên giảm dần theo tổng lương
	employeesSortBySalary := SortBySalary(list)

	// Lấy ra giá trị lương lớn thứ 2
	for i := 1; i < len(employeesSortBySalary); i++ {
		if employeesSortBySalary[i].getSalary() < employeesSortBySalary[0].getSalary() {
			result = employeesSortBySalary[i].getSalary()
			break
		}
	}

	return result
}