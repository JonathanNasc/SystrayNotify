package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/getlantern/systray"
	"github.com/JonathanNasc/systray-notify/icons"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	onExit := func() {
		fmt.Println("Starting onExit")
		now := time.Now()
		ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
		fmt.Println("Finished onExit")
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icons.NeutralData)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	go func() {
		systray.SetIcon(icons.NeutralData)
		systray.SetTitle("Awesome App")
		systray.SetTooltip("Pretty awesome棒棒嗒")
		mChange := systray.AddMenuItem("Change Me", "Change Me")
		mChecked := systray.AddMenuItem("Unchecked", "Check Me")
		mEnabled := systray.AddMenuItem("Enabled", "Enabled")
		systray.AddMenuItem("Ignored", "Ignored")
		mUrl := systray.AddMenuItem("Open Lantern.org", "my home")
		mQuit := systray.AddMenuItem("退出", "Quit the whole app")
		for {
			select {
			case <-mChange.ClickedCh:
				mChange.SetTitle("I've Changed")
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					mChecked.SetTitle("Unchecked")
				} else {
					mChecked.Check()
					mChecked.SetTitle("Checked")
				}
			case <-mEnabled.ClickedCh:
				mEnabled.SetTitle("Disabled")
				mEnabled.Disable()
			case <-mUrl.ClickedCh:
				open.Run("https://www.getlantern.org")
			case <-mQuit.ClickedCh:
				systray.Quit()
				fmt.Println("Quit2 now...")
				return
			}
		}
	}()
}
