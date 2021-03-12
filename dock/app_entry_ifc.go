/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package dock

import (
	"errors"
	"os"
	"syscall"
	"time"

	dbus "pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbusutil"
)

func (e *AppEntry) GetInterfaceName() string {
	return entryDBusInterface
}

func (entry *AppEntry) Activate(timestamp uint32) *dbus.Error {
	logger.Debug("Activate timestamp:", timestamp)
	m := entry.manager
	if HideModeType(m.HideMode.Get()) == HideModeSmartHide {
		m.setPropHideState(HideStateShow)
		m.updateHideState(true)
	}

	entry.PropsMu.RLock()
	hasWindow := entry.hasWindow()
	entry.PropsMu.RUnlock()

	if !hasWindow {
		entry.launchApp(timestamp)
		return nil
	}

	if entry.current == nil {
		err := errors.New("entry.current is nil")
		logger.Warning(err)
		return dbusutil.ToError(err)
	}

	winInfo := entry.current
	var err error
	if m.isActiveWindow(winInfo) {
		sd, _ := m.waylandWM.IsShowingDesktop(0)
		if winInfo.isMinimized() || sd {
			err = winInfo.activate()
		} else {
			if len(entry.windows) == 1 {
				err = winInfo.minimize()
			} else {
				nextWinInfo := entry.findNextLeader()
				if nextWinInfo != nil {
					err = nextWinInfo.activate()
				} else {
					err = errors.New("nextWinInfo is nil")
				}
			}
		}

	} else {
		err = winInfo.activate()
	}

	if err != nil {
		logger.Warning(err)
	}

	return dbusutil.ToError(err)
}

func (e *AppEntry) HandleMenuItem(timestamp uint32, id string) *dbus.Error {
	logger.Debugf("HandleMenuItem id: %q timestamp: %v", id, timestamp)
	menu := e.Menu.getMenu()
	if menu != nil {
		err := menu.HandleAction(id, timestamp)
		return dbusutil.ToError(err)
	}
	logger.Warning("HandleMenuItem failed: entry.coreMenu is nil")
	return nil
}

func (e *AppEntry) HandleDragDrop(timestamp uint32, files []string) *dbus.Error {
	logger.Debugf("handle drag drop files: %v, timestamp: %v", files, timestamp)

	ai := e.appInfo
	if ai != nil {
		e.manager.launch(ai.GetFileName(), timestamp, files)
	} else {
		logger.Warning("not supported")
	}
	return nil
}

// RequestDock 驻留
func (entry *AppEntry) RequestDock() *dbus.Error {
	docked, err := entry.manager.dockEntry(entry)
	if err != nil {
		return dbusutil.ToError(err)
	}
	if docked {
		entry.manager.saveDockedApps()
	}
	return nil
}

// RequestUndock 取消驻留
func (entry *AppEntry) RequestUndock() *dbus.Error {
	entry.manager.undockEntry(entry)
	return nil
}

func (entry *AppEntry) PresentWindows() *dbus.Error {
	entry.PropsMu.RLock()
	windowIds := entry.getWindowIds()
	entry.PropsMu.RUnlock()
	if len(windowIds) > 0 {
		entry.manager.wm.PresentWindows(dbus.FlagNoAutoStart, windowIds)
	}
	return nil
}

func (entry *AppEntry) NewInstance(timestamp uint32) *dbus.Error {
	entry.launchApp(timestamp)
	return nil
}

func (entry *AppEntry) Check() *dbus.Error {
	entry.PropsMu.RLock()
	winInfoSlice := entry.getWindowInfoSlice()
	entry.PropsMu.RUnlock()

	for _, winInfo := range winInfoSlice {
		entry.manager.attachOrDetachWindow(winInfo)
	}
	return nil
}

func (entry *AppEntry) ForceQuit() *dbus.Error {
	entry.PropsMu.RLock()
	winInfoSlice := entry.getWindowInfoSlice()
	entry.PropsMu.RUnlock()

	pidWinInfosMap := make(map[uint][]WindowInfo)
	for _, winInfo := range winInfoSlice {
		pid := winInfo.getPid()
		//winInfo.getProcess()有的时候会存在/proc/pid/下文件权限问题，导致获取的process为空，
		//改用isProcessAlive判断进程是否存在
		if pid != 0 && isProcessAlive(pid) {
			pidWinInfosMap[pid] = append(pidWinInfosMap[pid], winInfo)
		} else {
			err := winInfo.killClient()
			if err != nil {
				logger.Warning(err)
			}
		}
	}

	for pid, winInfoSlice := range pidWinInfosMap {
		err := killProcess(pid)
		if err != nil {
			logger.Warning(err)
			if os.IsPermission(err) {
				for _, winInfo := range winInfoSlice {
					err = winInfo.killClient()
					if err != nil {
						logger.Warning(err)
					}
				}
			}
		}
	}
	return nil
}

func killProcess(pid uint) error {
	if isProcessAlive(pid) {
		logger.Debug("kill process", pid)
		osP, err := os.FindProcess(int(pid))
		if err != nil {
			return err
		}
		err = osP.Signal(syscall.SIGTERM)
		if err != nil {
			logger.Warningf("failed to send signal TERM to process %d: %v",
				osP.Pid, err)
			return err
		}
		time.AfterFunc(5*time.Second, func() {
			if isProcessAlive(pid) {
				err := osP.Kill()
				if err != nil {
					logger.Warningf("failed to send signal KILL to process %d: %v",
						osP.Pid, err)
				}
			}
		})
	}
	return nil
}

// if sig is 0, then no signal is sent, but error checking is still performed;
// this can be used to check for the existence of a process ID or process group ID.
// quoted from https://man7.org/linux/man-pages/man2/kill.2.html
func isProcessAlive(pid uint) bool {
	p, err := os.FindProcess(int(pid))
	if err != nil {
		return false
	}
	err = p.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}
	return true
}

func (entry *AppEntry) GetAllowedCloseWindows() ([]uint32, *dbus.Error) {
	entry.PropsMu.RLock()
	ret := make([]uint32, len(entry.windows))
	winInfos := entry.getAllowedCloseWindows()
	for idx, winInfo := range winInfos {
		ret[idx] = uint32(winInfo.getXid())
	}
	entry.PropsMu.RUnlock()
	return ret, nil
}
