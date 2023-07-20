// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package inputdevices

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-gir/gio-2.0"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/gsprop"
	"github.com/linuxdeepin/go-lib/strv"
	dutils "github.com/linuxdeepin/go-lib/utils"
	power "github.com/linuxdeepin/go-dbus-factory/com.deepin.system.power"
	configManager "github.com/linuxdeepin/go-dbus-factory/org.desktopspec.ConfigManager"
)

const (
	tpadSchema = "com.deepin.dde.touchpad"

	tpadKeyEnabled            = "touchpad-enabled"
	tpadKeyLeftHanded         = "left-handed"
	tpadKeyDisableWhileTyping = "disable-while-typing"
	tpadKeyNaturalScroll      = "natural-scroll"
	tpadKeyEdgeScroll         = "edge-scroll-enabled"
	tpadKeyHorizScroll        = "horiz-scroll-enabled"
	tpadKeyVertScroll         = "vert-scroll-enabled"
	tpadKeyAcceleration       = "motion-acceleration"
	tpadKeyThreshold          = "motion-threshold"
	tpadKeyScaling            = "motion-scaling"
	tpadKeyTapClick           = "tap-to-click"
	tpadKeyScrollDelta        = "delta-scroll"
	tpadKeyWhileTypingCmd     = "disable-while-typing-cmd"
	tpadKeyPalmDetect         = "palm-detect"
	tpadKeyPalmMinWidth       = "palm-min-width"
	tpadKeyPalmMinZ           = "palm-min-pressure"

	dsettingsAppID            = "org.deepin.dde.daemon"
	dsettingsInputdevices     = "org.deepin.dde.daemon.inputdevices"
	dsettingsData             = "ps2MouseAsTouchPadEnabled"
)

const (
	syndaemonPidFile = "/tmp/syndaemon.pid"
)

type Touchpad struct {
	service    *dbusutil.Service
	PropsMu    sync.RWMutex
	Exist      bool
	//存储PS/2 Mouse指针对象，用于设置gsetting开启/关闭
	ps2MousesTouchPad *mouseInfo
	DeviceList string

	// dbusutil-gen: ignore-below
	TPadEnable      gsprop.Bool `prop:"access:rw"`
	LeftHanded      gsprop.Bool `prop:"access:rw"`
	DisableIfTyping gsprop.Bool `prop:"access:rw"`
	NaturalScroll   gsprop.Bool `prop:"access:rw"`
	EdgeScroll      gsprop.Bool `prop:"access:rw"`
	HorizScroll     gsprop.Bool `prop:"access:rw"`
	VertScroll      gsprop.Bool `prop:"access:rw"`
	TapClick        gsprop.Bool `prop:"access:rw"`
	PalmDetect      gsprop.Bool `prop:"access:rw"`

	MotionAcceleration gsprop.Double `prop:"access:rw"`
	MotionThreshold    gsprop.Double `prop:"access:rw"`
	MotionScaling      gsprop.Double `prop:"access:rw"`

	DoubleClick   gsprop.Int `prop:"access:rw"`
	DragThreshold gsprop.Int `prop:"access:rw"`
	DeltaScroll   gsprop.Int `prop:"access:rw"`
	PalmMinWidth  gsprop.Int `prop:"access:rw"`
	PalmMinZ      gsprop.Int `prop:"access:rw"`

	devInfos     Touchpads
	setting      *gio.Settings
	mouseSetting *gio.Settings
}

