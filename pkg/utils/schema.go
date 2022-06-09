package utils

/*
This is a temporary reference of the Open Feature schema
https://github.com/open-feature/playground/blob/main/schemas/flag.schema.json
*/
func GetSchema() string {
	return `
	{
		"type": "object",
		"schemas": {
			"BooleanFlag": {
				"type": "object",
				"properties": {
					"variants": {
						"description": "variants: struct.MaxFields(2) \u0026 {",
						"type": "object",
						"required": [
							"enabled",
							"disabled"
						],
						"properties": {
							"enabled": {
								"type": "boolean",
								"enum": [
									true
								]
							},
							"disabled": {
								"type": "boolean",
								"enum": [
									false
								]
							}
						}
					},
					"defaultVariant": {
						"type": "string",
						"enum": [
							"enabled",
							"disabled"
						],
						"default": "enabled"
					}
				},
				"allOf": [
					{
						"$ref": "#/schemas/Flag"
					},
					{
						"required": [
							"variants",
							"defaultVariant"
						]
					}
				]
			},
			"Flag": {
				"type": "object",
				"required": [
					"state",
					"variants",
					"defaultVariant",
					"rules"
				],
				"properties": {
					"state": {
						"type": "string",
						"enum": [
							"enabled",
							"disabled"
						]
					},
					"variants": {
						"description": "variants: struct.MinFields(2) \u0026 {",
						"type": "object"
					},
					"defaultVariant": {
						"description": "The default variant must match a defined variant",
						"type": "string"
					},
					"rules": {
						"type": "array",
						"items": {
							"type": "object",
							"required": [
								"action",
								"condition"
							],
							"properties": {
								"action": {
									"description": "TODO: only one action can be set at a time",
									"type": "object",
									"required": [
										"variant"
									],
									"properties": {
										"variant": {
											"description": "variant: or([for f, _ in variants {f}])",
											"type": "string"
										}
									}
								},
								"condition": {
									"type": "array",
									"items": {
										"type": "object",
										"required": [
											"context",
											"op",
											"value"
										],
										"properties": {
											"context": {
												"description": "The context key that should be evaluated in this condition",
												"type": "string"
											},
											"op": {
												"description": "The operation that should be performed",
												"type": "string",
												"enum": [
													"equals",
													"starts_with",
													"ends_with"
												]
											},
											"value": {
												"description": "The value that should be evaluated\nTODO see if more values are possible",
												"type": "string"
											}
										}
									}
								}
							}
						}
					}
				}
			},
			"Flagd": {
				"type": "object",
				"properties": {
					"booleanFlags": {
						"type": "object",
						"additionalProperties": {
							"$ref": "#/schemas/BooleanFlag"
						}
					},
					"stringFlags": {
						"type": "object",
						"additionalProperties": {
							"$ref": "#/schemas/StringFlag"
						}
					},
					"numericFlags": {
						"type": "object",
						"additionalProperties": {
							"$ref": "#/schemas/NumericFlag"
						}
					},
					"objectFlags": {
						"type": "object",
						"additionalProperties": {
							"$ref": "#/schemas/ObjectFlag"
						}
					}
				}
			},
			"NumericFlag": {
				"type": "object",
				"properties": {
					"variants": {
						"description": "defaultVariant: or([for f, _ in variants {f}])",
						"type": "object",
						"additionalProperties": {
							"type": "number"
						}
					}
				},
				"allOf": [
					{
						"$ref": "#/schemas/Flag"
					},
					{
						"required": [
							"variants"
						]
					}
				]
			},
			"ObjectFlag": {
				"type": "object",
				"properties": {
					"variants": {
						"description": "defaultVariant: or([for f, _ in variants {f}])",
						"type": "object",
						"additionalProperties": {
							"description": "Any valid structs will work",
							"type": "object"
						}
					}
				},
				"allOf": [
					{
						"$ref": "#/schemas/Flag"
					},
					{
						"required": [
							"variants"
						]
					}
				]
			},
			"StringFlag": {
				"type": "object",
				"properties": {
					"variants": {
						"description": "defaultVariant: or([for f, _ in variants {f}])",
						"type": "object",
						"additionalProperties": {
							"type": "string"
						}
					}
				},
				"allOf": [
					{
						"$ref": "#/schemas/Flag"
					},
					{
						"required": [
							"variants"
						]
					}
				]
			}
		},
		"$schema": "http://json-schema.org/draft-04/schema#",
		"properties": {
			"booleanFlags": {
				"type": "object",
				"additionalProperties": {
					"$ref": "#/schemas/BooleanFlag"
				}
			},
			"stringFlags": {
				"type": "object",
				"additionalProperties": {
					"$ref": "#/schemas/StringFlag"
				}
			},
			"numericFlags": {
				"type": "object",
				"additionalProperties": {
					"$ref": "#/schemas/NumericFlag"
				}
			},
			"objectFlags": {
				"type": "object",
				"additionalProperties": {
					"$ref": "#/schemas/ObjectFlag"
				}
			}
		}
	}
	`
}
