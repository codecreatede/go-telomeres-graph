package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-10-15

A kmer based approach to identify the telomeric regions and then find the upstream and the
downstream of the same. It will match the kmer to the telomers and where it is matching it
will extract the upstream and the downstream regions.


*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}

var (
	genomefasta string
	fastaid     string
	telomere    string
	upstream    int
	downstream  int
)

var rootCmd = &cobra.Command{
	Use:  "telomeres",
	Long: "This analyzes and reports the telomeres and the corresponding upstream and the downstream from the match telomeres",
	Run:  matchTelomer,
}

func init() {
	rootCmd.Flags().
		StringVarP(&genomefasta, "genome fasta", "G", "genome file in fasta format", "fasta genome")
	rootCmd.Flags().
		StringVarP(&fastaid, "fastaid", "F", "id of the fasta sequence that you want to look for the telomere", "sequence selection for teleomere")
	rootCmd.Flags().
		StringVarP(&telomere, "telomere to search for", "T", "string for the telomere", "teomere search space")
	rootCmd.Flags().
		IntVarP(&upstream, "upstream sequence", "U", 10, "upstream region for the seqence")
	rootCmd.Flags().IntVarP(&downstream, "downstream sequence", "D", 10, "downstream sequence")
}

func matchTelomer(cmd *cobra.Command, args []string) {
	type matchID struct {
		ID string
	}

	type matchSeq struct {
		seq string
	}

	mID := []matchID{}
	mSeq := []matchSeq{}

	fOpen, err := os.Open(genomefasta)
	if err != nil {
		log.Fatal(err)
	}
	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			id := strings.ReplaceAll(string(line), ">", "")
			mID = append(mID, matchID{
				ID: id,
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			mSeq = append(mSeq, matchSeq{
				seq: string(line),
			})
		}
	}

	fastaSelectedID := []string{}
	fastaSelectedSeq := []string{}

	for i := range mSeq {
		if mID[i].ID == fastaid {
			fastaSelectedID = append(fastaSelectedID, mID[i].ID)
			fastaSelectedSeq = append(fastaSelectedSeq, mSeq[i].seq)
		}
	}

	length := len(telomere)
	kmer := []string{}
	start := []int{}
	end := []int{}

	for i := 0; i <= len(fastaSelectedSeq)-1; i++ {
		for j := 0; j <= len(string(fastaSelectedSeq[i]))-length; j++ {
			kmer = append(kmer, fastaSelectedSeq[i][j:j+length])
			start = append(start, j)
			end = append(end, j+length)
		}
	}

	selectedOnes := []string{}
	selectedStart := []int{}
	selectedEnd := []int{}

	for i := range kmer {
		if kmer[i] == telomere {
			selectedOnes = append(selectedOnes, kmer[i])
			selectedStart = append(selectedStart, start[i])
			selectedEnd = append(selectedEnd, end[i])
		}
	}

	upStreamSelectedOnes := []string{}
	downStreamSelectedOnes := []string{}
	addUpDownStream := []string{}

	for i := range fastaSelectedSeq {
		for j := range selectedStart {
			addUpDownStream = append(
				addUpDownStream,
				fastaSelectedSeq[i][selectedStart[j]-upstream:selectedEnd[j]+downstream],
			)
			upStreamSelectedOnes = append(
				upStreamSelectedOnes,
				fastaSelectedSeq[i][selectedStart[j]-upstream:selectedStart[j]],
			)
			downStreamSelectedOnes = append(
				downStreamSelectedOnes,
				fastaSelectedSeq[i][selectedEnd[j]:selectedEnd[j]+downstream],
			)
		}
	}

	for i := range selectedOnes {
		fmt.Println(selectedOnes[i], selectedStart[i], selectedEnd[i])
	}
	for i := range upStreamSelectedOnes {
		fmt.Println(
			upStreamSelectedOnes[i],
			"\n",
			downStreamSelectedOnes[i],
			"\n",
			addUpDownStream[i],
		)
	}
}