func newTouchpad(service *dbusutil.Service) *Touchpad {
	var tpad = new(Touchpad)

	tpad.service = service
	tpad.setting = gio.NewSettings(tpadSchema)
	tpad.TPadEnable.Bind(tpad.setting, tpadKeyEnabled)
	tpad.LeftHanded.Bind(tpad.setting, tpadKeyLeftHanded)
	tpad.DisableIfTyping.Bind(tpad.setting, tpadKeyDisableWhileTyping)
	tpad.NaturalScroll.Bind(tpad.setting, tpadKeyNaturalScroll)
	tpad.EdgeScroll.Bind(tpad.setting, tpadKeyEdgeScroll)
	tpad.VertScroll.Bind(tpad.setting, tpadKeyVertScroll)
	tpad.HorizScroll.Bind(tpad.setting, tpadKeyHorizScroll)
	tpad.TapClick.Bind(tpad.setting, tpadKeyTapClick)
	tpad.PalmDetect.Bind(tpad.setting, tpadKeyPalmDetect)
	tpad.MotionAcceleration.Bind(tpad.setting, tpadKeyAcceleration)
	tpad.MotionThreshold.Bind(tpad.setting, tpadKeyThreshold)
	tpad.MotionScaling.Bind(tpad.setting, tpadKeyScaling)
	tpad.DeltaScroll.Bind(tpad.setting, tpadKeyScrollDelta)
	tpad.PalmMinWidth.Bind(tpad.setting, tpadKeyPalmMinWidth)
	tpad.PalmMinZ.Bind(tpad.setting, tpadKeyPalmMinZ)

	tpad.mouseSetting = gio.NewSettings(mouseSchema)
	tpad.DoubleClick.Bind(tpad.mouseSetting, mouseKeyDoubleClick)
	tpad.DragThreshold.Bind(tpad.mouseSetting, mouseKeyDragThreshold)

	tpad.updateDXTpads()

	return tpad
}

func (tpad *Touchpad) getDsgPS2MouseAsTouchPadEnable() bool {
	sysBus, err := dbus.SystemBus()
	if err != nil {
		return false
	}
	ds := configManager.NewConfigManager(sysBus)

	inputdevicesPath, err := ds.AcquireManager(0, dsettingsAppID, dsettingsInputdevices, "")
	if err != nil {
		logger.Warning(err)
		return false
	}
	inputdevicesDsg, err := configManager.NewManager(sysBus, inputdevicesPath)
	if err != nil {
		logger.Warning(err)
		return false
	}
	value, err := inputdevicesDsg.Value(0, dsettingsData)
	if err != nil {
		logger.Warning(err)
		return false
	}
	return value.Value().(bool)
}

func (tpad *Touchpad) init() {
	if !tpad.Exist {
		return
	}

	tpad.enable(tpad.TPadEnable.Get())
	tpad.enableLeftHanded()
	tpad.enableNaturalScroll()
	tpad.enableEdgeScroll()
	tpad.enableTapToClick()
	tpad.enableTwoFingerScroll()
	tpad.motionAcceleration()
	tpad.motionThreshold()
	tpad.motionScaling()
	tpad.disableWhileTyping()
	tpad.enablePalmDetect()
	tpad.setPalmDimensions()
}

func (tpad *Touchpad) handleDeviceChanged() {
	tpad.updateDXTpads()
	tpad.init()
}

func (tpad *Touchpad) isPS2Mouse() bool {
	if !tpad.getDsgPS2MouseAsTouchPadEnable() {
		logger.Info("isPS2Mouse Dsg org.deepin.dde.daemon.inputdevices ps2MouseAsTouchPadEnabled is false")
		return false
	}
	sysBus, err := dbus.SystemBus()
	if err != nil {
		logger.Warning(err)
		return false
	}
	sysPower := power.NewPower(sysBus)
	hasBattery, err := sysPower.HasBattery().Get(0)
	logger.Info("isPS2Mouse hasBattery : ", hasBattery)
	if !hasBattery {
		logger.Warning(err)
		return false
	}

	for _, info := range getMouseInfos(false) {
		if info.TrackPoint {
			continue
		}
		logger.Debugf("isPS2Mouse info : %v, name : %s, devNode : %s, phys : %s", info, info.Name, info.devNode, info.phys)
		if strings.Contains(info.Name, "PS/2") && strings.Contains(info.Name, "Mouse") {
			tpad.ps2MousesTouchPad = info
			break
		}
	}
	return tpad.ps2MousesTouchPad != nil
}

