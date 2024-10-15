# go-teleomeres-graph

- searching telomeres across the plant and human genome across tree of life for reassembly of those telomere specific regions. 
- a kmer based approach to search for the matching telomere and then extract the upstream and the downstream of the matching telomere for reassembly of those regions.
- the tolkit version is present here [tolkit](https://github.com/tolkit), which only plots the kmers and not tell where to look and doesnt provide support for reassembly.
- the clade specific sets are also present in this repository, which are used in the tolkit for the search. 
- tolkit doesnt allow for the telomere extraction for re-assembly and go-telomeres-graph provide that. 

```
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-telomeres-graph ±main⚡ » \
go run main.go -h                                                               
This analyzes and reports the telomeres and the corresponding upstream and the downstream from the match telomeres

Usage:
  telomeres [flags]

Flags:
  -D, --downstream sequence int         downstream sequence (default 10)
  -F, --fastaid string                  sequence selection for teleomere (default "id of the fasta sequence that you want to look for the telomere")
  -G, --genome fasta string             fasta genome (default "genome file in fasta format")
  -h, --help                            help for telomeres
  -I, --kmersize telomere int           telomer kmder define (default 4)
  -T, --telomere to search for string   teomere search space (default "string for the telomere")
  -U, --upstream sequence int           upstream region for the seqence (default 10)
exit status 1
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-telomeres-graph ±main⚡ » 
go run main.go -G fastafile.fasta -F chr10:66478458-66505490 -T ATCGG
ATCGG 33 38
GCCCAATGAC - upstream
 CGTGTCTAGC - downstream 
 GCCCAATGACATCGGCGTGTCTAGC - upstream and downstream with the telomere included
```
Gaurav Sablok
