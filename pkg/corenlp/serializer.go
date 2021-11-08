package corenlp

type Serializer string

// ProtobufSerializer writes the output to a protocol buffer.
const ProtobufSerializer Serializer = "edu.stanford.nlp.pipeline.ProtobufAnnotationSerializer"

// GenericSerializer writes the output to a Java serialized object
const GenericSerializer Serializer = "edu.stanford.nlp.pipeline.GenericAnnotationSerializer"

// CustomSerializer writes the output to a (lossy!) textual representation.
const CustomSerializer Serializer = "edu.stanford.nlp.pipeline.CustomAnnotationSerializer"
