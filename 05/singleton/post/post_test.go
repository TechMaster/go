package post

import "testing"

func TestGetInstance(t *testing.T) {
	instancePost := GetInstancePost()
	if instancePost == nil {
		t.Errorf("Post instance is not initial")
	}
}

func TestAddView(t *testing.T) {
	instancePost := GetInstancePost()
	
	instancePost.AddView()
	instancePost.AddView()

	view := instancePost.GetView()
	if view != 3 {
		t.Errorf("View was incorrect, got: %d, want: %d.", view, 3)
	 }
}

func TestCompareView(t *testing.T) {
	instancePost := GetInstancePost()
	
	for i := 0; i < 9; i++ {
		instancePost.AddView()
	}

	view := instancePost.GetView()
	if view != 10 {
		t.Errorf("View was incorrect, got: %d, want: %d.", view, 10)
	}
}