func (tpad *Touchpad) updateDXTpads() {
	tpad.devInfos = Touchpads{}
	for _, info := range getTPadInfos(false) {
		if !globalWayland {
			tmp := tpad.devInfos.get(info.Id)
			if tmp != nil {
				continue
			}
		}
		tpad.devInfos = append(tpad.devInfos, info)
	}

	tpad.PropsMu.Lock()
	var v string
	if len(tpad.devInfos) == 0 {
		tpad.setPropExist(tpad.isPS2Mouse())
	} else {
		tpad.setPropExist(true)
		v = tpad.devInfos.string()
	}
	tpad.setPropDeviceList(v)
	tpad.PropsMu.Unlock()
}

func (tpad *Touchpad) enable(enabled bool) {
	if len(tpad.devInfos) > 0 {
		for _, v := range tpad.devInfos {
			err := v.Enable(enabled)
			if err != nil {
				logger.Warningf("Enable '%v - %v' failed: %v",
					v.Id, v.Name, err)
			}
		}
	} else if tpad.ps2MousesTouchPad != nil {
		tpad.ps2MousesTouchPad.Enable(enabled)
	}

	enableGesture(enabled)
}

func (tpad *Touchpad) enableLeftHanded() {
	enabled := tpad.LeftHanded.Get()
	for _, v := range tpad.devInfos {
		err := v.EnableLeftHanded(enabled)
		if err != nil {
			logger.Debugf("Enable left handed '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
	setWMTPadBoolKey(wmTPadKeyLeftHanded, enabled)
}

func (tpad *Touchpad) enableNaturalScroll() {
	enabled := tpad.NaturalScroll.Get()
	for _, v := range tpad.devInfos {
		err := v.EnableNaturalScroll(enabled)
		if err != nil {
			logger.Debugf("Enable natural scroll '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
	setWMTPadBoolKey(wmTPadKeyNaturalScroll, enabled)
}

func (tpad *Touchpad) setScrollDistance() {
	delta := tpad.DeltaScroll.Get()
	for _, v := range tpad.devInfos {
		err := v.SetScrollDistance(delta, delta)
		if err != nil {
			logger.Debugf("Set natural scroll distance '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) enableEdgeScroll() {
	enabled := tpad.EdgeScroll.Get()
	for _, v := range tpad.devInfos {
		err := v.EnableEdgeScroll(enabled)
		if err != nil {
			logger.Debugf("Enable edge scroll '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
	setWMTPadBoolKey(wmTPadKeyEdgeScroll, enabled)
}

func (tpad *Touchpad) enableTwoFingerScroll() {
	vert := tpad.VertScroll.Get()
	horiz := tpad.HorizScroll.Get()
	for _, v := range tpad.devInfos {
		err := v.EnableTwoFingerScroll(vert, horiz)
		if err != nil {
			logger.Debugf("Enable two-finger scroll '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) enableTapToClick() {
	enabled := tpad.TapClick.Get()
	for _, v := range tpad.devInfos {
		err := v.EnableTapToClick(enabled)
		if err != nil {
			logger.Debugf("Enable tap to click '%v - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
	setWMTPadBoolKey(wmTPadKeyTapClick, enabled)
}

func (tpad *Touchpad) motionAcceleration() {
	accel := float32(tpad.MotionAcceleration.Get())
	for _, v := range tpad.devInfos {
		err := v.SetMotionAcceleration(accel)
		if err != nil {
			logger.Debugf("Set acceleration for '%d - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) motionThreshold() {
	thres := float32(tpad.MotionThreshold.Get())
	for _, v := range tpad.devInfos {
		err := v.SetMotionThreshold(thres)
		if err != nil {
			logger.Debugf("Set threshold for '%d - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) motionScaling() {
	scaling := float32(tpad.MotionScaling.Get())
	for _, v := range tpad.devInfos {
		err := v.SetMotionScaling(scaling)
		if err != nil {
			logger.Debugf("Set scaling for '%d - %v' failed: %v",
				v.Id, v.Name, err)
		}
	}
}

func (tpad *Touchpad) disableWhileTyping() {
	if !tpad.Exist {
		return
	}

	var usedLibinput bool = false
	enabled := tpad.DisableIfTyping.Get()
	for _, v := range tpad.devInfos {
		err := v.EnableDisableWhileTyping(enabled)
		if err != nil {
			continue
		}
		usedLibinput = true
	}
	if usedLibinput {
		return
	}

	if enabled {
		tpad.startSyndaemon()
	} else {
		tpad.stopSyndaemon()
	}
}

func (tpad *Touchpad) startSyndaemon() {
	if isSyndaemonExist(syndaemonPidFile) {
		logger.Debug("Syndaemon has running")
		return
	}

	syncmd := tpad.setting.GetString(tpadKeyWhileTypingCmd)
	if syncmd == "" {
		logger.Warning("Failed to start syndaemon, because no cmd is specified")
		return
	}
	logger.Debug("[startSyndaemon] will exec:", syncmd)
	args := strings.Split(syncmd, " ")
	argsLen := len(args)
	var cmd *exec.Cmd
	if argsLen == 1 {
		// pidfile will be created only in daemon mode
		cmd = exec.Command(args[0], "-d", "-p", syndaemonPidFile)
	} else {
		list := strv.Strv(args)
		if !list.Contains("-p") {
			if !list.Contains("-d") {
				args = append(args, "-d")
			}
			args = append(args, "-p", syndaemonPidFile)
		}
		argsLen = len(args)
		cmd = exec.Command(args[0], args[1:argsLen]...)
	}
	err := cmd.Start()
	if err != nil {
		err = os.Remove(syndaemonPidFile)
		if err != nil {
			logger.Warning("Remove error:", err)
		}
		logger.Debug("[disableWhileTyping] start syndaemon failed:", err)
		return
	}

	go func() {
		_ = cmd.Wait()
	}()
}

func (tpad *Touchpad) stopSyndaemon() {
	out, err := exec.Command("killall", "syndaemon").CombinedOutput()
	if err != nil {
		logger.Warning("[stopSyndaemon] failed:", string(out), err)
	}
	err = os.Remove(syndaemonPidFile)
	if err != nil {
		logger.Warning("remove error:", err)
	}
}

func (tpad *Touchpad) enablePalmDetect() {
	enabled := tpad.PalmDetect.Get()
	for _, dev := range tpad.devInfos {
		err := dev.EnablePalmDetect(enabled)
		if err != nil {
			logger.Warning("[enablePalmDetect] failed to enable:", dev.Id, enabled, err)
		}
	}
}

func (tpad *Touchpad) setPalmDimensions() {
	width := tpad.PalmMinWidth.Get()
	z := tpad.PalmMinZ.Get()
	for _, dev := range tpad.devInfos {
		err := dev.SetPalmDimensions(width, z)
		if err != nil {
			logger.Warning("[setPalmDimensions] failed to set:", dev.Id, width, z, err)
		}
	}
}

func isSyndaemonExist(pidFile string) bool {
	if !dutils.IsFileExist(pidFile) {
		out, err := exec.Command("pgrep", "syndaemon").CombinedOutput()
		if err != nil || len(out) < 2 {
			return false
		}
		return true
	}

	context, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return false
	}

	pid, err := strconv.ParseInt(strings.TrimSpace(string(context)), 10, 64)
	if err != nil {
		return false
	}
	var file = fmt.Sprintf("/proc/%v/cmdline", pid)
	return isProcessExist(file, "syndaemon")
}

func isProcessExist(file, name string) bool {
	context, err := ioutil.ReadFile(file)
	if err != nil {
		return false
	}

	return strings.Contains(string(context), name)
}

func enableGesture(enabled bool) {
	s, err := dutils.CheckAndNewGSettings("com.deepin.dde.gesture")
	if err != nil {
		return
	}
	if s.GetBoolean("touch-pad-enabled") == enabled {
		return
	}

	s.SetBoolean("touch-pad-enabled", enabled)
	s.Unref()
}
