package hosts

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/zelcion/focusmode/constants"
)

type HostFile struct {
	Entries map[string]HostEntry
}

type HostEntry struct {
	Domain       string
	RedirectTo   string
	Active       bool
	IsFmManaged  bool
	NonEntryLine string
}

func getHostsFilePath() (string, error) {
	if constants.IS_DEBUG {
		absPath, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get working directory: %v", err)
		}
		return path.Join(absPath, constants.HOSTS_FILE_PATH_DEBUG), nil
	}

	switch os := runtime.GOOS; os {
	case "windows":
		return constants.HOSTS_FILE_PATH_WINDOWS, nil
	case "darwin":
		return constants.HOSTS_FILE_PATH_MACOS, nil
	case "linux":
		return constants.HOSTS_FILE_PATH_LINUX, nil
	default:
		return "", fmt.Errorf("unsupported operating system: %s", os)
	}
}

func backupHostsFile() error {
	hostsFilePath, err := getHostsFilePath()
	if err != nil {
		return err
	}

	input, err := os.ReadFile(hostsFilePath)
	if err != nil {
		return fmt.Errorf("failed to read hosts file for backup: %v", err)
	}

	backupPath := hostsFilePath + constants.HOSTS_FILE_BACKUP_EXTENSION
	if err := os.WriteFile(backupPath, input, 0644); err != nil {
		return fmt.Errorf("failed to write backup hosts file: %v", err)
	}

	return nil
}

func LoadHostsFile() (HostFile, error) {
	hostsFilePath, err := getHostsFilePath()
	if err != nil {
		return HostFile{}, err
	}

	file, err := os.Open(hostsFilePath)
	if err != nil {
		return HostFile{}, fmt.Errorf("failed to open hosts file: %v", err)
	}
	defer file.Close()

	var entries []HostEntry
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		entry := readLine(line)

		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return HostFile{}, fmt.Errorf("failed to read hosts file: %v", err)
	}

	entriesMap := make(map[string]HostEntry)
	for _, entry := range entries {
		if entry.Domain != "" {
			entriesMap[entry.Domain+entry.RedirectTo] = entry
			continue
		}

		entriesMap[fmt.Sprintf("nonentry_%d", len(entriesMap))] = entry
	}

	result := HostFile{Entries: entriesMap}

	// Treats the case of first load where there is no annotation yet
	if fileHasFmAnnotation(&result) {
		return result, nil
	}

	backupErr := backupHostsFile()
	if backupErr != nil {
		return result, fmt.Errorf("failed to backup hosts file: %v", backupErr)
	}

	annotated_result := *addFmAnnotationToFile(&result)

	saveErr := annotated_result.Save()
	if saveErr != nil {
		return annotated_result, fmt.Errorf("failed to save annotated hosts file: %v", saveErr)
	}

	return annotated_result, nil
}

func readLine(line string) HostEntry {
	if line == "" || strings.HasPrefix(line, "#") {
		return HostEntry{
			NonEntryLine: line,
		}
	}

	fields := strings.Fields(line)
	if len(fields) < 2 {
		return HostEntry{
			NonEntryLine: line,
		}
	}

	entry := HostEntry{
		RedirectTo:   fields[0],
		Domain:       fields[1],
		Active:       !strings.HasPrefix(line, "#"),
		IsFmManaged:  strings.Contains(line, "# FocusMode managed"),
		NonEntryLine: "",
	}

	return entry
}

func (entry *HostEntry) toLine() string {
	if entry.NonEntryLine != "" {
		return entry.NonEntryLine
	}

	activePrefix := ""
	if !entry.Active {
		activePrefix = "# "
	}
	fmComment := ""
	if entry.IsFmManaged {
		fmComment = " # FocusMode managed"
	}
	return fmt.Sprintf("%s%s %s%s", activePrefix, entry.RedirectTo, entry.Domain, fmComment)
}

func (hf *HostFile) Save() error {
	hostsFilePath, err := getHostsFilePath()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(hostsFilePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open hosts file for writing: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, entry := range hf.Entries {
		var line = entry.toLine()
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write to hosts file: %v", err)
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writes to hosts file: %v", err)
	}

	return nil
}

func fileHasFmAnnotation(hf *HostFile) bool {
	var first_entry *HostEntry
	for _, entry := range hf.Entries {
		first_entry = &entry
		break
	}
	if first_entry != nil && strings.Contains(first_entry.NonEntryLine, constants.HOSTS_FILE_FM_ANNOTATION) {
		return true
	}

	return false
}

func addFmAnnotationToFile(hf *HostFile) *HostFile {
	hasFmAnnotation := fileHasFmAnnotation(hf)
	if hasFmAnnotation {
		return hf
	}

	resultEntries := make(map[string]HostEntry)

	resultEntries[fmt.Sprintf("fm_%d", len(resultEntries))] = HostEntry{
		NonEntryLine: constants.HOSTS_FILE_FM_ANNOTATION,
	}

	resultEntries[fmt.Sprintf("fm_%d", len(resultEntries))] = HostEntry{
		NonEntryLine: constants.HOSTS_FILE_FM_BACKUP_ANNOTATION,
	}

	maps.Copy(resultEntries, hf.Entries)
	return &HostFile{Entries: resultEntries}
}

func CanEditHostsFile() (bool, error) {
	hostsFilePath, err := getHostsFilePath()
	if err != nil {
		return false, err
	}

	file, err := os.OpenFile(hostsFilePath, os.O_WRONLY, 0644)
	if err != nil {
		return false, fmt.Errorf("failed to open hosts file for writing: %v", err)
	}
	file.Close()

	return true, nil
}
