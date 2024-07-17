/*
Copyright © 2024 Adam Cholewiński ITSRICHARDSHANK@GMAIL.COM
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Download a playlist from m3u8 file url",
	Run:   runUrl,
}

func init() {
	rootCmd.AddCommand(urlCmd)

	urlCmd.Flags().StringP("output", "o", "result", "Output file name without extension.")
}

func runUrl(cmd *cobra.Command, args []string) {
	red := color.New(color.Red)
	blue := color.New(color.Blue)
	blueBold := color.New(color.Blue, color.Bold)
	greenBold := color.New(color.Green, color.Bold)

	if len(args) < 1 {
		red.Println("Please enter the [url] argument.")
		return
	}

	urlArg := args[0]
	urlP, err := url.Parse(urlArg)
	if err != nil {
		red.Println("Error parsing URL")
		red.Println(err)
		return
	}

	outputName, _ := cmd.Flags().GetString("output")
	outputFile := outputName + ".ts"

	out, err := os.Create(outputFile)
	if err != nil {
		red.Println("Error creating outputfile")
		red.Println(err)
	}
	defer out.Close()

	resp, err := http.Get(urlArg)
	if err != nil {
		red.Println("Error downloading the M3U8 file. Please check the URL.")
		red.Println(err)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var segmentURLs []string
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			segmentURLs = append(segmentURLs, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading M3U8 file:", err)
		return
	}

	for i, segmentURL := range segmentURLs {
		blue.Printf("\rDownloading segment %d/%d ", i+1, len(segmentURLs))
		percent := (float64(float64(i)+1) / float64(len(segmentURLs))) * 100
		blueBold.Printf("(%.0f%%)", percent)
		wholeSegmentURL := ""
		segmentURLP, _ := url.Parse(segmentURL)
		if segmentURLP.Scheme != "" && segmentURLP.Host != "" {
			wholeSegmentURL = segmentURL
		} else {
			wholeSegmentURL = urlP.Scheme + "://" + urlP.Host + segmentURL
		}
		tempResp, err := http.Get(wholeSegmentURL)
		if err != nil {
			red.Println("\nError downloading a segment")
			red.Println(err)
			return
		}

		io.Copy(out, tempResp.Body)
		tempResp.Body.Close()
	}
	fmt.Println()

	greenBold.Printf("Video download completed successfully. File saved to %s\n", outputFile)
}
