# go-teleomeres-graph

- searching telomeres across the plant and human genome across tree of life for reassembly of those telomere specific regions. 
- a kmer based approach to search for the matching telomere and then extract the upstream and the downstream of the matching telomere for reassembly of those regions.
- the tolkit version is present here [tolkit](https://github.com/tolkit/telomeric-identifier), which only plots the kmers and not tell where to look and doesnt provide support for reassembly.
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

```
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-telomeres-graph ±main⚡ » 
./telomeres -h                                                                  
This analyzes and reports the telomeres and the corresponding upstream and the downstream from the match telomeres

Usage:
  telomeres [flags]

Flags:
  -D, --downstream sequence int         downstream sequence (default 10)
  -F, --fastaid string                  sequence selection for teleomere (default "id of the fasta sequence that you want to look for the telomere")
  -G, --genome fasta string             fasta genome (default "genome file in fasta format")
  -h, --help                            help for telomeres
  -T, --telomere to search for string   teomere search space (default "string for the telomere")
  -U, --upstream sequence int           upstream region for the seqence (default 10)
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-telomeres-graph ±main⚡ » 
./telomeres -G fastafile.fasta -F chr10:66478458-66505490 -T ATCGG
GCCCAATGAC
 CGTGTCTAGC
 GCCCAATGACATCGGCGTGTCTAGC
```
```
wxr-xr-x. 1 gauavsablok gauavsablok      40 Oct 14 14:58 clades
-rw-r--r--. 1 gauavsablok gauavsablok      13 Oct 15 13:51 downstream.fasta
-rwxr-xr-x. 1 gauavsablok gauavsablok     148 Oct 15 11:38 fastafile.fasta
drwxr-xr-x. 1 gauavsablok gauavsablok     164 Oct 15 13:51 .git
-rw-r--r--. 1 gauavsablok gauavsablok     204 Oct 15 09:49 go.mod
-rw-r--r--. 1 gauavsablok gauavsablok     896 Oct 15 09:49 go.sum
-rw-r--r--. 1 gauavsablok gauavsablok    4658 Oct 15 13:49 main.go
-rw-r--r--. 1 gauavsablok gauavsablok    3081 Oct 15 13:53 README.md
-rw-r--r--. 1 gauavsablok gauavsablok      28 Oct 15 13:51 reassemble.fasta
-rwxr-xr-x. 1 gauavsablok gauavsablok 5480185 Oct 15 13:51 telomeres
-rw-r--r--. 1 gauavsablok gauavsablok 5488640 Oct 15 13:52 telomeres.tar
-rw-r--r--. 1 gauavsablok gauavsablok     161 Oct 15 13:51 telomeres.txt
-rw-r--r--. 1 gauavsablok gauavsablok      13 Oct 15 13:51 upstream.fasta
```
Gaurav Sablok
