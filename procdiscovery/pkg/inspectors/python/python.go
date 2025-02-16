package python

import (
	"debug/elf"
	"os"
	"strings"

	"github.com/hashicorp/go-version"

	"github.com/odigos-io/odigos/common"
	"github.com/odigos-io/odigos/procdiscovery/pkg/process"
)

type PythonInspector struct{}

const (
	pythonProcessName = "python"
	libPythonStr      = "libpython3"
)

func (p *PythonInspector) Inspect(proc *process.Details) (common.ProgrammingLanguage, bool) {
	if strings.Contains(proc.ExePath, pythonProcessName) || strings.Contains(proc.CmdLine, pythonProcessName) {
		return common.PythonProgrammingLanguage, true
	}

	if p.isLibPythonLinked(proc) {
		return common.PythonProgrammingLanguage, true
	}

	return "", false
}

func (p *PythonInspector) GetRuntimeVersion(proc *process.Details, containerURL string) *version.Version {
	if value, exists := proc.GetDetailedEnvsValue(process.PythonVersionConst); exists {
		return common.GetVersion(value)
	}

	return nil
}

func (p *PythonInspector) isLibPythonLinked(proc *process.Details) bool {
	if proc.Exefile == nil {
		var err error
		proc.Exefile, err = os.Open(proc.ExePath)
		if err != nil {
			return false
		}
		defer proc.Exefile.Close()
	}

	elfFile, err := elf.NewFile(proc.Exefile)
	if err != nil {
		return false
	}
	defer elfFile.Close()

	dynamicSection, err := elfFile.DynString(elf.DT_NEEDED)
	if err != nil {
		return false
	}

	for _, dep := range dynamicSection {
		if strings.Contains(dep, libPythonStr) {
			return true
		}
	}

	return false
}
