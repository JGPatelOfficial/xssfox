package printing

import (
	"strconv"

	"github.com/JGPatelOfficial/xssfox/internal/utils"
	"github.com/JGPatelOfficial/xssfox/pkg/model"
)

// ScanSummary prints the summary of the scan.
func ScanSummary(scanResult model.Result, options model.Options) {
	XSSLog("SYSTEM-M", utils.GenerateTerminalWidthLine("-"), options)
	XSSLog("SYSTEM-M", "[duration: "+scanResult.Duration.String()+"][issues: "+strconv.Itoa(len(scanResult.PoCs))+"] Finish Scan!", options)
}
