![Gopher](gopher.svg)

# Jinn

Jinn is a collection of NLP related packages for the Go ecosystem. Jinn is not NLTK. 

## Packages

| Package   | Description                | Documentation                                                               |
|-----------|----------------------------|-----------------------------------------------------------------------------|
| `corenlp` | Simple CoreNLP[^1] wrapper | [pkg/corenlp](https://pkg.go.dev/github.com/jonasknobloch/jinn/pkg/corenlp) |
| `msrpc`   | MSRPC[^2] sample iterator  | [pkg/msrpc](https://pkg.go.dev/github.com/jonasknobloch/jinn/pkg/msrpc)     |
| `paws`    | PAWS[^3] sample iterator   | [pkg/paws](https://pkg.go.dev/github.com/jonasknobloch/jinn/pkg/paws)       |
| `score`   | BLEU[^4] score             | [pkg/score](https://pkg.go.dev/github.com/jonasknobloch/jinn/pkg/score)     |
| `tree`    | Basic tree implementation  | [pkg/tree](https://pkg.go.dev/github.com/jonasknobloch/jinn/pkg/tree)       |
| `utility` | Utility functions          | [pkg/utility](https://pkg.go.dev/github.com/jonasknobloch/jinn/pkg/utility) |

[^1]: [The Stanford CoreNLP Natural Language Processing Toolkit](https://nlp.stanford.edu/pubs/StanfordCoreNlp2014.pdf)
[^2]: [Automatically Constructing a Corpus of Sentential Paraphrases](https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/I05-50025B15D.pdf)
[^3]: [PAWS: Paraphrase Adversaries from Word Scrambling](https://arxiv.org/abs/1904.01130)
[^4]: [BLEU: a Method for Automatic Evaluation of Machine Translation](https://aclanthology.org/P02-1040.pdf)