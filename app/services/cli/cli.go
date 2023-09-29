package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/Bruno-10/base/business/core/base"
	"github.com/Bruno-10/base/foundation/logger"
)

func main() {
	core := base.NewCore(logger.New(os.Stdout, logger.LevelInfo, "BASE", func(ctx context.Context) string { return "" }))

	fmt.Print("Enter text: ")

	r := bufio.NewReader(os.Stdin)
	text, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, err := core.Execute(context.Background(), text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, op := range result.SumGroup {
		fmt.Printf("Group %d: %f", i, op)
		fmt.Println()
	}

	fmt.Printf("Total calculate: %f", result.Total)
	fmt.Println()
}
