package service

import (
	"concat/internal/util"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type DirService interface {
	ProcessDir() error
}

type dirService struct {
	inputDir   string
	outputFile string
	numWorkers int
}

func NewDirService(inputDir string, outputFile string, numWorkers int) DirService {
	return &dirService{
		inputDir:   inputDir,
		outputFile: outputFile,
		numWorkers: numWorkers,
	}
}

func (d *dirService) ProcessDir() error {
	linesCh := make(chan string)
	var wg sync.WaitGroup
	sem := NewSemaphore(d.numWorkers)
	filepath.Walk(d.inputDir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			wg.Add(1)
			go worker(util.GenerateID(), path, linesCh, &wg, sem)
		}
		return nil
	})
	go func() {
		wg.Wait()
		close(linesCh)
	}()

	of, err := os.Create(d.outputFile)
	if err != nil {
		return err
	}
	defer of.Close()

	for line := range linesCh {
		fmt.Fprintln(of, line)
	}
	return nil
}

func worker(id string, path string, ch chan string, wg *sync.WaitGroup, semaphore Semaphore) {
	semaphore.Acquire()
	log.Printf("[WORKER %s] Starting ...\n", id)
	defer func() {
		log.Printf("[WORKER %s] Finishing ...\n", id)
		semaphore.Release()
		wg.Done()
	}()

	util.ForEachLines(path, func(line string) {
		ch <- strings.ToUpper(line)
	})
}
