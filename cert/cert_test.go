package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid err=%v\n", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference got=nil\n")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Expected GOLANG COURSE")
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	name := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaailjbzirmognetuohrymijtiy,jy-kçil"
	_, err := New(name, "bob", "2018-05-31")
	if err == nil {
		t.Error("Error should be returned on a too long course")
	}
}

func TestEmptyName(t *testing.T) {
	_, err := New("Golang", "", "2018-05-31")
	if err == nil {
		t.Error("Error should be returned on an empty name")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "azjdfheehgrhte('ejtie)kfkgjfjrefdhdhshcdhfjjesjfeeiefjekgddkfg,niegsuigegiori,gigkgrkngnfni"
	_, err := New("Golang", name, "2018-05-31")
	if err == nil {
		t.Error("Error should be returned on a too long name")
	}
}
