// TestMake_03b25d77b9 generates a test suite for make()
func TestMake_03b25d77b9(t *testing.T) {
	// Test Case 1 : Passing nil as EnumValue slice 
	t.Run("Nil Enum value", func(t *testing.T) {
		ec := &executionContext{}
		res := ec.marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValueᚄ(context.Background(), gqlparser.AST{}, nil)

		if res != graphql.Null {
			t.Errorf("Expected %s but returned %s\n", graphql.Null, res)
		} 
	})

	// Test Case 2: Test with an enum value list which is not nil
	t.Run("Enum value list is not null", func(t *testing.T) {
		ec := &executionContext{}
		enum1 := introspection.EnumValue{
			Name:        "enum1",
			Description: "This is enum1",
			IsDeprecated: false,
			DeprecationReason:  "",
		}

		enum2 := introspection.EnumValue{
			Name:        "enum2",
			Description: "This is enum2",
			IsDeprecated: false,
			DeprecationReason:  "",
		}

		res := ec.marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValueᚄ(context.Background(), gqlparser.AST{}, []introspection.EnumValue{enum1, enum2})

		// Check if it's not returning null
		if res == graphql.Null {
			t.Errorf("Expected not null but returned null")
		}
	})

	// Test Case 3: Testing with a single Enum value 
	t.Run("Single Enum value", func(t *testing.T) {
		ec := &executionContext{}
		enum1 := introspection.EnumValue{
			Name:        "enum1",
			Description: "This is enum1",
			IsDeprecated: false,
			DeprecationReason:  "",
		}

		res := ec.marshalO__EnumValue2ᚕgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐEnumValueᚄ(context.Background(), gqlparser.AST{}, []introspection.EnumValue{enum1})

		// Check if it's not returning null
		if res == graphql.Null {
			t.Errorf("Expected not null but returned null")
		} 
	})
}
