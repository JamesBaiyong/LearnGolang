package sd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

const (
	B = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

//服务的运行情况检测
func HealthCheck(c *gin.Context)  {
	message := "OK~"
	c.String(http.StatusOK,"\n"+message)
	//c.String(200,"\n"+"Ok~~!")
}

//硬盘使用情况
func DiskCheck(c *gin.Context)  {
	u,_ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "STATUS: OK~"

	if usedPercent >= 95 {
		status =http.StatusInternalServerError
		text  = "STATUS: CRITICAL!!!"
	} else if usedPercent >= 90{
		status = http.StatusTooManyRequests
		text = "STATUS: WARING!!!"
	}

	message := fmt.Sprintf(
		"%s - Free space:%dMB(%dGB) / %dMB(%dGB) | Used: %d%%",
			text,usedMB,usedGB,totalMB,totalGB,usedPercent,
		)
	c.String(status,"\n"+message)
}

//CPU使用情况
func CPUCheck(c *gin.Context){
	cores,_ := cpu.Counts(false)
	a,_ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15


	status := http.StatusOK
	text := "STATUS: OK~"

	if l5 >= float64(cores - 1) {
		status = http.StatusInternalServerError
		text  = "STATUS: CRITICAL!!!"
	} else if l5 >= float64(cores - 2){
		status = http.StatusTooManyRequests
		text  = "STATUS: WARING!!!"
	}

	messages := fmt.Sprintf(
		"%s - Load average: %2.f,%2.f,%2.f | Cores:%d",
		text,l1,l5,l15,cores,
	)
	c.String(status,"\n"+messages)
}

//内存使用情况
func RAMCheck(c *gin.Context)  {
	u,_ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "STATUS: OK~"

	if usedPercent >= 95{
		status =http.StatusInternalServerError
		text  = "STATUS: CRITICAL!!!"
	} else if usedPercent >= 90{
		status = http.StatusTooManyRequests
		text = "STATUS: WARING!!!"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, "\n"+message)
}
