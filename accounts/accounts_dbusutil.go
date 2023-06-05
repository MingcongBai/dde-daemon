// Code generated by "dbusutil-gen -type Manager,User manager.go user.go"; DO NOT EDIT.

package accounts

func (v *User) setPropUserName(value string) (changed bool) {
	if v.UserName != value {
		v.UserName = value
		v.emitPropChangedUserName(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedUserName(value string) error {
	return v.service.EmitPropertyChanged(v, "UserName", value)
}

func (v *User) setPropUUID(value string) (changed bool) {
	if v.UUID != value {
		v.UUID = value
		v.emitPropChangedUUID(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedUUID(value string) error {
	return v.service.EmitPropertyChanged(v, "UUID", value)
}

func (v *User) setPropFullName(value string) (changed bool) {
	if v.FullName != value {
		v.FullName = value
		v.emitPropChangedFullName(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedFullName(value string) error {
	return v.service.EmitPropertyChanged(v, "FullName", value)
}

func (v *User) setPropUid(value string) (changed bool) {
	if v.Uid != value {
		v.Uid = value
		v.emitPropChangedUid(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedUid(value string) error {
	return v.service.EmitPropertyChanged(v, "Uid", value)
}

func (v *User) setPropGid(value string) (changed bool) {
	if v.Gid != value {
		v.Gid = value
		v.emitPropChangedGid(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedGid(value string) error {
	return v.service.EmitPropertyChanged(v, "Gid", value)
}

func (v *User) setPropHomeDir(value string) (changed bool) {
	if v.HomeDir != value {
		v.HomeDir = value
		v.emitPropChangedHomeDir(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedHomeDir(value string) error {
	return v.service.EmitPropertyChanged(v, "HomeDir", value)
}

func (v *User) setPropShell(value string) (changed bool) {
	if v.Shell != value {
		v.Shell = value
		v.emitPropChangedShell(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedShell(value string) error {
	return v.service.EmitPropertyChanged(v, "Shell", value)
}

func (v *User) setPropLocale(value string) (changed bool) {
	if v.Locale != value {
		v.Locale = value
		v.emitPropChangedLocale(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedLocale(value string) error {
	return v.service.EmitPropertyChanged(v, "Locale", value)
}

func (v *User) setPropLayout(value string) (changed bool) {
	if v.Layout != value {
		v.Layout = value
		v.emitPropChangedLayout(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedLayout(value string) error {
	return v.service.EmitPropertyChanged(v, "Layout", value)
}

func (v *User) setPropIconFile(value string) (changed bool) {
	if v.IconFile != value {
		v.IconFile = value
		v.emitPropChangedIconFile(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedIconFile(value string) error {
	return v.service.EmitPropertyChanged(v, "IconFile", value)
}

func (v *User) setPropPasswordHint(value string) (changed bool) {
	if v.PasswordHint != value {
		v.PasswordHint = value
		v.emitPropChangedPasswordHint(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedPasswordHint(value string) error {
	return v.service.EmitPropertyChanged(v, "PasswordHint", value)
}

func (v *User) setPropUse24HourFormat(value bool) (changed bool) {
	if v.Use24HourFormat != value {
		v.Use24HourFormat = value
		v.emitPropChangedUse24HourFormat(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedUse24HourFormat(value bool) error {
	return v.service.EmitPropertyChanged(v, "Use24HourFormat", value)
}

func (v *User) setPropWeekdayFormat(value int32) (changed bool) {
	if v.WeekdayFormat != value {
		v.WeekdayFormat = value
		v.emitPropChangedWeekdayFormat(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedWeekdayFormat(value int32) error {
	return v.service.EmitPropertyChanged(v, "WeekdayFormat", value)
}

func (v *User) setPropShortDateFormat(value int32) (changed bool) {
	if v.ShortDateFormat != value {
		v.ShortDateFormat = value
		v.emitPropChangedShortDateFormat(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedShortDateFormat(value int32) error {
	return v.service.EmitPropertyChanged(v, "ShortDateFormat", value)
}

func (v *User) setPropLongDateFormat(value int32) (changed bool) {
	if v.LongDateFormat != value {
		v.LongDateFormat = value
		v.emitPropChangedLongDateFormat(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedLongDateFormat(value int32) error {
	return v.service.EmitPropertyChanged(v, "LongDateFormat", value)
}

func (v *User) setPropShortTimeFormat(value int32) (changed bool) {
	if v.ShortTimeFormat != value {
		v.ShortTimeFormat = value
		v.emitPropChangedShortTimeFormat(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedShortTimeFormat(value int32) error {
	return v.service.EmitPropertyChanged(v, "ShortTimeFormat", value)
}

func (v *User) setPropLongTimeFormat(value int32) (changed bool) {
	if v.LongTimeFormat != value {
		v.LongTimeFormat = value
		v.emitPropChangedLongTimeFormat(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedLongTimeFormat(value int32) error {
	return v.service.EmitPropertyChanged(v, "LongTimeFormat", value)
}

func (v *User) setPropWeekBegins(value int32) (changed bool) {
	if v.WeekBegins != value {
		v.WeekBegins = value
		v.emitPropChangedWeekBegins(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedWeekBegins(value int32) error {
	return v.service.EmitPropertyChanged(v, "WeekBegins", value)
}

func (v *User) setPropDesktopBackgrounds(value []string) {
	v.DesktopBackgrounds = value
	v.emitPropChangedDesktopBackgrounds(value)
}

func (v *User) emitPropChangedDesktopBackgrounds(value []string) error {
	return v.service.EmitPropertyChanged(v, "DesktopBackgrounds", value)
}

func (v *User) setPropGroups(value []string) (changed bool) {
	if !isStrvEqual(v.Groups, value) {
		v.Groups = value
		v.emitPropChangedGroups(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedGroups(value []string) error {
	return v.service.EmitPropertyChanged(v, "Groups", value)
}

func (v *User) setPropGreeterBackground(value string) (changed bool) {
	if v.GreeterBackground != value {
		v.GreeterBackground = value
		v.emitPropChangedGreeterBackground(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedGreeterBackground(value string) error {
	return v.service.EmitPropertyChanged(v, "GreeterBackground", value)
}

func (v *User) setPropXSession(value string) (changed bool) {
	if v.XSession != value {
		v.XSession = value
		v.emitPropChangedXSession(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedXSession(value string) error {
	return v.service.EmitPropertyChanged(v, "XSession", value)
}

func (v *User) setPropPasswordStatus(value string) (changed bool) {
	if v.PasswordStatus != value {
		v.PasswordStatus = value
		v.emitPropChangedPasswordStatus(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedPasswordStatus(value string) error {
	return v.service.EmitPropertyChanged(v, "PasswordStatus", value)
}

func (v *User) setPropMaxPasswordAge(value int32) (changed bool) {
	if v.MaxPasswordAge != value {
		v.MaxPasswordAge = value
		v.emitPropChangedMaxPasswordAge(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedMaxPasswordAge(value int32) error {
	return v.service.EmitPropertyChanged(v, "MaxPasswordAge", value)
}

func (v *User) setPropPasswordLastChange(value int32) (changed bool) {
	if v.PasswordLastChange != value {
		v.PasswordLastChange = value
		v.emitPropChangedPasswordLastChange(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedPasswordLastChange(value int32) error {
	return v.service.EmitPropertyChanged(v, "PasswordLastChange", value)
}

func (v *User) setPropLocked(value bool) (changed bool) {
	if v.Locked != value {
		v.Locked = value
		v.emitPropChangedLocked(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedLocked(value bool) error {
	return v.service.EmitPropertyChanged(v, "Locked", value)
}

func (v *User) setPropAutomaticLogin(value bool) (changed bool) {
	if v.AutomaticLogin != value {
		v.AutomaticLogin = value
		v.emitPropChangedAutomaticLogin(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedAutomaticLogin(value bool) error {
	return v.service.EmitPropertyChanged(v, "AutomaticLogin", value)
}

func (v *User) setPropWorkspace(value int32) (changed bool) {
	if v.Workspace != value {
		v.Workspace = value
		v.emitPropChangedWorkspace(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedWorkspace(value int32) error {
	return v.service.EmitPropertyChanged(v, "Workspace", value)
}

func (v *User) setPropSystemAccount(value bool) (changed bool) {
	if v.SystemAccount != value {
		v.SystemAccount = value
		v.emitPropChangedSystemAccount(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedSystemAccount(value bool) error {
	return v.service.EmitPropertyChanged(v, "SystemAccount", value)
}

func (v *User) setPropNoPasswdLogin(value bool) (changed bool) {
	if v.NoPasswdLogin != value {
		v.NoPasswdLogin = value
		v.emitPropChangedNoPasswdLogin(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedNoPasswdLogin(value bool) error {
	return v.service.EmitPropertyChanged(v, "NoPasswdLogin", value)
}

func (v *User) setPropAccountType(value int32) (changed bool) {
	if v.AccountType != value {
		v.AccountType = value
		v.emitPropChangedAccountType(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedAccountType(value int32) error {
	return v.service.EmitPropertyChanged(v, "AccountType", value)
}

func (v *User) setPropLoginTime(value uint64) (changed bool) {
	if v.LoginTime != value {
		v.LoginTime = value
		v.emitPropChangedLoginTime(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedLoginTime(value uint64) error {
	return v.service.EmitPropertyChanged(v, "LoginTime", value)
}

func (v *User) setPropCreatedTime(value uint64) (changed bool) {
	if v.CreatedTime != value {
		v.CreatedTime = value
		v.emitPropChangedCreatedTime(value)
		return true
	}
	return false
}

func (v *User) emitPropChangedCreatedTime(value uint64) error {
	return v.service.EmitPropertyChanged(v, "CreatedTime", value)
}

func (v *User) setPropIconList(value []string) {
	v.IconList = value
	v.emitPropChangedIconList(value)
}

func (v *User) emitPropChangedIconList(value []string) error {
	return v.service.EmitPropertyChanged(v, "IconList", value)
}

func (v *User) setPropHistoryLayout(value []string) {
	v.HistoryLayout = value
	v.emitPropChangedHistoryLayout(value)
}

func (v *User) emitPropChangedHistoryLayout(value []string) error {
	return v.service.EmitPropertyChanged(v, "HistoryLayout", value)
}

func (v *Manager) setPropAllowGuest(value bool) (changed bool) {
	if v.AllowGuest != value {
		v.AllowGuest = value
		v.emitPropChangedAllowGuest(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedAllowGuest(value bool) error {
	return v.service.EmitPropertyChanged(v, "AllowGuest", value)
}

func (v *Manager) setPropGroupList(value []string) (changed bool) {
	if !isStrvEqual(v.GroupList, value) {
		v.GroupList = value
		v.emitPropChangedGroupList(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedGroupList(value []string) error {
	return v.service.EmitPropertyChanged(v, "GroupList", value)
}

func (v *Manager) setPropIsTerminalLocked(value bool) (changed bool) {
	if v.IsTerminalLocked != value {
		v.IsTerminalLocked = value
		v.emitPropChangedIsTerminalLocked(value)
		return true
	}
	return false
}

func (v *Manager) emitPropChangedIsTerminalLocked(value bool) error {
	return v.service.EmitPropertyChanged(v, "IsTerminalLocked", value)
}
