package main

import (
	"flag"
	"fmt"
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"strconv"
	"strings"
	"time"
	"wow-query-updater/connections"
	"wow-query-updater/updater"
)

func main() {
	config := &Config{}
	config.LoadFromFile("config.json")

	classic := flag.Bool("classic", false, "Classic mode")
	onlyTasks := flag.String("only", "", "Run only the tasks specified in this argument (comma separated)")
	startTask := flag.String("start", "", "Skip all tasks before this one")
	concurrency := flag.Int("concurrency", 80, "Concurrent API requests")
	cacheKey := flag.String("cache-key", "1", "Cache key, used to isolate cache session data")
	flag.Parse()

	var taskList []string
	if *onlyTasks != "" {
		taskList = strings.Split(*onlyTasks, ",")
	} else {
		taskList = []string{}
	}

	schema := "public"
	if *classic {
		schema = "classic"
	}

	connections.Connect(config.Username, config.Password, config.Database, 105, schema, config.Hostname, config.Port)
	connections.DatabaseSetup(*classic)
	connections.ReportingMode = false

	cacheProvider := connections.NewPostgresCache(*cacheKey)

	connections.WowClient = blizzard_api.NewWoWClient("us", cacheProvider, *classic, *concurrency)
	connections.WowClient.CreateAccessToken(config.ClientID, config.ClientSecret, "")

    taskManager := SetupTaskManager(*concurrency, updater.LtInfo, *classic)

	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	grid := ui.NewGrid()
	termWidth, _ := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, 12)

	taskLabel := NewLabel("Current task", "Items", ui.ColorYellow)
	startedLabel := NewLabel("Started", "0", ui.ColorBlue)
	elapsedLabel := NewLabel("Elapsed", "0", ui.ColorBlue)
	modeLabel := NewLabel("Mode", "", ui.ColorYellow)
	if *classic {
		modeLabel.Text = "Classic"
	} else {
		modeLabel.Text = "Retail"
	}

	cachedRequestsLabel := NewLabel("Cached requests", "0", ui.ColorGreen)
	uncachedRequestsLabel := NewLabel("Uncached requests", "0", ui.ColorMagenta)
	maxRequestsLabel := NewLabel("Max requests per second", "0", ui.ColorMagenta)
	failedRequestsLabel := NewLabel("Failed requests", "0", ui.ColorRed)
	statusLabel := NewLabel("Status", "Running", ui.ColorYellow)

	tasksGauge := widgets.NewGauge()
	tasksGauge.Title = "Total progress"
	tasksGauge.BarColor = ui.ColorBlue

	taskGauge := widgets.NewGauge()
	taskGauge.Title = "Current task progress"
	taskGauge.BarColor = ui.ColorMagenta

	grid.Set(
		ui.NewRow(3.0/12,
			ui.NewCol(1/2.0, taskLabel),
			ui.NewCol(.3/2.0, startedLabel),
			ui.NewCol(.4/2.0, elapsedLabel),
			ui.NewCol(.3/2.0, modeLabel),
		),
		ui.NewRow(3.0/12,
			ui.NewCol(1.0, tasksGauge),
		),
		ui.NewRow(3.0/12,
			ui.NewCol(1.0, taskGauge),
		),
		ui.NewRow(3.0/12,
			ui.NewCol(1.0/6, cachedRequestsLabel),
			ui.NewCol(1.0/6, uncachedRequestsLabel),
			ui.NewCol(1.0/6, maxRequestsLabel),
			ui.NewCol(1.0/6, failedRequestsLabel),
			ui.NewCol(2.0/6, statusLabel),
		),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Millisecond * 200).C
	taskDoneChannel := make(chan int)
	startedAt := time.Now()
	startedLabel.Text = startedAt.Format("15:04:05")

	// Average reqs / s (only uncached)
	var lastCount int32
	var reqsLastSecond int32
	var maxReqsSecond int32
	reqsTicker := time.Tick(time.Second)
	go taskManager.Run(*startTask, taskList, taskDoneChannel)

	for {
		select {
		case <-reqsTicker:
			reqsLastSecond = cacheProvider.UncachedRequests - lastCount
			lastCount = cacheProvider.UncachedRequests
			if reqsLastSecond > maxReqsSecond {
				maxReqsSecond = reqsLastSecond
			}
			maxRequestsLabel.Text = fmt.Sprintf("%d (%d)", maxReqsSecond, reqsLastSecond)
			ui.Render(grid)
		case <-ticker:
			taskLabel.Text = taskManager.CurrentTask.GetName()
			elapsedLabel.Text = time.Since(startedAt).String()
			tasksGauge.Percent = taskManager.Progress
			taskGauge.Percent = taskManager.CurrentTask.GetProgress()
			cachedRequestsLabel.Text = strconv.Itoa(int(cacheProvider.CachedRequests))
			uncachedRequestsLabel.Text = strconv.Itoa(int(cacheProvider.UncachedRequests))
			failedRequestsLabel.Text = strconv.Itoa(int(taskManager.FailedRequests))
			if taskManager.Status {
				statusLabel.Text = "Running"
				statusLabel.TextStyle.Fg = ui.ColorGreen
			} else {
				statusLabel.Text = fmt.Sprintf("Suspended until %s", taskManager.ResumeTimestamp.Format("15:04:05"))
				statusLabel.TextStyle.Fg = ui.ColorRed
			}
			ui.Render(grid)
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, 12)
				ui.Clear()
				ui.Render(grid)
			}
		case <-taskDoneChannel:
			return
		}
	}
}
