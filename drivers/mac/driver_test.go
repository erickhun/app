package mac

import (
	"testing"

	"github.com/murlokswarm/app"
	"github.com/murlokswarm/app/tests"
)

func TestDriver(t *testing.T) {
	d := &Driver{
		MenubarConfig: MenuBarConfig{
			URL: "tests.menubar",
		},
		DockURL: "tests.menu",
	}

	tests.TestDriver(t, func(onRun func()) app.Driver {
		d.OnRun = onRun
		return d
	}, d.Quit)
}
