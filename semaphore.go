/// batch processing in go lang
var files []string // example list file path

sem := make(chan struct{}, 1)
var wg sync.WaitGroup
for _, file := range files {
	wg.Add(1)
	//fmt.Println(file)
	sem <- struct{}{}
	go func() {
		defer wg.Done()
		defer func() { <-sem }()

		cmd := exec.Command("tesseract", file, "-")
		scan, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(scan))

	}()
}
wg.Wait()
