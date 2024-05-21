package main

import (
  "bufio"
  "fmt"
  "github.com/cloudwego/hertz/pkg/common/hlog"
  "os"
  "os/exec"
  "strings"
)

const requirementsFile = "requirements.txt"

func main() {
  if len(os.Args) < 3 {
    hlog.Info("Usage: mpip <install|uninstall> <package>")
    os.Exit(1)
  }

  command := os.Args[1]
  packageName := os.Args[2]
  hlog.Infof("%s %s", command, packageName)
  switch command {
  case "install":
    installPackage(packageName)
  case "uninstall":
    uninstallPackage(packageName)
  default:
    hlog.Info("Invalid command. Use 'install' or 'uninstall'.")
    os.Exit(1)
  }
}

func installPackage(packageName string) {
  cmd := exec.Command("pip", "install", packageName)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    hlog.Errorf("Failed to install package %s: %v\n", packageName, err)
    os.Exit(1)
  }

  addPackageToRequirements(packageName)
}

func uninstallPackage(packageName string) {
  cmd := exec.Command("pip", "uninstall", "-y", packageName)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    fmt.Printf("Failed to uninstall package %s: %v\n", packageName, err)
    os.Exit(1)
  }

  removePackageFromRequirements(packageName)
}

func addPackageToRequirements(packageName string) {
  file, err := os.OpenFile(requirementsFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
  if err != nil {
    hlog.Errorf("Failed to open %s: %v\n", requirementsFile, err)
    os.Exit(1)
  }
  defer file.Close()

  _, err = file.WriteString(packageName + "\n")
  if err != nil {
    hlog.Errorf("Failed to write to %s: %v\n", requirementsFile, err)
    os.Exit(1)
  }
}

func removePackageFromRequirements(packageName string) {
  file, err := os.Open(requirementsFile)
  if err != nil {
    hlog.Errorf("Failed to open %s: %v\n", requirementsFile, err)
    os.Exit(1)
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    if strings.TrimSpace(line) != packageName {
      lines = append(lines, line)
    }
  }

  if err := scanner.Err(); err != nil {
    hlog.Errorf("Failed to read %s: %v\n", requirementsFile, err)
    os.Exit(1)
  }

  file, err = os.Create(requirementsFile)
  if err != nil {
    hlog.Errorf("Failed to open %s: %v\n", requirementsFile, err)
    os.Exit(1)
  }
  defer file.Close()

  for _, line := range lines {
    _, err := file.WriteString(line + "\n")
    if err != nil {
      hlog.Errorf("Failed to write to %s: %v\n", requirementsFile, err)
      os.Exit(1)
    }
  }
}
