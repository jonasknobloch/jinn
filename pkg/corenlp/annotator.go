package corenlp

type Annotator string

// TokenizerAnnotator tokenizes the text.
const TokenizerAnnotator Annotator = "tokenizer"

// CleanXmlAnnotator removes XML tokens from the document.
const CleanXmlAnnotator Annotator = "cleanxml"

// DocDateAnnotator allows specifying dates for documents.
const DocDateAnnotator Annotator = "docdate"

// WordsToSentencesAnnotator splits a sequence of tokens into sentences.
const WordsToSentencesAnnotator Annotator = "ssplit"

// POSTaggerAnnotator labels tokens with their POS tag.
const POSTaggerAnnotator Annotator = "pos"

// MorphaAnnotator generates the word lemmas for all tokens in the corpus.
const MorphaAnnotator Annotator = "lemma"

// NERCombinerAnnotator recognizes named (PERSON, LOCATION, ORGANIZATION, MISC), numerical (MONEY, NUMBER, ORDINAL,
// PERCENT), and temporal (DATE, TIME, DURATION, SET) entities.
const NERCombinerAnnotator Annotator = "ner"

// EntityMentionsAnnotator groups NER tagged tokens together into mentions.
const EntityMentionsAnnotator Annotator = "entitymentions"

// TokensRegexNERAnnotator implements a simple, rule-based NER over token sequences using Java regular expressions.
const TokensRegexNERAnnotator Annotator = "regexner"

// TokensRegexAnnotator runs a TokensRegex pipeline within a full NLP pipeline.
const TokensRegexAnnotator Annotator = "tokensregex"

// ParserAnnotator provides full syntactic analysis, using both the constituent and the dependency representations.
const ParserAnnotator Annotator = "parse"

// DependencyParseAnnotator provides a fast syntactic dependency parser.
const DependencyParseAnnotator Annotator = "depparse"

// CorefAnnotator performs coreference resolution on a document, building links between entity mentions that refer to
// the same entity.
const CorefAnnotator Annotator = "coref"

// DeterministicCorefAnnotator implements both pronominal and nominal coreference resolution.
const DeterministicCorefAnnotator Annotator = "dcoref"

// RelationExtractorAnnotator finds relations between two entities.
const RelationExtractorAnnotator Annotator = "relation"

// NaturalLogicAnnotator marks quantifier scope and token polarity, according to natural logic semantics.
const NaturalLogicAnnotator Annotator = "natlog"

// OpenIEAnnotator extracts open-domain relation triples.
const OpenIEAnnotator Annotator = "openie"

// WikidictAnnotator links entity mentions to Wikipedia entities
const WikidictAnnotator Annotator = "entitylink"

// KBPAnnotator extracts (subject, relation, object) triples from sentences, using a combination of a statistical model,
// patterns over tokens, and patterns over dependencies.
const KBPAnnotator Annotator = "kbp"

// QuoteAnnotator deterministically picks out quotes delimited by “ or ‘ from a text.
const QuoteAnnotator Annotator = "quote"

// QuoteAttributionAnnotator attributes quotes to speakers in the document.
const QuoteAttributionAnnotator Annotator = "quote.attribution"

// SentimentAnnotator implements Socher et al’s sentiment model.
const SentimentAnnotator Annotator = "sentiment"

// TrueCaseAnnotator recognizes the true case of tokens in text where this information was lost, e.g., all upper case
// text.
const TrueCaseAnnotator Annotator = "truecase"

// UDFeatureAnnotator labels tokens with their Universal Dependencies universal part of speech (UPOS) and features.
const UDFeatureAnnotator Annotator = "udfeats"
