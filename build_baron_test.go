package buildbaron

import (
	"fmt"
	"github.com/evergreen-ci/evergreen/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	jiraFailure = "fake jira failed"
)

func TestTaskToJQL(t *testing.T) {
	Convey("Given a task with with two failed tests and one successful test, "+
		"the jql should contian only the failed test names", t, func() {
		task1 := model.Task{}
		task1.TestResults = []model.TestResult{
			{Status: "fail", TestFile: "foo.js"},
			{Status: "success", TestFile: "bar.js"},
			{Status: "fail", TestFile: "baz.js"},
		}
		task1.DisplayName = "foobar"
		jQL1 := taskToJQL(&task1)
		referenceJQL1 := fmt.Sprintf(JQLBFQuery, "text~\"foo.js\" or text~\"baz.js\"")
		So(jQL1, ShouldEqual, referenceJQL1)
	})

	Convey("Given a task with with oo failed tests, "+
		"the jql should contian only the failed task name", t, func() {
		task2 := model.Task{}
		task2.TestResults = []model.TestResult{}
		task2.DisplayName = "foobar"
		jQL2 := taskToJQL(&task2)
		referenceJQL2 := fmt.Sprintf(JQLBFQuery, "text~\"foobar\"")
		So(jQL2, ShouldEqual, referenceJQL2)
	})
}
