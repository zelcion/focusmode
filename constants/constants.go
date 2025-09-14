package constants

import "os"

var IS_DEBUG = os.Getenv("FOCUSMODE_DEBUG") == "1"

const CONFIG_FILE_NAME = "config.yaml"
const CURRENT_VERSION = "0.0.1"
const CONFIG_DIR_NAME = "focusmode" + CURRENT_VERSION

const HOSTS_FILE_PATH_LINUX = "/etc/hosts"
const HOSTS_FILE_PATH_WINDOWS = "C:\\Windows\\System32\\drivers\\etc\\hosts"
const HOSTS_FILE_PATH_MACOS = "/etc/hosts"
const HOSTS_FILE_FM_ANNOTATION = "# FocusMode edited this file"
const HOSTS_FILE_FM_BACKUP_ANNOTATION = "# Original file backup is created with .backup extension."
const FM_MANAGED_COMMENT = "# FocusMode Managed Entry"
const HOSTS_FILE_BACKUP_EXTENSION = ".backup"
const REDIRECT_IP = "127.0.0.1"

const HOSTS_FILE_PATH_DEBUG = "./hosts_debug"
