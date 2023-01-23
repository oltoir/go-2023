package main

import "fmt"

type Student struct {
	ID int
	Name string
	GPA float64
}

type Course struct {
	ID int
	Name string
	Credits int
	Students []Student
}

type Department struct {
	ID int
	Name string
	Courses []Course
}

type School struct {
	Name string
	Departments []Department
}

func (s *School) AddDepartment(department Department) {
	s.Departments = append(s.Departments, department)
}

func (s *School) RemoveDepartment(departmentID int) {
	for i, department := range s.Departments {
		if department.ID == departmentID {
			s.Departments = append(s.Departments[:i], s.Departments[i+1:]...)
			return
		}
	}
}

func (d *Department) AddCourse(course Course) {
	d.Courses = append(d.Courses, course)
}

func (d *Department) RemoveCourse(courseID int) {
	for i, course := range d.Courses {
		if course.ID == courseID {
			d.Courses = append(d.Courses[:i], d.Courses[i+1:]...)
			return
		}
	}
}

func (c *Course) AddStudent(student Student) {
	c.Students = append(c.Students, student)
}

func (c *Course) RemoveStudent(studentID int) {
	for i, student := range c.Students {
		if student.ID == studentID {
			c.Students = append(c.Students[:i], c.Students[i+1:]...)
			return
		}
	}
}

func PrintStudent(student Student) {
	fmt.Println("ID:", student.ID)
	fmt.Println("Name:", student.Name)
	fmt.Println("GPA:", student.GPA)
}

func PrintCourse(course Course) {
	fmt.Println("ID:", course.ID)
	fmt.Println("Name:", course.Name)
	fmt.Println("Credits:", course.Credits)
	fmt.Println("Students:")
	for _, student := range course.Students {
		PrintStudent(student)
	}
}

func PrintDepartment(department Department) {
	fmt.Println("ID:", department.ID)
	fmt.Println("Name:", department.Name)
	fmt.Println("Courses:")
	for _, course := range department.Courses {
		PrintCourse(course)
	}
}

func PrintSchool(school School) {
	fmt.Println("Name:", school.Name)
	fmt.Println("Departments:")
	for _, department := range school.Departments {
		PrintDepartment(department)
	}
}

func main() {
	school := School{Name: "School of Engineering"}

	department := Department{ID: 1, Name: "Computer Science"}
	department2 := Department{ID: 2, Name: "Electrical Engineering"}

	course := Course{ID: 1, Name: "Intro to Go", Credits: 3}
	course2 := Course{ID: 2, Name: "Intro to C++", Credits: 3}

	student := Student{ID: 1, Name: "John Doe", GPA: 4.0}
	student2 := Student{ID: 2, Name: "Jane Doe", GPA: 3.5}

	course.AddStudent(student)
	course.AddStudent(student2)

	department.AddCourse(course)
	department.AddCourse(course2)

	school.AddDepartment(department)
	school.AddDepartment(department2)

	// school.RemoveDepartment(2)
	PrintSchool(school)
}