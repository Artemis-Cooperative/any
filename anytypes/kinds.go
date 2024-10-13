package anytypes

type Kind struct {
	Name string
	Tags []string
}

var kinds = []Kind{
	{
		Name: "bool",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"ordered",
		},
	},
	{
		Name: "string",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"ordered",
		},
	},
	{
		Name: "int",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "int8",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "int16",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "int32",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "int64",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "uint",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "uint8",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "uint16",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "uint32",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "uint64",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"integer",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "uintptr",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"ordered",
		},
	},
	{
		Name: "float32",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "float64",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"ordered",
			"numeric",
		},
	},
	{
		Name: "complex64",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"numeric",
		},
	},
	{
		Name: "complex128",
		Tags: []string{
			"basic",
			"comparable",
			"element",
			"key",
			"numeric",
		},
	},
	{
		Name: "slice",
		Tags: []string{
			"element",
		},
	},
	{
		Name: "array",
		Tags: []string{
			"comparable",
			"element",
			"key",
		},
	},
	{
		Name: "map",
		Tags: []string{
			"element",
		},
	},
	{
		Name: "struct",
		Tags: []string{
			"comparable",
			"element",
			"key",
		},
	},
	{
		Name: "chan",
		Tags: []string{
			"comparable",
			"element",
			"key",
		},
	},
	{
		Name: "func",
		Tags: []string{
			"element",
		},
	},
	{
		Name: "ptr",
		Tags: []string{
			"comparable",
			"element",
			"key",
		},
	},
}

// Return kinds for their tag.
func kindsByTag(tag string) []any {
	var filtered []any

	for _, k := range kinds {
		for _, kTag := range k.Tags {
			if kTag == tag {
				filtered = append(filtered, k.Name)
			}
		}
	}

	return filtered
}