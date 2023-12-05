package main

// Define structures for YAML parsing
type Config struct {
    SDK      string    `yaml:"sdk"`
    Resource string    `yaml:"resource"`
    Service  string    `yaml:"service"`
    Fields   FieldSets `yaml:"fields"`
}

type FieldSets struct {
    TypeString []Field `yaml:"type_string"`
    TypeBool   []Field `yaml:"type_bool"`
    TypeInt    []Field `yaml:"type_int"`
}

type Field struct {
    Name  string `yaml:"name"`
    Value string `yaml:"value"`
}