package casing

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_ModelTitle_Ollama(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "nemotron-3-nano:latest", expected: "Nemotron 3 Nano (latest)"},
		{input: "nemotron-3-nano", expected: "Nemotron 3 Nano"},
		{input: "nemotron-3-nano:30b", expected: "Nemotron 3 Nano 30B"},
		{input: "nemotron-3-nano:30b-a3b-q4_K_M", expected: "Nemotron 3 Nano 30B A3B Q4_K_M"},
		{input: "nemotron-3-nano:30b-a3b-q8_0", expected: "Nemotron 3 Nano 30B A3B Q8_0"},
		{input: "nemotron-3-nano:30b-a3b-fp16", expected: "Nemotron 3 Nano 30B A3B FP16"},
		{input: "nemotron-3-nano:30b-cloud", expected: "Nemotron 3 Nano 30B Cloud"},
		{input: "functiongemma:latest", expected: "Functiongemma (latest)"},
		{input: "functiongemma", expected: "Functiongemma"},
		{input: "functiongemma:270m", expected: "Functiongemma 270M"},
		{input: "functiongemma:270m-it-q8_0", expected: "Functiongemma 270M Instruct Q8_0"},
		{input: "functiongemma:270m-it-fp16", expected: "Functiongemma 270M Instruct FP16"},
		{input: "olmo-3:latest", expected: "Olmo 3 (latest)"},
		{input: "olmo-3", expected: "Olmo 3"},
		{input: "olmo-3:7b", expected: "Olmo 3 7B"},
		{input: "olmo-3:32b", expected: "Olmo 3 32B"},
		{input: "olmo-3:7b-instruct", expected: "Olmo 3 7B Instruct"},
		{input: "olmo-3:7b-instruct-q4_K_M", expected: "Olmo 3 7B Instruct Q4_K_M"},
		{input: "olmo-3:7b-instruct-q8_0", expected: "Olmo 3 7B Instruct Q8_0"},
		{input: "olmo-3:7b-instruct-fp16", expected: "Olmo 3 7B Instruct FP16"},
		{input: "olmo-3:7b-think", expected: "Olmo 3 7B Think"},
		{input: "olmo-3:7b-think-q4_K_M", expected: "Olmo 3 7B Think Q4_K_M"},
		{input: "olmo-3:7b-think-q8_0", expected: "Olmo 3 7B Think Q8_0"},
		{input: "olmo-3:7b-think-fp16", expected: "Olmo 3 7B Think FP16"},
		{input: "olmo-3:32b-think", expected: "Olmo 3 32B Think"},
		{input: "olmo-3:32b-think-q4_K_M", expected: "Olmo 3 32B Think Q4_K_M"},
		{input: "olmo-3:32b-think-q8_0", expected: "Olmo 3 32B Think Q8_0"},
		{input: "olmo-3:32b-think-fp16", expected: "Olmo 3 32B Think FP16"},
		{input: "gemini-3-flash-preview:latest", expected: "Gemini 3 Flash Preview (latest)"},
		{input: "gemini-3-flash-preview", expected: "Gemini 3 Flash Preview"},
		{input: "gemini-3-flash-preview:cloud", expected: "Gemini 3 Flash Preview Cloud"},
		{input: "devstral-small-2:latest", expected: "Devstral Small 2 (latest)"},
		{input: "devstral-small-2", expected: "Devstral Small 2"},
		{input: "devstral-small-2:24b", expected: "Devstral Small 2 24B"},
		{input: "devstral-small-2:24b-cloud", expected: "Devstral Small 2 24B Cloud"},
		{input: "devstral-small-2:24b-instruct-2512-q4_K_M", expected: "Devstral Small 2 24B Instruct 2512 Q4_K_M"},
		{input: "devstral-small-2:24b-instruct-2512-q8_0", expected: "Devstral Small 2 24B Instruct 2512 Q8_0"},
		{input: "devstral-small-2:24b-instruct-2512-fp16", expected: "Devstral Small 2 24B Instruct 2512 FP16"},
		{input: "devstral-2:latest", expected: "Devstral 2 (latest)"},
		{input: "devstral-2", expected: "Devstral 2"},
		{input: "devstral-2:123b", expected: "Devstral 2 123B"},
		{input: "devstral-2:123b-cloud", expected: "Devstral 2 123B Cloud"},
		{input: "devstral-2:123b-instruct-2512-q4_K_M", expected: "Devstral 2 123B Instruct 2512 Q4_K_M"},
		{input: "devstral-2:123b-instruct-2512-q8_0", expected: "Devstral 2 123B Instruct 2512 Q8_0"},
		{input: "devstral-2:123b-instruct-2512-fp16", expected: "Devstral 2 123B Instruct 2512 FP16"},
		{input: "ministral-3:latest", expected: "Ministral 3 (latest)"},
		{input: "ministral-3", expected: "Ministral 3"},
		{input: "ministral-3:3b", expected: "Ministral 3 3B"},
		{input: "ministral-3:8b", expected: "Ministral 3 8B"},
		{input: "ministral-3:14b", expected: "Ministral 3 14B"},
		{input: "ministral-3:3b-cloud", expected: "Ministral 3 3B Cloud"},
		{input: "ministral-3:3b-instruct-2512-q4_K_M", expected: "Ministral 3 3B Instruct 2512 Q4_K_M"},
		{input: "ministral-3:3b-instruct-2512-q8_0", expected: "Ministral 3 3B Instruct 2512 Q8_0"},
		{input: "ministral-3:3b-instruct-2512-fp16", expected: "Ministral 3 3B Instruct 2512 FP16"},
		{input: "ministral-3:8b-cloud", expected: "Ministral 3 8B Cloud"},
		{input: "ministral-3:8b-instruct-2512-q4_K_M", expected: "Ministral 3 8B Instruct 2512 Q4_K_M"},
		{input: "ministral-3:8b-instruct-2512-q8_0", expected: "Ministral 3 8B Instruct 2512 Q8_0"},
		{input: "ministral-3:8b-instruct-2512-fp16", expected: "Ministral 3 8B Instruct 2512 FP16"},
		{input: "ministral-3:14b-cloud", expected: "Ministral 3 14B Cloud"},
		{input: "ministral-3:14b-instruct-2512-q4_K_M", expected: "Ministral 3 14B Instruct 2512 Q4_K_M"},
		{input: "ministral-3:14b-instruct-2512-q8_0", expected: "Ministral 3 14B Instruct 2512 Q8_0"},
		{input: "ministral-3:14b-instruct-2512-fp16", expected: "Ministral 3 14B Instruct 2512 FP16"},
		{input: "qwen3-vl:latest", expected: "Qwen3 Vl (latest)"},
		{input: "qwen3-vl", expected: "Qwen3 Vl"},
		{input: "qwen3-vl:2b", expected: "Qwen3 Vl 2B"},
		{input: "qwen3-vl:4b", expected: "Qwen3 Vl 4B"},
		{input: "qwen3-vl:8b", expected: "Qwen3 Vl 8B"},
		{input: "qwen3-vl:30b", expected: "Qwen3 Vl 30B"},
		{input: "qwen3-vl:32b", expected: "Qwen3 Vl 32B"},
		{input: "qwen3-vl:235b", expected: "Qwen3 Vl 235B"},
		{input: "qwen3-vl:2b-instruct", expected: "Qwen3 Vl 2B Instruct"},
		{input: "qwen3-vl:2b-instruct-q4_K_M", expected: "Qwen3 Vl 2B Instruct Q4_K_M"},
		{input: "qwen3-vl:2b-instruct-q8_0", expected: "Qwen3 Vl 2B Instruct Q8_0"},
		{input: "qwen3-vl:2b-instruct-bf16", expected: "Qwen3 Vl 2B Instruct BF16"},
		{input: "qwen3-vl:2b-thinking", expected: "Qwen3 Vl 2B Thinking"},
		{input: "qwen3-vl:2b-thinking-q4_K_M", expected: "Qwen3 Vl 2B Thinking Q4_K_M"},
		{input: "qwen3-vl:2b-thinking-q8_0", expected: "Qwen3 Vl 2B Thinking Q8_0"},
		{input: "qwen3-vl:2b-thinking-bf16", expected: "Qwen3 Vl 2B Thinking BF16"},
		{input: "qwen3-vl:4b-instruct", expected: "Qwen3 Vl 4B Instruct"},
		{input: "qwen3-vl:4b-instruct-q4_K_M", expected: "Qwen3 Vl 4B Instruct Q4_K_M"},
		{input: "qwen3-vl:4b-instruct-q8_0", expected: "Qwen3 Vl 4B Instruct Q8_0"},
		{input: "qwen3-vl:4b-instruct-bf16", expected: "Qwen3 Vl 4B Instruct BF16"},
		{input: "qwen3-vl:4b-thinking", expected: "Qwen3 Vl 4B Thinking"},
		{input: "qwen3-vl:4b-thinking-q4_K_M", expected: "Qwen3 Vl 4B Thinking Q4_K_M"},
		{input: "qwen3-vl:4b-thinking-q8_0", expected: "Qwen3 Vl 4B Thinking Q8_0"},
		{input: "qwen3-vl:4b-thinking-bf16", expected: "Qwen3 Vl 4B Thinking BF16"},
		{input: "qwen3-vl:8b-instruct", expected: "Qwen3 Vl 8B Instruct"},
		{input: "qwen3-vl:8b-instruct-q4_K_M", expected: "Qwen3 Vl 8B Instruct Q4_K_M"},
		{input: "qwen3-vl:8b-instruct-q8_0", expected: "Qwen3 Vl 8B Instruct Q8_0"},
		{input: "qwen3-vl:8b-instruct-bf16", expected: "Qwen3 Vl 8B Instruct BF16"},
		{input: "qwen3-vl:8b-thinking", expected: "Qwen3 Vl 8B Thinking"},
		{input: "qwen3-vl:8b-thinking-q4_K_M", expected: "Qwen3 Vl 8B Thinking Q4_K_M"},
		{input: "qwen3-vl:8b-thinking-q8_0", expected: "Qwen3 Vl 8B Thinking Q8_0"},
		{input: "qwen3-vl:8b-thinking-bf16", expected: "Qwen3 Vl 8B Thinking BF16"},
		{input: "qwen3-vl:30b-a3b", expected: "Qwen3 Vl 30B A3B"},
		{input: "qwen3-vl:30b-a3b-instruct", expected: "Qwen3 Vl 30B A3B Instruct"},
		{input: "qwen3-vl:30b-a3b-instruct-q4_K_M", expected: "Qwen3 Vl 30B A3B Instruct Q4_K_M"},
		{input: "qwen3-vl:30b-a3b-instruct-q8_0", expected: "Qwen3 Vl 30B A3B Instruct Q8_0"},
		{input: "qwen3-vl:30b-a3b-instruct-bf16", expected: "Qwen3 Vl 30B A3B Instruct BF16"},
		{input: "qwen3-vl:30b-a3b-thinking", expected: "Qwen3 Vl 30B A3B Thinking"},
		{input: "qwen3-vl:30b-a3b-thinking-q4_K_M", expected: "Qwen3 Vl 30B A3B Thinking Q4_K_M"},
		{input: "qwen3-vl:30b-a3b-thinking-q8_0", expected: "Qwen3 Vl 30B A3B Thinking Q8_0"},
		{input: "qwen3-vl:30b-a3b-thinking-bf16", expected: "Qwen3 Vl 30B A3B Thinking BF16"},
		{input: "qwen3-vl:32b-instruct", expected: "Qwen3 Vl 32B Instruct"},
		{input: "qwen3-vl:32b-instruct-q4_K_M", expected: "Qwen3 Vl 32B Instruct Q4_K_M"},
		{input: "qwen3-vl:32b-instruct-q8_0", expected: "Qwen3 Vl 32B Instruct Q8_0"},
		{input: "qwen3-vl:32b-instruct-bf16", expected: "Qwen3 Vl 32B Instruct BF16"},
		{input: "qwen3-vl:32b-thinking", expected: "Qwen3 Vl 32B Thinking"},
		{input: "qwen3-vl:32b-thinking-q4_K_M", expected: "Qwen3 Vl 32B Thinking Q4_K_M"},
		{input: "qwen3-vl:32b-thinking-q8_0", expected: "Qwen3 Vl 32B Thinking Q8_0"},
		{input: "qwen3-vl:32b-thinking-bf16", expected: "Qwen3 Vl 32B Thinking BF16"},
		{input: "qwen3-vl:235b-a22b", expected: "Qwen3 Vl 235B A22B"},
		{input: "qwen3-vl:235b-a22b-instruct", expected: "Qwen3 Vl 235B A22B Instruct"},
		{input: "qwen3-vl:235b-a22b-instruct-q4_K_M", expected: "Qwen3 Vl 235B A22B Instruct Q4_K_M"},
		{input: "qwen3-vl:235b-a22b-instruct-q8_0", expected: "Qwen3 Vl 235B A22B Instruct Q8_0"},
		{input: "qwen3-vl:235b-a22b-instruct-bf16", expected: "Qwen3 Vl 235B A22B Instruct BF16"},
		{input: "qwen3-vl:235b-a22b-thinking", expected: "Qwen3 Vl 235B A22B Thinking"},
		{input: "qwen3-vl:235b-a22b-thinking-q4_K_M", expected: "Qwen3 Vl 235B A22B Thinking Q4_K_M"},
		{input: "qwen3-vl:235b-a22b-thinking-q8_0", expected: "Qwen3 Vl 235B A22B Thinking Q8_0"},
		{input: "qwen3-vl:235b-a22b-thinking-bf16", expected: "Qwen3 Vl 235B A22B Thinking BF16"},
		{input: "qwen3-vl:235b-cloud", expected: "Qwen3 Vl 235B Cloud"},
		{input: "qwen3-vl:235b-instruct-cloud", expected: "Qwen3 Vl 235B Instruct Cloud"},
		{input: "gpt-oss:latest", expected: "GPT OSS (latest)"},
		{input: "gpt-oss", expected: "GPT OSS"},
		{input: "gpt-oss:20b", expected: "GPT OSS 20B"},
		{input: "gpt-oss:120b", expected: "GPT OSS 120B"},
		{input: "gpt-oss:20b-cloud", expected: "GPT OSS 20B Cloud"},
		{input: "gpt-oss:120b-cloud", expected: "GPT OSS 120B Cloud"},
		{input: "deepseek-r1:latest", expected: "Deepseek R1 (latest)"},
		{input: "deepseek-r1", expected: "Deepseek R1"},
		{input: "deepseek-r1:1.5b", expected: "Deepseek R1 1.5B"},
		{input: "deepseek-r1:7b", expected: "Deepseek R1 7B"},
		{input: "deepseek-r1:8b", expected: "Deepseek R1 8B"},
		{input: "deepseek-r1:14b", expected: "Deepseek R1 14B"},
		{input: "deepseek-r1:32b", expected: "Deepseek R1 32B"},
		{input: "deepseek-r1:70b", expected: "Deepseek R1 70B"},
		{input: "deepseek-r1:671b", expected: "Deepseek R1 671B"},
		{input: "deepseek-r1:1.5b-qwen-distill-q4_K_M", expected: "Deepseek R1 1.5B Qwen Distill Q4_K_M"},
		{input: "deepseek-r1:1.5b-qwen-distill-q8_0", expected: "Deepseek R1 1.5B Qwen Distill Q8_0"},
		{input: "deepseek-r1:1.5b-qwen-distill-fp16", expected: "Deepseek R1 1.5B Qwen Distill FP16"},
		{input: "deepseek-r1:7b-qwen-distill-q4_K_M", expected: "Deepseek R1 7B Qwen Distill Q4_K_M"},
		{input: "deepseek-r1:7b-qwen-distill-q8_0", expected: "Deepseek R1 7B Qwen Distill Q8_0"},
		{input: "deepseek-r1:7b-qwen-distill-fp16", expected: "Deepseek R1 7B Qwen Distill FP16"},
		{input: "deepseek-r1:8b-0528-qwen3-q4_K_M", expected: "Deepseek R1 8B 0528 Qwen3 Q4_K_M"},
		{input: "deepseek-r1:8b-0528-qwen3-q8_0", expected: "Deepseek R1 8B 0528 Qwen3 Q8_0"},
		{input: "deepseek-r1:8b-0528-qwen3-fp16", expected: "Deepseek R1 8B 0528 Qwen3 FP16"},
		{input: "deepseek-r1:8b-llama-distill-q4_K_M", expected: "Deepseek R1 8B Llama Distill Q4_K_M"},
		{input: "deepseek-r1:8b-llama-distill-q8_0", expected: "Deepseek R1 8B Llama Distill Q8_0"},
		{input: "deepseek-r1:8b-llama-distill-fp16", expected: "Deepseek R1 8B Llama Distill FP16"},
		{input: "deepseek-r1:14b-qwen-distill-q4_K_M", expected: "Deepseek R1 14B Qwen Distill Q4_K_M"},
		{input: "deepseek-r1:14b-qwen-distill-q8_0", expected: "Deepseek R1 14B Qwen Distill Q8_0"},
		{input: "deepseek-r1:14b-qwen-distill-fp16", expected: "Deepseek R1 14B Qwen Distill FP16"},
		{input: "deepseek-r1:32b-qwen-distill-q4_K_M", expected: "Deepseek R1 32B Qwen Distill Q4_K_M"},
		{input: "deepseek-r1:32b-qwen-distill-q8_0", expected: "Deepseek R1 32B Qwen Distill Q8_0"},
		{input: "deepseek-r1:32b-qwen-distill-fp16", expected: "Deepseek R1 32B Qwen Distill FP16"},
		{input: "deepseek-r1:70b-llama-distill-q4_K_M", expected: "Deepseek R1 70B Llama Distill Q4_K_M"},
		{input: "deepseek-r1:70b-llama-distill-q8_0", expected: "Deepseek R1 70B Llama Distill Q8_0"},
		{input: "deepseek-r1:70b-llama-distill-fp16", expected: "Deepseek R1 70B Llama Distill FP16"},
		{input: "deepseek-r1:671b-0528-q4_K_M", expected: "Deepseek R1 671B 0528 Q4_K_M"},
		{input: "deepseek-r1:671b-0528-q8_0", expected: "Deepseek R1 671B 0528 Q8_0"},
		{input: "deepseek-r1:671b-0528-fp16", expected: "Deepseek R1 671B 0528 FP16"},
		{input: "deepseek-r1:671b-q4_K_M", expected: "Deepseek R1 671B Q4_K_M"},
		{input: "deepseek-r1:671b-q8_0", expected: "Deepseek R1 671B Q8_0"},
		{input: "deepseek-r1:671b-fp16", expected: "Deepseek R1 671B FP16"},
		{input: "qwen3-coder:latest", expected: "Qwen3 Coder (latest)"},
		{input: "qwen3-coder", expected: "Qwen3 Coder"},
		{input: "qwen3-coder:30b", expected: "Qwen3 Coder 30B"},
		{input: "qwen3-coder:480b", expected: "Qwen3 Coder 480B"},
		{input: "qwen3-coder:30b-a3b-q4_K_M", expected: "Qwen3 Coder 30B A3B Q4_K_M"},
		{input: "qwen3-coder:30b-a3b-q8_0", expected: "Qwen3 Coder 30B A3B Q8_0"},
		{input: "qwen3-coder:30b-a3b-fp16", expected: "Qwen3 Coder 30B A3B FP16"},
		{input: "qwen3-coder:480b-a35b-q4_K_M", expected: "Qwen3 Coder 480B A35B Q4_K_M"},
		{input: "qwen3-coder:480b-a35b-q8_0", expected: "Qwen3 Coder 480B A35B Q8_0"},
		{input: "qwen3-coder:480b-a35b-fp16", expected: "Qwen3 Coder 480B A35B FP16"},
		{input: "qwen3-coder:480b-cloud", expected: "Qwen3 Coder 480B Cloud"},
		{input: "gemma3:latest", expected: "Gemma3 (latest)"},
		{input: "gemma3", expected: "Gemma3"},
		{input: "gemma3:270m", expected: "Gemma3 270M"},
		{input: "gemma3:1b", expected: "Gemma3 1B"},
		{input: "gemma3:4b", expected: "Gemma3 4B"},
		{input: "gemma3:12b", expected: "Gemma3 12B"},
		{input: "gemma3:27b", expected: "Gemma3 27B"},
		{input: "gemma3:270m-it-qat", expected: "Gemma3 270M Instruct Qat"},
		{input: "gemma3:270m-it-q8_0", expected: "Gemma3 270M Instruct Q8_0"},
		{input: "gemma3:270m-it-fp16", expected: "Gemma3 270M Instruct FP16"},
		{input: "gemma3:270m-it-bf16", expected: "Gemma3 270M Instruct BF16"},
		{input: "gemma3:1b-it-qat", expected: "Gemma3 1B Instruct Qat"},
		{input: "gemma3:1b-it-q4_K_M", expected: "Gemma3 1B Instruct Q4_K_M"},
		{input: "gemma3:1b-it-q8_0", expected: "Gemma3 1B Instruct Q8_0"},
		{input: "gemma3:1b-it-fp16", expected: "Gemma3 1B Instruct FP16"},
		{input: "gemma3:4b-cloud", expected: "Gemma3 4B Cloud"},
		{input: "gemma3:4b-it-qat", expected: "Gemma3 4B Instruct Qat"},
		{input: "gemma3:4b-it-q4_K_M", expected: "Gemma3 4B Instruct Q4_K_M"},
		{input: "gemma3:4b-it-q8_0", expected: "Gemma3 4B Instruct Q8_0"},
		{input: "gemma3:4b-it-fp16", expected: "Gemma3 4B Instruct FP16"},
		{input: "gemma3:12b-cloud", expected: "Gemma3 12B Cloud"},
		{input: "gemma3:12b-it-qat", expected: "Gemma3 12B Instruct Qat"},
		{input: "gemma3:12b-it-q4_K_M", expected: "Gemma3 12B Instruct Q4_K_M"},
		{input: "gemma3:12b-it-q8_0", expected: "Gemma3 12B Instruct Q8_0"},
		{input: "gemma3:12b-it-fp16", expected: "Gemma3 12B Instruct FP16"},
		{input: "gemma3:27b-cloud", expected: "Gemma3 27B Cloud"},
		{input: "gemma3:27b-it-qat", expected: "Gemma3 27B Instruct Qat"},
		{input: "gemma3:27b-it-q4_K_M", expected: "Gemma3 27B Instruct Q4_K_M"},
		{input: "gemma3:27b-it-q8_0", expected: "Gemma3 27B Instruct Q8_0"},
		{input: "gemma3:27b-it-fp16", expected: "Gemma3 27B Instruct FP16"},
		{input: "llama3.1:latest", expected: "Llama3.1 (latest)"},
		{input: "llama3.1", expected: "Llama3.1"},
		{input: "llama3.1:8b", expected: "Llama3.1 8B"},
		{input: "llama3.1:70b", expected: "Llama3.1 70B"},
		{input: "llama3.1:405b", expected: "Llama3.1 405B"},
		{input: "llama3.1:8b-instruct-q2_K", expected: "Llama3.1 8B Instruct Q2_K"},
		{input: "llama3.1:8b-instruct-q3_K_S", expected: "Llama3.1 8B Instruct Q3 K_S"},
		{input: "llama3.1:8b-instruct-q3_K_M", expected: "Llama3.1 8B Instruct Q3 K_M"},
		{input: "llama3.1:8b-instruct-q3_K_L", expected: "Llama3.1 8B Instruct Q3 K_L"},
		{input: "llama3.1:8b-instruct-q4_0", expected: "Llama3.1 8B Instruct Q4_0"},
		{input: "llama3.1:8b-instruct-q4_1", expected: "Llama3.1 8B Instruct Q4 1"},
		{input: "llama3.1:8b-instruct-q4_K_S", expected: "Llama3.1 8B Instruct Q4 K_S"},
		{input: "llama3.1:8b-instruct-q4_K_M", expected: "Llama3.1 8B Instruct Q4_K_M"},
		{input: "llama3.1:8b-instruct-q5_0", expected: "Llama3.1 8B Instruct Q5 0"},
		{input: "llama3.1:8b-instruct-q5_1", expected: "Llama3.1 8B Instruct Q5 1"},
		{input: "llama3.1:8b-instruct-q5_K_S", expected: "Llama3.1 8B Instruct Q5 K_S"},
		{input: "llama3.1:8b-instruct-q5_K_M", expected: "Llama3.1 8B Instruct Q5 K_M"},
		{input: "llama3.1:8b-instruct-q6_K", expected: "Llama3.1 8B Instruct Q6 K"},
		{input: "llama3.1:8b-instruct-q8_0", expected: "Llama3.1 8B Instruct Q8_0"},
		{input: "llama3.1:8b-instruct-fp16", expected: "Llama3.1 8B Instruct FP16"},
		{input: "llama3.1:8b-text-q2_K", expected: "Llama3.1 8B Text Q2_K"},
		{input: "llama3.1:8b-text-q3_K_S", expected: "Llama3.1 8B Text Q3 K_S"},
		{input: "llama3.1:8b-text-q3_K_M", expected: "Llama3.1 8B Text Q3 K_M"},
		{input: "llama3.1:8b-text-q3_K_L", expected: "Llama3.1 8B Text Q3 K_L"},
		{input: "llama3.1:8b-text-q4_0", expected: "Llama3.1 8B Text Q4_0"},
		{input: "llama3.1:8b-text-q4_1", expected: "Llama3.1 8B Text Q4 1"},
		{input: "llama3.1:8b-text-q4_K_S", expected: "Llama3.1 8B Text Q4 K_S"},
		{input: "llama3.1:8b-text-q4_K_M", expected: "Llama3.1 8B Text Q4_K_M"},
		{input: "llama3.1:8b-text-q5_0", expected: "Llama3.1 8B Text Q5 0"},
		{input: "llama3.1:8b-text-q5_1", expected: "Llama3.1 8B Text Q5 1"},
		{input: "llama3.1:8b-text-q5_K_S", expected: "Llama3.1 8B Text Q5 K_S"},
		{input: "llama3.1:8b-text-q5_K_M", expected: "Llama3.1 8B Text Q5 K_M"},
		{input: "llama3.1:8b-text-q6_K", expected: "Llama3.1 8B Text Q6 K"},
		{input: "llama3.1:8b-text-q8_0", expected: "Llama3.1 8B Text Q8_0"},
		{input: "llama3.1:8b-text-fp16", expected: "Llama3.1 8B Text FP16"},
		{input: "llama3.1:70b-instruct-q2_K", expected: "Llama3.1 70B Instruct Q2_K"},
		{input: "llama3.1:70b-instruct-q3_K_S", expected: "Llama3.1 70B Instruct Q3 K_S"},
		{input: "llama3.1:70b-instruct-q3_K_M", expected: "Llama3.1 70B Instruct Q3 K_M"},
		{input: "llama3.1:70b-instruct-q3_K_L", expected: "Llama3.1 70B Instruct Q3 K_L"},
		{input: "llama3.1:70b-instruct-q4_0", expected: "Llama3.1 70B Instruct Q4_0"},
		{input: "llama3.1:70b-instruct-q4_K_S", expected: "Llama3.1 70B Instruct Q4 K_S"},
		{input: "llama3.1:70b-instruct-q4_K_M", expected: "Llama3.1 70B Instruct Q4_K_M"},
		{input: "llama3.1:70b-instruct-q5_0", expected: "Llama3.1 70B Instruct Q5 0"},
		{input: "llama3.1:70b-instruct-q5_1", expected: "Llama3.1 70B Instruct Q5 1"},
		{input: "llama3.1:70b-instruct-q5_K_S", expected: "Llama3.1 70B Instruct Q5 K_S"},
		{input: "llama3.1:70b-instruct-q5_K_M", expected: "Llama3.1 70B Instruct Q5 K_M"},
		{input: "llama3.1:70b-instruct-q6_K", expected: "Llama3.1 70B Instruct Q6 K"},
		{input: "llama3.1:70b-instruct-q8_0", expected: "Llama3.1 70B Instruct Q8_0"},
		{input: "llama3.1:70b-instruct-fp16", expected: "Llama3.1 70B Instruct FP16"},
		{input: "llama3.1:70b-text-q2_K", expected: "Llama3.1 70B Text Q2_K"},
		{input: "llama3.1:70b-text-q3_K_S", expected: "Llama3.1 70B Text Q3 K_S"},
		{input: "llama3.1:70b-text-q3_K_M", expected: "Llama3.1 70B Text Q3 K_M"},
		{input: "llama3.1:70b-text-q3_K_L", expected: "Llama3.1 70B Text Q3 K_L"},
		{input: "llama3.1:70b-text-q4_0", expected: "Llama3.1 70B Text Q4_0"},
		{input: "llama3.1:70b-text-q4_1", expected: "Llama3.1 70B Text Q4 1"},
		{input: "llama3.1:70b-text-q4_K_S", expected: "Llama3.1 70B Text Q4 K_S"},
		{input: "llama3.1:70b-text-q4_K_M", expected: "Llama3.1 70B Text Q4_K_M"},
		{input: "llama3.1:70b-text-q5_0", expected: "Llama3.1 70B Text Q5 0"},
		{input: "llama3.1:70b-text-q5_1", expected: "Llama3.1 70B Text Q5 1"},
		{input: "llama3.1:70b-text-q5_K_S", expected: "Llama3.1 70B Text Q5 K_S"},
		{input: "llama3.1:70b-text-q5_K_M", expected: "Llama3.1 70B Text Q5 K_M"},
		{input: "llama3.1:70b-text-q6_K", expected: "Llama3.1 70B Text Q6 K"},
		{input: "llama3.1:70b-text-q8_0", expected: "Llama3.1 70B Text Q8_0"},
		{input: "llama3.1:70b-text-fp16", expected: "Llama3.1 70B Text FP16"},
		{input: "llama3.1:405b-instruct-q2_K", expected: "Llama3.1 405B Instruct Q2_K"},
		{input: "llama3.1:405b-instruct-q3_K_S", expected: "Llama3.1 405B Instruct Q3 K_S"},
		{input: "llama3.1:405b-instruct-q3_K_M", expected: "Llama3.1 405B Instruct Q3 K_M"},
		{input: "llama3.1:405b-instruct-q3_K_L", expected: "Llama3.1 405B Instruct Q3 K_L"},
		{input: "llama3.1:405b-instruct-q4_0", expected: "Llama3.1 405B Instruct Q4_0"},
		{input: "llama3.1:405b-instruct-q4_1", expected: "Llama3.1 405B Instruct Q4 1"},
		{input: "llama3.1:405b-instruct-q4_K_S", expected: "Llama3.1 405B Instruct Q4 K_S"},
		{input: "llama3.1:405b-instruct-q4_K_M", expected: "Llama3.1 405B Instruct Q4_K_M"},
		{input: "llama3.1:405b-instruct-q5_0", expected: "Llama3.1 405B Instruct Q5 0"},
		{input: "llama3.1:405b-instruct-q5_1", expected: "Llama3.1 405B Instruct Q5 1"},
		{input: "llama3.1:405b-instruct-q5_K_S", expected: "Llama3.1 405B Instruct Q5 K_S"},
		{input: "llama3.1:405b-instruct-q5_K_M", expected: "Llama3.1 405B Instruct Q5 K_M"},
		{input: "llama3.1:405b-instruct-q6_K", expected: "Llama3.1 405B Instruct Q6 K"},
		{input: "llama3.1:405b-instruct-q8_0", expected: "Llama3.1 405B Instruct Q8_0"},
		{input: "llama3.1:405b-instruct-fp16", expected: "Llama3.1 405B Instruct FP16"},
		{input: "llama3.1:405b-text-q2_K", expected: "Llama3.1 405B Text Q2_K"},
		{input: "llama3.1:405b-text-q3_K_S", expected: "Llama3.1 405B Text Q3 K_S"},
		{input: "llama3.1:405b-text-q3_K_M", expected: "Llama3.1 405B Text Q3 K_M"},
		{input: "llama3.1:405b-text-q3_K_L", expected: "Llama3.1 405B Text Q3 K_L"},
		{input: "llama3.1:405b-text-q4_0", expected: "Llama3.1 405B Text Q4_0"},
		{input: "llama3.1:405b-text-q4_1", expected: "Llama3.1 405B Text Q4 1"},
		{input: "llama3.1:405b-text-q4_K_S", expected: "Llama3.1 405B Text Q4 K_S"},
		{input: "llama3.1:405b-text-q4_K_M", expected: "Llama3.1 405B Text Q4_K_M"},
		{input: "llama3.1:405b-text-q5_0", expected: "Llama3.1 405B Text Q5 0"},
		{input: "llama3.1:405b-text-q5_1", expected: "Llama3.1 405B Text Q5 1"},
		{input: "llama3.1:405b-text-q5_K_S", expected: "Llama3.1 405B Text Q5 K_S"},
		{input: "llama3.1:405b-text-q5_K_M", expected: "Llama3.1 405B Text Q5 K_M"},
		{input: "llama3.1:405b-text-q6_K", expected: "Llama3.1 405B Text Q6 K"},
		{input: "llama3.1:405b-text-q8_0", expected: "Llama3.1 405B Text Q8_0"},
		{input: "llama3.1:405b-text-fp16", expected: "Llama3.1 405B Text FP16"},
		{input: "llama3.2:latest", expected: "Llama3.2 (latest)"},
		{input: "llama3.2", expected: "Llama3.2"},
		{input: "llama3.2:1b", expected: "Llama3.2 1B"},
		{input: "llama3.2:3b", expected: "Llama3.2 3B"},
		{input: "llama3.2:1b-instruct-q2_K", expected: "Llama3.2 1B Instruct Q2_K"},
		{input: "llama3.2:1b-instruct-q3_K_S", expected: "Llama3.2 1B Instruct Q3 K_S"},
		{input: "llama3.2:1b-instruct-q3_K_M", expected: "Llama3.2 1B Instruct Q3 K_M"},
		{input: "llama3.2:1b-instruct-q3_K_L", expected: "Llama3.2 1B Instruct Q3 K_L"},
		{input: "llama3.2:1b-instruct-q4_0", expected: "Llama3.2 1B Instruct Q4_0"},
		{input: "llama3.2:1b-instruct-q4_1", expected: "Llama3.2 1B Instruct Q4 1"},
		{input: "llama3.2:1b-instruct-q4_K_S", expected: "Llama3.2 1B Instruct Q4 K_S"},
		{input: "llama3.2:1b-instruct-q4_K_M", expected: "Llama3.2 1B Instruct Q4_K_M"},
		{input: "llama3.2:1b-instruct-q5_0", expected: "Llama3.2 1B Instruct Q5 0"},
		{input: "llama3.2:1b-instruct-q5_1", expected: "Llama3.2 1B Instruct Q5 1"},
		{input: "llama3.2:1b-instruct-q5_K_S", expected: "Llama3.2 1B Instruct Q5 K_S"},
		{input: "llama3.2:1b-instruct-q5_K_M", expected: "Llama3.2 1B Instruct Q5 K_M"},
		{input: "llama3.2:1b-instruct-q6_K", expected: "Llama3.2 1B Instruct Q6 K"},
		{input: "llama3.2:1b-instruct-q8_0", expected: "Llama3.2 1B Instruct Q8_0"},
		{input: "llama3.2:1b-instruct-fp16", expected: "Llama3.2 1B Instruct FP16"},
		{input: "llama3.2:1b-text-q2_K", expected: "Llama3.2 1B Text Q2_K"},
		{input: "llama3.2:1b-text-q3_K_S", expected: "Llama3.2 1B Text Q3 K_S"},
		{input: "llama3.2:1b-text-q3_K_M", expected: "Llama3.2 1B Text Q3 K_M"},
		{input: "llama3.2:1b-text-q3_K_L", expected: "Llama3.2 1B Text Q3 K_L"},
		{input: "llama3.2:1b-text-q4_0", expected: "Llama3.2 1B Text Q4_0"},
		{input: "llama3.2:1b-text-q4_1", expected: "Llama3.2 1B Text Q4 1"},
		{input: "llama3.2:1b-text-q4_K_S", expected: "Llama3.2 1B Text Q4 K_S"},
		{input: "llama3.2:1b-text-q4_K_M", expected: "Llama3.2 1B Text Q4_K_M"},
		{input: "llama3.2:1b-text-q5_0", expected: "Llama3.2 1B Text Q5 0"},
		{input: "llama3.2:1b-text-q5_1", expected: "Llama3.2 1B Text Q5 1"},
		{input: "llama3.2:1b-text-q5_K_S", expected: "Llama3.2 1B Text Q5 K_S"},
		{input: "llama3.2:1b-text-q5_K_M", expected: "Llama3.2 1B Text Q5 K_M"},
		{input: "llama3.2:1b-text-q6_K", expected: "Llama3.2 1B Text Q6 K"},
		{input: "llama3.2:1b-text-q8_0", expected: "Llama3.2 1B Text Q8_0"},
		{input: "llama3.2:1b-text-fp16", expected: "Llama3.2 1B Text FP16"},
		{input: "llama3.2:3b-instruct-q2_K", expected: "Llama3.2 3B Instruct Q2_K"},
		{input: "llama3.2:3b-instruct-q3_K_S", expected: "Llama3.2 3B Instruct Q3 K_S"},
		{input: "llama3.2:3b-instruct-q3_K_M", expected: "Llama3.2 3B Instruct Q3 K_M"},
		{input: "llama3.2:3b-instruct-q3_K_L", expected: "Llama3.2 3B Instruct Q3 K_L"},
		{input: "llama3.2:3b-instruct-q4_0", expected: "Llama3.2 3B Instruct Q4_0"},
		{input: "llama3.2:3b-instruct-q4_1", expected: "Llama3.2 3B Instruct Q4 1"},
		{input: "llama3.2:3b-instruct-q4_K_S", expected: "Llama3.2 3B Instruct Q4 K_S"},
		{input: "llama3.2:3b-instruct-q4_K_M", expected: "Llama3.2 3B Instruct Q4_K_M"},
		{input: "llama3.2:3b-instruct-q5_0", expected: "Llama3.2 3B Instruct Q5 0"},
		{input: "llama3.2:3b-instruct-q5_1", expected: "Llama3.2 3B Instruct Q5 1"},
		{input: "llama3.2:3b-instruct-q5_K_S", expected: "Llama3.2 3B Instruct Q5 K_S"},
		{input: "llama3.2:3b-instruct-q5_K_M", expected: "Llama3.2 3B Instruct Q5 K_M"},
		{input: "llama3.2:3b-instruct-q6_K", expected: "Llama3.2 3B Instruct Q6 K"},
		{input: "llama3.2:3b-instruct-q8_0", expected: "Llama3.2 3B Instruct Q8_0"},
		{input: "llama3.2:3b-instruct-fp16", expected: "Llama3.2 3B Instruct FP16"},
		{input: "llama3.2:3b-text-q2_K", expected: "Llama3.2 3B Text Q2_K"},
		{input: "llama3.2:3b-text-q3_K_S", expected: "Llama3.2 3B Text Q3 K_S"},
		{input: "llama3.2:3b-text-q3_K_M", expected: "Llama3.2 3B Text Q3 K_M"},
		{input: "llama3.2:3b-text-q3_K_L", expected: "Llama3.2 3B Text Q3 K_L"},
		{input: "llama3.2:3b-text-q4_0", expected: "Llama3.2 3B Text Q4_0"},
		{input: "llama3.2:3b-text-q4_1", expected: "Llama3.2 3B Text Q4 1"},
		{input: "llama3.2:3b-text-q4_K_S", expected: "Llama3.2 3B Text Q4 K_S"},
		{input: "llama3.2:3b-text-q4_K_M", expected: "Llama3.2 3B Text Q4_K_M"},
		{input: "llama3.2:3b-text-q5_0", expected: "Llama3.2 3B Text Q5 0"},
		{input: "llama3.2:3b-text-q5_1", expected: "Llama3.2 3B Text Q5 1"},
		{input: "llama3.2:3b-text-q5_K_S", expected: "Llama3.2 3B Text Q5 K_S"},
		{input: "llama3.2:3b-text-q5_K_M", expected: "Llama3.2 3B Text Q5 K_M"},
		{input: "llama3.2:3b-text-q6_K", expected: "Llama3.2 3B Text Q6 K"},
		{input: "llama3.2:3b-text-q8_0", expected: "Llama3.2 3B Text Q8_0"},
		{input: "llama3.2:3b-text-fp16", expected: "Llama3.2 3B Text FP16"},
		{input: "nomic-embed-text:latest", expected: "Nomic Embed Text (latest)"},
		{input: "nomic-embed-text", expected: "Nomic Embed Text"},
		{input: "nomic-embed-text:v1.5", expected: "Nomic Embed Text v1.5"},
		{input: "nomic-embed-text:137m-v1.5-fp16", expected: "Nomic Embed Text 137M v1.5 FP16"},
		{input: "mistral:latest", expected: "Mistral (latest)"},
		{input: "mistral", expected: "Mistral"},
		{input: "mistral:instruct", expected: "Mistral Instruct"},
		{input: "mistral:text", expected: "Mistral Text"},
		{input: "mistral:v0.1", expected: "Mistral v0.1"},
		{input: "mistral:v0.2", expected: "Mistral v0.2"},
		{input: "mistral:v0.3", expected: "Mistral v0.3"},
		{input: "mistral:7b", expected: "Mistral 7B"},
		{input: "mistral:7b-instruct", expected: "Mistral 7B Instruct"},
		{input: "mistral:7b-instruct-q2_K", expected: "Mistral 7B Instruct Q2_K"},
		{input: "mistral:7b-instruct-v0.2-q2_K", expected: "Mistral 7B Instruct v0.2 Q2_K"},
		{input: "mistral:7b-instruct-q3_K_S", expected: "Mistral 7B Instruct Q3 K_S"},
		{input: "mistral:7b-instruct-v0.2-q3_K_S", expected: "Mistral 7B Instruct v0.2 Q3 K_S"},
		{input: "mistral:7b-instruct-q3_K_M", expected: "Mistral 7B Instruct Q3 K_M"},
		{input: "mistral:7b-instruct-v0.2-q3_K_M", expected: "Mistral 7B Instruct v0.2 Q3 K_M"},
		{input: "mistral:7b-instruct-q3_K_L", expected: "Mistral 7B Instruct Q3 K_L"},
		{input: "mistral:7b-instruct-v0.2-q3_K_L", expected: "Mistral 7B Instruct v0.2 Q3 K_L"},
		{input: "mistral:7b-instruct-q4_0", expected: "Mistral 7B Instruct Q4_0"},
		{input: "mistral:7b-instruct-v0.2-q4_0", expected: "Mistral 7B Instruct v0.2 Q4_0"},
		{input: "mistral:7b-instruct-q4_1", expected: "Mistral 7B Instruct Q4 1"},
		{input: "mistral:7b-instruct-v0.2-q4_1", expected: "Mistral 7B Instruct v0.2 Q4 1"},
		{input: "mistral:7b-instruct-q4_K_S", expected: "Mistral 7B Instruct Q4 K_S"},
		{input: "mistral:7b-instruct-v0.2-q4_K_S", expected: "Mistral 7B Instruct v0.2 Q4 K_S"},
		{input: "mistral:7b-instruct-v0.2-q4_K_M", expected: "Mistral 7B Instruct v0.2 Q4_K_M"},
		{input: "mistral:7b-instruct-v0.2-q5_0", expected: "Mistral 7B Instruct v0.2 Q5 0"},
		{input: "mistral:7b-instruct-v0.2-q5_1", expected: "Mistral 7B Instruct v0.2 Q5 1"},
		{input: "mistral:7b-instruct-v0.2-q5_K_S", expected: "Mistral 7B Instruct v0.2 Q5 K_S"},
		{input: "mistral:7b-instruct-v0.2-q5_K_M", expected: "Mistral 7B Instruct v0.2 Q5 K_M"},
		{input: "mistral:7b-instruct-v0.2-q6_K", expected: "Mistral 7B Instruct v0.2 Q6 K"},
		{input: "mistral:7b-instruct-v0.2-q8_0", expected: "Mistral 7B Instruct v0.2 Q8_0"},
		{input: "mistral:7b-instruct-v0.2-fp16", expected: "Mistral 7B Instruct v0.2 FP16"},
		{input: "mistral:7b-instruct-v0.3-q2_K", expected: "Mistral 7B Instruct v0.3 Q2_K"},
		{input: "mistral:7b-instruct-v0.3-q3_K_S", expected: "Mistral 7B Instruct v0.3 Q3 K_S"},
		{input: "mistral:7b-instruct-v0.3-q3_K_M", expected: "Mistral 7B Instruct v0.3 Q3 K_M"},
		{input: "mistral:7b-instruct-v0.3-q3_K_L", expected: "Mistral 7B Instruct v0.3 Q3 K_L"},
		{input: "mistral:7b-instruct-v0.3-q4_0", expected: "Mistral 7B Instruct v0.3 Q4_0"},
		{input: "mistral:7b-instruct-v0.3-q4_1", expected: "Mistral 7B Instruct v0.3 Q4 1"},
		{input: "mistral:7b-instruct-v0.3-q4_K_S", expected: "Mistral 7B Instruct v0.3 Q4 K_S"},
		{input: "mistral:7b-instruct-q4_K_M", expected: "Mistral 7B Instruct Q4_K_M"},
		{input: "mistral:7b-instruct-v0.3-q4_K_M", expected: "Mistral 7B Instruct v0.3 Q4_K_M"},
		{input: "mistral:7b-instruct-q5_0", expected: "Mistral 7B Instruct Q5 0"},
		{input: "mistral:7b-instruct-v0.3-q5_0", expected: "Mistral 7B Instruct v0.3 Q5 0"},
		{input: "mistral:7b-instruct-q5_1", expected: "Mistral 7B Instruct Q5 1"},
		{input: "mistral:7b-instruct-v0.3-q5_1", expected: "Mistral 7B Instruct v0.3 Q5 1"},
		{input: "mistral:7b-instruct-q5_K_S", expected: "Mistral 7B Instruct Q5 K_S"},
		{input: "mistral:7b-instruct-v0.3-q5_K_S", expected: "Mistral 7B Instruct v0.3 Q5 K_S"},
		{input: "mistral:7b-instruct-q5_K_M", expected: "Mistral 7B Instruct Q5 K_M"},
		{input: "mistral:7b-instruct-v0.3-q5_K_M", expected: "Mistral 7B Instruct v0.3 Q5 K_M"},
		{input: "mistral:7b-instruct-q6_K", expected: "Mistral 7B Instruct Q6 K"},
		{input: "mistral:7b-instruct-v0.3-q6_K", expected: "Mistral 7B Instruct v0.3 Q6 K"},
		{input: "mistral:7b-instruct-q8_0", expected: "Mistral 7B Instruct Q8_0"},
		{input: "mistral:7b-instruct-v0.3-q8_0", expected: "Mistral 7B Instruct v0.3 Q8_0"},
		{input: "mistral:7b-instruct-fp16", expected: "Mistral 7B Instruct FP16"},
		{input: "mistral:7b-instruct-v0.3-fp16", expected: "Mistral 7B Instruct v0.3 FP16"},
		{input: "mistral:7b-text", expected: "Mistral 7B Text"},
		{input: "mistral:7b-text-q2_K", expected: "Mistral 7B Text Q2_K"},
		{input: "mistral:7b-text-v0.2-q2_K", expected: "Mistral 7B Text v0.2 Q2_K"},
		{input: "mistral:7b-text-q3_K_S", expected: "Mistral 7B Text Q3 K_S"},
		{input: "mistral:7b-text-v0.2-q3_K_S", expected: "Mistral 7B Text v0.2 Q3 K_S"},
		{input: "mistral:7b-text-q3_K_M", expected: "Mistral 7B Text Q3 K_M"},
		{input: "mistral:7b-text-v0.2-q3_K_M", expected: "Mistral 7B Text v0.2 Q3 K_M"},
		{input: "mistral:7b-text-q3_K_L", expected: "Mistral 7B Text Q3 K_L"},
		{input: "mistral:7b-text-v0.2-q3_K_L", expected: "Mistral 7B Text v0.2 Q3 K_L"},
		{input: "mistral:7b-text-q4_0", expected: "Mistral 7B Text Q4_0"},
		{input: "mistral:7b-text-v0.2-q4_0", expected: "Mistral 7B Text v0.2 Q4_0"},
		{input: "mistral:7b-text-q4_1", expected: "Mistral 7B Text Q4 1"},
		{input: "mistral:7b-text-v0.2-q4_1", expected: "Mistral 7B Text v0.2 Q4 1"},
		{input: "mistral:7b-text-q4_K_S", expected: "Mistral 7B Text Q4 K_S"},
		{input: "mistral:7b-text-v0.2-q4_K_S", expected: "Mistral 7B Text v0.2 Q4 K_S"},
		{input: "mistral:7b-text-q4_K_M", expected: "Mistral 7B Text Q4_K_M"},
		{input: "mistral:7b-text-v0.2-q4_K_M", expected: "Mistral 7B Text v0.2 Q4_K_M"},
		{input: "mistral:7b-text-q5_0", expected: "Mistral 7B Text Q5 0"},
		{input: "mistral:7b-text-v0.2-q5_0", expected: "Mistral 7B Text v0.2 Q5 0"},
		{input: "mistral:7b-text-q5_1", expected: "Mistral 7B Text Q5 1"},
		{input: "mistral:7b-text-v0.2-q5_1", expected: "Mistral 7B Text v0.2 Q5 1"},
		{input: "mistral:7b-text-q5_K_S", expected: "Mistral 7B Text Q5 K_S"},
		{input: "mistral:7b-text-v0.2-q5_K_S", expected: "Mistral 7B Text v0.2 Q5 K_S"},
		{input: "mistral:7b-text-q5_K_M", expected: "Mistral 7B Text Q5 K_M"},
		{input: "mistral:7b-text-v0.2-q5_K_M", expected: "Mistral 7B Text v0.2 Q5 K_M"},
		{input: "mistral:7b-text-q6_K", expected: "Mistral 7B Text Q6 K"},
		{input: "mistral:7b-text-v0.2-q6_K", expected: "Mistral 7B Text v0.2 Q6 K"},
		{input: "mistral:7b-text-q8_0", expected: "Mistral 7B Text Q8_0"},
		{input: "mistral:7b-text-v0.2-q8_0", expected: "Mistral 7B Text v0.2 Q8_0"},
		{input: "mistral:7b-text-fp16", expected: "Mistral 7B Text FP16"},
		{input: "mistral:7b-text-v0.2-fp16", expected: "Mistral 7B Text v0.2 FP16"},
		{input: "qwen2.5:latest", expected: "Qwen2.5 (latest)"},
		{input: "qwen2.5", expected: "Qwen2.5"},
		{input: "qwen2.5:0.5b", expected: "Qwen2.5 0.5B"},
		{input: "qwen2.5:1.5b", expected: "Qwen2.5 1.5B"},
		{input: "qwen2.5:3b", expected: "Qwen2.5 3B"},
		{input: "qwen2.5:7b", expected: "Qwen2.5 7B"},
		{input: "qwen2.5:14b", expected: "Qwen2.5 14B"},
		{input: "qwen2.5:32b", expected: "Qwen2.5 32B"},
		{input: "qwen2.5:72b", expected: "Qwen2.5 72B"},
		{input: "qwen2.5:0.5b-base", expected: "Qwen2.5 0.5B Base"},
		{input: "qwen2.5:0.5b-base-q2_K", expected: "Qwen2.5 0.5B Base Q2_K"},
		{input: "qwen2.5:0.5b-base-q3_K_S", expected: "Qwen2.5 0.5B Base Q3 K_S"},
		{input: "qwen2.5:0.5b-base-q3_K_M", expected: "Qwen2.5 0.5B Base Q3 K_M"},
		{input: "qwen2.5:0.5b-base-q3_K_L", expected: "Qwen2.5 0.5B Base Q3 K_L"},
		{input: "qwen2.5:0.5b-base-q4_0", expected: "Qwen2.5 0.5B Base Q4_0"},
		{input: "qwen2.5:0.5b-base-q4_1", expected: "Qwen2.5 0.5B Base Q4 1"},
		{input: "qwen2.5:0.5b-base-q4_K_S", expected: "Qwen2.5 0.5B Base Q4 K_S"},
		{input: "qwen2.5:0.5b-base-q4_K_M", expected: "Qwen2.5 0.5B Base Q4_K_M"},
		{input: "qwen2.5:0.5b-base-q5_0", expected: "Qwen2.5 0.5B Base Q5 0"},
		{input: "qwen2.5:0.5b-base-q5_1", expected: "Qwen2.5 0.5B Base Q5 1"},
		{input: "qwen2.5:0.5b-base-q5_K_S", expected: "Qwen2.5 0.5B Base Q5 K_S"},
		{input: "qwen2.5:0.5b-base-q8_0", expected: "Qwen2.5 0.5B Base Q8_0"},
		{input: "qwen2.5:0.5b-instruct", expected: "Qwen2.5 0.5B Instruct"},
		{input: "qwen2.5:0.5b-instruct-q2_K", expected: "Qwen2.5 0.5B Instruct Q2_K"},
		{input: "qwen2.5:0.5b-instruct-q3_K_S", expected: "Qwen2.5 0.5B Instruct Q3 K_S"},
		{input: "qwen2.5:0.5b-instruct-q3_K_M", expected: "Qwen2.5 0.5B Instruct Q3 K_M"},
		{input: "qwen2.5:0.5b-instruct-q3_K_L", expected: "Qwen2.5 0.5B Instruct Q3 K_L"},
		{input: "qwen2.5:0.5b-instruct-q4_0", expected: "Qwen2.5 0.5B Instruct Q4_0"},
		{input: "qwen2.5:0.5b-instruct-q4_1", expected: "Qwen2.5 0.5B Instruct Q4 1"},
		{input: "qwen2.5:0.5b-instruct-q4_K_S", expected: "Qwen2.5 0.5B Instruct Q4 K_S"},
		{input: "qwen2.5:0.5b-instruct-q4_K_M", expected: "Qwen2.5 0.5B Instruct Q4_K_M"},
		{input: "qwen2.5:0.5b-instruct-q5_0", expected: "Qwen2.5 0.5B Instruct Q5 0"},
		{input: "qwen2.5:0.5b-instruct-q5_1", expected: "Qwen2.5 0.5B Instruct Q5 1"},
		{input: "qwen2.5:0.5b-instruct-q5_K_S", expected: "Qwen2.5 0.5B Instruct Q5 K_S"},
		{input: "qwen2.5:0.5b-instruct-q5_K_M", expected: "Qwen2.5 0.5B Instruct Q5 K_M"},
		{input: "qwen2.5:0.5b-instruct-q6_K", expected: "Qwen2.5 0.5B Instruct Q6 K"},
		{input: "qwen2.5:0.5b-instruct-q8_0", expected: "Qwen2.5 0.5B Instruct Q8_0"},
		{input: "qwen2.5:0.5b-instruct-fp16", expected: "Qwen2.5 0.5B Instruct FP16"},
		{input: "qwen2.5:1.5b-instruct", expected: "Qwen2.5 1.5B Instruct"},
		{input: "qwen2.5:1.5b-instruct-q2_K", expected: "Qwen2.5 1.5B Instruct Q2_K"},
		{input: "qwen2.5:1.5b-instruct-q3_K_S", expected: "Qwen2.5 1.5B Instruct Q3 K_S"},
		{input: "qwen2.5:1.5b-instruct-q3_K_M", expected: "Qwen2.5 1.5B Instruct Q3 K_M"},
		{input: "qwen2.5:1.5b-instruct-q3_K_L", expected: "Qwen2.5 1.5B Instruct Q3 K_L"},
		{input: "qwen2.5:1.5b-instruct-q4_0", expected: "Qwen2.5 1.5B Instruct Q4_0"},
		{input: "qwen2.5:1.5b-instruct-q4_1", expected: "Qwen2.5 1.5B Instruct Q4 1"},
		{input: "qwen2.5:1.5b-instruct-q4_K_S", expected: "Qwen2.5 1.5B Instruct Q4 K_S"},
		{input: "qwen2.5:1.5b-instruct-q4_K_M", expected: "Qwen2.5 1.5B Instruct Q4_K_M"},
		{input: "qwen2.5:1.5b-instruct-q5_0", expected: "Qwen2.5 1.5B Instruct Q5 0"},
		{input: "qwen2.5:1.5b-instruct-q5_1", expected: "Qwen2.5 1.5B Instruct Q5 1"},
		{input: "qwen2.5:1.5b-instruct-q5_K_S", expected: "Qwen2.5 1.5B Instruct Q5 K_S"},
		{input: "qwen2.5:1.5b-instruct-q5_K_M", expected: "Qwen2.5 1.5B Instruct Q5 K_M"},
		{input: "qwen2.5:1.5b-instruct-q6_K", expected: "Qwen2.5 1.5B Instruct Q6 K"},
		{input: "qwen2.5:1.5b-instruct-q8_0", expected: "Qwen2.5 1.5B Instruct Q8_0"},
		{input: "qwen2.5:1.5b-instruct-fp16", expected: "Qwen2.5 1.5B Instruct FP16"},
		{input: "qwen2.5:3b-instruct", expected: "Qwen2.5 3B Instruct"},
		{input: "qwen2.5:3b-instruct-q2_K", expected: "Qwen2.5 3B Instruct Q2_K"},
		{input: "qwen2.5:3b-instruct-q3_K_S", expected: "Qwen2.5 3B Instruct Q3 K_S"},
		{input: "qwen2.5:3b-instruct-q3_K_M", expected: "Qwen2.5 3B Instruct Q3 K_M"},
		{input: "qwen2.5:3b-instruct-q3_K_L", expected: "Qwen2.5 3B Instruct Q3 K_L"},
		{input: "qwen2.5:3b-instruct-q4_0", expected: "Qwen2.5 3B Instruct Q4_0"},
		{input: "qwen2.5:3b-instruct-q4_1", expected: "Qwen2.5 3B Instruct Q4 1"},
		{input: "qwen2.5:3b-instruct-q4_K_S", expected: "Qwen2.5 3B Instruct Q4 K_S"},
		{input: "qwen2.5:3b-instruct-q4_K_M", expected: "Qwen2.5 3B Instruct Q4_K_M"},
		{input: "qwen2.5:3b-instruct-q5_0", expected: "Qwen2.5 3B Instruct Q5 0"},
		{input: "qwen2.5:3b-instruct-q5_1", expected: "Qwen2.5 3B Instruct Q5 1"},
		{input: "qwen2.5:3b-instruct-q5_K_S", expected: "Qwen2.5 3B Instruct Q5 K_S"},
		{input: "qwen2.5:3b-instruct-q5_K_M", expected: "Qwen2.5 3B Instruct Q5 K_M"},
		{input: "qwen2.5:3b-instruct-q6_K", expected: "Qwen2.5 3B Instruct Q6 K"},
		{input: "qwen2.5:3b-instruct-q8_0", expected: "Qwen2.5 3B Instruct Q8_0"},
		{input: "qwen2.5:3b-instruct-fp16", expected: "Qwen2.5 3B Instruct FP16"},
		{input: "qwen2.5:7b-instruct", expected: "Qwen2.5 7B Instruct"},
		{input: "qwen2.5:7b-instruct-q2_K", expected: "Qwen2.5 7B Instruct Q2_K"},
		{input: "qwen2.5:7b-instruct-q3_K_S", expected: "Qwen2.5 7B Instruct Q3 K_S"},
		{input: "qwen2.5:7b-instruct-q3_K_M", expected: "Qwen2.5 7B Instruct Q3 K_M"},
		{input: "qwen2.5:7b-instruct-q3_K_L", expected: "Qwen2.5 7B Instruct Q3 K_L"},
		{input: "qwen2.5:7b-instruct-q4_0", expected: "Qwen2.5 7B Instruct Q4_0"},
		{input: "qwen2.5:7b-instruct-q4_1", expected: "Qwen2.5 7B Instruct Q4 1"},
		{input: "qwen2.5:7b-instruct-q4_K_S", expected: "Qwen2.5 7B Instruct Q4 K_S"},
		{input: "qwen2.5:7b-instruct-q4_K_M", expected: "Qwen2.5 7B Instruct Q4_K_M"},
		{input: "qwen2.5:7b-instruct-q5_0", expected: "Qwen2.5 7B Instruct Q5 0"},
		{input: "qwen2.5:7b-instruct-q5_1", expected: "Qwen2.5 7B Instruct Q5 1"},
		{input: "qwen2.5:7b-instruct-q5_K_S", expected: "Qwen2.5 7B Instruct Q5 K_S"},
		{input: "qwen2.5:7b-instruct-q5_K_M", expected: "Qwen2.5 7B Instruct Q5 K_M"},
		{input: "qwen2.5:7b-instruct-q6_K", expected: "Qwen2.5 7B Instruct Q6 K"},
		{input: "qwen2.5:7b-instruct-q8_0", expected: "Qwen2.5 7B Instruct Q8_0"},
		{input: "qwen2.5:7b-instruct-fp16", expected: "Qwen2.5 7B Instruct FP16"},
		{input: "qwen2.5:14b-instruct", expected: "Qwen2.5 14B Instruct"},
		{input: "qwen2.5:14b-instruct-q2_K", expected: "Qwen2.5 14B Instruct Q2_K"},
		{input: "qwen2.5:14b-instruct-q3_K_S", expected: "Qwen2.5 14B Instruct Q3 K_S"},
		{input: "qwen2.5:14b-instruct-q3_K_M", expected: "Qwen2.5 14B Instruct Q3 K_M"},
		{input: "qwen2.5:14b-instruct-q3_K_L", expected: "Qwen2.5 14B Instruct Q3 K_L"},
		{input: "qwen2.5:14b-instruct-q4_0", expected: "Qwen2.5 14B Instruct Q4_0"},
		{input: "qwen2.5:14b-instruct-q4_1", expected: "Qwen2.5 14B Instruct Q4 1"},
		{input: "qwen2.5:14b-instruct-q4_K_S", expected: "Qwen2.5 14B Instruct Q4 K_S"},
		{input: "qwen2.5:14b-instruct-q4_K_M", expected: "Qwen2.5 14B Instruct Q4_K_M"},
		{input: "qwen2.5:14b-instruct-q5_0", expected: "Qwen2.5 14B Instruct Q5 0"},
		{input: "qwen2.5:14b-instruct-q5_1", expected: "Qwen2.5 14B Instruct Q5 1"},
		{input: "qwen2.5:14b-instruct-q5_K_S", expected: "Qwen2.5 14B Instruct Q5 K_S"},
		{input: "qwen2.5:14b-instruct-q5_K_M", expected: "Qwen2.5 14B Instruct Q5 K_M"},
		{input: "qwen2.5:14b-instruct-q6_K", expected: "Qwen2.5 14B Instruct Q6 K"},
		{input: "qwen2.5:14b-instruct-q8_0", expected: "Qwen2.5 14B Instruct Q8_0"},
		{input: "qwen2.5:14b-instruct-fp16", expected: "Qwen2.5 14B Instruct FP16"},
		{input: "qwen2.5:32b-instruct", expected: "Qwen2.5 32B Instruct"},
		{input: "qwen2.5:32b-instruct-q2_K", expected: "Qwen2.5 32B Instruct Q2_K"},
		{input: "qwen2.5:32b-instruct-q3_K_S", expected: "Qwen2.5 32B Instruct Q3 K_S"},
		{input: "qwen2.5:32b-instruct-q3_K_M", expected: "Qwen2.5 32B Instruct Q3 K_M"},
		{input: "qwen2.5:32b-instruct-q3_K_L", expected: "Qwen2.5 32B Instruct Q3 K_L"},
		{input: "qwen2.5:32b-instruct-q4_0", expected: "Qwen2.5 32B Instruct Q4_0"},
		{input: "qwen2.5:32b-instruct-q4_1", expected: "Qwen2.5 32B Instruct Q4 1"},
		{input: "qwen2.5:32b-instruct-q4_K_S", expected: "Qwen2.5 32B Instruct Q4 K_S"},
		{input: "qwen2.5:32b-instruct-q4_K_M", expected: "Qwen2.5 32B Instruct Q4_K_M"},
		{input: "qwen2.5:32b-instruct-q5_0", expected: "Qwen2.5 32B Instruct Q5 0"},
		{input: "qwen2.5:32b-instruct-q5_1", expected: "Qwen2.5 32B Instruct Q5 1"},
		{input: "qwen2.5:32b-instruct-q5_K_S", expected: "Qwen2.5 32B Instruct Q5 K_S"},
		{input: "qwen2.5:32b-instruct-q5_K_M", expected: "Qwen2.5 32B Instruct Q5 K_M"},
		{input: "qwen2.5:32b-instruct-q6_K", expected: "Qwen2.5 32B Instruct Q6 K"},
		{input: "qwen2.5:32b-instruct-q8_0", expected: "Qwen2.5 32B Instruct Q8_0"},
		{input: "qwen2.5:32b-instruct-fp16", expected: "Qwen2.5 32B Instruct FP16"},
		{input: "qwen2.5:72b-instruct", expected: "Qwen2.5 72B Instruct"},
		{input: "qwen2.5:72b-instruct-q2_K", expected: "Qwen2.5 72B Instruct Q2_K"},
		{input: "qwen2.5:72b-instruct-q3_K_S", expected: "Qwen2.5 72B Instruct Q3 K_S"},
		{input: "qwen2.5:72b-instruct-q3_K_M", expected: "Qwen2.5 72B Instruct Q3 K_M"},
		{input: "qwen2.5:72b-instruct-q3_K_L", expected: "Qwen2.5 72B Instruct Q3 K_L"},
		{input: "qwen2.5:72b-instruct-q4_0", expected: "Qwen2.5 72B Instruct Q4_0"},
		{input: "qwen2.5:72b-instruct-q4_1", expected: "Qwen2.5 72B Instruct Q4 1"},
		{input: "qwen2.5:72b-instruct-q4_K_S", expected: "Qwen2.5 72B Instruct Q4 K_S"},
		{input: "qwen2.5:72b-instruct-q4_K_M", expected: "Qwen2.5 72B Instruct Q4_K_M"},
		{input: "qwen2.5:72b-instruct-q5_0", expected: "Qwen2.5 72B Instruct Q5 0"},
		{input: "qwen2.5:72b-instruct-q5_1", expected: "Qwen2.5 72B Instruct Q5 1"},
		{input: "qwen2.5:72b-instruct-q5_K_S", expected: "Qwen2.5 72B Instruct Q5 K_S"},
		{input: "qwen2.5:72b-instruct-q5_K_M", expected: "Qwen2.5 72B Instruct Q5 K_M"},
		{input: "qwen2.5:72b-instruct-q6_K", expected: "Qwen2.5 72B Instruct Q6 K"},
		{input: "qwen2.5:72b-instruct-q8_0", expected: "Qwen2.5 72B Instruct Q8_0"},
		{input: "qwen2.5:72b-instruct-fp16", expected: "Qwen2.5 72B Instruct FP16"},
		{input: "qwen3:latest", expected: "Qwen3 (latest)"},
		{input: "qwen3", expected: "Qwen3"},
		{input: "qwen3:0.6b", expected: "Qwen3 0.6B"},
		{input: "qwen3:1.7b", expected: "Qwen3 1.7B"},
		{input: "qwen3:4b", expected: "Qwen3 4B"},
		{input: "qwen3:8b", expected: "Qwen3 8B"},
		{input: "qwen3:14b", expected: "Qwen3 14B"},
		{input: "qwen3:30b", expected: "Qwen3 30B"},
		{input: "qwen3:32b", expected: "Qwen3 32B"},
		{input: "qwen3:235b", expected: "Qwen3 235B"},
		{input: "qwen3:0.6b-q4_K_M", expected: "Qwen3 0.6B Q4_K_M"},
		{input: "qwen3:0.6b-q8_0", expected: "Qwen3 0.6B Q8_0"},
		{input: "qwen3:0.6b-fp16", expected: "Qwen3 0.6B FP16"},
		{input: "qwen3:1.7b-q4_K_M", expected: "Qwen3 1.7B Q4_K_M"},
		{input: "qwen3:1.7b-q8_0", expected: "Qwen3 1.7B Q8_0"},
		{input: "qwen3:1.7b-fp16", expected: "Qwen3 1.7B FP16"},
		{input: "qwen3:4b-instruct", expected: "Qwen3 4B Instruct"},
		{input: "qwen3:4b-instruct-2507-q4_K_M", expected: "Qwen3 4B Instruct 2507 Q4_K_M"},
		{input: "qwen3:4b-instruct-2507-q8_0", expected: "Qwen3 4B Instruct 2507 Q8_0"},
		{input: "qwen3:4b-instruct-2507-fp16", expected: "Qwen3 4B Instruct 2507 FP16"},
		{input: "qwen3:4b-thinking", expected: "Qwen3 4B Thinking"},
		{input: "qwen3:4b-thinking-2507-q4_K_M", expected: "Qwen3 4B Thinking 2507 Q4_K_M"},
		{input: "qwen3:4b-thinking-2507-q8_0", expected: "Qwen3 4B Thinking 2507 Q8_0"},
		{input: "qwen3:4b-thinking-2507-fp16", expected: "Qwen3 4B Thinking 2507 FP16"},
		{input: "qwen3:4b-q4_K_M", expected: "Qwen3 4B Q4_K_M"},
		{input: "qwen3:4b-q8_0", expected: "Qwen3 4B Q8_0"},
		{input: "qwen3:4b-fp16", expected: "Qwen3 4B FP16"},
		{input: "qwen3:8b-q4_K_M", expected: "Qwen3 8B Q4_K_M"},
		{input: "qwen3:8b-q8_0", expected: "Qwen3 8B Q8_0"},
		{input: "qwen3:8b-fp16", expected: "Qwen3 8B FP16"},
		{input: "qwen3:14b-q4_K_M", expected: "Qwen3 14B Q4_K_M"},
		{input: "qwen3:14b-q8_0", expected: "Qwen3 14B Q8_0"},
		{input: "qwen3:14b-fp16", expected: "Qwen3 14B FP16"},
		{input: "qwen3:30b-a3b", expected: "Qwen3 30B A3B"},
		{input: "qwen3:30b-a3b-instruct-2507-q4_K_M", expected: "Qwen3 30B A3B Instruct 2507 Q4_K_M"},
		{input: "qwen3:30b-a3b-q4_K_M", expected: "Qwen3 30B A3B Q4_K_M"},
		{input: "qwen3:30b-a3b-instruct-2507-q8_0", expected: "Qwen3 30B A3B Instruct 2507 Q8_0"},
		{input: "qwen3:30b-a3b-thinking-2507-q4_K_M", expected: "Qwen3 30B A3B Thinking 2507 Q4_K_M"},
		{input: "qwen3:30b-a3b-q8_0", expected: "Qwen3 30B A3B Q8_0"},
		{input: "qwen3:30b-a3b-thinking-2507-q8_0", expected: "Qwen3 30B A3B Thinking 2507 Q8_0"},
		{input: "qwen3:30b-a3b-fp16", expected: "Qwen3 30B A3B FP16"},
		{input: "qwen3:30b-a3b-instruct-2507-fp16", expected: "Qwen3 30B A3B Instruct 2507 FP16"},
		{input: "qwen3:30b-a3b-thinking-2507-fp16", expected: "Qwen3 30B A3B Thinking 2507 FP16"},
		{input: "qwen3:30b-instruct", expected: "Qwen3 30B Instruct"},
		{input: "qwen3:30b-thinking", expected: "Qwen3 30B Thinking"},
		{input: "qwen3:32b-q4_K_M", expected: "Qwen3 32B Q4_K_M"},
		{input: "qwen3:32b-q8_0", expected: "Qwen3 32B Q8_0"},
		{input: "qwen3:32b-fp16", expected: "Qwen3 32B FP16"},
		{input: "qwen3:235b-a22b", expected: "Qwen3 235B A22B"},
		{input: "qwen3:235b-a22b-instruct-2507-q4_K_M", expected: "Qwen3 235B A22B Instruct 2507 Q4_K_M"},
		{input: "qwen3:235b-a22b-q4_K_M", expected: "Qwen3 235B A22B Q4_K_M"},
		{input: "qwen3:235b-a22b-instruct-2507-q8_0", expected: "Qwen3 235B A22B Instruct 2507 Q8_0"},
		{input: "qwen3:235b-a22b-thinking-2507-q4_K_M", expected: "Qwen3 235B A22B Thinking 2507 Q4_K_M"},
		{input: "qwen3:235b-a22b-q8_0", expected: "Qwen3 235B A22B Q8_0"},
		{input: "qwen3:235b-a22b-thinking-2507-q8_0", expected: "Qwen3 235B A22B Thinking 2507 Q8_0"},
		{input: "qwen3:235b-a22b-fp16", expected: "Qwen3 235B A22B FP16"},
		{input: "qwen3:235b-a22b-thinking-2507-fp16", expected: "Qwen3 235B A22B Thinking 2507 FP16"},
		{input: "qwen3:235b-instruct", expected: "Qwen3 235B Instruct"},
		{input: "qwen3:235b-thinking", expected: "Qwen3 235B Thinking"},
		{input: "phi3:latest", expected: "Phi3 (latest)"},
		{input: "phi3", expected: "Phi3"},
		{input: "phi3:instruct", expected: "Phi3 Instruct"},
		{input: "phi3:medium", expected: "Phi3 Medium"},
		{input: "phi3:mini", expected: "Phi3 Mini"},
		{input: "phi3:3.8b", expected: "Phi3 3.8B"},
		{input: "phi3:14b", expected: "Phi3 14B"},
		{input: "phi3:3.8b-instruct", expected: "Phi3 3.8B Instruct"},
		{input: "phi3:3.8b-mini-128k-instruct-q2_K", expected: "Phi3 3.8B Mini 128K Instruct Q2_K"},
		{input: "phi3:3.8b-mini-128k-instruct-q3_K_S", expected: "Phi3 3.8B Mini 128K Instruct Q3 K_S"},
		{input: "phi3:3.8b-mini-128k-instruct-q3_K_M", expected: "Phi3 3.8B Mini 128K Instruct Q3 K_M"},
		{input: "phi3:3.8b-mini-128k-instruct-q3_K_L", expected: "Phi3 3.8B Mini 128K Instruct Q3 K_L"},
		{input: "phi3:3.8b-mini-128k-instruct-q4_0", expected: "Phi3 3.8B Mini 128K Instruct Q4_0"},
		{input: "phi3:3.8b-mini-128k-instruct-q4_1", expected: "Phi3 3.8B Mini 128K Instruct Q4 1"},
		{input: "phi3:3.8b-mini-128k-instruct-q4_K_S", expected: "Phi3 3.8B Mini 128K Instruct Q4 K_S"},
		{input: "phi3:3.8b-mini-128k-instruct-q4_K_M", expected: "Phi3 3.8B Mini 128K Instruct Q4_K_M"},
		{input: "phi3:3.8b-mini-128k-instruct-q5_0", expected: "Phi3 3.8B Mini 128K Instruct Q5 0"},
		{input: "phi3:3.8b-mini-128k-instruct-q5_1", expected: "Phi3 3.8B Mini 128K Instruct Q5 1"},
		{input: "phi3:3.8b-mini-128k-instruct-q5_K_S", expected: "Phi3 3.8B Mini 128K Instruct Q5 K_S"},
		{input: "phi3:3.8b-mini-128k-instruct-q5_K_M", expected: "Phi3 3.8B Mini 128K Instruct Q5 K_M"},
		{input: "phi3:3.8b-mini-128k-instruct-q6_K", expected: "Phi3 3.8B Mini 128K Instruct Q6 K"},
		{input: "phi3:3.8b-mini-128k-instruct-q8_0", expected: "Phi3 3.8B Mini 128K Instruct Q8_0"},
		{input: "phi3:3.8b-mini-128k-instruct-fp16", expected: "Phi3 3.8B Mini 128K Instruct FP16"},
		{input: "phi3:3.8b-mini-4k-instruct-q2_K", expected: "Phi3 3.8B Mini 4K Instruct Q2_K"},
		{input: "phi3:3.8b-mini-4k-instruct-q3_K_S", expected: "Phi3 3.8B Mini 4K Instruct Q3 K_S"},
		{input: "phi3:3.8b-mini-4k-instruct-q3_K_M", expected: "Phi3 3.8B Mini 4K Instruct Q3 K_M"},
		{input: "phi3:3.8b-mini-4k-instruct-q3_K_L", expected: "Phi3 3.8B Mini 4K Instruct Q3 K_L"},
		{input: "phi3:3.8b-mini-4k-instruct-q4_0", expected: "Phi3 3.8B Mini 4K Instruct Q4_0"},
		{input: "phi3:3.8b-mini-4k-instruct-q4_1", expected: "Phi3 3.8B Mini 4K Instruct Q4 1"},
		{input: "phi3:3.8b-mini-4k-instruct-q4_K_S", expected: "Phi3 3.8B Mini 4K Instruct Q4 K_S"},
		{input: "phi3:3.8b-mini-4k-instruct-q4_K_M", expected: "Phi3 3.8B Mini 4K Instruct Q4_K_M"},
		{input: "phi3:3.8b-mini-4k-instruct-q5_0", expected: "Phi3 3.8B Mini 4K Instruct Q5 0"},
		{input: "phi3:3.8b-mini-4k-instruct-q5_1", expected: "Phi3 3.8B Mini 4K Instruct Q5 1"},
		{input: "phi3:3.8b-mini-4k-instruct-q5_K_S", expected: "Phi3 3.8B Mini 4K Instruct Q5 K_S"},
		{input: "phi3:3.8b-mini-4k-instruct-q5_K_M", expected: "Phi3 3.8B Mini 4K Instruct Q5 K_M"},
		{input: "phi3:3.8b-mini-4k-instruct-q6_K", expected: "Phi3 3.8B Mini 4K Instruct Q6 K"},
		{input: "phi3:3.8b-mini-4k-instruct-q8_0", expected: "Phi3 3.8B Mini 4K Instruct Q8_0"},
		{input: "phi3:3.8b-mini-4k-instruct-fp16", expected: "Phi3 3.8B Mini 4K Instruct FP16"},
		{input: "phi3:14b-instruct", expected: "Phi3 14B Instruct"},
		{input: "phi3:14b-medium-128k-instruct-q2_K", expected: "Phi3 14B Medium 128K Instruct Q2_K"},
		{input: "phi3:14b-medium-128k-instruct-q3_K_S", expected: "Phi3 14B Medium 128K Instruct Q3 K_S"},
		{input: "phi3:14b-medium-128k-instruct-q3_K_M", expected: "Phi3 14B Medium 128K Instruct Q3 K_M"},
		{input: "phi3:14b-medium-128k-instruct-q3_K_L", expected: "Phi3 14B Medium 128K Instruct Q3 K_L"},
		{input: "phi3:14b-medium-128k-instruct-q4_0", expected: "Phi3 14B Medium 128K Instruct Q4_0"},
		{input: "phi3:14b-medium-128k-instruct-q4_1", expected: "Phi3 14B Medium 128K Instruct Q4 1"},
		{input: "phi3:14b-medium-128k-instruct-q4_K_S", expected: "Phi3 14B Medium 128K Instruct Q4 K_S"},
		{input: "phi3:14b-medium-128k-instruct-q4_K_M", expected: "Phi3 14B Medium 128K Instruct Q4_K_M"},
		{input: "phi3:14b-medium-128k-instruct-q5_0", expected: "Phi3 14B Medium 128K Instruct Q5 0"},
		{input: "phi3:14b-medium-128k-instruct-q5_1", expected: "Phi3 14B Medium 128K Instruct Q5 1"},
		{input: "phi3:14b-medium-128k-instruct-q5_K_S", expected: "Phi3 14B Medium 128K Instruct Q5 K_S"},
		{input: "phi3:14b-medium-128k-instruct-q5_K_M", expected: "Phi3 14B Medium 128K Instruct Q5 K_M"},
		{input: "phi3:14b-medium-128k-instruct-q6_K", expected: "Phi3 14B Medium 128K Instruct Q6 K"},
		{input: "phi3:14b-medium-128k-instruct-q8_0", expected: "Phi3 14B Medium 128K Instruct Q8_0"},
		{input: "phi3:14b-medium-128k-instruct-fp16", expected: "Phi3 14B Medium 128K Instruct FP16"},
		{input: "phi3:14b-medium-4k-instruct-q2_K", expected: "Phi3 14B Medium 4K Instruct Q2_K"},
		{input: "phi3:14b-medium-4k-instruct-q3_K_S", expected: "Phi3 14B Medium 4K Instruct Q3 K_S"},
		{input: "phi3:14b-medium-4k-instruct-q3_K_M", expected: "Phi3 14B Medium 4K Instruct Q3 K_M"},
		{input: "phi3:14b-medium-4k-instruct-q3_K_L", expected: "Phi3 14B Medium 4K Instruct Q3 K_L"},
		{input: "phi3:14b-medium-4k-instruct-q4_0", expected: "Phi3 14B Medium 4K Instruct Q4_0"},
		{input: "phi3:14b-medium-4k-instruct-q4_1", expected: "Phi3 14B Medium 4K Instruct Q4 1"},
		{input: "phi3:14b-medium-4k-instruct-q4_K_S", expected: "Phi3 14B Medium 4K Instruct Q4 K_S"},
		{input: "phi3:14b-medium-4k-instruct-q4_K_M", expected: "Phi3 14B Medium 4K Instruct Q4_K_M"},
		{input: "phi3:14b-medium-4k-instruct-q5_0", expected: "Phi3 14B Medium 4K Instruct Q5 0"},
		{input: "phi3:14b-medium-4k-instruct-q5_1", expected: "Phi3 14B Medium 4K Instruct Q5 1"},
		{input: "phi3:14b-medium-4k-instruct-q5_K_S", expected: "Phi3 14B Medium 4K Instruct Q5 K_S"},
		{input: "phi3:14b-medium-4k-instruct-q5_K_M", expected: "Phi3 14B Medium 4K Instruct Q5 K_M"},
		{input: "phi3:14b-medium-4k-instruct-q6_K", expected: "Phi3 14B Medium 4K Instruct Q6 K"},
		{input: "phi3:14b-medium-4k-instruct-q8_0", expected: "Phi3 14B Medium 4K Instruct Q8_0"},
		{input: "phi3:14b-medium-4k-instruct-fp16", expected: "Phi3 14B Medium 4K Instruct FP16"},
		{input: "phi3:medium-128k", expected: "Phi3 Medium 128K"},
		{input: "phi3:medium-4k", expected: "Phi3 Medium 4K"},
		{input: "phi3:mini-128k", expected: "Phi3 Mini 128K"},
		{input: "phi3:mini-4k", expected: "Phi3 Mini 4K"},
		{input: "llama3:latest", expected: "Llama3 (latest)"},
		{input: "llama3", expected: "Llama3"},
		{input: "llama3:instruct", expected: "Llama3 Instruct"},
		{input: "llama3:text", expected: "Llama3 Text"},
		{input: "llama3:8b", expected: "Llama3 8B"},
		{input: "llama3:70b", expected: "Llama3 70B"},
		{input: "llama3:8b-instruct-q2_K", expected: "Llama3 8B Instruct Q2_K"},
		{input: "llama3:8b-instruct-q3_K_S", expected: "Llama3 8B Instruct Q3 K_S"},
		{input: "llama3:8b-instruct-q3_K_M", expected: "Llama3 8B Instruct Q3 K_M"},
		{input: "llama3:8b-instruct-q3_K_L", expected: "Llama3 8B Instruct Q3 K_L"},
		{input: "llama3:8b-instruct-q4_0", expected: "Llama3 8B Instruct Q4_0"},
		{input: "llama3:8b-instruct-q4_1", expected: "Llama3 8B Instruct Q4 1"},
		{input: "llama3:8b-instruct-q4_K_S", expected: "Llama3 8B Instruct Q4 K_S"},
		{input: "llama3:8b-instruct-q4_K_M", expected: "Llama3 8B Instruct Q4_K_M"},
		{input: "llama3:8b-instruct-q5_0", expected: "Llama3 8B Instruct Q5 0"},
		{input: "llama3:8b-instruct-q5_1", expected: "Llama3 8B Instruct Q5 1"},
		{input: "llama3:8b-instruct-q5_K_S", expected: "Llama3 8B Instruct Q5 K_S"},
		{input: "llama3:8b-instruct-q5_K_M", expected: "Llama3 8B Instruct Q5 K_M"},
		{input: "llama3:8b-instruct-q6_K", expected: "Llama3 8B Instruct Q6 K"},
		{input: "llama3:8b-instruct-q8_0", expected: "Llama3 8B Instruct Q8_0"},
		{input: "llama3:8b-instruct-fp16", expected: "Llama3 8B Instruct FP16"},
		{input: "llama3:8b-text", expected: "Llama3 8B Text"},
		{input: "llama3:8b-text-q2_K", expected: "Llama3 8B Text Q2_K"},
		{input: "llama3:8b-text-q3_K_S", expected: "Llama3 8B Text Q3 K_S"},
		{input: "llama3:8b-text-q3_K_M", expected: "Llama3 8B Text Q3 K_M"},
		{input: "llama3:8b-text-q3_K_L", expected: "Llama3 8B Text Q3 K_L"},
		{input: "llama3:8b-text-q4_0", expected: "Llama3 8B Text Q4_0"},
		{input: "llama3:8b-text-q4_1", expected: "Llama3 8B Text Q4 1"},
		{input: "llama3:8b-text-q4_K_S", expected: "Llama3 8B Text Q4 K_S"},
		{input: "llama3:8b-text-q4_K_M", expected: "Llama3 8B Text Q4_K_M"},
		{input: "llama3:8b-text-q5_0", expected: "Llama3 8B Text Q5 0"},
		{input: "llama3:8b-text-q5_1", expected: "Llama3 8B Text Q5 1"},
		{input: "llama3:8b-text-q5_K_S", expected: "Llama3 8B Text Q5 K_S"},
		{input: "llama3:8b-text-q5_K_M", expected: "Llama3 8B Text Q5 K_M"},
		{input: "llama3:8b-text-q6_K", expected: "Llama3 8B Text Q6 K"},
		{input: "llama3:8b-text-q8_0", expected: "Llama3 8B Text Q8_0"},
		{input: "llama3:8b-text-fp16", expected: "Llama3 8B Text FP16"},
		{input: "llama3:70b-instruct", expected: "Llama3 70B Instruct"},
		{input: "llama3:70b-instruct-q2_K", expected: "Llama3 70B Instruct Q2_K"},
		{input: "llama3:70b-instruct-q3_K_S", expected: "Llama3 70B Instruct Q3 K_S"},
		{input: "llama3:70b-instruct-q3_K_M", expected: "Llama3 70B Instruct Q3 K_M"},
		{input: "llama3:70b-instruct-q3_K_L", expected: "Llama3 70B Instruct Q3 K_L"},
		{input: "llama3:70b-instruct-q4_0", expected: "Llama3 70B Instruct Q4_0"},
		{input: "llama3:70b-instruct-q4_1", expected: "Llama3 70B Instruct Q4 1"},
		{input: "llama3:70b-instruct-q4_K_S", expected: "Llama3 70B Instruct Q4 K_S"},
		{input: "llama3:70b-instruct-q4_K_M", expected: "Llama3 70B Instruct Q4_K_M"},
		{input: "llama3:70b-instruct-q5_0", expected: "Llama3 70B Instruct Q5 0"},
		{input: "llama3:70b-instruct-q5_1", expected: "Llama3 70B Instruct Q5 1"},
		{input: "llama3:70b-instruct-q5_K_S", expected: "Llama3 70B Instruct Q5 K_S"},
		{input: "llama3:70b-instruct-q5_K_M", expected: "Llama3 70B Instruct Q5 K_M"},
		{input: "llama3:70b-instruct-q6_K", expected: "Llama3 70B Instruct Q6 K"},
		{input: "llama3:70b-instruct-q8_0", expected: "Llama3 70B Instruct Q8_0"},
		{input: "llama3:70b-instruct-fp16", expected: "Llama3 70B Instruct FP16"},
		{input: "llama3:70b-text", expected: "Llama3 70B Text"},
		{input: "llama3:70b-text-q2_K", expected: "Llama3 70B Text Q2_K"},
		{input: "llama3:70b-text-q3_K_S", expected: "Llama3 70B Text Q3 K_S"},
		{input: "llama3:70b-text-q3_K_M", expected: "Llama3 70B Text Q3 K_M"},
		{input: "llama3:70b-text-q3_K_L", expected: "Llama3 70B Text Q3 K_L"},
		{input: "llama3:70b-text-q4_0", expected: "Llama3 70B Text Q4_0"},
		{input: "llama3:70b-text-q4_1", expected: "Llama3 70B Text Q4 1"},
		{input: "llama3:70b-text-q4_K_S", expected: "Llama3 70B Text Q4 K_S"},
		{input: "llama3:70b-text-q4_K_M", expected: "Llama3 70B Text Q4_K_M"},
		{input: "llama3:70b-text-q5_0", expected: "Llama3 70B Text Q5 0"},
		{input: "llama3:70b-text-q5_1", expected: "Llama3 70B Text Q5 1"},
		{input: "llama3:70b-text-q5_K_S", expected: "Llama3 70B Text Q5 K_S"},
		{input: "llama3:70b-text-q5_K_M", expected: "Llama3 70B Text Q5 K_M"},
		{input: "llama3:70b-text-q6_K", expected: "Llama3 70B Text Q6 K"},
		{input: "llama3:70b-text-q8_0", expected: "Llama3 70B Text Q8_0"},
		{input: "llama3:70b-text-fp16", expected: "Llama3 70B Text FP16"},
		{input: "llava:latest", expected: "Llava (latest)"},
		{input: "llava", expected: "Llava"},
		{input: "llava:v1.6", expected: "Llava v1.6"},
		{input: "llava:7b", expected: "Llava 7B"},
		{input: "llava:13b", expected: "Llava 13B"},
		{input: "llava:34b", expected: "Llava 34B"},
		{input: "llava:7b-v1.5-q2_K", expected: "Llava 7B v1.5 Q2_K"},
		{input: "llava:7b-v1.5-q3_K_S", expected: "Llava 7B v1.5 Q3 K_S"},
		{input: "llava:7b-v1.5-q3_K_M", expected: "Llava 7B v1.5 Q3 K_M"},
		{input: "llava:7b-v1.5-q3_K_L", expected: "Llava 7B v1.5 Q3 K_L"},
		{input: "llava:7b-v1.5-q4_0", expected: "Llava 7B v1.5 Q4_0"},
		{input: "llava:7b-v1.5-q4_1", expected: "Llava 7B v1.5 Q4 1"},
		{input: "llava:7b-v1.5-q4_K_S", expected: "Llava 7B v1.5 Q4 K_S"},
		{input: "llava:7b-v1.5-q4_K_M", expected: "Llava 7B v1.5 Q4_K_M"},
		{input: "llava:7b-v1.5-q5_0", expected: "Llava 7B v1.5 Q5 0"},
		{input: "llava:7b-v1.5-q5_1", expected: "Llava 7B v1.5 Q5 1"},
		{input: "llava:7b-v1.5-q5_K_S", expected: "Llava 7B v1.5 Q5 K_S"},
		{input: "llava:7b-v1.5-q5_K_M", expected: "Llava 7B v1.5 Q5 K_M"},
		{input: "llava:7b-v1.5-q6_K", expected: "Llava 7B v1.5 Q6 K"},
		{input: "llava:7b-v1.5-q8_0", expected: "Llava 7B v1.5 Q8_0"},
		{input: "llava:7b-v1.5-fp16", expected: "Llava 7B v1.5 FP16"},
		{input: "llava:7b-v1.6", expected: "Llava 7B v1.6"},
		{input: "llava:7b-v1.6-mistral-q2_K", expected: "Llava 7B v1.6 Mistral Q2_K"},
		{input: "llava:7b-v1.6-mistral-q3_K_S", expected: "Llava 7B v1.6 Mistral Q3 K_S"},
		{input: "llava:7b-v1.6-mistral-q3_K_M", expected: "Llava 7B v1.6 Mistral Q3 K_M"},
		{input: "llava:7b-v1.6-mistral-q3_K_L", expected: "Llava 7B v1.6 Mistral Q3 K_L"},
		{input: "llava:7b-v1.6-mistral-q4_0", expected: "Llava 7B v1.6 Mistral Q4_0"},
		{input: "llava:7b-v1.6-mistral-q4_1", expected: "Llava 7B v1.6 Mistral Q4 1"},
		{input: "llava:7b-v1.6-mistral-q4_K_S", expected: "Llava 7B v1.6 Mistral Q4 K_S"},
		{input: "llava:7b-v1.6-mistral-q4_K_M", expected: "Llava 7B v1.6 Mistral Q4_K_M"},
		{input: "llava:7b-v1.6-mistral-q5_0", expected: "Llava 7B v1.6 Mistral Q5 0"},
		{input: "llava:7b-v1.6-mistral-q5_1", expected: "Llava 7B v1.6 Mistral Q5 1"},
		{input: "llava:7b-v1.6-mistral-q5_K_S", expected: "Llava 7B v1.6 Mistral Q5 K_S"},
		{input: "llava:7b-v1.6-mistral-q5_K_M", expected: "Llava 7B v1.6 Mistral Q5 K_M"},
		{input: "llava:7b-v1.6-mistral-q6_K", expected: "Llava 7B v1.6 Mistral Q6 K"},
		{input: "llava:7b-v1.6-mistral-q8_0", expected: "Llava 7B v1.6 Mistral Q8_0"},
		{input: "llava:7b-v1.6-mistral-fp16", expected: "Llava 7B v1.6 Mistral FP16"},
		{input: "llava:7b-v1.6-vicuna-q2_K", expected: "Llava 7B v1.6 Vicuna Q2_K"},
		{input: "llava:7b-v1.6-vicuna-q3_K_S", expected: "Llava 7B v1.6 Vicuna Q3 K_S"},
		{input: "llava:7b-v1.6-vicuna-q3_K_M", expected: "Llava 7B v1.6 Vicuna Q3 K_M"},
		{input: "llava:7b-v1.6-vicuna-q3_K_L", expected: "Llava 7B v1.6 Vicuna Q3 K_L"},
		{input: "llava:7b-v1.6-vicuna-q4_0", expected: "Llava 7B v1.6 Vicuna Q4_0"},
		{input: "llava:7b-v1.6-vicuna-q4_1", expected: "Llava 7B v1.6 Vicuna Q4 1"},
		{input: "llava:7b-v1.6-vicuna-q4_K_S", expected: "Llava 7B v1.6 Vicuna Q4 K_S"},
		{input: "llava:7b-v1.6-vicuna-q4_K_M", expected: "Llava 7B v1.6 Vicuna Q4_K_M"},
		{input: "llava:7b-v1.6-vicuna-q5_0", expected: "Llava 7B v1.6 Vicuna Q5 0"},
		{input: "llava:7b-v1.6-vicuna-q5_1", expected: "Llava 7B v1.6 Vicuna Q5 1"},
		{input: "llava:7b-v1.6-vicuna-q5_K_S", expected: "Llava 7B v1.6 Vicuna Q5 K_S"},
		{input: "llava:7b-v1.6-vicuna-q5_K_M", expected: "Llava 7B v1.6 Vicuna Q5 K_M"},
		{input: "llava:7b-v1.6-vicuna-q6_K", expected: "Llava 7B v1.6 Vicuna Q6 K"},
		{input: "llava:7b-v1.6-vicuna-q8_0", expected: "Llava 7B v1.6 Vicuna Q8_0"},
		{input: "llava:7b-v1.6-vicuna-fp16", expected: "Llava 7B v1.6 Vicuna FP16"},
		{input: "llava:13b-v1.5-q2_K", expected: "Llava 13B v1.5 Q2_K"},
		{input: "llava:13b-v1.5-q3_K_S", expected: "Llava 13B v1.5 Q3 K_S"},
		{input: "llava:13b-v1.5-q3_K_M", expected: "Llava 13B v1.5 Q3 K_M"},
		{input: "llava:13b-v1.5-q3_K_L", expected: "Llava 13B v1.5 Q3 K_L"},
		{input: "llava:13b-v1.5-q4_0", expected: "Llava 13B v1.5 Q4_0"},
		{input: "llava:13b-v1.5-q4_1", expected: "Llava 13B v1.5 Q4 1"},
		{input: "llava:13b-v1.5-q4_K_S", expected: "Llava 13B v1.5 Q4 K_S"},
		{input: "llava:13b-v1.5-q4_K_M", expected: "Llava 13B v1.5 Q4_K_M"},
		{input: "llava:13b-v1.5-q5_0", expected: "Llava 13B v1.5 Q5 0"},
		{input: "llava:13b-v1.5-q5_1", expected: "Llava 13B v1.5 Q5 1"},
		{input: "llava:13b-v1.5-q5_K_S", expected: "Llava 13B v1.5 Q5 K_S"},
		{input: "llava:13b-v1.5-q5_K_M", expected: "Llava 13B v1.5 Q5 K_M"},
		{input: "llava:13b-v1.5-q6_K", expected: "Llava 13B v1.5 Q6 K"},
		{input: "llava:13b-v1.5-q8_0", expected: "Llava 13B v1.5 Q8_0"},
		{input: "llava:13b-v1.5-fp16", expected: "Llava 13B v1.5 FP16"},
		{input: "llava:13b-v1.6", expected: "Llava 13B v1.6"},
		{input: "llava:13b-v1.6-vicuna-q2_K", expected: "Llava 13B v1.6 Vicuna Q2_K"},
		{input: "llava:13b-v1.6-vicuna-q3_K_S", expected: "Llava 13B v1.6 Vicuna Q3 K_S"},
		{input: "llava:13b-v1.6-vicuna-q3_K_M", expected: "Llava 13B v1.6 Vicuna Q3 K_M"},
		{input: "llava:13b-v1.6-vicuna-q3_K_L", expected: "Llava 13B v1.6 Vicuna Q3 K_L"},
		{input: "llava:13b-v1.6-vicuna-q4_0", expected: "Llava 13B v1.6 Vicuna Q4_0"},
		{input: "llava:13b-v1.6-vicuna-q4_1", expected: "Llava 13B v1.6 Vicuna Q4 1"},
		{input: "llava:13b-v1.6-vicuna-q4_K_S", expected: "Llava 13B v1.6 Vicuna Q4 K_S"},
		{input: "llava:13b-v1.6-vicuna-q4_K_M", expected: "Llava 13B v1.6 Vicuna Q4_K_M"},
		{input: "llava:13b-v1.6-vicuna-q5_0", expected: "Llava 13B v1.6 Vicuna Q5 0"},
		{input: "llava:13b-v1.6-vicuna-q5_1", expected: "Llava 13B v1.6 Vicuna Q5 1"},
		{input: "llava:13b-v1.6-vicuna-q5_K_S", expected: "Llava 13B v1.6 Vicuna Q5 K_S"},
		{input: "llava:13b-v1.6-vicuna-q5_K_M", expected: "Llava 13B v1.6 Vicuna Q5 K_M"},
		{input: "llava:13b-v1.6-vicuna-q6_K", expected: "Llava 13B v1.6 Vicuna Q6 K"},
		{input: "llava:13b-v1.6-vicuna-q8_0", expected: "Llava 13B v1.6 Vicuna Q8_0"},
		{input: "llava:13b-v1.6-vicuna-fp16", expected: "Llava 13B v1.6 Vicuna FP16"},
		{input: "llava:34b-v1.6", expected: "Llava 34B v1.6"},
		{input: "llava:34b-v1.6-q2_K", expected: "Llava 34B v1.6 Q2_K"},
		{input: "llava:34b-v1.6-q3_K_S", expected: "Llava 34B v1.6 Q3 K_S"},
		{input: "llava:34b-v1.6-q3_K_M", expected: "Llava 34B v1.6 Q3 K_M"},
		{input: "llava:34b-v1.6-q3_K_L", expected: "Llava 34B v1.6 Q3 K_L"},
		{input: "llava:34b-v1.6-q4_0", expected: "Llava 34B v1.6 Q4_0"},
		{input: "llava:34b-v1.6-q4_1", expected: "Llava 34B v1.6 Q4 1"},
		{input: "llava:34b-v1.6-q4_K_S", expected: "Llava 34B v1.6 Q4 K_S"},
		{input: "llava:34b-v1.6-q4_K_M", expected: "Llava 34B v1.6 Q4_K_M"},
		{input: "llava:34b-v1.6-q5_0", expected: "Llava 34B v1.6 Q5 0"},
		{input: "llava:34b-v1.6-q5_1", expected: "Llava 34B v1.6 Q5 1"},
		{input: "llava:34b-v1.6-q5_K_S", expected: "Llava 34B v1.6 Q5 K_S"},
		{input: "llava:34b-v1.6-q5_K_M", expected: "Llava 34B v1.6 Q5 K_M"},
		{input: "llava:34b-v1.6-q6_K", expected: "Llava 34B v1.6 Q6 K"},
		{input: "llava:34b-v1.6-q8_0", expected: "Llava 34B v1.6 Q8_0"},
		{input: "llava:34b-v1.6-fp16", expected: "Llava 34B v1.6 FP16"},
		{input: "gemma2:latest", expected: "Gemma2 (latest)"},
		{input: "gemma2", expected: "Gemma2"},
		{input: "gemma2:2b", expected: "Gemma2 2B"},
		{input: "gemma2:9b", expected: "Gemma2 9B"},
		{input: "gemma2:27b", expected: "Gemma2 27B"},
		{input: "gemma2:2b-instruct-q2_K", expected: "Gemma2 2B Instruct Q2_K"},
		{input: "gemma2:2b-instruct-q3_K_S", expected: "Gemma2 2B Instruct Q3 K_S"},
		{input: "gemma2:2b-instruct-q3_K_M", expected: "Gemma2 2B Instruct Q3 K_M"},
		{input: "gemma2:2b-instruct-q3_K_L", expected: "Gemma2 2B Instruct Q3 K_L"},
		{input: "gemma2:2b-instruct-q4_0", expected: "Gemma2 2B Instruct Q4_0"},
		{input: "gemma2:2b-instruct-q4_1", expected: "Gemma2 2B Instruct Q4 1"},
		{input: "gemma2:2b-instruct-q4_K_S", expected: "Gemma2 2B Instruct Q4 K_S"},
		{input: "gemma2:2b-instruct-q4_K_M", expected: "Gemma2 2B Instruct Q4_K_M"},
		{input: "gemma2:2b-instruct-q5_0", expected: "Gemma2 2B Instruct Q5 0"},
		{input: "gemma2:2b-instruct-q5_1", expected: "Gemma2 2B Instruct Q5 1"},
		{input: "gemma2:2b-instruct-q5_K_S", expected: "Gemma2 2B Instruct Q5 K_S"},
		{input: "gemma2:2b-instruct-q5_K_M", expected: "Gemma2 2B Instruct Q5 K_M"},
		{input: "gemma2:2b-instruct-q6_K", expected: "Gemma2 2B Instruct Q6 K"},
		{input: "gemma2:2b-instruct-q8_0", expected: "Gemma2 2B Instruct Q8_0"},
		{input: "gemma2:2b-instruct-fp16", expected: "Gemma2 2B Instruct FP16"},
		{input: "gemma2:2b-text-q2_K", expected: "Gemma2 2B Text Q2_K"},
		{input: "gemma2:2b-text-q3_K_S", expected: "Gemma2 2B Text Q3 K_S"},
		{input: "gemma2:2b-text-q3_K_M", expected: "Gemma2 2B Text Q3 K_M"},
		{input: "gemma2:2b-text-q3_K_L", expected: "Gemma2 2B Text Q3 K_L"},
		{input: "gemma2:2b-text-q4_0", expected: "Gemma2 2B Text Q4_0"},
		{input: "gemma2:2b-text-q4_1", expected: "Gemma2 2B Text Q4 1"},
		{input: "gemma2:2b-text-q4_K_S", expected: "Gemma2 2B Text Q4 K_S"},
		{input: "gemma2:2b-text-q4_K_M", expected: "Gemma2 2B Text Q4_K_M"},
		{input: "gemma2:2b-text-q5_0", expected: "Gemma2 2B Text Q5 0"},
		{input: "gemma2:2b-text-q5_1", expected: "Gemma2 2B Text Q5 1"},
		{input: "gemma2:2b-text-q5_K_S", expected: "Gemma2 2B Text Q5 K_S"},
		{input: "gemma2:2b-text-q5_K_M", expected: "Gemma2 2B Text Q5 K_M"},
		{input: "gemma2:2b-text-q6_K", expected: "Gemma2 2B Text Q6 K"},
		{input: "gemma2:2b-text-q8_0", expected: "Gemma2 2B Text Q8_0"},
		{input: "gemma2:2b-text-fp16", expected: "Gemma2 2B Text FP16"},
		{input: "gemma2:9b-instruct-q2_K", expected: "Gemma2 9B Instruct Q2_K"},
		{input: "gemma2:9b-instruct-q3_K_S", expected: "Gemma2 9B Instruct Q3 K_S"},
		{input: "gemma2:9b-instruct-q3_K_M", expected: "Gemma2 9B Instruct Q3 K_M"},
		{input: "gemma2:9b-instruct-q3_K_L", expected: "Gemma2 9B Instruct Q3 K_L"},
		{input: "gemma2:9b-instruct-q4_0", expected: "Gemma2 9B Instruct Q4_0"},
		{input: "gemma2:9b-instruct-q4_1", expected: "Gemma2 9B Instruct Q4 1"},
		{input: "gemma2:9b-instruct-q4_K_S", expected: "Gemma2 9B Instruct Q4 K_S"},
		{input: "gemma2:9b-instruct-q4_K_M", expected: "Gemma2 9B Instruct Q4_K_M"},
		{input: "gemma2:9b-instruct-q5_0", expected: "Gemma2 9B Instruct Q5 0"},
		{input: "gemma2:9b-instruct-q5_1", expected: "Gemma2 9B Instruct Q5 1"},
		{input: "gemma2:9b-instruct-q5_K_S", expected: "Gemma2 9B Instruct Q5 K_S"},
		{input: "gemma2:9b-instruct-q5_K_M", expected: "Gemma2 9B Instruct Q5 K_M"},
		{input: "gemma2:9b-instruct-q6_K", expected: "Gemma2 9B Instruct Q6 K"},
		{input: "gemma2:9b-instruct-q8_0", expected: "Gemma2 9B Instruct Q8_0"},
		{input: "gemma2:9b-instruct-fp16", expected: "Gemma2 9B Instruct FP16"},
		{input: "gemma2:9b-text-q2_K", expected: "Gemma2 9B Text Q2_K"},
		{input: "gemma2:9b-text-q3_K_S", expected: "Gemma2 9B Text Q3 K_S"},
		{input: "gemma2:9b-text-q3_K_M", expected: "Gemma2 9B Text Q3 K_M"},
		{input: "gemma2:9b-text-q3_K_L", expected: "Gemma2 9B Text Q3 K_L"},
		{input: "gemma2:9b-text-q4_0", expected: "Gemma2 9B Text Q4_0"},
		{input: "gemma2:9b-text-q4_1", expected: "Gemma2 9B Text Q4 1"},
		{input: "gemma2:9b-text-q4_K_S", expected: "Gemma2 9B Text Q4 K_S"},
		{input: "gemma2:9b-text-q4_K_M", expected: "Gemma2 9B Text Q4_K_M"},
		{input: "gemma2:9b-text-q5_0", expected: "Gemma2 9B Text Q5 0"},
		{input: "gemma2:9b-text-q5_1", expected: "Gemma2 9B Text Q5 1"},
		{input: "gemma2:9b-text-q5_K_S", expected: "Gemma2 9B Text Q5 K_S"},
		{input: "gemma2:9b-text-q5_K_M", expected: "Gemma2 9B Text Q5 K_M"},
		{input: "gemma2:9b-text-q6_K", expected: "Gemma2 9B Text Q6 K"},
		{input: "gemma2:9b-text-q8_0", expected: "Gemma2 9B Text Q8_0"},
		{input: "gemma2:9b-text-fp16", expected: "Gemma2 9B Text FP16"},
		{input: "gemma2:27b-instruct-q2_K", expected: "Gemma2 27B Instruct Q2_K"},
		{input: "gemma2:27b-instruct-q3_K_S", expected: "Gemma2 27B Instruct Q3 K_S"},
		{input: "gemma2:27b-instruct-q3_K_M", expected: "Gemma2 27B Instruct Q3 K_M"},
		{input: "gemma2:27b-instruct-q3_K_L", expected: "Gemma2 27B Instruct Q3 K_L"},
		{input: "gemma2:27b-instruct-q4_0", expected: "Gemma2 27B Instruct Q4_0"},
		{input: "gemma2:27b-instruct-q4_1", expected: "Gemma2 27B Instruct Q4 1"},
		{input: "gemma2:27b-instruct-q4_K_S", expected: "Gemma2 27B Instruct Q4 K_S"},
		{input: "gemma2:27b-instruct-q4_K_M", expected: "Gemma2 27B Instruct Q4_K_M"},
		{input: "gemma2:27b-instruct-q5_0", expected: "Gemma2 27B Instruct Q5 0"},
		{input: "gemma2:27b-instruct-q5_1", expected: "Gemma2 27B Instruct Q5 1"},
		{input: "gemma2:27b-instruct-q5_K_S", expected: "Gemma2 27B Instruct Q5 K_S"},
		{input: "gemma2:27b-instruct-q5_K_M", expected: "Gemma2 27B Instruct Q5 K_M"},
		{input: "gemma2:27b-instruct-q6_K", expected: "Gemma2 27B Instruct Q6 K"},
		{input: "gemma2:27b-instruct-q8_0", expected: "Gemma2 27B Instruct Q8_0"},
		{input: "gemma2:27b-instruct-fp16", expected: "Gemma2 27B Instruct FP16"},
		{input: "gemma2:27b-text-q2_K", expected: "Gemma2 27B Text Q2_K"},
		{input: "gemma2:27b-text-q3_K_S", expected: "Gemma2 27B Text Q3 K_S"},
		{input: "gemma2:27b-text-q3_K_M", expected: "Gemma2 27B Text Q3 K_M"},
		{input: "gemma2:27b-text-q3_K_L", expected: "Gemma2 27B Text Q3 K_L"},
		{input: "gemma2:27b-text-q4_0", expected: "Gemma2 27B Text Q4_0"},
		{input: "gemma2:27b-text-q4_1", expected: "Gemma2 27B Text Q4 1"},
		{input: "gemma2:27b-text-q4_K_S", expected: "Gemma2 27B Text Q4 K_S"},
		{input: "gemma2:27b-text-q4_K_M", expected: "Gemma2 27B Text Q4_K_M"},
		{input: "gemma2:27b-text-q5_0", expected: "Gemma2 27B Text Q5 0"},
		{input: "gemma2:27b-text-q5_1", expected: "Gemma2 27B Text Q5 1"},
		{input: "gemma2:27b-text-q5_K_S", expected: "Gemma2 27B Text Q5 K_S"},
		{input: "gemma2:27b-text-q5_K_M", expected: "Gemma2 27B Text Q5 K_M"},
		{input: "gemma2:27b-text-q6_K", expected: "Gemma2 27B Text Q6 K"},
		{input: "gemma2:27b-text-q8_0", expected: "Gemma2 27B Text Q8_0"},
		{input: "gemma2:27b-text-fp16", expected: "Gemma2 27B Text FP16"},
		{input: "qwen2.5-coder:latest", expected: "Qwen2.5 Coder (latest)"},
		{input: "qwen2.5-coder", expected: "Qwen2.5 Coder"},
		{input: "qwen2.5-coder:0.5b", expected: "Qwen2.5 Coder 0.5B"},
		{input: "qwen2.5-coder:1.5b", expected: "Qwen2.5 Coder 1.5B"},
		{input: "qwen2.5-coder:3b", expected: "Qwen2.5 Coder 3B"},
		{input: "qwen2.5-coder:7b", expected: "Qwen2.5 Coder 7B"},
		{input: "qwen2.5-coder:14b", expected: "Qwen2.5 Coder 14B"},
		{input: "qwen2.5-coder:32b", expected: "Qwen2.5 Coder 32B"},
		{input: "qwen2.5-coder:0.5b-base", expected: "Qwen2.5 Coder 0.5B Base"},
		{input: "qwen2.5-coder:0.5b-base-q2_K", expected: "Qwen2.5 Coder 0.5B Base Q2_K"},
		{input: "qwen2.5-coder:0.5b-base-q3_K_S", expected: "Qwen2.5 Coder 0.5B Base Q3 K_S"},
		{input: "qwen2.5-coder:0.5b-base-q3_K_M", expected: "Qwen2.5 Coder 0.5B Base Q3 K_M"},
		{input: "qwen2.5-coder:0.5b-base-q3_K_L", expected: "Qwen2.5 Coder 0.5B Base Q3 K_L"},
		{input: "qwen2.5-coder:0.5b-base-q4_0", expected: "Qwen2.5 Coder 0.5B Base Q4_0"},
		{input: "qwen2.5-coder:0.5b-base-q4_1", expected: "Qwen2.5 Coder 0.5B Base Q4 1"},
		{input: "qwen2.5-coder:0.5b-base-q4_K_S", expected: "Qwen2.5 Coder 0.5B Base Q4 K_S"},
		{input: "qwen2.5-coder:0.5b-base-q4_K_M", expected: "Qwen2.5 Coder 0.5B Base Q4_K_M"},
		{input: "qwen2.5-coder:0.5b-base-q5_0", expected: "Qwen2.5 Coder 0.5B Base Q5 0"},
		{input: "qwen2.5-coder:0.5b-base-q5_1", expected: "Qwen2.5 Coder 0.5B Base Q5 1"},
		{input: "qwen2.5-coder:0.5b-base-q5_K_S", expected: "Qwen2.5 Coder 0.5B Base Q5 K_S"},
		{input: "qwen2.5-coder:0.5b-base-q5_K_M", expected: "Qwen2.5 Coder 0.5B Base Q5 K_M"},
		{input: "qwen2.5-coder:0.5b-base-q6_K", expected: "Qwen2.5 Coder 0.5B Base Q6 K"},
		{input: "qwen2.5-coder:0.5b-base-q8_0", expected: "Qwen2.5 Coder 0.5B Base Q8_0"},
		{input: "qwen2.5-coder:0.5b-base-fp16", expected: "Qwen2.5 Coder 0.5B Base FP16"},
		{input: "qwen2.5-coder:0.5b-instruct", expected: "Qwen2.5 Coder 0.5B Instruct"},
		{input: "qwen2.5-coder:0.5b-instruct-q2_K", expected: "Qwen2.5 Coder 0.5B Instruct Q2_K"},
		{input: "qwen2.5-coder:0.5b-instruct-q3_K_S", expected: "Qwen2.5 Coder 0.5B Instruct Q3 K_S"},
		{input: "qwen2.5-coder:0.5b-instruct-q3_K_M", expected: "Qwen2.5 Coder 0.5B Instruct Q3 K_M"},
		{input: "qwen2.5-coder:0.5b-instruct-q3_K_L", expected: "Qwen2.5 Coder 0.5B Instruct Q3 K_L"},
		{input: "qwen2.5-coder:0.5b-instruct-q4_0", expected: "Qwen2.5 Coder 0.5B Instruct Q4_0"},
		{input: "qwen2.5-coder:0.5b-instruct-q4_1", expected: "Qwen2.5 Coder 0.5B Instruct Q4 1"},
		{input: "qwen2.5-coder:0.5b-instruct-q4_K_S", expected: "Qwen2.5 Coder 0.5B Instruct Q4 K_S"},
		{input: "qwen2.5-coder:0.5b-instruct-q4_K_M", expected: "Qwen2.5 Coder 0.5B Instruct Q4_K_M"},
		{input: "qwen2.5-coder:0.5b-instruct-q5_0", expected: "Qwen2.5 Coder 0.5B Instruct Q5 0"},
		{input: "qwen2.5-coder:0.5b-instruct-q5_1", expected: "Qwen2.5 Coder 0.5B Instruct Q5 1"},
		{input: "qwen2.5-coder:0.5b-instruct-q5_K_S", expected: "Qwen2.5 Coder 0.5B Instruct Q5 K_S"},
		{input: "qwen2.5-coder:0.5b-instruct-q5_K_M", expected: "Qwen2.5 Coder 0.5B Instruct Q5 K_M"},
		{input: "qwen2.5-coder:0.5b-instruct-q6_K", expected: "Qwen2.5 Coder 0.5B Instruct Q6 K"},
		{input: "qwen2.5-coder:0.5b-instruct-q8_0", expected: "Qwen2.5 Coder 0.5B Instruct Q8_0"},
		{input: "qwen2.5-coder:0.5b-instruct-fp16", expected: "Qwen2.5 Coder 0.5B Instruct FP16"},
		{input: "qwen2.5-coder:1.5b-base", expected: "Qwen2.5 Coder 1.5B Base"},
		{input: "qwen2.5-coder:1.5b-base-q2_K", expected: "Qwen2.5 Coder 1.5B Base Q2_K"},
		{input: "qwen2.5-coder:1.5b-base-q3_K_S", expected: "Qwen2.5 Coder 1.5B Base Q3 K_S"},
		{input: "qwen2.5-coder:1.5b-base-q3_K_M", expected: "Qwen2.5 Coder 1.5B Base Q3 K_M"},
		{input: "qwen2.5-coder:1.5b-base-q3_K_L", expected: "Qwen2.5 Coder 1.5B Base Q3 K_L"},
		{input: "qwen2.5-coder:1.5b-base-q4_0", expected: "Qwen2.5 Coder 1.5B Base Q4_0"},
		{input: "qwen2.5-coder:1.5b-base-q4_1", expected: "Qwen2.5 Coder 1.5B Base Q4 1"},
		{input: "qwen2.5-coder:1.5b-base-q4_K_S", expected: "Qwen2.5 Coder 1.5B Base Q4 K_S"},
		{input: "qwen2.5-coder:1.5b-base-q4_K_M", expected: "Qwen2.5 Coder 1.5B Base Q4_K_M"},
		{input: "qwen2.5-coder:1.5b-base-q5_0", expected: "Qwen2.5 Coder 1.5B Base Q5 0"},
		{input: "qwen2.5-coder:1.5b-base-q5_1", expected: "Qwen2.5 Coder 1.5B Base Q5 1"},
		{input: "qwen2.5-coder:1.5b-base-q5_K_S", expected: "Qwen2.5 Coder 1.5B Base Q5 K_S"},
		{input: "qwen2.5-coder:1.5b-base-q5_K_M", expected: "Qwen2.5 Coder 1.5B Base Q5 K_M"},
		{input: "qwen2.5-coder:1.5b-base-q6_K", expected: "Qwen2.5 Coder 1.5B Base Q6 K"},
		{input: "qwen2.5-coder:1.5b-base-q8_0", expected: "Qwen2.5 Coder 1.5B Base Q8_0"},
		{input: "qwen2.5-coder:1.5b-base-fp16", expected: "Qwen2.5 Coder 1.5B Base FP16"},
		{input: "qwen2.5-coder:1.5b-instruct", expected: "Qwen2.5 Coder 1.5B Instruct"},
		{input: "qwen2.5-coder:1.5b-instruct-q2_K", expected: "Qwen2.5 Coder 1.5B Instruct Q2_K"},
		{input: "qwen2.5-coder:1.5b-instruct-q3_K_S", expected: "Qwen2.5 Coder 1.5B Instruct Q3 K_S"},
		{input: "qwen2.5-coder:1.5b-instruct-q3_K_M", expected: "Qwen2.5 Coder 1.5B Instruct Q3 K_M"},
		{input: "qwen2.5-coder:1.5b-instruct-q3_K_L", expected: "Qwen2.5 Coder 1.5B Instruct Q3 K_L"},
		{input: "qwen2.5-coder:1.5b-instruct-q4_0", expected: "Qwen2.5 Coder 1.5B Instruct Q4_0"},
		{input: "qwen2.5-coder:1.5b-instruct-q4_1", expected: "Qwen2.5 Coder 1.5B Instruct Q4 1"},
		{input: "qwen2.5-coder:1.5b-instruct-q4_K_S", expected: "Qwen2.5 Coder 1.5B Instruct Q4 K_S"},
		{input: "qwen2.5-coder:1.5b-instruct-q4_K_M", expected: "Qwen2.5 Coder 1.5B Instruct Q4_K_M"},
		{input: "qwen2.5-coder:1.5b-instruct-q5_0", expected: "Qwen2.5 Coder 1.5B Instruct Q5 0"},
		{input: "qwen2.5-coder:1.5b-instruct-q5_1", expected: "Qwen2.5 Coder 1.5B Instruct Q5 1"},
		{input: "qwen2.5-coder:1.5b-instruct-q5_K_S", expected: "Qwen2.5 Coder 1.5B Instruct Q5 K_S"},
		{input: "qwen2.5-coder:1.5b-instruct-q5_K_M", expected: "Qwen2.5 Coder 1.5B Instruct Q5 K_M"},
		{input: "qwen2.5-coder:1.5b-instruct-q6_K", expected: "Qwen2.5 Coder 1.5B Instruct Q6 K"},
		{input: "qwen2.5-coder:1.5b-instruct-q8_0", expected: "Qwen2.5 Coder 1.5B Instruct Q8_0"},
		{input: "qwen2.5-coder:1.5b-instruct-fp16", expected: "Qwen2.5 Coder 1.5B Instruct FP16"},
		{input: "qwen2.5-coder:3b-base", expected: "Qwen2.5 Coder 3B Base"},
		{input: "qwen2.5-coder:3b-base-q2_K", expected: "Qwen2.5 Coder 3B Base Q2_K"},
		{input: "qwen2.5-coder:3b-base-q3_K_S", expected: "Qwen2.5 Coder 3B Base Q3 K_S"},
		{input: "qwen2.5-coder:3b-base-q3_K_M", expected: "Qwen2.5 Coder 3B Base Q3 K_M"},
		{input: "qwen2.5-coder:3b-base-q3_K_L", expected: "Qwen2.5 Coder 3B Base Q3 K_L"},
		{input: "qwen2.5-coder:3b-base-q4_0", expected: "Qwen2.5 Coder 3B Base Q4_0"},
		{input: "qwen2.5-coder:3b-base-q4_1", expected: "Qwen2.5 Coder 3B Base Q4 1"},
		{input: "qwen2.5-coder:3b-base-q4_K_S", expected: "Qwen2.5 Coder 3B Base Q4 K_S"},
		{input: "qwen2.5-coder:3b-base-q4_K_M", expected: "Qwen2.5 Coder 3B Base Q4_K_M"},
		{input: "qwen2.5-coder:3b-base-q5_0", expected: "Qwen2.5 Coder 3B Base Q5 0"},
		{input: "qwen2.5-coder:3b-base-q5_1", expected: "Qwen2.5 Coder 3B Base Q5 1"},
		{input: "qwen2.5-coder:3b-base-q5_K_S", expected: "Qwen2.5 Coder 3B Base Q5 K_S"},
		{input: "qwen2.5-coder:3b-base-q5_K_M", expected: "Qwen2.5 Coder 3B Base Q5 K_M"},
		{input: "qwen2.5-coder:3b-base-q6_K", expected: "Qwen2.5 Coder 3B Base Q6 K"},
		{input: "qwen2.5-coder:3b-base-q8_0", expected: "Qwen2.5 Coder 3B Base Q8_0"},
		{input: "qwen2.5-coder:3b-base-fp16", expected: "Qwen2.5 Coder 3B Base FP16"},
		{input: "qwen2.5-coder:3b-instruct", expected: "Qwen2.5 Coder 3B Instruct"},
		{input: "qwen2.5-coder:3b-instruct-q2_K", expected: "Qwen2.5 Coder 3B Instruct Q2_K"},
		{input: "qwen2.5-coder:3b-instruct-q3_K_S", expected: "Qwen2.5 Coder 3B Instruct Q3 K_S"},
		{input: "qwen2.5-coder:3b-instruct-q3_K_M", expected: "Qwen2.5 Coder 3B Instruct Q3 K_M"},
		{input: "qwen2.5-coder:3b-instruct-q3_K_L", expected: "Qwen2.5 Coder 3B Instruct Q3 K_L"},
		{input: "qwen2.5-coder:3b-instruct-q4_0", expected: "Qwen2.5 Coder 3B Instruct Q4_0"},
		{input: "qwen2.5-coder:3b-instruct-q4_1", expected: "Qwen2.5 Coder 3B Instruct Q4 1"},
		{input: "qwen2.5-coder:3b-instruct-q4_K_S", expected: "Qwen2.5 Coder 3B Instruct Q4 K_S"},
		{input: "qwen2.5-coder:3b-instruct-q4_K_M", expected: "Qwen2.5 Coder 3B Instruct Q4_K_M"},
		{input: "qwen2.5-coder:3b-instruct-q5_0", expected: "Qwen2.5 Coder 3B Instruct Q5 0"},
		{input: "qwen2.5-coder:3b-instruct-q5_1", expected: "Qwen2.5 Coder 3B Instruct Q5 1"},
		{input: "qwen2.5-coder:3b-instruct-q5_K_S", expected: "Qwen2.5 Coder 3B Instruct Q5 K_S"},
		{input: "qwen2.5-coder:3b-instruct-q5_K_M", expected: "Qwen2.5 Coder 3B Instruct Q5 K_M"},
		{input: "qwen2.5-coder:3b-instruct-q6_K", expected: "Qwen2.5 Coder 3B Instruct Q6 K"},
		{input: "qwen2.5-coder:3b-instruct-q8_0", expected: "Qwen2.5 Coder 3B Instruct Q8_0"},
		{input: "qwen2.5-coder:3b-instruct-fp16", expected: "Qwen2.5 Coder 3B Instruct FP16"},
		{input: "qwen2.5-coder:7b-base", expected: "Qwen2.5 Coder 7B Base"},
		{input: "qwen2.5-coder:7b-base-q2_K", expected: "Qwen2.5 Coder 7B Base Q2_K"},
		{input: "qwen2.5-coder:7b-base-q3_K_S", expected: "Qwen2.5 Coder 7B Base Q3 K_S"},
		{input: "qwen2.5-coder:7b-base-q3_K_M", expected: "Qwen2.5 Coder 7B Base Q3 K_M"},
		{input: "qwen2.5-coder:7b-base-q3_K_L", expected: "Qwen2.5 Coder 7B Base Q3 K_L"},
		{input: "qwen2.5-coder:7b-base-q4_0", expected: "Qwen2.5 Coder 7B Base Q4_0"},
		{input: "qwen2.5-coder:7b-base-q4_1", expected: "Qwen2.5 Coder 7B Base Q4 1"},
		{input: "qwen2.5-coder:7b-base-q4_K_S", expected: "Qwen2.5 Coder 7B Base Q4 K_S"},
		{input: "qwen2.5-coder:7b-base-q4_K_M", expected: "Qwen2.5 Coder 7B Base Q4_K_M"},
		{input: "qwen2.5-coder:7b-base-q5_0", expected: "Qwen2.5 Coder 7B Base Q5 0"},
		{input: "qwen2.5-coder:7b-base-q5_1", expected: "Qwen2.5 Coder 7B Base Q5 1"},
		{input: "qwen2.5-coder:7b-base-q5_K_S", expected: "Qwen2.5 Coder 7B Base Q5 K_S"},
		{input: "qwen2.5-coder:7b-base-q5_K_M", expected: "Qwen2.5 Coder 7B Base Q5 K_M"},
		{input: "qwen2.5-coder:7b-base-q6_K", expected: "Qwen2.5 Coder 7B Base Q6 K"},
		{input: "qwen2.5-coder:7b-base-q8_0", expected: "Qwen2.5 Coder 7B Base Q8_0"},
		{input: "qwen2.5-coder:7b-base-fp16", expected: "Qwen2.5 Coder 7B Base FP16"},
		{input: "qwen2.5-coder:7b-instruct", expected: "Qwen2.5 Coder 7B Instruct"},
		{input: "qwen2.5-coder:7b-instruct-q2_K", expected: "Qwen2.5 Coder 7B Instruct Q2_K"},
		{input: "qwen2.5-coder:7b-instruct-q3_K_S", expected: "Qwen2.5 Coder 7B Instruct Q3 K_S"},
		{input: "qwen2.5-coder:7b-instruct-q3_K_M", expected: "Qwen2.5 Coder 7B Instruct Q3 K_M"},
		{input: "qwen2.5-coder:7b-instruct-q3_K_L", expected: "Qwen2.5 Coder 7B Instruct Q3 K_L"},
		{input: "qwen2.5-coder:7b-instruct-q4_0", expected: "Qwen2.5 Coder 7B Instruct Q4_0"},
		{input: "qwen2.5-coder:7b-instruct-q4_1", expected: "Qwen2.5 Coder 7B Instruct Q4 1"},
		{input: "qwen2.5-coder:7b-instruct-q4_K_S", expected: "Qwen2.5 Coder 7B Instruct Q4 K_S"},
		{input: "qwen2.5-coder:7b-instruct-q4_K_M", expected: "Qwen2.5 Coder 7B Instruct Q4_K_M"},
		{input: "qwen2.5-coder:7b-instruct-q5_0", expected: "Qwen2.5 Coder 7B Instruct Q5 0"},
		{input: "qwen2.5-coder:7b-instruct-q5_1", expected: "Qwen2.5 Coder 7B Instruct Q5 1"},
		{input: "qwen2.5-coder:7b-instruct-q5_K_S", expected: "Qwen2.5 Coder 7B Instruct Q5 K_S"},
		{input: "qwen2.5-coder:7b-instruct-q5_K_M", expected: "Qwen2.5 Coder 7B Instruct Q5 K_M"},
		{input: "qwen2.5-coder:7b-instruct-q6_K", expected: "Qwen2.5 Coder 7B Instruct Q6 K"},
		{input: "qwen2.5-coder:7b-instruct-q8_0", expected: "Qwen2.5 Coder 7B Instruct Q8_0"},
		{input: "qwen2.5-coder:7b-instruct-fp16", expected: "Qwen2.5 Coder 7B Instruct FP16"},
		{input: "qwen2.5-coder:14b-base", expected: "Qwen2.5 Coder 14B Base"},
		{input: "qwen2.5-coder:14b-base-q2_K", expected: "Qwen2.5 Coder 14B Base Q2_K"},
		{input: "qwen2.5-coder:14b-base-q3_K_S", expected: "Qwen2.5 Coder 14B Base Q3 K_S"},
		{input: "qwen2.5-coder:14b-base-q3_K_M", expected: "Qwen2.5 Coder 14B Base Q3 K_M"},
		{input: "qwen2.5-coder:14b-base-q3_K_L", expected: "Qwen2.5 Coder 14B Base Q3 K_L"},
		{input: "qwen2.5-coder:14b-base-q4_0", expected: "Qwen2.5 Coder 14B Base Q4_0"},
		{input: "qwen2.5-coder:14b-base-q4_1", expected: "Qwen2.5 Coder 14B Base Q4 1"},
		{input: "qwen2.5-coder:14b-base-q4_K_S", expected: "Qwen2.5 Coder 14B Base Q4 K_S"},
		{input: "qwen2.5-coder:14b-base-q4_K_M", expected: "Qwen2.5 Coder 14B Base Q4_K_M"},
		{input: "qwen2.5-coder:14b-base-q5_0", expected: "Qwen2.5 Coder 14B Base Q5 0"},
		{input: "qwen2.5-coder:14b-base-q5_1", expected: "Qwen2.5 Coder 14B Base Q5 1"},
		{input: "qwen2.5-coder:14b-base-q5_K_S", expected: "Qwen2.5 Coder 14B Base Q5 K_S"},
		{input: "qwen2.5-coder:14b-base-q5_K_M", expected: "Qwen2.5 Coder 14B Base Q5 K_M"},
		{input: "qwen2.5-coder:14b-base-q6_K", expected: "Qwen2.5 Coder 14B Base Q6 K"},
		{input: "qwen2.5-coder:14b-base-q8_0", expected: "Qwen2.5 Coder 14B Base Q8_0"},
		{input: "qwen2.5-coder:14b-base-fp16", expected: "Qwen2.5 Coder 14B Base FP16"},
		{input: "qwen2.5-coder:14b-instruct", expected: "Qwen2.5 Coder 14B Instruct"},
		{input: "qwen2.5-coder:14b-instruct-q2_K", expected: "Qwen2.5 Coder 14B Instruct Q2_K"},
		{input: "qwen2.5-coder:14b-instruct-q3_K_S", expected: "Qwen2.5 Coder 14B Instruct Q3 K_S"},
		{input: "qwen2.5-coder:14b-instruct-q3_K_M", expected: "Qwen2.5 Coder 14B Instruct Q3 K_M"},
		{input: "qwen2.5-coder:14b-instruct-q3_K_L", expected: "Qwen2.5 Coder 14B Instruct Q3 K_L"},
		{input: "qwen2.5-coder:14b-instruct-q4_0", expected: "Qwen2.5 Coder 14B Instruct Q4_0"},
		{input: "qwen2.5-coder:14b-instruct-q4_1", expected: "Qwen2.5 Coder 14B Instruct Q4 1"},
		{input: "qwen2.5-coder:14b-instruct-q4_K_S", expected: "Qwen2.5 Coder 14B Instruct Q4 K_S"},
		{input: "qwen2.5-coder:14b-instruct-q4_K_M", expected: "Qwen2.5 Coder 14B Instruct Q4_K_M"},
		{input: "qwen2.5-coder:14b-instruct-q5_0", expected: "Qwen2.5 Coder 14B Instruct Q5 0"},
		{input: "qwen2.5-coder:14b-instruct-q5_1", expected: "Qwen2.5 Coder 14B Instruct Q5 1"},
		{input: "qwen2.5-coder:14b-instruct-q5_K_S", expected: "Qwen2.5 Coder 14B Instruct Q5 K_S"},
		{input: "qwen2.5-coder:14b-instruct-q5_K_M", expected: "Qwen2.5 Coder 14B Instruct Q5 K_M"},
		{input: "qwen2.5-coder:14b-instruct-q6_K", expected: "Qwen2.5 Coder 14B Instruct Q6 K"},
		{input: "qwen2.5-coder:14b-instruct-q8_0", expected: "Qwen2.5 Coder 14B Instruct Q8_0"},
		{input: "qwen2.5-coder:14b-instruct-fp16", expected: "Qwen2.5 Coder 14B Instruct FP16"},
		{input: "qwen2.5-coder:32b-base", expected: "Qwen2.5 Coder 32B Base"},
		{input: "qwen2.5-coder:32b-base-q2_K", expected: "Qwen2.5 Coder 32B Base Q2_K"},
		{input: "qwen2.5-coder:32b-base-q3_K_S", expected: "Qwen2.5 Coder 32B Base Q3 K_S"},
		{input: "qwen2.5-coder:32b-base-q3_K_M", expected: "Qwen2.5 Coder 32B Base Q3 K_M"},
		{input: "qwen2.5-coder:32b-base-q3_K_L", expected: "Qwen2.5 Coder 32B Base Q3 K_L"},
		{input: "qwen2.5-coder:32b-base-q4_0", expected: "Qwen2.5 Coder 32B Base Q4_0"},
		{input: "qwen2.5-coder:32b-base-q4_1", expected: "Qwen2.5 Coder 32B Base Q4 1"},
		{input: "qwen2.5-coder:32b-base-q4_K_S", expected: "Qwen2.5 Coder 32B Base Q4 K_S"},
		{input: "qwen2.5-coder:32b-base-q4_K_M", expected: "Qwen2.5 Coder 32B Base Q4_K_M"},
		{input: "qwen2.5-coder:32b-base-q5_0", expected: "Qwen2.5 Coder 32B Base Q5 0"},
		{input: "qwen2.5-coder:32b-base-q5_1", expected: "Qwen2.5 Coder 32B Base Q5 1"},
		{input: "qwen2.5-coder:32b-base-q5_K_S", expected: "Qwen2.5 Coder 32B Base Q5 K_S"},
		{input: "qwen2.5-coder:32b-base-q5_K_M", expected: "Qwen2.5 Coder 32B Base Q5 K_M"},
		{input: "qwen2.5-coder:32b-base-q6_K", expected: "Qwen2.5 Coder 32B Base Q6 K"},
		{input: "qwen2.5-coder:32b-base-q8_0", expected: "Qwen2.5 Coder 32B Base Q8_0"},
		{input: "qwen2.5-coder:32b-base-fp16", expected: "Qwen2.5 Coder 32B Base FP16"},
		{input: "qwen2.5-coder:32b-instruct", expected: "Qwen2.5 Coder 32B Instruct"},
		{input: "qwen2.5-coder:32b-instruct-q2_K", expected: "Qwen2.5 Coder 32B Instruct Q2_K"},
		{input: "qwen2.5-coder:32b-instruct-q3_K_S", expected: "Qwen2.5 Coder 32B Instruct Q3 K_S"},
		{input: "qwen2.5-coder:32b-instruct-q3_K_M", expected: "Qwen2.5 Coder 32B Instruct Q3 K_M"},
		{input: "qwen2.5-coder:32b-instruct-q3_K_L", expected: "Qwen2.5 Coder 32B Instruct Q3 K_L"},
		{input: "qwen2.5-coder:32b-instruct-q4_0", expected: "Qwen2.5 Coder 32B Instruct Q4_0"},
		{input: "qwen2.5-coder:32b-instruct-q4_1", expected: "Qwen2.5 Coder 32B Instruct Q4 1"},
		{input: "qwen2.5-coder:32b-instruct-q4_K_S", expected: "Qwen2.5 Coder 32B Instruct Q4 K_S"},
		{input: "qwen2.5-coder:32b-instruct-q4_K_M", expected: "Qwen2.5 Coder 32B Instruct Q4_K_M"},
		{input: "qwen2.5-coder:32b-instruct-q5_0", expected: "Qwen2.5 Coder 32B Instruct Q5 0"},
		{input: "qwen2.5-coder:32b-instruct-q5_1", expected: "Qwen2.5 Coder 32B Instruct Q5 1"},
		{input: "qwen2.5-coder:32b-instruct-q5_K_S", expected: "Qwen2.5 Coder 32B Instruct Q5 K_S"},
		{input: "qwen2.5-coder:32b-instruct-q5_K_M", expected: "Qwen2.5 Coder 32B Instruct Q5 K_M"},
		{input: "qwen2.5-coder:32b-instruct-q6_K", expected: "Qwen2.5 Coder 32B Instruct Q6 K"},
		{input: "qwen2.5-coder:32b-instruct-q8_0", expected: "Qwen2.5 Coder 32B Instruct Q8_0"},
		{input: "qwen2.5-coder:32b-instruct-fp16", expected: "Qwen2.5 Coder 32B Instruct FP16"},
		{input: "phi4:latest", expected: "Phi4 (latest)"},
		{input: "phi4", expected: "Phi4"},
		{input: "phi4:14b", expected: "Phi4 14B"},
		{input: "phi4:14b-q4_K_M", expected: "Phi4 14B Q4_K_M"},
		{input: "phi4:14b-q8_0", expected: "Phi4 14B Q8_0"},
		{input: "phi4:14b-fp16", expected: "Phi4 14B FP16"},
		{input: "mxbai-embed-large:latest", expected: "Mxbai Embed Large (latest)"},
		{input: "mxbai-embed-large", expected: "Mxbai Embed Large"},
		{input: "mxbai-embed-large:v1", expected: "Mxbai Embed Large v1"},
		{input: "mxbai-embed-large:335m", expected: "Mxbai Embed Large 335M"},
		{input: "mxbai-embed-large:335m-v1-fp16", expected: "Mxbai Embed Large 335M v1 FP16"},
		{input: "gemma:latest", expected: "Gemma (latest)"},
		{input: "gemma", expected: "Gemma"},
		{input: "gemma:instruct", expected: "Gemma Instruct"},
		{input: "gemma:text", expected: "Gemma Text"},
		{input: "gemma:v1.1", expected: "Gemma v1.1"},
		{input: "gemma:2b", expected: "Gemma 2B"},
		{input: "gemma:7b", expected: "Gemma 7B"},
		{input: "gemma:2b-instruct", expected: "Gemma 2B Instruct"},
		{input: "gemma:2b-instruct-q2_K", expected: "Gemma 2B Instruct Q2_K"},
		{input: "gemma:2b-instruct-v1.1-q2_K", expected: "Gemma 2B Instruct v1.1 Q2_K"},
		{input: "gemma:2b-instruct-q3_K_S", expected: "Gemma 2B Instruct Q3 K_S"},
		{input: "gemma:2b-instruct-v1.1-q3_K_S", expected: "Gemma 2B Instruct v1.1 Q3 K_S"},
		{input: "gemma:2b-instruct-q3_K_M", expected: "Gemma 2B Instruct Q3 K_M"},
		{input: "gemma:2b-instruct-v1.1-q3_K_M", expected: "Gemma 2B Instruct v1.1 Q3 K_M"},
		{input: "gemma:2b-instruct-q3_K_L", expected: "Gemma 2B Instruct Q3 K_L"},
		{input: "gemma:2b-instruct-v1.1-q3_K_L", expected: "Gemma 2B Instruct v1.1 Q3 K_L"},
		{input: "gemma:2b-instruct-q4_0", expected: "Gemma 2B Instruct Q4_0"},
		{input: "gemma:2b-instruct-v1.1-q4_0", expected: "Gemma 2B Instruct v1.1 Q4_0"},
		{input: "gemma:2b-instruct-q4_1", expected: "Gemma 2B Instruct Q4 1"},
		{input: "gemma:2b-instruct-v1.1-q4_1", expected: "Gemma 2B Instruct v1.1 Q4 1"},
		{input: "gemma:2b-instruct-q4_K_S", expected: "Gemma 2B Instruct Q4 K_S"},
		{input: "gemma:2b-instruct-v1.1-q4_K_S", expected: "Gemma 2B Instruct v1.1 Q4 K_S"},
		{input: "gemma:2b-instruct-q4_K_M", expected: "Gemma 2B Instruct Q4_K_M"},
		{input: "gemma:2b-instruct-v1.1-q4_K_M", expected: "Gemma 2B Instruct v1.1 Q4_K_M"},
		{input: "gemma:2b-instruct-q5_0", expected: "Gemma 2B Instruct Q5 0"},
		{input: "gemma:2b-instruct-v1.1-q5_0", expected: "Gemma 2B Instruct v1.1 Q5 0"},
		{input: "gemma:2b-instruct-q5_1", expected: "Gemma 2B Instruct Q5 1"},
		{input: "gemma:2b-instruct-v1.1-q5_1", expected: "Gemma 2B Instruct v1.1 Q5 1"},
		{input: "gemma:2b-instruct-q5_K_S", expected: "Gemma 2B Instruct Q5 K_S"},
		{input: "gemma:2b-instruct-v1.1-q5_K_S", expected: "Gemma 2B Instruct v1.1 Q5 K_S"},
		{input: "gemma:2b-instruct-q5_K_M", expected: "Gemma 2B Instruct Q5 K_M"},
		{input: "gemma:2b-instruct-v1.1-q5_K_M", expected: "Gemma 2B Instruct v1.1 Q5 K_M"},
		{input: "gemma:2b-instruct-q6_K", expected: "Gemma 2B Instruct Q6 K"},
		{input: "gemma:2b-instruct-v1.1-q6_K", expected: "Gemma 2B Instruct v1.1 Q6 K"},
		{input: "gemma:2b-instruct-q8_0", expected: "Gemma 2B Instruct Q8_0"},
		{input: "gemma:2b-instruct-v1.1-q8_0", expected: "Gemma 2B Instruct v1.1 Q8_0"},
		{input: "gemma:2b-instruct-fp16", expected: "Gemma 2B Instruct FP16"},
		{input: "gemma:2b-instruct-v1.1-fp16", expected: "Gemma 2B Instruct v1.1 FP16"},
		{input: "gemma:2b-text", expected: "Gemma 2B Text"},
		{input: "gemma:2b-text-q2_K", expected: "Gemma 2B Text Q2_K"},
		{input: "gemma:2b-text-q3_K_S", expected: "Gemma 2B Text Q3 K_S"},
		{input: "gemma:2b-text-q3_K_M", expected: "Gemma 2B Text Q3 K_M"},
		{input: "gemma:2b-text-q3_K_L", expected: "Gemma 2B Text Q3 K_L"},
		{input: "gemma:2b-text-q4_0", expected: "Gemma 2B Text Q4_0"},
		{input: "gemma:2b-text-q4_1", expected: "Gemma 2B Text Q4 1"},
		{input: "gemma:2b-text-q4_K_S", expected: "Gemma 2B Text Q4 K_S"},
		{input: "gemma:2b-text-q4_K_M", expected: "Gemma 2B Text Q4_K_M"},
		{input: "gemma:2b-text-q5_0", expected: "Gemma 2B Text Q5 0"},
		{input: "gemma:2b-text-q5_1", expected: "Gemma 2B Text Q5 1"},
		{input: "gemma:2b-text-q5_K_S", expected: "Gemma 2B Text Q5 K_S"},
		{input: "gemma:2b-text-q5_K_M", expected: "Gemma 2B Text Q5 K_M"},
		{input: "gemma:2b-text-q6_K", expected: "Gemma 2B Text Q6 K"},
		{input: "gemma:2b-text-q8_0", expected: "Gemma 2B Text Q8_0"},
		{input: "gemma:2b-text-fp16", expected: "Gemma 2B Text FP16"},
		{input: "gemma:2b-v1.1", expected: "Gemma 2B v1.1"},
		{input: "gemma:7b-instruct", expected: "Gemma 7B Instruct"},
		{input: "gemma:7b-instruct-q2_K", expected: "Gemma 7B Instruct Q2_K"},
		{input: "gemma:7b-instruct-v1.1-q2_K", expected: "Gemma 7B Instruct v1.1 Q2_K"},
		{input: "gemma:7b-instruct-q3_K_S", expected: "Gemma 7B Instruct Q3 K_S"},
		{input: "gemma:7b-instruct-v1.1-q3_K_S", expected: "Gemma 7B Instruct v1.1 Q3 K_S"},
		{input: "gemma:7b-instruct-q3_K_M", expected: "Gemma 7B Instruct Q3 K_M"},
		{input: "gemma:7b-instruct-v1.1-q3_K_M", expected: "Gemma 7B Instruct v1.1 Q3 K_M"},
		{input: "gemma:7b-instruct-q3_K_L", expected: "Gemma 7B Instruct Q3 K_L"},
		{input: "gemma:7b-instruct-v1.1-q3_K_L", expected: "Gemma 7B Instruct v1.1 Q3 K_L"},
		{input: "gemma:7b-instruct-q4_0", expected: "Gemma 7B Instruct Q4_0"},
		{input: "gemma:7b-instruct-v1.1-q4_0", expected: "Gemma 7B Instruct v1.1 Q4_0"},
		{input: "gemma:7b-instruct-q4_1", expected: "Gemma 7B Instruct Q4 1"},
		{input: "gemma:7b-instruct-v1.1-q4_1", expected: "Gemma 7B Instruct v1.1 Q4 1"},
		{input: "gemma:7b-instruct-q4_K_S", expected: "Gemma 7B Instruct Q4 K_S"},
		{input: "gemma:7b-instruct-v1.1-q4_K_S", expected: "Gemma 7B Instruct v1.1 Q4 K_S"},
		{input: "gemma:7b-instruct-q4_K_M", expected: "Gemma 7B Instruct Q4_K_M"},
		{input: "gemma:7b-instruct-v1.1-q4_K_M", expected: "Gemma 7B Instruct v1.1 Q4_K_M"},
		{input: "gemma:7b-instruct-q5_0", expected: "Gemma 7B Instruct Q5 0"},
		{input: "gemma:7b-instruct-v1.1-q5_0", expected: "Gemma 7B Instruct v1.1 Q5 0"},
		{input: "gemma:7b-instruct-q5_1", expected: "Gemma 7B Instruct Q5 1"},
		{input: "gemma:7b-instruct-v1.1-q5_1", expected: "Gemma 7B Instruct v1.1 Q5 1"},
		{input: "gemma:7b-instruct-q5_K_S", expected: "Gemma 7B Instruct Q5 K_S"},
		{input: "gemma:7b-instruct-v1.1-q5_K_S", expected: "Gemma 7B Instruct v1.1 Q5 K_S"},
		{input: "gemma:7b-instruct-q5_K_M", expected: "Gemma 7B Instruct Q5 K_M"},
		{input: "gemma:7b-instruct-v1.1-q5_K_M", expected: "Gemma 7B Instruct v1.1 Q5 K_M"},
		{input: "gemma:7b-instruct-q6_K", expected: "Gemma 7B Instruct Q6 K"},
		{input: "gemma:7b-instruct-v1.1-q6_K", expected: "Gemma 7B Instruct v1.1 Q6 K"},
		{input: "gemma:7b-instruct-q8_0", expected: "Gemma 7B Instruct Q8_0"},
		{input: "gemma:7b-instruct-v1.1-q8_0", expected: "Gemma 7B Instruct v1.1 Q8_0"},
		{input: "gemma:7b-instruct-fp16", expected: "Gemma 7B Instruct FP16"},
		{input: "gemma:7b-instruct-v1.1-fp16", expected: "Gemma 7B Instruct v1.1 FP16"},
		{input: "gemma:7b-text", expected: "Gemma 7B Text"},
		{input: "gemma:7b-text-q2_K", expected: "Gemma 7B Text Q2_K"},
		{input: "gemma:7b-text-q3_K_S", expected: "Gemma 7B Text Q3 K_S"},
		{input: "gemma:7b-text-q3_K_M", expected: "Gemma 7B Text Q3 K_M"},
		{input: "gemma:7b-text-q3_K_L", expected: "Gemma 7B Text Q3 K_L"},
		{input: "gemma:7b-text-q4_0", expected: "Gemma 7B Text Q4_0"},
		{input: "gemma:7b-text-q4_1", expected: "Gemma 7B Text Q4 1"},
		{input: "gemma:7b-text-q4_K_S", expected: "Gemma 7B Text Q4 K_S"},
		{input: "gemma:7b-text-q4_K_M", expected: "Gemma 7B Text Q4_K_M"},
		{input: "gemma:7b-text-q5_0", expected: "Gemma 7B Text Q5 0"},
		{input: "gemma:7b-text-q5_1", expected: "Gemma 7B Text Q5 1"},
		{input: "gemma:7b-text-q5_K_S", expected: "Gemma 7B Text Q5 K_S"},
		{input: "gemma:7b-text-q5_K_M", expected: "Gemma 7B Text Q5 K_M"},
		{input: "gemma:7b-text-q6_K", expected: "Gemma 7B Text Q6 K"},
		{input: "gemma:7b-text-q8_0", expected: "Gemma 7B Text Q8_0"},
		{input: "gemma:7b-text-fp16", expected: "Gemma 7B Text FP16"},
		{input: "gemma:7b-v1.1", expected: "Gemma 7B v1.1"},
		{input: "qwen:latest", expected: "Qwen (latest)"},
		{input: "qwen", expected: "Qwen"},
		{input: "qwen:0.5b", expected: "Qwen 0.5B"},
		{input: "qwen:1.8b", expected: "Qwen 1.8B"},
		{input: "qwen:4b", expected: "Qwen 4B"},
		{input: "qwen:7b", expected: "Qwen 7B"},
		{input: "qwen:14b", expected: "Qwen 14B"},
		{input: "qwen:32b", expected: "Qwen 32B"},
		{input: "qwen:72b", expected: "Qwen 72B"},
		{input: "qwen:110b", expected: "Qwen 110B"},
		{input: "qwen:0.5b-chat", expected: "Qwen 0.5B Chat"},
		{input: "qwen:0.5b-chat-v1.5-q2_K", expected: "Qwen 0.5B Chat v1.5 Q2_K"},
		{input: "qwen:0.5b-chat-v1.5-q3_K_S", expected: "Qwen 0.5B Chat v1.5 Q3 K_S"},
		{input: "qwen:0.5b-chat-v1.5-q3_K_M", expected: "Qwen 0.5B Chat v1.5 Q3 K_M"},
		{input: "qwen:0.5b-chat-v1.5-q3_K_L", expected: "Qwen 0.5B Chat v1.5 Q3 K_L"},
		{input: "qwen:0.5b-chat-v1.5-q4_0", expected: "Qwen 0.5B Chat v1.5 Q4_0"},
		{input: "qwen:0.5b-chat-v1.5-q4_1", expected: "Qwen 0.5B Chat v1.5 Q4 1"},
		{input: "qwen:0.5b-chat-v1.5-q4_K_S", expected: "Qwen 0.5B Chat v1.5 Q4 K_S"},
		{input: "qwen:0.5b-chat-v1.5-q4_K_M", expected: "Qwen 0.5B Chat v1.5 Q4_K_M"},
		{input: "qwen:0.5b-chat-v1.5-q5_0", expected: "Qwen 0.5B Chat v1.5 Q5 0"},
		{input: "qwen:0.5b-chat-v1.5-q5_1", expected: "Qwen 0.5B Chat v1.5 Q5 1"},
		{input: "qwen:0.5b-chat-v1.5-q5_K_S", expected: "Qwen 0.5B Chat v1.5 Q5 K_S"},
		{input: "qwen:0.5b-chat-v1.5-q5_K_M", expected: "Qwen 0.5B Chat v1.5 Q5 K_M"},
		{input: "qwen:0.5b-chat-v1.5-q6_K", expected: "Qwen 0.5B Chat v1.5 Q6 K"},
		{input: "qwen:0.5b-chat-v1.5-q8_0", expected: "Qwen 0.5B Chat v1.5 Q8_0"},
		{input: "qwen:0.5b-chat-v1.5-fp16", expected: "Qwen 0.5B Chat v1.5 FP16"},
		{input: "qwen:0.5b-text", expected: "Qwen 0.5B Text"},
		{input: "qwen:0.5b-text-v1.5-q2_K", expected: "Qwen 0.5B Text v1.5 Q2_K"},
		{input: "qwen:0.5b-text-v1.5-q3_K_S", expected: "Qwen 0.5B Text v1.5 Q3 K_S"},
		{input: "qwen:0.5b-text-v1.5-q3_K_M", expected: "Qwen 0.5B Text v1.5 Q3 K_M"},
		{input: "qwen:0.5b-text-v1.5-q3_K_L", expected: "Qwen 0.5B Text v1.5 Q3 K_L"},
		{input: "qwen:0.5b-text-v1.5-q4_0", expected: "Qwen 0.5B Text v1.5 Q4_0"},
		{input: "qwen:0.5b-text-v1.5-q4_1", expected: "Qwen 0.5B Text v1.5 Q4 1"},
		{input: "qwen:0.5b-text-v1.5-q4_K_S", expected: "Qwen 0.5B Text v1.5 Q4 K_S"},
		{input: "qwen:0.5b-text-v1.5-q4_K_M", expected: "Qwen 0.5B Text v1.5 Q4_K_M"},
		{input: "qwen:0.5b-text-v1.5-q5_0", expected: "Qwen 0.5B Text v1.5 Q5 0"},
		{input: "qwen:0.5b-text-v1.5-q5_1", expected: "Qwen 0.5B Text v1.5 Q5 1"},
		{input: "qwen:0.5b-text-v1.5-q5_K_S", expected: "Qwen 0.5B Text v1.5 Q5 K_S"},
		{input: "qwen:0.5b-text-v1.5-q5_K_M", expected: "Qwen 0.5B Text v1.5 Q5 K_M"},
		{input: "qwen:0.5b-text-v1.5-q6_K", expected: "Qwen 0.5B Text v1.5 Q6 K"},
		{input: "qwen:0.5b-text-v1.5-q8_0", expected: "Qwen 0.5B Text v1.5 Q8_0"},
		{input: "qwen:0.5b-text-v1.5-fp16", expected: "Qwen 0.5B Text v1.5 FP16"},
		{input: "qwen:1.8b-chat", expected: "Qwen 1.8B Chat"},
		{input: "qwen:1.8b-chat-q2_K", expected: "Qwen 1.8B Chat Q2_K"},
		{input: "qwen:1.8b-chat-v1.5-q2_K", expected: "Qwen 1.8B Chat v1.5 Q2_K"},
		{input: "qwen:1.8b-chat-q3_K_S", expected: "Qwen 1.8B Chat Q3 K_S"},
		{input: "qwen:1.8b-chat-v1.5-q3_K_S", expected: "Qwen 1.8B Chat v1.5 Q3 K_S"},
		{input: "qwen:1.8b-chat-q3_K_M", expected: "Qwen 1.8B Chat Q3 K_M"},
		{input: "qwen:1.8b-chat-v1.5-q3_K_M", expected: "Qwen 1.8B Chat v1.5 Q3 K_M"},
		{input: "qwen:1.8b-chat-q3_K_L", expected: "Qwen 1.8B Chat Q3 K_L"},
		{input: "qwen:1.8b-chat-v1.5-q3_K_L", expected: "Qwen 1.8B Chat v1.5 Q3 K_L"},
		{input: "qwen:1.8b-chat-q4_0", expected: "Qwen 1.8B Chat Q4_0"},
		{input: "qwen:1.8b-chat-v1.5-q4_0", expected: "Qwen 1.8B Chat v1.5 Q4_0"},
		{input: "qwen:1.8b-chat-q4_1", expected: "Qwen 1.8B Chat Q4 1"},
		{input: "qwen:1.8b-chat-v1.5-q4_1", expected: "Qwen 1.8B Chat v1.5 Q4 1"},
		{input: "qwen:1.8b-chat-q4_K_S", expected: "Qwen 1.8B Chat Q4 K_S"},
		{input: "qwen:1.8b-chat-v1.5-q4_K_S", expected: "Qwen 1.8B Chat v1.5 Q4 K_S"},
		{input: "qwen:1.8b-chat-q4_K_M", expected: "Qwen 1.8B Chat Q4_K_M"},
		{input: "qwen:1.8b-chat-v1.5-q4_K_M", expected: "Qwen 1.8B Chat v1.5 Q4_K_M"},
		{input: "qwen:1.8b-chat-q5_0", expected: "Qwen 1.8B Chat Q5 0"},
		{input: "qwen:1.8b-chat-v1.5-q5_0", expected: "Qwen 1.8B Chat v1.5 Q5 0"},
		{input: "qwen:1.8b-chat-q5_1", expected: "Qwen 1.8B Chat Q5 1"},
		{input: "qwen:1.8b-chat-v1.5-q5_1", expected: "Qwen 1.8B Chat v1.5 Q5 1"},
		{input: "qwen:1.8b-chat-q5_K_S", expected: "Qwen 1.8B Chat Q5 K_S"},
		{input: "qwen:1.8b-chat-v1.5-q5_K_S", expected: "Qwen 1.8B Chat v1.5 Q5 K_S"},
		{input: "qwen:1.8b-chat-q5_K_M", expected: "Qwen 1.8B Chat Q5 K_M"},
		{input: "qwen:1.8b-chat-v1.5-q5_K_M", expected: "Qwen 1.8B Chat v1.5 Q5 K_M"},
		{input: "qwen:1.8b-chat-q6_K", expected: "Qwen 1.8B Chat Q6 K"},
		{input: "qwen:1.8b-chat-v1.5-q6_K", expected: "Qwen 1.8B Chat v1.5 Q6 K"},
		{input: "qwen:1.8b-chat-q8_0", expected: "Qwen 1.8B Chat Q8_0"},
		{input: "qwen:1.8b-chat-v1.5-q8_0", expected: "Qwen 1.8B Chat v1.5 Q8_0"},
		{input: "qwen:1.8b-chat-fp16", expected: "Qwen 1.8B Chat FP16"},
		{input: "qwen:1.8b-chat-v1.5-fp16", expected: "Qwen 1.8B Chat v1.5 FP16"},
		{input: "qwen:1.8b-text", expected: "Qwen 1.8B Text"},
		{input: "qwen:1.8b-text-q2_K", expected: "Qwen 1.8B Text Q2_K"},
		{input: "qwen:1.8b-text-v1.5-q2_K", expected: "Qwen 1.8B Text v1.5 Q2_K"},
		{input: "qwen:1.8b-text-q3_K_S", expected: "Qwen 1.8B Text Q3 K_S"},
		{input: "qwen:1.8b-text-v1.5-q3_K_S", expected: "Qwen 1.8B Text v1.5 Q3 K_S"},
		{input: "qwen:1.8b-text-q3_K_M", expected: "Qwen 1.8B Text Q3 K_M"},
		{input: "qwen:1.8b-text-v1.5-q3_K_M", expected: "Qwen 1.8B Text v1.5 Q3 K_M"},
		{input: "qwen:1.8b-text-q3_K_L", expected: "Qwen 1.8B Text Q3 K_L"},
		{input: "qwen:1.8b-text-v1.5-q3_K_L", expected: "Qwen 1.8B Text v1.5 Q3 K_L"},
		{input: "qwen:1.8b-text-q4_0", expected: "Qwen 1.8B Text Q4_0"},
		{input: "qwen:1.8b-text-v1.5-q4_0", expected: "Qwen 1.8B Text v1.5 Q4_0"},
		{input: "qwen:1.8b-text-q4_1", expected: "Qwen 1.8B Text Q4 1"},
		{input: "qwen:1.8b-text-v1.5-q4_1", expected: "Qwen 1.8B Text v1.5 Q4 1"},
		{input: "qwen:1.8b-text-q4_K_S", expected: "Qwen 1.8B Text Q4 K_S"},
		{input: "qwen:1.8b-text-v1.5-q4_K_S", expected: "Qwen 1.8B Text v1.5 Q4 K_S"},
		{input: "qwen:1.8b-text-q4_K_M", expected: "Qwen 1.8B Text Q4_K_M"},
		{input: "qwen:1.8b-text-v1.5-q4_K_M", expected: "Qwen 1.8B Text v1.5 Q4_K_M"},
		{input: "qwen:1.8b-text-q5_0", expected: "Qwen 1.8B Text Q5 0"},
		{input: "qwen:1.8b-text-v1.5-q5_0", expected: "Qwen 1.8B Text v1.5 Q5 0"},
		{input: "qwen:1.8b-text-q5_1", expected: "Qwen 1.8B Text Q5 1"},
		{input: "qwen:1.8b-text-v1.5-q5_1", expected: "Qwen 1.8B Text v1.5 Q5 1"},
		{input: "qwen:1.8b-text-q5_K_S", expected: "Qwen 1.8B Text Q5 K_S"},
		{input: "qwen:1.8b-text-v1.5-q5_K_S", expected: "Qwen 1.8B Text v1.5 Q5 K_S"},
		{input: "qwen:1.8b-text-q5_K_M", expected: "Qwen 1.8B Text Q5 K_M"},
		{input: "qwen:1.8b-text-v1.5-q5_K_M", expected: "Qwen 1.8B Text v1.5 Q5 K_M"},
		{input: "qwen:1.8b-text-q6_K", expected: "Qwen 1.8B Text Q6 K"},
		{input: "qwen:1.8b-text-v1.5-q6_K", expected: "Qwen 1.8B Text v1.5 Q6 K"},
		{input: "qwen:1.8b-text-q8_0", expected: "Qwen 1.8B Text Q8_0"},
		{input: "qwen:1.8b-text-v1.5-q8_0", expected: "Qwen 1.8B Text v1.5 Q8_0"},
		{input: "qwen:1.8b-text-fp16", expected: "Qwen 1.8B Text FP16"},
		{input: "qwen:1.8b-text-v1.5-fp16", expected: "Qwen 1.8B Text v1.5 FP16"},
		{input: "qwen:4b-chat", expected: "Qwen 4B Chat"},
		{input: "qwen:4b-chat-v1.5-q2_K", expected: "Qwen 4B Chat v1.5 Q2_K"},
		{input: "qwen:4b-chat-v1.5-q3_K_S", expected: "Qwen 4B Chat v1.5 Q3 K_S"},
		{input: "qwen:4b-chat-v1.5-q3_K_M", expected: "Qwen 4B Chat v1.5 Q3 K_M"},
		{input: "qwen:4b-chat-v1.5-q3_K_L", expected: "Qwen 4B Chat v1.5 Q3 K_L"},
		{input: "qwen:4b-chat-v1.5-q4_0", expected: "Qwen 4B Chat v1.5 Q4_0"},
		{input: "qwen:4b-chat-v1.5-q4_1", expected: "Qwen 4B Chat v1.5 Q4 1"},
		{input: "qwen:4b-chat-v1.5-q4_K_S", expected: "Qwen 4B Chat v1.5 Q4 K_S"},
		{input: "qwen:4b-chat-v1.5-q4_K_M", expected: "Qwen 4B Chat v1.5 Q4_K_M"},
		{input: "qwen:4b-chat-v1.5-q5_0", expected: "Qwen 4B Chat v1.5 Q5 0"},
		{input: "qwen:4b-chat-v1.5-q5_1", expected: "Qwen 4B Chat v1.5 Q5 1"},
		{input: "qwen:4b-chat-v1.5-q5_K_S", expected: "Qwen 4B Chat v1.5 Q5 K_S"},
		{input: "qwen:4b-chat-v1.5-q5_K_M", expected: "Qwen 4B Chat v1.5 Q5 K_M"},
		{input: "qwen:4b-chat-v1.5-q6_K", expected: "Qwen 4B Chat v1.5 Q6 K"},
		{input: "qwen:4b-chat-v1.5-q8_0", expected: "Qwen 4B Chat v1.5 Q8_0"},
		{input: "qwen:4b-chat-v1.5-fp16", expected: "Qwen 4B Chat v1.5 FP16"},
		{input: "qwen:4b-text", expected: "Qwen 4B Text"},
		{input: "qwen:4b-text-v1.5-q2_K", expected: "Qwen 4B Text v1.5 Q2_K"},
		{input: "qwen:4b-text-v1.5-q3_K_S", expected: "Qwen 4B Text v1.5 Q3 K_S"},
		{input: "qwen:4b-text-v1.5-q3_K_M", expected: "Qwen 4B Text v1.5 Q3 K_M"},
		{input: "qwen:4b-text-v1.5-q3_K_L", expected: "Qwen 4B Text v1.5 Q3 K_L"},
		{input: "qwen:4b-text-v1.5-q4_0", expected: "Qwen 4B Text v1.5 Q4_0"},
		{input: "qwen:4b-text-v1.5-q4_1", expected: "Qwen 4B Text v1.5 Q4 1"},
		{input: "qwen:4b-text-v1.5-q4_K_S", expected: "Qwen 4B Text v1.5 Q4 K_S"},
		{input: "qwen:4b-text-v1.5-q4_K_M", expected: "Qwen 4B Text v1.5 Q4_K_M"},
		{input: "qwen:4b-text-v1.5-q5_0", expected: "Qwen 4B Text v1.5 Q5 0"},
		{input: "qwen:4b-text-v1.5-q5_1", expected: "Qwen 4B Text v1.5 Q5 1"},
		{input: "qwen:4b-text-v1.5-q5_K_S", expected: "Qwen 4B Text v1.5 Q5 K_S"},
		{input: "qwen:4b-text-v1.5-q5_K_M", expected: "Qwen 4B Text v1.5 Q5 K_M"},
		{input: "qwen:4b-text-v1.5-q6_K", expected: "Qwen 4B Text v1.5 Q6 K"},
		{input: "qwen:4b-text-v1.5-q8_0", expected: "Qwen 4B Text v1.5 Q8_0"},
		{input: "qwen:4b-text-v1.5-fp16", expected: "Qwen 4B Text v1.5 FP16"},
		{input: "qwen:7b-chat", expected: "Qwen 7B Chat"},
		{input: "qwen:7b-chat-q2_K", expected: "Qwen 7B Chat Q2_K"},
		{input: "qwen:7b-chat-v1.5-q2_K", expected: "Qwen 7B Chat v1.5 Q2_K"},
		{input: "qwen:7b-chat-q3_K_S", expected: "Qwen 7B Chat Q3 K_S"},
		{input: "qwen:7b-chat-v1.5-q3_K_S", expected: "Qwen 7B Chat v1.5 Q3 K_S"},
		{input: "qwen:7b-chat-q3_K_M", expected: "Qwen 7B Chat Q3 K_M"},
		{input: "qwen:7b-chat-v1.5-q3_K_M", expected: "Qwen 7B Chat v1.5 Q3 K_M"},
		{input: "qwen:7b-chat-q3_K_L", expected: "Qwen 7B Chat Q3 K_L"},
		{input: "qwen:7b-chat-v1.5-q3_K_L", expected: "Qwen 7B Chat v1.5 Q3 K_L"},
		{input: "qwen:7b-chat-q4_0", expected: "Qwen 7B Chat Q4_0"},
		{input: "qwen:7b-chat-v1.5-q4_0", expected: "Qwen 7B Chat v1.5 Q4_0"},
		{input: "qwen:7b-chat-q4_1", expected: "Qwen 7B Chat Q4 1"},
		{input: "qwen:7b-chat-v1.5-q4_1", expected: "Qwen 7B Chat v1.5 Q4 1"},
		{input: "qwen:7b-chat-q4_K_S", expected: "Qwen 7B Chat Q4 K_S"},
		{input: "qwen:7b-chat-v1.5-q4_K_S", expected: "Qwen 7B Chat v1.5 Q4 K_S"},
		{input: "qwen:7b-chat-q4_K_M", expected: "Qwen 7B Chat Q4_K_M"},
		{input: "qwen:7b-chat-v1.5-q4_K_M", expected: "Qwen 7B Chat v1.5 Q4_K_M"},
		{input: "qwen:7b-chat-q5_0", expected: "Qwen 7B Chat Q5 0"},
		{input: "qwen:7b-chat-v1.5-q5_0", expected: "Qwen 7B Chat v1.5 Q5 0"},
		{input: "qwen:7b-chat-q5_1", expected: "Qwen 7B Chat Q5 1"},
		{input: "qwen:7b-chat-v1.5-q5_1", expected: "Qwen 7B Chat v1.5 Q5 1"},
		{input: "qwen:7b-chat-q5_K_S", expected: "Qwen 7B Chat Q5 K_S"},
		{input: "qwen:7b-chat-v1.5-q5_K_S", expected: "Qwen 7B Chat v1.5 Q5 K_S"},
		{input: "qwen:7b-chat-q5_K_M", expected: "Qwen 7B Chat Q5 K_M"},
		{input: "qwen:7b-chat-v1.5-q5_K_M", expected: "Qwen 7B Chat v1.5 Q5 K_M"},
		{input: "qwen:7b-chat-q6_K", expected: "Qwen 7B Chat Q6 K"},
		{input: "qwen:7b-chat-v1.5-q6_K", expected: "Qwen 7B Chat v1.5 Q6 K"},
		{input: "qwen:7b-chat-q8_0", expected: "Qwen 7B Chat Q8_0"},
		{input: "qwen:7b-chat-v1.5-q8_0", expected: "Qwen 7B Chat v1.5 Q8_0"},
		{input: "qwen:7b-chat-fp16", expected: "Qwen 7B Chat FP16"},
		{input: "qwen:7b-chat-v1.5-fp16", expected: "Qwen 7B Chat v1.5 FP16"},
		{input: "qwen:7b-text", expected: "Qwen 7B Text"},
		{input: "qwen:7b-text-v1.5-q2_K", expected: "Qwen 7B Text v1.5 Q2_K"},
		{input: "qwen:7b-text-v1.5-q3_K_S", expected: "Qwen 7B Text v1.5 Q3 K_S"},
		{input: "qwen:7b-text-v1.5-q3_K_M", expected: "Qwen 7B Text v1.5 Q3 K_M"},
		{input: "qwen:7b-text-v1.5-q3_K_L", expected: "Qwen 7B Text v1.5 Q3 K_L"},
		{input: "qwen:7b-text-v1.5-q4_0", expected: "Qwen 7B Text v1.5 Q4_0"},
		{input: "qwen:7b-text-v1.5-q4_1", expected: "Qwen 7B Text v1.5 Q4 1"},
		{input: "qwen:7b-text-v1.5-q4_K_S", expected: "Qwen 7B Text v1.5 Q4 K_S"},
		{input: "qwen:7b-text-v1.5-q4_K_M", expected: "Qwen 7B Text v1.5 Q4_K_M"},
		{input: "qwen:7b-text-v1.5-q5_0", expected: "Qwen 7B Text v1.5 Q5 0"},
		{input: "qwen:7b-text-v1.5-q5_1", expected: "Qwen 7B Text v1.5 Q5 1"},
		{input: "qwen:7b-text-v1.5-q5_K_S", expected: "Qwen 7B Text v1.5 Q5 K_S"},
		{input: "qwen:7b-text-v1.5-q5_K_M", expected: "Qwen 7B Text v1.5 Q5 K_M"},
		{input: "qwen:7b-text-v1.5-q6_K", expected: "Qwen 7B Text v1.5 Q6 K"},
		{input: "qwen:7b-text-v1.5-q8_0", expected: "Qwen 7B Text v1.5 Q8_0"},
		{input: "qwen:7b-text-v1.5-fp16", expected: "Qwen 7B Text v1.5 FP16"},
		{input: "qwen:7b-q2_K", expected: "Qwen 7B Q2_K"},
		{input: "qwen:7b-q3_K_S", expected: "Qwen 7B Q3 K_S"},
		{input: "qwen:7b-q3_K_M", expected: "Qwen 7B Q3 K_M"},
		{input: "qwen:7b-q3_K_L", expected: "Qwen 7B Q3 K_L"},
		{input: "qwen:7b-q4_0", expected: "Qwen 7B Q4_0"},
		{input: "qwen:7b-q4_1", expected: "Qwen 7B Q4 1"},
		{input: "qwen:7b-q4_K_S", expected: "Qwen 7B Q4 K_S"},
		{input: "qwen:7b-q4_K_M", expected: "Qwen 7B Q4_K_M"},
		{input: "qwen:7b-q5_0", expected: "Qwen 7B Q5 0"},
		{input: "qwen:7b-q5_1", expected: "Qwen 7B Q5 1"},
		{input: "qwen:7b-q5_K_S", expected: "Qwen 7B Q5 K_S"},
		{input: "qwen:7b-q5_K_M", expected: "Qwen 7B Q5 K_M"},
		{input: "qwen:7b-q6_K", expected: "Qwen 7B Q6 K"},
		{input: "qwen:7b-q8_0", expected: "Qwen 7B Q8_0"},
		{input: "qwen:7b-fp16", expected: "Qwen 7B FP16"},
		{input: "qwen:14b-chat", expected: "Qwen 14B Chat"},
		{input: "qwen:14b-chat-q2_K", expected: "Qwen 14B Chat Q2_K"},
		{input: "qwen:14b-chat-v1.5-q2_K", expected: "Qwen 14B Chat v1.5 Q2_K"},
		{input: "qwen:14b-chat-q3_K_S", expected: "Qwen 14B Chat Q3 K_S"},
		{input: "qwen:14b-chat-v1.5-q3_K_S", expected: "Qwen 14B Chat v1.5 Q3 K_S"},
		{input: "qwen:14b-chat-q3_K_M", expected: "Qwen 14B Chat Q3 K_M"},
		{input: "qwen:14b-chat-v1.5-q3_K_M", expected: "Qwen 14B Chat v1.5 Q3 K_M"},
		{input: "qwen:14b-chat-q3_K_L", expected: "Qwen 14B Chat Q3 K_L"},
		{input: "qwen:14b-chat-v1.5-q3_K_L", expected: "Qwen 14B Chat v1.5 Q3 K_L"},
		{input: "qwen:14b-chat-q4_0", expected: "Qwen 14B Chat Q4_0"},
		{input: "qwen:14b-chat-v1.5-q4_0", expected: "Qwen 14B Chat v1.5 Q4_0"},
		{input: "qwen:14b-chat-q4_1", expected: "Qwen 14B Chat Q4 1"},
		{input: "qwen:14b-chat-v1.5-q4_1", expected: "Qwen 14B Chat v1.5 Q4 1"},
		{input: "qwen:14b-chat-q4_K_S", expected: "Qwen 14B Chat Q4 K_S"},
		{input: "qwen:14b-chat-v1.5-q4_K_S", expected: "Qwen 14B Chat v1.5 Q4 K_S"},
		{input: "qwen:14b-chat-q4_K_M", expected: "Qwen 14B Chat Q4_K_M"},
		{input: "qwen:14b-chat-v1.5-q4_K_M", expected: "Qwen 14B Chat v1.5 Q4_K_M"},
		{input: "qwen:14b-chat-q5_0", expected: "Qwen 14B Chat Q5 0"},
		{input: "qwen:14b-chat-v1.5-q5_0", expected: "Qwen 14B Chat v1.5 Q5 0"},
		{input: "qwen:14b-chat-q5_1", expected: "Qwen 14B Chat Q5 1"},
		{input: "qwen:14b-chat-v1.5-q5_1", expected: "Qwen 14B Chat v1.5 Q5 1"},
		{input: "qwen:14b-chat-q5_K_S", expected: "Qwen 14B Chat Q5 K_S"},
		{input: "qwen:14b-chat-v1.5-q5_K_S", expected: "Qwen 14B Chat v1.5 Q5 K_S"},
		{input: "qwen:14b-chat-q5_K_M", expected: "Qwen 14B Chat Q5 K_M"},
		{input: "qwen:14b-chat-v1.5-q5_K_M", expected: "Qwen 14B Chat v1.5 Q5 K_M"},
		{input: "qwen:14b-chat-q6_K", expected: "Qwen 14B Chat Q6 K"},
		{input: "qwen:14b-chat-v1.5-q6_K", expected: "Qwen 14B Chat v1.5 Q6 K"},
		{input: "qwen:14b-chat-q8_0", expected: "Qwen 14B Chat Q8_0"},
		{input: "qwen:14b-chat-v1.5-q8_0", expected: "Qwen 14B Chat v1.5 Q8_0"},
		{input: "qwen:14b-chat-fp16", expected: "Qwen 14B Chat FP16"},
		{input: "qwen:14b-chat-v1.5-fp16", expected: "Qwen 14B Chat v1.5 FP16"},
		{input: "qwen:14b-text", expected: "Qwen 14B Text"},
		{input: "qwen:14b-text-q2_K", expected: "Qwen 14B Text Q2_K"},
		{input: "qwen:14b-text-v1.5-q2_K", expected: "Qwen 14B Text v1.5 Q2_K"},
		{input: "qwen:14b-text-q3_K_S", expected: "Qwen 14B Text Q3 K_S"},
		{input: "qwen:14b-text-v1.5-q3_K_S", expected: "Qwen 14B Text v1.5 Q3 K_S"},
		{input: "qwen:14b-text-q3_K_M", expected: "Qwen 14B Text Q3 K_M"},
		{input: "qwen:14b-text-v1.5-q3_K_M", expected: "Qwen 14B Text v1.5 Q3 K_M"},
		{input: "qwen:14b-text-q3_K_L", expected: "Qwen 14B Text Q3 K_L"},
		{input: "qwen:14b-text-v1.5-q3_K_L", expected: "Qwen 14B Text v1.5 Q3 K_L"},
		{input: "qwen:14b-text-q4_0", expected: "Qwen 14B Text Q4_0"},
		{input: "qwen:14b-text-v1.5-q4_0", expected: "Qwen 14B Text v1.5 Q4_0"},
		{input: "qwen:14b-text-q4_1", expected: "Qwen 14B Text Q4 1"},
		{input: "qwen:14b-text-v1.5-q4_1", expected: "Qwen 14B Text v1.5 Q4 1"},
		{input: "qwen:14b-text-q4_K_S", expected: "Qwen 14B Text Q4 K_S"},
		{input: "qwen:14b-text-v1.5-q4_K_S", expected: "Qwen 14B Text v1.5 Q4 K_S"},
		{input: "qwen:14b-text-q4_K_M", expected: "Qwen 14B Text Q4_K_M"},
		{input: "qwen:14b-text-v1.5-q4_K_M", expected: "Qwen 14B Text v1.5 Q4_K_M"},
		{input: "qwen:14b-text-q5_0", expected: "Qwen 14B Text Q5 0"},
		{input: "qwen:14b-text-v1.5-q5_0", expected: "Qwen 14B Text v1.5 Q5 0"},
		{input: "qwen:14b-text-q5_1", expected: "Qwen 14B Text Q5 1"},
		{input: "qwen:14b-text-v1.5-q5_1", expected: "Qwen 14B Text v1.5 Q5 1"},
		{input: "qwen:14b-text-q5_K_S", expected: "Qwen 14B Text Q5 K_S"},
		{input: "qwen:14b-text-v1.5-q5_K_S", expected: "Qwen 14B Text v1.5 Q5 K_S"},
		{input: "qwen:14b-text-q5_K_M", expected: "Qwen 14B Text Q5 K_M"},
		{input: "qwen:14b-text-v1.5-q5_K_M", expected: "Qwen 14B Text v1.5 Q5 K_M"},
		{input: "qwen:14b-text-q6_K", expected: "Qwen 14B Text Q6 K"},
		{input: "qwen:14b-text-v1.5-q6_K", expected: "Qwen 14B Text v1.5 Q6 K"},
		{input: "qwen:14b-text-q8_0", expected: "Qwen 14B Text Q8_0"},
		{input: "qwen:14b-text-v1.5-q8_0", expected: "Qwen 14B Text v1.5 Q8_0"},
		{input: "qwen:14b-text-fp16", expected: "Qwen 14B Text FP16"},
		{input: "qwen:14b-text-v1.5-fp16", expected: "Qwen 14B Text v1.5 FP16"},
		{input: "qwen:32b-chat", expected: "Qwen 32B Chat"},
		{input: "qwen:32b-chat-v1.5-q2_K", expected: "Qwen 32B Chat v1.5 Q2_K"},
		{input: "qwen:32b-chat-v1.5-q3_K_S", expected: "Qwen 32B Chat v1.5 Q3 K_S"},
		{input: "qwen:32b-chat-v1.5-q3_K_M", expected: "Qwen 32B Chat v1.5 Q3 K_M"},
		{input: "qwen:32b-chat-v1.5-q3_K_L", expected: "Qwen 32B Chat v1.5 Q3 K_L"},
		{input: "qwen:32b-chat-v1.5-q4_0", expected: "Qwen 32B Chat v1.5 Q4_0"},
		{input: "qwen:32b-chat-v1.5-q4_1", expected: "Qwen 32B Chat v1.5 Q4 1"},
		{input: "qwen:32b-chat-v1.5-q4_K_S", expected: "Qwen 32B Chat v1.5 Q4 K_S"},
		{input: "qwen:32b-chat-v1.5-q4_K_M", expected: "Qwen 32B Chat v1.5 Q4_K_M"},
		{input: "qwen:32b-chat-v1.5-q5_0", expected: "Qwen 32B Chat v1.5 Q5 0"},
		{input: "qwen:32b-chat-v1.5-q5_1", expected: "Qwen 32B Chat v1.5 Q5 1"},
		{input: "qwen:32b-chat-v1.5-q5_K_S", expected: "Qwen 32B Chat v1.5 Q5 K_S"},
		{input: "qwen:32b-chat-v1.5-q5_K_M", expected: "Qwen 32B Chat v1.5 Q5 K_M"},
		{input: "qwen:32b-chat-v1.5-q6_K", expected: "Qwen 32B Chat v1.5 Q6 K"},
		{input: "qwen:32b-chat-v1.5-q8_0", expected: "Qwen 32B Chat v1.5 Q8_0"},
		{input: "qwen:32b-chat-v1.5-fp16", expected: "Qwen 32B Chat v1.5 FP16"},
		{input: "qwen:32b-text", expected: "Qwen 32B Text"},
		{input: "qwen:32b-text-v1.5-q2_K", expected: "Qwen 32B Text v1.5 Q2_K"},
		{input: "qwen:32b-text-v1.5-q3_K_S", expected: "Qwen 32B Text v1.5 Q3 K_S"},
		{input: "qwen:32b-text-v1.5-q3_K_M", expected: "Qwen 32B Text v1.5 Q3 K_M"},
		{input: "qwen:32b-text-v1.5-q3_K_L", expected: "Qwen 32B Text v1.5 Q3 K_L"},
		{input: "qwen:32b-text-v1.5-q4_0", expected: "Qwen 32B Text v1.5 Q4_0"},
		{input: "qwen:32b-text-v1.5-q4_1", expected: "Qwen 32B Text v1.5 Q4 1"},
		{input: "qwen:32b-text-v1.5-q4_K_S", expected: "Qwen 32B Text v1.5 Q4 K_S"},
		{input: "qwen:32b-text-v1.5-q5_0", expected: "Qwen 32B Text v1.5 Q5 0"},
		{input: "qwen:32b-text-v1.5-q5_1", expected: "Qwen 32B Text v1.5 Q5 1"},
		{input: "qwen:32b-text-v1.5-q8_0", expected: "Qwen 32B Text v1.5 Q8_0"},
		{input: "qwen:72b-chat", expected: "Qwen 72B Chat"},
		{input: "qwen:72b-chat-q2_K", expected: "Qwen 72B Chat Q2_K"},
		{input: "qwen:72b-chat-v1.5-q2_K", expected: "Qwen 72B Chat v1.5 Q2_K"},
		{input: "qwen:72b-chat-q3_K_S", expected: "Qwen 72B Chat Q3 K_S"},
		{input: "qwen:72b-chat-v1.5-q3_K_S", expected: "Qwen 72B Chat v1.5 Q3 K_S"},
		{input: "qwen:72b-chat-q3_K_M", expected: "Qwen 72B Chat Q3 K_M"},
		{input: "qwen:72b-chat-v1.5-q3_K_M", expected: "Qwen 72B Chat v1.5 Q3 K_M"},
		{input: "qwen:72b-chat-q3_K_L", expected: "Qwen 72B Chat Q3 K_L"},
		{input: "qwen:72b-chat-v1.5-q3_K_L", expected: "Qwen 72B Chat v1.5 Q3 K_L"},
		{input: "qwen:72b-chat-q4_0", expected: "Qwen 72B Chat Q4_0"},
		{input: "qwen:72b-chat-v1.5-q4_0", expected: "Qwen 72B Chat v1.5 Q4_0"},
		{input: "qwen:72b-chat-q4_1", expected: "Qwen 72B Chat Q4 1"},
		{input: "qwen:72b-chat-v1.5-q4_1", expected: "Qwen 72B Chat v1.5 Q4 1"},
		{input: "qwen:72b-chat-q4_K_S", expected: "Qwen 72B Chat Q4 K_S"},
		{input: "qwen:72b-chat-v1.5-q4_K_S", expected: "Qwen 72B Chat v1.5 Q4 K_S"},
		{input: "qwen:72b-chat-q4_K_M", expected: "Qwen 72B Chat Q4_K_M"},
		{input: "qwen:72b-chat-v1.5-q4_K_M", expected: "Qwen 72B Chat v1.5 Q4_K_M"},
		{input: "qwen:72b-chat-q5_0", expected: "Qwen 72B Chat Q5 0"},
		{input: "qwen:72b-chat-v1.5-q5_0", expected: "Qwen 72B Chat v1.5 Q5 0"},
		{input: "qwen:72b-chat-q5_1", expected: "Qwen 72B Chat Q5 1"},
		{input: "qwen:72b-chat-v1.5-q5_1", expected: "Qwen 72B Chat v1.5 Q5 1"},
		{input: "qwen:72b-chat-q5_K_S", expected: "Qwen 72B Chat Q5 K_S"},
		{input: "qwen:72b-chat-v1.5-q5_K_S", expected: "Qwen 72B Chat v1.5 Q5 K_S"},
		{input: "qwen:72b-chat-q5_K_M", expected: "Qwen 72B Chat Q5 K_M"},
		{input: "qwen:72b-chat-v1.5-q5_K_M", expected: "Qwen 72B Chat v1.5 Q5 K_M"},
		{input: "qwen:72b-chat-q6_K", expected: "Qwen 72B Chat Q6 K"},
		{input: "qwen:72b-chat-v1.5-q6_K", expected: "Qwen 72B Chat v1.5 Q6 K"},
		{input: "qwen:72b-chat-q8_0", expected: "Qwen 72B Chat Q8_0"},
		{input: "qwen:72b-chat-v1.5-q8_0", expected: "Qwen 72B Chat v1.5 Q8_0"},
		{input: "qwen:72b-chat-fp16", expected: "Qwen 72B Chat FP16"},
		{input: "qwen:72b-chat-v1.5-fp16", expected: "Qwen 72B Chat v1.5 FP16"},
		{input: "qwen:72b-text", expected: "Qwen 72B Text"},
		{input: "qwen:72b-text-q2_K", expected: "Qwen 72B Text Q2_K"},
		{input: "qwen:72b-text-v1.5-q2_K", expected: "Qwen 72B Text v1.5 Q2_K"},
		{input: "qwen:72b-text-q3_K_S", expected: "Qwen 72B Text Q3 K_S"},
		{input: "qwen:72b-text-v1.5-q3_K_S", expected: "Qwen 72B Text v1.5 Q3 K_S"},
		{input: "qwen:72b-text-q3_K_M", expected: "Qwen 72B Text Q3 K_M"},
		{input: "qwen:72b-text-v1.5-q3_K_M", expected: "Qwen 72B Text v1.5 Q3 K_M"},
		{input: "qwen:72b-text-q3_K_L", expected: "Qwen 72B Text Q3 K_L"},
		{input: "qwen:72b-text-v1.5-q3_K_L", expected: "Qwen 72B Text v1.5 Q3 K_L"},
		{input: "qwen:72b-text-q4_0", expected: "Qwen 72B Text Q4_0"},
		{input: "qwen:72b-text-v1.5-q4_0", expected: "Qwen 72B Text v1.5 Q4_0"},
		{input: "qwen:72b-text-q4_1", expected: "Qwen 72B Text Q4 1"},
		{input: "qwen:72b-text-v1.5-q4_1", expected: "Qwen 72B Text v1.5 Q4 1"},
		{input: "qwen:72b-text-q4_K_S", expected: "Qwen 72B Text Q4 K_S"},
		{input: "qwen:72b-text-v1.5-q4_K_S", expected: "Qwen 72B Text v1.5 Q4 K_S"},
		{input: "qwen:72b-text-q4_K_M", expected: "Qwen 72B Text Q4_K_M"},
		{input: "qwen:72b-text-v1.5-q4_K_M", expected: "Qwen 72B Text v1.5 Q4_K_M"},
		{input: "qwen:72b-text-q5_0", expected: "Qwen 72B Text Q5 0"},
		{input: "qwen:72b-text-v1.5-q5_0", expected: "Qwen 72B Text v1.5 Q5 0"},
		{input: "qwen:72b-text-q5_1", expected: "Qwen 72B Text Q5 1"},
		{input: "qwen:72b-text-v1.5-q5_1", expected: "Qwen 72B Text v1.5 Q5 1"},
		{input: "qwen:72b-text-q5_K_S", expected: "Qwen 72B Text Q5 K_S"},
		{input: "qwen:72b-text-v1.5-q5_K_S", expected: "Qwen 72B Text v1.5 Q5 K_S"},
		{input: "qwen:72b-text-q5_K_M", expected: "Qwen 72B Text Q5 K_M"},
		{input: "qwen:72b-text-v1.5-q5_K_M", expected: "Qwen 72B Text v1.5 Q5 K_M"},
		{input: "qwen:72b-text-q6_K", expected: "Qwen 72B Text Q6 K"},
		{input: "qwen:72b-text-v1.5-q6_K", expected: "Qwen 72B Text v1.5 Q6 K"},
		{input: "qwen:72b-text-q8_0", expected: "Qwen 72B Text Q8_0"},
		{input: "qwen:72b-text-v1.5-q8_0", expected: "Qwen 72B Text v1.5 Q8_0"},
		{input: "qwen:72b-text-fp16", expected: "Qwen 72B Text FP16"},
		{input: "qwen:72b-text-v1.5-fp16", expected: "Qwen 72B Text v1.5 FP16"},
		{input: "qwen:110b-chat", expected: "Qwen 110B Chat"},
		{input: "qwen:110b-chat-v1.5-q2_K", expected: "Qwen 110B Chat v1.5 Q2_K"},
		{input: "qwen:110b-chat-v1.5-q3_K_S", expected: "Qwen 110B Chat v1.5 Q3 K_S"},
		{input: "qwen:110b-chat-v1.5-q3_K_M", expected: "Qwen 110B Chat v1.5 Q3 K_M"},
		{input: "qwen:110b-chat-v1.5-q3_K_L", expected: "Qwen 110B Chat v1.5 Q3 K_L"},
		{input: "qwen:110b-chat-v1.5-q4_0", expected: "Qwen 110B Chat v1.5 Q4_0"},
		{input: "qwen:110b-chat-v1.5-q4_1", expected: "Qwen 110B Chat v1.5 Q4 1"},
		{input: "qwen:110b-chat-v1.5-q4_K_S", expected: "Qwen 110B Chat v1.5 Q4 K_S"},
		{input: "qwen:110b-chat-v1.5-q4_K_M", expected: "Qwen 110B Chat v1.5 Q4_K_M"},
		{input: "qwen:110b-chat-v1.5-q5_0", expected: "Qwen 110B Chat v1.5 Q5 0"},
		{input: "qwen:110b-chat-v1.5-q5_1", expected: "Qwen 110B Chat v1.5 Q5 1"},
		{input: "qwen:110b-chat-v1.5-q5_K_S", expected: "Qwen 110B Chat v1.5 Q5 K_S"},
		{input: "qwen:110b-chat-v1.5-q5_K_M", expected: "Qwen 110B Chat v1.5 Q5 K_M"},
		{input: "qwen:110b-chat-v1.5-q6_K", expected: "Qwen 110B Chat v1.5 Q6 K"},
		{input: "qwen:110b-chat-v1.5-q8_0", expected: "Qwen 110B Chat v1.5 Q8_0"},
		{input: "qwen:110b-chat-v1.5-fp16", expected: "Qwen 110B Chat v1.5 FP16"},
		{input: "qwen:110b-text-v1.5-q2_K", expected: "Qwen 110B Text v1.5 Q2_K"},
		{input: "qwen:110b-text-v1.5-q3_K_S", expected: "Qwen 110B Text v1.5 Q3 K_S"},
		{input: "qwen:110b-text-v1.5-q3_K_M", expected: "Qwen 110B Text v1.5 Q3 K_M"},
		{input: "qwen:110b-text-v1.5-q3_K_L", expected: "Qwen 110B Text v1.5 Q3 K_L"},
		{input: "qwen:110b-text-v1.5-q4_0", expected: "Qwen 110B Text v1.5 Q4_0"},
		{input: "qwen:110b-text-v1.5-q4_1", expected: "Qwen 110B Text v1.5 Q4 1"},
		{input: "qwen:110b-text-v1.5-q4_K_S", expected: "Qwen 110B Text v1.5 Q4 K_S"},
		{input: "qwen:110b-text-v1.5-q4_K_M", expected: "Qwen 110B Text v1.5 Q4_K_M"},
		{input: "qwen:110b-text-v1.5-q5_0", expected: "Qwen 110B Text v1.5 Q5 0"},
		{input: "qwen:110b-text-v1.5-q5_1", expected: "Qwen 110B Text v1.5 Q5 1"},
		{input: "qwen:110b-text-v1.5-q5_K_S", expected: "Qwen 110B Text v1.5 Q5 K_S"},
		{input: "qwen:110b-text-v1.5-q5_K_M", expected: "Qwen 110B Text v1.5 Q5 K_M"},
		{input: "qwen:110b-text-v1.5-q6_K", expected: "Qwen 110B Text v1.5 Q6 K"},
		{input: "qwen:110b-text-v1.5-q8_0", expected: "Qwen 110B Text v1.5 Q8_0"},
		{input: "qwen:110b-text-v1.5-fp16", expected: "Qwen 110B Text v1.5 FP16"},
		{input: "llama2:latest", expected: "Llama2 (latest)"},
		{input: "llama2", expected: "Llama2"},
		{input: "llama2:chat", expected: "Llama2 Chat"},
		{input: "llama2:text", expected: "Llama2 Text"},
		{input: "llama2:7b", expected: "Llama2 7B"},
		{input: "llama2:13b", expected: "Llama2 13B"},
		{input: "llama2:70b", expected: "Llama2 70B"},
		{input: "llama2:7b-chat", expected: "Llama2 7B Chat"},
		{input: "llama2:7b-chat-q2_K", expected: "Llama2 7B Chat Q2_K"},
		{input: "llama2:7b-chat-q3_K_S", expected: "Llama2 7B Chat Q3 K_S"},
		{input: "llama2:7b-chat-q3_K_M", expected: "Llama2 7B Chat Q3 K_M"},
		{input: "llama2:7b-chat-q3_K_L", expected: "Llama2 7B Chat Q3 K_L"},
		{input: "llama2:7b-chat-q4_0", expected: "Llama2 7B Chat Q4_0"},
		{input: "llama2:7b-chat-q4_1", expected: "Llama2 7B Chat Q4 1"},
		{input: "llama2:7b-chat-q4_K_S", expected: "Llama2 7B Chat Q4 K_S"},
		{input: "llama2:7b-chat-q4_K_M", expected: "Llama2 7B Chat Q4_K_M"},
		{input: "llama2:7b-chat-q5_0", expected: "Llama2 7B Chat Q5 0"},
		{input: "llama2:7b-chat-q5_1", expected: "Llama2 7B Chat Q5 1"},
		{input: "llama2:7b-chat-q5_K_S", expected: "Llama2 7B Chat Q5 K_S"},
		{input: "llama2:7b-chat-q5_K_M", expected: "Llama2 7B Chat Q5 K_M"},
		{input: "llama2:7b-chat-q6_K", expected: "Llama2 7B Chat Q6 K"},
		{input: "llama2:7b-chat-q8_0", expected: "Llama2 7B Chat Q8_0"},
		{input: "llama2:7b-chat-fp16", expected: "Llama2 7B Chat FP16"},
		{input: "llama2:7b-text", expected: "Llama2 7B Text"},
		{input: "llama2:7b-text-q2_K", expected: "Llama2 7B Text Q2_K"},
		{input: "llama2:7b-text-q3_K_S", expected: "Llama2 7B Text Q3 K_S"},
		{input: "llama2:7b-text-q3_K_M", expected: "Llama2 7B Text Q3 K_M"},
		{input: "llama2:7b-text-q3_K_L", expected: "Llama2 7B Text Q3 K_L"},
		{input: "llama2:7b-text-q4_0", expected: "Llama2 7B Text Q4_0"},
		{input: "llama2:7b-text-q4_1", expected: "Llama2 7B Text Q4 1"},
		{input: "llama2:7b-text-q4_K_S", expected: "Llama2 7B Text Q4 K_S"},
		{input: "llama2:7b-text-q4_K_M", expected: "Llama2 7B Text Q4_K_M"},
		{input: "llama2:7b-text-q5_0", expected: "Llama2 7B Text Q5 0"},
		{input: "llama2:7b-text-q5_1", expected: "Llama2 7B Text Q5 1"},
		{input: "llama2:7b-text-q5_K_S", expected: "Llama2 7B Text Q5 K_S"},
		{input: "llama2:7b-text-q5_K_M", expected: "Llama2 7B Text Q5 K_M"},
		{input: "llama2:7b-text-q6_K", expected: "Llama2 7B Text Q6 K"},
		{input: "llama2:7b-text-q8_0", expected: "Llama2 7B Text Q8_0"},
		{input: "llama2:7b-text-fp16", expected: "Llama2 7B Text FP16"},
		{input: "llama2:13b-chat", expected: "Llama2 13B Chat"},
		{input: "llama2:13b-chat-q2_K", expected: "Llama2 13B Chat Q2_K"},
		{input: "llama2:13b-chat-q3_K_S", expected: "Llama2 13B Chat Q3 K_S"},
		{input: "llama2:13b-chat-q3_K_M", expected: "Llama2 13B Chat Q3 K_M"},
		{input: "llama2:13b-chat-q3_K_L", expected: "Llama2 13B Chat Q3 K_L"},
		{input: "llama2:13b-chat-q4_0", expected: "Llama2 13B Chat Q4_0"},
		{input: "llama2:13b-chat-q4_1", expected: "Llama2 13B Chat Q4 1"},
		{input: "llama2:13b-chat-q4_K_S", expected: "Llama2 13B Chat Q4 K_S"},
		{input: "llama2:13b-chat-q4_K_M", expected: "Llama2 13B Chat Q4_K_M"},
		{input: "llama2:13b-chat-q5_0", expected: "Llama2 13B Chat Q5 0"},
		{input: "llama2:13b-chat-q5_1", expected: "Llama2 13B Chat Q5 1"},
		{input: "llama2:13b-chat-q5_K_S", expected: "Llama2 13B Chat Q5 K_S"},
		{input: "llama2:13b-chat-q5_K_M", expected: "Llama2 13B Chat Q5 K_M"},
		{input: "llama2:13b-chat-q6_K", expected: "Llama2 13B Chat Q6 K"},
		{input: "llama2:13b-chat-q8_0", expected: "Llama2 13B Chat Q8_0"},
		{input: "llama2:13b-chat-fp16", expected: "Llama2 13B Chat FP16"},
		{input: "llama2:13b-text", expected: "Llama2 13B Text"},
		{input: "llama2:13b-text-q2_K", expected: "Llama2 13B Text Q2_K"},
		{input: "llama2:13b-text-q3_K_S", expected: "Llama2 13B Text Q3 K_S"},
		{input: "llama2:13b-text-q3_K_M", expected: "Llama2 13B Text Q3 K_M"},
		{input: "llama2:13b-text-q3_K_L", expected: "Llama2 13B Text Q3 K_L"},
		{input: "llama2:13b-text-q4_0", expected: "Llama2 13B Text Q4_0"},
		{input: "llama2:13b-text-q4_1", expected: "Llama2 13B Text Q4 1"},
		{input: "llama2:13b-text-q4_K_S", expected: "Llama2 13B Text Q4 K_S"},
		{input: "llama2:13b-text-q4_K_M", expected: "Llama2 13B Text Q4_K_M"},
		{input: "llama2:13b-text-q5_0", expected: "Llama2 13B Text Q5 0"},
		{input: "llama2:13b-text-q5_1", expected: "Llama2 13B Text Q5 1"},
		{input: "llama2:13b-text-q5_K_S", expected: "Llama2 13B Text Q5 K_S"},
		{input: "llama2:13b-text-q5_K_M", expected: "Llama2 13B Text Q5 K_M"},
		{input: "llama2:13b-text-q6_K", expected: "Llama2 13B Text Q6 K"},
		{input: "llama2:13b-text-q8_0", expected: "Llama2 13B Text Q8_0"},
		{input: "llama2:13b-text-fp16", expected: "Llama2 13B Text FP16"},
		{input: "llama2:70b-chat", expected: "Llama2 70B Chat"},
		{input: "llama2:70b-chat-q2_K", expected: "Llama2 70B Chat Q2_K"},
		{input: "llama2:70b-chat-q3_K_S", expected: "Llama2 70B Chat Q3 K_S"},
		{input: "llama2:70b-chat-q3_K_M", expected: "Llama2 70B Chat Q3 K_M"},
		{input: "llama2:70b-chat-q3_K_L", expected: "Llama2 70B Chat Q3 K_L"},
		{input: "llama2:70b-chat-q4_0", expected: "Llama2 70B Chat Q4_0"},
		{input: "llama2:70b-chat-q4_1", expected: "Llama2 70B Chat Q4 1"},
		{input: "llama2:70b-chat-q4_K_S", expected: "Llama2 70B Chat Q4 K_S"},
		{input: "llama2:70b-chat-q4_K_M", expected: "Llama2 70B Chat Q4_K_M"},
		{input: "llama2:70b-chat-q5_0", expected: "Llama2 70B Chat Q5 0"},
		{input: "llama2:70b-chat-q5_1", expected: "Llama2 70B Chat Q5 1"},
		{input: "llama2:70b-chat-q5_K_S", expected: "Llama2 70B Chat Q5 K_S"},
		{input: "llama2:70b-chat-q5_K_M", expected: "Llama2 70B Chat Q5 K_M"},
		{input: "llama2:70b-chat-q6_K", expected: "Llama2 70B Chat Q6 K"},
		{input: "llama2:70b-chat-q8_0", expected: "Llama2 70B Chat Q8_0"},
		{input: "llama2:70b-chat-fp16", expected: "Llama2 70B Chat FP16"},
		{input: "llama2:70b-text", expected: "Llama2 70B Text"},
		{input: "llama2:70b-text-q2_K", expected: "Llama2 70B Text Q2_K"},
		{input: "llama2:70b-text-q3_K_S", expected: "Llama2 70B Text Q3 K_S"},
		{input: "llama2:70b-text-q3_K_M", expected: "Llama2 70B Text Q3 K_M"},
		{input: "llama2:70b-text-q3_K_L", expected: "Llama2 70B Text Q3 K_L"},
		{input: "llama2:70b-text-q4_0", expected: "Llama2 70B Text Q4_0"},
		{input: "llama2:70b-text-q4_1", expected: "Llama2 70B Text Q4 1"},
		{input: "llama2:70b-text-q4_K_S", expected: "Llama2 70B Text Q4 K_S"},
		{input: "llama2:70b-text-q4_K_M", expected: "Llama2 70B Text Q4_K_M"},
		{input: "llama2:70b-text-q5_0", expected: "Llama2 70B Text Q5 0"},
		{input: "llama2:70b-text-q5_1", expected: "Llama2 70B Text Q5 1"},
		{input: "llama2:70b-text-q5_K_S", expected: "Llama2 70B Text Q5 K_S"},
		{input: "llama2:70b-text-q5_K_M", expected: "Llama2 70B Text Q5 K_M"},
		{input: "llama2:70b-text-q6_K", expected: "Llama2 70B Text Q6 K"},
		{input: "llama2:70b-text-q8_0", expected: "Llama2 70B Text Q8_0"},
		{input: "llama2:70b-text-fp16", expected: "Llama2 70B Text FP16"},
		{input: "qwen2:latest", expected: "Qwen2 (latest)"},
		{input: "qwen2", expected: "Qwen2"},
		{input: "qwen2:0.5b", expected: "Qwen2 0.5B"},
		{input: "qwen2:1.5b", expected: "Qwen2 1.5B"},
		{input: "qwen2:7b", expected: "Qwen2 7B"},
		{input: "qwen2:72b", expected: "Qwen2 72B"},
		{input: "qwen2:0.5b-instruct", expected: "Qwen2 0.5B Instruct"},
		{input: "qwen2:0.5b-instruct-q2_K", expected: "Qwen2 0.5B Instruct Q2_K"},
		{input: "qwen2:0.5b-instruct-q3_K_S", expected: "Qwen2 0.5B Instruct Q3 K_S"},
		{input: "qwen2:0.5b-instruct-q3_K_M", expected: "Qwen2 0.5B Instruct Q3 K_M"},
		{input: "qwen2:0.5b-instruct-q3_K_L", expected: "Qwen2 0.5B Instruct Q3 K_L"},
		{input: "qwen2:0.5b-instruct-q4_0", expected: "Qwen2 0.5B Instruct Q4_0"},
		{input: "qwen2:0.5b-instruct-q4_1", expected: "Qwen2 0.5B Instruct Q4 1"},
		{input: "qwen2:0.5b-instruct-q4_K_S", expected: "Qwen2 0.5B Instruct Q4 K_S"},
		{input: "qwen2:0.5b-instruct-q4_K_M", expected: "Qwen2 0.5B Instruct Q4_K_M"},
		{input: "qwen2:0.5b-instruct-q5_0", expected: "Qwen2 0.5B Instruct Q5 0"},
		{input: "qwen2:0.5b-instruct-q5_1", expected: "Qwen2 0.5B Instruct Q5 1"},
		{input: "qwen2:0.5b-instruct-q5_K_S", expected: "Qwen2 0.5B Instruct Q5 K_S"},
		{input: "qwen2:0.5b-instruct-q5_K_M", expected: "Qwen2 0.5B Instruct Q5 K_M"},
		{input: "qwen2:0.5b-instruct-q6_K", expected: "Qwen2 0.5B Instruct Q6 K"},
		{input: "qwen2:0.5b-instruct-q8_0", expected: "Qwen2 0.5B Instruct Q8_0"},
		{input: "qwen2:0.5b-instruct-fp16", expected: "Qwen2 0.5B Instruct FP16"},
		{input: "qwen2:1.5b-instruct", expected: "Qwen2 1.5B Instruct"},
		{input: "qwen2:1.5b-instruct-q2_K", expected: "Qwen2 1.5B Instruct Q2_K"},
		{input: "qwen2:1.5b-instruct-q3_K_S", expected: "Qwen2 1.5B Instruct Q3 K_S"},
		{input: "qwen2:1.5b-instruct-q3_K_M", expected: "Qwen2 1.5B Instruct Q3 K_M"},
		{input: "qwen2:1.5b-instruct-q3_K_L", expected: "Qwen2 1.5B Instruct Q3 K_L"},
		{input: "qwen2:1.5b-instruct-q4_0", expected: "Qwen2 1.5B Instruct Q4_0"},
		{input: "qwen2:1.5b-instruct-q4_1", expected: "Qwen2 1.5B Instruct Q4 1"},
		{input: "qwen2:1.5b-instruct-q4_K_S", expected: "Qwen2 1.5B Instruct Q4 K_S"},
		{input: "qwen2:1.5b-instruct-q4_K_M", expected: "Qwen2 1.5B Instruct Q4_K_M"},
		{input: "qwen2:1.5b-instruct-q5_0", expected: "Qwen2 1.5B Instruct Q5 0"},
		{input: "qwen2:1.5b-instruct-q5_1", expected: "Qwen2 1.5B Instruct Q5 1"},
		{input: "qwen2:1.5b-instruct-q5_K_S", expected: "Qwen2 1.5B Instruct Q5 K_S"},
		{input: "qwen2:1.5b-instruct-q5_K_M", expected: "Qwen2 1.5B Instruct Q5 K_M"},
		{input: "qwen2:1.5b-instruct-q6_K", expected: "Qwen2 1.5B Instruct Q6 K"},
		{input: "qwen2:1.5b-instruct-q8_0", expected: "Qwen2 1.5B Instruct Q8_0"},
		{input: "qwen2:1.5b-instruct-fp16", expected: "Qwen2 1.5B Instruct FP16"},
		{input: "qwen2:7b-instruct", expected: "Qwen2 7B Instruct"},
		{input: "qwen2:7b-instruct-q2_K", expected: "Qwen2 7B Instruct Q2_K"},
		{input: "qwen2:7b-instruct-q3_K_S", expected: "Qwen2 7B Instruct Q3 K_S"},
		{input: "qwen2:7b-instruct-q3_K_M", expected: "Qwen2 7B Instruct Q3 K_M"},
		{input: "qwen2:7b-instruct-q3_K_L", expected: "Qwen2 7B Instruct Q3 K_L"},
		{input: "qwen2:7b-instruct-q4_0", expected: "Qwen2 7B Instruct Q4_0"},
		{input: "qwen2:7b-instruct-q4_1", expected: "Qwen2 7B Instruct Q4 1"},
		{input: "qwen2:7b-instruct-q4_K_S", expected: "Qwen2 7B Instruct Q4 K_S"},
		{input: "qwen2:7b-instruct-q4_K_M", expected: "Qwen2 7B Instruct Q4_K_M"},
		{input: "qwen2:7b-instruct-q5_0", expected: "Qwen2 7B Instruct Q5 0"},
		{input: "qwen2:7b-instruct-q5_1", expected: "Qwen2 7B Instruct Q5 1"},
		{input: "qwen2:7b-instruct-q5_K_S", expected: "Qwen2 7B Instruct Q5 K_S"},
		{input: "qwen2:7b-instruct-q5_K_M", expected: "Qwen2 7B Instruct Q5 K_M"},
		{input: "qwen2:7b-instruct-q6_K", expected: "Qwen2 7B Instruct Q6 K"},
		{input: "qwen2:7b-instruct-q8_0", expected: "Qwen2 7B Instruct Q8_0"},
		{input: "qwen2:7b-instruct-fp16", expected: "Qwen2 7B Instruct FP16"},
		{input: "qwen2:7b-text", expected: "Qwen2 7B Text"},
		{input: "qwen2:7b-text-q2_K", expected: "Qwen2 7B Text Q2_K"},
		{input: "qwen2:7b-text-q3_K_S", expected: "Qwen2 7B Text Q3 K_S"},
		{input: "qwen2:7b-text-q3_K_M", expected: "Qwen2 7B Text Q3 K_M"},
		{input: "qwen2:7b-text-q3_K_L", expected: "Qwen2 7B Text Q3 K_L"},
		{input: "qwen2:7b-text-q4_0", expected: "Qwen2 7B Text Q4_0"},
		{input: "qwen2:7b-text-q4_1", expected: "Qwen2 7B Text Q4 1"},
		{input: "qwen2:7b-text-q4_K_S", expected: "Qwen2 7B Text Q4 K_S"},
		{input: "qwen2:7b-text-q4_K_M", expected: "Qwen2 7B Text Q4_K_M"},
		{input: "qwen2:7b-text-q5_0", expected: "Qwen2 7B Text Q5 0"},
		{input: "qwen2:7b-text-q5_1", expected: "Qwen2 7B Text Q5 1"},
		{input: "qwen2:7b-text-q8_0", expected: "Qwen2 7B Text Q8_0"},
		{input: "qwen2:72b-instruct", expected: "Qwen2 72B Instruct"},
		{input: "qwen2:72b-instruct-q2_K", expected: "Qwen2 72B Instruct Q2_K"},
		{input: "qwen2:72b-instruct-q3_K_S", expected: "Qwen2 72B Instruct Q3 K_S"},
		{input: "qwen2:72b-instruct-q3_K_M", expected: "Qwen2 72B Instruct Q3 K_M"},
		{input: "qwen2:72b-instruct-q3_K_L", expected: "Qwen2 72B Instruct Q3 K_L"},
		{input: "qwen2:72b-instruct-q4_0", expected: "Qwen2 72B Instruct Q4_0"},
		{input: "qwen2:72b-instruct-q4_1", expected: "Qwen2 72B Instruct Q4 1"},
		{input: "qwen2:72b-instruct-q4_K_S", expected: "Qwen2 72B Instruct Q4 K_S"},
		{input: "qwen2:72b-instruct-q4_K_M", expected: "Qwen2 72B Instruct Q4_K_M"},
		{input: "qwen2:72b-instruct-q5_0", expected: "Qwen2 72B Instruct Q5 0"},
		{input: "qwen2:72b-instruct-q5_1", expected: "Qwen2 72B Instruct Q5 1"},
		{input: "qwen2:72b-instruct-q5_K_S", expected: "Qwen2 72B Instruct Q5 K_S"},
		{input: "qwen2:72b-instruct-q5_K_M", expected: "Qwen2 72B Instruct Q5 K_M"},
		{input: "qwen2:72b-instruct-q6_K", expected: "Qwen2 72B Instruct Q6 K"},
		{input: "qwen2:72b-instruct-q8_0", expected: "Qwen2 72B Instruct Q8_0"},
		{input: "qwen2:72b-instruct-fp16", expected: "Qwen2 72B Instruct FP16"},
		{input: "qwen2:72b-text", expected: "Qwen2 72B Text"},
		{input: "qwen2:72b-text-q2_K", expected: "Qwen2 72B Text Q2_K"},
		{input: "qwen2:72b-text-q3_K_S", expected: "Qwen2 72B Text Q3 K_S"},
		{input: "qwen2:72b-text-q3_K_M", expected: "Qwen2 72B Text Q3 K_M"},
		{input: "qwen2:72b-text-q3_K_L", expected: "Qwen2 72B Text Q3 K_L"},
		{input: "qwen2:72b-text-q4_0", expected: "Qwen2 72B Text Q4_0"},
		{input: "qwen2:72b-text-q4_1", expected: "Qwen2 72B Text Q4 1"},
		{input: "qwen2:72b-text-q4_K_S", expected: "Qwen2 72B Text Q4 K_S"},
		{input: "qwen2:72b-text-q4_K_M", expected: "Qwen2 72B Text Q4_K_M"},
		{input: "qwen2:72b-text-q5_0", expected: "Qwen2 72B Text Q5 0"},
		{input: "qwen2:72b-text-q5_1", expected: "Qwen2 72B Text Q5 1"},
		{input: "qwen2:72b-text-q5_K_S", expected: "Qwen2 72B Text Q5 K_S"},
		{input: "qwen2:72b-text-q5_K_M", expected: "Qwen2 72B Text Q5 K_M"},
		{input: "qwen2:72b-text-q6_K", expected: "Qwen2 72B Text Q6 K"},
		{input: "qwen2:72b-text-q8_0", expected: "Qwen2 72B Text Q8_0"},
		{input: "qwen2:72b-text-fp16", expected: "Qwen2 72B Text FP16"},
		{input: "minicpm-v:latest", expected: "Minicpm V (latest)"},
		{input: "minicpm-v", expected: "Minicpm V"},
		{input: "minicpm-v:8b", expected: "Minicpm V 8B"},
		{input: "minicpm-v:8b-2.6-q2_K", expected: "Minicpm V 8B 2.6 Q2_K"},
		{input: "minicpm-v:8b-2.6-q3_K_S", expected: "Minicpm V 8B 2.6 Q3 K_S"},
		{input: "minicpm-v:8b-2.6-q3_K_M", expected: "Minicpm V 8B 2.6 Q3 K_M"},
		{input: "minicpm-v:8b-2.6-q3_K_L", expected: "Minicpm V 8B 2.6 Q3 K_L"},
		{input: "minicpm-v:8b-2.6-q4_0", expected: "Minicpm V 8B 2.6 Q4_0"},
		{input: "minicpm-v:8b-2.6-q4_1", expected: "Minicpm V 8B 2.6 Q4 1"},
		{input: "minicpm-v:8b-2.6-q4_K_S", expected: "Minicpm V 8B 2.6 Q4 K_S"},
		{input: "minicpm-v:8b-2.6-q4_K_M", expected: "Minicpm V 8B 2.6 Q4_K_M"},
		{input: "minicpm-v:8b-2.6-q5_0", expected: "Minicpm V 8B 2.6 Q5 0"},
		{input: "minicpm-v:8b-2.6-q5_1", expected: "Minicpm V 8B 2.6 Q5 1"},
		{input: "minicpm-v:8b-2.6-q5_K_S", expected: "Minicpm V 8B 2.6 Q5 K_S"},
		{input: "minicpm-v:8b-2.6-q5_K_M", expected: "Minicpm V 8B 2.6 Q5 K_M"},
		{input: "minicpm-v:8b-2.6-q6_K", expected: "Minicpm V 8B 2.6 Q6 K"},
		{input: "minicpm-v:8b-2.6-q8_0", expected: "Minicpm V 8B 2.6 Q8_0"},
		{input: "minicpm-v:8b-2.6-fp16", expected: "Minicpm V 8B 2.6 FP16"},
		{input: "codellama:latest", expected: "Codellama (latest)"},
		{input: "codellama", expected: "Codellama"},
		{input: "codellama:code", expected: "Codellama Code"},
		{input: "codellama:instruct", expected: "Codellama Instruct"},
		{input: "codellama:python", expected: "Codellama Python"},
		{input: "codellama:7b", expected: "Codellama 7B"},
		{input: "codellama:13b", expected: "Codellama 13B"},
		{input: "codellama:34b", expected: "Codellama 34B"},
		{input: "codellama:70b", expected: "Codellama 70B"},
		{input: "codellama:7b-code", expected: "Codellama 7B Code"},
		{input: "codellama:7b-code-q2_K", expected: "Codellama 7B Code Q2_K"},
		{input: "codellama:7b-code-q3_K_S", expected: "Codellama 7B Code Q3 K_S"},
		{input: "codellama:7b-code-q3_K_M", expected: "Codellama 7B Code Q3 K_M"},
		{input: "codellama:7b-code-q3_K_L", expected: "Codellama 7B Code Q3 K_L"},
		{input: "codellama:7b-code-q4_0", expected: "Codellama 7B Code Q4_0"},
		{input: "codellama:7b-code-q4_1", expected: "Codellama 7B Code Q4 1"},
		{input: "codellama:7b-code-q4_K_S", expected: "Codellama 7B Code Q4 K_S"},
		{input: "codellama:7b-code-q4_K_M", expected: "Codellama 7B Code Q4_K_M"},
		{input: "codellama:7b-code-q5_0", expected: "Codellama 7B Code Q5 0"},
		{input: "codellama:7b-code-q5_1", expected: "Codellama 7B Code Q5 1"},
		{input: "codellama:7b-code-q5_K_S", expected: "Codellama 7B Code Q5 K_S"},
		{input: "codellama:7b-code-q5_K_M", expected: "Codellama 7B Code Q5 K_M"},
		{input: "codellama:7b-code-q6_K", expected: "Codellama 7B Code Q6 K"},
		{input: "codellama:7b-code-q8_0", expected: "Codellama 7B Code Q8_0"},
		{input: "codellama:7b-code-fp16", expected: "Codellama 7B Code FP16"},
		{input: "codellama:7b-instruct", expected: "Codellama 7B Instruct"},
		{input: "codellama:7b-instruct-q2_K", expected: "Codellama 7B Instruct Q2_K"},
		{input: "codellama:7b-instruct-q3_K_S", expected: "Codellama 7B Instruct Q3 K_S"},
		{input: "codellama:7b-instruct-q3_K_M", expected: "Codellama 7B Instruct Q3 K_M"},
		{input: "codellama:7b-instruct-q3_K_L", expected: "Codellama 7B Instruct Q3 K_L"},
		{input: "codellama:7b-instruct-q4_0", expected: "Codellama 7B Instruct Q4_0"},
		{input: "codellama:7b-instruct-q4_1", expected: "Codellama 7B Instruct Q4 1"},
		{input: "codellama:7b-instruct-q4_K_S", expected: "Codellama 7B Instruct Q4 K_S"},
		{input: "codellama:7b-instruct-q4_K_M", expected: "Codellama 7B Instruct Q4_K_M"},
		{input: "codellama:7b-instruct-q5_0", expected: "Codellama 7B Instruct Q5 0"},
		{input: "codellama:7b-instruct-q5_1", expected: "Codellama 7B Instruct Q5 1"},
		{input: "codellama:7b-instruct-q5_K_S", expected: "Codellama 7B Instruct Q5 K_S"},
		{input: "codellama:7b-instruct-q5_K_M", expected: "Codellama 7B Instruct Q5 K_M"},
		{input: "codellama:7b-instruct-q6_K", expected: "Codellama 7B Instruct Q6 K"},
		{input: "codellama:7b-instruct-q8_0", expected: "Codellama 7B Instruct Q8_0"},
		{input: "codellama:7b-instruct-fp16", expected: "Codellama 7B Instruct FP16"},
		{input: "codellama:7b-python", expected: "Codellama 7B Python"},
		{input: "codellama:7b-python-q2_K", expected: "Codellama 7B Python Q2_K"},
		{input: "codellama:7b-python-q3_K_S", expected: "Codellama 7B Python Q3 K_S"},
		{input: "codellama:7b-python-q3_K_M", expected: "Codellama 7B Python Q3 K_M"},
		{input: "codellama:7b-python-q3_K_L", expected: "Codellama 7B Python Q3 K_L"},
		{input: "codellama:7b-python-q4_0", expected: "Codellama 7B Python Q4_0"},
		{input: "codellama:7b-python-q4_1", expected: "Codellama 7B Python Q4 1"},
		{input: "codellama:7b-python-q4_K_S", expected: "Codellama 7B Python Q4 K_S"},
		{input: "codellama:7b-python-q4_K_M", expected: "Codellama 7B Python Q4_K_M"},
		{input: "codellama:7b-python-q5_0", expected: "Codellama 7B Python Q5 0"},
		{input: "codellama:7b-python-q5_1", expected: "Codellama 7B Python Q5 1"},
		{input: "codellama:7b-python-q5_K_S", expected: "Codellama 7B Python Q5 K_S"},
		{input: "codellama:7b-python-q5_K_M", expected: "Codellama 7B Python Q5 K_M"},
		{input: "codellama:7b-python-q6_K", expected: "Codellama 7B Python Q6 K"},
		{input: "codellama:7b-python-q8_0", expected: "Codellama 7B Python Q8_0"},
		{input: "codellama:7b-python-fp16", expected: "Codellama 7B Python FP16"},
		{input: "codellama:13b-code", expected: "Codellama 13B Code"},
		{input: "codellama:13b-code-q2_K", expected: "Codellama 13B Code Q2_K"},
		{input: "codellama:13b-code-q3_K_S", expected: "Codellama 13B Code Q3 K_S"},
		{input: "codellama:13b-code-q3_K_M", expected: "Codellama 13B Code Q3 K_M"},
		{input: "codellama:13b-code-q3_K_L", expected: "Codellama 13B Code Q3 K_L"},
		{input: "codellama:13b-code-q4_0", expected: "Codellama 13B Code Q4_0"},
		{input: "codellama:13b-code-q4_1", expected: "Codellama 13B Code Q4 1"},
		{input: "codellama:13b-code-q4_K_S", expected: "Codellama 13B Code Q4 K_S"},
		{input: "codellama:13b-code-q4_K_M", expected: "Codellama 13B Code Q4_K_M"},
		{input: "codellama:13b-code-q5_0", expected: "Codellama 13B Code Q5 0"},
		{input: "codellama:13b-code-q5_1", expected: "Codellama 13B Code Q5 1"},
		{input: "codellama:13b-code-q5_K_S", expected: "Codellama 13B Code Q5 K_S"},
		{input: "codellama:13b-code-q5_K_M", expected: "Codellama 13B Code Q5 K_M"},
		{input: "codellama:13b-code-q6_K", expected: "Codellama 13B Code Q6 K"},
		{input: "codellama:13b-code-q8_0", expected: "Codellama 13B Code Q8_0"},
		{input: "codellama:13b-code-fp16", expected: "Codellama 13B Code FP16"},
		{input: "codellama:13b-instruct", expected: "Codellama 13B Instruct"},
		{input: "codellama:13b-instruct-q2_K", expected: "Codellama 13B Instruct Q2_K"},
		{input: "codellama:13b-instruct-q3_K_S", expected: "Codellama 13B Instruct Q3 K_S"},
		{input: "codellama:13b-instruct-q3_K_M", expected: "Codellama 13B Instruct Q3 K_M"},
		{input: "codellama:13b-instruct-q3_K_L", expected: "Codellama 13B Instruct Q3 K_L"},
		{input: "codellama:13b-instruct-q4_0", expected: "Codellama 13B Instruct Q4_0"},
		{input: "codellama:13b-instruct-q4_1", expected: "Codellama 13B Instruct Q4 1"},
		{input: "codellama:13b-instruct-q4_K_S", expected: "Codellama 13B Instruct Q4 K_S"},
		{input: "codellama:13b-instruct-q4_K_M", expected: "Codellama 13B Instruct Q4_K_M"},
		{input: "codellama:13b-instruct-q5_0", expected: "Codellama 13B Instruct Q5 0"},
		{input: "codellama:13b-instruct-q5_1", expected: "Codellama 13B Instruct Q5 1"},
		{input: "codellama:13b-instruct-q5_K_S", expected: "Codellama 13B Instruct Q5 K_S"},
		{input: "codellama:13b-instruct-q5_K_M", expected: "Codellama 13B Instruct Q5 K_M"},
		{input: "codellama:13b-instruct-q6_K", expected: "Codellama 13B Instruct Q6 K"},
		{input: "codellama:13b-instruct-q8_0", expected: "Codellama 13B Instruct Q8_0"},
		{input: "codellama:13b-instruct-fp16", expected: "Codellama 13B Instruct FP16"},
		{input: "codellama:13b-python", expected: "Codellama 13B Python"},
		{input: "codellama:13b-python-q2_K", expected: "Codellama 13B Python Q2_K"},
		{input: "codellama:13b-python-q3_K_S", expected: "Codellama 13B Python Q3 K_S"},
		{input: "codellama:13b-python-q3_K_M", expected: "Codellama 13B Python Q3 K_M"},
		{input: "codellama:13b-python-q3_K_L", expected: "Codellama 13B Python Q3 K_L"},
		{input: "codellama:13b-python-q4_0", expected: "Codellama 13B Python Q4_0"},
		{input: "codellama:13b-python-q4_1", expected: "Codellama 13B Python Q4 1"},
		{input: "codellama:13b-python-q4_K_S", expected: "Codellama 13B Python Q4 K_S"},
		{input: "codellama:13b-python-q4_K_M", expected: "Codellama 13B Python Q4_K_M"},
		{input: "codellama:13b-python-q5_0", expected: "Codellama 13B Python Q5 0"},
		{input: "codellama:13b-python-q5_1", expected: "Codellama 13B Python Q5 1"},
		{input: "codellama:13b-python-q5_K_S", expected: "Codellama 13B Python Q5 K_S"},
		{input: "codellama:13b-python-q5_K_M", expected: "Codellama 13B Python Q5 K_M"},
		{input: "codellama:13b-python-q6_K", expected: "Codellama 13B Python Q6 K"},
		{input: "codellama:13b-python-q8_0", expected: "Codellama 13B Python Q8_0"},
		{input: "codellama:13b-python-fp16", expected: "Codellama 13B Python FP16"},
		{input: "codellama:34b-code", expected: "Codellama 34B Code"},
		{input: "codellama:34b-code-q2_K", expected: "Codellama 34B Code Q2_K"},
		{input: "codellama:34b-code-q3_K_S", expected: "Codellama 34B Code Q3 K_S"},
		{input: "codellama:34b-code-q3_K_M", expected: "Codellama 34B Code Q3 K_M"},
		{input: "codellama:34b-code-q3_K_L", expected: "Codellama 34B Code Q3 K_L"},
		{input: "codellama:34b-code-q4_0", expected: "Codellama 34B Code Q4_0"},
		{input: "codellama:34b-code-q4_1", expected: "Codellama 34B Code Q4 1"},
		{input: "codellama:34b-code-q4_K_S", expected: "Codellama 34B Code Q4 K_S"},
		{input: "codellama:34b-code-q4_K_M", expected: "Codellama 34B Code Q4_K_M"},
		{input: "codellama:34b-code-q5_0", expected: "Codellama 34B Code Q5 0"},
		{input: "codellama:34b-code-q5_1", expected: "Codellama 34B Code Q5 1"},
		{input: "codellama:34b-code-q5_K_S", expected: "Codellama 34B Code Q5 K_S"},
		{input: "codellama:34b-code-q5_K_M", expected: "Codellama 34B Code Q5 K_M"},
		{input: "codellama:34b-code-q6_K", expected: "Codellama 34B Code Q6 K"},
		{input: "codellama:34b-code-q8_0", expected: "Codellama 34B Code Q8_0"},
		{input: "codellama:34b-instruct", expected: "Codellama 34B Instruct"},
		{input: "codellama:34b-instruct-q2_K", expected: "Codellama 34B Instruct Q2_K"},
		{input: "codellama:34b-instruct-q3_K_S", expected: "Codellama 34B Instruct Q3 K_S"},
		{input: "codellama:34b-instruct-q3_K_M", expected: "Codellama 34B Instruct Q3 K_M"},
		{input: "codellama:34b-instruct-q3_K_L", expected: "Codellama 34B Instruct Q3 K_L"},
		{input: "codellama:34b-instruct-q4_0", expected: "Codellama 34B Instruct Q4_0"},
		{input: "codellama:34b-instruct-q4_1", expected: "Codellama 34B Instruct Q4 1"},
		{input: "codellama:34b-instruct-q4_K_S", expected: "Codellama 34B Instruct Q4 K_S"},
		{input: "codellama:34b-instruct-q4_K_M", expected: "Codellama 34B Instruct Q4_K_M"},
		{input: "codellama:34b-instruct-q5_0", expected: "Codellama 34B Instruct Q5 0"},
		{input: "codellama:34b-instruct-q5_1", expected: "Codellama 34B Instruct Q5 1"},
		{input: "codellama:34b-instruct-q5_K_S", expected: "Codellama 34B Instruct Q5 K_S"},
		{input: "codellama:34b-instruct-q5_K_M", expected: "Codellama 34B Instruct Q5 K_M"},
		{input: "codellama:34b-instruct-q6_K", expected: "Codellama 34B Instruct Q6 K"},
		{input: "codellama:34b-instruct-q8_0", expected: "Codellama 34B Instruct Q8_0"},
		{input: "codellama:34b-instruct-fp16", expected: "Codellama 34B Instruct FP16"},
		{input: "codellama:34b-python", expected: "Codellama 34B Python"},
		{input: "codellama:34b-python-q2_K", expected: "Codellama 34B Python Q2_K"},
		{input: "codellama:34b-python-q3_K_S", expected: "Codellama 34B Python Q3 K_S"},
		{input: "codellama:34b-python-q3_K_M", expected: "Codellama 34B Python Q3 K_M"},
		{input: "codellama:34b-python-q3_K_L", expected: "Codellama 34B Python Q3 K_L"},
		{input: "codellama:34b-python-q4_0", expected: "Codellama 34B Python Q4_0"},
		{input: "codellama:34b-python-q4_1", expected: "Codellama 34B Python Q4 1"},
		{input: "codellama:34b-python-q4_K_S", expected: "Codellama 34B Python Q4 K_S"},
		{input: "codellama:34b-python-q4_K_M", expected: "Codellama 34B Python Q4_K_M"},
		{input: "codellama:34b-python-q5_0", expected: "Codellama 34B Python Q5 0"},
		{input: "codellama:34b-python-q5_1", expected: "Codellama 34B Python Q5 1"},
		{input: "codellama:34b-python-q5_K_S", expected: "Codellama 34B Python Q5 K_S"},
		{input: "codellama:34b-python-q5_K_M", expected: "Codellama 34B Python Q5 K_M"},
		{input: "codellama:34b-python-q6_K", expected: "Codellama 34B Python Q6 K"},
		{input: "codellama:34b-python-q8_0", expected: "Codellama 34B Python Q8_0"},
		{input: "codellama:34b-python-fp16", expected: "Codellama 34B Python FP16"},
		{input: "codellama:70b-code", expected: "Codellama 70B Code"},
		{input: "codellama:70b-code-q2_K", expected: "Codellama 70B Code Q2_K"},
		{input: "codellama:70b-code-q3_K_S", expected: "Codellama 70B Code Q3 K_S"},
		{input: "codellama:70b-code-q3_K_M", expected: "Codellama 70B Code Q3 K_M"},
		{input: "codellama:70b-code-q3_K_L", expected: "Codellama 70B Code Q3 K_L"},
		{input: "codellama:70b-code-q4_0", expected: "Codellama 70B Code Q4_0"},
		{input: "codellama:70b-code-q4_1", expected: "Codellama 70B Code Q4 1"},
		{input: "codellama:70b-code-q4_K_S", expected: "Codellama 70B Code Q4 K_S"},
		{input: "codellama:70b-code-q4_K_M", expected: "Codellama 70B Code Q4_K_M"},
		{input: "codellama:70b-code-q5_0", expected: "Codellama 70B Code Q5 0"},
		{input: "codellama:70b-code-q5_1", expected: "Codellama 70B Code Q5 1"},
		{input: "codellama:70b-code-q5_K_S", expected: "Codellama 70B Code Q5 K_S"},
		{input: "codellama:70b-code-q5_K_M", expected: "Codellama 70B Code Q5 K_M"},
		{input: "codellama:70b-code-q6_K", expected: "Codellama 70B Code Q6 K"},
		{input: "codellama:70b-code-q8_0", expected: "Codellama 70B Code Q8_0"},
		{input: "codellama:70b-code-fp16", expected: "Codellama 70B Code FP16"},
		{input: "codellama:70b-instruct", expected: "Codellama 70B Instruct"},
		{input: "codellama:70b-instruct-q2_K", expected: "Codellama 70B Instruct Q2_K"},
		{input: "codellama:70b-instruct-q3_K_S", expected: "Codellama 70B Instruct Q3 K_S"},
		{input: "codellama:70b-instruct-q3_K_M", expected: "Codellama 70B Instruct Q3 K_M"},
		{input: "codellama:70b-instruct-q3_K_L", expected: "Codellama 70B Instruct Q3 K_L"},
		{input: "codellama:70b-instruct-q4_0", expected: "Codellama 70B Instruct Q4_0"},
		{input: "codellama:70b-instruct-q4_1", expected: "Codellama 70B Instruct Q4 1"},
		{input: "codellama:70b-instruct-q4_K_S", expected: "Codellama 70B Instruct Q4 K_S"},
		{input: "codellama:70b-instruct-q4_K_M", expected: "Codellama 70B Instruct Q4_K_M"},
		{input: "codellama:70b-instruct-q5_0", expected: "Codellama 70B Instruct Q5 0"},
		{input: "codellama:70b-instruct-q5_1", expected: "Codellama 70B Instruct Q5 1"},
		{input: "codellama:70b-instruct-q5_K_S", expected: "Codellama 70B Instruct Q5 K_S"},
		{input: "codellama:70b-instruct-q5_K_M", expected: "Codellama 70B Instruct Q5 K_M"},
		{input: "codellama:70b-instruct-q6_K", expected: "Codellama 70B Instruct Q6 K"},
		{input: "codellama:70b-instruct-q8_0", expected: "Codellama 70B Instruct Q8_0"},
		{input: "codellama:70b-instruct-fp16", expected: "Codellama 70B Instruct FP16"},
		{input: "codellama:70b-python", expected: "Codellama 70B Python"},
		{input: "codellama:70b-python-q2_K", expected: "Codellama 70B Python Q2_K"},
		{input: "codellama:70b-python-q3_K_S", expected: "Codellama 70B Python Q3 K_S"},
		{input: "codellama:70b-python-q3_K_M", expected: "Codellama 70B Python Q3 K_M"},
		{input: "codellama:70b-python-q3_K_L", expected: "Codellama 70B Python Q3 K_L"},
		{input: "codellama:70b-python-q4_0", expected: "Codellama 70B Python Q4_0"},
		{input: "codellama:70b-python-q4_1", expected: "Codellama 70B Python Q4 1"},
		{input: "codellama:70b-python-q4_K_S", expected: "Codellama 70B Python Q4 K_S"},
		{input: "codellama:70b-python-q4_K_M", expected: "Codellama 70B Python Q4_K_M"},
		{input: "codellama:70b-python-q5_0", expected: "Codellama 70B Python Q5 0"},
		{input: "codellama:70b-python-q5_1", expected: "Codellama 70B Python Q5 1"},
		{input: "codellama:70b-python-q5_K_S", expected: "Codellama 70B Python Q5 K_S"},
		{input: "codellama:70b-python-q5_K_M", expected: "Codellama 70B Python Q5 K_M"},
		{input: "codellama:70b-python-q6_K", expected: "Codellama 70B Python Q6 K"},
		{input: "codellama:70b-python-q8_0", expected: "Codellama 70B Python Q8_0"},
		{input: "codellama:70b-python-fp16", expected: "Codellama 70B Python FP16"},
		{input: "dolphin3:latest", expected: "Dolphin3 (latest)"},
		{input: "dolphin3", expected: "Dolphin3"},
		{input: "dolphin3:8b", expected: "Dolphin3 8B"},
		{input: "dolphin3:8b-llama3.1-q4_K_M", expected: "Dolphin3 8B Llama3.1 Q4_K_M"},
		{input: "dolphin3:8b-llama3.1-q8_0", expected: "Dolphin3 8B Llama3.1 Q8_0"},
		{input: "dolphin3:8b-llama3.1-fp16", expected: "Dolphin3 8B Llama3.1 FP16"},
		{input: "olmo2:latest", expected: "Olmo2 (latest)"},
		{input: "olmo2", expected: "Olmo2"},
		{input: "olmo2:7b", expected: "Olmo2 7B"},
		{input: "olmo2:13b", expected: "Olmo2 13B"},
		{input: "olmo2:7b-1124-instruct-q4_K_M", expected: "Olmo2 7B 1124 Instruct Q4_K_M"},
		{input: "olmo2:7b-1124-instruct-q8_0", expected: "Olmo2 7B 1124 Instruct Q8_0"},
		{input: "olmo2:7b-1124-instruct-fp16", expected: "Olmo2 7B 1124 Instruct FP16"},
		{input: "olmo2:13b-1124-instruct-q4_K_M", expected: "Olmo2 13B 1124 Instruct Q4_K_M"},
		{input: "olmo2:13b-1124-instruct-q8_0", expected: "Olmo2 13B 1124 Instruct Q8_0"},
		{input: "olmo2:13b-1124-instruct-fp16", expected: "Olmo2 13B 1124 Instruct FP16"},
		{input: "llama3.2-vision:latest", expected: "Llama3.2 Vision (latest)"},
		{input: "llama3.2-vision", expected: "Llama3.2 Vision"},
		{input: "llama3.2-vision:11b", expected: "Llama3.2 Vision 11B"},
		{input: "llama3.2-vision:90b", expected: "Llama3.2 Vision 90B"},
		{input: "llama3.2-vision:11b-instruct-q4_K_M", expected: "Llama3.2 Vision 11B Instruct Q4_K_M"},
		{input: "llama3.2-vision:11b-instruct-q8_0", expected: "Llama3.2 Vision 11B Instruct Q8_0"},
		{input: "llama3.2-vision:11b-instruct-fp16", expected: "Llama3.2 Vision 11B Instruct FP16"},
		{input: "llama3.2-vision:90b-instruct-q4_K_M", expected: "Llama3.2 Vision 90B Instruct Q4_K_M"},
		{input: "llama3.2-vision:90b-instruct-q8_0", expected: "Llama3.2 Vision 90B Instruct Q8_0"},
		{input: "llama3.2-vision:90b-instruct-fp16", expected: "Llama3.2 Vision 90B Instruct FP16"},
		{input: "tinyllama:latest", expected: "Tinyllama (latest)"},
		{input: "tinyllama", expected: "Tinyllama"},
		{input: "tinyllama:chat", expected: "Tinyllama Chat"},
		{input: "tinyllama:v0.6", expected: "Tinyllama v0.6"},
		{input: "tinyllama:v1", expected: "Tinyllama v1"},
		{input: "tinyllama:1.1b", expected: "Tinyllama 1.1B"},
		{input: "tinyllama:1.1b-chat", expected: "Tinyllama 1.1B Chat"},
		{input: "tinyllama:1.1b-chat-v0.6-q2_K", expected: "Tinyllama 1.1B Chat v0.6 Q2_K"},
		{input: "tinyllama:1.1b-chat-v0.6-q3_K_S", expected: "Tinyllama 1.1B Chat v0.6 Q3 K_S"},
		{input: "tinyllama:1.1b-chat-v0.6-q3_K_M", expected: "Tinyllama 1.1B Chat v0.6 Q3 K_M"},
		{input: "tinyllama:1.1b-chat-v0.6-q3_K_L", expected: "Tinyllama 1.1B Chat v0.6 Q3 K_L"},
		{input: "tinyllama:1.1b-chat-v0.6-q4_0", expected: "Tinyllama 1.1B Chat v0.6 Q4_0"},
		{input: "tinyllama:1.1b-chat-v0.6-q4_1", expected: "Tinyllama 1.1B Chat v0.6 Q4 1"},
		{input: "tinyllama:1.1b-chat-v0.6-q4_K_S", expected: "Tinyllama 1.1B Chat v0.6 Q4 K_S"},
		{input: "tinyllama:1.1b-chat-v0.6-q4_K_M", expected: "Tinyllama 1.1B Chat v0.6 Q4_K_M"},
		{input: "tinyllama:1.1b-chat-v0.6-q5_0", expected: "Tinyllama 1.1B Chat v0.6 Q5 0"},
		{input: "tinyllama:1.1b-chat-v0.6-q5_1", expected: "Tinyllama 1.1B Chat v0.6 Q5 1"},
		{input: "tinyllama:1.1b-chat-v0.6-q5_K_S", expected: "Tinyllama 1.1B Chat v0.6 Q5 K_S"},
		{input: "tinyllama:1.1b-chat-v0.6-q5_K_M", expected: "Tinyllama 1.1B Chat v0.6 Q5 K_M"},
		{input: "tinyllama:1.1b-chat-v0.6-q6_K", expected: "Tinyllama 1.1B Chat v0.6 Q6 K"},
		{input: "tinyllama:1.1b-chat-v0.6-q8_0", expected: "Tinyllama 1.1B Chat v0.6 Q8_0"},
		{input: "tinyllama:1.1b-chat-v0.6-fp16", expected: "Tinyllama 1.1B Chat v0.6 FP16"},
		{input: "tinyllama:1.1b-chat-v1-q2_K", expected: "Tinyllama 1.1B Chat v1 Q2_K"},
		{input: "tinyllama:1.1b-chat-v1-q3_K_S", expected: "Tinyllama 1.1B Chat v1 Q3 K_S"},
		{input: "tinyllama:1.1b-chat-v1-q3_K_M", expected: "Tinyllama 1.1B Chat v1 Q3 K_M"},
		{input: "tinyllama:1.1b-chat-v1-q3_K_L", expected: "Tinyllama 1.1B Chat v1 Q3 K_L"},
		{input: "tinyllama:1.1b-chat-v1-q4_0", expected: "Tinyllama 1.1B Chat v1 Q4_0"},
		{input: "tinyllama:1.1b-chat-v1-q4_1", expected: "Tinyllama 1.1B Chat v1 Q4 1"},
		{input: "tinyllama:1.1b-chat-v1-q4_K_S", expected: "Tinyllama 1.1B Chat v1 Q4 K_S"},
		{input: "tinyllama:1.1b-chat-v1-q4_K_M", expected: "Tinyllama 1.1B Chat v1 Q4_K_M"},
		{input: "tinyllama:1.1b-chat-v1-q5_0", expected: "Tinyllama 1.1B Chat v1 Q5 0"},
		{input: "tinyllama:1.1b-chat-v1-q5_1", expected: "Tinyllama 1.1B Chat v1 Q5 1"},
		{input: "tinyllama:1.1b-chat-v1-q5_K_S", expected: "Tinyllama 1.1B Chat v1 Q5 K_S"},
		{input: "tinyllama:1.1b-chat-v1-q5_K_M", expected: "Tinyllama 1.1B Chat v1 Q5 K_M"},
		{input: "tinyllama:1.1b-chat-v1-q6_K", expected: "Tinyllama 1.1B Chat v1 Q6 K"},
		{input: "tinyllama:1.1b-chat-v1-q8_0", expected: "Tinyllama 1.1B Chat v1 Q8_0"},
		{input: "tinyllama:1.1b-chat-v1-fp16", expected: "Tinyllama 1.1B Chat v1 FP16"},
		{input: "mistral-nemo:latest", expected: "Mistral Nemo (latest)"},
		{input: "mistral-nemo", expected: "Mistral Nemo"},
		{input: "mistral-nemo:12b", expected: "Mistral Nemo 12B"},
		{input: "mistral-nemo:12b-instruct-2407-q2_K", expected: "Mistral Nemo 12B Instruct 2407 Q2_K"},
		{input: "mistral-nemo:12b-instruct-2407-q3_K_S", expected: "Mistral Nemo 12B Instruct 2407 Q3 K_S"},
		{input: "mistral-nemo:12b-instruct-2407-q3_K_M", expected: "Mistral Nemo 12B Instruct 2407 Q3 K_M"},
		{input: "mistral-nemo:12b-instruct-2407-q3_K_L", expected: "Mistral Nemo 12B Instruct 2407 Q3 K_L"},
		{input: "mistral-nemo:12b-instruct-2407-q4_0", expected: "Mistral Nemo 12B Instruct 2407 Q4_0"},
		{input: "mistral-nemo:12b-instruct-2407-q4_1", expected: "Mistral Nemo 12B Instruct 2407 Q4 1"},
		{input: "mistral-nemo:12b-instruct-2407-q4_K_S", expected: "Mistral Nemo 12B Instruct 2407 Q4 K_S"},
		{input: "mistral-nemo:12b-instruct-2407-q4_K_M", expected: "Mistral Nemo 12B Instruct 2407 Q4_K_M"},
		{input: "mistral-nemo:12b-instruct-2407-q5_0", expected: "Mistral Nemo 12B Instruct 2407 Q5 0"},
		{input: "mistral-nemo:12b-instruct-2407-q5_1", expected: "Mistral Nemo 12B Instruct 2407 Q5 1"},
		{input: "mistral-nemo:12b-instruct-2407-q5_K_S", expected: "Mistral Nemo 12B Instruct 2407 Q5 K_S"},
		{input: "mistral-nemo:12b-instruct-2407-q5_K_M", expected: "Mistral Nemo 12B Instruct 2407 Q5 K_M"},
		{input: "mistral-nemo:12b-instruct-2407-q6_K", expected: "Mistral Nemo 12B Instruct 2407 Q6 K"},
		{input: "mistral-nemo:12b-instruct-2407-q8_0", expected: "Mistral Nemo 12B Instruct 2407 Q8_0"},
		{input: "mistral-nemo:12b-instruct-2407-fp16", expected: "Mistral Nemo 12B Instruct 2407 FP16"},
		{input: "deepseek-v3:latest", expected: "Deepseek v3 (latest)"},
		{input: "deepseek-v3", expected: "Deepseek v3"},
		{input: "deepseek-v3:671b", expected: "Deepseek v3 671B"},
		{input: "deepseek-v3:671b-q4_K_M", expected: "Deepseek v3 671B Q4_K_M"},
		{input: "deepseek-v3:671b-q8_0", expected: "Deepseek v3 671B Q8_0"},
		{input: "deepseek-v3:671b-fp16", expected: "Deepseek v3 671B FP16"},
		{input: "bge-m3:latest", expected: "Bge M3 (latest)"},
		{input: "bge-m3", expected: "Bge M3"},
		{input: "bge-m3:567m", expected: "Bge M3 567M"},
		{input: "bge-m3:567m-fp16", expected: "Bge M3 567M FP16"},
		{input: "llama3.3:latest", expected: "Llama3.3 (latest)"},
		{input: "llama3.3", expected: "Llama3.3"},
		{input: "llama3.3:70b", expected: "Llama3.3 70B"},
		{input: "llama3.3:70b-instruct-q2_K", expected: "Llama3.3 70B Instruct Q2_K"},
		{input: "llama3.3:70b-instruct-q3_K_S", expected: "Llama3.3 70B Instruct Q3 K_S"},
		{input: "llama3.3:70b-instruct-q3_K_M", expected: "Llama3.3 70B Instruct Q3 K_M"},
		{input: "llama3.3:70b-instruct-q4_0", expected: "Llama3.3 70B Instruct Q4_0"},
		{input: "llama3.3:70b-instruct-q4_K_S", expected: "Llama3.3 70B Instruct Q4 K_S"},
		{input: "llama3.3:70b-instruct-q4_K_M", expected: "Llama3.3 70B Instruct Q4_K_M"},
		{input: "llama3.3:70b-instruct-q5_0", expected: "Llama3.3 70B Instruct Q5 0"},
		{input: "llama3.3:70b-instruct-q5_1", expected: "Llama3.3 70B Instruct Q5 1"},
		{input: "llama3.3:70b-instruct-q5_K_M", expected: "Llama3.3 70B Instruct Q5 K_M"},
		{input: "llama3.3:70b-instruct-q6_K", expected: "Llama3.3 70B Instruct Q6 K"},
		{input: "llama3.3:70b-instruct-q8_0", expected: "Llama3.3 70B Instruct Q8_0"},
		{input: "llama3.3:70b-instruct-fp16", expected: "Llama3.3 70B Instruct FP16"},
		{input: "deepseek-coder:latest", expected: "Deepseek Coder (latest)"},
		{input: "deepseek-coder", expected: "Deepseek Coder"},
		{input: "deepseek-coder:base", expected: "Deepseek Coder Base"},
		{input: "deepseek-coder:instruct", expected: "Deepseek Coder Instruct"},
		{input: "deepseek-coder:1.3b", expected: "Deepseek Coder 1.3B"},
		{input: "deepseek-coder:6.7b", expected: "Deepseek Coder 6.7B"},
		{input: "deepseek-coder:33b", expected: "Deepseek Coder 33B"},
		{input: "deepseek-coder:1.3b-base", expected: "Deepseek Coder 1.3B Base"},
		{input: "deepseek-coder:1.3b-base-q2_K", expected: "Deepseek Coder 1.3B Base Q2_K"},
		{input: "deepseek-coder:1.3b-base-q3_K_S", expected: "Deepseek Coder 1.3B Base Q3 K_S"},
		{input: "deepseek-coder:1.3b-base-q3_K_M", expected: "Deepseek Coder 1.3B Base Q3 K_M"},
		{input: "deepseek-coder:1.3b-base-q3_K_L", expected: "Deepseek Coder 1.3B Base Q3 K_L"},
		{input: "deepseek-coder:1.3b-base-q4_0", expected: "Deepseek Coder 1.3B Base Q4_0"},
		{input: "deepseek-coder:1.3b-base-q4_1", expected: "Deepseek Coder 1.3B Base Q4 1"},
		{input: "deepseek-coder:1.3b-base-q4_K_S", expected: "Deepseek Coder 1.3B Base Q4 K_S"},
		{input: "deepseek-coder:1.3b-base-q4_K_M", expected: "Deepseek Coder 1.3B Base Q4_K_M"},
		{input: "deepseek-coder:1.3b-base-q5_0", expected: "Deepseek Coder 1.3B Base Q5 0"},
		{input: "deepseek-coder:1.3b-base-q5_1", expected: "Deepseek Coder 1.3B Base Q5 1"},
		{input: "deepseek-coder:1.3b-base-q5_K_S", expected: "Deepseek Coder 1.3B Base Q5 K_S"},
		{input: "deepseek-coder:1.3b-base-q5_K_M", expected: "Deepseek Coder 1.3B Base Q5 K_M"},
		{input: "deepseek-coder:1.3b-base-q6_K", expected: "Deepseek Coder 1.3B Base Q6 K"},
		{input: "deepseek-coder:1.3b-base-q8_0", expected: "Deepseek Coder 1.3B Base Q8_0"},
		{input: "deepseek-coder:1.3b-base-fp16", expected: "Deepseek Coder 1.3B Base FP16"},
		{input: "deepseek-coder:1.3b-instruct", expected: "Deepseek Coder 1.3B Instruct"},
		{input: "deepseek-coder:1.3b-instruct-q2_K", expected: "Deepseek Coder 1.3B Instruct Q2_K"},
		{input: "deepseek-coder:1.3b-instruct-q3_K_S", expected: "Deepseek Coder 1.3B Instruct Q3 K_S"},
		{input: "deepseek-coder:1.3b-instruct-q3_K_M", expected: "Deepseek Coder 1.3B Instruct Q3 K_M"},
		{input: "deepseek-coder:1.3b-instruct-q3_K_L", expected: "Deepseek Coder 1.3B Instruct Q3 K_L"},
		{input: "deepseek-coder:1.3b-instruct-q4_0", expected: "Deepseek Coder 1.3B Instruct Q4_0"},
		{input: "deepseek-coder:1.3b-instruct-q4_1", expected: "Deepseek Coder 1.3B Instruct Q4 1"},
		{input: "deepseek-coder:1.3b-instruct-q4_K_S", expected: "Deepseek Coder 1.3B Instruct Q4 K_S"},
		{input: "deepseek-coder:1.3b-instruct-q4_K_M", expected: "Deepseek Coder 1.3B Instruct Q4_K_M"},
		{input: "deepseek-coder:1.3b-instruct-q5_0", expected: "Deepseek Coder 1.3B Instruct Q5 0"},
		{input: "deepseek-coder:1.3b-instruct-q5_1", expected: "Deepseek Coder 1.3B Instruct Q5 1"},
		{input: "deepseek-coder:1.3b-instruct-q5_K_S", expected: "Deepseek Coder 1.3B Instruct Q5 K_S"},
		{input: "deepseek-coder:1.3b-instruct-q5_K_M", expected: "Deepseek Coder 1.3B Instruct Q5 K_M"},
		{input: "deepseek-coder:1.3b-instruct-q6_K", expected: "Deepseek Coder 1.3B Instruct Q6 K"},
		{input: "deepseek-coder:1.3b-instruct-q8_0", expected: "Deepseek Coder 1.3B Instruct Q8_0"},
		{input: "deepseek-coder:1.3b-instruct-fp16", expected: "Deepseek Coder 1.3B Instruct FP16"},
		{input: "deepseek-coder:6.7b-base", expected: "Deepseek Coder 6.7B Base"},
		{input: "deepseek-coder:6.7b-base-q2_K", expected: "Deepseek Coder 6.7B Base Q2_K"},
		{input: "deepseek-coder:6.7b-base-q3_K_S", expected: "Deepseek Coder 6.7B Base Q3 K_S"},
		{input: "deepseek-coder:6.7b-base-q3_K_M", expected: "Deepseek Coder 6.7B Base Q3 K_M"},
		{input: "deepseek-coder:6.7b-base-q3_K_L", expected: "Deepseek Coder 6.7B Base Q3 K_L"},
		{input: "deepseek-coder:6.7b-base-q4_0", expected: "Deepseek Coder 6.7B Base Q4_0"},
		{input: "deepseek-coder:6.7b-base-q4_1", expected: "Deepseek Coder 6.7B Base Q4 1"},
		{input: "deepseek-coder:6.7b-base-q4_K_S", expected: "Deepseek Coder 6.7B Base Q4 K_S"},
		{input: "deepseek-coder:6.7b-base-q4_K_M", expected: "Deepseek Coder 6.7B Base Q4_K_M"},
		{input: "deepseek-coder:6.7b-base-q5_0", expected: "Deepseek Coder 6.7B Base Q5 0"},
		{input: "deepseek-coder:6.7b-base-q5_1", expected: "Deepseek Coder 6.7B Base Q5 1"},
		{input: "deepseek-coder:6.7b-base-q5_K_S", expected: "Deepseek Coder 6.7B Base Q5 K_S"},
		{input: "deepseek-coder:6.7b-base-q5_K_M", expected: "Deepseek Coder 6.7B Base Q5 K_M"},
		{input: "deepseek-coder:6.7b-base-q6_K", expected: "Deepseek Coder 6.7B Base Q6 K"},
		{input: "deepseek-coder:6.7b-base-q8_0", expected: "Deepseek Coder 6.7B Base Q8_0"},
		{input: "deepseek-coder:6.7b-base-fp16", expected: "Deepseek Coder 6.7B Base FP16"},
		{input: "deepseek-coder:6.7b-instruct", expected: "Deepseek Coder 6.7B Instruct"},
		{input: "deepseek-coder:6.7b-instruct-q2_K", expected: "Deepseek Coder 6.7B Instruct Q2_K"},
		{input: "deepseek-coder:6.7b-instruct-q3_K_S", expected: "Deepseek Coder 6.7B Instruct Q3 K_S"},
		{input: "deepseek-coder:6.7b-instruct-q3_K_M", expected: "Deepseek Coder 6.7B Instruct Q3 K_M"},
		{input: "deepseek-coder:6.7b-instruct-q3_K_L", expected: "Deepseek Coder 6.7B Instruct Q3 K_L"},
		{input: "deepseek-coder:6.7b-instruct-q4_0", expected: "Deepseek Coder 6.7B Instruct Q4_0"},
		{input: "deepseek-coder:6.7b-instruct-q4_1", expected: "Deepseek Coder 6.7B Instruct Q4 1"},
		{input: "deepseek-coder:6.7b-instruct-q4_K_S", expected: "Deepseek Coder 6.7B Instruct Q4 K_S"},
		{input: "deepseek-coder:6.7b-instruct-q4_K_M", expected: "Deepseek Coder 6.7B Instruct Q4_K_M"},
		{input: "deepseek-coder:6.7b-instruct-q5_0", expected: "Deepseek Coder 6.7B Instruct Q5 0"},
		{input: "deepseek-coder:6.7b-instruct-q5_1", expected: "Deepseek Coder 6.7B Instruct Q5 1"},
		{input: "deepseek-coder:6.7b-instruct-q5_K_S", expected: "Deepseek Coder 6.7B Instruct Q5 K_S"},
		{input: "deepseek-coder:6.7b-instruct-q5_K_M", expected: "Deepseek Coder 6.7B Instruct Q5 K_M"},
		{input: "deepseek-coder:6.7b-instruct-q6_K", expected: "Deepseek Coder 6.7B Instruct Q6 K"},
		{input: "deepseek-coder:6.7b-instruct-q8_0", expected: "Deepseek Coder 6.7B Instruct Q8_0"},
		{input: "deepseek-coder:6.7b-instruct-fp16", expected: "Deepseek Coder 6.7B Instruct FP16"},
		{input: "deepseek-coder:33b-base", expected: "Deepseek Coder 33B Base"},
		{input: "deepseek-coder:33b-base-q2_K", expected: "Deepseek Coder 33B Base Q2_K"},
		{input: "deepseek-coder:33b-base-q3_K_S", expected: "Deepseek Coder 33B Base Q3 K_S"},
		{input: "deepseek-coder:33b-base-q3_K_M", expected: "Deepseek Coder 33B Base Q3 K_M"},
		{input: "deepseek-coder:33b-base-q3_K_L", expected: "Deepseek Coder 33B Base Q3 K_L"},
		{input: "deepseek-coder:33b-base-q4_0", expected: "Deepseek Coder 33B Base Q4_0"},
		{input: "deepseek-coder:33b-base-q4_1", expected: "Deepseek Coder 33B Base Q4 1"},
		{input: "deepseek-coder:33b-base-q4_K_S", expected: "Deepseek Coder 33B Base Q4 K_S"},
		{input: "deepseek-coder:33b-base-q4_K_M", expected: "Deepseek Coder 33B Base Q4_K_M"},
		{input: "deepseek-coder:33b-base-q5_0", expected: "Deepseek Coder 33B Base Q5 0"},
		{input: "deepseek-coder:33b-base-q5_1", expected: "Deepseek Coder 33B Base Q5 1"},
		{input: "deepseek-coder:33b-base-q5_K_S", expected: "Deepseek Coder 33B Base Q5 K_S"},
		{input: "deepseek-coder:33b-base-q5_K_M", expected: "Deepseek Coder 33B Base Q5 K_M"},
		{input: "deepseek-coder:33b-base-q6_K", expected: "Deepseek Coder 33B Base Q6 K"},
		{input: "deepseek-coder:33b-base-q8_0", expected: "Deepseek Coder 33B Base Q8_0"},
		{input: "deepseek-coder:33b-base-fp16", expected: "Deepseek Coder 33B Base FP16"},
		{input: "deepseek-coder:33b-instruct", expected: "Deepseek Coder 33B Instruct"},
		{input: "deepseek-coder:33b-instruct-q2_K", expected: "Deepseek Coder 33B Instruct Q2_K"},
		{input: "deepseek-coder:33b-instruct-q3_K_S", expected: "Deepseek Coder 33B Instruct Q3 K_S"},
		{input: "deepseek-coder:33b-instruct-q3_K_M", expected: "Deepseek Coder 33B Instruct Q3 K_M"},
		{input: "deepseek-coder:33b-instruct-q3_K_L", expected: "Deepseek Coder 33B Instruct Q3 K_L"},
		{input: "deepseek-coder:33b-instruct-q4_0", expected: "Deepseek Coder 33B Instruct Q4_0"},
		{input: "deepseek-coder:33b-instruct-q4_1", expected: "Deepseek Coder 33B Instruct Q4 1"},
		{input: "deepseek-coder:33b-instruct-q4_K_S", expected: "Deepseek Coder 33B Instruct Q4 K_S"},
		{input: "deepseek-coder:33b-instruct-q4_K_M", expected: "Deepseek Coder 33B Instruct Q4_K_M"},
		{input: "deepseek-coder:33b-instruct-q5_0", expected: "Deepseek Coder 33B Instruct Q5 0"},
		{input: "deepseek-coder:33b-instruct-q5_1", expected: "Deepseek Coder 33B Instruct Q5 1"},
		{input: "deepseek-coder:33b-instruct-q5_K_S", expected: "Deepseek Coder 33B Instruct Q5 K_S"},
		{input: "deepseek-coder:33b-instruct-q5_K_M", expected: "Deepseek Coder 33B Instruct Q5 K_M"},
		{input: "deepseek-coder:33b-instruct-q6_K", expected: "Deepseek Coder 33B Instruct Q6 K"},
		{input: "deepseek-coder:33b-instruct-q8_0", expected: "Deepseek Coder 33B Instruct Q8_0"},
		{input: "deepseek-coder:33b-instruct-fp16", expected: "Deepseek Coder 33B Instruct FP16"},
		{input: "smollm2:latest", expected: "SmolLM2 (latest)"},
		{input: "smollm2", expected: "SmolLM2"},
		{input: "smollm2:135m", expected: "SmolLM2 135M"},
		{input: "smollm2:360m", expected: "SmolLM2 360M"},
		{input: "smollm2:1.7b", expected: "SmolLM2 1.7B"},
		{input: "smollm2:135m-instruct-q2_K", expected: "SmolLM2 135M Instruct Q2_K"},
		{input: "smollm2:135m-instruct-q3_K_S", expected: "SmolLM2 135M Instruct Q3 K_S"},
		{input: "smollm2:135m-instruct-q3_K_M", expected: "SmolLM2 135M Instruct Q3 K_M"},
		{input: "smollm2:135m-instruct-q3_K_L", expected: "SmolLM2 135M Instruct Q3 K_L"},
		{input: "smollm2:135m-instruct-q4_0", expected: "SmolLM2 135M Instruct Q4_0"},
		{input: "smollm2:135m-instruct-q4_1", expected: "SmolLM2 135M Instruct Q4 1"},
		{input: "smollm2:135m-instruct-q4_K_S", expected: "SmolLM2 135M Instruct Q4 K_S"},
		{input: "smollm2:135m-instruct-q4_K_M", expected: "SmolLM2 135M Instruct Q4_K_M"},
		{input: "smollm2:135m-instruct-q5_0", expected: "SmolLM2 135M Instruct Q5 0"},
		{input: "smollm2:135m-instruct-q5_1", expected: "SmolLM2 135M Instruct Q5 1"},
		{input: "smollm2:135m-instruct-q5_K_S", expected: "SmolLM2 135M Instruct Q5 K_S"},
		{input: "smollm2:135m-instruct-q5_K_M", expected: "SmolLM2 135M Instruct Q5 K_M"},
		{input: "smollm2:135m-instruct-q6_K", expected: "SmolLM2 135M Instruct Q6 K"},
		{input: "smollm2:135m-instruct-q8_0", expected: "SmolLM2 135M Instruct Q8_0"},
		{input: "smollm2:135m-instruct-fp16", expected: "SmolLM2 135M Instruct FP16"},
		{input: "smollm2:360m-instruct-q2_K", expected: "SmolLM2 360M Instruct Q2_K"},
		{input: "smollm2:360m-instruct-q3_K_S", expected: "SmolLM2 360M Instruct Q3 K_S"},
		{input: "smollm2:360m-instruct-q3_K_M", expected: "SmolLM2 360M Instruct Q3 K_M"},
		{input: "smollm2:360m-instruct-q3_K_L", expected: "SmolLM2 360M Instruct Q3 K_L"},
		{input: "smollm2:360m-instruct-q4_0", expected: "SmolLM2 360M Instruct Q4_0"},
		{input: "smollm2:360m-instruct-q4_1", expected: "SmolLM2 360M Instruct Q4 1"},
		{input: "smollm2:360m-instruct-q4_K_S", expected: "SmolLM2 360M Instruct Q4 K_S"},
		{input: "smollm2:360m-instruct-q4_K_M", expected: "SmolLM2 360M Instruct Q4_K_M"},
		{input: "smollm2:360m-instruct-q5_0", expected: "SmolLM2 360M Instruct Q5 0"},
		{input: "smollm2:360m-instruct-q5_1", expected: "SmolLM2 360M Instruct Q5 1"},
		{input: "smollm2:360m-instruct-q5_K_S", expected: "SmolLM2 360M Instruct Q5 K_S"},
		{input: "smollm2:360m-instruct-q5_K_M", expected: "SmolLM2 360M Instruct Q5 K_M"},
		{input: "smollm2:360m-instruct-q6_K", expected: "SmolLM2 360M Instruct Q6 K"},
		{input: "smollm2:360m-instruct-q8_0", expected: "SmolLM2 360M Instruct Q8_0"},
		{input: "smollm2:360m-instruct-fp16", expected: "SmolLM2 360M Instruct FP16"},
		{input: "smollm2:1.7b-instruct-q2_K", expected: "SmolLM2 1.7B Instruct Q2_K"},
		{input: "smollm2:1.7b-instruct-q3_K_S", expected: "SmolLM2 1.7B Instruct Q3 K_S"},
		{input: "smollm2:1.7b-instruct-q3_K_M", expected: "SmolLM2 1.7B Instruct Q3 K_M"},
		{input: "smollm2:1.7b-instruct-q3_K_L", expected: "SmolLM2 1.7B Instruct Q3 K_L"},
		{input: "smollm2:1.7b-instruct-q4_0", expected: "SmolLM2 1.7B Instruct Q4_0"},
		{input: "smollm2:1.7b-instruct-q4_1", expected: "SmolLM2 1.7B Instruct Q4 1"},
		{input: "smollm2:1.7b-instruct-q4_K_S", expected: "SmolLM2 1.7B Instruct Q4 K_S"},
		{input: "smollm2:1.7b-instruct-q4_K_M", expected: "SmolLM2 1.7B Instruct Q4_K_M"},
		{input: "smollm2:1.7b-instruct-q5_0", expected: "SmolLM2 1.7B Instruct Q5 0"},
		{input: "smollm2:1.7b-instruct-q5_1", expected: "SmolLM2 1.7B Instruct Q5 1"},
		{input: "smollm2:1.7b-instruct-q5_K_S", expected: "SmolLM2 1.7B Instruct Q5 K_S"},
		{input: "smollm2:1.7b-instruct-q5_K_M", expected: "SmolLM2 1.7B Instruct Q5 K_M"},
		{input: "smollm2:1.7b-instruct-q6_K", expected: "SmolLM2 1.7B Instruct Q6 K"},
		{input: "smollm2:1.7b-instruct-q8_0", expected: "SmolLM2 1.7B Instruct Q8_0"},
		{input: "smollm2:1.7b-instruct-fp16", expected: "SmolLM2 1.7B Instruct FP16"},
		{input: "mistral-small:latest", expected: "Mistral Small (latest)"},
		{input: "mistral-small", expected: "Mistral Small"},
		{input: "mistral-small:22b", expected: "Mistral Small 22B"},
		{input: "mistral-small:24b", expected: "Mistral Small 24B"},
		{input: "mistral-small:22b-instruct-2409-q2_K", expected: "Mistral Small 22B Instruct 2409 Q2_K"},
		{input: "mistral-small:22b-instruct-2409-q3_K_S", expected: "Mistral Small 22B Instruct 2409 Q3 K_S"},
		{input: "mistral-small:22b-instruct-2409-q3_K_M", expected: "Mistral Small 22B Instruct 2409 Q3 K_M"},
		{input: "mistral-small:22b-instruct-2409-q3_K_L", expected: "Mistral Small 22B Instruct 2409 Q3 K_L"},
		{input: "mistral-small:22b-instruct-2409-q4_0", expected: "Mistral Small 22B Instruct 2409 Q4_0"},
		{input: "mistral-small:22b-instruct-2409-q4_1", expected: "Mistral Small 22B Instruct 2409 Q4 1"},
		{input: "mistral-small:22b-instruct-2409-q4_K_S", expected: "Mistral Small 22B Instruct 2409 Q4 K_S"},
		{input: "mistral-small:22b-instruct-2409-q4_K_M", expected: "Mistral Small 22B Instruct 2409 Q4_K_M"},
		{input: "mistral-small:22b-instruct-2409-q5_0", expected: "Mistral Small 22B Instruct 2409 Q5 0"},
		{input: "mistral-small:22b-instruct-2409-q5_1", expected: "Mistral Small 22B Instruct 2409 Q5 1"},
		{input: "mistral-small:22b-instruct-2409-q5_K_S", expected: "Mistral Small 22B Instruct 2409 Q5 K_S"},
		{input: "mistral-small:22b-instruct-2409-q5_K_M", expected: "Mistral Small 22B Instruct 2409 Q5 K_M"},
		{input: "mistral-small:22b-instruct-2409-q6_K", expected: "Mistral Small 22B Instruct 2409 Q6 K"},
		{input: "mistral-small:22b-instruct-2409-q8_0", expected: "Mistral Small 22B Instruct 2409 Q8_0"},
		{input: "mistral-small:22b-instruct-2409-fp16", expected: "Mistral Small 22B Instruct 2409 FP16"},
		{input: "mistral-small:24b-instruct-2501-q4_K_M", expected: "Mistral Small 24B Instruct 2501 Q4_K_M"},
		{input: "mistral-small:24b-instruct-2501-q8_0", expected: "Mistral Small 24B Instruct 2501 Q8_0"},
		{input: "mistral-small:24b-instruct-2501-fp16", expected: "Mistral Small 24B Instruct 2501 FP16"},
		{input: "all-minilm:latest", expected: "All Minilm (latest)"},
		{input: "all-minilm", expected: "All Minilm"},
		{input: "all-minilm:l12", expected: "All Minilm L12"},
		{input: "all-minilm:l6", expected: "All Minilm L6"},
		{input: "all-minilm:v2", expected: "All Minilm v2"},
		{input: "all-minilm:22m", expected: "All Minilm 22M"},
		{input: "all-minilm:33m", expected: "All Minilm 33M"},
		{input: "all-minilm:22m-l6-v2-fp16", expected: "All Minilm 22M L6 v2 FP16"},
		{input: "all-minilm:33m-l12-v2-fp16", expected: "All Minilm 33M L12 v2 FP16"},
		{input: "all-minilm:l12-v2", expected: "All Minilm L12 v2"},
		{input: "all-minilm:l6-v2", expected: "All Minilm L6 v2"},
		{input: "llava-llama3:latest", expected: "Llava Llama3 (latest)"},
		{input: "llava-llama3", expected: "Llava Llama3"},
		{input: "llava-llama3:8b", expected: "Llava Llama3 8B"},
		{input: "llava-llama3:8b-v1.1-q4_0", expected: "Llava Llama3 8B v1.1 Q4_0"},
		{input: "llava-llama3:8b-v1.1-fp16", expected: "Llava Llama3 8B v1.1 FP16"},
		{input: "qwq:latest", expected: "Qwq (latest)"},
		{input: "qwq", expected: "Qwq"},
		{input: "qwq:32b", expected: "Qwq 32B"},
		{input: "qwq:32b-preview-q4_K_M", expected: "Qwq 32B Preview Q4_K_M"},
		{input: "qwq:32b-preview-q8_0", expected: "Qwq 32B Preview Q8_0"},
		{input: "qwq:32b-preview-fp16", expected: "Qwq 32B Preview FP16"},
		{input: "qwq:32b-q4_K_M", expected: "Qwq 32B Q4_K_M"},
		{input: "qwq:32b-q8_0", expected: "Qwq 32B Q8_0"},
		{input: "qwq:32b-fp16", expected: "Qwq 32B FP16"},
		{input: "codegemma:latest", expected: "Codegemma (latest)"},
		{input: "codegemma", expected: "Codegemma"},
		{input: "codegemma:code", expected: "Codegemma Code"},
		{input: "codegemma:instruct", expected: "Codegemma Instruct"},
		{input: "codegemma:2b", expected: "Codegemma 2B"},
		{input: "codegemma:7b", expected: "Codegemma 7B"},
		{input: "codegemma:2b-code", expected: "Codegemma 2B Code"},
		{input: "codegemma:2b-code-q2_K", expected: "Codegemma 2B Code Q2_K"},
		{input: "codegemma:2b-code-v1.1-q2_K", expected: "Codegemma 2B Code v1.1 Q2_K"},
		{input: "codegemma:2b-code-q3_K_S", expected: "Codegemma 2B Code Q3 K_S"},
		{input: "codegemma:2b-code-v1.1-q3_K_S", expected: "Codegemma 2B Code v1.1 Q3 K_S"},
		{input: "codegemma:2b-code-q3_K_M", expected: "Codegemma 2B Code Q3 K_M"},
		{input: "codegemma:2b-code-v1.1-q3_K_M", expected: "Codegemma 2B Code v1.1 Q3 K_M"},
		{input: "codegemma:2b-code-q3_K_L", expected: "Codegemma 2B Code Q3 K_L"},
		{input: "codegemma:2b-code-v1.1-q3_K_L", expected: "Codegemma 2B Code v1.1 Q3 K_L"},
		{input: "codegemma:2b-code-q4_0", expected: "Codegemma 2B Code Q4_0"},
		{input: "codegemma:2b-code-v1.1-q4_0", expected: "Codegemma 2B Code v1.1 Q4_0"},
		{input: "codegemma:2b-code-q4_1", expected: "Codegemma 2B Code Q4 1"},
		{input: "codegemma:2b-code-v1.1-q4_1", expected: "Codegemma 2B Code v1.1 Q4 1"},
		{input: "codegemma:2b-code-q4_K_S", expected: "Codegemma 2B Code Q4 K_S"},
		{input: "codegemma:2b-code-v1.1-q4_K_S", expected: "Codegemma 2B Code v1.1 Q4 K_S"},
		{input: "codegemma:2b-code-q4_K_M", expected: "Codegemma 2B Code Q4_K_M"},
		{input: "codegemma:2b-code-v1.1-q4_K_M", expected: "Codegemma 2B Code v1.1 Q4_K_M"},
		{input: "codegemma:2b-code-q5_0", expected: "Codegemma 2B Code Q5 0"},
		{input: "codegemma:2b-code-v1.1-q5_0", expected: "Codegemma 2B Code v1.1 Q5 0"},
		{input: "codegemma:2b-code-q5_1", expected: "Codegemma 2B Code Q5 1"},
		{input: "codegemma:2b-code-v1.1-q5_1", expected: "Codegemma 2B Code v1.1 Q5 1"},
		{input: "codegemma:2b-code-q5_K_S", expected: "Codegemma 2B Code Q5 K_S"},
		{input: "codegemma:2b-code-v1.1-q5_K_S", expected: "Codegemma 2B Code v1.1 Q5 K_S"},
		{input: "codegemma:2b-code-q5_K_M", expected: "Codegemma 2B Code Q5 K_M"},
		{input: "codegemma:2b-code-v1.1-q5_K_M", expected: "Codegemma 2B Code v1.1 Q5 K_M"},
		{input: "codegemma:2b-code-q6_K", expected: "Codegemma 2B Code Q6 K"},
		{input: "codegemma:2b-code-v1.1-q6_K", expected: "Codegemma 2B Code v1.1 Q6 K"},
		{input: "codegemma:2b-code-q8_0", expected: "Codegemma 2B Code Q8_0"},
		{input: "codegemma:2b-code-v1.1-q8_0", expected: "Codegemma 2B Code v1.1 Q8_0"},
		{input: "codegemma:2b-code-fp16", expected: "Codegemma 2B Code FP16"},
		{input: "codegemma:2b-code-v1.1-fp16", expected: "Codegemma 2B Code v1.1 FP16"},
		{input: "codegemma:2b-v1.1", expected: "Codegemma 2B v1.1"},
		{input: "codegemma:7b-code", expected: "Codegemma 7B Code"},
		{input: "codegemma:7b-code-q2_K", expected: "Codegemma 7B Code Q2_K"},
		{input: "codegemma:7b-code-q3_K_S", expected: "Codegemma 7B Code Q3 K_S"},
		{input: "codegemma:7b-code-q3_K_M", expected: "Codegemma 7B Code Q3 K_M"},
		{input: "codegemma:7b-code-q3_K_L", expected: "Codegemma 7B Code Q3 K_L"},
		{input: "codegemma:7b-code-q4_0", expected: "Codegemma 7B Code Q4_0"},
		{input: "codegemma:7b-code-q4_1", expected: "Codegemma 7B Code Q4 1"},
		{input: "codegemma:7b-code-q4_K_S", expected: "Codegemma 7B Code Q4 K_S"},
		{input: "codegemma:7b-code-q4_K_M", expected: "Codegemma 7B Code Q4_K_M"},
		{input: "codegemma:7b-code-q5_0", expected: "Codegemma 7B Code Q5 0"},
		{input: "codegemma:7b-code-q5_1", expected: "Codegemma 7B Code Q5 1"},
		{input: "codegemma:7b-code-q5_K_S", expected: "Codegemma 7B Code Q5 K_S"},
		{input: "codegemma:7b-code-q5_K_M", expected: "Codegemma 7B Code Q5 K_M"},
		{input: "codegemma:7b-code-q6_K", expected: "Codegemma 7B Code Q6 K"},
		{input: "codegemma:7b-code-q8_0", expected: "Codegemma 7B Code Q8_0"},
		{input: "codegemma:7b-code-fp16", expected: "Codegemma 7B Code FP16"},
		{input: "codegemma:7b-instruct", expected: "Codegemma 7B Instruct"},
		{input: "codegemma:7b-instruct-q2_K", expected: "Codegemma 7B Instruct Q2_K"},
		{input: "codegemma:7b-instruct-v1.1-q2_K", expected: "Codegemma 7B Instruct v1.1 Q2_K"},
		{input: "codegemma:7b-instruct-q3_K_S", expected: "Codegemma 7B Instruct Q3 K_S"},
		{input: "codegemma:7b-instruct-v1.1-q3_K_S", expected: "Codegemma 7B Instruct v1.1 Q3 K_S"},
		{input: "codegemma:7b-instruct-q3_K_M", expected: "Codegemma 7B Instruct Q3 K_M"},
		{input: "codegemma:7b-instruct-v1.1-q3_K_M", expected: "Codegemma 7B Instruct v1.1 Q3 K_M"},
		{input: "codegemma:7b-instruct-q3_K_L", expected: "Codegemma 7B Instruct Q3 K_L"},
		{input: "codegemma:7b-instruct-v1.1-q3_K_L", expected: "Codegemma 7B Instruct v1.1 Q3 K_L"},
		{input: "codegemma:7b-instruct-q4_0", expected: "Codegemma 7B Instruct Q4_0"},
		{input: "codegemma:7b-instruct-v1.1-q4_0", expected: "Codegemma 7B Instruct v1.1 Q4_0"},
		{input: "codegemma:7b-instruct-q4_1", expected: "Codegemma 7B Instruct Q4 1"},
		{input: "codegemma:7b-instruct-v1.1-q4_1", expected: "Codegemma 7B Instruct v1.1 Q4 1"},
		{input: "codegemma:7b-instruct-q4_K_S", expected: "Codegemma 7B Instruct Q4 K_S"},
		{input: "codegemma:7b-instruct-v1.1-q4_K_S", expected: "Codegemma 7B Instruct v1.1 Q4 K_S"},
		{input: "codegemma:7b-instruct-q4_K_M", expected: "Codegemma 7B Instruct Q4_K_M"},
		{input: "codegemma:7b-instruct-v1.1-q4_K_M", expected: "Codegemma 7B Instruct v1.1 Q4_K_M"},
		{input: "codegemma:7b-instruct-q5_0", expected: "Codegemma 7B Instruct Q5 0"},
		{input: "codegemma:7b-instruct-v1.1-q5_0", expected: "Codegemma 7B Instruct v1.1 Q5 0"},
		{input: "codegemma:7b-instruct-q5_1", expected: "Codegemma 7B Instruct Q5 1"},
		{input: "codegemma:7b-instruct-v1.1-q5_1", expected: "Codegemma 7B Instruct v1.1 Q5 1"},
		{input: "codegemma:7b-instruct-q5_K_S", expected: "Codegemma 7B Instruct Q5 K_S"},
		{input: "codegemma:7b-instruct-v1.1-q5_K_S", expected: "Codegemma 7B Instruct v1.1 Q5 K_S"},
		{input: "codegemma:7b-instruct-q5_K_M", expected: "Codegemma 7B Instruct Q5 K_M"},
		{input: "codegemma:7b-instruct-v1.1-q5_K_M", expected: "Codegemma 7B Instruct v1.1 Q5 K_M"},
		{input: "codegemma:7b-instruct-q6_K", expected: "Codegemma 7B Instruct Q6 K"},
		{input: "codegemma:7b-instruct-v1.1-q6_K", expected: "Codegemma 7B Instruct v1.1 Q6 K"},
		{input: "codegemma:7b-instruct-q8_0", expected: "Codegemma 7B Instruct Q8_0"},
		{input: "codegemma:7b-instruct-v1.1-q8_0", expected: "Codegemma 7B Instruct v1.1 Q8_0"},
		{input: "codegemma:7b-instruct-fp16", expected: "Codegemma 7B Instruct FP16"},
		{input: "codegemma:7b-instruct-v1.1-fp16", expected: "Codegemma 7B Instruct v1.1 FP16"},
		{input: "codegemma:7b-v1.1", expected: "Codegemma 7B v1.1"},
		{input: "starcoder2:latest", expected: "StarCoder2 (latest)"},
		{input: "starcoder2", expected: "StarCoder2"},
		{input: "starcoder2:instruct", expected: "StarCoder2 Instruct"},
		{input: "starcoder2:3b", expected: "StarCoder2 3B"},
		{input: "starcoder2:7b", expected: "StarCoder2 7B"},
		{input: "starcoder2:15b", expected: "StarCoder2 15B"},
		{input: "starcoder2:3b-q2_K", expected: "StarCoder2 3B Q2_K"},
		{input: "starcoder2:3b-q3_K_S", expected: "StarCoder2 3B Q3 K_S"},
		{input: "starcoder2:3b-q3_K_M", expected: "StarCoder2 3B Q3 K_M"},
		{input: "starcoder2:3b-q3_K_L", expected: "StarCoder2 3B Q3 K_L"},
		{input: "starcoder2:3b-q4_0", expected: "StarCoder2 3B Q4_0"},
		{input: "starcoder2:3b-q4_1", expected: "StarCoder2 3B Q4 1"},
		{input: "starcoder2:3b-q4_K_S", expected: "StarCoder2 3B Q4 K_S"},
		{input: "starcoder2:3b-q4_K_M", expected: "StarCoder2 3B Q4_K_M"},
		{input: "starcoder2:3b-q5_0", expected: "StarCoder2 3B Q5 0"},
		{input: "starcoder2:3b-q5_1", expected: "StarCoder2 3B Q5 1"},
		{input: "starcoder2:3b-q5_K_S", expected: "StarCoder2 3B Q5 K_S"},
		{input: "starcoder2:3b-q5_K_M", expected: "StarCoder2 3B Q5 K_M"},
		{input: "starcoder2:3b-q6_K", expected: "StarCoder2 3B Q6 K"},
		{input: "starcoder2:3b-q8_0", expected: "StarCoder2 3B Q8_0"},
		{input: "starcoder2:3b-fp16", expected: "StarCoder2 3B FP16"},
		{input: "starcoder2:7b-q2_K", expected: "StarCoder2 7B Q2_K"},
		{input: "starcoder2:7b-q3_K_S", expected: "StarCoder2 7B Q3 K_S"},
		{input: "starcoder2:7b-q3_K_M", expected: "StarCoder2 7B Q3 K_M"},
		{input: "starcoder2:7b-q3_K_L", expected: "StarCoder2 7B Q3 K_L"},
		{input: "starcoder2:7b-q4_0", expected: "StarCoder2 7B Q4_0"},
		{input: "starcoder2:7b-q4_1", expected: "StarCoder2 7B Q4 1"},
		{input: "starcoder2:7b-q4_K_S", expected: "StarCoder2 7B Q4 K_S"},
		{input: "starcoder2:7b-q4_K_M", expected: "StarCoder2 7B Q4_K_M"},
		{input: "starcoder2:7b-q5_0", expected: "StarCoder2 7B Q5 0"},
		{input: "starcoder2:7b-q5_1", expected: "StarCoder2 7B Q5 1"},
		{input: "starcoder2:7b-q5_K_S", expected: "StarCoder2 7B Q5 K_S"},
		{input: "starcoder2:7b-q5_K_M", expected: "StarCoder2 7B Q5 K_M"},
		{input: "starcoder2:7b-q6_K", expected: "StarCoder2 7B Q6 K"},
		{input: "starcoder2:7b-q8_0", expected: "StarCoder2 7B Q8_0"},
		{input: "starcoder2:7b-fp16", expected: "StarCoder2 7B FP16"},
		{input: "starcoder2:15b-instruct", expected: "StarCoder2 15B Instruct"},
		{input: "starcoder2:15b-instruct-v0.1-q2_K", expected: "StarCoder2 15B Instruct v0.1 Q2_K"},
		{input: "starcoder2:15b-instruct-v0.1-q3_K_S", expected: "StarCoder2 15B Instruct v0.1 Q3 K_S"},
		{input: "starcoder2:15b-instruct-v0.1-q3_K_M", expected: "StarCoder2 15B Instruct v0.1 Q3 K_M"},
		{input: "starcoder2:15b-instruct-v0.1-q3_K_L", expected: "StarCoder2 15B Instruct v0.1 Q3 K_L"},
		{input: "starcoder2:15b-instruct-q4_0", expected: "StarCoder2 15B Instruct Q4_0"},
		{input: "starcoder2:15b-instruct-v0.1-q4_0", expected: "StarCoder2 15B Instruct v0.1 Q4_0"},
		{input: "starcoder2:15b-instruct-v0.1-q4_1", expected: "StarCoder2 15B Instruct v0.1 Q4 1"},
		{input: "starcoder2:15b-instruct-v0.1-q4_K_S", expected: "StarCoder2 15B Instruct v0.1 Q4 K_S"},
		{input: "starcoder2:15b-instruct-v0.1-q4_K_M", expected: "StarCoder2 15B Instruct v0.1 Q4_K_M"},
		{input: "starcoder2:15b-instruct-v0.1-q5_0", expected: "StarCoder2 15B Instruct v0.1 Q5 0"},
		{input: "starcoder2:15b-instruct-v0.1-q5_1", expected: "StarCoder2 15B Instruct v0.1 Q5 1"},
		{input: "starcoder2:15b-instruct-v0.1-q5_K_S", expected: "StarCoder2 15B Instruct v0.1 Q5 K_S"},
		{input: "starcoder2:15b-instruct-v0.1-q5_K_M", expected: "StarCoder2 15B Instruct v0.1 Q5 K_M"},
		{input: "starcoder2:15b-instruct-v0.1-q6_K", expected: "StarCoder2 15B Instruct v0.1 Q6 K"},
		{input: "starcoder2:15b-instruct-v0.1-q8_0", expected: "StarCoder2 15B Instruct v0.1 Q8_0"},
		{input: "starcoder2:15b-instruct-v0.1-fp16", expected: "StarCoder2 15B Instruct v0.1 FP16"},
		{input: "starcoder2:15b-q2_K", expected: "StarCoder2 15B Q2_K"},
		{input: "starcoder2:15b-q3_K_S", expected: "StarCoder2 15B Q3 K_S"},
		{input: "starcoder2:15b-q3_K_M", expected: "StarCoder2 15B Q3 K_M"},
		{input: "starcoder2:15b-q3_K_L", expected: "StarCoder2 15B Q3 K_L"},
		{input: "starcoder2:15b-q4_0", expected: "StarCoder2 15B Q4_0"},
		{input: "starcoder2:15b-q4_1", expected: "StarCoder2 15B Q4 1"},
		{input: "starcoder2:15b-q4_K_S", expected: "StarCoder2 15B Q4 K_S"},
		{input: "starcoder2:15b-q4_K_M", expected: "StarCoder2 15B Q4_K_M"},
		{input: "starcoder2:15b-q5_0", expected: "StarCoder2 15B Q5 0"},
		{input: "starcoder2:15b-q5_1", expected: "StarCoder2 15B Q5 1"},
		{input: "starcoder2:15b-q5_K_S", expected: "StarCoder2 15B Q5 K_S"},
		{input: "starcoder2:15b-q5_K_M", expected: "StarCoder2 15B Q5 K_M"},
		{input: "starcoder2:15b-q6_K", expected: "StarCoder2 15B Q6 K"},
		{input: "starcoder2:15b-q8_0", expected: "StarCoder2 15B Q8_0"},
		{input: "starcoder2:15b-fp16", expected: "StarCoder2 15B FP16"},
		{input: "falcon3:latest", expected: "Falcon3 (latest)"},
		{input: "falcon3", expected: "Falcon3"},
		{input: "falcon3:1b", expected: "Falcon3 1B"},
		{input: "falcon3:3b", expected: "Falcon3 3B"},
		{input: "falcon3:7b", expected: "Falcon3 7B"},
		{input: "falcon3:10b", expected: "Falcon3 10B"},
		{input: "falcon3:1b-instruct-q4_K_M", expected: "Falcon3 1B Instruct Q4_K_M"},
		{input: "falcon3:1b-instruct-q8_0", expected: "Falcon3 1B Instruct Q8_0"},
		{input: "falcon3:1b-instruct-fp16", expected: "Falcon3 1B Instruct FP16"},
		{input: "falcon3:3b-instruct-q4_K_M", expected: "Falcon3 3B Instruct Q4_K_M"},
		{input: "falcon3:3b-instruct-q8_0", expected: "Falcon3 3B Instruct Q8_0"},
		{input: "falcon3:3b-instruct-fp16", expected: "Falcon3 3B Instruct FP16"},
		{input: "falcon3:7b-instruct-q4_K_M", expected: "Falcon3 7B Instruct Q4_K_M"},
		{input: "falcon3:7b-instruct-q8_0", expected: "Falcon3 7B Instruct Q8_0"},
		{input: "falcon3:7b-instruct-fp16", expected: "Falcon3 7B Instruct FP16"},
		{input: "falcon3:10b-instruct-q4_K_M", expected: "Falcon3 10B Instruct Q4_K_M"},
		{input: "falcon3:10b-instruct-q8_0", expected: "Falcon3 10B Instruct Q8_0"},
		{input: "falcon3:10b-instruct-fp16", expected: "Falcon3 10B Instruct FP16"},
		{input: "granite3.1-moe:latest", expected: "Granite3.1 Moe (latest)"},
		{input: "granite3.1-moe", expected: "Granite3.1 Moe"},
		{input: "granite3.1-moe:1b", expected: "Granite3.1 Moe 1B"},
		{input: "granite3.1-moe:3b", expected: "Granite3.1 Moe 3B"},
		{input: "granite3.1-moe:1b-instruct-q2_K", expected: "Granite3.1 Moe 1B Instruct Q2_K"},
		{input: "granite3.1-moe:1b-instruct-q3_K_S", expected: "Granite3.1 Moe 1B Instruct Q3 K_S"},
		{input: "granite3.1-moe:1b-instruct-q3_K_M", expected: "Granite3.1 Moe 1B Instruct Q3 K_M"},
		{input: "granite3.1-moe:1b-instruct-q3_K_L", expected: "Granite3.1 Moe 1B Instruct Q3 K_L"},
		{input: "granite3.1-moe:1b-instruct-q4_0", expected: "Granite3.1 Moe 1B Instruct Q4_0"},
		{input: "granite3.1-moe:1b-instruct-q4_1", expected: "Granite3.1 Moe 1B Instruct Q4 1"},
		{input: "granite3.1-moe:1b-instruct-q4_K_S", expected: "Granite3.1 Moe 1B Instruct Q4 K_S"},
		{input: "granite3.1-moe:1b-instruct-q4_K_M", expected: "Granite3.1 Moe 1B Instruct Q4_K_M"},
		{input: "granite3.1-moe:1b-instruct-q5_0", expected: "Granite3.1 Moe 1B Instruct Q5 0"},
		{input: "granite3.1-moe:1b-instruct-q5_1", expected: "Granite3.1 Moe 1B Instruct Q5 1"},
		{input: "granite3.1-moe:1b-instruct-q5_K_S", expected: "Granite3.1 Moe 1B Instruct Q5 K_S"},
		{input: "granite3.1-moe:1b-instruct-q5_K_M", expected: "Granite3.1 Moe 1B Instruct Q5 K_M"},
		{input: "granite3.1-moe:1b-instruct-q6_K", expected: "Granite3.1 Moe 1B Instruct Q6 K"},
		{input: "granite3.1-moe:1b-instruct-q8_0", expected: "Granite3.1 Moe 1B Instruct Q8_0"},
		{input: "granite3.1-moe:1b-instruct-fp16", expected: "Granite3.1 Moe 1B Instruct FP16"},
		{input: "granite3.1-moe:3b-instruct-q2_K", expected: "Granite3.1 Moe 3B Instruct Q2_K"},
		{input: "granite3.1-moe:3b-instruct-q3_K_S", expected: "Granite3.1 Moe 3B Instruct Q3 K_S"},
		{input: "granite3.1-moe:3b-instruct-q3_K_M", expected: "Granite3.1 Moe 3B Instruct Q3 K_M"},
		{input: "granite3.1-moe:3b-instruct-q3_K_L", expected: "Granite3.1 Moe 3B Instruct Q3 K_L"},
		{input: "granite3.1-moe:3b-instruct-q4_0", expected: "Granite3.1 Moe 3B Instruct Q4_0"},
		{input: "granite3.1-moe:3b-instruct-q4_1", expected: "Granite3.1 Moe 3B Instruct Q4 1"},
		{input: "granite3.1-moe:3b-instruct-q4_K_S", expected: "Granite3.1 Moe 3B Instruct Q4 K_S"},
		{input: "granite3.1-moe:3b-instruct-q4_K_M", expected: "Granite3.1 Moe 3B Instruct Q4_K_M"},
		{input: "granite3.1-moe:3b-instruct-q5_0", expected: "Granite3.1 Moe 3B Instruct Q5 0"},
		{input: "granite3.1-moe:3b-instruct-q5_1", expected: "Granite3.1 Moe 3B Instruct Q5 1"},
		{input: "granite3.1-moe:3b-instruct-q5_K_S", expected: "Granite3.1 Moe 3B Instruct Q5 K_S"},
		{input: "granite3.1-moe:3b-instruct-q5_K_M", expected: "Granite3.1 Moe 3B Instruct Q5 K_M"},
		{input: "granite3.1-moe:3b-instruct-q6_K", expected: "Granite3.1 Moe 3B Instruct Q6 K"},
		{input: "granite3.1-moe:3b-instruct-q8_0", expected: "Granite3.1 Moe 3B Instruct Q8_0"},
		{input: "granite3.1-moe:3b-instruct-fp16", expected: "Granite3.1 Moe 3B Instruct FP16"},
		{input: "mixtral:latest", expected: "Mixtral (latest)"},
		{input: "mixtral", expected: "Mixtral"},
		{input: "mixtral:instruct", expected: "Mixtral Instruct"},
		{input: "mixtral:text", expected: "Mixtral Text"},
		{input: "mixtral:v0.1", expected: "Mixtral v0.1"},
		{input: "mixtral:8x7b", expected: "Mixtral 8x7B"},
		{input: "mixtral:8x22b", expected: "Mixtral 8x22B"},
		{input: "mixtral:8x7b-instruct-v0.1-q2_K", expected: "Mixtral 8x7B Instruct v0.1 Q2_K"},
		{input: "mixtral:8x7b-instruct-v0.1-q3_K_S", expected: "Mixtral 8x7B Instruct v0.1 Q3 K_S"},
		{input: "mixtral:8x7b-instruct-v0.1-q3_K_M", expected: "Mixtral 8x7B Instruct v0.1 Q3 K_M"},
		{input: "mixtral:8x7b-instruct-v0.1-q3_K_L", expected: "Mixtral 8x7B Instruct v0.1 Q3 K_L"},
		{input: "mixtral:8x7b-instruct-v0.1-q4_0", expected: "Mixtral 8x7B Instruct v0.1 Q4_0"},
		{input: "mixtral:8x7b-instruct-v0.1-q4_1", expected: "Mixtral 8x7B Instruct v0.1 Q4 1"},
		{input: "mixtral:8x7b-instruct-v0.1-q4_K_S", expected: "Mixtral 8x7B Instruct v0.1 Q4 K_S"},
		{input: "mixtral:8x7b-instruct-v0.1-q4_K_M", expected: "Mixtral 8x7B Instruct v0.1 Q4_K_M"},
		{input: "mixtral:8x7b-instruct-v0.1-q5_0", expected: "Mixtral 8x7B Instruct v0.1 Q5 0"},
		{input: "mixtral:8x7b-instruct-v0.1-q5_1", expected: "Mixtral 8x7B Instruct v0.1 Q5 1"},
		{input: "mixtral:8x7b-instruct-v0.1-q5_K_S", expected: "Mixtral 8x7B Instruct v0.1 Q5 K_S"},
		{input: "mixtral:8x7b-instruct-v0.1-q5_K_M", expected: "Mixtral 8x7B Instruct v0.1 Q5 K_M"},
		{input: "mixtral:8x7b-instruct-v0.1-q6_K", expected: "Mixtral 8x7B Instruct v0.1 Q6 K"},
		{input: "mixtral:8x7b-instruct-v0.1-q8_0", expected: "Mixtral 8x7B Instruct v0.1 Q8_0"},
		{input: "mixtral:8x7b-instruct-v0.1-fp16", expected: "Mixtral 8x7B Instruct v0.1 FP16"},
		{input: "mixtral:8x7b-text", expected: "Mixtral 8x7B Text"},
		{input: "mixtral:8x7b-text-v0.1-q2_K", expected: "Mixtral 8x7B Text v0.1 Q2_K"},
		{input: "mixtral:8x7b-text-v0.1-q3_K_S", expected: "Mixtral 8x7B Text v0.1 Q3 K_S"},
		{input: "mixtral:8x7b-text-v0.1-q3_K_M", expected: "Mixtral 8x7B Text v0.1 Q3 K_M"},
		{input: "mixtral:8x7b-text-v0.1-q3_K_L", expected: "Mixtral 8x7B Text v0.1 Q3 K_L"},
		{input: "mixtral:8x7b-text-v0.1-q4_0", expected: "Mixtral 8x7B Text v0.1 Q4_0"},
		{input: "mixtral:8x7b-text-v0.1-q4_1", expected: "Mixtral 8x7B Text v0.1 Q4 1"},
		{input: "mixtral:8x7b-text-v0.1-q4_K_S", expected: "Mixtral 8x7B Text v0.1 Q4 K_S"},
		{input: "mixtral:8x7b-text-v0.1-q4_K_M", expected: "Mixtral 8x7B Text v0.1 Q4_K_M"},
		{input: "mixtral:8x7b-text-v0.1-q5_0", expected: "Mixtral 8x7B Text v0.1 Q5 0"},
		{input: "mixtral:8x7b-text-v0.1-q5_1", expected: "Mixtral 8x7B Text v0.1 Q5 1"},
		{input: "mixtral:8x7b-text-v0.1-q5_K_S", expected: "Mixtral 8x7B Text v0.1 Q5 K_S"},
		{input: "mixtral:8x7b-text-v0.1-q5_K_M", expected: "Mixtral 8x7B Text v0.1 Q5 K_M"},
		{input: "mixtral:8x7b-text-v0.1-q6_K", expected: "Mixtral 8x7B Text v0.1 Q6 K"},
		{input: "mixtral:8x7b-text-v0.1-q8_0", expected: "Mixtral 8x7B Text v0.1 Q8_0"},
		{input: "mixtral:8x7b-text-v0.1-fp16", expected: "Mixtral 8x7B Text v0.1 FP16"},
		{input: "mixtral:8x22b-instruct", expected: "Mixtral 8x22B Instruct"},
		{input: "mixtral:8x22b-instruct-v0.1-q2_K", expected: "Mixtral 8x22B Instruct v0.1 Q2_K"},
		{input: "mixtral:8x22b-instruct-v0.1-q3_K_S", expected: "Mixtral 8x22B Instruct v0.1 Q3 K_S"},
		{input: "mixtral:8x22b-instruct-v0.1-q3_K_M", expected: "Mixtral 8x22B Instruct v0.1 Q3 K_M"},
		{input: "mixtral:8x22b-instruct-v0.1-q3_K_L", expected: "Mixtral 8x22B Instruct v0.1 Q3 K_L"},
		{input: "mixtral:8x22b-instruct-v0.1-q4_0", expected: "Mixtral 8x22B Instruct v0.1 Q4_0"},
		{input: "mixtral:8x22b-instruct-v0.1-q4_1", expected: "Mixtral 8x22B Instruct v0.1 Q4 1"},
		{input: "mixtral:8x22b-instruct-v0.1-q4_K_S", expected: "Mixtral 8x22B Instruct v0.1 Q4 K_S"},
		{input: "mixtral:8x22b-instruct-v0.1-q4_K_M", expected: "Mixtral 8x22B Instruct v0.1 Q4_K_M"},
		{input: "mixtral:8x22b-instruct-v0.1-q5_0", expected: "Mixtral 8x22B Instruct v0.1 Q5 0"},
		{input: "mixtral:8x22b-instruct-v0.1-q5_1", expected: "Mixtral 8x22B Instruct v0.1 Q5 1"},
		{input: "mixtral:8x22b-instruct-v0.1-q5_K_S", expected: "Mixtral 8x22B Instruct v0.1 Q5 K_S"},
		{input: "mixtral:8x22b-instruct-v0.1-q5_K_M", expected: "Mixtral 8x22B Instruct v0.1 Q5 K_M"},
		{input: "mixtral:8x22b-instruct-v0.1-q6_K", expected: "Mixtral 8x22B Instruct v0.1 Q6 K"},
		{input: "mixtral:8x22b-instruct-v0.1-q8_0", expected: "Mixtral 8x22B Instruct v0.1 Q8_0"},
		{input: "mixtral:8x22b-instruct-v0.1-fp16", expected: "Mixtral 8x22B Instruct v0.1 FP16"},
		{input: "mixtral:8x22b-text", expected: "Mixtral 8x22B Text"},
		{input: "mixtral:8x22b-text-v0.1-q2_K", expected: "Mixtral 8x22B Text v0.1 Q2_K"},
		{input: "mixtral:8x22b-text-v0.1-q3_K_S", expected: "Mixtral 8x22B Text v0.1 Q3 K_S"},
		{input: "mixtral:8x22b-text-v0.1-q3_K_M", expected: "Mixtral 8x22B Text v0.1 Q3 K_M"},
		{input: "mixtral:8x22b-text-v0.1-q3_K_L", expected: "Mixtral 8x22B Text v0.1 Q3 K_L"},
		{input: "mixtral:8x22b-text-v0.1-q4_0", expected: "Mixtral 8x22B Text v0.1 Q4_0"},
		{input: "mixtral:8x22b-text-v0.1-q4_1", expected: "Mixtral 8x22B Text v0.1 Q4 1"},
		{input: "mixtral:8x22b-text-v0.1-q4_K_S", expected: "Mixtral 8x22B Text v0.1 Q4 K_S"},
		{input: "mixtral:8x22b-text-v0.1-q4_K_M", expected: "Mixtral 8x22B Text v0.1 Q4_K_M"},
		{input: "mixtral:8x22b-text-v0.1-q5_0", expected: "Mixtral 8x22B Text v0.1 Q5 0"},
		{input: "mixtral:8x22b-text-v0.1-q5_1", expected: "Mixtral 8x22B Text v0.1 Q5 1"},
		{input: "mixtral:8x22b-text-v0.1-q5_K_S", expected: "Mixtral 8x22B Text v0.1 Q5 K_S"},
		{input: "mixtral:8x22b-text-v0.1-q5_K_M", expected: "Mixtral 8x22B Text v0.1 Q5 K_M"},
		{input: "mixtral:8x22b-text-v0.1-q6_K", expected: "Mixtral 8x22B Text v0.1 Q6 K"},
		{input: "mixtral:8x22b-text-v0.1-q8_0", expected: "Mixtral 8x22B Text v0.1 Q8_0"},
		{input: "mixtral:8x22b-text-v0.1-fp16", expected: "Mixtral 8x22B Text v0.1 FP16"},
		{input: "mixtral:v0.1-instruct", expected: "Mixtral v0.1 Instruct"},
		{input: "llama2-uncensored:latest", expected: "Llama2 Uncensored (latest)"},
		{input: "llama2-uncensored", expected: "Llama2 Uncensored"},
		{input: "llama2-uncensored:7b", expected: "Llama2 Uncensored 7B"},
		{input: "llama2-uncensored:70b", expected: "Llama2 Uncensored 70B"},
		{input: "llama2-uncensored:7b-chat", expected: "Llama2 Uncensored 7B Chat"},
		{input: "llama2-uncensored:7b-chat-q2_K", expected: "Llama2 Uncensored 7B Chat Q2_K"},
		{input: "llama2-uncensored:7b-chat-q3_K_S", expected: "Llama2 Uncensored 7B Chat Q3 K_S"},
		{input: "llama2-uncensored:7b-chat-q3_K_M", expected: "Llama2 Uncensored 7B Chat Q3 K_M"},
		{input: "llama2-uncensored:7b-chat-q3_K_L", expected: "Llama2 Uncensored 7B Chat Q3 K_L"},
		{input: "llama2-uncensored:7b-chat-q4_0", expected: "Llama2 Uncensored 7B Chat Q4_0"},
		{input: "llama2-uncensored:7b-chat-q4_1", expected: "Llama2 Uncensored 7B Chat Q4 1"},
		{input: "llama2-uncensored:7b-chat-q4_K_S", expected: "Llama2 Uncensored 7B Chat Q4 K_S"},
		{input: "llama2-uncensored:7b-chat-q4_K_M", expected: "Llama2 Uncensored 7B Chat Q4_K_M"},
		{input: "llama2-uncensored:7b-chat-q5_0", expected: "Llama2 Uncensored 7B Chat Q5 0"},
		{input: "llama2-uncensored:7b-chat-q5_1", expected: "Llama2 Uncensored 7B Chat Q5 1"},
		{input: "llama2-uncensored:7b-chat-q5_K_S", expected: "Llama2 Uncensored 7B Chat Q5 K_S"},
		{input: "llama2-uncensored:7b-chat-q5_K_M", expected: "Llama2 Uncensored 7B Chat Q5 K_M"},
		{input: "llama2-uncensored:7b-chat-q6_K", expected: "Llama2 Uncensored 7B Chat Q6 K"},
		{input: "llama2-uncensored:7b-chat-q8_0", expected: "Llama2 Uncensored 7B Chat Q8_0"},
		{input: "llama2-uncensored:7b-chat-fp16", expected: "Llama2 Uncensored 7B Chat FP16"},
		{input: "llama2-uncensored:70b-chat", expected: "Llama2 Uncensored 70B Chat"},
		{input: "llama2-uncensored:70b-chat-q2_K", expected: "Llama2 Uncensored 70B Chat Q2_K"},
		{input: "llama2-uncensored:70b-chat-q3_K_S", expected: "Llama2 Uncensored 70B Chat Q3 K_S"},
		{input: "llama2-uncensored:70b-chat-q3_K_M", expected: "Llama2 Uncensored 70B Chat Q3 K_M"},
		{input: "llama2-uncensored:70b-chat-q3_K_L", expected: "Llama2 Uncensored 70B Chat Q3 K_L"},
		{input: "llama2-uncensored:70b-chat-q4_0", expected: "Llama2 Uncensored 70B Chat Q4_0"},
		{input: "llama2-uncensored:70b-chat-q4_1", expected: "Llama2 Uncensored 70B Chat Q4 1"},
		{input: "llama2-uncensored:70b-chat-q4_K_S", expected: "Llama2 Uncensored 70B Chat Q4 K_S"},
		{input: "llama2-uncensored:70b-chat-q4_K_M", expected: "Llama2 Uncensored 70B Chat Q4_K_M"},
		{input: "llama2-uncensored:70b-chat-q5_0", expected: "Llama2 Uncensored 70B Chat Q5 0"},
		{input: "llama2-uncensored:70b-chat-q5_1", expected: "Llama2 Uncensored 70B Chat Q5 1"},
		{input: "llama2-uncensored:70b-chat-q5_K_S", expected: "Llama2 Uncensored 70B Chat Q5 K_S"},
		{input: "llama2-uncensored:70b-chat-q5_K_M", expected: "Llama2 Uncensored 70B Chat Q5 K_M"},
		{input: "llama2-uncensored:70b-chat-q6_K", expected: "Llama2 Uncensored 70B Chat Q6 K"},
		{input: "llama2-uncensored:70b-chat-q8_0", expected: "Llama2 Uncensored 70B Chat Q8_0"},
		{input: "orca-mini:latest", expected: "Orca Mini (latest)"},
		{input: "orca-mini", expected: "Orca Mini"},
		{input: "orca-mini:3b", expected: "Orca Mini 3B"},
		{input: "orca-mini:7b", expected: "Orca Mini 7B"},
		{input: "orca-mini:13b", expected: "Orca Mini 13B"},
		{input: "orca-mini:70b", expected: "Orca Mini 70B"},
		{input: "orca-mini:3b-q4_0", expected: "Orca Mini 3B Q4_0"},
		{input: "orca-mini:3b-q4_1", expected: "Orca Mini 3B Q4 1"},
		{input: "orca-mini:3b-q5_0", expected: "Orca Mini 3B Q5 0"},
		{input: "orca-mini:3b-q5_1", expected: "Orca Mini 3B Q5 1"},
		{input: "orca-mini:3b-q8_0", expected: "Orca Mini 3B Q8_0"},
		{input: "orca-mini:3b-fp16", expected: "Orca Mini 3B FP16"},
		{input: "orca-mini:7b-v2-q2_K", expected: "Orca Mini 7B v2 Q2_K"},
		{input: "orca-mini:7b-v2-q3_K_S", expected: "Orca Mini 7B v2 Q3 K_S"},
		{input: "orca-mini:7b-v2-q3_K_M", expected: "Orca Mini 7B v2 Q3 K_M"},
		{input: "orca-mini:7b-v2-q3_K_L", expected: "Orca Mini 7B v2 Q3 K_L"},
		{input: "orca-mini:7b-v2-q4_0", expected: "Orca Mini 7B v2 Q4_0"},
		{input: "orca-mini:7b-v2-q4_1", expected: "Orca Mini 7B v2 Q4 1"},
		{input: "orca-mini:7b-v2-q4_K_S", expected: "Orca Mini 7B v2 Q4 K_S"},
		{input: "orca-mini:7b-v2-q4_K_M", expected: "Orca Mini 7B v2 Q4_K_M"},
		{input: "orca-mini:7b-v2-q5_0", expected: "Orca Mini 7B v2 Q5 0"},
		{input: "orca-mini:7b-v2-q5_1", expected: "Orca Mini 7B v2 Q5 1"},
		{input: "orca-mini:7b-v2-q5_K_S", expected: "Orca Mini 7B v2 Q5 K_S"},
		{input: "orca-mini:7b-v2-q5_K_M", expected: "Orca Mini 7B v2 Q5 K_M"},
		{input: "orca-mini:7b-v2-q6_K", expected: "Orca Mini 7B v2 Q6 K"},
		{input: "orca-mini:7b-v2-q8_0", expected: "Orca Mini 7B v2 Q8_0"},
		{input: "orca-mini:7b-v2-fp16", expected: "Orca Mini 7B v2 FP16"},
		{input: "orca-mini:7b-v3", expected: "Orca Mini 7B v3"},
		{input: "orca-mini:7b-v3-q2_K", expected: "Orca Mini 7B v3 Q2_K"},
		{input: "orca-mini:7b-v3-q3_K_S", expected: "Orca Mini 7B v3 Q3 K_S"},
		{input: "orca-mini:7b-v3-q3_K_M", expected: "Orca Mini 7B v3 Q3 K_M"},
		{input: "orca-mini:7b-v3-q3_K_L", expected: "Orca Mini 7B v3 Q3 K_L"},
		{input: "orca-mini:7b-v3-q4_0", expected: "Orca Mini 7B v3 Q4_0"},
		{input: "orca-mini:7b-v3-q4_1", expected: "Orca Mini 7B v3 Q4 1"},
		{input: "orca-mini:7b-v3-q4_K_S", expected: "Orca Mini 7B v3 Q4 K_S"},
		{input: "orca-mini:7b-v3-q4_K_M", expected: "Orca Mini 7B v3 Q4_K_M"},
		{input: "orca-mini:7b-v3-q5_0", expected: "Orca Mini 7B v3 Q5 0"},
		{input: "orca-mini:7b-v3-q5_1", expected: "Orca Mini 7B v3 Q5 1"},
		{input: "orca-mini:7b-v3-q5_K_S", expected: "Orca Mini 7B v3 Q5 K_S"},
		{input: "orca-mini:7b-v3-q5_K_M", expected: "Orca Mini 7B v3 Q5 K_M"},
		{input: "orca-mini:7b-v3-q6_K", expected: "Orca Mini 7B v3 Q6 K"},
		{input: "orca-mini:7b-v3-q8_0", expected: "Orca Mini 7B v3 Q8_0"},
		{input: "orca-mini:7b-v3-fp16", expected: "Orca Mini 7B v3 FP16"},
		{input: "orca-mini:7b-q2_K", expected: "Orca Mini 7B Q2_K"},
		{input: "orca-mini:7b-q3_K_S", expected: "Orca Mini 7B Q3 K_S"},
		{input: "orca-mini:7b-q3_K_M", expected: "Orca Mini 7B Q3 K_M"},
		{input: "orca-mini:7b-q3_K_L", expected: "Orca Mini 7B Q3 K_L"},
		{input: "orca-mini:7b-q4_0", expected: "Orca Mini 7B Q4_0"},
		{input: "orca-mini:7b-q4_1", expected: "Orca Mini 7B Q4 1"},
		{input: "orca-mini:7b-q4_K_S", expected: "Orca Mini 7B Q4 K_S"},
		{input: "orca-mini:7b-q4_K_M", expected: "Orca Mini 7B Q4_K_M"},
		{input: "orca-mini:7b-q5_0", expected: "Orca Mini 7B Q5 0"},
		{input: "orca-mini:7b-q5_1", expected: "Orca Mini 7B Q5 1"},
		{input: "orca-mini:7b-q5_K_S", expected: "Orca Mini 7B Q5 K_S"},
		{input: "orca-mini:7b-q5_K_M", expected: "Orca Mini 7B Q5 K_M"},
		{input: "orca-mini:7b-q6_K", expected: "Orca Mini 7B Q6 K"},
		{input: "orca-mini:7b-q8_0", expected: "Orca Mini 7B Q8_0"},
		{input: "orca-mini:7b-fp16", expected: "Orca Mini 7B FP16"},
		{input: "orca-mini:13b-v2-q2_K", expected: "Orca Mini 13B v2 Q2_K"},
		{input: "orca-mini:13b-v2-q3_K_S", expected: "Orca Mini 13B v2 Q3 K_S"},
		{input: "orca-mini:13b-v2-q3_K_M", expected: "Orca Mini 13B v2 Q3 K_M"},
		{input: "orca-mini:13b-v2-q3_K_L", expected: "Orca Mini 13B v2 Q3 K_L"},
		{input: "orca-mini:13b-v2-q4_0", expected: "Orca Mini 13B v2 Q4_0"},
		{input: "orca-mini:13b-v2-q4_1", expected: "Orca Mini 13B v2 Q4 1"},
		{input: "orca-mini:13b-v2-q4_K_S", expected: "Orca Mini 13B v2 Q4 K_S"},
		{input: "orca-mini:13b-v2-q4_K_M", expected: "Orca Mini 13B v2 Q4_K_M"},
		{input: "orca-mini:13b-v2-q5_0", expected: "Orca Mini 13B v2 Q5 0"},
		{input: "orca-mini:13b-v2-q5_1", expected: "Orca Mini 13B v2 Q5 1"},
		{input: "orca-mini:13b-v2-q5_K_S", expected: "Orca Mini 13B v2 Q5 K_S"},
		{input: "orca-mini:13b-v2-q5_K_M", expected: "Orca Mini 13B v2 Q5 K_M"},
		{input: "orca-mini:13b-v2-q6_K", expected: "Orca Mini 13B v2 Q6 K"},
		{input: "orca-mini:13b-v2-q8_0", expected: "Orca Mini 13B v2 Q8_0"},
		{input: "orca-mini:13b-v2-fp16", expected: "Orca Mini 13B v2 FP16"},
		{input: "orca-mini:13b-v3", expected: "Orca Mini 13B v3"},
		{input: "orca-mini:13b-v3-q2_K", expected: "Orca Mini 13B v3 Q2_K"},
		{input: "orca-mini:13b-v3-q3_K_S", expected: "Orca Mini 13B v3 Q3 K_S"},
		{input: "orca-mini:13b-v3-q3_K_M", expected: "Orca Mini 13B v3 Q3 K_M"},
		{input: "orca-mini:13b-v3-q3_K_L", expected: "Orca Mini 13B v3 Q3 K_L"},
		{input: "orca-mini:13b-v3-q4_0", expected: "Orca Mini 13B v3 Q4_0"},
		{input: "orca-mini:13b-v3-q4_1", expected: "Orca Mini 13B v3 Q4 1"},
		{input: "orca-mini:13b-v3-q4_K_S", expected: "Orca Mini 13B v3 Q4 K_S"},
		{input: "orca-mini:13b-v3-q4_K_M", expected: "Orca Mini 13B v3 Q4_K_M"},
		{input: "orca-mini:13b-v3-q5_0", expected: "Orca Mini 13B v3 Q5 0"},
		{input: "orca-mini:13b-v3-q5_1", expected: "Orca Mini 13B v3 Q5 1"},
		{input: "orca-mini:13b-v3-q5_K_S", expected: "Orca Mini 13B v3 Q5 K_S"},
		{input: "orca-mini:13b-v3-q5_K_M", expected: "Orca Mini 13B v3 Q5 K_M"},
		{input: "orca-mini:13b-v3-q6_K", expected: "Orca Mini 13B v3 Q6 K"},
		{input: "orca-mini:13b-v3-q8_0", expected: "Orca Mini 13B v3 Q8_0"},
		{input: "orca-mini:13b-v3-fp16", expected: "Orca Mini 13B v3 FP16"},
		{input: "orca-mini:13b-q2_K", expected: "Orca Mini 13B Q2_K"},
		{input: "orca-mini:13b-q3_K_S", expected: "Orca Mini 13B Q3 K_S"},
		{input: "orca-mini:13b-q3_K_M", expected: "Orca Mini 13B Q3 K_M"},
		{input: "orca-mini:13b-q3_K_L", expected: "Orca Mini 13B Q3 K_L"},
		{input: "orca-mini:13b-q4_0", expected: "Orca Mini 13B Q4_0"},
		{input: "orca-mini:13b-q4_1", expected: "Orca Mini 13B Q4 1"},
		{input: "orca-mini:13b-q4_K_S", expected: "Orca Mini 13B Q4 K_S"},
		{input: "orca-mini:13b-q4_K_M", expected: "Orca Mini 13B Q4_K_M"},
		{input: "orca-mini:13b-q5_0", expected: "Orca Mini 13B Q5 0"},
		{input: "orca-mini:13b-q5_1", expected: "Orca Mini 13B Q5 1"},
		{input: "orca-mini:13b-q5_K_S", expected: "Orca Mini 13B Q5 K_S"},
		{input: "orca-mini:13b-q5_K_M", expected: "Orca Mini 13B Q5 K_M"},
		{input: "orca-mini:13b-q6_K", expected: "Orca Mini 13B Q6 K"},
		{input: "orca-mini:13b-q8_0", expected: "Orca Mini 13B Q8_0"},
		{input: "orca-mini:13b-fp16", expected: "Orca Mini 13B FP16"},
		{input: "orca-mini:70b-v3", expected: "Orca Mini 70B v3"},
		{input: "orca-mini:70b-v3-q2_K", expected: "Orca Mini 70B v3 Q2_K"},
		{input: "orca-mini:70b-v3-q3_K_S", expected: "Orca Mini 70B v3 Q3 K_S"},
		{input: "orca-mini:70b-v3-q3_K_M", expected: "Orca Mini 70B v3 Q3 K_M"},
		{input: "orca-mini:70b-v3-q3_K_L", expected: "Orca Mini 70B v3 Q3 K_L"},
		{input: "orca-mini:70b-v3-q4_0", expected: "Orca Mini 70B v3 Q4_0"},
		{input: "orca-mini:70b-v3-q4_1", expected: "Orca Mini 70B v3 Q4 1"},
		{input: "orca-mini:70b-v3-q4_K_S", expected: "Orca Mini 70B v3 Q4 K_S"},
		{input: "orca-mini:70b-v3-q4_K_M", expected: "Orca Mini 70B v3 Q4_K_M"},
		{input: "orca-mini:70b-v3-q5_0", expected: "Orca Mini 70B v3 Q5 0"},
		{input: "orca-mini:70b-v3-q5_1", expected: "Orca Mini 70B v3 Q5 1"},
		{input: "orca-mini:70b-v3-q5_K_S", expected: "Orca Mini 70B v3 Q5 K_S"},
		{input: "orca-mini:70b-v3-q5_K_M", expected: "Orca Mini 70B v3 Q5 K_M"},
		{input: "orca-mini:70b-v3-q6_K", expected: "Orca Mini 70B v3 Q6 K"},
		{input: "orca-mini:70b-v3-q8_0", expected: "Orca Mini 70B v3 Q8_0"},
		{input: "orca-mini:70b-v3-fp16", expected: "Orca Mini 70B v3 FP16"},
		{input: "snowflake-arctic-embed:latest", expected: "Snowflake Arctic Embed (latest)"},
		{input: "snowflake-arctic-embed", expected: "Snowflake Arctic Embed"},
		{input: "snowflake-arctic-embed:l", expected: "Snowflake Arctic Embed L"},
		{input: "snowflake-arctic-embed:m", expected: "Snowflake Arctic Embed M"},
		{input: "snowflake-arctic-embed:s", expected: "Snowflake Arctic Embed S"},
		{input: "snowflake-arctic-embed:xs", expected: "Snowflake Arctic Embed Xs"},
		{input: "snowflake-arctic-embed:22m", expected: "Snowflake Arctic Embed 22M"},
		{input: "snowflake-arctic-embed:33m", expected: "Snowflake Arctic Embed 33M"},
		{input: "snowflake-arctic-embed:110m", expected: "Snowflake Arctic Embed 110M"},
		{input: "snowflake-arctic-embed:137m", expected: "Snowflake Arctic Embed 137M"},
		{input: "snowflake-arctic-embed:335m", expected: "Snowflake Arctic Embed 335M"},
		{input: "snowflake-arctic-embed:22m-xs-fp16", expected: "Snowflake Arctic Embed 22M Xs FP16"},
		{input: "snowflake-arctic-embed:33m-s-fp16", expected: "Snowflake Arctic Embed 33M S FP16"},
		{input: "snowflake-arctic-embed:110m-m-fp16", expected: "Snowflake Arctic Embed 110M M FP16"},
		{input: "snowflake-arctic-embed:137m-m-long-fp16", expected: "Snowflake Arctic Embed 137M M Long FP16"},
		{input: "snowflake-arctic-embed:335m-l-fp16", expected: "Snowflake Arctic Embed 335M L FP16"},
		{input: "snowflake-arctic-embed:m-long", expected: "Snowflake Arctic Embed M Long"},
		{input: "deepseek-coder-v2:latest", expected: "Deepseek Coder v2 (latest)"},
		{input: "deepseek-coder-v2", expected: "Deepseek Coder v2"},
		{input: "deepseek-coder-v2:lite", expected: "Deepseek Coder v2 Lite"},
		{input: "deepseek-coder-v2:16b", expected: "Deepseek Coder v2 16B"},
		{input: "deepseek-coder-v2:236b", expected: "Deepseek Coder v2 236B"},
		{input: "deepseek-coder-v2:16b-lite-base-q2_K", expected: "Deepseek Coder v2 16B Lite Base Q2_K"},
		{input: "deepseek-coder-v2:16b-lite-base-q3_K_S", expected: "Deepseek Coder v2 16B Lite Base Q3 K_S"},
		{input: "deepseek-coder-v2:16b-lite-base-q3_K_M", expected: "Deepseek Coder v2 16B Lite Base Q3 K_M"},
		{input: "deepseek-coder-v2:16b-lite-base-q3_K_L", expected: "Deepseek Coder v2 16B Lite Base Q3 K_L"},
		{input: "deepseek-coder-v2:16b-lite-base-q4_0", expected: "Deepseek Coder v2 16B Lite Base Q4_0"},
		{input: "deepseek-coder-v2:16b-lite-base-q4_1", expected: "Deepseek Coder v2 16B Lite Base Q4 1"},
		{input: "deepseek-coder-v2:16b-lite-base-q4_K_S", expected: "Deepseek Coder v2 16B Lite Base Q4 K_S"},
		{input: "deepseek-coder-v2:16b-lite-base-q4_K_M", expected: "Deepseek Coder v2 16B Lite Base Q4_K_M"},
		{input: "deepseek-coder-v2:16b-lite-base-q5_0", expected: "Deepseek Coder v2 16B Lite Base Q5 0"},
		{input: "deepseek-coder-v2:16b-lite-base-q5_1", expected: "Deepseek Coder v2 16B Lite Base Q5 1"},
		{input: "deepseek-coder-v2:16b-lite-base-q5_K_S", expected: "Deepseek Coder v2 16B Lite Base Q5 K_S"},
		{input: "deepseek-coder-v2:16b-lite-base-q5_K_M", expected: "Deepseek Coder v2 16B Lite Base Q5 K_M"},
		{input: "deepseek-coder-v2:16b-lite-base-q6_K", expected: "Deepseek Coder v2 16B Lite Base Q6 K"},
		{input: "deepseek-coder-v2:16b-lite-base-q8_0", expected: "Deepseek Coder v2 16B Lite Base Q8_0"},
		{input: "deepseek-coder-v2:16b-lite-base-fp16", expected: "Deepseek Coder v2 16B Lite Base FP16"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q2_K", expected: "Deepseek Coder v2 16B Lite Instruct Q2_K"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q3_K_S", expected: "Deepseek Coder v2 16B Lite Instruct Q3 K_S"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q3_K_M", expected: "Deepseek Coder v2 16B Lite Instruct Q3 K_M"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q3_K_L", expected: "Deepseek Coder v2 16B Lite Instruct Q3 K_L"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q4_0", expected: "Deepseek Coder v2 16B Lite Instruct Q4_0"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q4_1", expected: "Deepseek Coder v2 16B Lite Instruct Q4 1"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q4_K_S", expected: "Deepseek Coder v2 16B Lite Instruct Q4 K_S"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q4_K_M", expected: "Deepseek Coder v2 16B Lite Instruct Q4_K_M"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q5_0", expected: "Deepseek Coder v2 16B Lite Instruct Q5 0"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q5_1", expected: "Deepseek Coder v2 16B Lite Instruct Q5 1"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q5_K_S", expected: "Deepseek Coder v2 16B Lite Instruct Q5 K_S"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q5_K_M", expected: "Deepseek Coder v2 16B Lite Instruct Q5 K_M"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q6_K", expected: "Deepseek Coder v2 16B Lite Instruct Q6 K"},
		{input: "deepseek-coder-v2:16b-lite-instruct-q8_0", expected: "Deepseek Coder v2 16B Lite Instruct Q8_0"},
		{input: "deepseek-coder-v2:16b-lite-instruct-fp16", expected: "Deepseek Coder v2 16B Lite Instruct FP16"},
		{input: "deepseek-coder-v2:236b-base-q2_K", expected: "Deepseek Coder v2 236B Base Q2_K"},
		{input: "deepseek-coder-v2:236b-base-q3_K_S", expected: "Deepseek Coder v2 236B Base Q3 K_S"},
		{input: "deepseek-coder-v2:236b-base-q3_K_M", expected: "Deepseek Coder v2 236B Base Q3 K_M"},
		{input: "deepseek-coder-v2:236b-base-q3_K_L", expected: "Deepseek Coder v2 236B Base Q3 K_L"},
		{input: "deepseek-coder-v2:236b-base-q4_0", expected: "Deepseek Coder v2 236B Base Q4_0"},
		{input: "deepseek-coder-v2:236b-base-q4_1", expected: "Deepseek Coder v2 236B Base Q4 1"},
		{input: "deepseek-coder-v2:236b-base-q4_K_S", expected: "Deepseek Coder v2 236B Base Q4 K_S"},
		{input: "deepseek-coder-v2:236b-base-q4_K_M", expected: "Deepseek Coder v2 236B Base Q4_K_M"},
		{input: "deepseek-coder-v2:236b-base-q5_0", expected: "Deepseek Coder v2 236B Base Q5 0"},
		{input: "deepseek-coder-v2:236b-base-q5_1", expected: "Deepseek Coder v2 236B Base Q5 1"},
		{input: "deepseek-coder-v2:236b-base-q5_K_S", expected: "Deepseek Coder v2 236B Base Q5 K_S"},
		{input: "deepseek-coder-v2:236b-base-q5_K_M", expected: "Deepseek Coder v2 236B Base Q5 K_M"},
		{input: "deepseek-coder-v2:236b-base-q6_K", expected: "Deepseek Coder v2 236B Base Q6 K"},
		{input: "deepseek-coder-v2:236b-base-q8_0", expected: "Deepseek Coder v2 236B Base Q8_0"},
		{input: "deepseek-coder-v2:236b-base-fp16", expected: "Deepseek Coder v2 236B Base FP16"},
		{input: "deepseek-coder-v2:236b-instruct-q2_K", expected: "Deepseek Coder v2 236B Instruct Q2_K"},
		{input: "deepseek-coder-v2:236b-instruct-q3_K_S", expected: "Deepseek Coder v2 236B Instruct Q3 K_S"},
		{input: "deepseek-coder-v2:236b-instruct-q3_K_M", expected: "Deepseek Coder v2 236B Instruct Q3 K_M"},
		{input: "deepseek-coder-v2:236b-instruct-q3_K_L", expected: "Deepseek Coder v2 236B Instruct Q3 K_L"},
		{input: "deepseek-coder-v2:236b-instruct-q4_0", expected: "Deepseek Coder v2 236B Instruct Q4_0"},
		{input: "deepseek-coder-v2:236b-instruct-q4_1", expected: "Deepseek Coder v2 236B Instruct Q4 1"},
		{input: "deepseek-coder-v2:236b-instruct-q4_K_S", expected: "Deepseek Coder v2 236B Instruct Q4 K_S"},
		{input: "deepseek-coder-v2:236b-instruct-q4_K_M", expected: "Deepseek Coder v2 236B Instruct Q4_K_M"},
		{input: "deepseek-coder-v2:236b-instruct-q5_0", expected: "Deepseek Coder v2 236B Instruct Q5 0"},
		{input: "deepseek-coder-v2:236b-instruct-q5_1", expected: "Deepseek Coder v2 236B Instruct Q5 1"},
		{input: "deepseek-coder-v2:236b-instruct-q5_K_S", expected: "Deepseek Coder v2 236B Instruct Q5 K_S"},
		{input: "deepseek-coder-v2:236b-instruct-q5_K_M", expected: "Deepseek Coder v2 236B Instruct Q5 K_M"},
		{input: "deepseek-coder-v2:236b-instruct-q6_K", expected: "Deepseek Coder v2 236B Instruct Q6 K"},
		{input: "deepseek-coder-v2:236b-instruct-q8_0", expected: "Deepseek Coder v2 236B Instruct Q8_0"},
		{input: "deepseek-coder-v2:236b-instruct-fp16", expected: "Deepseek Coder v2 236B Instruct FP16"},
		{input: "qwen2.5vl:latest", expected: "QWEN2.5VL (latest)"},
		{input: "qwen2.5vl", expected: "QWEN2.5VL"},
		{input: "qwen2.5vl:3b", expected: "QWEN2.5VL 3B"},
		{input: "qwen2.5vl:7b", expected: "QWEN2.5VL 7B"},
		{input: "qwen2.5vl:32b", expected: "QWEN2.5VL 32B"},
		{input: "qwen2.5vl:72b", expected: "QWEN2.5VL 72B"},
		{input: "qwen2.5vl:3b-q4_K_M", expected: "QWEN2.5VL 3B Q4_K_M"},
		{input: "qwen2.5vl:3b-q8_0", expected: "QWEN2.5VL 3B Q8_0"},
		{input: "qwen2.5vl:3b-fp16", expected: "QWEN2.5VL 3B FP16"},
		{input: "qwen2.5vl:7b-q4_K_M", expected: "QWEN2.5VL 7B Q4_K_M"},
		{input: "qwen2.5vl:7b-q8_0", expected: "QWEN2.5VL 7B Q8_0"},
		{input: "qwen2.5vl:7b-fp16", expected: "QWEN2.5VL 7B FP16"},
		{input: "qwen2.5vl:32b-q4_K_M", expected: "QWEN2.5VL 32B Q4_K_M"},
		{input: "qwen2.5vl:32b-q8_0", expected: "QWEN2.5VL 32B Q8_0"},
		{input: "qwen2.5vl:32b-fp16", expected: "QWEN2.5VL 32B FP16"},
		{input: "qwen2.5vl:72b-q4_K_M", expected: "QWEN2.5VL 72B Q4_K_M"},
		{input: "qwen2.5vl:72b-q8_0", expected: "QWEN2.5VL 72B Q8_0"},
		{input: "qwen2.5vl:72b-fp16", expected: "QWEN2.5VL 72B FP16"},
		{input: "cogito:latest", expected: "Cogito (latest)"},
		{input: "cogito", expected: "Cogito"},
		{input: "cogito:3b", expected: "Cogito 3B"},
		{input: "cogito:8b", expected: "Cogito 8B"},
		{input: "cogito:14b", expected: "Cogito 14B"},
		{input: "cogito:32b", expected: "Cogito 32B"},
		{input: "cogito:70b", expected: "Cogito 70B"},
		{input: "cogito:3b-v1-preview-llama-q4_K_M", expected: "Cogito 3B v1 Preview Llama Q4_K_M"},
		{input: "cogito:3b-v1-preview-llama-q8_0", expected: "Cogito 3B v1 Preview Llama Q8_0"},
		{input: "cogito:3b-v1-preview-llama-fp16", expected: "Cogito 3B v1 Preview Llama FP16"},
		{input: "cogito:8b-v1-preview-llama-q4_K_M", expected: "Cogito 8B v1 Preview Llama Q4_K_M"},
		{input: "cogito:8b-v1-preview-llama-q8_0", expected: "Cogito 8B v1 Preview Llama Q8_0"},
		{input: "cogito:14b-v1-preview-qwen-q4_K_M", expected: "Cogito 14B v1 Preview Qwen Q4_K_M"},
		{input: "cogito:14b-v1-preview-qwen-q8_0", expected: "Cogito 14B v1 Preview Qwen Q8_0"},
		{input: "cogito:14b-v1-preview-qwen-fp16", expected: "Cogito 14B v1 Preview Qwen FP16"},
		{input: "cogito:32b-v1-preview-qwen-q4_K_M", expected: "Cogito 32B v1 Preview Qwen Q4_K_M"},
		{input: "cogito:32b-v1-preview-qwen-q8_0", expected: "Cogito 32B v1 Preview Qwen Q8_0"},
		{input: "cogito:32b-v1-preview-qwen-fp16", expected: "Cogito 32B v1 Preview Qwen FP16"},
		{input: "cogito:70b-v1-preview-llama-q4_K_M", expected: "Cogito 70B v1 Preview Llama Q4_K_M"},
		{input: "cogito:70b-v1-preview-llama-q8_0", expected: "Cogito 70B v1 Preview Llama Q8_0"},
		{input: "cogito:70b-v1-preview-llama-fp16", expected: "Cogito 70B v1 Preview Llama FP16"},
		{input: "gemma3n:latest", expected: "Gemma3n (latest)"},
		{input: "gemma3n", expected: "Gemma3n"},
		{input: "gemma3n:e2b", expected: "Gemma3n E2B"},
		{input: "gemma3n:e4b", expected: "Gemma3n E4B"},
		{input: "gemma3n:e2b-it-q4_K_M", expected: "Gemma3n E2B Instruct Q4_K_M"},
		{input: "gemma3n:e2b-it-q8_0", expected: "Gemma3n E2B Instruct Q8_0"},
		{input: "gemma3n:e2b-it-fp16", expected: "Gemma3n E2B Instruct FP16"},
		{input: "gemma3n:e4b-it-q4_K_M", expected: "Gemma3n E4B Instruct Q4_K_M"},
		{input: "gemma3n:e4b-it-q8_0", expected: "Gemma3n E4B Instruct Q8_0"},
		{input: "gemma3n:e4b-it-fp16", expected: "Gemma3n E4B Instruct FP16"},
		{input: "llama4:latest", expected: "Llama4 (latest)"},
		{input: "llama4", expected: "Llama4"},
		{input: "llama4:maverick", expected: "Llama4 Maverick"},
		{input: "llama4:scout", expected: "Llama4 Scout"},
		{input: "llama4:16x17b", expected: "Llama4 16x17B"},
		{input: "llama4:128x17b", expected: "Llama4 128x17B"},
		{input: "llama4:17b-maverick-128e-instruct-q4_K_M", expected: "Llama4 17B Maverick 128E Instruct Q4_K_M"},
		{input: "llama4:17b-maverick-128e-instruct-q8_0", expected: "Llama4 17B Maverick 128E Instruct Q8_0"},
		{input: "llama4:17b-maverick-128e-instruct-fp16", expected: "Llama4 17B Maverick 128E Instruct FP16"},
		{input: "llama4:17b-scout-16e-instruct-q4_K_M", expected: "Llama4 17B Scout 16E Instruct Q4_K_M"},
		{input: "llama4:17b-scout-16e-instruct-q8_0", expected: "Llama4 17B Scout 16E Instruct Q8_0"},
		{input: "llama4:17b-scout-16e-instruct-fp16", expected: "Llama4 17B Scout 16E Instruct FP16"},
		{input: "mistral-small3.2:latest", expected: "Mistral Small3.2 (latest)"},
		{input: "mistral-small3.2", expected: "Mistral Small3.2"},
		{input: "mistral-small3.2:24b", expected: "Mistral Small3.2 24B"},
		{input: "mistral-small3.2:24b-instruct-2506-q4_K_M", expected: "Mistral Small3.2 24B Instruct 2506 Q4_K_M"},
		{input: "mistral-small3.2:24b-instruct-2506-q8_0", expected: "Mistral Small3.2 24B Instruct 2506 Q8_0"},
		{input: "mistral-small3.2:24b-instruct-2506-fp16", expected: "Mistral Small3.2 24B Instruct 2506 FP16"},
		{input: "deepscaler:latest", expected: "Deepscaler (latest)"},
		{input: "deepscaler", expected: "Deepscaler"},
		{input: "deepscaler:1.5b", expected: "Deepscaler 1.5B"},
		{input: "deepscaler:1.5b-preview-q4_K_M", expected: "Deepscaler 1.5B Preview Q4_K_M"},
		{input: "deepscaler:1.5b-preview-q8_0", expected: "Deepscaler 1.5B Preview Q8_0"},
		{input: "deepscaler:1.5b-preview-fp16", expected: "Deepscaler 1.5B Preview FP16"},
		{input: "phi4-reasoning:latest", expected: "Phi4 Reasoning (latest)"},
		{input: "phi4-reasoning", expected: "Phi4 Reasoning"},
		{input: "phi4-reasoning:plus", expected: "Phi4 Reasoning Plus"},
		{input: "phi4-reasoning:14b", expected: "Phi4 Reasoning 14B"},
		{input: "phi4-reasoning:14b-plus-q4_K_M", expected: "Phi4 Reasoning 14B Plus Q4_K_M"},
		{input: "phi4-reasoning:14b-plus-q8_0", expected: "Phi4 Reasoning 14B Plus Q8_0"},
		{input: "phi4-reasoning:14b-plus-fp16", expected: "Phi4 Reasoning 14B Plus FP16"},
		{input: "phi4-reasoning:14b-q4_K_M", expected: "Phi4 Reasoning 14B Q4_K_M"},
		{input: "phi4-reasoning:14b-q8_0", expected: "Phi4 Reasoning 14B Q8_0"},
		{input: "phi4-reasoning:14b-fp16", expected: "Phi4 Reasoning 14B FP16"},
		{input: "dolphin-phi:latest", expected: "Dolphin Phi (latest)"},
		{input: "dolphin-phi", expected: "Dolphin Phi"},
		{input: "dolphin-phi:2.7b", expected: "Dolphin Phi 2.7B"},
		{input: "dolphin-phi:2.7b-v2.6", expected: "Dolphin Phi 2.7B v2.6"},
		{input: "dolphin-phi:2.7b-v2.6-q2_K", expected: "Dolphin Phi 2.7B v2.6 Q2_K"},
		{input: "dolphin-phi:2.7b-v2.6-q3_K_S", expected: "Dolphin Phi 2.7B v2.6 Q3 K_S"},
		{input: "dolphin-phi:2.7b-v2.6-q3_K_M", expected: "Dolphin Phi 2.7B v2.6 Q3 K_M"},
		{input: "dolphin-phi:2.7b-v2.6-q3_K_L", expected: "Dolphin Phi 2.7B v2.6 Q3 K_L"},
		{input: "dolphin-phi:2.7b-v2.6-q4_0", expected: "Dolphin Phi 2.7B v2.6 Q4_0"},
		{input: "dolphin-phi:2.7b-v2.6-q4_K_S", expected: "Dolphin Phi 2.7B v2.6 Q4 K_S"},
		{input: "dolphin-phi:2.7b-v2.6-q4_K_M", expected: "Dolphin Phi 2.7B v2.6 Q4_K_M"},
		{input: "dolphin-phi:2.7b-v2.6-q5_0", expected: "Dolphin Phi 2.7B v2.6 Q5 0"},
		{input: "dolphin-phi:2.7b-v2.6-q5_K_S", expected: "Dolphin Phi 2.7B v2.6 Q5 K_S"},
		{input: "dolphin-phi:2.7b-v2.6-q5_K_M", expected: "Dolphin Phi 2.7B v2.6 Q5 K_M"},
		{input: "dolphin-phi:2.7b-v2.6-q6_K", expected: "Dolphin Phi 2.7B v2.6 Q6 K"},
		{input: "dolphin-phi:2.7b-v2.6-q8_0", expected: "Dolphin Phi 2.7B v2.6 Q8_0"},
		{input: "magistral:latest", expected: "Magistral (latest)"},
		{input: "magistral", expected: "Magistral"},
		{input: "magistral:24b", expected: "Magistral 24B"},
		{input: "magistral:24b-small-2506-q4_K_M", expected: "Magistral 24B Small 2506 Q4_K_M"},
		{input: "magistral:24b-small-2506-q8_0", expected: "Magistral 24B Small 2506 Q8_0"},
		{input: "magistral:24b-small-2506-fp16", expected: "Magistral 24B Small 2506 FP16"},
		{input: "phi:latest", expected: "Phi (latest)"},
		{input: "phi", expected: "Phi"},
		{input: "phi:chat", expected: "Phi Chat"},
		{input: "phi:2.7b", expected: "Phi 2.7B"},
		{input: "phi:2.7b-chat-v2-q2_K", expected: "Phi 2.7B Chat v2 Q2_K"},
		{input: "phi:2.7b-chat-v2-q3_K_S", expected: "Phi 2.7B Chat v2 Q3 K_S"},
		{input: "phi:2.7b-chat-v2-q3_K_M", expected: "Phi 2.7B Chat v2 Q3 K_M"},
		{input: "phi:2.7b-chat-v2-q3_K_L", expected: "Phi 2.7B Chat v2 Q3 K_L"},
		{input: "phi:2.7b-chat-v2-q4_0", expected: "Phi 2.7B Chat v2 Q4_0"},
		{input: "phi:2.7b-chat-v2-q4_1", expected: "Phi 2.7B Chat v2 Q4 1"},
		{input: "phi:2.7b-chat-v2-q4_K_S", expected: "Phi 2.7B Chat v2 Q4 K_S"},
		{input: "phi:2.7b-chat-v2-q4_K_M", expected: "Phi 2.7B Chat v2 Q4_K_M"},
		{input: "phi:2.7b-chat-v2-q5_0", expected: "Phi 2.7B Chat v2 Q5 0"},
		{input: "phi:2.7b-chat-v2-q5_1", expected: "Phi 2.7B Chat v2 Q5 1"},
		{input: "phi:2.7b-chat-v2-q5_K_S", expected: "Phi 2.7B Chat v2 Q5 K_S"},
		{input: "phi:2.7b-chat-v2-q5_K_M", expected: "Phi 2.7B Chat v2 Q5 K_M"},
		{input: "phi:2.7b-chat-v2-q6_K", expected: "Phi 2.7B Chat v2 Q6 K"},
		{input: "phi:2.7b-chat-v2-q8_0", expected: "Phi 2.7B Chat v2 Q8_0"},
		{input: "phi:2.7b-chat-v2-fp16", expected: "Phi 2.7B Chat v2 FP16"},
		{input: "granite3.3:latest", expected: "Granite3.3 (latest)"},
		{input: "granite3.3", expected: "Granite3.3"},
		{input: "granite3.3:2b", expected: "Granite3.3 2B"},
		{input: "granite3.3:8b", expected: "Granite3.3 8B"},
		{input: "dolphin-mixtral:latest", expected: "Dolphin Mixtral (latest)"},
		{input: "dolphin-mixtral", expected: "Dolphin Mixtral"},
		{input: "dolphin-mixtral:v2.5", expected: "Dolphin Mixtral v2.5"},
		{input: "dolphin-mixtral:v2.6", expected: "Dolphin Mixtral v2.6"},
		{input: "dolphin-mixtral:v2.7", expected: "Dolphin Mixtral v2.7"},
		{input: "dolphin-mixtral:8x7b", expected: "Dolphin Mixtral 8x7B"},
		{input: "dolphin-mixtral:8x22b", expected: "Dolphin Mixtral 8x22B"},
		{input: "dolphin-mixtral:8x7b-v2.5", expected: "Dolphin Mixtral 8x7B v2.5"},
		{input: "dolphin-mixtral:8x7b-v2.5-q2_K", expected: "Dolphin Mixtral 8x7B v2.5 Q2_K"},
		{input: "dolphin-mixtral:8x7b-v2.5-q3_K_S", expected: "Dolphin Mixtral 8x7B v2.5 Q3 K_S"},
		{input: "dolphin-mixtral:8x7b-v2.5-q3_K_M", expected: "Dolphin Mixtral 8x7B v2.5 Q3 K_M"},
		{input: "dolphin-mixtral:8x7b-v2.5-q3_K_L", expected: "Dolphin Mixtral 8x7B v2.5 Q3 K_L"},
		{input: "dolphin-mixtral:8x7b-v2.5-q4_0", expected: "Dolphin Mixtral 8x7B v2.5 Q4_0"},
		{input: "dolphin-mixtral:8x7b-v2.5-q4_1", expected: "Dolphin Mixtral 8x7B v2.5 Q4 1"},
		{input: "dolphin-mixtral:8x7b-v2.5-q4_K_S", expected: "Dolphin Mixtral 8x7B v2.5 Q4 K_S"},
		{input: "dolphin-mixtral:8x7b-v2.5-q4_K_M", expected: "Dolphin Mixtral 8x7B v2.5 Q4_K_M"},
		{input: "dolphin-mixtral:8x7b-v2.5-q5_0", expected: "Dolphin Mixtral 8x7B v2.5 Q5 0"},
		{input: "dolphin-mixtral:8x7b-v2.5-q5_1", expected: "Dolphin Mixtral 8x7B v2.5 Q5 1"},
		{input: "dolphin-mixtral:8x7b-v2.5-q5_K_S", expected: "Dolphin Mixtral 8x7B v2.5 Q5 K_S"},
		{input: "dolphin-mixtral:8x7b-v2.5-q5_K_M", expected: "Dolphin Mixtral 8x7B v2.5 Q5 K_M"},
		{input: "dolphin-mixtral:8x7b-v2.5-q6_K", expected: "Dolphin Mixtral 8x7B v2.5 Q6 K"},
		{input: "dolphin-mixtral:8x7b-v2.5-q8_0", expected: "Dolphin Mixtral 8x7B v2.5 Q8_0"},
		{input: "dolphin-mixtral:8x7b-v2.5-fp16", expected: "Dolphin Mixtral 8x7B v2.5 FP16"},
		{input: "dolphin-mixtral:8x7b-v2.6", expected: "Dolphin Mixtral 8x7B v2.6"},
		{input: "dolphin-mixtral:8x7b-v2.6-q2_K", expected: "Dolphin Mixtral 8x7B v2.6 Q2_K"},
		{input: "dolphin-mixtral:8x7b-v2.6-q3_K_S", expected: "Dolphin Mixtral 8x7B v2.6 Q3 K_S"},
		{input: "dolphin-mixtral:8x7b-v2.6-q3_K_M", expected: "Dolphin Mixtral 8x7B v2.6 Q3 K_M"},
		{input: "dolphin-mixtral:8x7b-v2.6-q3_K_L", expected: "Dolphin Mixtral 8x7B v2.6 Q3 K_L"},
		{input: "dolphin-mixtral:8x7b-v2.6-q4_0", expected: "Dolphin Mixtral 8x7B v2.6 Q4_0"},
		{input: "dolphin-mixtral:8x7b-v2.6-q4_1", expected: "Dolphin Mixtral 8x7B v2.6 Q4 1"},
		{input: "dolphin-mixtral:8x7b-v2.6-q4_K_S", expected: "Dolphin Mixtral 8x7B v2.6 Q4 K_S"},
		{input: "dolphin-mixtral:8x7b-v2.6-q4_K_M", expected: "Dolphin Mixtral 8x7B v2.6 Q4_K_M"},
		{input: "dolphin-mixtral:8x7b-v2.6-q5_0", expected: "Dolphin Mixtral 8x7B v2.6 Q5 0"},
		{input: "dolphin-mixtral:8x7b-v2.6-q5_1", expected: "Dolphin Mixtral 8x7B v2.6 Q5 1"},
		{input: "dolphin-mixtral:8x7b-v2.6-q5_K_S", expected: "Dolphin Mixtral 8x7B v2.6 Q5 K_S"},
		{input: "dolphin-mixtral:8x7b-v2.6-q5_K_M", expected: "Dolphin Mixtral 8x7B v2.6 Q5 K_M"},
		{input: "dolphin-mixtral:8x7b-v2.6-q6_K", expected: "Dolphin Mixtral 8x7B v2.6 Q6 K"},
		{input: "dolphin-mixtral:8x7b-v2.6-q8_0", expected: "Dolphin Mixtral 8x7B v2.6 Q8_0"},
		{input: "dolphin-mixtral:8x7b-v2.6-fp16", expected: "Dolphin Mixtral 8x7B v2.6 FP16"},
		{input: "dolphin-mixtral:8x7b-v2.7", expected: "Dolphin Mixtral 8x7B v2.7"},
		{input: "dolphin-mixtral:8x7b-v2.7-q2_K", expected: "Dolphin Mixtral 8x7B v2.7 Q2_K"},
		{input: "dolphin-mixtral:8x7b-v2.7-q3_K_S", expected: "Dolphin Mixtral 8x7B v2.7 Q3 K_S"},
		{input: "dolphin-mixtral:8x7b-v2.7-q3_K_M", expected: "Dolphin Mixtral 8x7B v2.7 Q3 K_M"},
		{input: "dolphin-mixtral:8x7b-v2.7-q3_K_L", expected: "Dolphin Mixtral 8x7B v2.7 Q3 K_L"},
		{input: "dolphin-mixtral:8x7b-v2.7-q4_0", expected: "Dolphin Mixtral 8x7B v2.7 Q4_0"},
		{input: "dolphin-mixtral:8x7b-v2.7-q4_1", expected: "Dolphin Mixtral 8x7B v2.7 Q4 1"},
		{input: "dolphin-mixtral:8x7b-v2.7-q4_K_S", expected: "Dolphin Mixtral 8x7B v2.7 Q4 K_S"},
		{input: "dolphin-mixtral:8x7b-v2.7-q4_K_M", expected: "Dolphin Mixtral 8x7B v2.7 Q4_K_M"},
		{input: "dolphin-mixtral:8x7b-v2.7-q5_0", expected: "Dolphin Mixtral 8x7B v2.7 Q5 0"},
		{input: "dolphin-mixtral:8x7b-v2.7-q5_1", expected: "Dolphin Mixtral 8x7B v2.7 Q5 1"},
		{input: "dolphin-mixtral:8x7b-v2.7-q5_K_S", expected: "Dolphin Mixtral 8x7B v2.7 Q5 K_S"},
		{input: "dolphin-mixtral:8x7b-v2.7-q5_K_M", expected: "Dolphin Mixtral 8x7B v2.7 Q5 K_M"},
		{input: "dolphin-mixtral:8x7b-v2.7-q6_K", expected: "Dolphin Mixtral 8x7B v2.7 Q6 K"},
		{input: "dolphin-mixtral:8x7b-v2.7-q8_0", expected: "Dolphin Mixtral 8x7B v2.7 Q8_0"},
		{input: "dolphin-mixtral:8x7b-v2.7-fp16", expected: "Dolphin Mixtral 8x7B v2.7 FP16"},
		{input: "dolphin-mixtral:8x22b-v2.9", expected: "Dolphin Mixtral 8x22B v2.9"},
		{input: "dolphin-mixtral:8x22b-v2.9-q2_K", expected: "Dolphin Mixtral 8x22B v2.9 Q2_K"},
		{input: "dolphin-mixtral:8x22b-v2.9-q3_K_S", expected: "Dolphin Mixtral 8x22B v2.9 Q3 K_S"},
		{input: "dolphin-mixtral:8x22b-v2.9-q3_K_M", expected: "Dolphin Mixtral 8x22B v2.9 Q3 K_M"},
		{input: "dolphin-mixtral:8x22b-v2.9-q3_K_L", expected: "Dolphin Mixtral 8x22B v2.9 Q3 K_L"},
		{input: "dolphin-mixtral:8x22b-v2.9-q4_0", expected: "Dolphin Mixtral 8x22B v2.9 Q4_0"},
		{input: "dolphin-mixtral:8x22b-v2.9-q4_1", expected: "Dolphin Mixtral 8x22B v2.9 Q4 1"},
		{input: "dolphin-mixtral:8x22b-v2.9-q4_K_S", expected: "Dolphin Mixtral 8x22B v2.9 Q4 K_S"},
		{input: "dolphin-mixtral:8x22b-v2.9-q4_K_M", expected: "Dolphin Mixtral 8x22B v2.9 Q4_K_M"},
		{input: "dolphin-mixtral:8x22b-v2.9-q5_0", expected: "Dolphin Mixtral 8x22B v2.9 Q5 0"},
		{input: "dolphin-mixtral:8x22b-v2.9-q5_1", expected: "Dolphin Mixtral 8x22B v2.9 Q5 1"},
		{input: "dolphin-mixtral:8x22b-v2.9-q5_K_S", expected: "Dolphin Mixtral 8x22B v2.9 Q5 K_S"},
		{input: "dolphin-mixtral:8x22b-v2.9-q5_K_M", expected: "Dolphin Mixtral 8x22B v2.9 Q5 K_M"},
		{input: "dolphin-mixtral:8x22b-v2.9-q6_K", expected: "Dolphin Mixtral 8x22B v2.9 Q6 K"},
		{input: "dolphin-mixtral:8x22b-v2.9-q8_0", expected: "Dolphin Mixtral 8x22B v2.9 Q8_0"},
		{input: "dolphin-mixtral:8x22b-v2.9-fp16", expected: "Dolphin Mixtral 8x22B v2.9 FP16"},
		{input: "phi4-mini:latest", expected: "Phi4 Mini (latest)"},
		{input: "phi4-mini", expected: "Phi4 Mini"},
		{input: "phi4-mini:3.8b", expected: "Phi4 Mini 3.8B"},
		{input: "phi4-mini:3.8b-q4_K_M", expected: "Phi4 Mini 3.8B Q4_K_M"},
		{input: "phi4-mini:3.8b-q8_0", expected: "Phi4 Mini 3.8B Q8_0"},
		{input: "phi4-mini:3.8b-fp16", expected: "Phi4 Mini 3.8B FP16"},
		{input: "dolphin-llama3:latest", expected: "Dolphin Llama3 (latest)"},
		{input: "dolphin-llama3", expected: "Dolphin Llama3"},
		{input: "dolphin-llama3:v2.9", expected: "Dolphin Llama3 v2.9"},
		{input: "dolphin-llama3:8b", expected: "Dolphin Llama3 8B"},
		{input: "dolphin-llama3:70b", expected: "Dolphin Llama3 70B"},
		{input: "dolphin-llama3:8b-256k", expected: "Dolphin Llama3 8B 256K"},
		{input: "dolphin-llama3:8b-256k-v2.9", expected: "Dolphin Llama3 8B 256K v2.9"},
		{input: "dolphin-llama3:8b-256k-v2.9-q2_K", expected: "Dolphin Llama3 8B 256K v2.9 Q2_K"},
		{input: "dolphin-llama3:8b-256k-v2.9-q3_K_S", expected: "Dolphin Llama3 8B 256K v2.9 Q3 K_S"},
		{input: "dolphin-llama3:8b-256k-v2.9-q3_K_M", expected: "Dolphin Llama3 8B 256K v2.9 Q3 K_M"},
		{input: "dolphin-llama3:8b-256k-v2.9-q3_K_L", expected: "Dolphin Llama3 8B 256K v2.9 Q3 K_L"},
		{input: "dolphin-llama3:8b-256k-v2.9-q4_0", expected: "Dolphin Llama3 8B 256K v2.9 Q4_0"},
		{input: "dolphin-llama3:8b-256k-v2.9-q4_1", expected: "Dolphin Llama3 8B 256K v2.9 Q4 1"},
		{input: "dolphin-llama3:8b-256k-v2.9-q4_K_S", expected: "Dolphin Llama3 8B 256K v2.9 Q4 K_S"},
		{input: "dolphin-llama3:8b-256k-v2.9-q4_K_M", expected: "Dolphin Llama3 8B 256K v2.9 Q4_K_M"},
		{input: "dolphin-llama3:8b-256k-v2.9-q5_0", expected: "Dolphin Llama3 8B 256K v2.9 Q5 0"},
		{input: "dolphin-llama3:8b-256k-v2.9-q5_1", expected: "Dolphin Llama3 8B 256K v2.9 Q5 1"},
		{input: "dolphin-llama3:8b-256k-v2.9-q5_K_S", expected: "Dolphin Llama3 8B 256K v2.9 Q5 K_S"},
		{input: "dolphin-llama3:8b-256k-v2.9-q5_K_M", expected: "Dolphin Llama3 8B 256K v2.9 Q5 K_M"},
		{input: "dolphin-llama3:8b-256k-v2.9-q6_K", expected: "Dolphin Llama3 8B 256K v2.9 Q6 K"},
		{input: "dolphin-llama3:8b-256k-v2.9-q8_0", expected: "Dolphin Llama3 8B 256K v2.9 Q8_0"},
		{input: "dolphin-llama3:8b-256k-v2.9-fp16", expected: "Dolphin Llama3 8B 256K v2.9 FP16"},
		{input: "dolphin-llama3:8b-v2.9", expected: "Dolphin Llama3 8B v2.9"},
		{input: "dolphin-llama3:8b-v2.9-q2_K", expected: "Dolphin Llama3 8B v2.9 Q2_K"},
		{input: "dolphin-llama3:8b-v2.9-q3_K_S", expected: "Dolphin Llama3 8B v2.9 Q3 K_S"},
		{input: "dolphin-llama3:8b-v2.9-q3_K_M", expected: "Dolphin Llama3 8B v2.9 Q3 K_M"},
		{input: "dolphin-llama3:8b-v2.9-q3_K_L", expected: "Dolphin Llama3 8B v2.9 Q3 K_L"},
		{input: "dolphin-llama3:8b-v2.9-q4_0", expected: "Dolphin Llama3 8B v2.9 Q4_0"},
		{input: "dolphin-llama3:8b-v2.9-q4_1", expected: "Dolphin Llama3 8B v2.9 Q4 1"},
		{input: "dolphin-llama3:8b-v2.9-q4_K_S", expected: "Dolphin Llama3 8B v2.9 Q4 K_S"},
		{input: "dolphin-llama3:8b-v2.9-q4_K_M", expected: "Dolphin Llama3 8B v2.9 Q4_K_M"},
		{input: "dolphin-llama3:8b-v2.9-q5_0", expected: "Dolphin Llama3 8B v2.9 Q5 0"},
		{input: "dolphin-llama3:8b-v2.9-q5_1", expected: "Dolphin Llama3 8B v2.9 Q5 1"},
		{input: "dolphin-llama3:8b-v2.9-q5_K_S", expected: "Dolphin Llama3 8B v2.9 Q5 K_S"},
		{input: "dolphin-llama3:8b-v2.9-q5_K_M", expected: "Dolphin Llama3 8B v2.9 Q5 K_M"},
		{input: "dolphin-llama3:8b-v2.9-q6_K", expected: "Dolphin Llama3 8B v2.9 Q6 K"},
		{input: "dolphin-llama3:8b-v2.9-q8_0", expected: "Dolphin Llama3 8B v2.9 Q8_0"},
		{input: "dolphin-llama3:8b-v2.9-fp16", expected: "Dolphin Llama3 8B v2.9 FP16"},
		{input: "dolphin-llama3:70b-v2.9", expected: "Dolphin Llama3 70B v2.9"},
		{input: "dolphin-llama3:70b-v2.9-q2_K", expected: "Dolphin Llama3 70B v2.9 Q2_K"},
		{input: "dolphin-llama3:70b-v2.9-q3_K_S", expected: "Dolphin Llama3 70B v2.9 Q3 K_S"},
		{input: "dolphin-llama3:70b-v2.9-q3_K_M", expected: "Dolphin Llama3 70B v2.9 Q3 K_M"},
		{input: "dolphin-llama3:70b-v2.9-q3_K_L", expected: "Dolphin Llama3 70B v2.9 Q3 K_L"},
		{input: "dolphin-llama3:70b-v2.9-q4_0", expected: "Dolphin Llama3 70B v2.9 Q4_0"},
		{input: "dolphin-llama3:70b-v2.9-q4_1", expected: "Dolphin Llama3 70B v2.9 Q4 1"},
		{input: "dolphin-llama3:70b-v2.9-q4_K_S", expected: "Dolphin Llama3 70B v2.9 Q4 K_S"},
		{input: "dolphin-llama3:70b-v2.9-q4_K_M", expected: "Dolphin Llama3 70B v2.9 Q4_K_M"},
		{input: "dolphin-llama3:70b-v2.9-q5_0", expected: "Dolphin Llama3 70B v2.9 Q5 0"},
		{input: "dolphin-llama3:70b-v2.9-q5_1", expected: "Dolphin Llama3 70B v2.9 Q5 1"},
		{input: "dolphin-llama3:70b-v2.9-q5_K_S", expected: "Dolphin Llama3 70B v2.9 Q5 K_S"},
		{input: "dolphin-llama3:70b-v2.9-q5_K_M", expected: "Dolphin Llama3 70B v2.9 Q5 K_M"},
		{input: "dolphin-llama3:70b-v2.9-q6_K", expected: "Dolphin Llama3 70B v2.9 Q6 K"},
		{input: "dolphin-llama3:70b-v2.9-q8_0", expected: "Dolphin Llama3 70B v2.9 Q8_0"},
		{input: "dolphin-llama3:70b-v2.9-fp16", expected: "Dolphin Llama3 70B v2.9 FP16"},
		{input: "openthinker:latest", expected: "Openthinker (latest)"},
		{input: "openthinker", expected: "Openthinker"},
		{input: "openthinker:7b", expected: "Openthinker 7B"},
		{input: "openthinker:32b", expected: "Openthinker 32B"},
		{input: "openthinker:7b-v2-q4_K_M", expected: "Openthinker 7B v2 Q4_K_M"},
		{input: "openthinker:7b-v2-q8_0", expected: "Openthinker 7B v2 Q8_0"},
		{input: "openthinker:7b-v2-fp16", expected: "Openthinker 7B v2 FP16"},
		{input: "openthinker:7b-q4_K_M", expected: "Openthinker 7B Q4_K_M"},
		{input: "openthinker:7b-q8_0", expected: "Openthinker 7B Q8_0"},
		{input: "openthinker:7b-fp16", expected: "Openthinker 7B FP16"},
		{input: "openthinker:32b-v2-q4_K_M", expected: "Openthinker 32B v2 Q4_K_M"},
		{input: "openthinker:32b-v2-q8_0", expected: "Openthinker 32B v2 Q8_0"},
		{input: "openthinker:32b-v2-fp16", expected: "Openthinker 32B v2 FP16"},
		{input: "openthinker:32b-q4_K_M", expected: "Openthinker 32B Q4_K_M"},
		{input: "openthinker:32b-q8_0", expected: "Openthinker 32B Q8_0"},
		{input: "openthinker:32b-fp16", expected: "Openthinker 32B FP16"},
		{input: "codestral:latest", expected: "Codestral (latest)"},
		{input: "codestral", expected: "Codestral"},
		{input: "codestral:v0.1", expected: "Codestral v0.1"},
		{input: "codestral:22b", expected: "Codestral 22B"},
		{input: "codestral:22b-v0.1-q2_K", expected: "Codestral 22B v0.1 Q2_K"},
		{input: "codestral:22b-v0.1-q3_K_S", expected: "Codestral 22B v0.1 Q3 K_S"},
		{input: "codestral:22b-v0.1-q3_K_M", expected: "Codestral 22B v0.1 Q3 K_M"},
		{input: "codestral:22b-v0.1-q3_K_L", expected: "Codestral 22B v0.1 Q3 K_L"},
		{input: "codestral:22b-v0.1-q4_0", expected: "Codestral 22B v0.1 Q4_0"},
		{input: "codestral:22b-v0.1-q4_1", expected: "Codestral 22B v0.1 Q4 1"},
		{input: "codestral:22b-v0.1-q4_K_S", expected: "Codestral 22B v0.1 Q4 K_S"},
		{input: "codestral:22b-v0.1-q4_K_M", expected: "Codestral 22B v0.1 Q4_K_M"},
		{input: "codestral:22b-v0.1-q5_0", expected: "Codestral 22B v0.1 Q5 0"},
		{input: "codestral:22b-v0.1-q5_1", expected: "Codestral 22B v0.1 Q5 1"},
		{input: "codestral:22b-v0.1-q5_K_S", expected: "Codestral 22B v0.1 Q5 K_S"},
		{input: "codestral:22b-v0.1-q5_K_M", expected: "Codestral 22B v0.1 Q5 K_M"},
		{input: "codestral:22b-v0.1-q6_K", expected: "Codestral 22B v0.1 Q6 K"},
		{input: "codestral:22b-v0.1-q8_0", expected: "Codestral 22B v0.1 Q8_0"},
		{input: "smollm:latest", expected: "SmolLM (latest)"},
		{input: "smollm", expected: "SmolLM"},
		{input: "smollm:135m", expected: "SmolLM 135M"},
		{input: "smollm:360m", expected: "SmolLM 360M"},
		{input: "smollm:1.7b", expected: "SmolLM 1.7B"},
		{input: "smollm:135m-base-v0.2-q2_K", expected: "SmolLM 135M Base v0.2 Q2_K"},
		{input: "smollm:135m-base-v0.2-q3_K_S", expected: "SmolLM 135M Base v0.2 Q3 K_S"},
		{input: "smollm:135m-base-v0.2-q3_K_M", expected: "SmolLM 135M Base v0.2 Q3 K_M"},
		{input: "smollm:135m-base-v0.2-q3_K_L", expected: "SmolLM 135M Base v0.2 Q3 K_L"},
		{input: "smollm:135m-base-v0.2-q4_0", expected: "SmolLM 135M Base v0.2 Q4_0"},
		{input: "smollm:135m-base-v0.2-q4_1", expected: "SmolLM 135M Base v0.2 Q4 1"},
		{input: "smollm:135m-base-v0.2-q4_K_S", expected: "SmolLM 135M Base v0.2 Q4 K_S"},
		{input: "smollm:135m-base-v0.2-q4_K_M", expected: "SmolLM 135M Base v0.2 Q4_K_M"},
		{input: "smollm:135m-base-v0.2-q5_0", expected: "SmolLM 135M Base v0.2 Q5 0"},
		{input: "smollm:135m-base-v0.2-q5_1", expected: "SmolLM 135M Base v0.2 Q5 1"},
		{input: "smollm:135m-base-v0.2-q5_K_S", expected: "SmolLM 135M Base v0.2 Q5 K_S"},
		{input: "smollm:135m-base-v0.2-q5_K_M", expected: "SmolLM 135M Base v0.2 Q5 K_M"},
		{input: "smollm:135m-base-v0.2-q6_K", expected: "SmolLM 135M Base v0.2 Q6 K"},
		{input: "smollm:135m-base-v0.2-q8_0", expected: "SmolLM 135M Base v0.2 Q8_0"},
		{input: "smollm:135m-base-v0.2-fp16", expected: "SmolLM 135M Base v0.2 FP16"},
		{input: "smollm:135m-instruct-v0.2-q2_K", expected: "SmolLM 135M Instruct v0.2 Q2_K"},
		{input: "smollm:135m-instruct-v0.2-q3_K_S", expected: "SmolLM 135M Instruct v0.2 Q3 K_S"},
		{input: "smollm:135m-instruct-v0.2-q3_K_M", expected: "SmolLM 135M Instruct v0.2 Q3 K_M"},
		{input: "smollm:135m-instruct-v0.2-q3_K_L", expected: "SmolLM 135M Instruct v0.2 Q3 K_L"},
		{input: "smollm:135m-instruct-v0.2-q4_0", expected: "SmolLM 135M Instruct v0.2 Q4_0"},
		{input: "smollm:135m-instruct-v0.2-q4_1", expected: "SmolLM 135M Instruct v0.2 Q4 1"},
		{input: "smollm:135m-instruct-v0.2-q4_K_S", expected: "SmolLM 135M Instruct v0.2 Q4 K_S"},
		{input: "smollm:135m-instruct-v0.2-q4_K_M", expected: "SmolLM 135M Instruct v0.2 Q4_K_M"},
		{input: "smollm:135m-instruct-v0.2-q5_0", expected: "SmolLM 135M Instruct v0.2 Q5 0"},
		{input: "smollm:135m-instruct-v0.2-q5_1", expected: "SmolLM 135M Instruct v0.2 Q5 1"},
		{input: "smollm:135m-instruct-v0.2-q5_K_S", expected: "SmolLM 135M Instruct v0.2 Q5 K_S"},
		{input: "smollm:135m-instruct-v0.2-q5_K_M", expected: "SmolLM 135M Instruct v0.2 Q5 K_M"},
		{input: "smollm:135m-instruct-v0.2-q6_K", expected: "SmolLM 135M Instruct v0.2 Q6 K"},
		{input: "smollm:135m-instruct-v0.2-q8_0", expected: "SmolLM 135M Instruct v0.2 Q8_0"},
		{input: "smollm:135m-instruct-v0.2-fp16", expected: "SmolLM 135M Instruct v0.2 FP16"},
		{input: "smollm:360m-base-v0.2-q2_K", expected: "SmolLM 360M Base v0.2 Q2_K"},
		{input: "smollm:360m-base-v0.2-q3_K_S", expected: "SmolLM 360M Base v0.2 Q3 K_S"},
		{input: "smollm:360m-base-v0.2-q3_K_M", expected: "SmolLM 360M Base v0.2 Q3 K_M"},
		{input: "smollm:360m-base-v0.2-q3_K_L", expected: "SmolLM 360M Base v0.2 Q3 K_L"},
		{input: "smollm:360m-base-v0.2-q4_0", expected: "SmolLM 360M Base v0.2 Q4_0"},
		{input: "smollm:360m-base-v0.2-q4_1", expected: "SmolLM 360M Base v0.2 Q4 1"},
		{input: "smollm:360m-base-v0.2-q4_K_S", expected: "SmolLM 360M Base v0.2 Q4 K_S"},
		{input: "smollm:360m-base-v0.2-q4_K_M", expected: "SmolLM 360M Base v0.2 Q4_K_M"},
		{input: "smollm:360m-base-v0.2-q5_0", expected: "SmolLM 360M Base v0.2 Q5 0"},
		{input: "smollm:360m-base-v0.2-q5_1", expected: "SmolLM 360M Base v0.2 Q5 1"},
		{input: "smollm:360m-base-v0.2-q5_K_S", expected: "SmolLM 360M Base v0.2 Q5 K_S"},
		{input: "smollm:360m-base-v0.2-q5_K_M", expected: "SmolLM 360M Base v0.2 Q5 K_M"},
		{input: "smollm:360m-base-v0.2-q6_K", expected: "SmolLM 360M Base v0.2 Q6 K"},
		{input: "smollm:360m-base-v0.2-q8_0", expected: "SmolLM 360M Base v0.2 Q8_0"},
		{input: "smollm:360m-base-v0.2-fp16", expected: "SmolLM 360M Base v0.2 FP16"},
		{input: "smollm:360m-instruct-v0.2-q2_K", expected: "SmolLM 360M Instruct v0.2 Q2_K"},
		{input: "smollm:360m-instruct-v0.2-q3_K_S", expected: "SmolLM 360M Instruct v0.2 Q3 K_S"},
		{input: "smollm:360m-instruct-v0.2-q3_K_M", expected: "SmolLM 360M Instruct v0.2 Q3 K_M"},
		{input: "smollm:360m-instruct-v0.2-q3_K_L", expected: "SmolLM 360M Instruct v0.2 Q3 K_L"},
		{input: "smollm:360m-instruct-v0.2-q4_0", expected: "SmolLM 360M Instruct v0.2 Q4_0"},
		{input: "smollm:360m-instruct-v0.2-q4_1", expected: "SmolLM 360M Instruct v0.2 Q4 1"},
		{input: "smollm:360m-instruct-v0.2-q4_K_S", expected: "SmolLM 360M Instruct v0.2 Q4 K_S"},
		{input: "smollm:360m-instruct-v0.2-q4_K_M", expected: "SmolLM 360M Instruct v0.2 Q4_K_M"},
		{input: "smollm:360m-instruct-v0.2-q5_0", expected: "SmolLM 360M Instruct v0.2 Q5 0"},
		{input: "smollm:360m-instruct-v0.2-q5_1", expected: "SmolLM 360M Instruct v0.2 Q5 1"},
		{input: "smollm:360m-instruct-v0.2-q5_K_S", expected: "SmolLM 360M Instruct v0.2 Q5 K_S"},
		{input: "smollm:360m-instruct-v0.2-q5_K_M", expected: "SmolLM 360M Instruct v0.2 Q5 K_M"},
		{input: "smollm:360m-instruct-v0.2-q6_K", expected: "SmolLM 360M Instruct v0.2 Q6 K"},
		{input: "smollm:360m-instruct-v0.2-q8_0", expected: "SmolLM 360M Instruct v0.2 Q8_0"},
		{input: "smollm:360m-instruct-v0.2-fp16", expected: "SmolLM 360M Instruct v0.2 FP16"},
		{input: "smollm:1.7b-base-v0.2-q2_K", expected: "SmolLM 1.7B Base v0.2 Q2_K"},
		{input: "smollm:1.7b-base-v0.2-q3_K_S", expected: "SmolLM 1.7B Base v0.2 Q3 K_S"},
		{input: "smollm:1.7b-base-v0.2-q3_K_M", expected: "SmolLM 1.7B Base v0.2 Q3 K_M"},
		{input: "smollm:1.7b-base-v0.2-q3_K_L", expected: "SmolLM 1.7B Base v0.2 Q3 K_L"},
		{input: "smollm:1.7b-base-v0.2-q4_0", expected: "SmolLM 1.7B Base v0.2 Q4_0"},
		{input: "smollm:1.7b-base-v0.2-q4_1", expected: "SmolLM 1.7B Base v0.2 Q4 1"},
		{input: "smollm:1.7b-base-v0.2-q4_K_S", expected: "SmolLM 1.7B Base v0.2 Q4 K_S"},
		{input: "smollm:1.7b-base-v0.2-q4_K_M", expected: "SmolLM 1.7B Base v0.2 Q4_K_M"},
		{input: "smollm:1.7b-base-v0.2-q5_0", expected: "SmolLM 1.7B Base v0.2 Q5 0"},
		{input: "smollm:1.7b-base-v0.2-q5_1", expected: "SmolLM 1.7B Base v0.2 Q5 1"},
		{input: "smollm:1.7b-base-v0.2-q5_K_S", expected: "SmolLM 1.7B Base v0.2 Q5 K_S"},
		{input: "smollm:1.7b-base-v0.2-q5_K_M", expected: "SmolLM 1.7B Base v0.2 Q5 K_M"},
		{input: "smollm:1.7b-base-v0.2-q6_K", expected: "SmolLM 1.7B Base v0.2 Q6 K"},
		{input: "smollm:1.7b-base-v0.2-q8_0", expected: "SmolLM 1.7B Base v0.2 Q8_0"},
		{input: "smollm:1.7b-base-v0.2-fp16", expected: "SmolLM 1.7B Base v0.2 FP16"},
		{input: "smollm:1.7b-instruct-v0.2-q2_K", expected: "SmolLM 1.7B Instruct v0.2 Q2_K"},
		{input: "smollm:1.7b-instruct-v0.2-q3_K_S", expected: "SmolLM 1.7B Instruct v0.2 Q3 K_S"},
		{input: "smollm:1.7b-instruct-v0.2-q3_K_M", expected: "SmolLM 1.7B Instruct v0.2 Q3 K_M"},
		{input: "smollm:1.7b-instruct-v0.2-q3_K_L", expected: "SmolLM 1.7B Instruct v0.2 Q3 K_L"},
		{input: "smollm:1.7b-instruct-v0.2-q4_0", expected: "SmolLM 1.7B Instruct v0.2 Q4_0"},
		{input: "smollm:1.7b-instruct-v0.2-q4_1", expected: "SmolLM 1.7B Instruct v0.2 Q4 1"},
		{input: "smollm:1.7b-instruct-v0.2-q4_K_S", expected: "SmolLM 1.7B Instruct v0.2 Q4 K_S"},
		{input: "smollm:1.7b-instruct-v0.2-q4_K_M", expected: "SmolLM 1.7B Instruct v0.2 Q4_K_M"},
		{input: "smollm:1.7b-instruct-v0.2-q5_0", expected: "SmolLM 1.7B Instruct v0.2 Q5 0"},
		{input: "smollm:1.7b-instruct-v0.2-q5_1", expected: "SmolLM 1.7B Instruct v0.2 Q5 1"},
		{input: "smollm:1.7b-instruct-v0.2-q5_K_S", expected: "SmolLM 1.7B Instruct v0.2 Q5 K_S"},
		{input: "smollm:1.7b-instruct-v0.2-q5_K_M", expected: "SmolLM 1.7B Instruct v0.2 Q5 K_M"},
		{input: "smollm:1.7b-instruct-v0.2-q6_K", expected: "SmolLM 1.7B Instruct v0.2 Q6 K"},
		{input: "smollm:1.7b-instruct-v0.2-q8_0", expected: "SmolLM 1.7B Instruct v0.2 Q8_0"},
		{input: "smollm:1.7b-instruct-v0.2-fp16", expected: "SmolLM 1.7B Instruct v0.2 FP16"},
		{input: "granite3.2-vision:latest", expected: "Granite3.2 Vision (latest)"},
		{input: "granite3.2-vision", expected: "Granite3.2 Vision"},
		{input: "granite3.2-vision:2b", expected: "Granite3.2 Vision 2B"},
		{input: "granite3.2-vision:2b-q4_K_M", expected: "Granite3.2 Vision 2B Q4_K_M"},
		{input: "granite3.2-vision:2b-q8_0", expected: "Granite3.2 Vision 2B Q8_0"},
		{input: "granite3.2-vision:2b-fp16", expected: "Granite3.2 Vision 2B FP16"},
		{input: "devstral:latest", expected: "Devstral (latest)"},
		{input: "devstral", expected: "Devstral"},
		{input: "devstral:24b", expected: "Devstral 24B"},
		{input: "devstral:24b-small-2505-q4_K_M", expected: "Devstral 24B Small 2505 Q4_K_M"},
		{input: "devstral:24b-small-2505-q8_0", expected: "Devstral 24B Small 2505 Q8_0"},
		{input: "devstral:24b-small-2505-fp16", expected: "Devstral 24B Small 2505 FP16"},
		{input: "wizardlm2:latest", expected: "WizardLM2 (latest)"},
		{input: "wizardlm2", expected: "WizardLM2"},
		{input: "wizardlm2:7b", expected: "WizardLM2 7B"},
		{input: "wizardlm2:8x22b", expected: "WizardLM2 8x22B"},
		{input: "wizardlm2:7b-q2_K", expected: "WizardLM2 7B Q2_K"},
		{input: "wizardlm2:7b-q3_K_S", expected: "WizardLM2 7B Q3 K_S"},
		{input: "wizardlm2:7b-q3_K_M", expected: "WizardLM2 7B Q3 K_M"},
		{input: "wizardlm2:7b-q3_K_L", expected: "WizardLM2 7B Q3 K_L"},
		{input: "wizardlm2:7b-q4_0", expected: "WizardLM2 7B Q4_0"},
		{input: "wizardlm2:7b-q4_1", expected: "WizardLM2 7B Q4 1"},
		{input: "wizardlm2:7b-q4_K_S", expected: "WizardLM2 7B Q4 K_S"},
		{input: "wizardlm2:7b-q4_K_M", expected: "WizardLM2 7B Q4_K_M"},
		{input: "wizardlm2:7b-q5_0", expected: "WizardLM2 7B Q5 0"},
		{input: "wizardlm2:7b-q5_1", expected: "WizardLM2 7B Q5 1"},
		{input: "wizardlm2:7b-q5_K_S", expected: "WizardLM2 7B Q5 K_S"},
		{input: "wizardlm2:7b-q5_K_M", expected: "WizardLM2 7B Q5 K_M"},
		{input: "wizardlm2:7b-q6_K", expected: "WizardLM2 7B Q6 K"},
		{input: "wizardlm2:7b-q8_0", expected: "WizardLM2 7B Q8_0"},
		{input: "wizardlm2:7b-fp16", expected: "WizardLM2 7B FP16"},
		{input: "wizardlm2:8x22b-q2_K", expected: "WizardLM2 8x22B Q2_K"},
		{input: "wizardlm2:8x22b-q4_0", expected: "WizardLM2 8x22B Q4_0"},
		{input: "wizardlm2:8x22b-q8_0", expected: "WizardLM2 8x22B Q8_0"},
		{input: "wizardlm2:8x22b-fp16", expected: "WizardLM2 8x22B FP16"},
		{input: "dolphin-mistral:latest", expected: "Dolphin Mistral (latest)"},
		{input: "dolphin-mistral", expected: "Dolphin Mistral"},
		{input: "dolphin-mistral:v2", expected: "Dolphin Mistral v2"},
		{input: "dolphin-mistral:v2.1", expected: "Dolphin Mistral v2.1"},
		{input: "dolphin-mistral:v2.2", expected: "Dolphin Mistral v2.2"},
		{input: "dolphin-mistral:v2.2.1", expected: "Dolphin Mistral v2.2.1"},
		{input: "dolphin-mistral:v2.6", expected: "Dolphin Mistral v2.6"},
		{input: "dolphin-mistral:v2.8", expected: "Dolphin Mistral v2.8"},
		{input: "dolphin-mistral:7b", expected: "Dolphin Mistral 7B"},
		{input: "dolphin-mistral:7b-v2", expected: "Dolphin Mistral 7B v2"},
		{input: "dolphin-mistral:7b-v2-q2_K", expected: "Dolphin Mistral 7B v2 Q2_K"},
		{input: "dolphin-mistral:7b-v2-q3_K_S", expected: "Dolphin Mistral 7B v2 Q3 K_S"},
		{input: "dolphin-mistral:7b-v2-q3_K_M", expected: "Dolphin Mistral 7B v2 Q3 K_M"},
		{input: "dolphin-mistral:7b-v2-q3_K_L", expected: "Dolphin Mistral 7B v2 Q3 K_L"},
		{input: "dolphin-mistral:7b-v2-q4_0", expected: "Dolphin Mistral 7B v2 Q4_0"},
		{input: "dolphin-mistral:7b-v2-q4_1", expected: "Dolphin Mistral 7B v2 Q4 1"},
		{input: "dolphin-mistral:7b-v2-q4_K_S", expected: "Dolphin Mistral 7B v2 Q4 K_S"},
		{input: "dolphin-mistral:7b-v2-q4_K_M", expected: "Dolphin Mistral 7B v2 Q4_K_M"},
		{input: "dolphin-mistral:7b-v2-q5_0", expected: "Dolphin Mistral 7B v2 Q5 0"},
		{input: "dolphin-mistral:7b-v2-q5_1", expected: "Dolphin Mistral 7B v2 Q5 1"},
		{input: "dolphin-mistral:7b-v2-q5_K_S", expected: "Dolphin Mistral 7B v2 Q5 K_S"},
		{input: "dolphin-mistral:7b-v2-q5_K_M", expected: "Dolphin Mistral 7B v2 Q5 K_M"},
		{input: "dolphin-mistral:7b-v2-q6_K", expected: "Dolphin Mistral 7B v2 Q6 K"},
		{input: "dolphin-mistral:7b-v2-q8_0", expected: "Dolphin Mistral 7B v2 Q8_0"},
		{input: "dolphin-mistral:7b-v2-fp16", expected: "Dolphin Mistral 7B v2 FP16"},
		{input: "dolphin-mistral:7b-v2.1", expected: "Dolphin Mistral 7B v2.1"},
		{input: "dolphin-mistral:7b-v2.1-q2_K", expected: "Dolphin Mistral 7B v2.1 Q2_K"},
		{input: "dolphin-mistral:7b-v2.1-q3_K_S", expected: "Dolphin Mistral 7B v2.1 Q3 K_S"},
		{input: "dolphin-mistral:7b-v2.1-q3_K_M", expected: "Dolphin Mistral 7B v2.1 Q3 K_M"},
		{input: "dolphin-mistral:7b-v2.1-q3_K_L", expected: "Dolphin Mistral 7B v2.1 Q3 K_L"},
		{input: "dolphin-mistral:7b-v2.1-q4_0", expected: "Dolphin Mistral 7B v2.1 Q4_0"},
		{input: "dolphin-mistral:7b-v2.1-q4_1", expected: "Dolphin Mistral 7B v2.1 Q4 1"},
		{input: "dolphin-mistral:7b-v2.1-q4_K_S", expected: "Dolphin Mistral 7B v2.1 Q4 K_S"},
		{input: "dolphin-mistral:7b-v2.1-q4_K_M", expected: "Dolphin Mistral 7B v2.1 Q4_K_M"},
		{input: "dolphin-mistral:7b-v2.1-q5_0", expected: "Dolphin Mistral 7B v2.1 Q5 0"},
		{input: "dolphin-mistral:7b-v2.1-q5_1", expected: "Dolphin Mistral 7B v2.1 Q5 1"},
		{input: "dolphin-mistral:7b-v2.1-q5_K_S", expected: "Dolphin Mistral 7B v2.1 Q5 K_S"},
		{input: "dolphin-mistral:7b-v2.1-q5_K_M", expected: "Dolphin Mistral 7B v2.1 Q5 K_M"},
		{input: "dolphin-mistral:7b-v2.1-q6_K", expected: "Dolphin Mistral 7B v2.1 Q6 K"},
		{input: "dolphin-mistral:7b-v2.1-q8_0", expected: "Dolphin Mistral 7B v2.1 Q8_0"},
		{input: "dolphin-mistral:7b-v2.1-fp16", expected: "Dolphin Mistral 7B v2.1 FP16"},
		{input: "dolphin-mistral:7b-v2.2", expected: "Dolphin Mistral 7B v2.2"},
		{input: "dolphin-mistral:7b-v2.2-q2_K", expected: "Dolphin Mistral 7B v2.2 Q2_K"},
		{input: "dolphin-mistral:7b-v2.2-q3_K_S", expected: "Dolphin Mistral 7B v2.2 Q3 K_S"},
		{input: "dolphin-mistral:7b-v2.2-q3_K_M", expected: "Dolphin Mistral 7B v2.2 Q3 K_M"},
		{input: "dolphin-mistral:7b-v2.2-q3_K_L", expected: "Dolphin Mistral 7B v2.2 Q3 K_L"},
		{input: "dolphin-mistral:7b-v2.2-q4_0", expected: "Dolphin Mistral 7B v2.2 Q4_0"},
		{input: "dolphin-mistral:7b-v2.2-q4_1", expected: "Dolphin Mistral 7B v2.2 Q4 1"},
		{input: "dolphin-mistral:7b-v2.2-q4_K_S", expected: "Dolphin Mistral 7B v2.2 Q4 K_S"},
		{input: "dolphin-mistral:7b-v2.2-q4_K_M", expected: "Dolphin Mistral 7B v2.2 Q4_K_M"},
		{input: "dolphin-mistral:7b-v2.2-q5_0", expected: "Dolphin Mistral 7B v2.2 Q5 0"},
		{input: "dolphin-mistral:7b-v2.2-q5_1", expected: "Dolphin Mistral 7B v2.2 Q5 1"},
		{input: "dolphin-mistral:7b-v2.2-q5_K_S", expected: "Dolphin Mistral 7B v2.2 Q5 K_S"},
		{input: "dolphin-mistral:7b-v2.2-q5_K_M", expected: "Dolphin Mistral 7B v2.2 Q5 K_M"},
		{input: "dolphin-mistral:7b-v2.2-q6_K", expected: "Dolphin Mistral 7B v2.2 Q6 K"},
		{input: "dolphin-mistral:7b-v2.2-q8_0", expected: "Dolphin Mistral 7B v2.2 Q8_0"},
		{input: "dolphin-mistral:7b-v2.2-fp16", expected: "Dolphin Mistral 7B v2.2 FP16"},
		{input: "dolphin-mistral:7b-v2.2.1", expected: "Dolphin Mistral 7B v2.2.1"},
		{input: "dolphin-mistral:7b-v2.2.1-q2_K", expected: "Dolphin Mistral 7B v2.2.1 Q2_K"},
		{input: "dolphin-mistral:7b-v2.2.1-q3_K_S", expected: "Dolphin Mistral 7B v2.2.1 Q3 K_S"},
		{input: "dolphin-mistral:7b-v2.2.1-q3_K_M", expected: "Dolphin Mistral 7B v2.2.1 Q3 K_M"},
		{input: "dolphin-mistral:7b-v2.2.1-q3_K_L", expected: "Dolphin Mistral 7B v2.2.1 Q3 K_L"},
		{input: "dolphin-mistral:7b-v2.2.1-q4_0", expected: "Dolphin Mistral 7B v2.2.1 Q4_0"},
		{input: "dolphin-mistral:7b-v2.2.1-q4_1", expected: "Dolphin Mistral 7B v2.2.1 Q4 1"},
		{input: "dolphin-mistral:7b-v2.2.1-q4_K_S", expected: "Dolphin Mistral 7B v2.2.1 Q4 K_S"},
		{input: "dolphin-mistral:7b-v2.2.1-q4_K_M", expected: "Dolphin Mistral 7B v2.2.1 Q4_K_M"},
		{input: "dolphin-mistral:7b-v2.2.1-q5_0", expected: "Dolphin Mistral 7B v2.2.1 Q5 0"},
		{input: "dolphin-mistral:7b-v2.2.1-q5_1", expected: "Dolphin Mistral 7B v2.2.1 Q5 1"},
		{input: "dolphin-mistral:7b-v2.2.1-q5_K_S", expected: "Dolphin Mistral 7B v2.2.1 Q5 K_S"},
		{input: "dolphin-mistral:7b-v2.2.1-q5_K_M", expected: "Dolphin Mistral 7B v2.2.1 Q5 K_M"},
		{input: "dolphin-mistral:7b-v2.2.1-q6_K", expected: "Dolphin Mistral 7B v2.2.1 Q6 K"},
		{input: "dolphin-mistral:7b-v2.2.1-q8_0", expected: "Dolphin Mistral 7B v2.2.1 Q8_0"},
		{input: "dolphin-mistral:7b-v2.2.1-fp16", expected: "Dolphin Mistral 7B v2.2.1 FP16"},
		{input: "dolphin-mistral:7b-v2.6", expected: "Dolphin Mistral 7B v2.6"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser", expected: "Dolphin Mistral 7B v2.6 Dpo Laser"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q2_K", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q2_K"},
		{input: "dolphin-mistral:7b-v2.6-q2_K", expected: "Dolphin Mistral 7B v2.6 Q2_K"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q3_K_S", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q3 K_S"},
		{input: "dolphin-mistral:7b-v2.6-q3_K_S", expected: "Dolphin Mistral 7B v2.6 Q3 K_S"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q3_K_M", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q3 K_M"},
		{input: "dolphin-mistral:7b-v2.6-q3_K_M", expected: "Dolphin Mistral 7B v2.6 Q3 K_M"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q3_K_L", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q3 K_L"},
		{input: "dolphin-mistral:7b-v2.6-q3_K_L", expected: "Dolphin Mistral 7B v2.6 Q3 K_L"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q4_0", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q4_0"},
		{input: "dolphin-mistral:7b-v2.6-q4_0", expected: "Dolphin Mistral 7B v2.6 Q4_0"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q4_1", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q4 1"},
		{input: "dolphin-mistral:7b-v2.6-q4_1", expected: "Dolphin Mistral 7B v2.6 Q4 1"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q4_K_S", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q4 K_S"},
		{input: "dolphin-mistral:7b-v2.6-q4_K_S", expected: "Dolphin Mistral 7B v2.6 Q4 K_S"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q4_K_M", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q4_K_M"},
		{input: "dolphin-mistral:7b-v2.6-q4_K_M", expected: "Dolphin Mistral 7B v2.6 Q4_K_M"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q5_0", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q5 0"},
		{input: "dolphin-mistral:7b-v2.6-q5_0", expected: "Dolphin Mistral 7B v2.6 Q5 0"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q5_1", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q5 1"},
		{input: "dolphin-mistral:7b-v2.6-q5_1", expected: "Dolphin Mistral 7B v2.6 Q5 1"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q5_K_S", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q5 K_S"},
		{input: "dolphin-mistral:7b-v2.6-q5_K_S", expected: "Dolphin Mistral 7B v2.6 Q5 K_S"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q5_K_M", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q5 K_M"},
		{input: "dolphin-mistral:7b-v2.6-q5_K_M", expected: "Dolphin Mistral 7B v2.6 Q5 K_M"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q6_K", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q6 K"},
		{input: "dolphin-mistral:7b-v2.6-q6_K", expected: "Dolphin Mistral 7B v2.6 Q6 K"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-q8_0", expected: "Dolphin Mistral 7B v2.6 Dpo Laser Q8_0"},
		{input: "dolphin-mistral:7b-v2.6-q8_0", expected: "Dolphin Mistral 7B v2.6 Q8_0"},
		{input: "dolphin-mistral:7b-v2.6-dpo-laser-fp16", expected: "Dolphin Mistral 7B v2.6 Dpo Laser FP16"},
		{input: "dolphin-mistral:7b-v2.6-fp16", expected: "Dolphin Mistral 7B v2.6 FP16"},
		{input: "dolphin-mistral:7b-v2.8", expected: "Dolphin Mistral 7B v2.8"},
		{input: "dolphin-mistral:7b-v2.8-q2_K", expected: "Dolphin Mistral 7B v2.8 Q2_K"},
		{input: "dolphin-mistral:7b-v2.8-q3_K_S", expected: "Dolphin Mistral 7B v2.8 Q3 K_S"},
		{input: "dolphin-mistral:7b-v2.8-q3_K_M", expected: "Dolphin Mistral 7B v2.8 Q3 K_M"},
		{input: "dolphin-mistral:7b-v2.8-q3_K_L", expected: "Dolphin Mistral 7B v2.8 Q3 K_L"},
		{input: "dolphin-mistral:7b-v2.8-q4_0", expected: "Dolphin Mistral 7B v2.8 Q4_0"},
		{input: "dolphin-mistral:7b-v2.8-q4_1", expected: "Dolphin Mistral 7B v2.8 Q4 1"},
		{input: "dolphin-mistral:7b-v2.8-q4_K_S", expected: "Dolphin Mistral 7B v2.8 Q4 K_S"},
		{input: "dolphin-mistral:7b-v2.8-q4_K_M", expected: "Dolphin Mistral 7B v2.8 Q4_K_M"},
		{input: "dolphin-mistral:7b-v2.8-q5_0", expected: "Dolphin Mistral 7B v2.8 Q5 0"},
		{input: "dolphin-mistral:7b-v2.8-q5_1", expected: "Dolphin Mistral 7B v2.8 Q5 1"},
		{input: "dolphin-mistral:7b-v2.8-q5_K_S", expected: "Dolphin Mistral 7B v2.8 Q5 K_S"},
		{input: "dolphin-mistral:7b-v2.8-q5_K_M", expected: "Dolphin Mistral 7B v2.8 Q5 K_M"},
		{input: "dolphin-mistral:7b-v2.8-q6_K", expected: "Dolphin Mistral 7B v2.8 Q6 K"},
		{input: "dolphin-mistral:7b-v2.8-q8_0", expected: "Dolphin Mistral 7B v2.8 Q8_0"},
		{input: "dolphin-mistral:7b-v2.8-fp16", expected: "Dolphin Mistral 7B v2.8 FP16"},
		{input: "deepcoder:latest", expected: "Deepcoder (latest)"},
		{input: "deepcoder", expected: "Deepcoder"},
		{input: "deepcoder:1.5b", expected: "Deepcoder 1.5B"},
		{input: "deepcoder:14b", expected: "Deepcoder 14B"},
		{input: "deepcoder:1.5b-preview-q4_K_M", expected: "Deepcoder 1.5B Preview Q4_K_M"},
		{input: "deepcoder:1.5b-preview-q8_0", expected: "Deepcoder 1.5B Preview Q8_0"},
		{input: "deepcoder:1.5b-preview-fp16", expected: "Deepcoder 1.5B Preview FP16"},
		{input: "deepcoder:14b-preview-q4_K_M", expected: "Deepcoder 14B Preview Q4_K_M"},
		{input: "deepcoder:14b-preview-q8_0", expected: "Deepcoder 14B Preview Q8_0"},
		{input: "deepcoder:14b-preview-fp16", expected: "Deepcoder 14B Preview FP16"},
		{input: "moondream:latest", expected: "Moondream (latest)"},
		{input: "moondream", expected: "Moondream"},
		{input: "moondream:v2", expected: "Moondream v2"},
		{input: "moondream:1.8b", expected: "Moondream 1.8B"},
		{input: "moondream:1.8b-v2-q2_K", expected: "Moondream 1.8B v2 Q2_K"},
		{input: "moondream:1.8b-v2-q3_K_S", expected: "Moondream 1.8B v2 Q3 K_S"},
		{input: "moondream:1.8b-v2-q3_K_M", expected: "Moondream 1.8B v2 Q3 K_M"},
		{input: "moondream:1.8b-v2-q3_K_L", expected: "Moondream 1.8B v2 Q3 K_L"},
		{input: "moondream:1.8b-v2-q4_0", expected: "Moondream 1.8B v2 Q4_0"},
		{input: "moondream:1.8b-v2-q4_1", expected: "Moondream 1.8B v2 Q4 1"},
		{input: "moondream:1.8b-v2-q4_K_S", expected: "Moondream 1.8B v2 Q4 K_S"},
		{input: "moondream:1.8b-v2-q4_K_M", expected: "Moondream 1.8B v2 Q4_K_M"},
		{input: "moondream:1.8b-v2-q5_0", expected: "Moondream 1.8B v2 Q5 0"},
		{input: "moondream:1.8b-v2-q5_1", expected: "Moondream 1.8B v2 Q5 1"},
		{input: "moondream:1.8b-v2-q5_K_S", expected: "Moondream 1.8B v2 Q5 K_S"},
		{input: "moondream:1.8b-v2-q5_K_M", expected: "Moondream 1.8B v2 Q5 K_M"},
		{input: "moondream:1.8b-v2-q6_K", expected: "Moondream 1.8B v2 Q6 K"},
		{input: "moondream:1.8b-v2-q8_0", expected: "Moondream 1.8B v2 Q8_0"},
		{input: "moondream:1.8b-v2-fp16", expected: "Moondream 1.8B v2 FP16"},
		{input: "mistral-small3.1:latest", expected: "Mistral Small3.1 (latest)"},
		{input: "mistral-small3.1", expected: "Mistral Small3.1"},
		{input: "mistral-small3.1:24b", expected: "Mistral Small3.1 24B"},
		{input: "mistral-small3.1:24b-instruct-2503-q4_K_M", expected: "Mistral Small3.1 24B Instruct 2503 Q4_K_M"},
		{input: "mistral-small3.1:24b-instruct-2503-q8_0", expected: "Mistral Small3.1 24B Instruct 2503 Q8_0"},
		{input: "mistral-small3.1:24b-instruct-2503-fp16", expected: "Mistral Small3.1 24B Instruct 2503 FP16"},
		{input: "command-r:latest", expected: "Command R (latest)"},
		{input: "command-r", expected: "Command R"},
		{input: "command-r:v0.1", expected: "Command R v0.1"},
		{input: "command-r:35b", expected: "Command R 35B"},
		{input: "command-r:35b-08-2024-q2_K", expected: "Command R 35B (2024-08) Q2_K"},
		{input: "command-r:35b-08-2024-q3_K_S", expected: "Command R 35B (2024-08) Q3 K_S"},
		{input: "command-r:35b-08-2024-q3_K_M", expected: "Command R 35B (2024-08) Q3 K_M"},
		{input: "command-r:35b-08-2024-q3_K_L", expected: "Command R 35B (2024-08) Q3 K_L"},
		{input: "command-r:35b-08-2024-q4_0", expected: "Command R 35B (2024-08) Q4_0"},
		{input: "command-r:35b-08-2024-q4_1", expected: "Command R 35B (2024-08) Q4 1"},
		{input: "command-r:35b-08-2024-q4_K_S", expected: "Command R 35B (2024-08) Q4 K_S"},
		{input: "command-r:35b-08-2024-q4_K_M", expected: "Command R 35B (2024-08) Q4_K_M"},
		{input: "command-r:35b-08-2024-q5_0", expected: "Command R 35B (2024-08) Q5 0"},
		{input: "command-r:35b-08-2024-q5_1", expected: "Command R 35B (2024-08) Q5 1"},
		{input: "command-r:35b-08-2024-q5_K_S", expected: "Command R 35B (2024-08) Q5 K_S"},
		{input: "command-r:35b-08-2024-q5_K_M", expected: "Command R 35B (2024-08) Q5 K_M"},
		{input: "command-r:35b-08-2024-q6_K", expected: "Command R 35B (2024-08) Q6 K"},
		{input: "command-r:35b-08-2024-q8_0", expected: "Command R 35B (2024-08) Q8_0"},
		{input: "command-r:35b-08-2024-fp16", expected: "Command R 35B (2024-08) FP16"},
		{input: "command-r:35b-v0.1-q2_K", expected: "Command R 35B v0.1 Q2_K"},
		{input: "command-r:35b-v0.1-q3_K_S", expected: "Command R 35B v0.1 Q3 K_S"},
		{input: "command-r:35b-v0.1-q3_K_M", expected: "Command R 35B v0.1 Q3 K_M"},
		{input: "command-r:35b-v0.1-q3_K_L", expected: "Command R 35B v0.1 Q3 K_L"},
		{input: "command-r:35b-v0.1-q4_0", expected: "Command R 35B v0.1 Q4_0"},
		{input: "command-r:35b-v0.1-q4_1", expected: "Command R 35B v0.1 Q4 1"},
		{input: "command-r:35b-v0.1-q4_K_S", expected: "Command R 35B v0.1 Q4 K_S"},
		{input: "command-r:35b-v0.1-q4_K_M", expected: "Command R 35B v0.1 Q4_K_M"},
		{input: "command-r:35b-v0.1-q5_1", expected: "Command R 35B v0.1 Q5 1"},
		{input: "command-r:35b-v0.1-q5_K_S", expected: "Command R 35B v0.1 Q5 K_S"},
		{input: "command-r:35b-v0.1-q5_K_M", expected: "Command R 35B v0.1 Q5 K_M"},
		{input: "command-r:35b-v0.1-q6_K", expected: "Command R 35B v0.1 Q6 K"},
		{input: "command-r:35b-v0.1-q8_0", expected: "Command R 35B v0.1 Q8_0"},
		{input: "command-r:35b-v0.1-fp16", expected: "Command R 35B v0.1 FP16"},
		{input: "granite-code:latest", expected: "Granite Code (latest)"},
		{input: "granite-code", expected: "Granite Code"},
		{input: "granite-code:3b", expected: "Granite Code 3B"},
		{input: "granite-code:8b", expected: "Granite Code 8B"},
		{input: "granite-code:20b", expected: "Granite Code 20B"},
		{input: "granite-code:34b", expected: "Granite Code 34B"},
		{input: "granite-code:3b-base", expected: "Granite Code 3B Base"},
		{input: "granite-code:3b-base-q2_K", expected: "Granite Code 3B Base Q2_K"},
		{input: "granite-code:3b-base-q3_K_S", expected: "Granite Code 3B Base Q3 K_S"},
		{input: "granite-code:3b-base-q3_K_M", expected: "Granite Code 3B Base Q3 K_M"},
		{input: "granite-code:3b-base-q3_K_L", expected: "Granite Code 3B Base Q3 K_L"},
		{input: "granite-code:3b-base-q4_0", expected: "Granite Code 3B Base Q4_0"},
		{input: "granite-code:3b-base-q4_1", expected: "Granite Code 3B Base Q4 1"},
		{input: "granite-code:3b-base-q4_K_S", expected: "Granite Code 3B Base Q4 K_S"},
		{input: "granite-code:3b-base-q4_K_M", expected: "Granite Code 3B Base Q4_K_M"},
		{input: "granite-code:3b-base-q5_0", expected: "Granite Code 3B Base Q5 0"},
		{input: "granite-code:3b-base-q5_1", expected: "Granite Code 3B Base Q5 1"},
		{input: "granite-code:3b-base-q5_K_S", expected: "Granite Code 3B Base Q5 K_S"},
		{input: "granite-code:3b-base-q5_K_M", expected: "Granite Code 3B Base Q5 K_M"},
		{input: "granite-code:3b-base-q6_K", expected: "Granite Code 3B Base Q6 K"},
		{input: "granite-code:3b-base-q8_0", expected: "Granite Code 3B Base Q8_0"},
		{input: "granite-code:3b-base-fp16", expected: "Granite Code 3B Base FP16"},
		{input: "granite-code:3b-instruct", expected: "Granite Code 3B Instruct"},
		{input: "granite-code:3b-instruct-128k-q2_K", expected: "Granite Code 3B Instruct 128K Q2_K"},
		{input: "granite-code:3b-instruct-q2_K", expected: "Granite Code 3B Instruct Q2_K"},
		{input: "granite-code:3b-instruct-128k-q3_K_S", expected: "Granite Code 3B Instruct 128K Q3 K_S"},
		{input: "granite-code:3b-instruct-q3_K_S", expected: "Granite Code 3B Instruct Q3 K_S"},
		{input: "granite-code:3b-instruct-128k-q3_K_M", expected: "Granite Code 3B Instruct 128K Q3 K_M"},
		{input: "granite-code:3b-instruct-q3_K_M", expected: "Granite Code 3B Instruct Q3 K_M"},
		{input: "granite-code:3b-instruct-128k-q3_K_L", expected: "Granite Code 3B Instruct 128K Q3 K_L"},
		{input: "granite-code:3b-instruct-q3_K_L", expected: "Granite Code 3B Instruct Q3 K_L"},
		{input: "granite-code:3b-instruct-128k-q4_0", expected: "Granite Code 3B Instruct 128K Q4_0"},
		{input: "granite-code:3b-instruct-q4_0", expected: "Granite Code 3B Instruct Q4_0"},
		{input: "granite-code:3b-instruct-128k-q4_1", expected: "Granite Code 3B Instruct 128K Q4 1"},
		{input: "granite-code:3b-instruct-q4_1", expected: "Granite Code 3B Instruct Q4 1"},
		{input: "granite-code:3b-instruct-128k-q4_K_S", expected: "Granite Code 3B Instruct 128K Q4 K_S"},
		{input: "granite-code:3b-instruct-q4_K_S", expected: "Granite Code 3B Instruct Q4 K_S"},
		{input: "granite-code:3b-instruct-128k-q4_K_M", expected: "Granite Code 3B Instruct 128K Q4_K_M"},
		{input: "granite-code:3b-instruct-q4_K_M", expected: "Granite Code 3B Instruct Q4_K_M"},
		{input: "granite-code:3b-instruct-128k-q5_0", expected: "Granite Code 3B Instruct 128K Q5 0"},
		{input: "granite-code:3b-instruct-q5_0", expected: "Granite Code 3B Instruct Q5 0"},
		{input: "granite-code:3b-instruct-128k-q5_1", expected: "Granite Code 3B Instruct 128K Q5 1"},
		{input: "granite-code:3b-instruct-q5_1", expected: "Granite Code 3B Instruct Q5 1"},
		{input: "granite-code:3b-instruct-128k-q5_K_S", expected: "Granite Code 3B Instruct 128K Q5 K_S"},
		{input: "granite-code:3b-instruct-q5_K_S", expected: "Granite Code 3B Instruct Q5 K_S"},
		{input: "granite-code:3b-instruct-128k-q5_K_M", expected: "Granite Code 3B Instruct 128K Q5 K_M"},
		{input: "granite-code:3b-instruct-q5_K_M", expected: "Granite Code 3B Instruct Q5 K_M"},
		{input: "granite-code:3b-instruct-128k-q6_K", expected: "Granite Code 3B Instruct 128K Q6 K"},
		{input: "granite-code:3b-instruct-q6_K", expected: "Granite Code 3B Instruct Q6 K"},
		{input: "granite-code:3b-instruct-128k-q8_0", expected: "Granite Code 3B Instruct 128K Q8_0"},
		{input: "granite-code:3b-instruct-q8_0", expected: "Granite Code 3B Instruct Q8_0"},
		{input: "granite-code:3b-instruct-128k-fp16", expected: "Granite Code 3B Instruct 128K FP16"},
		{input: "granite-code:3b-instruct-fp16", expected: "Granite Code 3B Instruct FP16"},
		{input: "granite-code:8b-base", expected: "Granite Code 8B Base"},
		{input: "granite-code:8b-base-q2_K", expected: "Granite Code 8B Base Q2_K"},
		{input: "granite-code:8b-base-q3_K_S", expected: "Granite Code 8B Base Q3 K_S"},
		{input: "granite-code:8b-base-q3_K_M", expected: "Granite Code 8B Base Q3 K_M"},
		{input: "granite-code:8b-base-q3_K_L", expected: "Granite Code 8B Base Q3 K_L"},
		{input: "granite-code:8b-base-q4_0", expected: "Granite Code 8B Base Q4_0"},
		{input: "granite-code:8b-base-q4_1", expected: "Granite Code 8B Base Q4 1"},
		{input: "granite-code:8b-base-q4_K_S", expected: "Granite Code 8B Base Q4 K_S"},
		{input: "granite-code:8b-base-q4_K_M", expected: "Granite Code 8B Base Q4_K_M"},
		{input: "granite-code:8b-base-q5_0", expected: "Granite Code 8B Base Q5 0"},
		{input: "granite-code:8b-base-q5_1", expected: "Granite Code 8B Base Q5 1"},
		{input: "granite-code:8b-base-q5_K_S", expected: "Granite Code 8B Base Q5 K_S"},
		{input: "granite-code:8b-base-q5_K_M", expected: "Granite Code 8B Base Q5 K_M"},
		{input: "granite-code:8b-base-q6_K", expected: "Granite Code 8B Base Q6 K"},
		{input: "granite-code:8b-base-q8_0", expected: "Granite Code 8B Base Q8_0"},
		{input: "granite-code:8b-base-fp16", expected: "Granite Code 8B Base FP16"},
		{input: "granite-code:8b-instruct", expected: "Granite Code 8B Instruct"},
		{input: "granite-code:8b-instruct-q2_K", expected: "Granite Code 8B Instruct Q2_K"},
		{input: "granite-code:8b-instruct-q3_K_S", expected: "Granite Code 8B Instruct Q3 K_S"},
		{input: "granite-code:8b-instruct-q3_K_M", expected: "Granite Code 8B Instruct Q3 K_M"},
		{input: "granite-code:8b-instruct-q3_K_L", expected: "Granite Code 8B Instruct Q3 K_L"},
		{input: "granite-code:8b-instruct-128k-q4_0", expected: "Granite Code 8B Instruct 128K Q4_0"},
		{input: "granite-code:8b-instruct-q4_0", expected: "Granite Code 8B Instruct Q4_0"},
		{input: "granite-code:8b-instruct-128k-q4_1", expected: "Granite Code 8B Instruct 128K Q4 1"},
		{input: "granite-code:8b-instruct-q4_1", expected: "Granite Code 8B Instruct Q4 1"},
		{input: "granite-code:8b-instruct-q4_K_S", expected: "Granite Code 8B Instruct Q4 K_S"},
		{input: "granite-code:8b-instruct-q4_K_M", expected: "Granite Code 8B Instruct Q4_K_M"},
		{input: "granite-code:8b-instruct-q5_0", expected: "Granite Code 8B Instruct Q5 0"},
		{input: "granite-code:8b-instruct-q5_1", expected: "Granite Code 8B Instruct Q5 1"},
		{input: "granite-code:8b-instruct-q5_K_S", expected: "Granite Code 8B Instruct Q5 K_S"},
		{input: "granite-code:8b-instruct-q5_K_M", expected: "Granite Code 8B Instruct Q5 K_M"},
		{input: "granite-code:8b-instruct-q6_K", expected: "Granite Code 8B Instruct Q6 K"},
		{input: "granite-code:8b-instruct-q8_0", expected: "Granite Code 8B Instruct Q8_0"},
		{input: "granite-code:8b-instruct-fp16", expected: "Granite Code 8B Instruct FP16"},
		{input: "granite-code:20b-base", expected: "Granite Code 20B Base"},
		{input: "granite-code:20b-base-q2_K", expected: "Granite Code 20B Base Q2_K"},
		{input: "granite-code:20b-base-q3_K_S", expected: "Granite Code 20B Base Q3 K_S"},
		{input: "granite-code:20b-base-q3_K_M", expected: "Granite Code 20B Base Q3 K_M"},
		{input: "granite-code:20b-base-q3_K_L", expected: "Granite Code 20B Base Q3 K_L"},
		{input: "granite-code:20b-base-q4_0", expected: "Granite Code 20B Base Q4_0"},
		{input: "granite-code:20b-base-q4_1", expected: "Granite Code 20B Base Q4 1"},
		{input: "granite-code:20b-base-q4_K_S", expected: "Granite Code 20B Base Q4 K_S"},
		{input: "granite-code:20b-base-q4_K_M", expected: "Granite Code 20B Base Q4_K_M"},
		{input: "granite-code:20b-base-q5_0", expected: "Granite Code 20B Base Q5 0"},
		{input: "granite-code:20b-base-q5_1", expected: "Granite Code 20B Base Q5 1"},
		{input: "granite-code:20b-base-q5_K_S", expected: "Granite Code 20B Base Q5 K_S"},
		{input: "granite-code:20b-base-q5_K_M", expected: "Granite Code 20B Base Q5 K_M"},
		{input: "granite-code:20b-base-q6_K", expected: "Granite Code 20B Base Q6 K"},
		{input: "granite-code:20b-base-q8_0", expected: "Granite Code 20B Base Q8_0"},
		{input: "granite-code:20b-base-fp16", expected: "Granite Code 20B Base FP16"},
		{input: "granite-code:20b-instruct", expected: "Granite Code 20B Instruct"},
		{input: "granite-code:20b-instruct-8k-q2_K", expected: "Granite Code 20B Instruct 8K Q2_K"},
		{input: "granite-code:20b-instruct-q2_K", expected: "Granite Code 20B Instruct Q2_K"},
		{input: "granite-code:20b-instruct-8k-q3_K_S", expected: "Granite Code 20B Instruct 8K Q3 K_S"},
		{input: "granite-code:20b-instruct-q3_K_S", expected: "Granite Code 20B Instruct Q3 K_S"},
		{input: "granite-code:20b-instruct-8k-q3_K_M", expected: "Granite Code 20B Instruct 8K Q3 K_M"},
		{input: "granite-code:20b-instruct-q3_K_M", expected: "Granite Code 20B Instruct Q3 K_M"},
		{input: "granite-code:20b-instruct-8k-q3_K_L", expected: "Granite Code 20B Instruct 8K Q3 K_L"},
		{input: "granite-code:20b-instruct-q3_K_L", expected: "Granite Code 20B Instruct Q3 K_L"},
		{input: "granite-code:20b-instruct-8k-q4_0", expected: "Granite Code 20B Instruct 8K Q4_0"},
		{input: "granite-code:20b-instruct-q4_0", expected: "Granite Code 20B Instruct Q4_0"},
		{input: "granite-code:20b-instruct-8k-q4_1", expected: "Granite Code 20B Instruct 8K Q4 1"},
		{input: "granite-code:20b-instruct-q4_1", expected: "Granite Code 20B Instruct Q4 1"},
		{input: "granite-code:20b-instruct-8k-q4_K_S", expected: "Granite Code 20B Instruct 8K Q4 K_S"},
		{input: "granite-code:20b-instruct-q4_K_S", expected: "Granite Code 20B Instruct Q4 K_S"},
		{input: "granite-code:20b-instruct-8k-q4_K_M", expected: "Granite Code 20B Instruct 8K Q4_K_M"},
		{input: "granite-code:20b-instruct-q4_K_M", expected: "Granite Code 20B Instruct Q4_K_M"},
		{input: "granite-code:20b-instruct-8k-q5_0", expected: "Granite Code 20B Instruct 8K Q5 0"},
		{input: "granite-code:20b-instruct-q5_0", expected: "Granite Code 20B Instruct Q5 0"},
		{input: "granite-code:20b-instruct-8k-q5_1", expected: "Granite Code 20B Instruct 8K Q5 1"},
		{input: "granite-code:20b-instruct-q5_1", expected: "Granite Code 20B Instruct Q5 1"},
		{input: "granite-code:20b-instruct-8k-q5_K_S", expected: "Granite Code 20B Instruct 8K Q5 K_S"},
		{input: "granite-code:20b-instruct-q5_K_S", expected: "Granite Code 20B Instruct Q5 K_S"},
		{input: "granite-code:20b-instruct-8k-q5_K_M", expected: "Granite Code 20B Instruct 8K Q5 K_M"},
		{input: "granite-code:20b-instruct-q5_K_M", expected: "Granite Code 20B Instruct Q5 K_M"},
		{input: "granite-code:20b-instruct-8k-q6_K", expected: "Granite Code 20B Instruct 8K Q6 K"},
		{input: "granite-code:20b-instruct-q6_K", expected: "Granite Code 20B Instruct Q6 K"},
		{input: "granite-code:20b-instruct-8k-q8_0", expected: "Granite Code 20B Instruct 8K Q8_0"},
		{input: "granite-code:20b-instruct-q8_0", expected: "Granite Code 20B Instruct Q8_0"},
		{input: "granite-code:20b-instruct-8k-fp16", expected: "Granite Code 20B Instruct 8K FP16"},
		{input: "granite-code:34b-base", expected: "Granite Code 34B Base"},
		{input: "granite-code:34b-base-q2_K", expected: "Granite Code 34B Base Q2_K"},
		{input: "granite-code:34b-base-q3_K_S", expected: "Granite Code 34B Base Q3 K_S"},
		{input: "granite-code:34b-base-q3_K_M", expected: "Granite Code 34B Base Q3 K_M"},
		{input: "granite-code:34b-base-q3_K_L", expected: "Granite Code 34B Base Q3 K_L"},
		{input: "granite-code:34b-base-q4_0", expected: "Granite Code 34B Base Q4_0"},
		{input: "granite-code:34b-base-q4_1", expected: "Granite Code 34B Base Q4 1"},
		{input: "granite-code:34b-base-q4_K_S", expected: "Granite Code 34B Base Q4 K_S"},
		{input: "granite-code:34b-base-q4_K_M", expected: "Granite Code 34B Base Q4_K_M"},
		{input: "granite-code:34b-base-q5_0", expected: "Granite Code 34B Base Q5 0"},
		{input: "granite-code:34b-base-q5_1", expected: "Granite Code 34B Base Q5 1"},
		{input: "granite-code:34b-base-q5_K_S", expected: "Granite Code 34B Base Q5 K_S"},
		{input: "granite-code:34b-base-q5_K_M", expected: "Granite Code 34B Base Q5 K_M"},
		{input: "granite-code:34b-base-q6_K", expected: "Granite Code 34B Base Q6 K"},
		{input: "granite-code:34b-base-q8_0", expected: "Granite Code 34B Base Q8_0"},
		{input: "granite-code:34b-instruct", expected: "Granite Code 34B Instruct"},
		{input: "granite-code:34b-instruct-q2_K", expected: "Granite Code 34B Instruct Q2_K"},
		{input: "granite-code:34b-instruct-q3_K_S", expected: "Granite Code 34B Instruct Q3 K_S"},
		{input: "granite-code:34b-instruct-q3_K_M", expected: "Granite Code 34B Instruct Q3 K_M"},
		{input: "granite-code:34b-instruct-q3_K_L", expected: "Granite Code 34B Instruct Q3 K_L"},
		{input: "granite-code:34b-instruct-q4_0", expected: "Granite Code 34B Instruct Q4_0"},
		{input: "granite-code:34b-instruct-q4_1", expected: "Granite Code 34B Instruct Q4 1"},
		{input: "granite-code:34b-instruct-q4_K_S", expected: "Granite Code 34B Instruct Q4 K_S"},
		{input: "granite-code:34b-instruct-q4_K_M", expected: "Granite Code 34B Instruct Q4_K_M"},
		{input: "granite-code:34b-instruct-q5_0", expected: "Granite Code 34B Instruct Q5 0"},
		{input: "granite-code:34b-instruct-q5_1", expected: "Granite Code 34B Instruct Q5 1"},
		{input: "granite-code:34b-instruct-q5_K_S", expected: "Granite Code 34B Instruct Q5 K_S"},
		{input: "granite-code:34b-instruct-q5_K_M", expected: "Granite Code 34B Instruct Q5 K_M"},
		{input: "granite-code:34b-instruct-q6_K", expected: "Granite Code 34B Instruct Q6 K"},
		{input: "granite-code:34b-instruct-q8_0", expected: "Granite Code 34B Instruct Q8_0"},
		{input: "phi3.5:latest", expected: "Phi3.5 (latest)"},
		{input: "phi3.5", expected: "Phi3.5"},
		{input: "phi3.5:3.8b", expected: "Phi3.5 3.8B"},
		{input: "phi3.5:3.8b-mini-instruct-q2_K", expected: "Phi3.5 3.8B Mini Instruct Q2_K"},
		{input: "phi3.5:3.8b-mini-instruct-q3_K_S", expected: "Phi3.5 3.8B Mini Instruct Q3 K_S"},
		{input: "phi3.5:3.8b-mini-instruct-q3_K_M", expected: "Phi3.5 3.8B Mini Instruct Q3 K_M"},
		{input: "phi3.5:3.8b-mini-instruct-q3_K_L", expected: "Phi3.5 3.8B Mini Instruct Q3 K_L"},
		{input: "phi3.5:3.8b-mini-instruct-q4_0", expected: "Phi3.5 3.8B Mini Instruct Q4_0"},
		{input: "phi3.5:3.8b-mini-instruct-q4_1", expected: "Phi3.5 3.8B Mini Instruct Q4 1"},
		{input: "phi3.5:3.8b-mini-instruct-q4_K_S", expected: "Phi3.5 3.8B Mini Instruct Q4 K_S"},
		{input: "phi3.5:3.8b-mini-instruct-q4_K_M", expected: "Phi3.5 3.8B Mini Instruct Q4_K_M"},
		{input: "phi3.5:3.8b-mini-instruct-q5_0", expected: "Phi3.5 3.8B Mini Instruct Q5 0"},
		{input: "phi3.5:3.8b-mini-instruct-q5_1", expected: "Phi3.5 3.8B Mini Instruct Q5 1"},
		{input: "phi3.5:3.8b-mini-instruct-q5_K_S", expected: "Phi3.5 3.8B Mini Instruct Q5 K_S"},
		{input: "phi3.5:3.8b-mini-instruct-q5_K_M", expected: "Phi3.5 3.8B Mini Instruct Q5 K_M"},
		{input: "phi3.5:3.8b-mini-instruct-q6_K", expected: "Phi3.5 3.8B Mini Instruct Q6 K"},
		{input: "phi3.5:3.8b-mini-instruct-q8_0", expected: "Phi3.5 3.8B Mini Instruct Q8_0"},
		{input: "phi3.5:3.8b-mini-instruct-fp16", expected: "Phi3.5 3.8B Mini Instruct FP16"},
		{input: "hermes3:latest", expected: "Hermes3 (latest)"},
		{input: "hermes3", expected: "Hermes3"},
		{input: "hermes3:3b", expected: "Hermes3 3B"},
		{input: "hermes3:8b", expected: "Hermes3 8B"},
		{input: "hermes3:70b", expected: "Hermes3 70B"},
		{input: "hermes3:405b", expected: "Hermes3 405B"},
		{input: "hermes3:3b-llama3.2-q2_K", expected: "Hermes3 3B Llama3.2 Q2_K"},
		{input: "hermes3:3b-llama3.2-q3_K_S", expected: "Hermes3 3B Llama3.2 Q3 K_S"},
		{input: "hermes3:3b-llama3.2-q3_K_M", expected: "Hermes3 3B Llama3.2 Q3 K_M"},
		{input: "hermes3:3b-llama3.2-q3_K_L", expected: "Hermes3 3B Llama3.2 Q3 K_L"},
		{input: "hermes3:3b-llama3.2-q4_0", expected: "Hermes3 3B Llama3.2 Q4_0"},
		{input: "hermes3:3b-llama3.2-q4_1", expected: "Hermes3 3B Llama3.2 Q4 1"},
		{input: "hermes3:3b-llama3.2-q4_K_S", expected: "Hermes3 3B Llama3.2 Q4 K_S"},
		{input: "hermes3:3b-llama3.2-q4_K_M", expected: "Hermes3 3B Llama3.2 Q4_K_M"},
		{input: "hermes3:3b-llama3.2-q5_0", expected: "Hermes3 3B Llama3.2 Q5 0"},
		{input: "hermes3:3b-llama3.2-q5_1", expected: "Hermes3 3B Llama3.2 Q5 1"},
		{input: "hermes3:3b-llama3.2-q5_K_S", expected: "Hermes3 3B Llama3.2 Q5 K_S"},
		{input: "hermes3:3b-llama3.2-q5_K_M", expected: "Hermes3 3B Llama3.2 Q5 K_M"},
		{input: "hermes3:3b-llama3.2-q6_K", expected: "Hermes3 3B Llama3.2 Q6 K"},
		{input: "hermes3:3b-llama3.2-q8_0", expected: "Hermes3 3B Llama3.2 Q8_0"},
		{input: "hermes3:3b-llama3.2-fp16", expected: "Hermes3 3B Llama3.2 FP16"},
		{input: "hermes3:8b-llama3.1-q2_K", expected: "Hermes3 8B Llama3.1 Q2_K"},
		{input: "hermes3:8b-llama3.1-q3_K_S", expected: "Hermes3 8B Llama3.1 Q3 K_S"},
		{input: "hermes3:8b-llama3.1-q3_K_M", expected: "Hermes3 8B Llama3.1 Q3 K_M"},
		{input: "hermes3:8b-llama3.1-q3_K_L", expected: "Hermes3 8B Llama3.1 Q3 K_L"},
		{input: "hermes3:8b-llama3.1-q4_0", expected: "Hermes3 8B Llama3.1 Q4_0"},
		{input: "hermes3:8b-llama3.1-q4_1", expected: "Hermes3 8B Llama3.1 Q4 1"},
		{input: "hermes3:8b-llama3.1-q4_K_S", expected: "Hermes3 8B Llama3.1 Q4 K_S"},
		{input: "hermes3:8b-llama3.1-q4_K_M", expected: "Hermes3 8B Llama3.1 Q4_K_M"},
		{input: "hermes3:8b-llama3.1-q5_0", expected: "Hermes3 8B Llama3.1 Q5 0"},
		{input: "hermes3:8b-llama3.1-q5_1", expected: "Hermes3 8B Llama3.1 Q5 1"},
		{input: "hermes3:8b-llama3.1-q5_K_S", expected: "Hermes3 8B Llama3.1 Q5 K_S"},
		{input: "hermes3:8b-llama3.1-q5_K_M", expected: "Hermes3 8B Llama3.1 Q5 K_M"},
		{input: "hermes3:8b-llama3.1-q6_K", expected: "Hermes3 8B Llama3.1 Q6 K"},
		{input: "hermes3:8b-llama3.1-q8_0", expected: "Hermes3 8B Llama3.1 Q8_0"},
		{input: "hermes3:8b-llama3.1-fp16", expected: "Hermes3 8B Llama3.1 FP16"},
		{input: "hermes3:70b-llama3.1-q2_K", expected: "Hermes3 70B Llama3.1 Q2_K"},
		{input: "hermes3:70b-llama3.1-q3_K_S", expected: "Hermes3 70B Llama3.1 Q3 K_S"},
		{input: "hermes3:70b-llama3.1-q3_K_M", expected: "Hermes3 70B Llama3.1 Q3 K_M"},
		{input: "hermes3:70b-llama3.1-q3_K_L", expected: "Hermes3 70B Llama3.1 Q3 K_L"},
		{input: "hermes3:70b-llama3.1-q4_0", expected: "Hermes3 70B Llama3.1 Q4_0"},
		{input: "hermes3:70b-llama3.1-q4_1", expected: "Hermes3 70B Llama3.1 Q4 1"},
		{input: "hermes3:70b-llama3.1-q4_K_S", expected: "Hermes3 70B Llama3.1 Q4 K_S"},
		{input: "hermes3:70b-llama3.1-q4_K_M", expected: "Hermes3 70B Llama3.1 Q4_K_M"},
		{input: "hermes3:70b-llama3.1-q5_0", expected: "Hermes3 70B Llama3.1 Q5 0"},
		{input: "hermes3:70b-llama3.1-q5_1", expected: "Hermes3 70B Llama3.1 Q5 1"},
		{input: "hermes3:70b-llama3.1-q5_K_S", expected: "Hermes3 70B Llama3.1 Q5 K_S"},
		{input: "hermes3:70b-llama3.1-q5_K_M", expected: "Hermes3 70B Llama3.1 Q5 K_M"},
		{input: "hermes3:70b-llama3.1-q6_K", expected: "Hermes3 70B Llama3.1 Q6 K"},
		{input: "hermes3:70b-llama3.1-q8_0", expected: "Hermes3 70B Llama3.1 Q8_0"},
		{input: "hermes3:70b-llama3.1-fp16", expected: "Hermes3 70B Llama3.1 FP16"},
		{input: "hermes3:405b-llama3.1-q2_K", expected: "Hermes3 405B Llama3.1 Q2_K"},
		{input: "hermes3:405b-llama3.1-q3_K_S", expected: "Hermes3 405B Llama3.1 Q3 K_S"},
		{input: "hermes3:405b-llama3.1-q3_K_M", expected: "Hermes3 405B Llama3.1 Q3 K_M"},
		{input: "hermes3:405b-llama3.1-q3_K_L", expected: "Hermes3 405B Llama3.1 Q3 K_L"},
		{input: "hermes3:405b-llama3.1-q4_0", expected: "Hermes3 405B Llama3.1 Q4_0"},
		{input: "hermes3:405b-llama3.1-q4_1", expected: "Hermes3 405B Llama3.1 Q4 1"},
		{input: "hermes3:405b-llama3.1-q4_K_S", expected: "Hermes3 405B Llama3.1 Q4 K_S"},
		{input: "hermes3:405b-llama3.1-q4_K_M", expected: "Hermes3 405B Llama3.1 Q4_K_M"},
		{input: "hermes3:405b-llama3.1-q5_0", expected: "Hermes3 405B Llama3.1 Q5 0"},
		{input: "hermes3:405b-llama3.1-q5_1", expected: "Hermes3 405B Llama3.1 Q5 1"},
		{input: "hermes3:405b-llama3.1-q5_K_S", expected: "Hermes3 405B Llama3.1 Q5 K_S"},
		{input: "hermes3:405b-llama3.1-q5_K_M", expected: "Hermes3 405B Llama3.1 Q5 K_M"},
		{input: "hermes3:405b-llama3.1-q6_K", expected: "Hermes3 405B Llama3.1 Q6 K"},
		{input: "hermes3:405b-llama3.1-q8_0", expected: "Hermes3 405B Llama3.1 Q8_0"},
		{input: "hermes3:405b-llama3.1-fp16", expected: "Hermes3 405B Llama3.1 FP16"},
		{input: "yi:latest", expected: "Yi (latest)"},
		{input: "yi", expected: "Yi"},
		{input: "yi:v1.5", expected: "Yi v1.5"},
		{input: "yi:6b", expected: "Yi 6B"},
		{input: "yi:9b", expected: "Yi 9B"},
		{input: "yi:34b", expected: "Yi 34B"},
		{input: "yi:6b-200k", expected: "Yi 6B 200K"},
		{input: "yi:6b-200k-q2_K", expected: "Yi 6B 200K Q2_K"},
		{input: "yi:6b-200k-q3_K_S", expected: "Yi 6B 200K Q3 K_S"},
		{input: "yi:6b-200k-q3_K_M", expected: "Yi 6B 200K Q3 K_M"},
		{input: "yi:6b-200k-q3_K_L", expected: "Yi 6B 200K Q3 K_L"},
		{input: "yi:6b-200k-q4_0", expected: "Yi 6B 200K Q4_0"},
		{input: "yi:6b-200k-q4_1", expected: "Yi 6B 200K Q4 1"},
		{input: "yi:6b-200k-q4_K_S", expected: "Yi 6B 200K Q4 K_S"},
		{input: "yi:6b-200k-q4_K_M", expected: "Yi 6B 200K Q4_K_M"},
		{input: "yi:6b-200k-q5_0", expected: "Yi 6B 200K Q5 0"},
		{input: "yi:6b-200k-q5_1", expected: "Yi 6B 200K Q5 1"},
		{input: "yi:6b-200k-q5_K_S", expected: "Yi 6B 200K Q5 K_S"},
		{input: "yi:6b-200k-q5_K_M", expected: "Yi 6B 200K Q5 K_M"},
		{input: "yi:6b-200k-q6_K", expected: "Yi 6B 200K Q6 K"},
		{input: "yi:6b-200k-q8_0", expected: "Yi 6B 200K Q8_0"},
		{input: "yi:6b-200k-fp16", expected: "Yi 6B 200K FP16"},
		{input: "yi:6b-chat", expected: "Yi 6B Chat"},
		{input: "yi:6b-chat-q2_K", expected: "Yi 6B Chat Q2_K"},
		{input: "yi:6b-chat-v1.5-q2_K", expected: "Yi 6B Chat v1.5 Q2_K"},
		{input: "yi:6b-chat-q3_K_S", expected: "Yi 6B Chat Q3 K_S"},
		{input: "yi:6b-chat-v1.5-q3_K_S", expected: "Yi 6B Chat v1.5 Q3 K_S"},
		{input: "yi:6b-chat-q3_K_M", expected: "Yi 6B Chat Q3 K_M"},
		{input: "yi:6b-chat-v1.5-q3_K_M", expected: "Yi 6B Chat v1.5 Q3 K_M"},
		{input: "yi:6b-chat-q3_K_L", expected: "Yi 6B Chat Q3 K_L"},
		{input: "yi:6b-chat-v1.5-q3_K_L", expected: "Yi 6B Chat v1.5 Q3 K_L"},
		{input: "yi:6b-chat-q4_0", expected: "Yi 6B Chat Q4_0"},
		{input: "yi:6b-chat-v1.5-q4_0", expected: "Yi 6B Chat v1.5 Q4_0"},
		{input: "yi:6b-chat-q4_1", expected: "Yi 6B Chat Q4 1"},
		{input: "yi:6b-chat-v1.5-q4_1", expected: "Yi 6B Chat v1.5 Q4 1"},
		{input: "yi:6b-chat-q4_K_S", expected: "Yi 6B Chat Q4 K_S"},
		{input: "yi:6b-chat-v1.5-q4_K_S", expected: "Yi 6B Chat v1.5 Q4 K_S"},
		{input: "yi:6b-chat-q4_K_M", expected: "Yi 6B Chat Q4_K_M"},
		{input: "yi:6b-chat-v1.5-q4_K_M", expected: "Yi 6B Chat v1.5 Q4_K_M"},
		{input: "yi:6b-chat-q5_0", expected: "Yi 6B Chat Q5 0"},
		{input: "yi:6b-chat-v1.5-q5_0", expected: "Yi 6B Chat v1.5 Q5 0"},
		{input: "yi:6b-chat-q5_1", expected: "Yi 6B Chat Q5 1"},
		{input: "yi:6b-chat-v1.5-q5_1", expected: "Yi 6B Chat v1.5 Q5 1"},
		{input: "yi:6b-chat-q5_K_S", expected: "Yi 6B Chat Q5 K_S"},
		{input: "yi:6b-chat-v1.5-q5_K_S", expected: "Yi 6B Chat v1.5 Q5 K_S"},
		{input: "yi:6b-chat-q5_K_M", expected: "Yi 6B Chat Q5 K_M"},
		{input: "yi:6b-chat-v1.5-q5_K_M", expected: "Yi 6B Chat v1.5 Q5 K_M"},
		{input: "yi:6b-chat-q6_K", expected: "Yi 6B Chat Q6 K"},
		{input: "yi:6b-chat-v1.5-q6_K", expected: "Yi 6B Chat v1.5 Q6 K"},
		{input: "yi:6b-chat-q8_0", expected: "Yi 6B Chat Q8_0"},
		{input: "yi:6b-chat-v1.5-q8_0", expected: "Yi 6B Chat v1.5 Q8_0"},
		{input: "yi:6b-chat-fp16", expected: "Yi 6B Chat FP16"},
		{input: "yi:6b-chat-v1.5-fp16", expected: "Yi 6B Chat v1.5 FP16"},
		{input: "yi:6b-v1.5", expected: "Yi 6B v1.5"},
		{input: "yi:6b-v1.5-q2_K", expected: "Yi 6B v1.5 Q2_K"},
		{input: "yi:6b-v1.5-q3_K_S", expected: "Yi 6B v1.5 Q3 K_S"},
		{input: "yi:6b-v1.5-q3_K_M", expected: "Yi 6B v1.5 Q3 K_M"},
		{input: "yi:6b-v1.5-q3_K_L", expected: "Yi 6B v1.5 Q3 K_L"},
		{input: "yi:6b-v1.5-q4_0", expected: "Yi 6B v1.5 Q4_0"},
		{input: "yi:6b-v1.5-q4_1", expected: "Yi 6B v1.5 Q4 1"},
		{input: "yi:6b-v1.5-q4_K_S", expected: "Yi 6B v1.5 Q4 K_S"},
		{input: "yi:6b-v1.5-q4_K_M", expected: "Yi 6B v1.5 Q4_K_M"},
		{input: "yi:6b-v1.5-q5_0", expected: "Yi 6B v1.5 Q5 0"},
		{input: "yi:6b-v1.5-q5_1", expected: "Yi 6B v1.5 Q5 1"},
		{input: "yi:6b-v1.5-q5_K_S", expected: "Yi 6B v1.5 Q5 K_S"},
		{input: "yi:6b-v1.5-q5_K_M", expected: "Yi 6B v1.5 Q5 K_M"},
		{input: "yi:6b-v1.5-q6_K", expected: "Yi 6B v1.5 Q6 K"},
		{input: "yi:6b-v1.5-q8_0", expected: "Yi 6B v1.5 Q8_0"},
		{input: "yi:6b-v1.5-fp16", expected: "Yi 6B v1.5 FP16"},
		{input: "yi:6b-q2_K", expected: "Yi 6B Q2_K"},
		{input: "yi:6b-q3_K_S", expected: "Yi 6B Q3 K_S"},
		{input: "yi:6b-q3_K_M", expected: "Yi 6B Q3 K_M"},
		{input: "yi:6b-q3_K_L", expected: "Yi 6B Q3 K_L"},
		{input: "yi:6b-q4_0", expected: "Yi 6B Q4_0"},
		{input: "yi:6b-q4_1", expected: "Yi 6B Q4 1"},
		{input: "yi:6b-q4_K_S", expected: "Yi 6B Q4 K_S"},
		{input: "yi:6b-q4_K_M", expected: "Yi 6B Q4_K_M"},
		{input: "yi:6b-q5_0", expected: "Yi 6B Q5 0"},
		{input: "yi:6b-q5_1", expected: "Yi 6B Q5 1"},
		{input: "yi:6b-q5_K_S", expected: "Yi 6B Q5 K_S"},
		{input: "yi:6b-q5_K_M", expected: "Yi 6B Q5 K_M"},
		{input: "yi:6b-q6_K", expected: "Yi 6B Q6 K"},
		{input: "yi:6b-q8_0", expected: "Yi 6B Q8_0"},
		{input: "yi:6b-fp16", expected: "Yi 6B FP16"},
		{input: "yi:9b-chat", expected: "Yi 9B Chat"},
		{input: "yi:9b-chat-v1.5-q2_K", expected: "Yi 9B Chat v1.5 Q2_K"},
		{input: "yi:9b-chat-v1.5-q3_K_S", expected: "Yi 9B Chat v1.5 Q3 K_S"},
		{input: "yi:9b-chat-v1.5-q3_K_M", expected: "Yi 9B Chat v1.5 Q3 K_M"},
		{input: "yi:9b-chat-v1.5-q3_K_L", expected: "Yi 9B Chat v1.5 Q3 K_L"},
		{input: "yi:9b-chat-v1.5-q4_0", expected: "Yi 9B Chat v1.5 Q4_0"},
		{input: "yi:9b-chat-v1.5-q4_1", expected: "Yi 9B Chat v1.5 Q4 1"},
		{input: "yi:9b-chat-v1.5-q4_K_S", expected: "Yi 9B Chat v1.5 Q4 K_S"},
		{input: "yi:9b-chat-v1.5-q4_K_M", expected: "Yi 9B Chat v1.5 Q4_K_M"},
		{input: "yi:9b-chat-v1.5-q5_0", expected: "Yi 9B Chat v1.5 Q5 0"},
		{input: "yi:9b-chat-v1.5-q5_1", expected: "Yi 9B Chat v1.5 Q5 1"},
		{input: "yi:9b-chat-v1.5-q5_K_S", expected: "Yi 9B Chat v1.5 Q5 K_S"},
		{input: "yi:9b-chat-v1.5-q5_K_M", expected: "Yi 9B Chat v1.5 Q5 K_M"},
		{input: "yi:9b-chat-v1.5-q6_K", expected: "Yi 9B Chat v1.5 Q6 K"},
		{input: "yi:9b-chat-v1.5-q8_0", expected: "Yi 9B Chat v1.5 Q8_0"},
		{input: "yi:9b-chat-v1.5-fp16", expected: "Yi 9B Chat v1.5 FP16"},
		{input: "yi:9b-v1.5", expected: "Yi 9B v1.5"},
		{input: "yi:9b-v1.5-q2_K", expected: "Yi 9B v1.5 Q2_K"},
		{input: "yi:9b-v1.5-q3_K_S", expected: "Yi 9B v1.5 Q3 K_S"},
		{input: "yi:9b-v1.5-q3_K_M", expected: "Yi 9B v1.5 Q3 K_M"},
		{input: "yi:9b-v1.5-q3_K_L", expected: "Yi 9B v1.5 Q3 K_L"},
		{input: "yi:9b-v1.5-q4_0", expected: "Yi 9B v1.5 Q4_0"},
		{input: "yi:9b-v1.5-q4_1", expected: "Yi 9B v1.5 Q4 1"},
		{input: "yi:9b-v1.5-q4_K_S", expected: "Yi 9B v1.5 Q4 K_S"},
		{input: "yi:9b-v1.5-q4_K_M", expected: "Yi 9B v1.5 Q4_K_M"},
		{input: "yi:9b-v1.5-q5_0", expected: "Yi 9B v1.5 Q5 0"},
		{input: "yi:9b-v1.5-q5_1", expected: "Yi 9B v1.5 Q5 1"},
		{input: "yi:9b-v1.5-q5_K_S", expected: "Yi 9B v1.5 Q5 K_S"},
		{input: "yi:9b-v1.5-q5_K_M", expected: "Yi 9B v1.5 Q5 K_M"},
		{input: "yi:9b-v1.5-q6_K", expected: "Yi 9B v1.5 Q6 K"},
		{input: "yi:9b-v1.5-q8_0", expected: "Yi 9B v1.5 Q8_0"},
		{input: "yi:9b-v1.5-fp16", expected: "Yi 9B v1.5 FP16"},
		{input: "yi:34b-chat", expected: "Yi 34B Chat"},
		{input: "yi:34b-chat-q2_K", expected: "Yi 34B Chat Q2_K"},
		{input: "yi:34b-chat-v1.5-q2_K", expected: "Yi 34B Chat v1.5 Q2_K"},
		{input: "yi:34b-chat-q3_K_S", expected: "Yi 34B Chat Q3 K_S"},
		{input: "yi:34b-chat-v1.5-q3_K_S", expected: "Yi 34B Chat v1.5 Q3 K_S"},
		{input: "yi:34b-chat-q3_K_M", expected: "Yi 34B Chat Q3 K_M"},
		{input: "yi:34b-chat-v1.5-q3_K_M", expected: "Yi 34B Chat v1.5 Q3 K_M"},
		{input: "yi:34b-chat-q3_K_L", expected: "Yi 34B Chat Q3 K_L"},
		{input: "yi:34b-chat-v1.5-q3_K_L", expected: "Yi 34B Chat v1.5 Q3 K_L"},
		{input: "yi:34b-chat-q4_0", expected: "Yi 34B Chat Q4_0"},
		{input: "yi:34b-chat-v1.5-q4_0", expected: "Yi 34B Chat v1.5 Q4_0"},
		{input: "yi:34b-chat-q4_1", expected: "Yi 34B Chat Q4 1"},
		{input: "yi:34b-chat-v1.5-q4_1", expected: "Yi 34B Chat v1.5 Q4 1"},
		{input: "yi:34b-chat-q4_K_S", expected: "Yi 34B Chat Q4 K_S"},
		{input: "yi:34b-chat-v1.5-q4_K_S", expected: "Yi 34B Chat v1.5 Q4 K_S"},
		{input: "yi:34b-chat-q4_K_M", expected: "Yi 34B Chat Q4_K_M"},
		{input: "yi:34b-chat-v1.5-q4_K_M", expected: "Yi 34B Chat v1.5 Q4_K_M"},
		{input: "yi:34b-chat-q5_0", expected: "Yi 34B Chat Q5 0"},
		{input: "yi:34b-chat-v1.5-q5_0", expected: "Yi 34B Chat v1.5 Q5 0"},
		{input: "yi:34b-chat-q5_1", expected: "Yi 34B Chat Q5 1"},
		{input: "yi:34b-chat-v1.5-q5_1", expected: "Yi 34B Chat v1.5 Q5 1"},
		{input: "yi:34b-chat-q5_K_S", expected: "Yi 34B Chat Q5 K_S"},
		{input: "yi:34b-chat-v1.5-q5_K_S", expected: "Yi 34B Chat v1.5 Q5 K_S"},
		{input: "yi:34b-chat-q5_K_M", expected: "Yi 34B Chat Q5 K_M"},
		{input: "yi:34b-chat-v1.5-q5_K_M", expected: "Yi 34B Chat v1.5 Q5 K_M"},
		{input: "yi:34b-chat-q6_K", expected: "Yi 34B Chat Q6 K"},
		{input: "yi:34b-chat-v1.5-q6_K", expected: "Yi 34B Chat v1.5 Q6 K"},
		{input: "yi:34b-chat-q8_0", expected: "Yi 34B Chat Q8_0"},
		{input: "yi:34b-chat-v1.5-q8_0", expected: "Yi 34B Chat v1.5 Q8_0"},
		{input: "yi:34b-chat-fp16", expected: "Yi 34B Chat FP16"},
		{input: "yi:34b-chat-v1.5-fp16", expected: "Yi 34B Chat v1.5 FP16"},
		{input: "yi:34b-v1.5", expected: "Yi 34B v1.5"},
		{input: "yi:34b-v1.5-q2_K", expected: "Yi 34B v1.5 Q2_K"},
		{input: "yi:34b-v1.5-q3_K_S", expected: "Yi 34B v1.5 Q3 K_S"},
		{input: "yi:34b-v1.5-q3_K_M", expected: "Yi 34B v1.5 Q3 K_M"},
		{input: "yi:34b-v1.5-q3_K_L", expected: "Yi 34B v1.5 Q3 K_L"},
		{input: "yi:34b-v1.5-q4_0", expected: "Yi 34B v1.5 Q4_0"},
		{input: "yi:34b-v1.5-q4_1", expected: "Yi 34B v1.5 Q4 1"},
		{input: "yi:34b-v1.5-q4_K_S", expected: "Yi 34B v1.5 Q4 K_S"},
		{input: "yi:34b-v1.5-q4_K_M", expected: "Yi 34B v1.5 Q4_K_M"},
		{input: "yi:34b-v1.5-q5_0", expected: "Yi 34B v1.5 Q5 0"},
		{input: "yi:34b-v1.5-q5_1", expected: "Yi 34B v1.5 Q5 1"},
		{input: "yi:34b-v1.5-q5_K_S", expected: "Yi 34B v1.5 Q5 K_S"},
		{input: "yi:34b-v1.5-q5_K_M", expected: "Yi 34B v1.5 Q5 K_M"},
		{input: "yi:34b-v1.5-q6_K", expected: "Yi 34B v1.5 Q6 K"},
		{input: "yi:34b-v1.5-q8_0", expected: "Yi 34B v1.5 Q8_0"},
		{input: "yi:34b-v1.5-fp16", expected: "Yi 34B v1.5 FP16"},
		{input: "yi:34b-q2_K", expected: "Yi 34B Q2_K"},
		{input: "yi:34b-q3_K_S", expected: "Yi 34B Q3 K_S"},
		{input: "yi:34b-q3_K_M", expected: "Yi 34B Q3 K_M"},
		{input: "yi:34b-q3_K_L", expected: "Yi 34B Q3 K_L"},
		{input: "yi:34b-q4_0", expected: "Yi 34B Q4_0"},
		{input: "yi:34b-q4_1", expected: "Yi 34B Q4 1"},
		{input: "yi:34b-q4_K_S", expected: "Yi 34B Q4 K_S"},
		{input: "yi:34b-q4_K_M", expected: "Yi 34B Q4_K_M"},
		{input: "yi:34b-q5_0", expected: "Yi 34B Q5 0"},
		{input: "yi:34b-q5_1", expected: "Yi 34B Q5 1"},
		{input: "yi:34b-q5_K_S", expected: "Yi 34B Q5 K_S"},
		{input: "yi:34b-q6_K", expected: "Yi 34B Q6 K"},
		{input: "embeddinggemma:latest", expected: "Embeddinggemma (latest)"},
		{input: "embeddinggemma", expected: "Embeddinggemma"},
		{input: "embeddinggemma:300m", expected: "Embeddinggemma 300M"},
		{input: "embeddinggemma:300m-qat-q4_0", expected: "Embeddinggemma 300M Qat Q4_0"},
		{input: "embeddinggemma:300m-qat-q8_0", expected: "Embeddinggemma 300M Qat Q8_0"},
		{input: "embeddinggemma:300m-bf16", expected: "Embeddinggemma 300M BF16"},
		{input: "zephyr:latest", expected: "Zephyr (latest)"},
		{input: "zephyr", expected: "Zephyr"},
		{input: "zephyr:7b", expected: "Zephyr 7B"},
		{input: "zephyr:141b", expected: "Zephyr 141B"},
		{input: "zephyr:7b-alpha", expected: "Zephyr 7B Alpha"},
		{input: "zephyr:7b-alpha-q2_K", expected: "Zephyr 7B Alpha Q2_K"},
		{input: "zephyr:7b-alpha-q3_K_S", expected: "Zephyr 7B Alpha Q3 K_S"},
		{input: "zephyr:7b-alpha-q3_K_M", expected: "Zephyr 7B Alpha Q3 K_M"},
		{input: "zephyr:7b-alpha-q3_K_L", expected: "Zephyr 7B Alpha Q3 K_L"},
		{input: "zephyr:7b-alpha-q4_0", expected: "Zephyr 7B Alpha Q4_0"},
		{input: "zephyr:7b-alpha-q4_1", expected: "Zephyr 7B Alpha Q4 1"},
		{input: "zephyr:7b-alpha-q4_K_S", expected: "Zephyr 7B Alpha Q4 K_S"},
		{input: "zephyr:7b-alpha-q4_K_M", expected: "Zephyr 7B Alpha Q4_K_M"},
		{input: "zephyr:7b-alpha-q5_0", expected: "Zephyr 7B Alpha Q5 0"},
		{input: "zephyr:7b-alpha-q5_1", expected: "Zephyr 7B Alpha Q5 1"},
		{input: "zephyr:7b-alpha-q5_K_S", expected: "Zephyr 7B Alpha Q5 K_S"},
		{input: "zephyr:7b-alpha-q5_K_M", expected: "Zephyr 7B Alpha Q5 K_M"},
		{input: "zephyr:7b-alpha-q6_K", expected: "Zephyr 7B Alpha Q6 K"},
		{input: "zephyr:7b-alpha-q8_0", expected: "Zephyr 7B Alpha Q8_0"},
		{input: "zephyr:7b-alpha-fp16", expected: "Zephyr 7B Alpha FP16"},
		{input: "zephyr:7b-beta", expected: "Zephyr 7B Beta"},
		{input: "zephyr:7b-beta-q2_K", expected: "Zephyr 7B Beta Q2_K"},
		{input: "zephyr:7b-beta-q3_K_S", expected: "Zephyr 7B Beta Q3 K_S"},
		{input: "zephyr:7b-beta-q3_K_M", expected: "Zephyr 7B Beta Q3 K_M"},
		{input: "zephyr:7b-beta-q3_K_L", expected: "Zephyr 7B Beta Q3 K_L"},
		{input: "zephyr:7b-beta-q4_0", expected: "Zephyr 7B Beta Q4_0"},
		{input: "zephyr:7b-beta-q4_1", expected: "Zephyr 7B Beta Q4 1"},
		{input: "zephyr:7b-beta-q4_K_S", expected: "Zephyr 7B Beta Q4 K_S"},
		{input: "zephyr:7b-beta-q4_K_M", expected: "Zephyr 7B Beta Q4_K_M"},
		{input: "zephyr:7b-beta-q5_0", expected: "Zephyr 7B Beta Q5 0"},
		{input: "zephyr:7b-beta-q5_1", expected: "Zephyr 7B Beta Q5 1"},
		{input: "zephyr:7b-beta-q5_K_S", expected: "Zephyr 7B Beta Q5 K_S"},
		{input: "zephyr:7b-beta-q5_K_M", expected: "Zephyr 7B Beta Q5 K_M"},
		{input: "zephyr:7b-beta-q6_K", expected: "Zephyr 7B Beta Q6 K"},
		{input: "zephyr:7b-beta-q8_0", expected: "Zephyr 7B Beta Q8_0"},
		{input: "zephyr:7b-beta-fp16", expected: "Zephyr 7B Beta FP16"},
		{input: "zephyr:141b-v0.1", expected: "Zephyr 141B v0.1"},
		{input: "zephyr:141b-v0.1-q2_K", expected: "Zephyr 141B v0.1 Q2_K"},
		{input: "zephyr:141b-v0.1-q4_0", expected: "Zephyr 141B v0.1 Q4_0"},
		{input: "zephyr:141b-v0.1-q8_0", expected: "Zephyr 141B v0.1 Q8_0"},
		{input: "zephyr:141b-v0.1-fp16", expected: "Zephyr 141B v0.1 FP16"},
		{input: "exaone-deep:latest", expected: "Exaone Deep (latest)"},
		{input: "exaone-deep", expected: "Exaone Deep"},
		{input: "exaone-deep:2.4b", expected: "Exaone Deep 2.4B"},
		{input: "exaone-deep:7.8b", expected: "Exaone Deep 7.8B"},
		{input: "exaone-deep:32b", expected: "Exaone Deep 32B"},
		{input: "exaone-deep:2.4b-q4_K_M", expected: "Exaone Deep 2.4B Q4_K_M"},
		{input: "exaone-deep:2.4b-q8_0", expected: "Exaone Deep 2.4B Q8_0"},
		{input: "exaone-deep:2.4b-fp16", expected: "Exaone Deep 2.4B FP16"},
		{input: "exaone-deep:7.8b-q4_K_M", expected: "Exaone Deep 7.8B Q4_K_M"},
		{input: "exaone-deep:7.8b-q8_0", expected: "Exaone Deep 7.8B Q8_0"},
		{input: "exaone-deep:7.8b-fp16", expected: "Exaone Deep 7.8B FP16"},
		{input: "exaone-deep:32b-q4_K_M", expected: "Exaone Deep 32B Q4_K_M"},
		{input: "exaone-deep:32b-q8_0", expected: "Exaone Deep 32B Q8_0"},
		{input: "exaone-deep:32b-fp16", expected: "Exaone Deep 32B FP16"},
		{input: "granite4:latest", expected: "Granite4 (latest)"},
		{input: "granite4", expected: "Granite4"},
		{input: "granite4:micro", expected: "Granite4 Micro"},
		{input: "granite4:350m", expected: "Granite4 350M"},
		{input: "granite4:1b", expected: "Granite4 1B"},
		{input: "granite4:3b", expected: "Granite4 3B"},
		{input: "granite4:350m-h", expected: "Granite4 350M H"},
		{input: "granite4:350m-h-q8_0", expected: "Granite4 350M H Q8_0"},
		{input: "granite4:350m-bf16", expected: "Granite4 350M BF16"},
		{input: "granite4:1b-h", expected: "Granite4 1B H"},
		{input: "granite4:1b-h-q8_0", expected: "Granite4 1B H Q8_0"},
		{input: "granite4:1b-bf16", expected: "Granite4 1B BF16"},
		{input: "granite4:3b-h", expected: "Granite4 3B H"},
		{input: "granite4:7b-a1b-h", expected: "Granite4 7B A1B H"},
		{input: "granite4:32b-a9b-h", expected: "Granite4 32B A9B H"},
		{input: "granite4:micro-h", expected: "Granite4 Micro H"},
		{input: "granite4:small-h", expected: "Granite4 Small H"},
		{input: "granite4:tiny-h", expected: "Granite4 Tiny H"},
		{input: "mistral-large:latest", expected: "Mistral Large (latest)"},
		{input: "mistral-large", expected: "Mistral Large"},
		{input: "mistral-large:123b", expected: "Mistral Large 123B"},
		{input: "mistral-large:123b-instruct-2407-q2_K", expected: "Mistral Large 123B Instruct 2407 Q2_K"},
		{input: "mistral-large:123b-instruct-2407-q3_K_S", expected: "Mistral Large 123B Instruct 2407 Q3 K_S"},
		{input: "mistral-large:123b-instruct-2407-q3_K_M", expected: "Mistral Large 123B Instruct 2407 Q3 K_M"},
		{input: "mistral-large:123b-instruct-2407-q3_K_L", expected: "Mistral Large 123B Instruct 2407 Q3 K_L"},
		{input: "mistral-large:123b-instruct-2407-q4_0", expected: "Mistral Large 123B Instruct 2407 Q4_0"},
		{input: "mistral-large:123b-instruct-2407-q4_1", expected: "Mistral Large 123B Instruct 2407 Q4 1"},
		{input: "mistral-large:123b-instruct-2407-q4_K_S", expected: "Mistral Large 123B Instruct 2407 Q4 K_S"},
		{input: "mistral-large:123b-instruct-2407-q4_K_M", expected: "Mistral Large 123B Instruct 2407 Q4_K_M"},
		{input: "mistral-large:123b-instruct-2407-q5_0", expected: "Mistral Large 123B Instruct 2407 Q5 0"},
		{input: "mistral-large:123b-instruct-2407-q5_1", expected: "Mistral Large 123B Instruct 2407 Q5 1"},
		{input: "mistral-large:123b-instruct-2407-q5_K_S", expected: "Mistral Large 123B Instruct 2407 Q5 K_S"},
		{input: "mistral-large:123b-instruct-2407-q5_K_M", expected: "Mistral Large 123B Instruct 2407 Q5 K_M"},
		{input: "mistral-large:123b-instruct-2407-q6_K", expected: "Mistral Large 123B Instruct 2407 Q6 K"},
		{input: "mistral-large:123b-instruct-2407-q8_0", expected: "Mistral Large 123B Instruct 2407 Q8_0"},
		{input: "mistral-large:123b-instruct-2407-fp16", expected: "Mistral Large 123B Instruct 2407 FP16"},
		{input: "mistral-large:123b-instruct-2411-q2_K", expected: "Mistral Large 123B Instruct 2411 Q2_K"},
		{input: "mistral-large:123b-instruct-2411-q3_K_S", expected: "Mistral Large 123B Instruct 2411 Q3 K_S"},
		{input: "mistral-large:123b-instruct-2411-q3_K_M", expected: "Mistral Large 123B Instruct 2411 Q3 K_M"},
		{input: "mistral-large:123b-instruct-2411-q3_K_L", expected: "Mistral Large 123B Instruct 2411 Q3 K_L"},
		{input: "mistral-large:123b-instruct-2411-q4_0", expected: "Mistral Large 123B Instruct 2411 Q4_0"},
		{input: "mistral-large:123b-instruct-2411-q4_1", expected: "Mistral Large 123B Instruct 2411 Q4 1"},
		{input: "mistral-large:123b-instruct-2411-q4_K_S", expected: "Mistral Large 123B Instruct 2411 Q4 K_S"},
		{input: "mistral-large:123b-instruct-2411-q4_K_M", expected: "Mistral Large 123B Instruct 2411 Q4_K_M"},
		{input: "mistral-large:123b-instruct-2411-q5_0", expected: "Mistral Large 123B Instruct 2411 Q5 0"},
		{input: "mistral-large:123b-instruct-2411-q5_1", expected: "Mistral Large 123B Instruct 2411 Q5 1"},
		{input: "mistral-large:123b-instruct-2411-q5_K_S", expected: "Mistral Large 123B Instruct 2411 Q5 K_S"},
		{input: "mistral-large:123b-instruct-2411-q5_K_M", expected: "Mistral Large 123B Instruct 2411 Q5 K_M"},
		{input: "mistral-large:123b-instruct-2411-q6_K", expected: "Mistral Large 123B Instruct 2411 Q6 K"},
		{input: "mistral-large:123b-instruct-2411-q8_0", expected: "Mistral Large 123B Instruct 2411 Q8_0"},
		{input: "mistral-large:123b-instruct-2411-fp16", expected: "Mistral Large 123B Instruct 2411 FP16"},
		{input: "wizard-vicuna-uncensored:latest", expected: "Wizard Vicuna Uncensored (latest)"},
		{input: "wizard-vicuna-uncensored", expected: "Wizard Vicuna Uncensored"},
		{input: "wizard-vicuna-uncensored:7b", expected: "Wizard Vicuna Uncensored 7B"},
		{input: "wizard-vicuna-uncensored:13b", expected: "Wizard Vicuna Uncensored 13B"},
		{input: "wizard-vicuna-uncensored:30b", expected: "Wizard Vicuna Uncensored 30B"},
		{input: "wizard-vicuna-uncensored:7b-q2_K", expected: "Wizard Vicuna Uncensored 7B Q2_K"},
		{input: "wizard-vicuna-uncensored:7b-q3_K_S", expected: "Wizard Vicuna Uncensored 7B Q3 K_S"},
		{input: "wizard-vicuna-uncensored:7b-q3_K_M", expected: "Wizard Vicuna Uncensored 7B Q3 K_M"},
		{input: "wizard-vicuna-uncensored:7b-q3_K_L", expected: "Wizard Vicuna Uncensored 7B Q3 K_L"},
		{input: "wizard-vicuna-uncensored:7b-q4_0", expected: "Wizard Vicuna Uncensored 7B Q4_0"},
		{input: "wizard-vicuna-uncensored:7b-q4_1", expected: "Wizard Vicuna Uncensored 7B Q4 1"},
		{input: "wizard-vicuna-uncensored:7b-q4_K_S", expected: "Wizard Vicuna Uncensored 7B Q4 K_S"},
		{input: "wizard-vicuna-uncensored:7b-q4_K_M", expected: "Wizard Vicuna Uncensored 7B Q4_K_M"},
		{input: "wizard-vicuna-uncensored:7b-q5_0", expected: "Wizard Vicuna Uncensored 7B Q5 0"},
		{input: "wizard-vicuna-uncensored:7b-q5_1", expected: "Wizard Vicuna Uncensored 7B Q5 1"},
		{input: "wizard-vicuna-uncensored:7b-q5_K_S", expected: "Wizard Vicuna Uncensored 7B Q5 K_S"},
		{input: "wizard-vicuna-uncensored:7b-q5_K_M", expected: "Wizard Vicuna Uncensored 7B Q5 K_M"},
		{input: "wizard-vicuna-uncensored:7b-q6_K", expected: "Wizard Vicuna Uncensored 7B Q6 K"},
		{input: "wizard-vicuna-uncensored:7b-q8_0", expected: "Wizard Vicuna Uncensored 7B Q8_0"},
		{input: "wizard-vicuna-uncensored:7b-fp16", expected: "Wizard Vicuna Uncensored 7B FP16"},
		{input: "wizard-vicuna-uncensored:13b-q2_K", expected: "Wizard Vicuna Uncensored 13B Q2_K"},
		{input: "wizard-vicuna-uncensored:13b-q3_K_S", expected: "Wizard Vicuna Uncensored 13B Q3 K_S"},
		{input: "wizard-vicuna-uncensored:13b-q3_K_M", expected: "Wizard Vicuna Uncensored 13B Q3 K_M"},
		{input: "wizard-vicuna-uncensored:13b-q3_K_L", expected: "Wizard Vicuna Uncensored 13B Q3 K_L"},
		{input: "wizard-vicuna-uncensored:13b-q4_0", expected: "Wizard Vicuna Uncensored 13B Q4_0"},
		{input: "wizard-vicuna-uncensored:13b-q4_1", expected: "Wizard Vicuna Uncensored 13B Q4 1"},
		{input: "wizard-vicuna-uncensored:13b-q4_K_S", expected: "Wizard Vicuna Uncensored 13B Q4 K_S"},
		{input: "wizard-vicuna-uncensored:13b-q4_K_M", expected: "Wizard Vicuna Uncensored 13B Q4_K_M"},
		{input: "wizard-vicuna-uncensored:13b-q5_0", expected: "Wizard Vicuna Uncensored 13B Q5 0"},
		{input: "wizard-vicuna-uncensored:13b-q5_1", expected: "Wizard Vicuna Uncensored 13B Q5 1"},
		{input: "wizard-vicuna-uncensored:13b-q5_K_S", expected: "Wizard Vicuna Uncensored 13B Q5 K_S"},
		{input: "wizard-vicuna-uncensored:13b-q5_K_M", expected: "Wizard Vicuna Uncensored 13B Q5 K_M"},
		{input: "wizard-vicuna-uncensored:13b-q6_K", expected: "Wizard Vicuna Uncensored 13B Q6 K"},
		{input: "wizard-vicuna-uncensored:13b-q8_0", expected: "Wizard Vicuna Uncensored 13B Q8_0"},
		{input: "wizard-vicuna-uncensored:13b-fp16", expected: "Wizard Vicuna Uncensored 13B FP16"},
		{input: "wizard-vicuna-uncensored:30b-q2_K", expected: "Wizard Vicuna Uncensored 30B Q2_K"},
		{input: "wizard-vicuna-uncensored:30b-q3_K_S", expected: "Wizard Vicuna Uncensored 30B Q3 K_S"},
		{input: "wizard-vicuna-uncensored:30b-q3_K_M", expected: "Wizard Vicuna Uncensored 30B Q3 K_M"},
		{input: "wizard-vicuna-uncensored:30b-q3_K_L", expected: "Wizard Vicuna Uncensored 30B Q3 K_L"},
		{input: "wizard-vicuna-uncensored:30b-q4_0", expected: "Wizard Vicuna Uncensored 30B Q4_0"},
		{input: "wizard-vicuna-uncensored:30b-q4_1", expected: "Wizard Vicuna Uncensored 30B Q4 1"},
		{input: "wizard-vicuna-uncensored:30b-q4_K_S", expected: "Wizard Vicuna Uncensored 30B Q4 K_S"},
		{input: "wizard-vicuna-uncensored:30b-q4_K_M", expected: "Wizard Vicuna Uncensored 30B Q4_K_M"},
		{input: "wizard-vicuna-uncensored:30b-q5_0", expected: "Wizard Vicuna Uncensored 30B Q5 0"},
		{input: "wizard-vicuna-uncensored:30b-q5_1", expected: "Wizard Vicuna Uncensored 30B Q5 1"},
		{input: "wizard-vicuna-uncensored:30b-q5_K_S", expected: "Wizard Vicuna Uncensored 30B Q5 K_S"},
		{input: "wizard-vicuna-uncensored:30b-q5_K_M", expected: "Wizard Vicuna Uncensored 30B Q5 K_M"},
		{input: "wizard-vicuna-uncensored:30b-q6_K", expected: "Wizard Vicuna Uncensored 30B Q6 K"},
		{input: "wizard-vicuna-uncensored:30b-q8_0", expected: "Wizard Vicuna Uncensored 30B Q8_0"},
		{input: "wizard-vicuna-uncensored:30b-fp16", expected: "Wizard Vicuna Uncensored 30B FP16"},
		{input: "opencoder:latest", expected: "Opencoder (latest)"},
		{input: "opencoder", expected: "Opencoder"},
		{input: "opencoder:1.5b", expected: "Opencoder 1.5B"},
		{input: "opencoder:8b", expected: "Opencoder 8B"},
		{input: "opencoder:1.5b-instruct-q4_K_M", expected: "Opencoder 1.5B Instruct Q4_K_M"},
		{input: "opencoder:1.5b-instruct-q8_0", expected: "Opencoder 1.5B Instruct Q8_0"},
		{input: "opencoder:1.5b-instruct-fp16", expected: "Opencoder 1.5B Instruct FP16"},
		{input: "opencoder:8b-instruct-q4_K_M", expected: "Opencoder 8B Instruct Q4_K_M"},
		{input: "opencoder:8b-instruct-q8_0", expected: "Opencoder 8B Instruct Q8_0"},
		{input: "opencoder:8b-instruct-fp16", expected: "Opencoder 8B Instruct FP16"},
		{input: "starcoder:latest", expected: "Starcoder (latest)"},
		{input: "starcoder", expected: "Starcoder"},
		{input: "starcoder:1b", expected: "Starcoder 1B"},
		{input: "starcoder:3b", expected: "Starcoder 3B"},
		{input: "starcoder:7b", expected: "Starcoder 7B"},
		{input: "starcoder:15b", expected: "Starcoder 15B"},
		{input: "starcoder:1b-base", expected: "Starcoder 1B Base"},
		{input: "starcoder:1b-base-q2_K", expected: "Starcoder 1B Base Q2_K"},
		{input: "starcoder:1b-base-q3_K_S", expected: "Starcoder 1B Base Q3 K_S"},
		{input: "starcoder:1b-base-q3_K_M", expected: "Starcoder 1B Base Q3 K_M"},
		{input: "starcoder:1b-base-q3_K_L", expected: "Starcoder 1B Base Q3 K_L"},
		{input: "starcoder:1b-base-q4_0", expected: "Starcoder 1B Base Q4_0"},
		{input: "starcoder:1b-base-q4_1", expected: "Starcoder 1B Base Q4 1"},
		{input: "starcoder:1b-base-q4_K_S", expected: "Starcoder 1B Base Q4 K_S"},
		{input: "starcoder:1b-base-q4_K_M", expected: "Starcoder 1B Base Q4_K_M"},
		{input: "starcoder:1b-base-q5_0", expected: "Starcoder 1B Base Q5 0"},
		{input: "starcoder:1b-base-q5_1", expected: "Starcoder 1B Base Q5 1"},
		{input: "starcoder:1b-base-q5_K_S", expected: "Starcoder 1B Base Q5 K_S"},
		{input: "starcoder:1b-base-q5_K_M", expected: "Starcoder 1B Base Q5 K_M"},
		{input: "starcoder:1b-base-q6_K", expected: "Starcoder 1B Base Q6 K"},
		{input: "starcoder:1b-base-q8_0", expected: "Starcoder 1B Base Q8_0"},
		{input: "starcoder:1b-base-fp16", expected: "Starcoder 1B Base FP16"},
		{input: "starcoder:3b-base", expected: "Starcoder 3B Base"},
		{input: "starcoder:3b-base-q2_K", expected: "Starcoder 3B Base Q2_K"},
		{input: "starcoder:3b-base-q3_K_S", expected: "Starcoder 3B Base Q3 K_S"},
		{input: "starcoder:3b-base-q3_K_M", expected: "Starcoder 3B Base Q3 K_M"},
		{input: "starcoder:3b-base-q3_K_L", expected: "Starcoder 3B Base Q3 K_L"},
		{input: "starcoder:3b-base-q4_0", expected: "Starcoder 3B Base Q4_0"},
		{input: "starcoder:3b-base-q4_1", expected: "Starcoder 3B Base Q4 1"},
		{input: "starcoder:3b-base-q4_K_S", expected: "Starcoder 3B Base Q4 K_S"},
		{input: "starcoder:3b-base-q4_K_M", expected: "Starcoder 3B Base Q4_K_M"},
		{input: "starcoder:3b-base-q5_0", expected: "Starcoder 3B Base Q5 0"},
		{input: "starcoder:3b-base-q5_1", expected: "Starcoder 3B Base Q5 1"},
		{input: "starcoder:3b-base-q5_K_S", expected: "Starcoder 3B Base Q5 K_S"},
		{input: "starcoder:3b-base-q5_K_M", expected: "Starcoder 3B Base Q5 K_M"},
		{input: "starcoder:3b-base-q6_K", expected: "Starcoder 3B Base Q6 K"},
		{input: "starcoder:3b-base-q8_0", expected: "Starcoder 3B Base Q8_0"},
		{input: "starcoder:3b-base-fp16", expected: "Starcoder 3B Base FP16"},
		{input: "starcoder:7b-base", expected: "Starcoder 7B Base"},
		{input: "starcoder:7b-base-q2_K", expected: "Starcoder 7B Base Q2_K"},
		{input: "starcoder:7b-base-q3_K_S", expected: "Starcoder 7B Base Q3 K_S"},
		{input: "starcoder:7b-base-q3_K_M", expected: "Starcoder 7B Base Q3 K_M"},
		{input: "starcoder:7b-base-q3_K_L", expected: "Starcoder 7B Base Q3 K_L"},
		{input: "starcoder:7b-base-q4_0", expected: "Starcoder 7B Base Q4_0"},
		{input: "starcoder:7b-base-q4_1", expected: "Starcoder 7B Base Q4 1"},
		{input: "starcoder:7b-base-q4_K_S", expected: "Starcoder 7B Base Q4 K_S"},
		{input: "starcoder:7b-base-q4_K_M", expected: "Starcoder 7B Base Q4_K_M"},
		{input: "starcoder:7b-base-q5_0", expected: "Starcoder 7B Base Q5 0"},
		{input: "starcoder:7b-base-q5_1", expected: "Starcoder 7B Base Q5 1"},
		{input: "starcoder:7b-base-q5_K_S", expected: "Starcoder 7B Base Q5 K_S"},
		{input: "starcoder:7b-base-q5_K_M", expected: "Starcoder 7B Base Q5 K_M"},
		{input: "starcoder:7b-base-q6_K", expected: "Starcoder 7B Base Q6 K"},
		{input: "starcoder:7b-base-q8_0", expected: "Starcoder 7B Base Q8_0"},
		{input: "starcoder:7b-base-fp16", expected: "Starcoder 7B Base FP16"},
		{input: "starcoder:15b-base", expected: "Starcoder 15B Base"},
		{input: "starcoder:15b-base-q2_K", expected: "Starcoder 15B Base Q2_K"},
		{input: "starcoder:15b-base-q3_K_S", expected: "Starcoder 15B Base Q3 K_S"},
		{input: "starcoder:15b-base-q3_K_M", expected: "Starcoder 15B Base Q3 K_M"},
		{input: "starcoder:15b-base-q3_K_L", expected: "Starcoder 15B Base Q3 K_L"},
		{input: "starcoder:15b-base-q4_0", expected: "Starcoder 15B Base Q4_0"},
		{input: "starcoder:15b-base-q4_1", expected: "Starcoder 15B Base Q4 1"},
		{input: "starcoder:15b-base-q4_K_S", expected: "Starcoder 15B Base Q4 K_S"},
		{input: "starcoder:15b-base-q4_K_M", expected: "Starcoder 15B Base Q4_K_M"},
		{input: "starcoder:15b-base-q5_0", expected: "Starcoder 15B Base Q5 0"},
		{input: "starcoder:15b-base-q5_1", expected: "Starcoder 15B Base Q5 1"},
		{input: "starcoder:15b-base-q5_K_S", expected: "Starcoder 15B Base Q5 K_S"},
		{input: "starcoder:15b-base-q5_K_M", expected: "Starcoder 15B Base Q5 K_M"},
		{input: "starcoder:15b-base-q6_K", expected: "Starcoder 15B Base Q6 K"},
		{input: "starcoder:15b-base-q8_0", expected: "Starcoder 15B Base Q8_0"},
		{input: "starcoder:15b-base-fp16", expected: "Starcoder 15B Base FP16"},
		{input: "starcoder:15b-plus", expected: "Starcoder 15B Plus"},
		{input: "starcoder:15b-plus-q2_K", expected: "Starcoder 15B Plus Q2_K"},
		{input: "starcoder:15b-plus-q3_K_S", expected: "Starcoder 15B Plus Q3 K_S"},
		{input: "starcoder:15b-plus-q3_K_M", expected: "Starcoder 15B Plus Q3 K_M"},
		{input: "starcoder:15b-plus-q3_K_L", expected: "Starcoder 15B Plus Q3 K_L"},
		{input: "starcoder:15b-plus-q4_0", expected: "Starcoder 15B Plus Q4_0"},
		{input: "starcoder:15b-plus-q4_1", expected: "Starcoder 15B Plus Q4 1"},
		{input: "starcoder:15b-plus-q4_K_S", expected: "Starcoder 15B Plus Q4 K_S"},
		{input: "starcoder:15b-plus-q4_K_M", expected: "Starcoder 15B Plus Q4_K_M"},
		{input: "starcoder:15b-plus-q5_0", expected: "Starcoder 15B Plus Q5 0"},
		{input: "starcoder:15b-plus-q5_1", expected: "Starcoder 15B Plus Q5 1"},
		{input: "starcoder:15b-plus-q5_K_S", expected: "Starcoder 15B Plus Q5 K_S"},
		{input: "starcoder:15b-plus-q5_K_M", expected: "Starcoder 15B Plus Q5 K_M"},
		{input: "starcoder:15b-plus-q6_K", expected: "Starcoder 15B Plus Q6 K"},
		{input: "starcoder:15b-plus-q8_0", expected: "Starcoder 15B Plus Q8_0"},
		{input: "starcoder:15b-plus-fp16", expected: "Starcoder 15B Plus FP16"},
		{input: "starcoder:15b-q2_K", expected: "Starcoder 15B Q2_K"},
		{input: "starcoder:15b-q3_K_S", expected: "Starcoder 15B Q3 K_S"},
		{input: "starcoder:15b-q3_K_M", expected: "Starcoder 15B Q3 K_M"},
		{input: "starcoder:15b-q3_K_L", expected: "Starcoder 15B Q3 K_L"},
		{input: "starcoder:15b-q4_0", expected: "Starcoder 15B Q4_0"},
		{input: "starcoder:15b-q4_1", expected: "Starcoder 15B Q4 1"},
		{input: "starcoder:15b-q4_K_S", expected: "Starcoder 15B Q4 K_S"},
		{input: "starcoder:15b-q4_K_M", expected: "Starcoder 15B Q4_K_M"},
		{input: "starcoder:15b-q5_0", expected: "Starcoder 15B Q5 0"},
		{input: "starcoder:15b-q5_1", expected: "Starcoder 15B Q5 1"},
		{input: "starcoder:15b-q5_K_S", expected: "Starcoder 15B Q5 K_S"},
		{input: "starcoder:15b-q5_K_M", expected: "Starcoder 15B Q5 K_M"},
		{input: "starcoder:15b-q6_K", expected: "Starcoder 15B Q6 K"},
		{input: "starcoder:15b-q8_0", expected: "Starcoder 15B Q8_0"},
		{input: "starcoder:15b-fp16", expected: "Starcoder 15B FP16"},
		{input: "nous-hermes:latest", expected: "Nous Hermes (latest)"},
		{input: "nous-hermes", expected: "Nous Hermes"},
		{input: "nous-hermes:7b", expected: "Nous Hermes 7B"},
		{input: "nous-hermes:13b", expected: "Nous Hermes 13B"},
		{input: "nous-hermes:7b-llama2", expected: "Nous Hermes 7B Llama2"},
		{input: "nous-hermes:7b-llama2-q2_K", expected: "Nous Hermes 7B Llama2 Q2_K"},
		{input: "nous-hermes:7b-llama2-q3_K_S", expected: "Nous Hermes 7B Llama2 Q3 K_S"},
		{input: "nous-hermes:7b-llama2-q3_K_M", expected: "Nous Hermes 7B Llama2 Q3 K_M"},
		{input: "nous-hermes:7b-llama2-q3_K_L", expected: "Nous Hermes 7B Llama2 Q3 K_L"},
		{input: "nous-hermes:7b-llama2-q4_0", expected: "Nous Hermes 7B Llama2 Q4_0"},
		{input: "nous-hermes:7b-llama2-q4_1", expected: "Nous Hermes 7B Llama2 Q4 1"},
		{input: "nous-hermes:7b-llama2-q4_K_S", expected: "Nous Hermes 7B Llama2 Q4 K_S"},
		{input: "nous-hermes:7b-llama2-q4_K_M", expected: "Nous Hermes 7B Llama2 Q4_K_M"},
		{input: "nous-hermes:7b-llama2-q5_0", expected: "Nous Hermes 7B Llama2 Q5 0"},
		{input: "nous-hermes:7b-llama2-q5_1", expected: "Nous Hermes 7B Llama2 Q5 1"},
		{input: "nous-hermes:7b-llama2-q5_K_S", expected: "Nous Hermes 7B Llama2 Q5 K_S"},
		{input: "nous-hermes:7b-llama2-q5_K_M", expected: "Nous Hermes 7B Llama2 Q5 K_M"},
		{input: "nous-hermes:7b-llama2-q6_K", expected: "Nous Hermes 7B Llama2 Q6 K"},
		{input: "nous-hermes:7b-llama2-q8_0", expected: "Nous Hermes 7B Llama2 Q8_0"},
		{input: "nous-hermes:7b-llama2-fp16", expected: "Nous Hermes 7B Llama2 FP16"},
		{input: "nous-hermes:13b-llama2", expected: "Nous Hermes 13B Llama2"},
		{input: "nous-hermes:13b-llama2-q2_K", expected: "Nous Hermes 13B Llama2 Q2_K"},
		{input: "nous-hermes:13b-llama2-q3_K_S", expected: "Nous Hermes 13B Llama2 Q3 K_S"},
		{input: "nous-hermes:13b-llama2-q3_K_M", expected: "Nous Hermes 13B Llama2 Q3 K_M"},
		{input: "nous-hermes:13b-llama2-q3_K_L", expected: "Nous Hermes 13B Llama2 Q3 K_L"},
		{input: "nous-hermes:13b-llama2-q4_0", expected: "Nous Hermes 13B Llama2 Q4_0"},
		{input: "nous-hermes:13b-llama2-q4_1", expected: "Nous Hermes 13B Llama2 Q4 1"},
		{input: "nous-hermes:13b-llama2-q4_K_S", expected: "Nous Hermes 13B Llama2 Q4 K_S"},
		{input: "nous-hermes:13b-llama2-q4_K_M", expected: "Nous Hermes 13B Llama2 Q4_K_M"},
		{input: "nous-hermes:13b-llama2-q5_0", expected: "Nous Hermes 13B Llama2 Q5 0"},
		{input: "nous-hermes:13b-llama2-q5_1", expected: "Nous Hermes 13B Llama2 Q5 1"},
		{input: "nous-hermes:13b-llama2-q5_K_S", expected: "Nous Hermes 13B Llama2 Q5 K_S"},
		{input: "nous-hermes:13b-llama2-q5_K_M", expected: "Nous Hermes 13B Llama2 Q5 K_M"},
		{input: "nous-hermes:13b-llama2-q6_K", expected: "Nous Hermes 13B Llama2 Q6 K"},
		{input: "nous-hermes:13b-llama2-q8_0", expected: "Nous Hermes 13B Llama2 Q8_0"},
		{input: "nous-hermes:13b-llama2-fp16", expected: "Nous Hermes 13B Llama2 FP16"},
		{input: "nous-hermes:13b-q2_K", expected: "Nous Hermes 13B Q2_K"},
		{input: "nous-hermes:13b-q3_K_S", expected: "Nous Hermes 13B Q3 K_S"},
		{input: "nous-hermes:13b-q3_K_M", expected: "Nous Hermes 13B Q3 K_M"},
		{input: "nous-hermes:13b-q3_K_L", expected: "Nous Hermes 13B Q3 K_L"},
		{input: "nous-hermes:13b-q4_0", expected: "Nous Hermes 13B Q4_0"},
		{input: "nous-hermes:13b-q4_1", expected: "Nous Hermes 13B Q4 1"},
		{input: "nous-hermes:13b-q4_K_S", expected: "Nous Hermes 13B Q4 K_S"},
		{input: "nous-hermes:13b-q4_K_M", expected: "Nous Hermes 13B Q4_K_M"},
		{input: "nous-hermes:13b-q5_0", expected: "Nous Hermes 13B Q5 0"},
		{input: "nous-hermes:13b-q5_1", expected: "Nous Hermes 13B Q5 1"},
		{input: "nous-hermes:13b-q5_K_S", expected: "Nous Hermes 13B Q5 K_S"},
		{input: "nous-hermes:13b-q5_K_M", expected: "Nous Hermes 13B Q5 K_M"},
		{input: "nous-hermes:13b-q6_K", expected: "Nous Hermes 13B Q6 K"},
		{input: "nous-hermes:13b-q8_0", expected: "Nous Hermes 13B Q8_0"},
		{input: "nous-hermes:13b-fp16", expected: "Nous Hermes 13B FP16"},
		{input: "nous-hermes:70b-llama2-q2_K", expected: "Nous Hermes 70B Llama2 Q2_K"},
		{input: "nous-hermes:70b-llama2-q3_K_S", expected: "Nous Hermes 70B Llama2 Q3 K_S"},
		{input: "nous-hermes:70b-llama2-q3_K_M", expected: "Nous Hermes 70B Llama2 Q3 K_M"},
		{input: "nous-hermes:70b-llama2-q3_K_L", expected: "Nous Hermes 70B Llama2 Q3 K_L"},
		{input: "nous-hermes:70b-llama2-q4_0", expected: "Nous Hermes 70B Llama2 Q4_0"},
		{input: "nous-hermes:70b-llama2-q4_1", expected: "Nous Hermes 70B Llama2 Q4 1"},
		{input: "nous-hermes:70b-llama2-q4_K_S", expected: "Nous Hermes 70B Llama2 Q4 K_S"},
		{input: "nous-hermes:70b-llama2-q4_K_M", expected: "Nous Hermes 70B Llama2 Q4_K_M"},
		{input: "nous-hermes:70b-llama2-q5_0", expected: "Nous Hermes 70B Llama2 Q5 0"},
		{input: "nous-hermes:70b-llama2-q5_1", expected: "Nous Hermes 70B Llama2 Q5 1"},
		{input: "nous-hermes:70b-llama2-q5_K_M", expected: "Nous Hermes 70B Llama2 Q5 K_M"},
		{input: "nous-hermes:70b-llama2-q6_K", expected: "Nous Hermes 70B Llama2 Q6 K"},
		{input: "nous-hermes:70b-llama2-fp16", expected: "Nous Hermes 70B Llama2 FP16"},
		{input: "falcon:latest", expected: "Falcon (latest)"},
		{input: "falcon", expected: "Falcon"},
		{input: "falcon:instruct", expected: "Falcon Instruct"},
		{input: "falcon:text", expected: "Falcon Text"},
		{input: "falcon:7b", expected: "Falcon 7B"},
		{input: "falcon:40b", expected: "Falcon 40B"},
		{input: "falcon:180b", expected: "Falcon 180B"},
		{input: "falcon:7b-instruct", expected: "Falcon 7B Instruct"},
		{input: "falcon:7b-instruct-q4_0", expected: "Falcon 7B Instruct Q4_0"},
		{input: "falcon:7b-instruct-q4_1", expected: "Falcon 7B Instruct Q4 1"},
		{input: "falcon:7b-instruct-q5_0", expected: "Falcon 7B Instruct Q5 0"},
		{input: "falcon:7b-instruct-q5_1", expected: "Falcon 7B Instruct Q5 1"},
		{input: "falcon:7b-instruct-q8_0", expected: "Falcon 7B Instruct Q8_0"},
		{input: "falcon:7b-instruct-fp16", expected: "Falcon 7B Instruct FP16"},
		{input: "falcon:7b-text", expected: "Falcon 7B Text"},
		{input: "falcon:7b-text-q4_0", expected: "Falcon 7B Text Q4_0"},
		{input: "falcon:7b-text-q4_1", expected: "Falcon 7B Text Q4 1"},
		{input: "falcon:7b-text-q5_0", expected: "Falcon 7B Text Q5 0"},
		{input: "falcon:7b-text-q5_1", expected: "Falcon 7B Text Q5 1"},
		{input: "falcon:7b-text-q8_0", expected: "Falcon 7B Text Q8_0"},
		{input: "falcon:7b-text-fp16", expected: "Falcon 7B Text FP16"},
		{input: "falcon:40b-instruct", expected: "Falcon 40B Instruct"},
		{input: "falcon:40b-instruct-q4_0", expected: "Falcon 40B Instruct Q4_0"},
		{input: "falcon:40b-instruct-q4_1", expected: "Falcon 40B Instruct Q4 1"},
		{input: "falcon:40b-instruct-q5_0", expected: "Falcon 40B Instruct Q5 0"},
		{input: "falcon:40b-instruct-q5_1", expected: "Falcon 40B Instruct Q5 1"},
		{input: "falcon:40b-instruct-q8_0", expected: "Falcon 40B Instruct Q8_0"},
		{input: "falcon:40b-instruct-fp16", expected: "Falcon 40B Instruct FP16"},
		{input: "falcon:40b-text", expected: "Falcon 40B Text"},
		{input: "falcon:40b-text-q4_0", expected: "Falcon 40B Text Q4_0"},
		{input: "falcon:40b-text-q4_1", expected: "Falcon 40B Text Q4 1"},
		{input: "falcon:40b-text-q5_0", expected: "Falcon 40B Text Q5 0"},
		{input: "falcon:40b-text-q5_1", expected: "Falcon 40B Text Q5 1"},
		{input: "falcon:40b-text-q8_0", expected: "Falcon 40B Text Q8_0"},
		{input: "falcon:40b-text-fp16", expected: "Falcon 40B Text FP16"},
		{input: "falcon:180b-chat", expected: "Falcon 180B Chat"},
		{input: "falcon:180b-chat-q4_0", expected: "Falcon 180B Chat Q4_0"},
		{input: "falcon:180b-text", expected: "Falcon 180B Text"},
		{input: "falcon:180b-text-q4_0", expected: "Falcon 180B Text Q4_0"},
		{input: "deepseek-llm:latest", expected: "Deepseek Llm (latest)"},
		{input: "deepseek-llm", expected: "Deepseek Llm"},
		{input: "deepseek-llm:7b", expected: "Deepseek Llm 7B"},
		{input: "deepseek-llm:67b", expected: "Deepseek Llm 67B"},
		{input: "deepseek-llm:7b-base", expected: "Deepseek Llm 7B Base"},
		{input: "deepseek-llm:7b-base-q2_K", expected: "Deepseek Llm 7B Base Q2_K"},
		{input: "deepseek-llm:7b-base-q3_K_S", expected: "Deepseek Llm 7B Base Q3 K_S"},
		{input: "deepseek-llm:7b-base-q3_K_M", expected: "Deepseek Llm 7B Base Q3 K_M"},
		{input: "deepseek-llm:7b-base-q3_K_L", expected: "Deepseek Llm 7B Base Q3 K_L"},
		{input: "deepseek-llm:7b-base-q4_0", expected: "Deepseek Llm 7B Base Q4_0"},
		{input: "deepseek-llm:7b-base-q4_1", expected: "Deepseek Llm 7B Base Q4 1"},
		{input: "deepseek-llm:7b-base-q4_K_S", expected: "Deepseek Llm 7B Base Q4 K_S"},
		{input: "deepseek-llm:7b-base-q4_K_M", expected: "Deepseek Llm 7B Base Q4_K_M"},
		{input: "deepseek-llm:7b-base-q5_0", expected: "Deepseek Llm 7B Base Q5 0"},
		{input: "deepseek-llm:7b-base-q5_1", expected: "Deepseek Llm 7B Base Q5 1"},
		{input: "deepseek-llm:7b-base-q5_K_S", expected: "Deepseek Llm 7B Base Q5 K_S"},
		{input: "deepseek-llm:7b-base-q5_K_M", expected: "Deepseek Llm 7B Base Q5 K_M"},
		{input: "deepseek-llm:7b-base-q6_K", expected: "Deepseek Llm 7B Base Q6 K"},
		{input: "deepseek-llm:7b-base-q8_0", expected: "Deepseek Llm 7B Base Q8_0"},
		{input: "deepseek-llm:7b-base-fp16", expected: "Deepseek Llm 7B Base FP16"},
		{input: "deepseek-llm:7b-chat", expected: "Deepseek Llm 7B Chat"},
		{input: "deepseek-llm:7b-chat-q2_K", expected: "Deepseek Llm 7B Chat Q2_K"},
		{input: "deepseek-llm:7b-chat-q3_K_S", expected: "Deepseek Llm 7B Chat Q3 K_S"},
		{input: "deepseek-llm:7b-chat-q3_K_M", expected: "Deepseek Llm 7B Chat Q3 K_M"},
		{input: "deepseek-llm:7b-chat-q3_K_L", expected: "Deepseek Llm 7B Chat Q3 K_L"},
		{input: "deepseek-llm:7b-chat-q4_0", expected: "Deepseek Llm 7B Chat Q4_0"},
		{input: "deepseek-llm:7b-chat-q4_1", expected: "Deepseek Llm 7B Chat Q4 1"},
		{input: "deepseek-llm:7b-chat-q4_K_S", expected: "Deepseek Llm 7B Chat Q4 K_S"},
		{input: "deepseek-llm:7b-chat-q4_K_M", expected: "Deepseek Llm 7B Chat Q4_K_M"},
		{input: "deepseek-llm:7b-chat-q5_0", expected: "Deepseek Llm 7B Chat Q5 0"},
		{input: "deepseek-llm:7b-chat-q5_1", expected: "Deepseek Llm 7B Chat Q5 1"},
		{input: "deepseek-llm:7b-chat-q5_K_S", expected: "Deepseek Llm 7B Chat Q5 K_S"},
		{input: "deepseek-llm:7b-chat-q5_K_M", expected: "Deepseek Llm 7B Chat Q5 K_M"},
		{input: "deepseek-llm:7b-chat-q6_K", expected: "Deepseek Llm 7B Chat Q6 K"},
		{input: "deepseek-llm:7b-chat-q8_0", expected: "Deepseek Llm 7B Chat Q8_0"},
		{input: "deepseek-llm:7b-chat-fp16", expected: "Deepseek Llm 7B Chat FP16"},
		{input: "deepseek-llm:67b-base", expected: "Deepseek Llm 67B Base"},
		{input: "deepseek-llm:67b-base-q2_K", expected: "Deepseek Llm 67B Base Q2_K"},
		{input: "deepseek-llm:67b-base-q3_K_S", expected: "Deepseek Llm 67B Base Q3 K_S"},
		{input: "deepseek-llm:67b-base-q3_K_M", expected: "Deepseek Llm 67B Base Q3 K_M"},
		{input: "deepseek-llm:67b-base-q3_K_L", expected: "Deepseek Llm 67B Base Q3 K_L"},
		{input: "deepseek-llm:67b-base-q4_0", expected: "Deepseek Llm 67B Base Q4_0"},
		{input: "deepseek-llm:67b-base-q4_1", expected: "Deepseek Llm 67B Base Q4 1"},
		{input: "deepseek-llm:67b-base-q4_K_S", expected: "Deepseek Llm 67B Base Q4 K_S"},
		{input: "deepseek-llm:67b-base-q4_K_M", expected: "Deepseek Llm 67B Base Q4_K_M"},
		{input: "deepseek-llm:67b-base-q5_0", expected: "Deepseek Llm 67B Base Q5 0"},
		{input: "deepseek-llm:67b-base-q5_1", expected: "Deepseek Llm 67B Base Q5 1"},
		{input: "deepseek-llm:67b-base-q5_K_S", expected: "Deepseek Llm 67B Base Q5 K_S"},
		{input: "deepseek-llm:67b-base-q5_K_M", expected: "Deepseek Llm 67B Base Q5 K_M"},
		{input: "deepseek-llm:67b-base-q6_K", expected: "Deepseek Llm 67B Base Q6 K"},
		{input: "deepseek-llm:67b-base-q8_0", expected: "Deepseek Llm 67B Base Q8_0"},
		{input: "deepseek-llm:67b-base-fp16", expected: "Deepseek Llm 67B Base FP16"},
		{input: "deepseek-llm:67b-chat", expected: "Deepseek Llm 67B Chat"},
		{input: "deepseek-llm:67b-chat-q2_K", expected: "Deepseek Llm 67B Chat Q2_K"},
		{input: "deepseek-llm:67b-chat-q3_K_S", expected: "Deepseek Llm 67B Chat Q3 K_S"},
		{input: "deepseek-llm:67b-chat-q3_K_M", expected: "Deepseek Llm 67B Chat Q3 K_M"},
		{input: "deepseek-llm:67b-chat-q3_K_L", expected: "Deepseek Llm 67B Chat Q3 K_L"},
		{input: "deepseek-llm:67b-chat-q4_0", expected: "Deepseek Llm 67B Chat Q4_0"},
		{input: "deepseek-llm:67b-chat-q4_1", expected: "Deepseek Llm 67B Chat Q4 1"},
		{input: "deepseek-llm:67b-chat-q4_K_S", expected: "Deepseek Llm 67B Chat Q4 K_S"},
		{input: "deepseek-llm:67b-chat-q4_K_M", expected: "Deepseek Llm 67B Chat Q4_K_M"},
		{input: "deepseek-llm:67b-chat-q5_0", expected: "Deepseek Llm 67B Chat Q5 0"},
		{input: "deepseek-llm:67b-chat-q5_1", expected: "Deepseek Llm 67B Chat Q5 1"},
		{input: "deepseek-llm:67b-chat-q5_K_S", expected: "Deepseek Llm 67B Chat Q5 K_S"},
		{input: "deepseek-llm:67b-chat-fp16", expected: "Deepseek Llm 67B Chat FP16"},
		{input: "openchat:latest", expected: "Openchat (latest)"},
		{input: "openchat", expected: "Openchat"},
		{input: "openchat:7b", expected: "Openchat 7B"},
		{input: "openchat:7b-v3.5", expected: "Openchat 7B v3.5"},
		{input: "openchat:7b-v3.5-0106", expected: "Openchat 7B v3.5 0106"},
		{input: "openchat:7b-v3.5-0106-q2_K", expected: "Openchat 7B v3.5 0106 Q2_K"},
		{input: "openchat:7b-v3.5-q2_K", expected: "Openchat 7B v3.5 Q2_K"},
		{input: "openchat:7b-v3.5-0106-q3_K_S", expected: "Openchat 7B v3.5 0106 Q3 K_S"},
		{input: "openchat:7b-v3.5-q3_K_S", expected: "Openchat 7B v3.5 Q3 K_S"},
		{input: "openchat:7b-v3.5-0106-q3_K_M", expected: "Openchat 7B v3.5 0106 Q3 K_M"},
		{input: "openchat:7b-v3.5-q3_K_M", expected: "Openchat 7B v3.5 Q3 K_M"},
		{input: "openchat:7b-v3.5-0106-q3_K_L", expected: "Openchat 7B v3.5 0106 Q3 K_L"},
		{input: "openchat:7b-v3.5-q3_K_L", expected: "Openchat 7B v3.5 Q3 K_L"},
		{input: "openchat:7b-v3.5-0106-q4_0", expected: "Openchat 7B v3.5 0106 Q4_0"},
		{input: "openchat:7b-v3.5-q4_0", expected: "Openchat 7B v3.5 Q4_0"},
		{input: "openchat:7b-v3.5-0106-q4_1", expected: "Openchat 7B v3.5 0106 Q4 1"},
		{input: "openchat:7b-v3.5-q4_1", expected: "Openchat 7B v3.5 Q4 1"},
		{input: "openchat:7b-v3.5-0106-q4_K_S", expected: "Openchat 7B v3.5 0106 Q4 K_S"},
		{input: "openchat:7b-v3.5-q4_K_S", expected: "Openchat 7B v3.5 Q4 K_S"},
		{input: "openchat:7b-v3.5-0106-q4_K_M", expected: "Openchat 7B v3.5 0106 Q4_K_M"},
		{input: "openchat:7b-v3.5-q4_K_M", expected: "Openchat 7B v3.5 Q4_K_M"},
		{input: "openchat:7b-v3.5-0106-q5_0", expected: "Openchat 7B v3.5 0106 Q5 0"},
		{input: "openchat:7b-v3.5-q5_0", expected: "Openchat 7B v3.5 Q5 0"},
		{input: "openchat:7b-v3.5-0106-q5_1", expected: "Openchat 7B v3.5 0106 Q5 1"},
		{input: "openchat:7b-v3.5-q5_1", expected: "Openchat 7B v3.5 Q5 1"},
		{input: "openchat:7b-v3.5-0106-q5_K_S", expected: "Openchat 7B v3.5 0106 Q5 K_S"},
		{input: "openchat:7b-v3.5-0106-q5_K_M", expected: "Openchat 7B v3.5 0106 Q5 K_M"},
		{input: "openchat:7b-v3.5-0106-q6_K", expected: "Openchat 7B v3.5 0106 Q6 K"},
		{input: "openchat:7b-v3.5-0106-q8_0", expected: "Openchat 7B v3.5 0106 Q8_0"},
		{input: "openchat:7b-v3.5-0106-fp16", expected: "Openchat 7B v3.5 0106 FP16"},
		{input: "openchat:7b-v3.5-1210", expected: "Openchat 7B v3.5 1210"},
		{input: "openchat:7b-v3.5-1210-q2_K", expected: "Openchat 7B v3.5 1210 Q2_K"},
		{input: "openchat:7b-v3.5-1210-q3_K_S", expected: "Openchat 7B v3.5 1210 Q3 K_S"},
		{input: "openchat:7b-v3.5-1210-q3_K_M", expected: "Openchat 7B v3.5 1210 Q3 K_M"},
		{input: "openchat:7b-v3.5-1210-q3_K_L", expected: "Openchat 7B v3.5 1210 Q3 K_L"},
		{input: "openchat:7b-v3.5-1210-q4_0", expected: "Openchat 7B v3.5 1210 Q4_0"},
		{input: "openchat:7b-v3.5-1210-q4_1", expected: "Openchat 7B v3.5 1210 Q4 1"},
		{input: "openchat:7b-v3.5-1210-q4_K_S", expected: "Openchat 7B v3.5 1210 Q4 K_S"},
		{input: "openchat:7b-v3.5-1210-q4_K_M", expected: "Openchat 7B v3.5 1210 Q4_K_M"},
		{input: "openchat:7b-v3.5-1210-q5_0", expected: "Openchat 7B v3.5 1210 Q5 0"},
		{input: "openchat:7b-v3.5-1210-q5_1", expected: "Openchat 7B v3.5 1210 Q5 1"},
		{input: "openchat:7b-v3.5-1210-q5_K_S", expected: "Openchat 7B v3.5 1210 Q5 K_S"},
		{input: "openchat:7b-v3.5-q5_K_S", expected: "Openchat 7B v3.5 Q5 K_S"},
		{input: "openchat:7b-v3.5-1210-q5_K_M", expected: "Openchat 7B v3.5 1210 Q5 K_M"},
		{input: "openchat:7b-v3.5-q5_K_M", expected: "Openchat 7B v3.5 Q5 K_M"},
		{input: "openchat:7b-v3.5-1210-q6_K", expected: "Openchat 7B v3.5 1210 Q6 K"},
		{input: "openchat:7b-v3.5-q6_K", expected: "Openchat 7B v3.5 Q6 K"},
		{input: "openchat:7b-v3.5-1210-q8_0", expected: "Openchat 7B v3.5 1210 Q8_0"},
		{input: "openchat:7b-v3.5-q8_0", expected: "Openchat 7B v3.5 Q8_0"},
		{input: "openchat:7b-v3.5-1210-fp16", expected: "Openchat 7B v3.5 1210 FP16"},
		{input: "openchat:7b-v3.5-fp16", expected: "Openchat 7B v3.5 FP16"},
		{input: "vicuna:latest", expected: "Vicuna (latest)"},
		{input: "vicuna", expected: "Vicuna"},
		{input: "vicuna:7b", expected: "Vicuna 7B"},
		{input: "vicuna:13b", expected: "Vicuna 13B"},
		{input: "vicuna:33b", expected: "Vicuna 33B"},
		{input: "vicuna:7b-16k", expected: "Vicuna 7B 16K"},
		{input: "vicuna:7b-v1.5-16k-q2_K", expected: "Vicuna 7B v1.5 16K Q2_K"},
		{input: "vicuna:7b-v1.5-q2_K", expected: "Vicuna 7B v1.5 Q2_K"},
		{input: "vicuna:7b-v1.5-16k-q3_K_S", expected: "Vicuna 7B v1.5 16K Q3 K_S"},
		{input: "vicuna:7b-v1.5-q3_K_S", expected: "Vicuna 7B v1.5 Q3 K_S"},
		{input: "vicuna:7b-v1.5-16k-q3_K_M", expected: "Vicuna 7B v1.5 16K Q3 K_M"},
		{input: "vicuna:7b-v1.5-q3_K_M", expected: "Vicuna 7B v1.5 Q3 K_M"},
		{input: "vicuna:7b-v1.5-16k-q3_K_L", expected: "Vicuna 7B v1.5 16K Q3 K_L"},
		{input: "vicuna:7b-v1.5-q3_K_L", expected: "Vicuna 7B v1.5 Q3 K_L"},
		{input: "vicuna:7b-v1.5-16k-q4_0", expected: "Vicuna 7B v1.5 16K Q4_0"},
		{input: "vicuna:7b-v1.5-q4_0", expected: "Vicuna 7B v1.5 Q4_0"},
		{input: "vicuna:7b-v1.5-16k-q4_1", expected: "Vicuna 7B v1.5 16K Q4 1"},
		{input: "vicuna:7b-v1.5-q4_1", expected: "Vicuna 7B v1.5 Q4 1"},
		{input: "vicuna:7b-v1.5-16k-q4_K_S", expected: "Vicuna 7B v1.5 16K Q4 K_S"},
		{input: "vicuna:7b-v1.5-q4_K_S", expected: "Vicuna 7B v1.5 Q4 K_S"},
		{input: "vicuna:7b-v1.5-16k-q4_K_M", expected: "Vicuna 7B v1.5 16K Q4_K_M"},
		{input: "vicuna:7b-v1.5-q4_K_M", expected: "Vicuna 7B v1.5 Q4_K_M"},
		{input: "vicuna:7b-v1.5-16k-q5_0", expected: "Vicuna 7B v1.5 16K Q5 0"},
		{input: "vicuna:7b-v1.5-q5_0", expected: "Vicuna 7B v1.5 Q5 0"},
		{input: "vicuna:7b-v1.5-16k-q5_1", expected: "Vicuna 7B v1.5 16K Q5 1"},
		{input: "vicuna:7b-v1.5-q5_1", expected: "Vicuna 7B v1.5 Q5 1"},
		{input: "vicuna:7b-v1.5-16k-q5_K_S", expected: "Vicuna 7B v1.5 16K Q5 K_S"},
		{input: "vicuna:7b-v1.5-q5_K_S", expected: "Vicuna 7B v1.5 Q5 K_S"},
		{input: "vicuna:7b-v1.5-16k-q5_K_M", expected: "Vicuna 7B v1.5 16K Q5 K_M"},
		{input: "vicuna:7b-v1.5-q5_K_M", expected: "Vicuna 7B v1.5 Q5 K_M"},
		{input: "vicuna:7b-v1.5-16k-q6_K", expected: "Vicuna 7B v1.5 16K Q6 K"},
		{input: "vicuna:7b-v1.5-q6_K", expected: "Vicuna 7B v1.5 Q6 K"},
		{input: "vicuna:7b-v1.5-16k-q8_0", expected: "Vicuna 7B v1.5 16K Q8_0"},
		{input: "vicuna:7b-v1.5-q8_0", expected: "Vicuna 7B v1.5 Q8_0"},
		{input: "vicuna:7b-v1.5-16k-fp16", expected: "Vicuna 7B v1.5 16K FP16"},
		{input: "vicuna:7b-v1.5-fp16", expected: "Vicuna 7B v1.5 FP16"},
		{input: "vicuna:7b-q2_K", expected: "Vicuna 7B Q2_K"},
		{input: "vicuna:7b-q3_K_S", expected: "Vicuna 7B Q3 K_S"},
		{input: "vicuna:7b-q3_K_M", expected: "Vicuna 7B Q3 K_M"},
		{input: "vicuna:7b-q3_K_L", expected: "Vicuna 7B Q3 K_L"},
		{input: "vicuna:7b-q4_0", expected: "Vicuna 7B Q4_0"},
		{input: "vicuna:7b-q4_1", expected: "Vicuna 7B Q4 1"},
		{input: "vicuna:7b-q4_K_S", expected: "Vicuna 7B Q4 K_S"},
		{input: "vicuna:7b-q4_K_M", expected: "Vicuna 7B Q4_K_M"},
		{input: "vicuna:7b-q5_0", expected: "Vicuna 7B Q5 0"},
		{input: "vicuna:7b-q5_1", expected: "Vicuna 7B Q5 1"},
		{input: "vicuna:7b-q5_K_S", expected: "Vicuna 7B Q5 K_S"},
		{input: "vicuna:7b-q5_K_M", expected: "Vicuna 7B Q5 K_M"},
		{input: "vicuna:7b-q6_K", expected: "Vicuna 7B Q6 K"},
		{input: "vicuna:7b-q8_0", expected: "Vicuna 7B Q8_0"},
		{input: "vicuna:7b-fp16", expected: "Vicuna 7B FP16"},
		{input: "vicuna:13b-16k", expected: "Vicuna 13B 16K"},
		{input: "vicuna:13b-v1.5-16k-q2_K", expected: "Vicuna 13B v1.5 16K Q2_K"},
		{input: "vicuna:13b-v1.5-q2_K", expected: "Vicuna 13B v1.5 Q2_K"},
		{input: "vicuna:13b-v1.5-16k-q3_K_S", expected: "Vicuna 13B v1.5 16K Q3 K_S"},
		{input: "vicuna:13b-v1.5-q3_K_S", expected: "Vicuna 13B v1.5 Q3 K_S"},
		{input: "vicuna:13b-v1.5-16k-q3_K_M", expected: "Vicuna 13B v1.5 16K Q3 K_M"},
		{input: "vicuna:13b-v1.5-q3_K_M", expected: "Vicuna 13B v1.5 Q3 K_M"},
		{input: "vicuna:13b-v1.5-16k-q3_K_L", expected: "Vicuna 13B v1.5 16K Q3 K_L"},
		{input: "vicuna:13b-v1.5-q3_K_L", expected: "Vicuna 13B v1.5 Q3 K_L"},
		{input: "vicuna:13b-v1.5-16k-q4_0", expected: "Vicuna 13B v1.5 16K Q4_0"},
		{input: "vicuna:13b-v1.5-q4_0", expected: "Vicuna 13B v1.5 Q4_0"},
		{input: "vicuna:13b-v1.5-16k-q4_1", expected: "Vicuna 13B v1.5 16K Q4 1"},
		{input: "vicuna:13b-v1.5-q4_1", expected: "Vicuna 13B v1.5 Q4 1"},
		{input: "vicuna:13b-v1.5-16k-q4_K_S", expected: "Vicuna 13B v1.5 16K Q4 K_S"},
		{input: "vicuna:13b-v1.5-q4_K_S", expected: "Vicuna 13B v1.5 Q4 K_S"},
		{input: "vicuna:13b-v1.5-16k-q4_K_M", expected: "Vicuna 13B v1.5 16K Q4_K_M"},
		{input: "vicuna:13b-v1.5-q4_K_M", expected: "Vicuna 13B v1.5 Q4_K_M"},
		{input: "vicuna:13b-v1.5-16k-q5_0", expected: "Vicuna 13B v1.5 16K Q5 0"},
		{input: "vicuna:13b-v1.5-q5_0", expected: "Vicuna 13B v1.5 Q5 0"},
		{input: "vicuna:13b-v1.5-16k-q5_1", expected: "Vicuna 13B v1.5 16K Q5 1"},
		{input: "vicuna:13b-v1.5-q5_1", expected: "Vicuna 13B v1.5 Q5 1"},
		{input: "vicuna:13b-v1.5-16k-q5_K_S", expected: "Vicuna 13B v1.5 16K Q5 K_S"},
		{input: "vicuna:13b-v1.5-q5_K_S", expected: "Vicuna 13B v1.5 Q5 K_S"},
		{input: "vicuna:13b-v1.5-16k-q5_K_M", expected: "Vicuna 13B v1.5 16K Q5 K_M"},
		{input: "vicuna:13b-v1.5-q5_K_M", expected: "Vicuna 13B v1.5 Q5 K_M"},
		{input: "vicuna:13b-v1.5-16k-q6_K", expected: "Vicuna 13B v1.5 16K Q6 K"},
		{input: "vicuna:13b-v1.5-q6_K", expected: "Vicuna 13B v1.5 Q6 K"},
		{input: "vicuna:13b-v1.5-16k-q8_0", expected: "Vicuna 13B v1.5 16K Q8_0"},
		{input: "vicuna:13b-v1.5-q8_0", expected: "Vicuna 13B v1.5 Q8_0"},
		{input: "vicuna:13b-v1.5-16k-fp16", expected: "Vicuna 13B v1.5 16K FP16"},
		{input: "vicuna:13b-v1.5-fp16", expected: "Vicuna 13B v1.5 FP16"},
		{input: "vicuna:13b-q2_K", expected: "Vicuna 13B Q2_K"},
		{input: "vicuna:13b-q3_K_S", expected: "Vicuna 13B Q3 K_S"},
		{input: "vicuna:13b-q3_K_M", expected: "Vicuna 13B Q3 K_M"},
		{input: "vicuna:13b-q3_K_L", expected: "Vicuna 13B Q3 K_L"},
		{input: "vicuna:13b-q4_0", expected: "Vicuna 13B Q4_0"},
		{input: "vicuna:13b-q4_1", expected: "Vicuna 13B Q4 1"},
		{input: "vicuna:13b-q4_K_S", expected: "Vicuna 13B Q4 K_S"},
		{input: "vicuna:13b-q4_K_M", expected: "Vicuna 13B Q4_K_M"},
		{input: "vicuna:13b-q5_0", expected: "Vicuna 13B Q5 0"},
		{input: "vicuna:13b-q5_1", expected: "Vicuna 13B Q5 1"},
		{input: "vicuna:13b-q5_K_S", expected: "Vicuna 13B Q5 K_S"},
		{input: "vicuna:13b-q5_K_M", expected: "Vicuna 13B Q5 K_M"},
		{input: "vicuna:13b-q6_K", expected: "Vicuna 13B Q6 K"},
		{input: "vicuna:13b-q8_0", expected: "Vicuna 13B Q8_0"},
		{input: "vicuna:13b-fp16", expected: "Vicuna 13B FP16"},
		{input: "vicuna:33b-q2_K", expected: "Vicuna 33B Q2_K"},
		{input: "vicuna:33b-q3_K_S", expected: "Vicuna 33B Q3 K_S"},
		{input: "vicuna:33b-q3_K_M", expected: "Vicuna 33B Q3 K_M"},
		{input: "vicuna:33b-q3_K_L", expected: "Vicuna 33B Q3 K_L"},
		{input: "vicuna:33b-q4_0", expected: "Vicuna 33B Q4_0"},
		{input: "vicuna:33b-q4_1", expected: "Vicuna 33B Q4 1"},
		{input: "vicuna:33b-q4_K_S", expected: "Vicuna 33B Q4 K_S"},
		{input: "vicuna:33b-q4_K_M", expected: "Vicuna 33B Q4_K_M"},
		{input: "vicuna:33b-q5_0", expected: "Vicuna 33B Q5 0"},
		{input: "vicuna:33b-q5_1", expected: "Vicuna 33B Q5 1"},
		{input: "vicuna:33b-q5_K_S", expected: "Vicuna 33B Q5 K_S"},
		{input: "vicuna:33b-q5_K_M", expected: "Vicuna 33B Q5 K_M"},
		{input: "vicuna:33b-q6_K", expected: "Vicuna 33B Q6 K"},
		{input: "vicuna:33b-q8_0", expected: "Vicuna 33B Q8_0"},
		{input: "vicuna:33b-fp16", expected: "Vicuna 33B FP16"},
		{input: "deepseek-v2:latest", expected: "Deepseek v2 (latest)"},
		{input: "deepseek-v2", expected: "Deepseek v2"},
		{input: "deepseek-v2:lite", expected: "Deepseek v2 Lite"},
		{input: "deepseek-v2:16b", expected: "Deepseek v2 16B"},
		{input: "deepseek-v2:236b", expected: "Deepseek v2 236B"},
		{input: "deepseek-v2:16b-lite-chat-q2_K", expected: "Deepseek v2 16B Lite Chat Q2_K"},
		{input: "deepseek-v2:16b-lite-chat-q3_K_S", expected: "Deepseek v2 16B Lite Chat Q3 K_S"},
		{input: "deepseek-v2:16b-lite-chat-q3_K_M", expected: "Deepseek v2 16B Lite Chat Q3 K_M"},
		{input: "deepseek-v2:16b-lite-chat-q3_K_L", expected: "Deepseek v2 16B Lite Chat Q3 K_L"},
		{input: "deepseek-v2:16b-lite-chat-q4_0", expected: "Deepseek v2 16B Lite Chat Q4_0"},
		{input: "deepseek-v2:16b-lite-chat-q4_1", expected: "Deepseek v2 16B Lite Chat Q4 1"},
		{input: "deepseek-v2:16b-lite-chat-q4_K_S", expected: "Deepseek v2 16B Lite Chat Q4 K_S"},
		{input: "deepseek-v2:16b-lite-chat-q4_K_M", expected: "Deepseek v2 16B Lite Chat Q4_K_M"},
		{input: "deepseek-v2:16b-lite-chat-q5_0", expected: "Deepseek v2 16B Lite Chat Q5 0"},
		{input: "deepseek-v2:16b-lite-chat-q5_1", expected: "Deepseek v2 16B Lite Chat Q5 1"},
		{input: "deepseek-v2:16b-lite-chat-q5_K_S", expected: "Deepseek v2 16B Lite Chat Q5 K_S"},
		{input: "deepseek-v2:16b-lite-chat-q5_K_M", expected: "Deepseek v2 16B Lite Chat Q5 K_M"},
		{input: "deepseek-v2:16b-lite-chat-q6_K", expected: "Deepseek v2 16B Lite Chat Q6 K"},
		{input: "deepseek-v2:16b-lite-chat-q8_0", expected: "Deepseek v2 16B Lite Chat Q8_0"},
		{input: "deepseek-v2:16b-lite-chat-fp16", expected: "Deepseek v2 16B Lite Chat FP16"},
		{input: "deepseek-v2:236b-chat-q2_K", expected: "Deepseek v2 236B Chat Q2_K"},
		{input: "deepseek-v2:236b-chat-q3_K_S", expected: "Deepseek v2 236B Chat Q3 K_S"},
		{input: "deepseek-v2:236b-chat-q3_K_M", expected: "Deepseek v2 236B Chat Q3 K_M"},
		{input: "deepseek-v2:236b-chat-q3_K_L", expected: "Deepseek v2 236B Chat Q3 K_L"},
		{input: "deepseek-v2:236b-chat-q4_0", expected: "Deepseek v2 236B Chat Q4_0"},
		{input: "deepseek-v2:236b-chat-q4_1", expected: "Deepseek v2 236B Chat Q4 1"},
		{input: "deepseek-v2:236b-chat-q4_K_S", expected: "Deepseek v2 236B Chat Q4 K_S"},
		{input: "deepseek-v2:236b-chat-q4_K_M", expected: "Deepseek v2 236B Chat Q4_K_M"},
		{input: "deepseek-v2:236b-chat-q5_0", expected: "Deepseek v2 236B Chat Q5 0"},
		{input: "deepseek-v2:236b-chat-q5_1", expected: "Deepseek v2 236B Chat Q5 1"},
		{input: "deepseek-v2:236b-chat-q5_K_S", expected: "Deepseek v2 236B Chat Q5 K_S"},
		{input: "deepseek-v2:236b-chat-q5_K_M", expected: "Deepseek v2 236B Chat Q5 K_M"},
		{input: "deepseek-v2:236b-chat-q6_K", expected: "Deepseek v2 236B Chat Q6 K"},
		{input: "deepseek-v2:236b-chat-q8_0", expected: "Deepseek v2 236B Chat Q8_0"},
		{input: "deepseek-v2:236b-chat-fp16", expected: "Deepseek v2 236B Chat FP16"},
		{input: "openhermes:latest", expected: "Openhermes (latest)"},
		{input: "openhermes", expected: "Openhermes"},
		{input: "openhermes:v2", expected: "Openhermes v2"},
		{input: "openhermes:v2.5", expected: "Openhermes v2.5"},
		{input: "openhermes:7b-mistral-v2-q2_K", expected: "Openhermes 7B Mistral v2 Q2_K"},
		{input: "openhermes:7b-mistral-v2-q3_K_S", expected: "Openhermes 7B Mistral v2 Q3 K_S"},
		{input: "openhermes:7b-mistral-v2-q3_K_M", expected: "Openhermes 7B Mistral v2 Q3 K_M"},
		{input: "openhermes:7b-mistral-v2-q3_K_L", expected: "Openhermes 7B Mistral v2 Q3 K_L"},
		{input: "openhermes:7b-mistral-v2-q4_0", expected: "Openhermes 7B Mistral v2 Q4_0"},
		{input: "openhermes:7b-mistral-v2-q4_1", expected: "Openhermes 7B Mistral v2 Q4 1"},
		{input: "openhermes:7b-mistral-v2-q4_K_S", expected: "Openhermes 7B Mistral v2 Q4 K_S"},
		{input: "openhermes:7b-mistral-v2-q4_K_M", expected: "Openhermes 7B Mistral v2 Q4_K_M"},
		{input: "openhermes:7b-mistral-v2-q5_0", expected: "Openhermes 7B Mistral v2 Q5 0"},
		{input: "openhermes:7b-mistral-v2-q5_1", expected: "Openhermes 7B Mistral v2 Q5 1"},
		{input: "openhermes:7b-mistral-v2-q5_K_S", expected: "Openhermes 7B Mistral v2 Q5 K_S"},
		{input: "openhermes:7b-mistral-v2-q5_K_M", expected: "Openhermes 7B Mistral v2 Q5 K_M"},
		{input: "openhermes:7b-mistral-v2-q6_K", expected: "Openhermes 7B Mistral v2 Q6 K"},
		{input: "openhermes:7b-mistral-v2-q8_0", expected: "Openhermes 7B Mistral v2 Q8_0"},
		{input: "openhermes:7b-mistral-v2-fp16", expected: "Openhermes 7B Mistral v2 FP16"},
		{input: "openhermes:7b-mistral-v2.5-q2_K", expected: "Openhermes 7B Mistral v2.5 Q2_K"},
		{input: "openhermes:7b-mistral-v2.5-q3_K_S", expected: "Openhermes 7B Mistral v2.5 Q3 K_S"},
		{input: "openhermes:7b-mistral-v2.5-q3_K_M", expected: "Openhermes 7B Mistral v2.5 Q3 K_M"},
		{input: "openhermes:7b-mistral-v2.5-q3_K_L", expected: "Openhermes 7B Mistral v2.5 Q3 K_L"},
		{input: "openhermes:7b-mistral-v2.5-q4_0", expected: "Openhermes 7B Mistral v2.5 Q4_0"},
		{input: "openhermes:7b-mistral-v2.5-q4_1", expected: "Openhermes 7B Mistral v2.5 Q4 1"},
		{input: "openhermes:7b-mistral-v2.5-q4_K_S", expected: "Openhermes 7B Mistral v2.5 Q4 K_S"},
		{input: "openhermes:7b-mistral-v2.5-q4_K_M", expected: "Openhermes 7B Mistral v2.5 Q4_K_M"},
		{input: "openhermes:7b-mistral-v2.5-q5_0", expected: "Openhermes 7B Mistral v2.5 Q5 0"},
		{input: "openhermes:7b-mistral-v2.5-q5_1", expected: "Openhermes 7B Mistral v2.5 Q5 1"},
		{input: "openhermes:7b-mistral-v2.5-q5_K_S", expected: "Openhermes 7B Mistral v2.5 Q5 K_S"},
		{input: "openhermes:7b-mistral-v2.5-q5_K_M", expected: "Openhermes 7B Mistral v2.5 Q5 K_M"},
		{input: "openhermes:7b-mistral-v2.5-q6_K", expected: "Openhermes 7B Mistral v2.5 Q6 K"},
		{input: "openhermes:7b-mistral-v2.5-q8_0", expected: "Openhermes 7B Mistral v2.5 Q8_0"},
		{input: "openhermes:7b-mistral-v2.5-fp16", expected: "Openhermes 7B Mistral v2.5 FP16"},
		{input: "openhermes:7b-v2", expected: "Openhermes 7B v2"},
		{input: "openhermes:7b-v2.5", expected: "Openhermes 7B v2.5"},
		{input: "codegeex4:latest", expected: "CodeGeeX4 (latest)"},
		{input: "codegeex4", expected: "CodeGeeX4"},
		{input: "codegeex4:9b", expected: "CodeGeeX4 9B"},
		{input: "codegeex4:9b-all-q2_K", expected: "CodeGeeX4 9B All Q2_K"},
		{input: "codegeex4:9b-all-q3_K_S", expected: "CodeGeeX4 9B All Q3 K_S"},
		{input: "codegeex4:9b-all-q3_K_M", expected: "CodeGeeX4 9B All Q3 K_M"},
		{input: "codegeex4:9b-all-q3_K_L", expected: "CodeGeeX4 9B All Q3 K_L"},
		{input: "codegeex4:9b-all-q4_0", expected: "CodeGeeX4 9B All Q4_0"},
		{input: "codegeex4:9b-all-q4_1", expected: "CodeGeeX4 9B All Q4 1"},
		{input: "codegeex4:9b-all-q4_K_S", expected: "CodeGeeX4 9B All Q4 K_S"},
		{input: "codegeex4:9b-all-q4_K_M", expected: "CodeGeeX4 9B All Q4_K_M"},
		{input: "codegeex4:9b-all-q5_0", expected: "CodeGeeX4 9B All Q5 0"},
		{input: "codegeex4:9b-all-q5_1", expected: "CodeGeeX4 9B All Q5 1"},
		{input: "codegeex4:9b-all-q5_K_S", expected: "CodeGeeX4 9B All Q5 K_S"},
		{input: "codegeex4:9b-all-q5_K_M", expected: "CodeGeeX4 9B All Q5 K_M"},
		{input: "codegeex4:9b-all-q6_K", expected: "CodeGeeX4 9B All Q6 K"},
		{input: "codegeex4:9b-all-q8_0", expected: "CodeGeeX4 9B All Q8_0"},
		{input: "codegeex4:9b-all-fp16", expected: "CodeGeeX4 9B All FP16"},
		{input: "mistral-openorca:latest", expected: "Mistral Openorca (latest)"},
		{input: "mistral-openorca", expected: "Mistral Openorca"},
		{input: "mistral-openorca:7b", expected: "Mistral Openorca 7B"},
		{input: "mistral-openorca:7b-q2_K", expected: "Mistral Openorca 7B Q2_K"},
		{input: "mistral-openorca:7b-q3_K_S", expected: "Mistral Openorca 7B Q3 K_S"},
		{input: "mistral-openorca:7b-q3_K_M", expected: "Mistral Openorca 7B Q3 K_M"},
		{input: "mistral-openorca:7b-q3_K_L", expected: "Mistral Openorca 7B Q3 K_L"},
		{input: "mistral-openorca:7b-q4_0", expected: "Mistral Openorca 7B Q4_0"},
		{input: "mistral-openorca:7b-q4_1", expected: "Mistral Openorca 7B Q4 1"},
		{input: "mistral-openorca:7b-q4_K_S", expected: "Mistral Openorca 7B Q4 K_S"},
		{input: "mistral-openorca:7b-q4_K_M", expected: "Mistral Openorca 7B Q4_K_M"},
		{input: "mistral-openorca:7b-q5_0", expected: "Mistral Openorca 7B Q5 0"},
		{input: "mistral-openorca:7b-q5_1", expected: "Mistral Openorca 7B Q5 1"},
		{input: "mistral-openorca:7b-q5_K_S", expected: "Mistral Openorca 7B Q5 K_S"},
		{input: "mistral-openorca:7b-q5_K_M", expected: "Mistral Openorca 7B Q5 K_M"},
		{input: "mistral-openorca:7b-q6_K", expected: "Mistral Openorca 7B Q6 K"},
		{input: "mistral-openorca:7b-q8_0", expected: "Mistral Openorca 7B Q8_0"},
		{input: "mistral-openorca:7b-fp16", expected: "Mistral Openorca 7B FP16"},
		{input: "deepseek-v3.1:latest", expected: "Deepseek v3.1 (latest)"},
		{input: "deepseek-v3.1", expected: "Deepseek v3.1"},
		{input: "deepseek-v3.1:671b", expected: "Deepseek v3.1 671B"},
		{input: "deepseek-v3.1:671b-cloud", expected: "Deepseek v3.1 671B Cloud"},
		{input: "deepseek-v3.1:671b-terminus-q4_K_M", expected: "Deepseek v3.1 671B Terminus Q4_K_M"},
		{input: "deepseek-v3.1:671b-terminus-q8_0", expected: "Deepseek v3.1 671B Terminus Q8_0"},
		{input: "deepseek-v3.1:671b-terminus-fp16", expected: "Deepseek v3.1 671B Terminus FP16"},
		{input: "deepseek-v3.1:671b-q8_0", expected: "Deepseek v3.1 671B Q8_0"},
		{input: "deepseek-v3.1:671b-fp16", expected: "Deepseek v3.1 671B FP16"},
		{input: "codeqwen:latest", expected: "Codeqwen (latest)"},
		{input: "codeqwen", expected: "Codeqwen"},
		{input: "codeqwen:chat", expected: "Codeqwen Chat"},
		{input: "codeqwen:code", expected: "Codeqwen Code"},
		{input: "codeqwen:v1.5", expected: "Codeqwen v1.5"},
		{input: "codeqwen:7b", expected: "Codeqwen 7B"},
		{input: "codeqwen:7b-chat", expected: "Codeqwen 7B Chat"},
		{input: "codeqwen:7b-chat-v1.5-q2_K", expected: "Codeqwen 7B Chat v1.5 Q2_K"},
		{input: "codeqwen:7b-chat-v1.5-q3_K_S", expected: "Codeqwen 7B Chat v1.5 Q3 K_S"},
		{input: "codeqwen:7b-chat-v1.5-q3_K_M", expected: "Codeqwen 7B Chat v1.5 Q3 K_M"},
		{input: "codeqwen:7b-chat-v1.5-q3_K_L", expected: "Codeqwen 7B Chat v1.5 Q3 K_L"},
		{input: "codeqwen:7b-chat-v1.5-q4_0", expected: "Codeqwen 7B Chat v1.5 Q4_0"},
		{input: "codeqwen:7b-chat-v1.5-q4_1", expected: "Codeqwen 7B Chat v1.5 Q4 1"},
		{input: "codeqwen:7b-chat-v1.5-q4_K_S", expected: "Codeqwen 7B Chat v1.5 Q4 K_S"},
		{input: "codeqwen:7b-chat-v1.5-q4_K_M", expected: "Codeqwen 7B Chat v1.5 Q4_K_M"},
		{input: "codeqwen:7b-chat-v1.5-q5_0", expected: "Codeqwen 7B Chat v1.5 Q5 0"},
		{input: "codeqwen:7b-chat-v1.5-q5_1", expected: "Codeqwen 7B Chat v1.5 Q5 1"},
		{input: "codeqwen:7b-chat-v1.5-q5_K_S", expected: "Codeqwen 7B Chat v1.5 Q5 K_S"},
		{input: "codeqwen:7b-chat-v1.5-q5_K_M", expected: "Codeqwen 7B Chat v1.5 Q5 K_M"},
		{input: "codeqwen:7b-chat-v1.5-q6_K", expected: "Codeqwen 7B Chat v1.5 Q6 K"},
		{input: "codeqwen:7b-chat-v1.5-q8_0", expected: "Codeqwen 7B Chat v1.5 Q8_0"},
		{input: "codeqwen:7b-chat-v1.5-fp16", expected: "Codeqwen 7B Chat v1.5 FP16"},
		{input: "codeqwen:7b-code", expected: "Codeqwen 7B Code"},
		{input: "codeqwen:7b-code-v1.5-q4_0", expected: "Codeqwen 7B Code v1.5 Q4_0"},
		{input: "codeqwen:7b-code-v1.5-q4_1", expected: "Codeqwen 7B Code v1.5 Q4 1"},
		{input: "codeqwen:7b-code-v1.5-q5_0", expected: "Codeqwen 7B Code v1.5 Q5 0"},
		{input: "codeqwen:7b-code-v1.5-q5_1", expected: "Codeqwen 7B Code v1.5 Q5 1"},
		{input: "codeqwen:7b-code-v1.5-q8_0", expected: "Codeqwen 7B Code v1.5 Q8_0"},
		{input: "codeqwen:7b-code-v1.5-fp16", expected: "Codeqwen 7B Code v1.5 FP16"},
		{input: "codeqwen:v1.5-chat", expected: "Codeqwen v1.5 Chat"},
		{input: "codeqwen:v1.5-code", expected: "Codeqwen v1.5 Code"},
		{input: "snowflake-arctic-embed2:latest", expected: "Snowflake Arctic EMBED2 (latest)"},
		{input: "snowflake-arctic-embed2", expected: "Snowflake Arctic EMBED2"},
		{input: "snowflake-arctic-embed2:568m", expected: "Snowflake Arctic EMBED2 568M"},
		{input: "snowflake-arctic-embed2:568m-l-fp16", expected: "Snowflake Arctic EMBED2 568M L FP16"},
		{input: "qwen3-next:latest", expected: "Qwen3 Next (latest)"},
		{input: "qwen3-next", expected: "Qwen3 Next"},
		{input: "qwen3-next:80b", expected: "Qwen3 Next 80B"},
		{input: "qwen3-next:80b-a3b-instruct-q4_K_M", expected: "Qwen3 Next 80B A3B Instruct Q4_K_M"},
		{input: "qwen3-next:80b-a3b-instruct-q8_0", expected: "Qwen3 Next 80B A3B Instruct Q8_0"},
		{input: "qwen3-next:80b-a3b-instruct-fp16", expected: "Qwen3 Next 80B A3B Instruct FP16"},
		{input: "qwen3-next:80b-a3b-thinking", expected: "Qwen3 Next 80B A3B Thinking"},
		{input: "qwen3-next:80b-a3b-thinking-q4_K_M", expected: "Qwen3 Next 80B A3B Thinking Q4_K_M"},
		{input: "qwen3-next:80b-a3b-thinking-q8_0", expected: "Qwen3 Next 80B A3B Thinking Q8_0"},
		{input: "qwen3-next:80b-a3b-thinking-fp16", expected: "Qwen3 Next 80B A3B Thinking FP16"},
		{input: "qwen3-next:80b-cloud", expected: "Qwen3 Next 80B Cloud"},
		{input: "command-r-plus:latest", expected: "Command R Plus (latest)"},
		{input: "command-r-plus", expected: "Command R Plus"},
		{input: "command-r-plus:104b", expected: "Command R Plus 104B"},
		{input: "command-r-plus:104b-08-2024-q2_K", expected: "Command R Plus 104B (2024-08) Q2_K"},
		{input: "command-r-plus:104b-08-2024-q3_K_S", expected: "Command R Plus 104B (2024-08) Q3 K_S"},
		{input: "command-r-plus:104b-08-2024-q3_K_M", expected: "Command R Plus 104B (2024-08) Q3 K_M"},
		{input: "command-r-plus:104b-08-2024-q3_K_L", expected: "Command R Plus 104B (2024-08) Q3 K_L"},
		{input: "command-r-plus:104b-08-2024-q4_0", expected: "Command R Plus 104B (2024-08) Q4_0"},
		{input: "command-r-plus:104b-08-2024-q4_1", expected: "Command R Plus 104B (2024-08) Q4 1"},
		{input: "command-r-plus:104b-08-2024-q4_K_S", expected: "Command R Plus 104B (2024-08) Q4 K_S"},
		{input: "command-r-plus:104b-08-2024-q4_K_M", expected: "Command R Plus 104B (2024-08) Q4_K_M"},
		{input: "command-r-plus:104b-08-2024-q5_0", expected: "Command R Plus 104B (2024-08) Q5 0"},
		{input: "command-r-plus:104b-08-2024-q5_1", expected: "Command R Plus 104B (2024-08) Q5 1"},
		{input: "command-r-plus:104b-08-2024-q5_K_S", expected: "Command R Plus 104B (2024-08) Q5 K_S"},
		{input: "command-r-plus:104b-08-2024-q5_K_M", expected: "Command R Plus 104B (2024-08) Q5 K_M"},
		{input: "command-r-plus:104b-08-2024-q6_K", expected: "Command R Plus 104B (2024-08) Q6 K"},
		{input: "command-r-plus:104b-08-2024-q8_0", expected: "Command R Plus 104B (2024-08) Q8_0"},
		{input: "command-r-plus:104b-08-2024-fp16", expected: "Command R Plus 104B (2024-08) FP16"},
		{input: "command-r-plus:104b-q2_K", expected: "Command R Plus 104B Q2_K"},
		{input: "command-r-plus:104b-q4_0", expected: "Command R Plus 104B Q4_0"},
		{input: "command-r-plus:104b-q8_0", expected: "Command R Plus 104B Q8_0"},
		{input: "command-r-plus:104b-fp16", expected: "Command R Plus 104B FP16"},
		{input: "qwen2-math:latest", expected: "Qwen2 Math (latest)"},
		{input: "qwen2-math", expected: "Qwen2 Math"},
		{input: "qwen2-math:1.5b", expected: "Qwen2 Math 1.5B"},
		{input: "qwen2-math:7b", expected: "Qwen2 Math 7B"},
		{input: "qwen2-math:72b", expected: "Qwen2 Math 72B"},
		{input: "qwen2-math:1.5b-instruct", expected: "Qwen2 Math 1.5B Instruct"},
		{input: "qwen2-math:1.5b-instruct-q2_K", expected: "Qwen2 Math 1.5B Instruct Q2_K"},
		{input: "qwen2-math:1.5b-instruct-q3_K_S", expected: "Qwen2 Math 1.5B Instruct Q3 K_S"},
		{input: "qwen2-math:1.5b-instruct-q3_K_M", expected: "Qwen2 Math 1.5B Instruct Q3 K_M"},
		{input: "qwen2-math:1.5b-instruct-q3_K_L", expected: "Qwen2 Math 1.5B Instruct Q3 K_L"},
		{input: "qwen2-math:1.5b-instruct-q4_0", expected: "Qwen2 Math 1.5B Instruct Q4_0"},
		{input: "qwen2-math:1.5b-instruct-q4_1", expected: "Qwen2 Math 1.5B Instruct Q4 1"},
		{input: "qwen2-math:1.5b-instruct-q4_K_S", expected: "Qwen2 Math 1.5B Instruct Q4 K_S"},
		{input: "qwen2-math:1.5b-instruct-q4_K_M", expected: "Qwen2 Math 1.5B Instruct Q4_K_M"},
		{input: "qwen2-math:1.5b-instruct-q5_0", expected: "Qwen2 Math 1.5B Instruct Q5 0"},
		{input: "qwen2-math:1.5b-instruct-q5_1", expected: "Qwen2 Math 1.5B Instruct Q5 1"},
		{input: "qwen2-math:1.5b-instruct-q5_K_S", expected: "Qwen2 Math 1.5B Instruct Q5 K_S"},
		{input: "qwen2-math:1.5b-instruct-q5_K_M", expected: "Qwen2 Math 1.5B Instruct Q5 K_M"},
		{input: "qwen2-math:1.5b-instruct-q6_K", expected: "Qwen2 Math 1.5B Instruct Q6 K"},
		{input: "qwen2-math:1.5b-instruct-q8_0", expected: "Qwen2 Math 1.5B Instruct Q8_0"},
		{input: "qwen2-math:1.5b-instruct-fp16", expected: "Qwen2 Math 1.5B Instruct FP16"},
		{input: "qwen2-math:7b-instruct", expected: "Qwen2 Math 7B Instruct"},
		{input: "qwen2-math:7b-instruct-q2_K", expected: "Qwen2 Math 7B Instruct Q2_K"},
		{input: "qwen2-math:7b-instruct-q3_K_S", expected: "Qwen2 Math 7B Instruct Q3 K_S"},
		{input: "qwen2-math:7b-instruct-q3_K_M", expected: "Qwen2 Math 7B Instruct Q3 K_M"},
		{input: "qwen2-math:7b-instruct-q3_K_L", expected: "Qwen2 Math 7B Instruct Q3 K_L"},
		{input: "qwen2-math:7b-instruct-q4_0", expected: "Qwen2 Math 7B Instruct Q4_0"},
		{input: "qwen2-math:7b-instruct-q4_1", expected: "Qwen2 Math 7B Instruct Q4 1"},
		{input: "qwen2-math:7b-instruct-q4_K_S", expected: "Qwen2 Math 7B Instruct Q4 K_S"},
		{input: "qwen2-math:7b-instruct-q4_K_M", expected: "Qwen2 Math 7B Instruct Q4_K_M"},
		{input: "qwen2-math:7b-instruct-q5_0", expected: "Qwen2 Math 7B Instruct Q5 0"},
		{input: "qwen2-math:7b-instruct-q5_1", expected: "Qwen2 Math 7B Instruct Q5 1"},
		{input: "qwen2-math:7b-instruct-q5_K_S", expected: "Qwen2 Math 7B Instruct Q5 K_S"},
		{input: "qwen2-math:7b-instruct-q5_K_M", expected: "Qwen2 Math 7B Instruct Q5 K_M"},
		{input: "qwen2-math:7b-instruct-q6_K", expected: "Qwen2 Math 7B Instruct Q6 K"},
		{input: "qwen2-math:7b-instruct-q8_0", expected: "Qwen2 Math 7B Instruct Q8_0"},
		{input: "qwen2-math:7b-instruct-fp16", expected: "Qwen2 Math 7B Instruct FP16"},
		{input: "qwen2-math:72b-instruct", expected: "Qwen2 Math 72B Instruct"},
		{input: "qwen2-math:72b-instruct-q2_K", expected: "Qwen2 Math 72B Instruct Q2_K"},
		{input: "qwen2-math:72b-instruct-q3_K_S", expected: "Qwen2 Math 72B Instruct Q3 K_S"},
		{input: "qwen2-math:72b-instruct-q3_K_M", expected: "Qwen2 Math 72B Instruct Q3 K_M"},
		{input: "qwen2-math:72b-instruct-q3_K_L", expected: "Qwen2 Math 72B Instruct Q3 K_L"},
		{input: "qwen2-math:72b-instruct-q4_0", expected: "Qwen2 Math 72B Instruct Q4_0"},
		{input: "qwen2-math:72b-instruct-q4_1", expected: "Qwen2 Math 72B Instruct Q4 1"},
		{input: "qwen2-math:72b-instruct-q4_K_S", expected: "Qwen2 Math 72B Instruct Q4 K_S"},
		{input: "qwen2-math:72b-instruct-q4_K_M", expected: "Qwen2 Math 72B Instruct Q4_K_M"},
		{input: "qwen2-math:72b-instruct-q5_0", expected: "Qwen2 Math 72B Instruct Q5 0"},
		{input: "qwen2-math:72b-instruct-q5_1", expected: "Qwen2 Math 72B Instruct Q5 1"},
		{input: "qwen2-math:72b-instruct-q5_K_S", expected: "Qwen2 Math 72B Instruct Q5 K_S"},
		{input: "qwen2-math:72b-instruct-q5_K_M", expected: "Qwen2 Math 72B Instruct Q5 K_M"},
		{input: "qwen2-math:72b-instruct-q6_K", expected: "Qwen2 Math 72B Instruct Q6 K"},
		{input: "qwen2-math:72b-instruct-q8_0", expected: "Qwen2 Math 72B Instruct Q8_0"},
		{input: "qwen2-math:72b-instruct-fp16", expected: "Qwen2 Math 72B Instruct FP16"},
		{input: "qwen3-embedding:latest", expected: "Qwen3 Embedding (latest)"},
		{input: "qwen3-embedding", expected: "Qwen3 Embedding"},
		{input: "qwen3-embedding:0.6b", expected: "Qwen3 Embedding 0.6B"},
		{input: "qwen3-embedding:4b", expected: "Qwen3 Embedding 4B"},
		{input: "qwen3-embedding:8b", expected: "Qwen3 Embedding 8B"},
		{input: "qwen3-embedding:0.6b-q8_0", expected: "Qwen3 Embedding 0.6B Q8_0"},
		{input: "qwen3-embedding:0.6b-fp16", expected: "Qwen3 Embedding 0.6B FP16"},
		{input: "qwen3-embedding:4b-q4_K_M", expected: "Qwen3 Embedding 4B Q4_K_M"},
		{input: "qwen3-embedding:4b-q8_0", expected: "Qwen3 Embedding 4B Q8_0"},
		{input: "qwen3-embedding:4b-fp16", expected: "Qwen3 Embedding 4B FP16"},
		{input: "qwen3-embedding:8b-q4_K_M", expected: "Qwen3 Embedding 8B Q4_K_M"},
		{input: "qwen3-embedding:8b-q8_0", expected: "Qwen3 Embedding 8B Q8_0"},
		{input: "qwen3-embedding:8b-fp16", expected: "Qwen3 Embedding 8B FP16"},
		{input: "tinydolphin:latest", expected: "Tinydolphin (latest)"},
		{input: "tinydolphin", expected: "Tinydolphin"},
		{input: "tinydolphin:v2.8", expected: "Tinydolphin v2.8"},
		{input: "tinydolphin:1.1b", expected: "Tinydolphin 1.1B"},
		{input: "tinydolphin:1.1b-v2.8-q2_K", expected: "Tinydolphin 1.1B v2.8 Q2_K"},
		{input: "tinydolphin:1.1b-v2.8-q3_K_S", expected: "Tinydolphin 1.1B v2.8 Q3 K_S"},
		{input: "tinydolphin:1.1b-v2.8-q3_K_M", expected: "Tinydolphin 1.1B v2.8 Q3 K_M"},
		{input: "tinydolphin:1.1b-v2.8-q3_K_L", expected: "Tinydolphin 1.1B v2.8 Q3 K_L"},
		{input: "tinydolphin:1.1b-v2.8-q4_0", expected: "Tinydolphin 1.1B v2.8 Q4_0"},
		{input: "tinydolphin:1.1b-v2.8-q4_1", expected: "Tinydolphin 1.1B v2.8 Q4 1"},
		{input: "tinydolphin:1.1b-v2.8-q4_K_S", expected: "Tinydolphin 1.1B v2.8 Q4 K_S"},
		{input: "tinydolphin:1.1b-v2.8-q4_K_M", expected: "Tinydolphin 1.1B v2.8 Q4_K_M"},
		{input: "tinydolphin:1.1b-v2.8-q5_0", expected: "Tinydolphin 1.1B v2.8 Q5 0"},
		{input: "tinydolphin:1.1b-v2.8-q5_1", expected: "Tinydolphin 1.1B v2.8 Q5 1"},
		{input: "tinydolphin:1.1b-v2.8-q5_K_S", expected: "Tinydolphin 1.1B v2.8 Q5 K_S"},
		{input: "tinydolphin:1.1b-v2.8-q5_K_M", expected: "Tinydolphin 1.1B v2.8 Q5 K_M"},
		{input: "tinydolphin:1.1b-v2.8-q6_K", expected: "Tinydolphin 1.1B v2.8 Q6 K"},
		{input: "tinydolphin:1.1b-v2.8-q8_0", expected: "Tinydolphin 1.1B v2.8 Q8_0"},
		{input: "tinydolphin:1.1b-v2.8-fp16", expected: "Tinydolphin 1.1B v2.8 FP16"},
		{input: "aya:latest", expected: "Aya (latest)"},
		{input: "aya", expected: "Aya"},
		{input: "aya:8b", expected: "Aya 8B"},
		{input: "aya:35b", expected: "Aya 35B"},
		{input: "aya:8b-23", expected: "Aya 8B 23"},
		{input: "aya:8b-23-q2_K", expected: "Aya 8B 23 Q2_K"},
		{input: "aya:8b-23-q3_K_S", expected: "Aya 8B 23 Q3 K_S"},
		{input: "aya:8b-23-q3_K_M", expected: "Aya 8B 23 Q3 K_M"},
		{input: "aya:8b-23-q3_K_L", expected: "Aya 8B 23 Q3 K_L"},
		{input: "aya:8b-23-q4_0", expected: "Aya 8B 23 Q4_0"},
		{input: "aya:8b-23-q4_1", expected: "Aya 8B 23 Q4 1"},
		{input: "aya:8b-23-q4_K_S", expected: "Aya 8B 23 Q4 K_S"},
		{input: "aya:8b-23-q4_K_M", expected: "Aya 8B 23 Q4_K_M"},
		{input: "aya:8b-23-q5_0", expected: "Aya 8B 23 Q5 0"},
		{input: "aya:8b-23-q5_1", expected: "Aya 8B 23 Q5 1"},
		{input: "aya:8b-23-q5_K_S", expected: "Aya 8B 23 Q5 K_S"},
		{input: "aya:8b-23-q5_K_M", expected: "Aya 8B 23 Q5 K_M"},
		{input: "aya:8b-23-q6_K", expected: "Aya 8B 23 Q6 K"},
		{input: "aya:8b-23-q8_0", expected: "Aya 8B 23 Q8_0"},
		{input: "aya:35b-23", expected: "Aya 35B 23"},
		{input: "aya:35b-23-q2_K", expected: "Aya 35B 23 Q2_K"},
		{input: "aya:35b-23-q3_K_S", expected: "Aya 35B 23 Q3 K_S"},
		{input: "aya:35b-23-q3_K_M", expected: "Aya 35B 23 Q3 K_M"},
		{input: "aya:35b-23-q3_K_L", expected: "Aya 35B 23 Q3 K_L"},
		{input: "aya:35b-23-q4_0", expected: "Aya 35B 23 Q4_0"},
		{input: "aya:35b-23-q4_1", expected: "Aya 35B 23 Q4 1"},
		{input: "aya:35b-23-q4_K_S", expected: "Aya 35B 23 Q4 K_S"},
		{input: "aya:35b-23-q4_K_M", expected: "Aya 35B 23 Q4_K_M"},
		{input: "aya:35b-23-q5_0", expected: "Aya 35B 23 Q5 0"},
		{input: "aya:35b-23-q5_1", expected: "Aya 35B 23 Q5 1"},
		{input: "aya:35b-23-q5_K_S", expected: "Aya 35B 23 Q5 K_S"},
		{input: "aya:35b-23-q5_K_M", expected: "Aya 35B 23 Q5 K_M"},
		{input: "aya:35b-23-q6_K", expected: "Aya 35B 23 Q6 K"},
		{input: "aya:35b-23-q8_0", expected: "Aya 35B 23 Q8_0"},
		{input: "glm4:latest", expected: "GLM4 (latest)"},
		{input: "glm4", expected: "GLM4"},
		{input: "glm4:9b", expected: "GLM4 9B"},
		{input: "glm4:9b-chat-q2_K", expected: "GLM4 9B Chat Q2_K"},
		{input: "glm4:9b-chat-q3_K_S", expected: "GLM4 9B Chat Q3 K_S"},
		{input: "glm4:9b-chat-q3_K_M", expected: "GLM4 9B Chat Q3 K_M"},
		{input: "glm4:9b-chat-q3_K_L", expected: "GLM4 9B Chat Q3 K_L"},
		{input: "glm4:9b-chat-q4_0", expected: "GLM4 9B Chat Q4_0"},
		{input: "glm4:9b-chat-q4_1", expected: "GLM4 9B Chat Q4 1"},
		{input: "glm4:9b-chat-q4_K_S", expected: "GLM4 9B Chat Q4 K_S"},
		{input: "glm4:9b-chat-q4_K_M", expected: "GLM4 9B Chat Q4_K_M"},
		{input: "glm4:9b-chat-q5_0", expected: "GLM4 9B Chat Q5 0"},
		{input: "glm4:9b-chat-q5_1", expected: "GLM4 9B Chat Q5 1"},
		{input: "glm4:9b-chat-q5_K_S", expected: "GLM4 9B Chat Q5 K_S"},
		{input: "glm4:9b-chat-q5_K_M", expected: "GLM4 9B Chat Q5 K_M"},
		{input: "glm4:9b-chat-q6_K", expected: "GLM4 9B Chat Q6 K"},
		{input: "glm4:9b-chat-q8_0", expected: "GLM4 9B Chat Q8_0"},
		{input: "glm4:9b-chat-fp16", expected: "GLM4 9B Chat FP16"},
		{input: "glm4:9b-text-q2_K", expected: "GLM4 9B Text Q2_K"},
		{input: "glm4:9b-text-q3_K_S", expected: "GLM4 9B Text Q3 K_S"},
		{input: "glm4:9b-text-q3_K_M", expected: "GLM4 9B Text Q3 K_M"},
		{input: "glm4:9b-text-q3_K_L", expected: "GLM4 9B Text Q3 K_L"},
		{input: "glm4:9b-text-q4_0", expected: "GLM4 9B Text Q4_0"},
		{input: "glm4:9b-text-q4_1", expected: "GLM4 9B Text Q4 1"},
		{input: "glm4:9b-text-q4_K_S", expected: "GLM4 9B Text Q4 K_S"},
		{input: "glm4:9b-text-q4_K_M", expected: "GLM4 9B Text Q4_K_M"},
		{input: "glm4:9b-text-q5_0", expected: "GLM4 9B Text Q5 0"},
		{input: "glm4:9b-text-q5_1", expected: "GLM4 9B Text Q5 1"},
		{input: "glm4:9b-text-q5_K_S", expected: "GLM4 9B Text Q5 K_S"},
		{input: "glm4:9b-text-q5_K_M", expected: "GLM4 9B Text Q5 K_M"},
		{input: "glm4:9b-text-q6_K", expected: "GLM4 9B Text Q6 K"},
		{input: "glm4:9b-text-q8_0", expected: "GLM4 9B Text Q8_0"},
		{input: "glm4:9b-text-fp16", expected: "GLM4 9B Text FP16"},
		{input: "llama2-chinese:latest", expected: "Llama2 Chinese (latest)"},
		{input: "llama2-chinese", expected: "Llama2 Chinese"},
		{input: "llama2-chinese:7b", expected: "Llama2 Chinese 7B"},
		{input: "llama2-chinese:13b", expected: "Llama2 Chinese 13B"},
		{input: "llama2-chinese:7b-chat", expected: "Llama2 Chinese 7B Chat"},
		{input: "llama2-chinese:7b-chat-q2_K", expected: "Llama2 Chinese 7B Chat Q2_K"},
		{input: "llama2-chinese:7b-chat-q3_K_S", expected: "Llama2 Chinese 7B Chat Q3 K_S"},
		{input: "llama2-chinese:7b-chat-q3_K_M", expected: "Llama2 Chinese 7B Chat Q3 K_M"},
		{input: "llama2-chinese:7b-chat-q3_K_L", expected: "Llama2 Chinese 7B Chat Q3 K_L"},
		{input: "llama2-chinese:7b-chat-q4_0", expected: "Llama2 Chinese 7B Chat Q4_0"},
		{input: "llama2-chinese:7b-chat-q4_1", expected: "Llama2 Chinese 7B Chat Q4 1"},
		{input: "llama2-chinese:7b-chat-q4_K_S", expected: "Llama2 Chinese 7B Chat Q4 K_S"},
		{input: "llama2-chinese:7b-chat-q4_K_M", expected: "Llama2 Chinese 7B Chat Q4_K_M"},
		{input: "llama2-chinese:7b-chat-q5_0", expected: "Llama2 Chinese 7B Chat Q5 0"},
		{input: "llama2-chinese:7b-chat-q5_1", expected: "Llama2 Chinese 7B Chat Q5 1"},
		{input: "llama2-chinese:7b-chat-q5_K_S", expected: "Llama2 Chinese 7B Chat Q5 K_S"},
		{input: "llama2-chinese:7b-chat-q5_K_M", expected: "Llama2 Chinese 7B Chat Q5 K_M"},
		{input: "llama2-chinese:7b-chat-q6_K", expected: "Llama2 Chinese 7B Chat Q6 K"},
		{input: "llama2-chinese:7b-chat-q8_0", expected: "Llama2 Chinese 7B Chat Q8_0"},
		{input: "llama2-chinese:7b-chat-fp16", expected: "Llama2 Chinese 7B Chat FP16"},
		{input: "llama2-chinese:13b-chat", expected: "Llama2 Chinese 13B Chat"},
		{input: "llama2-chinese:13b-chat-q2_K", expected: "Llama2 Chinese 13B Chat Q2_K"},
		{input: "llama2-chinese:13b-chat-q3_K_S", expected: "Llama2 Chinese 13B Chat Q3 K_S"},
		{input: "llama2-chinese:13b-chat-q3_K_M", expected: "Llama2 Chinese 13B Chat Q3 K_M"},
		{input: "llama2-chinese:13b-chat-q3_K_L", expected: "Llama2 Chinese 13B Chat Q3 K_L"},
		{input: "llama2-chinese:13b-chat-q4_0", expected: "Llama2 Chinese 13B Chat Q4_0"},
		{input: "llama2-chinese:13b-chat-q4_1", expected: "Llama2 Chinese 13B Chat Q4 1"},
		{input: "llama2-chinese:13b-chat-q4_K_S", expected: "Llama2 Chinese 13B Chat Q4 K_S"},
		{input: "llama2-chinese:13b-chat-q4_K_M", expected: "Llama2 Chinese 13B Chat Q4_K_M"},
		{input: "llama2-chinese:13b-chat-q5_0", expected: "Llama2 Chinese 13B Chat Q5 0"},
		{input: "llama2-chinese:13b-chat-q5_1", expected: "Llama2 Chinese 13B Chat Q5 1"},
		{input: "llama2-chinese:13b-chat-q5_K_S", expected: "Llama2 Chinese 13B Chat Q5 K_S"},
		{input: "llama2-chinese:13b-chat-q5_K_M", expected: "Llama2 Chinese 13B Chat Q5 K_M"},
		{input: "llama2-chinese:13b-chat-q6_K", expected: "Llama2 Chinese 13B Chat Q6 K"},
		{input: "llama2-chinese:13b-chat-q8_0", expected: "Llama2 Chinese 13B Chat Q8_0"},
		{input: "llama2-chinese:13b-chat-fp16", expected: "Llama2 Chinese 13B Chat FP16"},
		{input: "granite3.2:latest", expected: "Granite3.2 (latest)"},
		{input: "granite3.2", expected: "Granite3.2"},
		{input: "granite3.2:2b", expected: "Granite3.2 2B"},
		{input: "granite3.2:8b", expected: "Granite3.2 8B"},
		{input: "granite3.2:2b-instruct-q4_K_M", expected: "Granite3.2 2B Instruct Q4_K_M"},
		{input: "granite3.2:2b-instruct-q8_0", expected: "Granite3.2 2B Instruct Q8_0"},
		{input: "granite3.2:2b-instruct-fp16", expected: "Granite3.2 2B Instruct FP16"},
		{input: "granite3.2:8b-instruct-q4_K_M", expected: "Granite3.2 8B Instruct Q4_K_M"},
		{input: "granite3.2:8b-instruct-q8_0", expected: "Granite3.2 8B Instruct Q8_0"},
		{input: "granite3.2:8b-instruct-fp16", expected: "Granite3.2 8B Instruct FP16"},
		{input: "paraphrase-multilingual:latest", expected: "Paraphrase Multilingual (latest)"},
		{input: "paraphrase-multilingual", expected: "Paraphrase Multilingual"},
		{input: "paraphrase-multilingual:278m", expected: "Paraphrase Multilingual 278M"},
		{input: "paraphrase-multilingual:278m-mpnet-base-v2-fp16", expected: "Paraphrase Multilingual 278M Mpnet Base v2 FP16"},
		{input: "stable-code:latest", expected: "Stable Code (latest)"},
		{input: "stable-code", expected: "Stable Code"},
		{input: "stable-code:code", expected: "Stable Code Code"},
		{input: "stable-code:instruct", expected: "Stable Code Instruct"},
		{input: "stable-code:3b", expected: "Stable Code 3B"},
		{input: "stable-code:3b-code", expected: "Stable Code 3B Code"},
		{input: "stable-code:3b-code-q2_K", expected: "Stable Code 3B Code Q2_K"},
		{input: "stable-code:3b-code-q3_K_S", expected: "Stable Code 3B Code Q3 K_S"},
		{input: "stable-code:3b-code-q3_K_M", expected: "Stable Code 3B Code Q3 K_M"},
		{input: "stable-code:3b-code-q3_K_L", expected: "Stable Code 3B Code Q3 K_L"},
		{input: "stable-code:3b-code-q4_0", expected: "Stable Code 3B Code Q4_0"},
		{input: "stable-code:3b-code-q4_1", expected: "Stable Code 3B Code Q4 1"},
		{input: "stable-code:3b-code-q4_K_S", expected: "Stable Code 3B Code Q4 K_S"},
		{input: "stable-code:3b-code-q4_K_M", expected: "Stable Code 3B Code Q4_K_M"},
		{input: "stable-code:3b-code-q5_0", expected: "Stable Code 3B Code Q5 0"},
		{input: "stable-code:3b-code-q5_1", expected: "Stable Code 3B Code Q5 1"},
		{input: "stable-code:3b-code-q5_K_S", expected: "Stable Code 3B Code Q5 K_S"},
		{input: "stable-code:3b-code-q5_K_M", expected: "Stable Code 3B Code Q5 K_M"},
		{input: "stable-code:3b-code-q6_K", expected: "Stable Code 3B Code Q6 K"},
		{input: "stable-code:3b-code-q8_0", expected: "Stable Code 3B Code Q8_0"},
		{input: "stable-code:3b-code-fp16", expected: "Stable Code 3B Code FP16"},
		{input: "stable-code:3b-instruct", expected: "Stable Code 3B Instruct"},
		{input: "stable-code:3b-instruct-q2_K", expected: "Stable Code 3B Instruct Q2_K"},
		{input: "stable-code:3b-instruct-q3_K_S", expected: "Stable Code 3B Instruct Q3 K_S"},
		{input: "stable-code:3b-instruct-q3_K_M", expected: "Stable Code 3B Instruct Q3 K_M"},
		{input: "stable-code:3b-instruct-q3_K_L", expected: "Stable Code 3B Instruct Q3 K_L"},
		{input: "stable-code:3b-instruct-q4_0", expected: "Stable Code 3B Instruct Q4_0"},
		{input: "stable-code:3b-instruct-q4_1", expected: "Stable Code 3B Instruct Q4 1"},
		{input: "stable-code:3b-instruct-q4_K_S", expected: "Stable Code 3B Instruct Q4 K_S"},
		{input: "stable-code:3b-instruct-q4_K_M", expected: "Stable Code 3B Instruct Q4_K_M"},
		{input: "stable-code:3b-instruct-q5_0", expected: "Stable Code 3B Instruct Q5 0"},
		{input: "stable-code:3b-instruct-q5_1", expected: "Stable Code 3B Instruct Q5 1"},
		{input: "stable-code:3b-instruct-q5_K_S", expected: "Stable Code 3B Instruct Q5 K_S"},
		{input: "stable-code:3b-instruct-q5_K_M", expected: "Stable Code 3B Instruct Q5 K_M"},
		{input: "stable-code:3b-instruct-q6_K", expected: "Stable Code 3B Instruct Q6 K"},
		{input: "stable-code:3b-instruct-q8_0", expected: "Stable Code 3B Instruct Q8_0"},
		{input: "stable-code:3b-instruct-fp16", expected: "Stable Code 3B Instruct FP16"},
		{input: "neural-chat:latest", expected: "Neural Chat (latest)"},
		{input: "neural-chat", expected: "Neural Chat"},
		{input: "neural-chat:7b", expected: "Neural Chat 7B"},
		{input: "neural-chat:7b-v3.1", expected: "Neural Chat 7B v3.1"},
		{input: "neural-chat:7b-v3.1-q2_K", expected: "Neural Chat 7B v3.1 Q2_K"},
		{input: "neural-chat:7b-v3.1-q3_K_S", expected: "Neural Chat 7B v3.1 Q3 K_S"},
		{input: "neural-chat:7b-v3.1-q3_K_M", expected: "Neural Chat 7B v3.1 Q3 K_M"},
		{input: "neural-chat:7b-v3.1-q3_K_L", expected: "Neural Chat 7B v3.1 Q3 K_L"},
		{input: "neural-chat:7b-v3.1-q4_0", expected: "Neural Chat 7B v3.1 Q4_0"},
		{input: "neural-chat:7b-v3.1-q4_1", expected: "Neural Chat 7B v3.1 Q4 1"},
		{input: "neural-chat:7b-v3.1-q4_K_S", expected: "Neural Chat 7B v3.1 Q4 K_S"},
		{input: "neural-chat:7b-v3.1-q4_K_M", expected: "Neural Chat 7B v3.1 Q4_K_M"},
		{input: "neural-chat:7b-v3.1-q5_0", expected: "Neural Chat 7B v3.1 Q5 0"},
		{input: "neural-chat:7b-v3.1-q5_1", expected: "Neural Chat 7B v3.1 Q5 1"},
		{input: "neural-chat:7b-v3.1-q5_K_S", expected: "Neural Chat 7B v3.1 Q5 K_S"},
		{input: "neural-chat:7b-v3.1-q5_K_M", expected: "Neural Chat 7B v3.1 Q5 K_M"},
		{input: "neural-chat:7b-v3.1-q6_K", expected: "Neural Chat 7B v3.1 Q6 K"},
		{input: "neural-chat:7b-v3.1-q8_0", expected: "Neural Chat 7B v3.1 Q8_0"},
		{input: "neural-chat:7b-v3.1-fp16", expected: "Neural Chat 7B v3.1 FP16"},
		{input: "neural-chat:7b-v3.2", expected: "Neural Chat 7B v3.2"},
		{input: "neural-chat:7b-v3.2-q2_K", expected: "Neural Chat 7B v3.2 Q2_K"},
		{input: "neural-chat:7b-v3.2-q3_K_S", expected: "Neural Chat 7B v3.2 Q3 K_S"},
		{input: "neural-chat:7b-v3.2-q3_K_M", expected: "Neural Chat 7B v3.2 Q3 K_M"},
		{input: "neural-chat:7b-v3.2-q3_K_L", expected: "Neural Chat 7B v3.2 Q3 K_L"},
		{input: "neural-chat:7b-v3.2-q4_0", expected: "Neural Chat 7B v3.2 Q4_0"},
		{input: "neural-chat:7b-v3.2-q4_1", expected: "Neural Chat 7B v3.2 Q4 1"},
		{input: "neural-chat:7b-v3.2-q4_K_S", expected: "Neural Chat 7B v3.2 Q4 K_S"},
		{input: "neural-chat:7b-v3.2-q4_K_M", expected: "Neural Chat 7B v3.2 Q4_K_M"},
		{input: "neural-chat:7b-v3.2-q5_0", expected: "Neural Chat 7B v3.2 Q5 0"},
		{input: "neural-chat:7b-v3.2-q5_1", expected: "Neural Chat 7B v3.2 Q5 1"},
		{input: "neural-chat:7b-v3.2-q5_K_S", expected: "Neural Chat 7B v3.2 Q5 K_S"},
		{input: "neural-chat:7b-v3.2-q5_K_M", expected: "Neural Chat 7B v3.2 Q5 K_M"},
		{input: "neural-chat:7b-v3.2-q6_K", expected: "Neural Chat 7B v3.2 Q6 K"},
		{input: "neural-chat:7b-v3.2-q8_0", expected: "Neural Chat 7B v3.2 Q8_0"},
		{input: "neural-chat:7b-v3.2-fp16", expected: "Neural Chat 7B v3.2 FP16"},
		{input: "neural-chat:7b-v3.3", expected: "Neural Chat 7B v3.3"},
		{input: "neural-chat:7b-v3.3-q2_K", expected: "Neural Chat 7B v3.3 Q2_K"},
		{input: "neural-chat:7b-v3.3-q3_K_S", expected: "Neural Chat 7B v3.3 Q3 K_S"},
		{input: "neural-chat:7b-v3.3-q3_K_M", expected: "Neural Chat 7B v3.3 Q3 K_M"},
		{input: "neural-chat:7b-v3.3-q3_K_L", expected: "Neural Chat 7B v3.3 Q3 K_L"},
		{input: "neural-chat:7b-v3.3-q4_0", expected: "Neural Chat 7B v3.3 Q4_0"},
		{input: "neural-chat:7b-v3.3-q4_1", expected: "Neural Chat 7B v3.3 Q4 1"},
		{input: "neural-chat:7b-v3.3-q4_K_S", expected: "Neural Chat 7B v3.3 Q4 K_S"},
		{input: "neural-chat:7b-v3.3-q4_K_M", expected: "Neural Chat 7B v3.3 Q4_K_M"},
		{input: "neural-chat:7b-v3.3-q5_0", expected: "Neural Chat 7B v3.3 Q5 0"},
		{input: "neural-chat:7b-v3.3-q5_1", expected: "Neural Chat 7B v3.3 Q5 1"},
		{input: "neural-chat:7b-v3.3-q5_K_S", expected: "Neural Chat 7B v3.3 Q5 K_S"},
		{input: "neural-chat:7b-v3.3-q5_K_M", expected: "Neural Chat 7B v3.3 Q5 K_M"},
		{input: "neural-chat:7b-v3.3-q6_K", expected: "Neural Chat 7B v3.3 Q6 K"},
		{input: "neural-chat:7b-v3.3-q8_0", expected: "Neural Chat 7B v3.3 Q8_0"},
		{input: "neural-chat:7b-v3.3-fp16", expected: "Neural Chat 7B v3.3 FP16"},
		{input: "nous-hermes2:latest", expected: "Nous Hermes2 (latest)"},
		{input: "nous-hermes2", expected: "Nous Hermes2"},
		{input: "nous-hermes2:10.7b", expected: "Nous Hermes2 10.7B"},
		{input: "nous-hermes2:34b", expected: "Nous Hermes2 34B"},
		{input: "nous-hermes2:10.7b-solar-q2_K", expected: "Nous Hermes2 10.7B Solar Q2_K"},
		{input: "nous-hermes2:10.7b-solar-q3_K_S", expected: "Nous Hermes2 10.7B Solar Q3 K_S"},
		{input: "nous-hermes2:10.7b-solar-q3_K_M", expected: "Nous Hermes2 10.7B Solar Q3 K_M"},
		{input: "nous-hermes2:10.7b-solar-q3_K_L", expected: "Nous Hermes2 10.7B Solar Q3 K_L"},
		{input: "nous-hermes2:10.7b-solar-q4_0", expected: "Nous Hermes2 10.7B Solar Q4_0"},
		{input: "nous-hermes2:10.7b-solar-q4_1", expected: "Nous Hermes2 10.7B Solar Q4 1"},
		{input: "nous-hermes2:10.7b-solar-q4_K_S", expected: "Nous Hermes2 10.7B Solar Q4 K_S"},
		{input: "nous-hermes2:10.7b-solar-q4_K_M", expected: "Nous Hermes2 10.7B Solar Q4_K_M"},
		{input: "nous-hermes2:10.7b-solar-q5_0", expected: "Nous Hermes2 10.7B Solar Q5 0"},
		{input: "nous-hermes2:10.7b-solar-q5_1", expected: "Nous Hermes2 10.7B Solar Q5 1"},
		{input: "nous-hermes2:10.7b-solar-q5_K_S", expected: "Nous Hermes2 10.7B Solar Q5 K_S"},
		{input: "nous-hermes2:10.7b-solar-q5_K_M", expected: "Nous Hermes2 10.7B Solar Q5 K_M"},
		{input: "nous-hermes2:10.7b-solar-q6_K", expected: "Nous Hermes2 10.7B Solar Q6 K"},
		{input: "nous-hermes2:10.7b-solar-q8_0", expected: "Nous Hermes2 10.7B Solar Q8_0"},
		{input: "nous-hermes2:10.7b-solar-fp16", expected: "Nous Hermes2 10.7B Solar FP16"},
		{input: "nous-hermes2:34b-yi-q2_K", expected: "Nous Hermes2 34B Yi Q2_K"},
		{input: "nous-hermes2:34b-yi-q3_K_S", expected: "Nous Hermes2 34B Yi Q3 K_S"},
		{input: "nous-hermes2:34b-yi-q3_K_M", expected: "Nous Hermes2 34B Yi Q3 K_M"},
		{input: "nous-hermes2:34b-yi-q3_K_L", expected: "Nous Hermes2 34B Yi Q3 K_L"},
		{input: "nous-hermes2:34b-yi-q4_0", expected: "Nous Hermes2 34B Yi Q4_0"},
		{input: "nous-hermes2:34b-yi-q4_1", expected: "Nous Hermes2 34B Yi Q4 1"},
		{input: "nous-hermes2:34b-yi-q4_K_S", expected: "Nous Hermes2 34B Yi Q4 K_S"},
		{input: "nous-hermes2:34b-yi-q4_K_M", expected: "Nous Hermes2 34B Yi Q4_K_M"},
		{input: "nous-hermes2:34b-yi-q5_0", expected: "Nous Hermes2 34B Yi Q5 0"},
		{input: "nous-hermes2:34b-yi-q5_1", expected: "Nous Hermes2 34B Yi Q5 1"},
		{input: "nous-hermes2:34b-yi-q5_K_S", expected: "Nous Hermes2 34B Yi Q5 K_S"},
		{input: "nous-hermes2:34b-yi-q5_K_M", expected: "Nous Hermes2 34B Yi Q5 K_M"},
		{input: "nous-hermes2:34b-yi-q6_K", expected: "Nous Hermes2 34B Yi Q6 K"},
		{input: "nous-hermes2:34b-yi-q8_0", expected: "Nous Hermes2 34B Yi Q8_0"},
		{input: "nous-hermes2:34b-yi-fp16", expected: "Nous Hermes2 34B Yi FP16"},
		{input: "bakllava:latest", expected: "Bakllava (latest)"},
		{input: "bakllava", expected: "Bakllava"},
		{input: "bakllava:7b", expected: "Bakllava 7B"},
		{input: "bakllava:7b-v1-q2_K", expected: "Bakllava 7B v1 Q2_K"},
		{input: "bakllava:7b-v1-q3_K_S", expected: "Bakllava 7B v1 Q3 K_S"},
		{input: "bakllava:7b-v1-q3_K_M", expected: "Bakllava 7B v1 Q3 K_M"},
		{input: "bakllava:7b-v1-q3_K_L", expected: "Bakllava 7B v1 Q3 K_L"},
		{input: "bakllava:7b-v1-q4_0", expected: "Bakllava 7B v1 Q4_0"},
		{input: "bakllava:7b-v1-q4_1", expected: "Bakllava 7B v1 Q4 1"},
		{input: "bakllava:7b-v1-q4_K_S", expected: "Bakllava 7B v1 Q4 K_S"},
		{input: "bakllava:7b-v1-q4_K_M", expected: "Bakllava 7B v1 Q4_K_M"},
		{input: "bakllava:7b-v1-q5_0", expected: "Bakllava 7B v1 Q5 0"},
		{input: "bakllava:7b-v1-q5_1", expected: "Bakllava 7B v1 Q5 1"},
		{input: "bakllava:7b-v1-q5_K_S", expected: "Bakllava 7B v1 Q5 K_S"},
		{input: "bakllava:7b-v1-q5_K_M", expected: "Bakllava 7B v1 Q5 K_M"},
		{input: "bakllava:7b-v1-q6_K", expected: "Bakllava 7B v1 Q6 K"},
		{input: "bakllava:7b-v1-q8_0", expected: "Bakllava 7B v1 Q8_0"},
		{input: "bakllava:7b-v1-fp16", expected: "Bakllava 7B v1 FP16"},
		{input: "wizardcoder:latest", expected: "Wizardcoder (latest)"},
		{input: "wizardcoder", expected: "Wizardcoder"},
		{input: "wizardcoder:python", expected: "Wizardcoder Python"},
		{input: "wizardcoder:33b", expected: "Wizardcoder 33B"},
		{input: "wizardcoder:7b-python", expected: "Wizardcoder 7B Python"},
		{input: "wizardcoder:7b-python-q2_K", expected: "Wizardcoder 7B Python Q2_K"},
		{input: "wizardcoder:7b-python-q3_K_S", expected: "Wizardcoder 7B Python Q3 K_S"},
		{input: "wizardcoder:7b-python-q3_K_M", expected: "Wizardcoder 7B Python Q3 K_M"},
		{input: "wizardcoder:7b-python-q3_K_L", expected: "Wizardcoder 7B Python Q3 K_L"},
		{input: "wizardcoder:7b-python-q4_0", expected: "Wizardcoder 7B Python Q4_0"},
		{input: "wizardcoder:7b-python-q4_1", expected: "Wizardcoder 7B Python Q4 1"},
		{input: "wizardcoder:7b-python-q4_K_S", expected: "Wizardcoder 7B Python Q4 K_S"},
		{input: "wizardcoder:7b-python-q4_K_M", expected: "Wizardcoder 7B Python Q4_K_M"},
		{input: "wizardcoder:7b-python-q5_0", expected: "Wizardcoder 7B Python Q5 0"},
		{input: "wizardcoder:7b-python-q5_1", expected: "Wizardcoder 7B Python Q5 1"},
		{input: "wizardcoder:7b-python-q5_K_S", expected: "Wizardcoder 7B Python Q5 K_S"},
		{input: "wizardcoder:7b-python-q5_K_M", expected: "Wizardcoder 7B Python Q5 K_M"},
		{input: "wizardcoder:7b-python-q6_K", expected: "Wizardcoder 7B Python Q6 K"},
		{input: "wizardcoder:7b-python-q8_0", expected: "Wizardcoder 7B Python Q8_0"},
		{input: "wizardcoder:7b-python-fp16", expected: "Wizardcoder 7B Python FP16"},
		{input: "wizardcoder:13b-python", expected: "Wizardcoder 13B Python"},
		{input: "wizardcoder:13b-python-q2_K", expected: "Wizardcoder 13B Python Q2_K"},
		{input: "wizardcoder:13b-python-q3_K_S", expected: "Wizardcoder 13B Python Q3 K_S"},
		{input: "wizardcoder:13b-python-q3_K_M", expected: "Wizardcoder 13B Python Q3 K_M"},
		{input: "wizardcoder:13b-python-q3_K_L", expected: "Wizardcoder 13B Python Q3 K_L"},
		{input: "wizardcoder:13b-python-q4_0", expected: "Wizardcoder 13B Python Q4_0"},
		{input: "wizardcoder:13b-python-q4_1", expected: "Wizardcoder 13B Python Q4 1"},
		{input: "wizardcoder:13b-python-q4_K_S", expected: "Wizardcoder 13B Python Q4 K_S"},
		{input: "wizardcoder:13b-python-q4_K_M", expected: "Wizardcoder 13B Python Q4_K_M"},
		{input: "wizardcoder:13b-python-q5_0", expected: "Wizardcoder 13B Python Q5 0"},
		{input: "wizardcoder:13b-python-q5_1", expected: "Wizardcoder 13B Python Q5 1"},
		{input: "wizardcoder:13b-python-q5_K_S", expected: "Wizardcoder 13B Python Q5 K_S"},
		{input: "wizardcoder:13b-python-q5_K_M", expected: "Wizardcoder 13B Python Q5 K_M"},
		{input: "wizardcoder:13b-python-q6_K", expected: "Wizardcoder 13B Python Q6 K"},
		{input: "wizardcoder:13b-python-q8_0", expected: "Wizardcoder 13B Python Q8_0"},
		{input: "wizardcoder:13b-python-fp16", expected: "Wizardcoder 13B Python FP16"},
		{input: "wizardcoder:33b-v1.1", expected: "Wizardcoder 33B v1.1"},
		{input: "wizardcoder:33b-v1.1-q2_K", expected: "Wizardcoder 33B v1.1 Q2_K"},
		{input: "wizardcoder:33b-v1.1-q3_K_S", expected: "Wizardcoder 33B v1.1 Q3 K_S"},
		{input: "wizardcoder:33b-v1.1-q3_K_M", expected: "Wizardcoder 33B v1.1 Q3 K_M"},
		{input: "wizardcoder:33b-v1.1-q3_K_L", expected: "Wizardcoder 33B v1.1 Q3 K_L"},
		{input: "wizardcoder:33b-v1.1-q4_0", expected: "Wizardcoder 33B v1.1 Q4_0"},
		{input: "wizardcoder:33b-v1.1-q4_1", expected: "Wizardcoder 33B v1.1 Q4 1"},
		{input: "wizardcoder:33b-v1.1-q4_K_S", expected: "Wizardcoder 33B v1.1 Q4 K_S"},
		{input: "wizardcoder:33b-v1.1-q4_K_M", expected: "Wizardcoder 33B v1.1 Q4_K_M"},
		{input: "wizardcoder:33b-v1.1-q5_0", expected: "Wizardcoder 33B v1.1 Q5 0"},
		{input: "wizardcoder:33b-v1.1-q5_1", expected: "Wizardcoder 33B v1.1 Q5 1"},
		{input: "wizardcoder:33b-v1.1-q5_K_S", expected: "Wizardcoder 33B v1.1 Q5 K_S"},
		{input: "wizardcoder:33b-v1.1-q5_K_M", expected: "Wizardcoder 33B v1.1 Q5 K_M"},
		{input: "wizardcoder:33b-v1.1-q6_K", expected: "Wizardcoder 33B v1.1 Q6 K"},
		{input: "wizardcoder:33b-v1.1-q8_0", expected: "Wizardcoder 33B v1.1 Q8_0"},
		{input: "wizardcoder:33b-v1.1-fp16", expected: "Wizardcoder 33B v1.1 FP16"},
		{input: "wizardcoder:34b-python", expected: "Wizardcoder 34B Python"},
		{input: "wizardcoder:34b-python-q2_K", expected: "Wizardcoder 34B Python Q2_K"},
		{input: "wizardcoder:34b-python-q3_K_S", expected: "Wizardcoder 34B Python Q3 K_S"},
		{input: "wizardcoder:34b-python-q3_K_M", expected: "Wizardcoder 34B Python Q3 K_M"},
		{input: "wizardcoder:34b-python-q3_K_L", expected: "Wizardcoder 34B Python Q3 K_L"},
		{input: "wizardcoder:34b-python-q4_0", expected: "Wizardcoder 34B Python Q4_0"},
		{input: "wizardcoder:34b-python-q4_1", expected: "Wizardcoder 34B Python Q4 1"},
		{input: "wizardcoder:34b-python-q4_K_S", expected: "Wizardcoder 34B Python Q4 K_S"},
		{input: "wizardcoder:34b-python-q4_K_M", expected: "Wizardcoder 34B Python Q4_K_M"},
		{input: "wizardcoder:34b-python-q5_0", expected: "Wizardcoder 34B Python Q5 0"},
		{input: "wizardcoder:34b-python-q5_1", expected: "Wizardcoder 34B Python Q5 1"},
		{input: "wizardcoder:34b-python-q5_K_S", expected: "Wizardcoder 34B Python Q5 K_S"},
		{input: "wizardcoder:34b-python-q5_K_M", expected: "Wizardcoder 34B Python Q5 K_M"},
		{input: "wizardcoder:34b-python-q6_K", expected: "Wizardcoder 34B Python Q6 K"},
		{input: "wizardcoder:34b-python-q8_0", expected: "Wizardcoder 34B Python Q8_0"},
		{input: "wizardcoder:34b-python-fp16", expected: "Wizardcoder 34B Python FP16"},
		{input: "sqlcoder:latest", expected: "Sqlcoder (latest)"},
		{input: "sqlcoder", expected: "Sqlcoder"},
		{input: "sqlcoder:7b", expected: "Sqlcoder 7B"},
		{input: "sqlcoder:15b", expected: "Sqlcoder 15B"},
		{input: "sqlcoder:7b-q2_K", expected: "Sqlcoder 7B Q2_K"},
		{input: "sqlcoder:7b-q3_K_S", expected: "Sqlcoder 7B Q3 K_S"},
		{input: "sqlcoder:7b-q3_K_M", expected: "Sqlcoder 7B Q3 K_M"},
		{input: "sqlcoder:7b-q3_K_L", expected: "Sqlcoder 7B Q3 K_L"},
		{input: "sqlcoder:7b-q4_0", expected: "Sqlcoder 7B Q4_0"},
		{input: "sqlcoder:7b-q4_1", expected: "Sqlcoder 7B Q4 1"},
		{input: "sqlcoder:7b-q4_K_S", expected: "Sqlcoder 7B Q4 K_S"},
		{input: "sqlcoder:7b-q4_K_M", expected: "Sqlcoder 7B Q4_K_M"},
		{input: "sqlcoder:7b-q5_0", expected: "Sqlcoder 7B Q5 0"},
		{input: "sqlcoder:7b-q5_1", expected: "Sqlcoder 7B Q5 1"},
		{input: "sqlcoder:7b-q5_K_S", expected: "Sqlcoder 7B Q5 K_S"},
		{input: "sqlcoder:7b-q5_K_M", expected: "Sqlcoder 7B Q5 K_M"},
		{input: "sqlcoder:7b-q6_K", expected: "Sqlcoder 7B Q6 K"},
		{input: "sqlcoder:7b-q8_0", expected: "Sqlcoder 7B Q8_0"},
		{input: "sqlcoder:7b-fp16", expected: "Sqlcoder 7B FP16"},
		{input: "sqlcoder:15b-q2_K", expected: "Sqlcoder 15B Q2_K"},
		{input: "sqlcoder:15b-q3_K_S", expected: "Sqlcoder 15B Q3 K_S"},
		{input: "sqlcoder:15b-q3_K_M", expected: "Sqlcoder 15B Q3 K_M"},
		{input: "sqlcoder:15b-q3_K_L", expected: "Sqlcoder 15B Q3 K_L"},
		{input: "sqlcoder:15b-q4_0", expected: "Sqlcoder 15B Q4_0"},
		{input: "sqlcoder:15b-q4_1", expected: "Sqlcoder 15B Q4 1"},
		{input: "sqlcoder:15b-q4_K_S", expected: "Sqlcoder 15B Q4 K_S"},
		{input: "sqlcoder:15b-q4_K_M", expected: "Sqlcoder 15B Q4_K_M"},
		{input: "sqlcoder:15b-q5_0", expected: "Sqlcoder 15B Q5 0"},
		{input: "sqlcoder:15b-q5_1", expected: "Sqlcoder 15B Q5 1"},
		{input: "sqlcoder:15b-q5_K_S", expected: "Sqlcoder 15B Q5 K_S"},
		{input: "sqlcoder:15b-q5_K_M", expected: "Sqlcoder 15B Q5 K_M"},
		{input: "sqlcoder:15b-q6_K", expected: "Sqlcoder 15B Q6 K"},
		{input: "sqlcoder:15b-q8_0", expected: "Sqlcoder 15B Q8_0"},
		{input: "sqlcoder:15b-fp16", expected: "Sqlcoder 15B FP16"},
		{input: "sqlcoder:70b-alpha-q2_K", expected: "Sqlcoder 70B Alpha Q2_K"},
		{input: "sqlcoder:70b-alpha-q3_K_S", expected: "Sqlcoder 70B Alpha Q3 K_S"},
		{input: "sqlcoder:70b-alpha-q3_K_M", expected: "Sqlcoder 70B Alpha Q3 K_M"},
		{input: "sqlcoder:70b-alpha-q3_K_L", expected: "Sqlcoder 70B Alpha Q3 K_L"},
		{input: "sqlcoder:70b-alpha-q4_0", expected: "Sqlcoder 70B Alpha Q4_0"},
		{input: "sqlcoder:70b-alpha-q4_1", expected: "Sqlcoder 70B Alpha Q4 1"},
		{input: "sqlcoder:70b-alpha-q4_K_S", expected: "Sqlcoder 70B Alpha Q4 K_S"},
		{input: "sqlcoder:70b-alpha-q4_K_M", expected: "Sqlcoder 70B Alpha Q4_K_M"},
		{input: "sqlcoder:70b-alpha-q5_0", expected: "Sqlcoder 70B Alpha Q5 0"},
		{input: "sqlcoder:70b-alpha-q5_1", expected: "Sqlcoder 70B Alpha Q5 1"},
		{input: "sqlcoder:70b-alpha-q5_K_S", expected: "Sqlcoder 70B Alpha Q5 K_S"},
		{input: "sqlcoder:70b-alpha-q5_K_M", expected: "Sqlcoder 70B Alpha Q5 K_M"},
		{input: "sqlcoder:70b-alpha-q6_K", expected: "Sqlcoder 70B Alpha Q6 K"},
		{input: "sqlcoder:70b-alpha-q8_0", expected: "Sqlcoder 70B Alpha Q8_0"},
		{input: "sqlcoder:70b-alpha-fp16", expected: "Sqlcoder 70B Alpha FP16"},
		{input: "bge-large:latest", expected: "Bge Large (latest)"},
		{input: "bge-large", expected: "Bge Large"},
		{input: "bge-large:335m", expected: "Bge Large 335M"},
		{input: "bge-large:335m-en-v1.5-fp16", expected: "Bge Large 335M En v1.5 FP16"},
		{input: "stablelm2:latest", expected: "StableLM2 (latest)"},
		{input: "stablelm2", expected: "StableLM2"},
		{input: "stablelm2:chat", expected: "StableLM2 Chat"},
		{input: "stablelm2:zephyr", expected: "StableLM2 Zephyr"},
		{input: "stablelm2:1.6b", expected: "StableLM2 1.6B"},
		{input: "stablelm2:12b", expected: "StableLM2 12B"},
		{input: "stablelm2:1.6b-chat", expected: "StableLM2 1.6B Chat"},
		{input: "stablelm2:1.6b-chat-q2_K", expected: "StableLM2 1.6B Chat Q2_K"},
		{input: "stablelm2:1.6b-chat-q3_K_S", expected: "StableLM2 1.6B Chat Q3 K_S"},
		{input: "stablelm2:1.6b-chat-q3_K_M", expected: "StableLM2 1.6B Chat Q3 K_M"},
		{input: "stablelm2:1.6b-chat-q3_K_L", expected: "StableLM2 1.6B Chat Q3 K_L"},
		{input: "stablelm2:1.6b-chat-q4_0", expected: "StableLM2 1.6B Chat Q4_0"},
		{input: "stablelm2:1.6b-chat-q4_1", expected: "StableLM2 1.6B Chat Q4 1"},
		{input: "stablelm2:1.6b-chat-q4_K_S", expected: "StableLM2 1.6B Chat Q4 K_S"},
		{input: "stablelm2:1.6b-chat-q4_K_M", expected: "StableLM2 1.6B Chat Q4_K_M"},
		{input: "stablelm2:1.6b-chat-q5_0", expected: "StableLM2 1.6B Chat Q5 0"},
		{input: "stablelm2:1.6b-chat-q5_1", expected: "StableLM2 1.6B Chat Q5 1"},
		{input: "stablelm2:1.6b-chat-q5_K_S", expected: "StableLM2 1.6B Chat Q5 K_S"},
		{input: "stablelm2:1.6b-chat-q5_K_M", expected: "StableLM2 1.6B Chat Q5 K_M"},
		{input: "stablelm2:1.6b-chat-q6_K", expected: "StableLM2 1.6B Chat Q6 K"},
		{input: "stablelm2:1.6b-chat-q8_0", expected: "StableLM2 1.6B Chat Q8_0"},
		{input: "stablelm2:1.6b-chat-fp16", expected: "StableLM2 1.6B Chat FP16"},
		{input: "stablelm2:1.6b-zephyr", expected: "StableLM2 1.6B Zephyr"},
		{input: "stablelm2:1.6b-zephyr-q2_K", expected: "StableLM2 1.6B Zephyr Q2_K"},
		{input: "stablelm2:1.6b-zephyr-q3_K_S", expected: "StableLM2 1.6B Zephyr Q3 K_S"},
		{input: "stablelm2:1.6b-zephyr-q3_K_M", expected: "StableLM2 1.6B Zephyr Q3 K_M"},
		{input: "stablelm2:1.6b-zephyr-q3_K_L", expected: "StableLM2 1.6B Zephyr Q3 K_L"},
		{input: "stablelm2:1.6b-zephyr-q4_0", expected: "StableLM2 1.6B Zephyr Q4_0"},
		{input: "stablelm2:1.6b-zephyr-q4_1", expected: "StableLM2 1.6B Zephyr Q4 1"},
		{input: "stablelm2:1.6b-zephyr-q4_K_S", expected: "StableLM2 1.6B Zephyr Q4 K_S"},
		{input: "stablelm2:1.6b-zephyr-q4_K_M", expected: "StableLM2 1.6B Zephyr Q4_K_M"},
		{input: "stablelm2:1.6b-zephyr-q5_0", expected: "StableLM2 1.6B Zephyr Q5 0"},
		{input: "stablelm2:1.6b-zephyr-q5_1", expected: "StableLM2 1.6B Zephyr Q5 1"},
		{input: "stablelm2:1.6b-zephyr-q5_K_S", expected: "StableLM2 1.6B Zephyr Q5 K_S"},
		{input: "stablelm2:1.6b-zephyr-q5_K_M", expected: "StableLM2 1.6B Zephyr Q5 K_M"},
		{input: "stablelm2:1.6b-zephyr-q6_K", expected: "StableLM2 1.6B Zephyr Q6 K"},
		{input: "stablelm2:1.6b-zephyr-q8_0", expected: "StableLM2 1.6B Zephyr Q8_0"},
		{input: "stablelm2:1.6b-zephyr-fp16", expected: "StableLM2 1.6B Zephyr FP16"},
		{input: "stablelm2:1.6b-q2_K", expected: "StableLM2 1.6B Q2_K"},
		{input: "stablelm2:1.6b-q3_K_S", expected: "StableLM2 1.6B Q3 K_S"},
		{input: "stablelm2:1.6b-q3_K_M", expected: "StableLM2 1.6B Q3 K_M"},
		{input: "stablelm2:1.6b-q3_K_L", expected: "StableLM2 1.6B Q3 K_L"},
		{input: "stablelm2:1.6b-q4_0", expected: "StableLM2 1.6B Q4_0"},
		{input: "stablelm2:1.6b-q4_1", expected: "StableLM2 1.6B Q4 1"},
		{input: "stablelm2:1.6b-q4_K_S", expected: "StableLM2 1.6B Q4 K_S"},
		{input: "stablelm2:1.6b-q4_K_M", expected: "StableLM2 1.6B Q4_K_M"},
		{input: "stablelm2:1.6b-q5_0", expected: "StableLM2 1.6B Q5 0"},
		{input: "stablelm2:1.6b-q5_1", expected: "StableLM2 1.6B Q5 1"},
		{input: "stablelm2:1.6b-q5_K_S", expected: "StableLM2 1.6B Q5 K_S"},
		{input: "stablelm2:1.6b-q5_K_M", expected: "StableLM2 1.6B Q5 K_M"},
		{input: "stablelm2:1.6b-q6_K", expected: "StableLM2 1.6B Q6 K"},
		{input: "stablelm2:1.6b-q8_0", expected: "StableLM2 1.6B Q8_0"},
		{input: "stablelm2:1.6b-fp16", expected: "StableLM2 1.6B FP16"},
		{input: "stablelm2:12b-chat", expected: "StableLM2 12B Chat"},
		{input: "stablelm2:12b-chat-q2_K", expected: "StableLM2 12B Chat Q2_K"},
		{input: "stablelm2:12b-chat-q3_K_S", expected: "StableLM2 12B Chat Q3 K_S"},
		{input: "stablelm2:12b-chat-q3_K_M", expected: "StableLM2 12B Chat Q3 K_M"},
		{input: "stablelm2:12b-chat-q3_K_L", expected: "StableLM2 12B Chat Q3 K_L"},
		{input: "stablelm2:12b-chat-q4_0", expected: "StableLM2 12B Chat Q4_0"},
		{input: "stablelm2:12b-chat-q4_1", expected: "StableLM2 12B Chat Q4 1"},
		{input: "stablelm2:12b-chat-q4_K_S", expected: "StableLM2 12B Chat Q4 K_S"},
		{input: "stablelm2:12b-chat-q4_K_M", expected: "StableLM2 12B Chat Q4_K_M"},
		{input: "stablelm2:12b-chat-q5_0", expected: "StableLM2 12B Chat Q5 0"},
		{input: "stablelm2:12b-chat-q5_1", expected: "StableLM2 12B Chat Q5 1"},
		{input: "stablelm2:12b-chat-q5_K_S", expected: "StableLM2 12B Chat Q5 K_S"},
		{input: "stablelm2:12b-chat-q5_K_M", expected: "StableLM2 12B Chat Q5 K_M"},
		{input: "stablelm2:12b-chat-q6_K", expected: "StableLM2 12B Chat Q6 K"},
		{input: "stablelm2:12b-chat-q8_0", expected: "StableLM2 12B Chat Q8_0"},
		{input: "stablelm2:12b-chat-fp16", expected: "StableLM2 12B Chat FP16"},
		{input: "stablelm2:12b-text", expected: "StableLM2 12B Text"},
		{input: "stablelm2:12b-q2_K", expected: "StableLM2 12B Q2_K"},
		{input: "stablelm2:12b-q3_K_S", expected: "StableLM2 12B Q3 K_S"},
		{input: "stablelm2:12b-q3_K_M", expected: "StableLM2 12B Q3 K_M"},
		{input: "stablelm2:12b-q3_K_L", expected: "StableLM2 12B Q3 K_L"},
		{input: "stablelm2:12b-q4_0", expected: "StableLM2 12B Q4_0"},
		{input: "stablelm2:12b-q4_1", expected: "StableLM2 12B Q4 1"},
		{input: "stablelm2:12b-q4_K_S", expected: "StableLM2 12B Q4 K_S"},
		{input: "stablelm2:12b-q4_K_M", expected: "StableLM2 12B Q4_K_M"},
		{input: "stablelm2:12b-q5_0", expected: "StableLM2 12B Q5 0"},
		{input: "stablelm2:12b-q5_1", expected: "StableLM2 12B Q5 1"},
		{input: "stablelm2:12b-q5_K_S", expected: "StableLM2 12B Q5 K_S"},
		{input: "stablelm2:12b-q5_K_M", expected: "StableLM2 12B Q5 K_M"},
		{input: "stablelm2:12b-q6_K", expected: "StableLM2 12B Q6 K"},
		{input: "stablelm2:12b-q8_0", expected: "StableLM2 12B Q8_0"},
		{input: "stablelm2:12b-fp16", expected: "StableLM2 12B FP16"},
		{input: "r1-1776:latest", expected: "R1 1776 (latest)"},
		{input: "r1-1776", expected: "R1 1776"},
		{input: "r1-1776:70b", expected: "R1 1776 70B"},
		{input: "r1-1776:671b", expected: "R1 1776 671B"},
		{input: "r1-1776:70b-distill-llama-q4_K_M", expected: "R1 1776 70B Distill Llama Q4_K_M"},
		{input: "r1-1776:70b-distill-llama-q8_0", expected: "R1 1776 70B Distill Llama Q8_0"},
		{input: "r1-1776:70b-distill-llama-fp16", expected: "R1 1776 70B Distill Llama FP16"},
		{input: "r1-1776:671b-q4_K_M", expected: "R1 1776 671B Q4_K_M"},
		{input: "r1-1776:671b-q8_0", expected: "R1 1776 671B Q8_0"},
		{input: "r1-1776:671b-fp16", expected: "R1 1776 671B FP16"},
		{input: "yi-coder:latest", expected: "Yi Coder (latest)"},
		{input: "yi-coder", expected: "Yi Coder"},
		{input: "yi-coder:1.5b", expected: "Yi Coder 1.5B"},
		{input: "yi-coder:9b", expected: "Yi Coder 9B"},
		{input: "yi-coder:1.5b-base", expected: "Yi Coder 1.5B Base"},
		{input: "yi-coder:1.5b-base-q2_K", expected: "Yi Coder 1.5B Base Q2_K"},
		{input: "yi-coder:1.5b-base-q3_K_S", expected: "Yi Coder 1.5B Base Q3 K_S"},
		{input: "yi-coder:1.5b-base-q3_K_M", expected: "Yi Coder 1.5B Base Q3 K_M"},
		{input: "yi-coder:1.5b-base-q3_K_L", expected: "Yi Coder 1.5B Base Q3 K_L"},
		{input: "yi-coder:1.5b-base-q4_0", expected: "Yi Coder 1.5B Base Q4_0"},
		{input: "yi-coder:1.5b-base-q4_1", expected: "Yi Coder 1.5B Base Q4 1"},
		{input: "yi-coder:1.5b-base-q4_K_S", expected: "Yi Coder 1.5B Base Q4 K_S"},
		{input: "yi-coder:1.5b-base-q4_K_M", expected: "Yi Coder 1.5B Base Q4_K_M"},
		{input: "yi-coder:1.5b-base-q5_0", expected: "Yi Coder 1.5B Base Q5 0"},
		{input: "yi-coder:1.5b-base-q5_1", expected: "Yi Coder 1.5B Base Q5 1"},
		{input: "yi-coder:1.5b-base-q5_K_S", expected: "Yi Coder 1.5B Base Q5 K_S"},
		{input: "yi-coder:1.5b-base-q5_K_M", expected: "Yi Coder 1.5B Base Q5 K_M"},
		{input: "yi-coder:1.5b-base-q6_K", expected: "Yi Coder 1.5B Base Q6 K"},
		{input: "yi-coder:1.5b-base-q8_0", expected: "Yi Coder 1.5B Base Q8_0"},
		{input: "yi-coder:1.5b-base-fp16", expected: "Yi Coder 1.5B Base FP16"},
		{input: "yi-coder:1.5b-chat", expected: "Yi Coder 1.5B Chat"},
		{input: "yi-coder:1.5b-chat-q2_K", expected: "Yi Coder 1.5B Chat Q2_K"},
		{input: "yi-coder:1.5b-chat-q3_K_S", expected: "Yi Coder 1.5B Chat Q3 K_S"},
		{input: "yi-coder:1.5b-chat-q3_K_M", expected: "Yi Coder 1.5B Chat Q3 K_M"},
		{input: "yi-coder:1.5b-chat-q3_K_L", expected: "Yi Coder 1.5B Chat Q3 K_L"},
		{input: "yi-coder:1.5b-chat-q4_0", expected: "Yi Coder 1.5B Chat Q4_0"},
		{input: "yi-coder:1.5b-chat-q4_1", expected: "Yi Coder 1.5B Chat Q4 1"},
		{input: "yi-coder:1.5b-chat-q4_K_S", expected: "Yi Coder 1.5B Chat Q4 K_S"},
		{input: "yi-coder:1.5b-chat-q4_K_M", expected: "Yi Coder 1.5B Chat Q4_K_M"},
		{input: "yi-coder:1.5b-chat-q5_0", expected: "Yi Coder 1.5B Chat Q5 0"},
		{input: "yi-coder:1.5b-chat-q5_1", expected: "Yi Coder 1.5B Chat Q5 1"},
		{input: "yi-coder:1.5b-chat-q5_K_S", expected: "Yi Coder 1.5B Chat Q5 K_S"},
		{input: "yi-coder:1.5b-chat-q5_K_M", expected: "Yi Coder 1.5B Chat Q5 K_M"},
		{input: "yi-coder:1.5b-chat-q6_K", expected: "Yi Coder 1.5B Chat Q6 K"},
		{input: "yi-coder:1.5b-chat-q8_0", expected: "Yi Coder 1.5B Chat Q8_0"},
		{input: "yi-coder:1.5b-chat-fp16", expected: "Yi Coder 1.5B Chat FP16"},
		{input: "yi-coder:9b-base", expected: "Yi Coder 9B Base"},
		{input: "yi-coder:9b-base-q2_K", expected: "Yi Coder 9B Base Q2_K"},
		{input: "yi-coder:9b-base-q3_K_S", expected: "Yi Coder 9B Base Q3 K_S"},
		{input: "yi-coder:9b-base-q3_K_M", expected: "Yi Coder 9B Base Q3 K_M"},
		{input: "yi-coder:9b-base-q3_K_L", expected: "Yi Coder 9B Base Q3 K_L"},
		{input: "yi-coder:9b-base-q4_0", expected: "Yi Coder 9B Base Q4_0"},
		{input: "yi-coder:9b-base-q4_1", expected: "Yi Coder 9B Base Q4 1"},
		{input: "yi-coder:9b-base-q4_K_S", expected: "Yi Coder 9B Base Q4 K_S"},
		{input: "yi-coder:9b-base-q4_K_M", expected: "Yi Coder 9B Base Q4_K_M"},
		{input: "yi-coder:9b-base-q5_0", expected: "Yi Coder 9B Base Q5 0"},
		{input: "yi-coder:9b-base-q5_1", expected: "Yi Coder 9B Base Q5 1"},
		{input: "yi-coder:9b-base-q5_K_S", expected: "Yi Coder 9B Base Q5 K_S"},
		{input: "yi-coder:9b-base-q5_K_M", expected: "Yi Coder 9B Base Q5 K_M"},
		{input: "yi-coder:9b-base-q6_K", expected: "Yi Coder 9B Base Q6 K"},
		{input: "yi-coder:9b-base-q8_0", expected: "Yi Coder 9B Base Q8_0"},
		{input: "yi-coder:9b-base-fp16", expected: "Yi Coder 9B Base FP16"},
		{input: "yi-coder:9b-chat", expected: "Yi Coder 9B Chat"},
		{input: "yi-coder:9b-chat-q2_K", expected: "Yi Coder 9B Chat Q2_K"},
		{input: "yi-coder:9b-chat-q3_K_S", expected: "Yi Coder 9B Chat Q3 K_S"},
		{input: "yi-coder:9b-chat-q3_K_M", expected: "Yi Coder 9B Chat Q3 K_M"},
		{input: "yi-coder:9b-chat-q3_K_L", expected: "Yi Coder 9B Chat Q3 K_L"},
		{input: "yi-coder:9b-chat-q4_0", expected: "Yi Coder 9B Chat Q4_0"},
		{input: "yi-coder:9b-chat-q4_1", expected: "Yi Coder 9B Chat Q4 1"},
		{input: "yi-coder:9b-chat-q4_K_S", expected: "Yi Coder 9B Chat Q4 K_S"},
		{input: "yi-coder:9b-chat-q4_K_M", expected: "Yi Coder 9B Chat Q4_K_M"},
		{input: "yi-coder:9b-chat-q5_0", expected: "Yi Coder 9B Chat Q5 0"},
		{input: "yi-coder:9b-chat-q5_1", expected: "Yi Coder 9B Chat Q5 1"},
		{input: "yi-coder:9b-chat-q5_K_S", expected: "Yi Coder 9B Chat Q5 K_S"},
		{input: "yi-coder:9b-chat-q5_K_M", expected: "Yi Coder 9B Chat Q5 K_M"},
		{input: "yi-coder:9b-chat-q6_K", expected: "Yi Coder 9B Chat Q6 K"},
		{input: "yi-coder:9b-chat-q8_0", expected: "Yi Coder 9B Chat Q8_0"},
		{input: "yi-coder:9b-chat-fp16", expected: "Yi Coder 9B Chat FP16"},
		{input: "llava-phi3:latest", expected: "Llava Phi3 (latest)"},
		{input: "llava-phi3", expected: "Llava Phi3"},
		{input: "llava-phi3:3.8b", expected: "Llava Phi3 3.8B"},
		{input: "llava-phi3:3.8b-mini-q4_0", expected: "Llava Phi3 3.8B Mini Q4_0"},
		{input: "llava-phi3:3.8b-mini-fp16", expected: "Llava Phi3 3.8B Mini FP16"},
		{input: "llama3-chatqa:latest", expected: "Llama3 Chatqa (latest)"},
		{input: "llama3-chatqa", expected: "Llama3 Chatqa"},
		{input: "llama3-chatqa:8b", expected: "Llama3 Chatqa 8B"},
		{input: "llama3-chatqa:70b", expected: "Llama3 Chatqa 70B"},
		{input: "llama3-chatqa:8b-v1.5", expected: "Llama3 Chatqa 8B v1.5"},
		{input: "llama3-chatqa:8b-v1.5-q2_K", expected: "Llama3 Chatqa 8B v1.5 Q2_K"},
		{input: "llama3-chatqa:8b-v1.5-q3_K_S", expected: "Llama3 Chatqa 8B v1.5 Q3 K_S"},
		{input: "llama3-chatqa:8b-v1.5-q3_K_M", expected: "Llama3 Chatqa 8B v1.5 Q3 K_M"},
		{input: "llama3-chatqa:8b-v1.5-q3_K_L", expected: "Llama3 Chatqa 8B v1.5 Q3 K_L"},
		{input: "llama3-chatqa:8b-v1.5-q4_0", expected: "Llama3 Chatqa 8B v1.5 Q4_0"},
		{input: "llama3-chatqa:8b-v1.5-q4_1", expected: "Llama3 Chatqa 8B v1.5 Q4 1"},
		{input: "llama3-chatqa:8b-v1.5-q4_K_S", expected: "Llama3 Chatqa 8B v1.5 Q4 K_S"},
		{input: "llama3-chatqa:8b-v1.5-q4_K_M", expected: "Llama3 Chatqa 8B v1.5 Q4_K_M"},
		{input: "llama3-chatqa:8b-v1.5-q5_0", expected: "Llama3 Chatqa 8B v1.5 Q5 0"},
		{input: "llama3-chatqa:8b-v1.5-q5_1", expected: "Llama3 Chatqa 8B v1.5 Q5 1"},
		{input: "llama3-chatqa:8b-v1.5-q5_K_S", expected: "Llama3 Chatqa 8B v1.5 Q5 K_S"},
		{input: "llama3-chatqa:8b-v1.5-q5_K_M", expected: "Llama3 Chatqa 8B v1.5 Q5 K_M"},
		{input: "llama3-chatqa:8b-v1.5-q6_K", expected: "Llama3 Chatqa 8B v1.5 Q6 K"},
		{input: "llama3-chatqa:8b-v1.5-q8_0", expected: "Llama3 Chatqa 8B v1.5 Q8_0"},
		{input: "llama3-chatqa:8b-v1.5-fp16", expected: "Llama3 Chatqa 8B v1.5 FP16"},
		{input: "llama3-chatqa:70b-v1.5", expected: "Llama3 Chatqa 70B v1.5"},
		{input: "llama3-chatqa:70b-v1.5-q2_K", expected: "Llama3 Chatqa 70B v1.5 Q2_K"},
		{input: "llama3-chatqa:70b-v1.5-q3_K_S", expected: "Llama3 Chatqa 70B v1.5 Q3 K_S"},
		{input: "llama3-chatqa:70b-v1.5-q3_K_M", expected: "Llama3 Chatqa 70B v1.5 Q3 K_M"},
		{input: "llama3-chatqa:70b-v1.5-q3_K_L", expected: "Llama3 Chatqa 70B v1.5 Q3 K_L"},
		{input: "llama3-chatqa:70b-v1.5-q4_0", expected: "Llama3 Chatqa 70B v1.5 Q4_0"},
		{input: "llama3-chatqa:70b-v1.5-q4_1", expected: "Llama3 Chatqa 70B v1.5 Q4 1"},
		{input: "llama3-chatqa:70b-v1.5-q4_K_S", expected: "Llama3 Chatqa 70B v1.5 Q4 K_S"},
		{input: "llama3-chatqa:70b-v1.5-q4_K_M", expected: "Llama3 Chatqa 70B v1.5 Q4_K_M"},
		{input: "llama3-chatqa:70b-v1.5-q5_0", expected: "Llama3 Chatqa 70B v1.5 Q5 0"},
		{input: "llama3-chatqa:70b-v1.5-q5_1", expected: "Llama3 Chatqa 70B v1.5 Q5 1"},
		{input: "llama3-chatqa:70b-v1.5-q5_K_S", expected: "Llama3 Chatqa 70B v1.5 Q5 K_S"},
		{input: "llama3-chatqa:70b-v1.5-q5_K_M", expected: "Llama3 Chatqa 70B v1.5 Q5 K_M"},
		{input: "llama3-chatqa:70b-v1.5-q6_K", expected: "Llama3 Chatqa 70B v1.5 Q6 K"},
		{input: "llama3-chatqa:70b-v1.5-q8_0", expected: "Llama3 Chatqa 70B v1.5 Q8_0"},
		{input: "llama3-chatqa:70b-v1.5-fp16", expected: "Llama3 Chatqa 70B v1.5 FP16"},
		{input: "granite3-dense:latest", expected: "Granite3 Dense (latest)"},
		{input: "granite3-dense", expected: "Granite3 Dense"},
		{input: "granite3-dense:2b", expected: "Granite3 Dense 2B"},
		{input: "granite3-dense:8b", expected: "Granite3 Dense 8B"},
		{input: "granite3-dense:2b-instruct-q2_K", expected: "Granite3 Dense 2B Instruct Q2_K"},
		{input: "granite3-dense:2b-instruct-q3_K_S", expected: "Granite3 Dense 2B Instruct Q3 K_S"},
		{input: "granite3-dense:2b-instruct-q3_K_M", expected: "Granite3 Dense 2B Instruct Q3 K_M"},
		{input: "granite3-dense:2b-instruct-q3_K_L", expected: "Granite3 Dense 2B Instruct Q3 K_L"},
		{input: "granite3-dense:2b-instruct-q4_0", expected: "Granite3 Dense 2B Instruct Q4_0"},
		{input: "granite3-dense:2b-instruct-q4_1", expected: "Granite3 Dense 2B Instruct Q4 1"},
		{input: "granite3-dense:2b-instruct-q4_K_S", expected: "Granite3 Dense 2B Instruct Q4 K_S"},
		{input: "granite3-dense:2b-instruct-q4_K_M", expected: "Granite3 Dense 2B Instruct Q4_K_M"},
		{input: "granite3-dense:2b-instruct-q5_0", expected: "Granite3 Dense 2B Instruct Q5 0"},
		{input: "granite3-dense:2b-instruct-q5_1", expected: "Granite3 Dense 2B Instruct Q5 1"},
		{input: "granite3-dense:2b-instruct-q5_K_S", expected: "Granite3 Dense 2B Instruct Q5 K_S"},
		{input: "granite3-dense:2b-instruct-q5_K_M", expected: "Granite3 Dense 2B Instruct Q5 K_M"},
		{input: "granite3-dense:2b-instruct-q6_K", expected: "Granite3 Dense 2B Instruct Q6 K"},
		{input: "granite3-dense:2b-instruct-q8_0", expected: "Granite3 Dense 2B Instruct Q8_0"},
		{input: "granite3-dense:2b-instruct-fp16", expected: "Granite3 Dense 2B Instruct FP16"},
		{input: "granite3-dense:8b-instruct-q2_K", expected: "Granite3 Dense 8B Instruct Q2_K"},
		{input: "granite3-dense:8b-instruct-q3_K_S", expected: "Granite3 Dense 8B Instruct Q3 K_S"},
		{input: "granite3-dense:8b-instruct-q3_K_M", expected: "Granite3 Dense 8B Instruct Q3 K_M"},
		{input: "granite3-dense:8b-instruct-q3_K_L", expected: "Granite3 Dense 8B Instruct Q3 K_L"},
		{input: "granite3-dense:8b-instruct-q4_0", expected: "Granite3 Dense 8B Instruct Q4_0"},
		{input: "granite3-dense:8b-instruct-q4_1", expected: "Granite3 Dense 8B Instruct Q4 1"},
		{input: "granite3-dense:8b-instruct-q4_K_S", expected: "Granite3 Dense 8B Instruct Q4 K_S"},
		{input: "granite3-dense:8b-instruct-q4_K_M", expected: "Granite3 Dense 8B Instruct Q4_K_M"},
		{input: "granite3-dense:8b-instruct-q5_0", expected: "Granite3 Dense 8B Instruct Q5 0"},
		{input: "granite3-dense:8b-instruct-q5_1", expected: "Granite3 Dense 8B Instruct Q5 1"},
		{input: "granite3-dense:8b-instruct-q5_K_S", expected: "Granite3 Dense 8B Instruct Q5 K_S"},
		{input: "granite3-dense:8b-instruct-q5_K_M", expected: "Granite3 Dense 8B Instruct Q5 K_M"},
		{input: "granite3-dense:8b-instruct-q6_K", expected: "Granite3 Dense 8B Instruct Q6 K"},
		{input: "granite3-dense:8b-instruct-q8_0", expected: "Granite3 Dense 8B Instruct Q8_0"},
		{input: "granite3-dense:8b-instruct-fp16", expected: "Granite3 Dense 8B Instruct FP16"},
		{input: "granite3.1-dense:latest", expected: "Granite3.1 Dense (latest)"},
		{input: "granite3.1-dense", expected: "Granite3.1 Dense"},
		{input: "granite3.1-dense:2b", expected: "Granite3.1 Dense 2B"},
		{input: "granite3.1-dense:8b", expected: "Granite3.1 Dense 8B"},
		{input: "granite3.1-dense:2b-instruct-q2_K", expected: "Granite3.1 Dense 2B Instruct Q2_K"},
		{input: "granite3.1-dense:2b-instruct-q3_K_S", expected: "Granite3.1 Dense 2B Instruct Q3 K_S"},
		{input: "granite3.1-dense:2b-instruct-q3_K_M", expected: "Granite3.1 Dense 2B Instruct Q3 K_M"},
		{input: "granite3.1-dense:2b-instruct-q3_K_L", expected: "Granite3.1 Dense 2B Instruct Q3 K_L"},
		{input: "granite3.1-dense:2b-instruct-q4_0", expected: "Granite3.1 Dense 2B Instruct Q4_0"},
		{input: "granite3.1-dense:2b-instruct-q4_1", expected: "Granite3.1 Dense 2B Instruct Q4 1"},
		{input: "granite3.1-dense:2b-instruct-q4_K_S", expected: "Granite3.1 Dense 2B Instruct Q4 K_S"},
		{input: "granite3.1-dense:2b-instruct-q4_K_M", expected: "Granite3.1 Dense 2B Instruct Q4_K_M"},
		{input: "granite3.1-dense:2b-instruct-q5_0", expected: "Granite3.1 Dense 2B Instruct Q5 0"},
		{input: "granite3.1-dense:2b-instruct-q5_1", expected: "Granite3.1 Dense 2B Instruct Q5 1"},
		{input: "granite3.1-dense:2b-instruct-q5_K_S", expected: "Granite3.1 Dense 2B Instruct Q5 K_S"},
		{input: "granite3.1-dense:2b-instruct-q5_K_M", expected: "Granite3.1 Dense 2B Instruct Q5 K_M"},
		{input: "granite3.1-dense:2b-instruct-q6_K", expected: "Granite3.1 Dense 2B Instruct Q6 K"},
		{input: "granite3.1-dense:2b-instruct-q8_0", expected: "Granite3.1 Dense 2B Instruct Q8_0"},
		{input: "granite3.1-dense:2b-instruct-fp16", expected: "Granite3.1 Dense 2B Instruct FP16"},
		{input: "granite3.1-dense:8b-instruct-q2_K", expected: "Granite3.1 Dense 8B Instruct Q2_K"},
		{input: "granite3.1-dense:8b-instruct-q3_K_S", expected: "Granite3.1 Dense 8B Instruct Q3 K_S"},
		{input: "granite3.1-dense:8b-instruct-q3_K_M", expected: "Granite3.1 Dense 8B Instruct Q3 K_M"},
		{input: "granite3.1-dense:8b-instruct-q3_K_L", expected: "Granite3.1 Dense 8B Instruct Q3 K_L"},
		{input: "granite3.1-dense:8b-instruct-q4_0", expected: "Granite3.1 Dense 8B Instruct Q4_0"},
		{input: "granite3.1-dense:8b-instruct-q4_1", expected: "Granite3.1 Dense 8B Instruct Q4 1"},
		{input: "granite3.1-dense:8b-instruct-q4_K_S", expected: "Granite3.1 Dense 8B Instruct Q4 K_S"},
		{input: "granite3.1-dense:8b-instruct-q4_K_M", expected: "Granite3.1 Dense 8B Instruct Q4_K_M"},
		{input: "granite3.1-dense:8b-instruct-q5_0", expected: "Granite3.1 Dense 8B Instruct Q5 0"},
		{input: "granite3.1-dense:8b-instruct-q5_1", expected: "Granite3.1 Dense 8B Instruct Q5 1"},
		{input: "granite3.1-dense:8b-instruct-q5_K_S", expected: "Granite3.1 Dense 8B Instruct Q5 K_S"},
		{input: "granite3.1-dense:8b-instruct-q5_K_M", expected: "Granite3.1 Dense 8B Instruct Q5 K_M"},
		{input: "granite3.1-dense:8b-instruct-q6_K", expected: "Granite3.1 Dense 8B Instruct Q6 K"},
		{input: "granite3.1-dense:8b-instruct-q8_0", expected: "Granite3.1 Dense 8B Instruct Q8_0"},
		{input: "granite3.1-dense:8b-instruct-fp16", expected: "Granite3.1 Dense 8B Instruct FP16"},
		{input: "exaone3.5:latest", expected: "ExaOne3.5 (latest)"},
		{input: "exaone3.5", expected: "ExaOne3.5"},
		{input: "exaone3.5:2.4b", expected: "ExaOne3.5 2.4B"},
		{input: "exaone3.5:7.8b", expected: "ExaOne3.5 7.8B"},
		{input: "exaone3.5:32b", expected: "ExaOne3.5 32B"},
		{input: "exaone3.5:2.4b-instruct-q4_K_M", expected: "ExaOne3.5 2.4B Instruct Q4_K_M"},
		{input: "exaone3.5:2.4b-instruct-q8_0", expected: "ExaOne3.5 2.4B Instruct Q8_0"},
		{input: "exaone3.5:2.4b-instruct-fp16", expected: "ExaOne3.5 2.4B Instruct FP16"},
		{input: "exaone3.5:7.8b-instruct-q4_K_M", expected: "ExaOne3.5 7.8B Instruct Q4_K_M"},
		{input: "exaone3.5:7.8b-instruct-q8_0", expected: "ExaOne3.5 7.8B Instruct Q8_0"},
		{input: "exaone3.5:7.8b-instruct-fp16", expected: "ExaOne3.5 7.8B Instruct FP16"},
		{input: "exaone3.5:32b-instruct-q4_K_M", expected: "ExaOne3.5 32B Instruct Q4_K_M"},
		{input: "exaone3.5:32b-instruct-q8_0", expected: "ExaOne3.5 32B Instruct Q8_0"},
		{input: "exaone3.5:32b-instruct-fp16", expected: "ExaOne3.5 32B Instruct FP16"},
		{input: "granite-embedding:latest", expected: "Granite Embedding (latest)"},
		{input: "granite-embedding", expected: "Granite Embedding"},
		{input: "granite-embedding:30m", expected: "Granite Embedding 30M"},
		{input: "granite-embedding:278m", expected: "Granite Embedding 278M"},
		{input: "granite-embedding:30m-en", expected: "Granite Embedding 30M En"},
		{input: "granite-embedding:30m-en-fp16", expected: "Granite Embedding 30M En FP16"},
		{input: "granite-embedding:278m-fp16", expected: "Granite Embedding 278M FP16"},
		{input: "reflection:latest", expected: "Reflection (latest)"},
		{input: "reflection", expected: "Reflection"},
		{input: "reflection:70b", expected: "Reflection 70B"},
		{input: "reflection:70b-q2_K", expected: "Reflection 70B Q2_K"},
		{input: "reflection:70b-q3_K_S", expected: "Reflection 70B Q3 K_S"},
		{input: "reflection:70b-q3_K_M", expected: "Reflection 70B Q3 K_M"},
		{input: "reflection:70b-q3_K_L", expected: "Reflection 70B Q3 K_L"},
		{input: "reflection:70b-q4_0", expected: "Reflection 70B Q4_0"},
		{input: "reflection:70b-q4_1", expected: "Reflection 70B Q4 1"},
		{input: "reflection:70b-q4_K_S", expected: "Reflection 70B Q4 K_S"},
		{input: "reflection:70b-q4_K_M", expected: "Reflection 70B Q4_K_M"},
		{input: "reflection:70b-q5_0", expected: "Reflection 70B Q5 0"},
		{input: "reflection:70b-q5_1", expected: "Reflection 70B Q5 1"},
		{input: "reflection:70b-q5_K_S", expected: "Reflection 70B Q5 K_S"},
		{input: "reflection:70b-q5_K_M", expected: "Reflection 70B Q5 K_M"},
		{input: "reflection:70b-q6_K", expected: "Reflection 70B Q6 K"},
		{input: "reflection:70b-q8_0", expected: "Reflection 70B Q8_0"},
		{input: "reflection:70b-fp16", expected: "Reflection 70B FP16"},
		{input: "wizard-math:latest", expected: "Wizard Math (latest)"},
		{input: "wizard-math", expected: "Wizard Math"},
		{input: "wizard-math:7b", expected: "Wizard Math 7B"},
		{input: "wizard-math:13b", expected: "Wizard Math 13B"},
		{input: "wizard-math:70b", expected: "Wizard Math 70B"},
		{input: "wizard-math:7b-v1.1-q2_K", expected: "Wizard Math 7B v1.1 Q2_K"},
		{input: "wizard-math:7b-v1.1-q3_K_S", expected: "Wizard Math 7B v1.1 Q3 K_S"},
		{input: "wizard-math:7b-v1.1-q3_K_M", expected: "Wizard Math 7B v1.1 Q3 K_M"},
		{input: "wizard-math:7b-v1.1-q3_K_L", expected: "Wizard Math 7B v1.1 Q3 K_L"},
		{input: "wizard-math:7b-v1.1-q4_0", expected: "Wizard Math 7B v1.1 Q4_0"},
		{input: "wizard-math:7b-v1.1-q4_1", expected: "Wizard Math 7B v1.1 Q4 1"},
		{input: "wizard-math:7b-v1.1-q4_K_S", expected: "Wizard Math 7B v1.1 Q4 K_S"},
		{input: "wizard-math:7b-v1.1-q4_K_M", expected: "Wizard Math 7B v1.1 Q4_K_M"},
		{input: "wizard-math:7b-v1.1-q5_0", expected: "Wizard Math 7B v1.1 Q5 0"},
		{input: "wizard-math:7b-v1.1-q5_1", expected: "Wizard Math 7B v1.1 Q5 1"},
		{input: "wizard-math:7b-v1.1-q5_K_S", expected: "Wizard Math 7B v1.1 Q5 K_S"},
		{input: "wizard-math:7b-v1.1-q5_K_M", expected: "Wizard Math 7B v1.1 Q5 K_M"},
		{input: "wizard-math:7b-v1.1-q6_K", expected: "Wizard Math 7B v1.1 Q6 K"},
		{input: "wizard-math:7b-v1.1-q8_0", expected: "Wizard Math 7B v1.1 Q8_0"},
		{input: "wizard-math:7b-v1.1-fp16", expected: "Wizard Math 7B v1.1 FP16"},
		{input: "wizard-math:7b-q2_K", expected: "Wizard Math 7B Q2_K"},
		{input: "wizard-math:7b-q3_K_S", expected: "Wizard Math 7B Q3 K_S"},
		{input: "wizard-math:7b-q3_K_M", expected: "Wizard Math 7B Q3 K_M"},
		{input: "wizard-math:7b-q3_K_L", expected: "Wizard Math 7B Q3 K_L"},
		{input: "wizard-math:7b-q4_0", expected: "Wizard Math 7B Q4_0"},
		{input: "wizard-math:7b-q4_1", expected: "Wizard Math 7B Q4 1"},
		{input: "wizard-math:7b-q4_K_S", expected: "Wizard Math 7B Q4 K_S"},
		{input: "wizard-math:7b-q4_K_M", expected: "Wizard Math 7B Q4_K_M"},
		{input: "wizard-math:7b-q5_0", expected: "Wizard Math 7B Q5 0"},
		{input: "wizard-math:7b-q5_1", expected: "Wizard Math 7B Q5 1"},
		{input: "wizard-math:7b-q5_K_S", expected: "Wizard Math 7B Q5 K_S"},
		{input: "wizard-math:7b-q5_K_M", expected: "Wizard Math 7B Q5 K_M"},
		{input: "wizard-math:7b-q6_K", expected: "Wizard Math 7B Q6 K"},
		{input: "wizard-math:7b-q8_0", expected: "Wizard Math 7B Q8_0"},
		{input: "wizard-math:7b-fp16", expected: "Wizard Math 7B FP16"},
		{input: "wizard-math:13b-q2_K", expected: "Wizard Math 13B Q2_K"},
		{input: "wizard-math:13b-q3_K_S", expected: "Wizard Math 13B Q3 K_S"},
		{input: "wizard-math:13b-q3_K_M", expected: "Wizard Math 13B Q3 K_M"},
		{input: "wizard-math:13b-q3_K_L", expected: "Wizard Math 13B Q3 K_L"},
		{input: "wizard-math:13b-q4_0", expected: "Wizard Math 13B Q4_0"},
		{input: "wizard-math:13b-q4_1", expected: "Wizard Math 13B Q4 1"},
		{input: "wizard-math:13b-q4_K_S", expected: "Wizard Math 13B Q4 K_S"},
		{input: "wizard-math:13b-q4_K_M", expected: "Wizard Math 13B Q4_K_M"},
		{input: "wizard-math:13b-q5_0", expected: "Wizard Math 13B Q5 0"},
		{input: "wizard-math:13b-q5_1", expected: "Wizard Math 13B Q5 1"},
		{input: "wizard-math:13b-q5_K_S", expected: "Wizard Math 13B Q5 K_S"},
		{input: "wizard-math:13b-q5_K_M", expected: "Wizard Math 13B Q5 K_M"},
		{input: "wizard-math:13b-q6_K", expected: "Wizard Math 13B Q6 K"},
		{input: "wizard-math:13b-q8_0", expected: "Wizard Math 13B Q8_0"},
		{input: "wizard-math:13b-fp16", expected: "Wizard Math 13B FP16"},
		{input: "wizard-math:70b-q2_K", expected: "Wizard Math 70B Q2_K"},
		{input: "wizard-math:70b-q3_K_S", expected: "Wizard Math 70B Q3 K_S"},
		{input: "wizard-math:70b-q3_K_M", expected: "Wizard Math 70B Q3 K_M"},
		{input: "wizard-math:70b-q3_K_L", expected: "Wizard Math 70B Q3 K_L"},
		{input: "wizard-math:70b-q4_0", expected: "Wizard Math 70B Q4_0"},
		{input: "wizard-math:70b-q4_1", expected: "Wizard Math 70B Q4 1"},
		{input: "wizard-math:70b-q4_K_S", expected: "Wizard Math 70B Q4 K_S"},
		{input: "wizard-math:70b-q4_K_M", expected: "Wizard Math 70B Q4_K_M"},
		{input: "wizard-math:70b-q5_0", expected: "Wizard Math 70B Q5 0"},
		{input: "wizard-math:70b-q5_1", expected: "Wizard Math 70B Q5 1"},
		{input: "wizard-math:70b-q5_K_S", expected: "Wizard Math 70B Q5 K_S"},
		{input: "wizard-math:70b-q5_K_M", expected: "Wizard Math 70B Q5 K_M"},
		{input: "wizard-math:70b-q6_K", expected: "Wizard Math 70B Q6 K"},
		{input: "wizard-math:70b-q8_0", expected: "Wizard Math 70B Q8_0"},
		{input: "wizard-math:70b-fp16", expected: "Wizard Math 70B FP16"},
		{input: "llama3-gradient:latest", expected: "Llama3 Gradient (latest)"},
		{input: "llama3-gradient", expected: "Llama3 Gradient"},
		{input: "llama3-gradient:instruct", expected: "Llama3 Gradient Instruct"},
		{input: "llama3-gradient:1048k", expected: "Llama3 Gradient 1048K"},
		{input: "llama3-gradient:8b", expected: "Llama3 Gradient 8B"},
		{input: "llama3-gradient:70b", expected: "Llama3 Gradient 70B"},
		{input: "llama3-gradient:8b-instruct-1048k-q2_K", expected: "Llama3 Gradient 8B Instruct 1048K Q2_K"},
		{input: "llama3-gradient:8b-instruct-1048k-q3_K_S", expected: "Llama3 Gradient 8B Instruct 1048K Q3 K_S"},
		{input: "llama3-gradient:8b-instruct-1048k-q3_K_M", expected: "Llama3 Gradient 8B Instruct 1048K Q3 K_M"},
		{input: "llama3-gradient:8b-instruct-1048k-q3_K_L", expected: "Llama3 Gradient 8B Instruct 1048K Q3 K_L"},
		{input: "llama3-gradient:8b-instruct-1048k-q4_0", expected: "Llama3 Gradient 8B Instruct 1048K Q4_0"},
		{input: "llama3-gradient:8b-instruct-1048k-q4_1", expected: "Llama3 Gradient 8B Instruct 1048K Q4 1"},
		{input: "llama3-gradient:8b-instruct-1048k-q4_K_S", expected: "Llama3 Gradient 8B Instruct 1048K Q4 K_S"},
		{input: "llama3-gradient:8b-instruct-1048k-q4_K_M", expected: "Llama3 Gradient 8B Instruct 1048K Q4_K_M"},
		{input: "llama3-gradient:8b-instruct-1048k-q5_0", expected: "Llama3 Gradient 8B Instruct 1048K Q5 0"},
		{input: "llama3-gradient:8b-instruct-1048k-q5_1", expected: "Llama3 Gradient 8B Instruct 1048K Q5 1"},
		{input: "llama3-gradient:8b-instruct-1048k-q5_K_S", expected: "Llama3 Gradient 8B Instruct 1048K Q5 K_S"},
		{input: "llama3-gradient:8b-instruct-1048k-q5_K_M", expected: "Llama3 Gradient 8B Instruct 1048K Q5 K_M"},
		{input: "llama3-gradient:8b-instruct-1048k-q6_K", expected: "Llama3 Gradient 8B Instruct 1048K Q6 K"},
		{input: "llama3-gradient:8b-instruct-1048k-q8_0", expected: "Llama3 Gradient 8B Instruct 1048K Q8_0"},
		{input: "llama3-gradient:8b-instruct-1048k-fp16", expected: "Llama3 Gradient 8B Instruct 1048K FP16"},
		{input: "llama3-gradient:70b-instruct-1048k-q2_K", expected: "Llama3 Gradient 70B Instruct 1048K Q2_K"},
		{input: "llama3-gradient:70b-instruct-1048k-q3_K_S", expected: "Llama3 Gradient 70B Instruct 1048K Q3 K_S"},
		{input: "llama3-gradient:70b-instruct-1048k-q3_K_M", expected: "Llama3 Gradient 70B Instruct 1048K Q3 K_M"},
		{input: "llama3-gradient:70b-instruct-1048k-q3_K_L", expected: "Llama3 Gradient 70B Instruct 1048K Q3 K_L"},
		{input: "llama3-gradient:70b-instruct-1048k-q4_0", expected: "Llama3 Gradient 70B Instruct 1048K Q4_0"},
		{input: "llama3-gradient:70b-instruct-1048k-q4_1", expected: "Llama3 Gradient 70B Instruct 1048K Q4 1"},
		{input: "llama3-gradient:70b-instruct-1048k-q4_K_S", expected: "Llama3 Gradient 70B Instruct 1048K Q4 K_S"},
		{input: "llama3-gradient:70b-instruct-1048k-q4_K_M", expected: "Llama3 Gradient 70B Instruct 1048K Q4_K_M"},
		{input: "llama3-gradient:70b-instruct-1048k-q5_0", expected: "Llama3 Gradient 70B Instruct 1048K Q5 0"},
		{input: "llama3-gradient:70b-instruct-1048k-q5_1", expected: "Llama3 Gradient 70B Instruct 1048K Q5 1"},
		{input: "llama3-gradient:70b-instruct-1048k-q5_K_S", expected: "Llama3 Gradient 70B Instruct 1048K Q5 K_S"},
		{input: "llama3-gradient:70b-instruct-1048k-q5_K_M", expected: "Llama3 Gradient 70B Instruct 1048K Q5 K_M"},
		{input: "llama3-gradient:70b-instruct-1048k-q6_K", expected: "Llama3 Gradient 70B Instruct 1048K Q6 K"},
		{input: "llama3-gradient:70b-instruct-1048k-q8_0", expected: "Llama3 Gradient 70B Instruct 1048K Q8_0"},
		{input: "llama3-gradient:70b-instruct-1048k-fp16", expected: "Llama3 Gradient 70B Instruct 1048K FP16"},
		{input: "dolphincoder:latest", expected: "Dolphincoder (latest)"},
		{input: "dolphincoder", expected: "Dolphincoder"},
		{input: "dolphincoder:7b", expected: "Dolphincoder 7B"},
		{input: "dolphincoder:15b", expected: "Dolphincoder 15B"},
		{input: "dolphincoder:7b-starcoder2", expected: "Dolphincoder 7B StarCoder2"},
		{input: "dolphincoder:7b-starcoder2-q2_K", expected: "Dolphincoder 7B StarCoder2 Q2_K"},
		{input: "dolphincoder:7b-starcoder2-q3_K_S", expected: "Dolphincoder 7B StarCoder2 Q3 K_S"},
		{input: "dolphincoder:7b-starcoder2-q3_K_M", expected: "Dolphincoder 7B StarCoder2 Q3 K_M"},
		{input: "dolphincoder:7b-starcoder2-q3_K_L", expected: "Dolphincoder 7B StarCoder2 Q3 K_L"},
		{input: "dolphincoder:7b-starcoder2-q4_0", expected: "Dolphincoder 7B StarCoder2 Q4_0"},
		{input: "dolphincoder:7b-starcoder2-q4_1", expected: "Dolphincoder 7B StarCoder2 Q4 1"},
		{input: "dolphincoder:7b-starcoder2-q4_K_S", expected: "Dolphincoder 7B StarCoder2 Q4 K_S"},
		{input: "dolphincoder:7b-starcoder2-q4_K_M", expected: "Dolphincoder 7B StarCoder2 Q4_K_M"},
		{input: "dolphincoder:7b-starcoder2-q5_0", expected: "Dolphincoder 7B StarCoder2 Q5 0"},
		{input: "dolphincoder:7b-starcoder2-q5_1", expected: "Dolphincoder 7B StarCoder2 Q5 1"},
		{input: "dolphincoder:7b-starcoder2-q5_K_S", expected: "Dolphincoder 7B StarCoder2 Q5 K_S"},
		{input: "dolphincoder:7b-starcoder2-q5_K_M", expected: "Dolphincoder 7B StarCoder2 Q5 K_M"},
		{input: "dolphincoder:7b-starcoder2-q6_K", expected: "Dolphincoder 7B StarCoder2 Q6 K"},
		{input: "dolphincoder:7b-starcoder2-q8_0", expected: "Dolphincoder 7B StarCoder2 Q8_0"},
		{input: "dolphincoder:7b-starcoder2-fp16", expected: "Dolphincoder 7B StarCoder2 FP16"},
		{input: "dolphincoder:15b-starcoder2", expected: "Dolphincoder 15B StarCoder2"},
		{input: "dolphincoder:15b-starcoder2-q2_K", expected: "Dolphincoder 15B StarCoder2 Q2_K"},
		{input: "dolphincoder:15b-starcoder2-q3_K_S", expected: "Dolphincoder 15B StarCoder2 Q3 K_S"},
		{input: "dolphincoder:15b-starcoder2-q3_K_M", expected: "Dolphincoder 15B StarCoder2 Q3 K_M"},
		{input: "dolphincoder:15b-starcoder2-q3_K_L", expected: "Dolphincoder 15B StarCoder2 Q3 K_L"},
		{input: "dolphincoder:15b-starcoder2-q4_0", expected: "Dolphincoder 15B StarCoder2 Q4_0"},
		{input: "dolphincoder:15b-starcoder2-q4_1", expected: "Dolphincoder 15B StarCoder2 Q4 1"},
		{input: "dolphincoder:15b-starcoder2-q4_K_S", expected: "Dolphincoder 15B StarCoder2 Q4 K_S"},
		{input: "dolphincoder:15b-starcoder2-q4_K_M", expected: "Dolphincoder 15B StarCoder2 Q4_K_M"},
		{input: "dolphincoder:15b-starcoder2-q5_0", expected: "Dolphincoder 15B StarCoder2 Q5 0"},
		{input: "dolphincoder:15b-starcoder2-q5_1", expected: "Dolphincoder 15B StarCoder2 Q5 1"},
		{input: "dolphincoder:15b-starcoder2-q5_K_S", expected: "Dolphincoder 15B StarCoder2 Q5 K_S"},
		{input: "dolphincoder:15b-starcoder2-q5_K_M", expected: "Dolphincoder 15B StarCoder2 Q5 K_M"},
		{input: "dolphincoder:15b-starcoder2-q6_K", expected: "Dolphincoder 15B StarCoder2 Q6 K"},
		{input: "dolphincoder:15b-starcoder2-q8_0", expected: "Dolphincoder 15B StarCoder2 Q8_0"},
		{input: "dolphincoder:15b-starcoder2-fp16", expected: "Dolphincoder 15B StarCoder2 FP16"},
		{input: "samantha-mistral:latest", expected: "Samantha Mistral (latest)"},
		{input: "samantha-mistral", expected: "Samantha Mistral"},
		{input: "samantha-mistral:7b", expected: "Samantha Mistral 7B"},
		{input: "samantha-mistral:7b-instruct-q2_K", expected: "Samantha Mistral 7B Instruct Q2_K"},
		{input: "samantha-mistral:7b-instruct-q3_K_S", expected: "Samantha Mistral 7B Instruct Q3 K_S"},
		{input: "samantha-mistral:7b-instruct-q3_K_M", expected: "Samantha Mistral 7B Instruct Q3 K_M"},
		{input: "samantha-mistral:7b-instruct-q3_K_L", expected: "Samantha Mistral 7B Instruct Q3 K_L"},
		{input: "samantha-mistral:7b-instruct-q4_0", expected: "Samantha Mistral 7B Instruct Q4_0"},
		{input: "samantha-mistral:7b-instruct-q4_1", expected: "Samantha Mistral 7B Instruct Q4 1"},
		{input: "samantha-mistral:7b-instruct-q4_K_S", expected: "Samantha Mistral 7B Instruct Q4 K_S"},
		{input: "samantha-mistral:7b-instruct-q4_K_M", expected: "Samantha Mistral 7B Instruct Q4_K_M"},
		{input: "samantha-mistral:7b-instruct-q5_0", expected: "Samantha Mistral 7B Instruct Q5 0"},
		{input: "samantha-mistral:7b-instruct-q5_1", expected: "Samantha Mistral 7B Instruct Q5 1"},
		{input: "samantha-mistral:7b-instruct-q5_K_S", expected: "Samantha Mistral 7B Instruct Q5 K_S"},
		{input: "samantha-mistral:7b-instruct-q5_K_M", expected: "Samantha Mistral 7B Instruct Q5 K_M"},
		{input: "samantha-mistral:7b-instruct-q6_K", expected: "Samantha Mistral 7B Instruct Q6 K"},
		{input: "samantha-mistral:7b-instruct-q8_0", expected: "Samantha Mistral 7B Instruct Q8_0"},
		{input: "samantha-mistral:7b-instruct-fp16", expected: "Samantha Mistral 7B Instruct FP16"},
		{input: "samantha-mistral:7b-text", expected: "Samantha Mistral 7B Text"},
		{input: "samantha-mistral:7b-text-q2_K", expected: "Samantha Mistral 7B Text Q2_K"},
		{input: "samantha-mistral:7b-text-q3_K_S", expected: "Samantha Mistral 7B Text Q3 K_S"},
		{input: "samantha-mistral:7b-text-q3_K_M", expected: "Samantha Mistral 7B Text Q3 K_M"},
		{input: "samantha-mistral:7b-text-q3_K_L", expected: "Samantha Mistral 7B Text Q3 K_L"},
		{input: "samantha-mistral:7b-text-q4_0", expected: "Samantha Mistral 7B Text Q4_0"},
		{input: "samantha-mistral:7b-text-q4_1", expected: "Samantha Mistral 7B Text Q4 1"},
		{input: "samantha-mistral:7b-text-q4_K_S", expected: "Samantha Mistral 7B Text Q4 K_S"},
		{input: "samantha-mistral:7b-text-q4_K_M", expected: "Samantha Mistral 7B Text Q4_K_M"},
		{input: "samantha-mistral:7b-text-q5_0", expected: "Samantha Mistral 7B Text Q5 0"},
		{input: "samantha-mistral:7b-text-q5_1", expected: "Samantha Mistral 7B Text Q5 1"},
		{input: "samantha-mistral:7b-text-q5_K_S", expected: "Samantha Mistral 7B Text Q5 K_S"},
		{input: "samantha-mistral:7b-text-q5_K_M", expected: "Samantha Mistral 7B Text Q5 K_M"},
		{input: "samantha-mistral:7b-text-q6_K", expected: "Samantha Mistral 7B Text Q6 K"},
		{input: "samantha-mistral:7b-text-q8_0", expected: "Samantha Mistral 7B Text Q8_0"},
		{input: "samantha-mistral:7b-text-fp16", expected: "Samantha Mistral 7B Text FP16"},
		{input: "samantha-mistral:7b-v1.2-text", expected: "Samantha Mistral 7B v1.2 Text"},
		{input: "samantha-mistral:7b-v1.2-text-q2_K", expected: "Samantha Mistral 7B v1.2 Text Q2_K"},
		{input: "samantha-mistral:7b-v1.2-text-q3_K_S", expected: "Samantha Mistral 7B v1.2 Text Q3 K_S"},
		{input: "samantha-mistral:7b-v1.2-text-q3_K_M", expected: "Samantha Mistral 7B v1.2 Text Q3 K_M"},
		{input: "samantha-mistral:7b-v1.2-text-q3_K_L", expected: "Samantha Mistral 7B v1.2 Text Q3 K_L"},
		{input: "samantha-mistral:7b-v1.2-text-q4_0", expected: "Samantha Mistral 7B v1.2 Text Q4_0"},
		{input: "samantha-mistral:7b-v1.2-text-q4_1", expected: "Samantha Mistral 7B v1.2 Text Q4 1"},
		{input: "samantha-mistral:7b-v1.2-text-q4_K_S", expected: "Samantha Mistral 7B v1.2 Text Q4 K_S"},
		{input: "samantha-mistral:7b-v1.2-text-q4_K_M", expected: "Samantha Mistral 7B v1.2 Text Q4_K_M"},
		{input: "samantha-mistral:7b-v1.2-text-q5_0", expected: "Samantha Mistral 7B v1.2 Text Q5 0"},
		{input: "samantha-mistral:7b-v1.2-text-q5_1", expected: "Samantha Mistral 7B v1.2 Text Q5 1"},
		{input: "samantha-mistral:7b-v1.2-text-q5_K_S", expected: "Samantha Mistral 7B v1.2 Text Q5 K_S"},
		{input: "samantha-mistral:7b-v1.2-text-q5_K_M", expected: "Samantha Mistral 7B v1.2 Text Q5 K_M"},
		{input: "samantha-mistral:7b-v1.2-text-q6_K", expected: "Samantha Mistral 7B v1.2 Text Q6 K"},
		{input: "samantha-mistral:7b-v1.2-text-q8_0", expected: "Samantha Mistral 7B v1.2 Text Q8_0"},
		{input: "samantha-mistral:7b-v1.2-text-fp16", expected: "Samantha Mistral 7B v1.2 Text FP16"},
		{input: "nemotron-mini:latest", expected: "Nemotron Mini (latest)"},
		{input: "nemotron-mini", expected: "Nemotron Mini"},
		{input: "nemotron-mini:4b", expected: "Nemotron Mini 4B"},
		{input: "nemotron-mini:4b-instruct-q2_K", expected: "Nemotron Mini 4B Instruct Q2_K"},
		{input: "nemotron-mini:4b-instruct-q3_K_S", expected: "Nemotron Mini 4B Instruct Q3 K_S"},
		{input: "nemotron-mini:4b-instruct-q3_K_M", expected: "Nemotron Mini 4B Instruct Q3 K_M"},
		{input: "nemotron-mini:4b-instruct-q3_K_L", expected: "Nemotron Mini 4B Instruct Q3 K_L"},
		{input: "nemotron-mini:4b-instruct-q4_0", expected: "Nemotron Mini 4B Instruct Q4_0"},
		{input: "nemotron-mini:4b-instruct-q4_1", expected: "Nemotron Mini 4B Instruct Q4 1"},
		{input: "nemotron-mini:4b-instruct-q4_K_S", expected: "Nemotron Mini 4B Instruct Q4 K_S"},
		{input: "nemotron-mini:4b-instruct-q4_K_M", expected: "Nemotron Mini 4B Instruct Q4_K_M"},
		{input: "nemotron-mini:4b-instruct-q5_0", expected: "Nemotron Mini 4B Instruct Q5 0"},
		{input: "nemotron-mini:4b-instruct-q5_1", expected: "Nemotron Mini 4B Instruct Q5 1"},
		{input: "nemotron-mini:4b-instruct-q5_K_S", expected: "Nemotron Mini 4B Instruct Q5 K_S"},
		{input: "nemotron-mini:4b-instruct-q5_K_M", expected: "Nemotron Mini 4B Instruct Q5 K_M"},
		{input: "nemotron-mini:4b-instruct-q6_K", expected: "Nemotron Mini 4B Instruct Q6 K"},
		{input: "nemotron-mini:4b-instruct-q8_0", expected: "Nemotron Mini 4B Instruct Q8_0"},
		{input: "nemotron-mini:4b-instruct-fp16", expected: "Nemotron Mini 4B Instruct FP16"},
		{input: "dbrx:latest", expected: "Dbrx (latest)"},
		{input: "dbrx", expected: "Dbrx"},
		{input: "dbrx:instruct", expected: "Dbrx Instruct"},
		{input: "dbrx:132b", expected: "Dbrx 132B"},
		{input: "dbrx:132b-instruct-q2_K", expected: "Dbrx 132B Instruct Q2_K"},
		{input: "dbrx:132b-instruct-q4_0", expected: "Dbrx 132B Instruct Q4_0"},
		{input: "dbrx:132b-instruct-q8_0", expected: "Dbrx 132B Instruct Q8_0"},
		{input: "dbrx:132b-instruct-fp16", expected: "Dbrx 132B Instruct FP16"},
		{input: "internlm2:latest", expected: "InternLM2 (latest)"},
		{input: "internlm2", expected: "InternLM2"},
		{input: "internlm2:1m", expected: "InternLM2 1M"},
		{input: "internlm2:1.8b", expected: "InternLM2 1.8B"},
		{input: "internlm2:7b", expected: "InternLM2 7B"},
		{input: "internlm2:20b", expected: "InternLM2 20B"},
		{input: "internlm2:1.8b-chat-v2.5-q2_K", expected: "InternLM2 1.8B Chat v2.5 Q2_K"},
		{input: "internlm2:1.8b-chat-v2.5-q3_K_S", expected: "InternLM2 1.8B Chat v2.5 Q3 K_S"},
		{input: "internlm2:1.8b-chat-v2.5-q3_K_M", expected: "InternLM2 1.8B Chat v2.5 Q3 K_M"},
		{input: "internlm2:1.8b-chat-v2.5-q3_K_L", expected: "InternLM2 1.8B Chat v2.5 Q3 K_L"},
		{input: "internlm2:1.8b-chat-v2.5-q4_0", expected: "InternLM2 1.8B Chat v2.5 Q4_0"},
		{input: "internlm2:1.8b-chat-v2.5-q4_1", expected: "InternLM2 1.8B Chat v2.5 Q4 1"},
		{input: "internlm2:1.8b-chat-v2.5-q4_K_S", expected: "InternLM2 1.8B Chat v2.5 Q4 K_S"},
		{input: "internlm2:1.8b-chat-v2.5-q4_K_M", expected: "InternLM2 1.8B Chat v2.5 Q4_K_M"},
		{input: "internlm2:1.8b-chat-v2.5-q5_0", expected: "InternLM2 1.8B Chat v2.5 Q5 0"},
		{input: "internlm2:1.8b-chat-v2.5-q5_1", expected: "InternLM2 1.8B Chat v2.5 Q5 1"},
		{input: "internlm2:1.8b-chat-v2.5-q5_K_S", expected: "InternLM2 1.8B Chat v2.5 Q5 K_S"},
		{input: "internlm2:1.8b-chat-v2.5-q5_K_M", expected: "InternLM2 1.8B Chat v2.5 Q5 K_M"},
		{input: "internlm2:1.8b-chat-v2.5-q6_K", expected: "InternLM2 1.8B Chat v2.5 Q6 K"},
		{input: "internlm2:1.8b-chat-v2.5-q8_0", expected: "InternLM2 1.8B Chat v2.5 Q8_0"},
		{input: "internlm2:1.8b-chat-v2.5-fp16", expected: "InternLM2 1.8B Chat v2.5 FP16"},
		{input: "internlm2:7b-chat-1m-v2.5-q2_K", expected: "InternLM2 7B Chat 1M v2.5 Q2_K"},
		{input: "internlm2:7b-chat-1m-v2.5-q3_K_S", expected: "InternLM2 7B Chat 1M v2.5 Q3 K_S"},
		{input: "internlm2:7b-chat-1m-v2.5-q3_K_M", expected: "InternLM2 7B Chat 1M v2.5 Q3 K_M"},
		{input: "internlm2:7b-chat-1m-v2.5-q3_K_L", expected: "InternLM2 7B Chat 1M v2.5 Q3 K_L"},
		{input: "internlm2:7b-chat-1m-v2.5-q4_0", expected: "InternLM2 7B Chat 1M v2.5 Q4_0"},
		{input: "internlm2:7b-chat-1m-v2.5-q4_1", expected: "InternLM2 7B Chat 1M v2.5 Q4 1"},
		{input: "internlm2:7b-chat-1m-v2.5-q4_K_S", expected: "InternLM2 7B Chat 1M v2.5 Q4 K_S"},
		{input: "internlm2:7b-chat-1m-v2.5-q4_K_M", expected: "InternLM2 7B Chat 1M v2.5 Q4_K_M"},
		{input: "internlm2:7b-chat-1m-v2.5-q5_0", expected: "InternLM2 7B Chat 1M v2.5 Q5 0"},
		{input: "internlm2:7b-chat-1m-v2.5-q5_1", expected: "InternLM2 7B Chat 1M v2.5 Q5 1"},
		{input: "internlm2:7b-chat-1m-v2.5-q5_K_S", expected: "InternLM2 7B Chat 1M v2.5 Q5 K_S"},
		{input: "internlm2:7b-chat-1m-v2.5-q5_K_M", expected: "InternLM2 7B Chat 1M v2.5 Q5 K_M"},
		{input: "internlm2:7b-chat-1m-v2.5-q6_K", expected: "InternLM2 7B Chat 1M v2.5 Q6 K"},
		{input: "internlm2:7b-chat-1m-v2.5-q8_0", expected: "InternLM2 7B Chat 1M v2.5 Q8_0"},
		{input: "internlm2:7b-chat-1m-v2.5-fp16", expected: "InternLM2 7B Chat 1M v2.5 FP16"},
		{input: "internlm2:7b-chat-v2.5-q2_K", expected: "InternLM2 7B Chat v2.5 Q2_K"},
		{input: "internlm2:7b-chat-v2.5-q3_K_S", expected: "InternLM2 7B Chat v2.5 Q3 K_S"},
		{input: "internlm2:7b-chat-v2.5-q3_K_M", expected: "InternLM2 7B Chat v2.5 Q3 K_M"},
		{input: "internlm2:7b-chat-v2.5-q3_K_L", expected: "InternLM2 7B Chat v2.5 Q3 K_L"},
		{input: "internlm2:7b-chat-v2.5-q4_0", expected: "InternLM2 7B Chat v2.5 Q4_0"},
		{input: "internlm2:7b-chat-v2.5-q4_1", expected: "InternLM2 7B Chat v2.5 Q4 1"},
		{input: "internlm2:7b-chat-v2.5-q4_K_S", expected: "InternLM2 7B Chat v2.5 Q4 K_S"},
		{input: "internlm2:7b-chat-v2.5-q4_K_M", expected: "InternLM2 7B Chat v2.5 Q4_K_M"},
		{input: "internlm2:7b-chat-v2.5-q5_0", expected: "InternLM2 7B Chat v2.5 Q5 0"},
		{input: "internlm2:7b-chat-v2.5-q5_1", expected: "InternLM2 7B Chat v2.5 Q5 1"},
		{input: "internlm2:7b-chat-v2.5-q5_K_S", expected: "InternLM2 7B Chat v2.5 Q5 K_S"},
		{input: "internlm2:7b-chat-v2.5-q5_K_M", expected: "InternLM2 7B Chat v2.5 Q5 K_M"},
		{input: "internlm2:7b-chat-v2.5-q6_K", expected: "InternLM2 7B Chat v2.5 Q6 K"},
		{input: "internlm2:7b-chat-v2.5-q8_0", expected: "InternLM2 7B Chat v2.5 Q8_0"},
		{input: "internlm2:7b-chat-v2.5-fp16", expected: "InternLM2 7B Chat v2.5 FP16"},
		{input: "internlm2:20b-chat-v2.5-q2_K", expected: "InternLM2 20B Chat v2.5 Q2_K"},
		{input: "internlm2:20b-chat-v2.5-q3_K_S", expected: "InternLM2 20B Chat v2.5 Q3 K_S"},
		{input: "internlm2:20b-chat-v2.5-q3_K_M", expected: "InternLM2 20B Chat v2.5 Q3 K_M"},
		{input: "internlm2:20b-chat-v2.5-q3_K_L", expected: "InternLM2 20B Chat v2.5 Q3 K_L"},
		{input: "internlm2:20b-chat-v2.5-q4_0", expected: "InternLM2 20B Chat v2.5 Q4_0"},
		{input: "internlm2:20b-chat-v2.5-q4_1", expected: "InternLM2 20B Chat v2.5 Q4 1"},
		{input: "internlm2:20b-chat-v2.5-q4_K_S", expected: "InternLM2 20B Chat v2.5 Q4 K_S"},
		{input: "internlm2:20b-chat-v2.5-q4_K_M", expected: "InternLM2 20B Chat v2.5 Q4_K_M"},
		{input: "internlm2:20b-chat-v2.5-q5_0", expected: "InternLM2 20B Chat v2.5 Q5 0"},
		{input: "internlm2:20b-chat-v2.5-q5_1", expected: "InternLM2 20B Chat v2.5 Q5 1"},
		{input: "internlm2:20b-chat-v2.5-q5_K_S", expected: "InternLM2 20B Chat v2.5 Q5 K_S"},
		{input: "internlm2:20b-chat-v2.5-q5_K_M", expected: "InternLM2 20B Chat v2.5 Q5 K_M"},
		{input: "internlm2:20b-chat-v2.5-q6_K", expected: "InternLM2 20B Chat v2.5 Q6 K"},
		{input: "internlm2:20b-chat-v2.5-q8_0", expected: "InternLM2 20B Chat v2.5 Q8_0"},
		{input: "internlm2:20b-chat-v2.5-fp16", expected: "InternLM2 20B Chat v2.5 FP16"},
		{input: "tulu3:latest", expected: "Tulu3 (latest)"},
		{input: "tulu3", expected: "Tulu3"},
		{input: "tulu3:8b", expected: "Tulu3 8B"},
		{input: "tulu3:70b", expected: "Tulu3 70B"},
		{input: "tulu3:8b-q4_K_M", expected: "Tulu3 8B Q4_K_M"},
		{input: "tulu3:8b-q8_0", expected: "Tulu3 8B Q8_0"},
		{input: "tulu3:8b-fp16", expected: "Tulu3 8B FP16"},
		{input: "tulu3:70b-q4_K_M", expected: "Tulu3 70B Q4_K_M"},
		{input: "tulu3:70b-q8_0", expected: "Tulu3 70B Q8_0"},
		{input: "tulu3:70b-fp16", expected: "Tulu3 70B FP16"},
		{input: "starling-lm:latest", expected: "Starling Lm (latest)"},
		{input: "starling-lm", expected: "Starling Lm"},
		{input: "starling-lm:alpha", expected: "Starling Lm Alpha"},
		{input: "starling-lm:beta", expected: "Starling Lm Beta"},
		{input: "starling-lm:7b", expected: "Starling Lm 7B"},
		{input: "starling-lm:7b-alpha", expected: "Starling Lm 7B Alpha"},
		{input: "starling-lm:7b-alpha-q2_K", expected: "Starling Lm 7B Alpha Q2_K"},
		{input: "starling-lm:7b-alpha-q3_K_S", expected: "Starling Lm 7B Alpha Q3 K_S"},
		{input: "starling-lm:7b-alpha-q3_K_M", expected: "Starling Lm 7B Alpha Q3 K_M"},
		{input: "starling-lm:7b-alpha-q3_K_L", expected: "Starling Lm 7B Alpha Q3 K_L"},
		{input: "starling-lm:7b-alpha-q4_0", expected: "Starling Lm 7B Alpha Q4_0"},
		{input: "starling-lm:7b-alpha-q4_1", expected: "Starling Lm 7B Alpha Q4 1"},
		{input: "starling-lm:7b-alpha-q4_K_S", expected: "Starling Lm 7B Alpha Q4 K_S"},
		{input: "starling-lm:7b-alpha-q4_K_M", expected: "Starling Lm 7B Alpha Q4_K_M"},
		{input: "starling-lm:7b-alpha-q5_0", expected: "Starling Lm 7B Alpha Q5 0"},
		{input: "starling-lm:7b-alpha-q5_1", expected: "Starling Lm 7B Alpha Q5 1"},
		{input: "starling-lm:7b-alpha-q5_K_S", expected: "Starling Lm 7B Alpha Q5 K_S"},
		{input: "starling-lm:7b-alpha-q5_K_M", expected: "Starling Lm 7B Alpha Q5 K_M"},
		{input: "starling-lm:7b-alpha-q6_K", expected: "Starling Lm 7B Alpha Q6 K"},
		{input: "starling-lm:7b-alpha-q8_0", expected: "Starling Lm 7B Alpha Q8_0"},
		{input: "starling-lm:7b-alpha-fp16", expected: "Starling Lm 7B Alpha FP16"},
		{input: "starling-lm:7b-beta", expected: "Starling Lm 7B Beta"},
		{input: "starling-lm:7b-beta-q2_K", expected: "Starling Lm 7B Beta Q2_K"},
		{input: "starling-lm:7b-beta-q3_K_S", expected: "Starling Lm 7B Beta Q3 K_S"},
		{input: "starling-lm:7b-beta-q3_K_M", expected: "Starling Lm 7B Beta Q3 K_M"},
		{input: "starling-lm:7b-beta-q3_K_L", expected: "Starling Lm 7B Beta Q3 K_L"},
		{input: "starling-lm:7b-beta-q4_0", expected: "Starling Lm 7B Beta Q4_0"},
		{input: "starling-lm:7b-beta-q4_1", expected: "Starling Lm 7B Beta Q4 1"},
		{input: "starling-lm:7b-beta-q4_K_S", expected: "Starling Lm 7B Beta Q4 K_S"},
		{input: "starling-lm:7b-beta-q4_K_M", expected: "Starling Lm 7B Beta Q4_K_M"},
		{input: "starling-lm:7b-beta-q5_0", expected: "Starling Lm 7B Beta Q5 0"},
		{input: "starling-lm:7b-beta-q5_1", expected: "Starling Lm 7B Beta Q5 1"},
		{input: "starling-lm:7b-beta-q5_K_S", expected: "Starling Lm 7B Beta Q5 K_S"},
		{input: "starling-lm:7b-beta-q5_K_M", expected: "Starling Lm 7B Beta Q5 K_M"},
		{input: "starling-lm:7b-beta-q6_K", expected: "Starling Lm 7B Beta Q6 K"},
		{input: "starling-lm:7b-beta-q8_0", expected: "Starling Lm 7B Beta Q8_0"},
		{input: "starling-lm:7b-beta-fp16", expected: "Starling Lm 7B Beta FP16"},
		{input: "llama3-groq-tool-use:latest", expected: "Llama3 Groq Tool Use (latest)"},
		{input: "llama3-groq-tool-use", expected: "Llama3 Groq Tool Use"},
		{input: "llama3-groq-tool-use:8b", expected: "Llama3 Groq Tool Use 8B"},
		{input: "llama3-groq-tool-use:70b", expected: "Llama3 Groq Tool Use 70B"},
		{input: "llama3-groq-tool-use:8b-q2_K", expected: "Llama3 Groq Tool Use 8B Q2_K"},
		{input: "llama3-groq-tool-use:8b-q3_K_S", expected: "Llama3 Groq Tool Use 8B Q3 K_S"},
		{input: "llama3-groq-tool-use:8b-q3_K_M", expected: "Llama3 Groq Tool Use 8B Q3 K_M"},
		{input: "llama3-groq-tool-use:8b-q3_K_L", expected: "Llama3 Groq Tool Use 8B Q3 K_L"},
		{input: "llama3-groq-tool-use:8b-q4_0", expected: "Llama3 Groq Tool Use 8B Q4_0"},
		{input: "llama3-groq-tool-use:8b-q4_1", expected: "Llama3 Groq Tool Use 8B Q4 1"},
		{input: "llama3-groq-tool-use:8b-q4_K_S", expected: "Llama3 Groq Tool Use 8B Q4 K_S"},
		{input: "llama3-groq-tool-use:8b-q4_K_M", expected: "Llama3 Groq Tool Use 8B Q4_K_M"},
		{input: "llama3-groq-tool-use:8b-q5_0", expected: "Llama3 Groq Tool Use 8B Q5 0"},
		{input: "llama3-groq-tool-use:8b-q5_1", expected: "Llama3 Groq Tool Use 8B Q5 1"},
		{input: "llama3-groq-tool-use:8b-q5_K_S", expected: "Llama3 Groq Tool Use 8B Q5 K_S"},
		{input: "llama3-groq-tool-use:8b-q5_K_M", expected: "Llama3 Groq Tool Use 8B Q5 K_M"},
		{input: "llama3-groq-tool-use:8b-q6_K", expected: "Llama3 Groq Tool Use 8B Q6 K"},
		{input: "llama3-groq-tool-use:8b-q8_0", expected: "Llama3 Groq Tool Use 8B Q8_0"},
		{input: "llama3-groq-tool-use:8b-fp16", expected: "Llama3 Groq Tool Use 8B FP16"},
		{input: "llama3-groq-tool-use:70b-q2_K", expected: "Llama3 Groq Tool Use 70B Q2_K"},
		{input: "llama3-groq-tool-use:70b-q3_K_S", expected: "Llama3 Groq Tool Use 70B Q3 K_S"},
		{input: "llama3-groq-tool-use:70b-q3_K_M", expected: "Llama3 Groq Tool Use 70B Q3 K_M"},
		{input: "llama3-groq-tool-use:70b-q3_K_L", expected: "Llama3 Groq Tool Use 70B Q3 K_L"},
		{input: "llama3-groq-tool-use:70b-q4_0", expected: "Llama3 Groq Tool Use 70B Q4_0"},
		{input: "llama3-groq-tool-use:70b-q4_1", expected: "Llama3 Groq Tool Use 70B Q4 1"},
		{input: "llama3-groq-tool-use:70b-q4_K_S", expected: "Llama3 Groq Tool Use 70B Q4 K_S"},
		{input: "llama3-groq-tool-use:70b-q4_K_M", expected: "Llama3 Groq Tool Use 70B Q4_K_M"},
		{input: "llama3-groq-tool-use:70b-q5_0", expected: "Llama3 Groq Tool Use 70B Q5 0"},
		{input: "llama3-groq-tool-use:70b-q5_1", expected: "Llama3 Groq Tool Use 70B Q5 1"},
		{input: "llama3-groq-tool-use:70b-q5_K_S", expected: "Llama3 Groq Tool Use 70B Q5 K_S"},
		{input: "llama3-groq-tool-use:70b-q5_K_M", expected: "Llama3 Groq Tool Use 70B Q5 K_M"},
		{input: "llama3-groq-tool-use:70b-q6_K", expected: "Llama3 Groq Tool Use 70B Q6 K"},
		{input: "llama3-groq-tool-use:70b-q8_0", expected: "Llama3 Groq Tool Use 70B Q8_0"},
		{input: "llama3-groq-tool-use:70b-fp16", expected: "Llama3 Groq Tool Use 70B FP16"},
		{input: "athene-v2:latest", expected: "Athene v2 (latest)"},
		{input: "athene-v2", expected: "Athene v2"},
		{input: "athene-v2:72b", expected: "Athene v2 72B"},
		{input: "athene-v2:72b-q2_K", expected: "Athene v2 72B Q2_K"},
		{input: "athene-v2:72b-q3_K_S", expected: "Athene v2 72B Q3 K_S"},
		{input: "athene-v2:72b-q3_K_M", expected: "Athene v2 72B Q3 K_M"},
		{input: "athene-v2:72b-q3_K_L", expected: "Athene v2 72B Q3 K_L"},
		{input: "athene-v2:72b-q4_0", expected: "Athene v2 72B Q4_0"},
		{input: "athene-v2:72b-q4_1", expected: "Athene v2 72B Q4 1"},
		{input: "athene-v2:72b-q4_K_S", expected: "Athene v2 72B Q4 K_S"},
		{input: "athene-v2:72b-q4_K_M", expected: "Athene v2 72B Q4_K_M"},
		{input: "athene-v2:72b-q5_0", expected: "Athene v2 72B Q5 0"},
		{input: "athene-v2:72b-q5_1", expected: "Athene v2 72B Q5 1"},
		{input: "athene-v2:72b-q5_K_S", expected: "Athene v2 72B Q5 K_S"},
		{input: "athene-v2:72b-q5_K_M", expected: "Athene v2 72B Q5 K_M"},
		{input: "athene-v2:72b-q6_K", expected: "Athene v2 72B Q6 K"},
		{input: "athene-v2:72b-q8_0", expected: "Athene v2 72B Q8_0"},
		{input: "athene-v2:72b-fp16", expected: "Athene v2 72B FP16"},
		{input: "phind-codellama:latest", expected: "Phind Codellama (latest)"},
		{input: "phind-codellama", expected: "Phind Codellama"},
		{input: "phind-codellama:34b", expected: "Phind Codellama 34B"},
		{input: "phind-codellama:34b-python", expected: "Phind Codellama 34B Python"},
		{input: "phind-codellama:34b-python-q2_K", expected: "Phind Codellama 34B Python Q2_K"},
		{input: "phind-codellama:34b-python-q3_K_S", expected: "Phind Codellama 34B Python Q3 K_S"},
		{input: "phind-codellama:34b-python-q3_K_M", expected: "Phind Codellama 34B Python Q3 K_M"},
		{input: "phind-codellama:34b-python-q3_K_L", expected: "Phind Codellama 34B Python Q3 K_L"},
		{input: "phind-codellama:34b-python-q4_0", expected: "Phind Codellama 34B Python Q4_0"},
		{input: "phind-codellama:34b-python-q4_1", expected: "Phind Codellama 34B Python Q4 1"},
		{input: "phind-codellama:34b-python-q4_K_S", expected: "Phind Codellama 34B Python Q4 K_S"},
		{input: "phind-codellama:34b-python-q4_K_M", expected: "Phind Codellama 34B Python Q4_K_M"},
		{input: "phind-codellama:34b-python-q5_0", expected: "Phind Codellama 34B Python Q5 0"},
		{input: "phind-codellama:34b-python-q5_1", expected: "Phind Codellama 34B Python Q5 1"},
		{input: "phind-codellama:34b-python-q5_K_S", expected: "Phind Codellama 34B Python Q5 K_S"},
		{input: "phind-codellama:34b-python-q5_K_M", expected: "Phind Codellama 34B Python Q5 K_M"},
		{input: "phind-codellama:34b-python-q6_K", expected: "Phind Codellama 34B Python Q6 K"},
		{input: "phind-codellama:34b-python-q8_0", expected: "Phind Codellama 34B Python Q8_0"},
		{input: "phind-codellama:34b-python-fp16", expected: "Phind Codellama 34B Python FP16"},
		{input: "phind-codellama:34b-v2", expected: "Phind Codellama 34B v2"},
		{input: "phind-codellama:34b-v2-q2_K", expected: "Phind Codellama 34B v2 Q2_K"},
		{input: "phind-codellama:34b-v2-q3_K_S", expected: "Phind Codellama 34B v2 Q3 K_S"},
		{input: "phind-codellama:34b-v2-q3_K_M", expected: "Phind Codellama 34B v2 Q3 K_M"},
		{input: "phind-codellama:34b-v2-q3_K_L", expected: "Phind Codellama 34B v2 Q3 K_L"},
		{input: "phind-codellama:34b-v2-q4_0", expected: "Phind Codellama 34B v2 Q4_0"},
		{input: "phind-codellama:34b-v2-q4_1", expected: "Phind Codellama 34B v2 Q4 1"},
		{input: "phind-codellama:34b-v2-q4_K_S", expected: "Phind Codellama 34B v2 Q4 K_S"},
		{input: "phind-codellama:34b-v2-q4_K_M", expected: "Phind Codellama 34B v2 Q4_K_M"},
		{input: "phind-codellama:34b-v2-q5_0", expected: "Phind Codellama 34B v2 Q5 0"},
		{input: "phind-codellama:34b-v2-q5_1", expected: "Phind Codellama 34B v2 Q5 1"},
		{input: "phind-codellama:34b-v2-q5_K_S", expected: "Phind Codellama 34B v2 Q5 K_S"},
		{input: "phind-codellama:34b-v2-q5_K_M", expected: "Phind Codellama 34B v2 Q5 K_M"},
		{input: "phind-codellama:34b-v2-q6_K", expected: "Phind Codellama 34B v2 Q6 K"},
		{input: "phind-codellama:34b-v2-q8_0", expected: "Phind Codellama 34B v2 Q8_0"},
		{input: "phind-codellama:34b-v2-fp16", expected: "Phind Codellama 34B v2 FP16"},
		{input: "phind-codellama:34b-q2_K", expected: "Phind Codellama 34B Q2_K"},
		{input: "phind-codellama:34b-q3_K_S", expected: "Phind Codellama 34B Q3 K_S"},
		{input: "phind-codellama:34b-q3_K_M", expected: "Phind Codellama 34B Q3 K_M"},
		{input: "phind-codellama:34b-q3_K_L", expected: "Phind Codellama 34B Q3 K_L"},
		{input: "phind-codellama:34b-q4_0", expected: "Phind Codellama 34B Q4_0"},
		{input: "phind-codellama:34b-q4_1", expected: "Phind Codellama 34B Q4 1"},
		{input: "phind-codellama:34b-q4_K_S", expected: "Phind Codellama 34B Q4 K_S"},
		{input: "phind-codellama:34b-q4_K_M", expected: "Phind Codellama 34B Q4_K_M"},
		{input: "phind-codellama:34b-q5_0", expected: "Phind Codellama 34B Q5 0"},
		{input: "phind-codellama:34b-q5_1", expected: "Phind Codellama 34B Q5 1"},
		{input: "phind-codellama:34b-q5_K_S", expected: "Phind Codellama 34B Q5 K_S"},
		{input: "phind-codellama:34b-q5_K_M", expected: "Phind Codellama 34B Q5 K_M"},
		{input: "phind-codellama:34b-q6_K", expected: "Phind Codellama 34B Q6 K"},
		{input: "phind-codellama:34b-q8_0", expected: "Phind Codellama 34B Q8_0"},
		{input: "phind-codellama:34b-fp16", expected: "Phind Codellama 34B FP16"},
		{input: "solar:latest", expected: "Solar (latest)"},
		{input: "solar", expected: "Solar"},
		{input: "solar:10.7b", expected: "Solar 10.7B"},
		{input: "solar:10.7b-instruct-v1-q2_K", expected: "Solar 10.7B Instruct v1 Q2_K"},
		{input: "solar:10.7b-instruct-v1-q3_K_S", expected: "Solar 10.7B Instruct v1 Q3 K_S"},
		{input: "solar:10.7b-instruct-v1-q3_K_M", expected: "Solar 10.7B Instruct v1 Q3 K_M"},
		{input: "solar:10.7b-instruct-v1-q3_K_L", expected: "Solar 10.7B Instruct v1 Q3 K_L"},
		{input: "solar:10.7b-instruct-v1-q4_0", expected: "Solar 10.7B Instruct v1 Q4_0"},
		{input: "solar:10.7b-instruct-v1-q4_1", expected: "Solar 10.7B Instruct v1 Q4 1"},
		{input: "solar:10.7b-instruct-v1-q4_K_S", expected: "Solar 10.7B Instruct v1 Q4 K_S"},
		{input: "solar:10.7b-instruct-v1-q4_K_M", expected: "Solar 10.7B Instruct v1 Q4_K_M"},
		{input: "solar:10.7b-instruct-v1-q5_0", expected: "Solar 10.7B Instruct v1 Q5 0"},
		{input: "solar:10.7b-instruct-v1-q5_1", expected: "Solar 10.7B Instruct v1 Q5 1"},
		{input: "solar:10.7b-instruct-v1-q5_K_S", expected: "Solar 10.7B Instruct v1 Q5 K_S"},
		{input: "solar:10.7b-instruct-v1-q5_K_M", expected: "Solar 10.7B Instruct v1 Q5 K_M"},
		{input: "solar:10.7b-instruct-v1-q6_K", expected: "Solar 10.7B Instruct v1 Q6 K"},
		{input: "solar:10.7b-instruct-v1-q8_0", expected: "Solar 10.7B Instruct v1 Q8_0"},
		{input: "solar:10.7b-instruct-v1-fp16", expected: "Solar 10.7B Instruct v1 FP16"},
		{input: "solar:10.7b-text-v1-q2_K", expected: "Solar 10.7B Text v1 Q2_K"},
		{input: "solar:10.7b-text-v1-q3_K_S", expected: "Solar 10.7B Text v1 Q3 K_S"},
		{input: "solar:10.7b-text-v1-q3_K_M", expected: "Solar 10.7B Text v1 Q3 K_M"},
		{input: "solar:10.7b-text-v1-q3_K_L", expected: "Solar 10.7B Text v1 Q3 K_L"},
		{input: "solar:10.7b-text-v1-q4_0", expected: "Solar 10.7B Text v1 Q4_0"},
		{input: "solar:10.7b-text-v1-q4_1", expected: "Solar 10.7B Text v1 Q4 1"},
		{input: "solar:10.7b-text-v1-q4_K_S", expected: "Solar 10.7B Text v1 Q4 K_S"},
		{input: "solar:10.7b-text-v1-q4_K_M", expected: "Solar 10.7B Text v1 Q4_K_M"},
		{input: "solar:10.7b-text-v1-q5_0", expected: "Solar 10.7B Text v1 Q5 0"},
		{input: "solar:10.7b-text-v1-q5_1", expected: "Solar 10.7B Text v1 Q5 1"},
		{input: "solar:10.7b-text-v1-q5_K_S", expected: "Solar 10.7B Text v1 Q5 K_S"},
		{input: "solar:10.7b-text-v1-q5_K_M", expected: "Solar 10.7B Text v1 Q5 K_M"},
		{input: "solar:10.7b-text-v1-q6_K", expected: "Solar 10.7B Text v1 Q6 K"},
		{input: "solar:10.7b-text-v1-q8_0", expected: "Solar 10.7B Text v1 Q8_0"},
		{input: "solar:10.7b-text-v1-fp16", expected: "Solar 10.7B Text v1 FP16"},
		{input: "xwinlm:latest", expected: "Xwinlm (latest)"},
		{input: "xwinlm", expected: "Xwinlm"},
		{input: "xwinlm:7b", expected: "Xwinlm 7B"},
		{input: "xwinlm:13b", expected: "Xwinlm 13B"},
		{input: "xwinlm:7b-v0.1", expected: "Xwinlm 7B v0.1"},
		{input: "xwinlm:7b-v0.1-q2_K", expected: "Xwinlm 7B v0.1 Q2_K"},
		{input: "xwinlm:7b-v0.1-q3_K_S", expected: "Xwinlm 7B v0.1 Q3 K_S"},
		{input: "xwinlm:7b-v0.1-q3_K_M", expected: "Xwinlm 7B v0.1 Q3 K_M"},
		{input: "xwinlm:7b-v0.1-q3_K_L", expected: "Xwinlm 7B v0.1 Q3 K_L"},
		{input: "xwinlm:7b-v0.1-q4_0", expected: "Xwinlm 7B v0.1 Q4_0"},
		{input: "xwinlm:7b-v0.1-q4_1", expected: "Xwinlm 7B v0.1 Q4 1"},
		{input: "xwinlm:7b-v0.1-q4_K_S", expected: "Xwinlm 7B v0.1 Q4 K_S"},
		{input: "xwinlm:7b-v0.1-q4_K_M", expected: "Xwinlm 7B v0.1 Q4_K_M"},
		{input: "xwinlm:7b-v0.1-q5_0", expected: "Xwinlm 7B v0.1 Q5 0"},
		{input: "xwinlm:7b-v0.1-q5_1", expected: "Xwinlm 7B v0.1 Q5 1"},
		{input: "xwinlm:7b-v0.1-q5_K_S", expected: "Xwinlm 7B v0.1 Q5 K_S"},
		{input: "xwinlm:7b-v0.1-q5_K_M", expected: "Xwinlm 7B v0.1 Q5 K_M"},
		{input: "xwinlm:7b-v0.1-q6_K", expected: "Xwinlm 7B v0.1 Q6 K"},
		{input: "xwinlm:7b-v0.1-q8_0", expected: "Xwinlm 7B v0.1 Q8_0"},
		{input: "xwinlm:7b-v0.1-fp16", expected: "Xwinlm 7B v0.1 FP16"},
		{input: "xwinlm:7b-v0.2", expected: "Xwinlm 7B v0.2"},
		{input: "xwinlm:7b-v0.2-q2_K", expected: "Xwinlm 7B v0.2 Q2_K"},
		{input: "xwinlm:7b-v0.2-q3_K_S", expected: "Xwinlm 7B v0.2 Q3 K_S"},
		{input: "xwinlm:7b-v0.2-q3_K_L", expected: "Xwinlm 7B v0.2 Q3 K_L"},
		{input: "xwinlm:7b-v0.2-q4_0", expected: "Xwinlm 7B v0.2 Q4_0"},
		{input: "xwinlm:7b-v0.2-q4_1", expected: "Xwinlm 7B v0.2 Q4 1"},
		{input: "xwinlm:7b-v0.2-q4_K_S", expected: "Xwinlm 7B v0.2 Q4 K_S"},
		{input: "xwinlm:7b-v0.2-q4_K_M", expected: "Xwinlm 7B v0.2 Q4_K_M"},
		{input: "xwinlm:7b-v0.2-q5_0", expected: "Xwinlm 7B v0.2 Q5 0"},
		{input: "xwinlm:7b-v0.2-q5_K_S", expected: "Xwinlm 7B v0.2 Q5 K_S"},
		{input: "xwinlm:7b-v0.2-q5_K_M", expected: "Xwinlm 7B v0.2 Q5 K_M"},
		{input: "xwinlm:7b-v0.2-q6_K", expected: "Xwinlm 7B v0.2 Q6 K"},
		{input: "xwinlm:7b-v0.2-q8_0", expected: "Xwinlm 7B v0.2 Q8_0"},
		{input: "xwinlm:7b-v0.2-fp16", expected: "Xwinlm 7B v0.2 FP16"},
		{input: "xwinlm:13b-v0.1", expected: "Xwinlm 13B v0.1"},
		{input: "xwinlm:13b-v0.1-q2_K", expected: "Xwinlm 13B v0.1 Q2_K"},
		{input: "xwinlm:13b-v0.1-q3_K_S", expected: "Xwinlm 13B v0.1 Q3 K_S"},
		{input: "xwinlm:13b-v0.1-q3_K_M", expected: "Xwinlm 13B v0.1 Q3 K_M"},
		{input: "xwinlm:13b-v0.1-q3_K_L", expected: "Xwinlm 13B v0.1 Q3 K_L"},
		{input: "xwinlm:13b-v0.1-q4_0", expected: "Xwinlm 13B v0.1 Q4_0"},
		{input: "xwinlm:13b-v0.1-q4_1", expected: "Xwinlm 13B v0.1 Q4 1"},
		{input: "xwinlm:13b-v0.1-q4_K_S", expected: "Xwinlm 13B v0.1 Q4 K_S"},
		{input: "xwinlm:13b-v0.1-q4_K_M", expected: "Xwinlm 13B v0.1 Q4_K_M"},
		{input: "xwinlm:13b-v0.1-q5_0", expected: "Xwinlm 13B v0.1 Q5 0"},
		{input: "xwinlm:13b-v0.1-q5_1", expected: "Xwinlm 13B v0.1 Q5 1"},
		{input: "xwinlm:13b-v0.1-q5_K_S", expected: "Xwinlm 13B v0.1 Q5 K_S"},
		{input: "xwinlm:13b-v0.1-q5_K_M", expected: "Xwinlm 13B v0.1 Q5 K_M"},
		{input: "xwinlm:13b-v0.1-q6_K", expected: "Xwinlm 13B v0.1 Q6 K"},
		{input: "xwinlm:13b-v0.1-q8_0", expected: "Xwinlm 13B v0.1 Q8_0"},
		{input: "xwinlm:13b-v0.1-fp16", expected: "Xwinlm 13B v0.1 FP16"},
		{input: "xwinlm:13b-v0.2", expected: "Xwinlm 13B v0.2"},
		{input: "xwinlm:13b-v0.2-q2_K", expected: "Xwinlm 13B v0.2 Q2_K"},
		{input: "xwinlm:13b-v0.2-q3_K_S", expected: "Xwinlm 13B v0.2 Q3 K_S"},
		{input: "xwinlm:13b-v0.2-q3_K_M", expected: "Xwinlm 13B v0.2 Q3 K_M"},
		{input: "xwinlm:13b-v0.2-q3_K_L", expected: "Xwinlm 13B v0.2 Q3 K_L"},
		{input: "xwinlm:13b-v0.2-q4_0", expected: "Xwinlm 13B v0.2 Q4_0"},
		{input: "xwinlm:13b-v0.2-q4_1", expected: "Xwinlm 13B v0.2 Q4 1"},
		{input: "xwinlm:13b-v0.2-q4_K_S", expected: "Xwinlm 13B v0.2 Q4 K_S"},
		{input: "xwinlm:13b-v0.2-q4_K_M", expected: "Xwinlm 13B v0.2 Q4_K_M"},
		{input: "xwinlm:13b-v0.2-q5_0", expected: "Xwinlm 13B v0.2 Q5 0"},
		{input: "xwinlm:13b-v0.2-q5_1", expected: "Xwinlm 13B v0.2 Q5 1"},
		{input: "xwinlm:13b-v0.2-q5_K_S", expected: "Xwinlm 13B v0.2 Q5 K_S"},
		{input: "xwinlm:13b-v0.2-q5_K_M", expected: "Xwinlm 13B v0.2 Q5 K_M"},
		{input: "xwinlm:13b-v0.2-q6_K", expected: "Xwinlm 13B v0.2 Q6 K"},
		{input: "xwinlm:13b-v0.2-q8_0", expected: "Xwinlm 13B v0.2 Q8_0"},
		{input: "xwinlm:13b-v0.2-fp16", expected: "Xwinlm 13B v0.2 FP16"},
		{input: "xwinlm:70b-v0.1", expected: "Xwinlm 70B v0.1"},
		{input: "xwinlm:70b-v0.1-q2_K", expected: "Xwinlm 70B v0.1 Q2_K"},
		{input: "xwinlm:70b-v0.1-q3_K_S", expected: "Xwinlm 70B v0.1 Q3 K_S"},
		{input: "xwinlm:70b-v0.1-q3_K_M", expected: "Xwinlm 70B v0.1 Q3 K_M"},
		{input: "xwinlm:70b-v0.1-q3_K_L", expected: "Xwinlm 70B v0.1 Q3 K_L"},
		{input: "xwinlm:70b-v0.1-q4_0", expected: "Xwinlm 70B v0.1 Q4_0"},
		{input: "xwinlm:70b-v0.1-q4_1", expected: "Xwinlm 70B v0.1 Q4 1"},
		{input: "xwinlm:70b-v0.1-q4_K_S", expected: "Xwinlm 70B v0.1 Q4 K_S"},
		{input: "xwinlm:70b-v0.1-q4_K_M", expected: "Xwinlm 70B v0.1 Q4_K_M"},
		{input: "xwinlm:70b-v0.1-q5_0", expected: "Xwinlm 70B v0.1 Q5 0"},
		{input: "xwinlm:70b-v0.1-q5_1", expected: "Xwinlm 70B v0.1 Q5 1"},
		{input: "xwinlm:70b-v0.1-q5_K_S", expected: "Xwinlm 70B v0.1 Q5 K_S"},
		{input: "xwinlm:70b-v0.1-q6_K", expected: "Xwinlm 70B v0.1 Q6 K"},
		{input: "xwinlm:70b-v0.1-q8_0", expected: "Xwinlm 70B v0.1 Q8_0"},
		{input: "xwinlm:70b-v0.1-fp16", expected: "Xwinlm 70B v0.1 FP16"},
		{input: "nemotron:latest", expected: "Nemotron (latest)"},
		{input: "nemotron", expected: "Nemotron"},
		{input: "nemotron:70b", expected: "Nemotron 70B"},
		{input: "nemotron:70b-instruct-q2_K", expected: "Nemotron 70B Instruct Q2_K"},
		{input: "nemotron:70b-instruct-q3_K_S", expected: "Nemotron 70B Instruct Q3 K_S"},
		{input: "nemotron:70b-instruct-q3_K_M", expected: "Nemotron 70B Instruct Q3 K_M"},
		{input: "nemotron:70b-instruct-q3_K_L", expected: "Nemotron 70B Instruct Q3 K_L"},
		{input: "nemotron:70b-instruct-q4_0", expected: "Nemotron 70B Instruct Q4_0"},
		{input: "nemotron:70b-instruct-q4_1", expected: "Nemotron 70B Instruct Q4 1"},
		{input: "nemotron:70b-instruct-q4_K_S", expected: "Nemotron 70B Instruct Q4 K_S"},
		{input: "nemotron:70b-instruct-q4_K_M", expected: "Nemotron 70B Instruct Q4_K_M"},
		{input: "nemotron:70b-instruct-q5_0", expected: "Nemotron 70B Instruct Q5 0"},
		{input: "nemotron:70b-instruct-q5_1", expected: "Nemotron 70B Instruct Q5 1"},
		{input: "nemotron:70b-instruct-q5_K_S", expected: "Nemotron 70B Instruct Q5 K_S"},
		{input: "nemotron:70b-instruct-q5_K_M", expected: "Nemotron 70B Instruct Q5 K_M"},
		{input: "nemotron:70b-instruct-q6_K", expected: "Nemotron 70B Instruct Q6 K"},
		{input: "nemotron:70b-instruct-q8_0", expected: "Nemotron 70B Instruct Q8_0"},
		{input: "nemotron:70b-instruct-fp16", expected: "Nemotron 70B Instruct FP16"},
		{input: "llama-guard3:latest", expected: "Llama Guard3 (latest)"},
		{input: "llama-guard3", expected: "Llama Guard3"},
		{input: "llama-guard3:1b", expected: "Llama Guard3 1B"},
		{input: "llama-guard3:8b", expected: "Llama Guard3 8B"},
		{input: "llama-guard3:1b-q2_K", expected: "Llama Guard3 1B Q2_K"},
		{input: "llama-guard3:1b-q3_K_S", expected: "Llama Guard3 1B Q3 K_S"},
		{input: "llama-guard3:1b-q3_K_M", expected: "Llama Guard3 1B Q3 K_M"},
		{input: "llama-guard3:1b-q3_K_L", expected: "Llama Guard3 1B Q3 K_L"},
		{input: "llama-guard3:1b-q4_0", expected: "Llama Guard3 1B Q4_0"},
		{input: "llama-guard3:1b-q4_1", expected: "Llama Guard3 1B Q4 1"},
		{input: "llama-guard3:1b-q4_K_S", expected: "Llama Guard3 1B Q4 K_S"},
		{input: "llama-guard3:1b-q4_K_M", expected: "Llama Guard3 1B Q4_K_M"},
		{input: "llama-guard3:1b-q5_0", expected: "Llama Guard3 1B Q5 0"},
		{input: "llama-guard3:1b-q5_1", expected: "Llama Guard3 1B Q5 1"},
		{input: "llama-guard3:1b-q5_K_S", expected: "Llama Guard3 1B Q5 K_S"},
		{input: "llama-guard3:1b-q5_K_M", expected: "Llama Guard3 1B Q5 K_M"},
		{input: "llama-guard3:1b-q6_K", expected: "Llama Guard3 1B Q6 K"},
		{input: "llama-guard3:1b-q8_0", expected: "Llama Guard3 1B Q8_0"},
		{input: "llama-guard3:1b-fp16", expected: "Llama Guard3 1B FP16"},
		{input: "llama-guard3:8b-q2_K", expected: "Llama Guard3 8B Q2_K"},
		{input: "llama-guard3:8b-q3_K_S", expected: "Llama Guard3 8B Q3 K_S"},
		{input: "llama-guard3:8b-q3_K_M", expected: "Llama Guard3 8B Q3 K_M"},
		{input: "llama-guard3:8b-q3_K_L", expected: "Llama Guard3 8B Q3 K_L"},
		{input: "llama-guard3:8b-q4_0", expected: "Llama Guard3 8B Q4_0"},
		{input: "llama-guard3:8b-q4_1", expected: "Llama Guard3 8B Q4 1"},
		{input: "llama-guard3:8b-q4_K_S", expected: "Llama Guard3 8B Q4 K_S"},
		{input: "llama-guard3:8b-q4_K_M", expected: "Llama Guard3 8B Q4_K_M"},
		{input: "llama-guard3:8b-q5_0", expected: "Llama Guard3 8B Q5 0"},
		{input: "llama-guard3:8b-q5_1", expected: "Llama Guard3 8B Q5 1"},
		{input: "llama-guard3:8b-q5_K_S", expected: "Llama Guard3 8B Q5 K_S"},
		{input: "llama-guard3:8b-q5_K_M", expected: "Llama Guard3 8B Q5 K_M"},
		{input: "llama-guard3:8b-q6_K", expected: "Llama Guard3 8B Q6 K"},
		{input: "llama-guard3:8b-q8_0", expected: "Llama Guard3 8B Q8_0"},
		{input: "llama-guard3:8b-fp16", expected: "Llama Guard3 8B FP16"},
		{input: "meditron:latest", expected: "Meditron (latest)"},
		{input: "meditron", expected: "Meditron"},
		{input: "meditron:7b", expected: "Meditron 7B"},
		{input: "meditron:70b", expected: "Meditron 70B"},
		{input: "meditron:7b-q2_K", expected: "Meditron 7B Q2_K"},
		{input: "meditron:7b-q3_K_S", expected: "Meditron 7B Q3 K_S"},
		{input: "meditron:7b-q3_K_M", expected: "Meditron 7B Q3 K_M"},
		{input: "meditron:7b-q3_K_L", expected: "Meditron 7B Q3 K_L"},
		{input: "meditron:7b-q4_0", expected: "Meditron 7B Q4_0"},
		{input: "meditron:7b-q4_1", expected: "Meditron 7B Q4 1"},
		{input: "meditron:7b-q4_K_S", expected: "Meditron 7B Q4 K_S"},
		{input: "meditron:7b-q4_K_M", expected: "Meditron 7B Q4_K_M"},
		{input: "meditron:7b-q5_0", expected: "Meditron 7B Q5 0"},
		{input: "meditron:7b-q5_1", expected: "Meditron 7B Q5 1"},
		{input: "meditron:7b-q5_K_S", expected: "Meditron 7B Q5 K_S"},
		{input: "meditron:7b-q5_K_M", expected: "Meditron 7B Q5 K_M"},
		{input: "meditron:7b-q6_K", expected: "Meditron 7B Q6 K"},
		{input: "meditron:7b-q8_0", expected: "Meditron 7B Q8_0"},
		{input: "meditron:7b-fp16", expected: "Meditron 7B FP16"},
		{input: "meditron:70b-q4_0", expected: "Meditron 70B Q4_0"},
		{input: "meditron:70b-q4_1", expected: "Meditron 70B Q4 1"},
		{input: "meditron:70b-q4_K_S", expected: "Meditron 70B Q4 K_S"},
		{input: "meditron:70b-q5_1", expected: "Meditron 70B Q5 1"},
		{input: "yarn-llama2:latest", expected: "Yarn Llama2 (latest)"},
		{input: "yarn-llama2", expected: "Yarn Llama2"},
		{input: "yarn-llama2:7b", expected: "Yarn Llama2 7B"},
		{input: "yarn-llama2:13b", expected: "Yarn Llama2 13B"},
		{input: "yarn-llama2:7b-128k", expected: "Yarn Llama2 7B 128K"},
		{input: "yarn-llama2:7b-128k-q2_K", expected: "Yarn Llama2 7B 128K Q2_K"},
		{input: "yarn-llama2:7b-128k-q3_K_S", expected: "Yarn Llama2 7B 128K Q3 K_S"},
		{input: "yarn-llama2:7b-128k-q3_K_M", expected: "Yarn Llama2 7B 128K Q3 K_M"},
		{input: "yarn-llama2:7b-128k-q3_K_L", expected: "Yarn Llama2 7B 128K Q3 K_L"},
		{input: "yarn-llama2:7b-128k-q4_0", expected: "Yarn Llama2 7B 128K Q4_0"},
		{input: "yarn-llama2:7b-128k-q4_1", expected: "Yarn Llama2 7B 128K Q4 1"},
		{input: "yarn-llama2:7b-128k-q4_K_S", expected: "Yarn Llama2 7B 128K Q4 K_S"},
		{input: "yarn-llama2:7b-128k-q4_K_M", expected: "Yarn Llama2 7B 128K Q4_K_M"},
		{input: "yarn-llama2:7b-128k-q5_0", expected: "Yarn Llama2 7B 128K Q5 0"},
		{input: "yarn-llama2:7b-128k-q5_1", expected: "Yarn Llama2 7B 128K Q5 1"},
		{input: "yarn-llama2:7b-128k-q5_K_S", expected: "Yarn Llama2 7B 128K Q5 K_S"},
		{input: "yarn-llama2:7b-128k-q5_K_M", expected: "Yarn Llama2 7B 128K Q5 K_M"},
		{input: "yarn-llama2:7b-128k-q6_K", expected: "Yarn Llama2 7B 128K Q6 K"},
		{input: "yarn-llama2:7b-128k-q8_0", expected: "Yarn Llama2 7B 128K Q8_0"},
		{input: "yarn-llama2:7b-128k-fp16", expected: "Yarn Llama2 7B 128K FP16"},
		{input: "yarn-llama2:7b-64k", expected: "Yarn Llama2 7B 64K"},
		{input: "yarn-llama2:7b-64k-q2_K", expected: "Yarn Llama2 7B 64K Q2_K"},
		{input: "yarn-llama2:7b-64k-q3_K_S", expected: "Yarn Llama2 7B 64K Q3 K_S"},
		{input: "yarn-llama2:7b-64k-q3_K_M", expected: "Yarn Llama2 7B 64K Q3 K_M"},
		{input: "yarn-llama2:7b-64k-q3_K_L", expected: "Yarn Llama2 7B 64K Q3 K_L"},
		{input: "yarn-llama2:7b-64k-q4_0", expected: "Yarn Llama2 7B 64K Q4_0"},
		{input: "yarn-llama2:7b-64k-q4_1", expected: "Yarn Llama2 7B 64K Q4 1"},
		{input: "yarn-llama2:7b-64k-q4_K_S", expected: "Yarn Llama2 7B 64K Q4 K_S"},
		{input: "yarn-llama2:7b-64k-q4_K_M", expected: "Yarn Llama2 7B 64K Q4_K_M"},
		{input: "yarn-llama2:7b-64k-q5_0", expected: "Yarn Llama2 7B 64K Q5 0"},
		{input: "yarn-llama2:7b-64k-q5_1", expected: "Yarn Llama2 7B 64K Q5 1"},
		{input: "yarn-llama2:7b-64k-q5_K_S", expected: "Yarn Llama2 7B 64K Q5 K_S"},
		{input: "yarn-llama2:7b-64k-q5_K_M", expected: "Yarn Llama2 7B 64K Q5 K_M"},
		{input: "yarn-llama2:7b-64k-q6_K", expected: "Yarn Llama2 7B 64K Q6 K"},
		{input: "yarn-llama2:7b-64k-q8_0", expected: "Yarn Llama2 7B 64K Q8_0"},
		{input: "yarn-llama2:7b-64k-fp16", expected: "Yarn Llama2 7B 64K FP16"},
		{input: "yarn-llama2:13b-128k", expected: "Yarn Llama2 13B 128K"},
		{input: "yarn-llama2:13b-128k-q2_K", expected: "Yarn Llama2 13B 128K Q2_K"},
		{input: "yarn-llama2:13b-128k-q3_K_S", expected: "Yarn Llama2 13B 128K Q3 K_S"},
		{input: "yarn-llama2:13b-128k-q3_K_M", expected: "Yarn Llama2 13B 128K Q3 K_M"},
		{input: "yarn-llama2:13b-128k-q3_K_L", expected: "Yarn Llama2 13B 128K Q3 K_L"},
		{input: "yarn-llama2:13b-128k-q4_0", expected: "Yarn Llama2 13B 128K Q4_0"},
		{input: "yarn-llama2:13b-128k-q4_1", expected: "Yarn Llama2 13B 128K Q4 1"},
		{input: "yarn-llama2:13b-128k-q4_K_S", expected: "Yarn Llama2 13B 128K Q4 K_S"},
		{input: "yarn-llama2:13b-128k-q4_K_M", expected: "Yarn Llama2 13B 128K Q4_K_M"},
		{input: "yarn-llama2:13b-128k-q5_0", expected: "Yarn Llama2 13B 128K Q5 0"},
		{input: "yarn-llama2:13b-128k-q5_1", expected: "Yarn Llama2 13B 128K Q5 1"},
		{input: "yarn-llama2:13b-128k-q5_K_S", expected: "Yarn Llama2 13B 128K Q5 K_S"},
		{input: "yarn-llama2:13b-128k-q5_K_M", expected: "Yarn Llama2 13B 128K Q5 K_M"},
		{input: "yarn-llama2:13b-128k-q6_K", expected: "Yarn Llama2 13B 128K Q6 K"},
		{input: "yarn-llama2:13b-128k-q8_0", expected: "Yarn Llama2 13B 128K Q8_0"},
		{input: "yarn-llama2:13b-128k-fp16", expected: "Yarn Llama2 13B 128K FP16"},
		{input: "yarn-llama2:13b-64k", expected: "Yarn Llama2 13B 64K"},
		{input: "yarn-llama2:13b-64k-q2_K", expected: "Yarn Llama2 13B 64K Q2_K"},
		{input: "yarn-llama2:13b-64k-q3_K_S", expected: "Yarn Llama2 13B 64K Q3 K_S"},
		{input: "yarn-llama2:13b-64k-q3_K_M", expected: "Yarn Llama2 13B 64K Q3 K_M"},
		{input: "yarn-llama2:13b-64k-q3_K_L", expected: "Yarn Llama2 13B 64K Q3 K_L"},
		{input: "yarn-llama2:13b-64k-q4_0", expected: "Yarn Llama2 13B 64K Q4_0"},
		{input: "yarn-llama2:13b-64k-q4_1", expected: "Yarn Llama2 13B 64K Q4 1"},
		{input: "yarn-llama2:13b-64k-q4_K_S", expected: "Yarn Llama2 13B 64K Q4 K_S"},
		{input: "yarn-llama2:13b-64k-q4_K_M", expected: "Yarn Llama2 13B 64K Q4_K_M"},
		{input: "yarn-llama2:13b-64k-q5_0", expected: "Yarn Llama2 13B 64K Q5 0"},
		{input: "yarn-llama2:13b-64k-q5_1", expected: "Yarn Llama2 13B 64K Q5 1"},
		{input: "yarn-llama2:13b-64k-q5_K_S", expected: "Yarn Llama2 13B 64K Q5 K_S"},
		{input: "yarn-llama2:13b-64k-q5_K_M", expected: "Yarn Llama2 13B 64K Q5 K_M"},
		{input: "yarn-llama2:13b-64k-q6_K", expected: "Yarn Llama2 13B 64K Q6 K"},
		{input: "yarn-llama2:13b-64k-q8_0", expected: "Yarn Llama2 13B 64K Q8_0"},
		{input: "yarn-llama2:13b-64k-fp16", expected: "Yarn Llama2 13B 64K FP16"},
		{input: "aya-expanse:latest", expected: "Aya Expanse (latest)"},
		{input: "aya-expanse", expected: "Aya Expanse"},
		{input: "aya-expanse:8b", expected: "Aya Expanse 8B"},
		{input: "aya-expanse:32b", expected: "Aya Expanse 32B"},
		{input: "aya-expanse:8b-q2_K", expected: "Aya Expanse 8B Q2_K"},
		{input: "aya-expanse:8b-q3_K_S", expected: "Aya Expanse 8B Q3 K_S"},
		{input: "aya-expanse:8b-q3_K_M", expected: "Aya Expanse 8B Q3 K_M"},
		{input: "aya-expanse:8b-q3_K_L", expected: "Aya Expanse 8B Q3 K_L"},
		{input: "aya-expanse:8b-q4_0", expected: "Aya Expanse 8B Q4_0"},
		{input: "aya-expanse:8b-q4_1", expected: "Aya Expanse 8B Q4 1"},
		{input: "aya-expanse:8b-q4_K_S", expected: "Aya Expanse 8B Q4 K_S"},
		{input: "aya-expanse:8b-q4_K_M", expected: "Aya Expanse 8B Q4_K_M"},
		{input: "aya-expanse:8b-q5_0", expected: "Aya Expanse 8B Q5 0"},
		{input: "aya-expanse:8b-q5_1", expected: "Aya Expanse 8B Q5 1"},
		{input: "aya-expanse:8b-q5_K_S", expected: "Aya Expanse 8B Q5 K_S"},
		{input: "aya-expanse:8b-q5_K_M", expected: "Aya Expanse 8B Q5 K_M"},
		{input: "aya-expanse:8b-q6_K", expected: "Aya Expanse 8B Q6 K"},
		{input: "aya-expanse:8b-q8_0", expected: "Aya Expanse 8B Q8_0"},
		{input: "aya-expanse:8b-fp16", expected: "Aya Expanse 8B FP16"},
		{input: "aya-expanse:32b-q2_K", expected: "Aya Expanse 32B Q2_K"},
		{input: "aya-expanse:32b-q3_K_S", expected: "Aya Expanse 32B Q3 K_S"},
		{input: "aya-expanse:32b-q3_K_M", expected: "Aya Expanse 32B Q3 K_M"},
		{input: "aya-expanse:32b-q3_K_L", expected: "Aya Expanse 32B Q3 K_L"},
		{input: "aya-expanse:32b-q4_0", expected: "Aya Expanse 32B Q4_0"},
		{input: "aya-expanse:32b-q4_1", expected: "Aya Expanse 32B Q4 1"},
		{input: "aya-expanse:32b-q4_K_S", expected: "Aya Expanse 32B Q4 K_S"},
		{input: "aya-expanse:32b-q4_K_M", expected: "Aya Expanse 32B Q4_K_M"},
		{input: "aya-expanse:32b-q5_0", expected: "Aya Expanse 32B Q5 0"},
		{input: "aya-expanse:32b-q5_1", expected: "Aya Expanse 32B Q5 1"},
		{input: "aya-expanse:32b-q5_K_S", expected: "Aya Expanse 32B Q5 K_S"},
		{input: "aya-expanse:32b-q5_K_M", expected: "Aya Expanse 32B Q5 K_M"},
		{input: "aya-expanse:32b-q6_K", expected: "Aya Expanse 32B Q6 K"},
		{input: "aya-expanse:32b-q8_0", expected: "Aya Expanse 32B Q8_0"},
		{input: "aya-expanse:32b-fp16", expected: "Aya Expanse 32B FP16"},
		{input: "wizardlm-uncensored:latest", expected: "WizardLM Uncensored (latest)"},
		{input: "wizardlm-uncensored", expected: "WizardLM Uncensored"},
		{input: "wizardlm-uncensored:13b", expected: "WizardLM Uncensored 13B"},
		{input: "wizardlm-uncensored:13b-llama2", expected: "WizardLM Uncensored 13B Llama2"},
		{input: "wizardlm-uncensored:13b-llama2-q2_K", expected: "WizardLM Uncensored 13B Llama2 Q2_K"},
		{input: "wizardlm-uncensored:13b-llama2-q3_K_S", expected: "WizardLM Uncensored 13B Llama2 Q3 K_S"},
		{input: "wizardlm-uncensored:13b-llama2-q3_K_M", expected: "WizardLM Uncensored 13B Llama2 Q3 K_M"},
		{input: "wizardlm-uncensored:13b-llama2-q3_K_L", expected: "WizardLM Uncensored 13B Llama2 Q3 K_L"},
		{input: "wizardlm-uncensored:13b-llama2-q4_0", expected: "WizardLM Uncensored 13B Llama2 Q4_0"},
		{input: "wizardlm-uncensored:13b-llama2-q4_1", expected: "WizardLM Uncensored 13B Llama2 Q4 1"},
		{input: "wizardlm-uncensored:13b-llama2-q4_K_S", expected: "WizardLM Uncensored 13B Llama2 Q4 K_S"},
		{input: "wizardlm-uncensored:13b-llama2-q4_K_M", expected: "WizardLM Uncensored 13B Llama2 Q4_K_M"},
		{input: "wizardlm-uncensored:13b-llama2-q5_0", expected: "WizardLM Uncensored 13B Llama2 Q5 0"},
		{input: "wizardlm-uncensored:13b-llama2-q5_1", expected: "WizardLM Uncensored 13B Llama2 Q5 1"},
		{input: "wizardlm-uncensored:13b-llama2-q5_K_S", expected: "WizardLM Uncensored 13B Llama2 Q5 K_S"},
		{input: "wizardlm-uncensored:13b-llama2-q5_K_M", expected: "WizardLM Uncensored 13B Llama2 Q5 K_M"},
		{input: "wizardlm-uncensored:13b-llama2-q6_K", expected: "WizardLM Uncensored 13B Llama2 Q6 K"},
		{input: "wizardlm-uncensored:13b-llama2-q8_0", expected: "WizardLM Uncensored 13B Llama2 Q8_0"},
		{input: "wizardlm-uncensored:13b-llama2-fp16", expected: "WizardLM Uncensored 13B Llama2 FP16"},
		{input: "granite3-moe:latest", expected: "Granite3 Moe (latest)"},
		{input: "granite3-moe", expected: "Granite3 Moe"},
		{input: "granite3-moe:1b", expected: "Granite3 Moe 1B"},
		{input: "granite3-moe:3b", expected: "Granite3 Moe 3B"},
		{input: "granite3-moe:1b-instruct-q2_K", expected: "Granite3 Moe 1B Instruct Q2_K"},
		{input: "granite3-moe:1b-instruct-q3_K_S", expected: "Granite3 Moe 1B Instruct Q3 K_S"},
		{input: "granite3-moe:1b-instruct-q3_K_M", expected: "Granite3 Moe 1B Instruct Q3 K_M"},
		{input: "granite3-moe:1b-instruct-q3_K_L", expected: "Granite3 Moe 1B Instruct Q3 K_L"},
		{input: "granite3-moe:1b-instruct-q4_0", expected: "Granite3 Moe 1B Instruct Q4_0"},
		{input: "granite3-moe:1b-instruct-q4_1", expected: "Granite3 Moe 1B Instruct Q4 1"},
		{input: "granite3-moe:1b-instruct-q4_K_S", expected: "Granite3 Moe 1B Instruct Q4 K_S"},
		{input: "granite3-moe:1b-instruct-q4_K_M", expected: "Granite3 Moe 1B Instruct Q4_K_M"},
		{input: "granite3-moe:1b-instruct-q5_0", expected: "Granite3 Moe 1B Instruct Q5 0"},
		{input: "granite3-moe:1b-instruct-q5_1", expected: "Granite3 Moe 1B Instruct Q5 1"},
		{input: "granite3-moe:1b-instruct-q5_K_S", expected: "Granite3 Moe 1B Instruct Q5 K_S"},
		{input: "granite3-moe:1b-instruct-q5_K_M", expected: "Granite3 Moe 1B Instruct Q5 K_M"},
		{input: "granite3-moe:1b-instruct-q6_K", expected: "Granite3 Moe 1B Instruct Q6 K"},
		{input: "granite3-moe:1b-instruct-q8_0", expected: "Granite3 Moe 1B Instruct Q8_0"},
		{input: "granite3-moe:1b-instruct-fp16", expected: "Granite3 Moe 1B Instruct FP16"},
		{input: "granite3-moe:3b-instruct-q2_K", expected: "Granite3 Moe 3B Instruct Q2_K"},
		{input: "granite3-moe:3b-instruct-q3_K_S", expected: "Granite3 Moe 3B Instruct Q3 K_S"},
		{input: "granite3-moe:3b-instruct-q3_K_M", expected: "Granite3 Moe 3B Instruct Q3 K_M"},
		{input: "granite3-moe:3b-instruct-q3_K_L", expected: "Granite3 Moe 3B Instruct Q3 K_L"},
		{input: "granite3-moe:3b-instruct-q4_0", expected: "Granite3 Moe 3B Instruct Q4_0"},
		{input: "granite3-moe:3b-instruct-q4_1", expected: "Granite3 Moe 3B Instruct Q4 1"},
		{input: "granite3-moe:3b-instruct-q4_K_S", expected: "Granite3 Moe 3B Instruct Q4 K_S"},
		{input: "granite3-moe:3b-instruct-q4_K_M", expected: "Granite3 Moe 3B Instruct Q4_K_M"},
		{input: "granite3-moe:3b-instruct-q5_0", expected: "Granite3 Moe 3B Instruct Q5 0"},
		{input: "granite3-moe:3b-instruct-q5_1", expected: "Granite3 Moe 3B Instruct Q5 1"},
		{input: "granite3-moe:3b-instruct-q5_K_S", expected: "Granite3 Moe 3B Instruct Q5 K_S"},
		{input: "granite3-moe:3b-instruct-q5_K_M", expected: "Granite3 Moe 3B Instruct Q5 K_M"},
		{input: "granite3-moe:3b-instruct-q6_K", expected: "Granite3 Moe 3B Instruct Q6 K"},
		{input: "granite3-moe:3b-instruct-q8_0", expected: "Granite3 Moe 3B Instruct Q8_0"},
		{input: "granite3-moe:3b-instruct-fp16", expected: "Granite3 Moe 3B Instruct FP16"},
		{input: "smallthinker:latest", expected: "Smallthinker (latest)"},
		{input: "smallthinker", expected: "Smallthinker"},
		{input: "smallthinker:3b", expected: "Smallthinker 3B"},
		{input: "smallthinker:3b-preview-q4_K_M", expected: "Smallthinker 3B Preview Q4_K_M"},
		{input: "smallthinker:3b-preview-q8_0", expected: "Smallthinker 3B Preview Q8_0"},
		{input: "smallthinker:3b-preview-fp16", expected: "Smallthinker 3B Preview FP16"},
		{input: "orca2:latest", expected: "Orca2 (latest)"},
		{input: "orca2", expected: "Orca2"},
		{input: "orca2:7b", expected: "Orca2 7B"},
		{input: "orca2:13b", expected: "Orca2 13B"},
		{input: "orca2:7b-q2_K", expected: "Orca2 7B Q2_K"},
		{input: "orca2:7b-q3_K_S", expected: "Orca2 7B Q3 K_S"},
		{input: "orca2:7b-q3_K_M", expected: "Orca2 7B Q3 K_M"},
		{input: "orca2:7b-q3_K_L", expected: "Orca2 7B Q3 K_L"},
		{input: "orca2:7b-q4_0", expected: "Orca2 7B Q4_0"},
		{input: "orca2:7b-q4_1", expected: "Orca2 7B Q4 1"},
		{input: "orca2:7b-q4_K_S", expected: "Orca2 7B Q4 K_S"},
		{input: "orca2:7b-q4_K_M", expected: "Orca2 7B Q4_K_M"},
		{input: "orca2:7b-q5_0", expected: "Orca2 7B Q5 0"},
		{input: "orca2:7b-q5_1", expected: "Orca2 7B Q5 1"},
		{input: "orca2:7b-q5_K_S", expected: "Orca2 7B Q5 K_S"},
		{input: "orca2:7b-q5_K_M", expected: "Orca2 7B Q5 K_M"},
		{input: "orca2:7b-q6_K", expected: "Orca2 7B Q6 K"},
		{input: "orca2:7b-q8_0", expected: "Orca2 7B Q8_0"},
		{input: "orca2:7b-fp16", expected: "Orca2 7B FP16"},
		{input: "orca2:13b-q2_K", expected: "Orca2 13B Q2_K"},
		{input: "orca2:13b-q3_K_S", expected: "Orca2 13B Q3 K_S"},
		{input: "orca2:13b-q3_K_M", expected: "Orca2 13B Q3 K_M"},
		{input: "orca2:13b-q3_K_L", expected: "Orca2 13B Q3 K_L"},
		{input: "orca2:13b-q4_0", expected: "Orca2 13B Q4_0"},
		{input: "orca2:13b-q4_1", expected: "Orca2 13B Q4 1"},
		{input: "orca2:13b-q4_K_S", expected: "Orca2 13B Q4 K_S"},
		{input: "orca2:13b-q4_K_M", expected: "Orca2 13B Q4_K_M"},
		{input: "orca2:13b-q5_0", expected: "Orca2 13B Q5 0"},
		{input: "orca2:13b-q5_1", expected: "Orca2 13B Q5 1"},
		{input: "orca2:13b-q5_K_S", expected: "Orca2 13B Q5 K_S"},
		{input: "orca2:13b-q5_K_M", expected: "Orca2 13B Q5 K_M"},
		{input: "orca2:13b-q6_K", expected: "Orca2 13B Q6 K"},
		{input: "orca2:13b-q8_0", expected: "Orca2 13B Q8_0"},
		{input: "orca2:13b-fp16", expected: "Orca2 13B FP16"},
		{input: "medllama2:latest", expected: "Medllama2 (latest)"},
		{input: "medllama2", expected: "Medllama2"},
		{input: "medllama2:7b", expected: "Medllama2 7B"},
		{input: "medllama2:7b-q2_K", expected: "Medllama2 7B Q2_K"},
		{input: "medllama2:7b-q3_K_S", expected: "Medllama2 7B Q3 K_S"},
		{input: "medllama2:7b-q3_K_M", expected: "Medllama2 7B Q3 K_M"},
		{input: "medllama2:7b-q3_K_L", expected: "Medllama2 7B Q3 K_L"},
		{input: "medllama2:7b-q4_0", expected: "Medllama2 7B Q4_0"},
		{input: "medllama2:7b-q4_1", expected: "Medllama2 7B Q4 1"},
		{input: "medllama2:7b-q4_K_S", expected: "Medllama2 7B Q4 K_S"},
		{input: "medllama2:7b-q4_K_M", expected: "Medllama2 7B Q4_K_M"},
		{input: "medllama2:7b-q5_0", expected: "Medllama2 7B Q5 0"},
		{input: "medllama2:7b-q5_1", expected: "Medllama2 7B Q5 1"},
		{input: "medllama2:7b-q5_K_S", expected: "Medllama2 7B Q5 K_S"},
		{input: "medllama2:7b-q5_K_M", expected: "Medllama2 7B Q5 K_M"},
		{input: "medllama2:7b-q6_K", expected: "Medllama2 7B Q6 K"},
		{input: "medllama2:7b-q8_0", expected: "Medllama2 7B Q8_0"},
		{input: "medllama2:7b-fp16", expected: "Medllama2 7B FP16"},
		{input: "command-r7b:latest", expected: "Command R7B (latest)"},
		{input: "command-r7b", expected: "Command R7B"},
		{input: "command-r7b:7b", expected: "Command R7B 7B"},
		{input: "command-r7b:7b-12-2024-q4_K_M", expected: "Command R7B 7B (2024-12) Q4_K_M"},
		{input: "command-r7b:7b-12-2024-q8_0", expected: "Command R7B 7B (2024-12) Q8_0"},
		{input: "command-r7b:7b-12-2024-fp16", expected: "Command R7B 7B (2024-12) FP16"},
		{input: "phi4-mini-reasoning:latest", expected: "Phi4 Mini Reasoning (latest)"},
		{input: "phi4-mini-reasoning", expected: "Phi4 Mini Reasoning"},
		{input: "phi4-mini-reasoning:3.8b", expected: "Phi4 Mini Reasoning 3.8B"},
		{input: "phi4-mini-reasoning:3.8b-q4_K_M", expected: "Phi4 Mini Reasoning 3.8B Q4_K_M"},
		{input: "phi4-mini-reasoning:3.8b-q8_0", expected: "Phi4 Mini Reasoning 3.8B Q8_0"},
		{input: "phi4-mini-reasoning:3.8b-fp16", expected: "Phi4 Mini Reasoning 3.8B FP16"},
		{input: "nous-hermes2-mixtral:latest", expected: "Nous Hermes2 Mixtral (latest)"},
		{input: "nous-hermes2-mixtral", expected: "Nous Hermes2 Mixtral"},
		{input: "nous-hermes2-mixtral:dpo", expected: "Nous Hermes2 Mixtral Dpo"},
		{input: "nous-hermes2-mixtral:8x7b", expected: "Nous Hermes2 Mixtral 8x7B"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q2_K", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q2_K"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q3_K_S", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q3 K_S"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q3_K_M", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q3 K_M"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q3_K_L", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q3 K_L"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q4_0", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q4_0"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q4_1", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q4 1"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q4_K_S", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q4 K_S"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q4_K_M", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q4_K_M"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q5_0", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q5 0"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q5_1", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q5 1"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q5_K_S", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q5 K_S"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q5_K_M", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q5 K_M"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q6_K", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q6 K"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-q8_0", expected: "Nous Hermes2 Mixtral 8x7B Dpo Q8_0"},
		{input: "nous-hermes2-mixtral:8x7b-dpo-fp16", expected: "Nous Hermes2 Mixtral 8x7B Dpo FP16"},
		{input: "stable-beluga:latest", expected: "Stable Beluga (latest)"},
		{input: "stable-beluga", expected: "Stable Beluga"},
		{input: "stable-beluga:7b", expected: "Stable Beluga 7B"},
		{input: "stable-beluga:13b", expected: "Stable Beluga 13B"},
		{input: "stable-beluga:70b", expected: "Stable Beluga 70B"},
		{input: "stable-beluga:7b-q2_K", expected: "Stable Beluga 7B Q2_K"},
		{input: "stable-beluga:7b-q3_K_S", expected: "Stable Beluga 7B Q3 K_S"},
		{input: "stable-beluga:7b-q3_K_M", expected: "Stable Beluga 7B Q3 K_M"},
		{input: "stable-beluga:7b-q3_K_L", expected: "Stable Beluga 7B Q3 K_L"},
		{input: "stable-beluga:7b-q4_0", expected: "Stable Beluga 7B Q4_0"},
		{input: "stable-beluga:7b-q4_1", expected: "Stable Beluga 7B Q4 1"},
		{input: "stable-beluga:7b-q4_K_S", expected: "Stable Beluga 7B Q4 K_S"},
		{input: "stable-beluga:7b-q4_K_M", expected: "Stable Beluga 7B Q4_K_M"},
		{input: "stable-beluga:7b-q5_0", expected: "Stable Beluga 7B Q5 0"},
		{input: "stable-beluga:7b-q5_1", expected: "Stable Beluga 7B Q5 1"},
		{input: "stable-beluga:7b-q5_K_S", expected: "Stable Beluga 7B Q5 K_S"},
		{input: "stable-beluga:7b-q5_K_M", expected: "Stable Beluga 7B Q5 K_M"},
		{input: "stable-beluga:7b-q6_K", expected: "Stable Beluga 7B Q6 K"},
		{input: "stable-beluga:7b-q8_0", expected: "Stable Beluga 7B Q8_0"},
		{input: "stable-beluga:7b-fp16", expected: "Stable Beluga 7B FP16"},
		{input: "stable-beluga:13b-q2_K", expected: "Stable Beluga 13B Q2_K"},
		{input: "stable-beluga:13b-q3_K_S", expected: "Stable Beluga 13B Q3 K_S"},
		{input: "stable-beluga:13b-q3_K_M", expected: "Stable Beluga 13B Q3 K_M"},
		{input: "stable-beluga:13b-q3_K_L", expected: "Stable Beluga 13B Q3 K_L"},
		{input: "stable-beluga:13b-q4_0", expected: "Stable Beluga 13B Q4_0"},
		{input: "stable-beluga:13b-q4_1", expected: "Stable Beluga 13B Q4 1"},
		{input: "stable-beluga:13b-q4_K_S", expected: "Stable Beluga 13B Q4 K_S"},
		{input: "stable-beluga:13b-q4_K_M", expected: "Stable Beluga 13B Q4_K_M"},
		{input: "stable-beluga:13b-q5_0", expected: "Stable Beluga 13B Q5 0"},
		{input: "stable-beluga:13b-q5_1", expected: "Stable Beluga 13B Q5 1"},
		{input: "stable-beluga:13b-q5_K_S", expected: "Stable Beluga 13B Q5 K_S"},
		{input: "stable-beluga:13b-q5_K_M", expected: "Stable Beluga 13B Q5 K_M"},
		{input: "stable-beluga:13b-q6_K", expected: "Stable Beluga 13B Q6 K"},
		{input: "stable-beluga:13b-q8_0", expected: "Stable Beluga 13B Q8_0"},
		{input: "stable-beluga:13b-fp16", expected: "Stable Beluga 13B FP16"},
		{input: "stable-beluga:70b-q2_K", expected: "Stable Beluga 70B Q2_K"},
		{input: "stable-beluga:70b-q3_K_S", expected: "Stable Beluga 70B Q3 K_S"},
		{input: "stable-beluga:70b-q3_K_M", expected: "Stable Beluga 70B Q3 K_M"},
		{input: "stable-beluga:70b-q3_K_L", expected: "Stable Beluga 70B Q3 K_L"},
		{input: "stable-beluga:70b-q4_0", expected: "Stable Beluga 70B Q4_0"},
		{input: "stable-beluga:70b-q4_1", expected: "Stable Beluga 70B Q4 1"},
		{input: "stable-beluga:70b-q4_K_S", expected: "Stable Beluga 70B Q4 K_S"},
		{input: "stable-beluga:70b-q4_K_M", expected: "Stable Beluga 70B Q4_K_M"},
		{input: "stable-beluga:70b-q5_0", expected: "Stable Beluga 70B Q5 0"},
		{input: "stable-beluga:70b-q5_1", expected: "Stable Beluga 70B Q5 1"},
		{input: "stable-beluga:70b-q5_K_S", expected: "Stable Beluga 70B Q5 K_S"},
		{input: "stable-beluga:70b-q5_K_M", expected: "Stable Beluga 70B Q5 K_M"},
		{input: "stable-beluga:70b-q6_K", expected: "Stable Beluga 70B Q6 K"},
		{input: "stable-beluga:70b-q8_0", expected: "Stable Beluga 70B Q8_0"},
		{input: "stable-beluga:70b-fp16", expected: "Stable Beluga 70B FP16"},
		{input: "deepseek-v2.5:latest", expected: "Deepseek v2.5 (latest)"},
		{input: "deepseek-v2.5", expected: "Deepseek v2.5"},
		{input: "deepseek-v2.5:236b", expected: "Deepseek v2.5 236B"},
		{input: "deepseek-v2.5:236b-q4_0", expected: "Deepseek v2.5 236B Q4_0"},
		{input: "deepseek-v2.5:236b-q4_1", expected: "Deepseek v2.5 236B Q4 1"},
		{input: "deepseek-v2.5:236b-q5_0", expected: "Deepseek v2.5 236B Q5 0"},
		{input: "deepseek-v2.5:236b-q5_1", expected: "Deepseek v2.5 236B Q5 1"},
		{input: "deepseek-v2.5:236b-q8_0", expected: "Deepseek v2.5 236B Q8_0"},
		{input: "reader-lm:latest", expected: "Reader Lm (latest)"},
		{input: "reader-lm", expected: "Reader Lm"},
		{input: "reader-lm:0.5b", expected: "Reader Lm 0.5B"},
		{input: "reader-lm:1.5b", expected: "Reader Lm 1.5B"},
		{input: "reader-lm:0.5b-q2_K", expected: "Reader Lm 0.5B Q2_K"},
		{input: "reader-lm:0.5b-q3_K_S", expected: "Reader Lm 0.5B Q3 K_S"},
		{input: "reader-lm:0.5b-q3_K_M", expected: "Reader Lm 0.5B Q3 K_M"},
		{input: "reader-lm:0.5b-q3_K_L", expected: "Reader Lm 0.5B Q3 K_L"},
		{input: "reader-lm:0.5b-q4_0", expected: "Reader Lm 0.5B Q4_0"},
		{input: "reader-lm:0.5b-q4_1", expected: "Reader Lm 0.5B Q4 1"},
		{input: "reader-lm:0.5b-q4_K_S", expected: "Reader Lm 0.5B Q4 K_S"},
		{input: "reader-lm:0.5b-q4_K_M", expected: "Reader Lm 0.5B Q4_K_M"},
		{input: "reader-lm:0.5b-q5_0", expected: "Reader Lm 0.5B Q5 0"},
		{input: "reader-lm:0.5b-q5_1", expected: "Reader Lm 0.5B Q5 1"},
		{input: "reader-lm:0.5b-q5_K_S", expected: "Reader Lm 0.5B Q5 K_S"},
		{input: "reader-lm:0.5b-q5_K_M", expected: "Reader Lm 0.5B Q5 K_M"},
		{input: "reader-lm:0.5b-q6_K", expected: "Reader Lm 0.5B Q6 K"},
		{input: "reader-lm:0.5b-q8_0", expected: "Reader Lm 0.5B Q8_0"},
		{input: "reader-lm:0.5b-fp16", expected: "Reader Lm 0.5B FP16"},
		{input: "reader-lm:1.5b-q2_K", expected: "Reader Lm 1.5B Q2_K"},
		{input: "reader-lm:1.5b-q3_K_S", expected: "Reader Lm 1.5B Q3 K_S"},
		{input: "reader-lm:1.5b-q3_K_M", expected: "Reader Lm 1.5B Q3 K_M"},
		{input: "reader-lm:1.5b-q3_K_L", expected: "Reader Lm 1.5B Q3 K_L"},
		{input: "reader-lm:1.5b-q4_0", expected: "Reader Lm 1.5B Q4_0"},
		{input: "reader-lm:1.5b-q4_1", expected: "Reader Lm 1.5B Q4 1"},
		{input: "reader-lm:1.5b-q4_K_S", expected: "Reader Lm 1.5B Q4 K_S"},
		{input: "reader-lm:1.5b-q4_K_M", expected: "Reader Lm 1.5B Q4_K_M"},
		{input: "reader-lm:1.5b-q5_0", expected: "Reader Lm 1.5B Q5 0"},
		{input: "reader-lm:1.5b-q5_1", expected: "Reader Lm 1.5B Q5 1"},
		{input: "reader-lm:1.5b-q5_K_S", expected: "Reader Lm 1.5B Q5 K_S"},
		{input: "reader-lm:1.5b-q5_K_M", expected: "Reader Lm 1.5B Q5 K_M"},
		{input: "reader-lm:1.5b-q6_K", expected: "Reader Lm 1.5B Q6 K"},
		{input: "reader-lm:1.5b-q8_0", expected: "Reader Lm 1.5B Q8_0"},
		{input: "reader-lm:1.5b-fp16", expected: "Reader Lm 1.5B FP16"},
		{input: "shieldgemma:latest", expected: "Shieldgemma (latest)"},
		{input: "shieldgemma", expected: "Shieldgemma"},
		{input: "shieldgemma:2b", expected: "Shieldgemma 2B"},
		{input: "shieldgemma:9b", expected: "Shieldgemma 9B"},
		{input: "shieldgemma:27b", expected: "Shieldgemma 27B"},
		{input: "shieldgemma:2b-q2_K", expected: "Shieldgemma 2B Q2_K"},
		{input: "shieldgemma:2b-q3_K_S", expected: "Shieldgemma 2B Q3 K_S"},
		{input: "shieldgemma:2b-q3_K_M", expected: "Shieldgemma 2B Q3 K_M"},
		{input: "shieldgemma:2b-q3_K_L", expected: "Shieldgemma 2B Q3 K_L"},
		{input: "shieldgemma:2b-q4_0", expected: "Shieldgemma 2B Q4_0"},
		{input: "shieldgemma:2b-q4_1", expected: "Shieldgemma 2B Q4 1"},
		{input: "shieldgemma:2b-q4_K_S", expected: "Shieldgemma 2B Q4 K_S"},
		{input: "shieldgemma:2b-q4_K_M", expected: "Shieldgemma 2B Q4_K_M"},
		{input: "shieldgemma:2b-q5_0", expected: "Shieldgemma 2B Q5 0"},
		{input: "shieldgemma:2b-q5_1", expected: "Shieldgemma 2B Q5 1"},
		{input: "shieldgemma:2b-q5_K_S", expected: "Shieldgemma 2B Q5 K_S"},
		{input: "shieldgemma:2b-q5_K_M", expected: "Shieldgemma 2B Q5 K_M"},
		{input: "shieldgemma:2b-q6_K", expected: "Shieldgemma 2B Q6 K"},
		{input: "shieldgemma:2b-q8_0", expected: "Shieldgemma 2B Q8_0"},
		{input: "shieldgemma:2b-fp16", expected: "Shieldgemma 2B FP16"},
		{input: "shieldgemma:9b-q2_K", expected: "Shieldgemma 9B Q2_K"},
		{input: "shieldgemma:9b-q3_K_S", expected: "Shieldgemma 9B Q3 K_S"},
		{input: "shieldgemma:9b-q3_K_M", expected: "Shieldgemma 9B Q3 K_M"},
		{input: "shieldgemma:9b-q3_K_L", expected: "Shieldgemma 9B Q3 K_L"},
		{input: "shieldgemma:9b-q4_0", expected: "Shieldgemma 9B Q4_0"},
		{input: "shieldgemma:9b-q4_1", expected: "Shieldgemma 9B Q4 1"},
		{input: "shieldgemma:9b-q4_K_S", expected: "Shieldgemma 9B Q4 K_S"},
		{input: "shieldgemma:9b-q4_K_M", expected: "Shieldgemma 9B Q4_K_M"},
		{input: "shieldgemma:9b-q5_0", expected: "Shieldgemma 9B Q5 0"},
		{input: "shieldgemma:9b-q5_1", expected: "Shieldgemma 9B Q5 1"},
		{input: "shieldgemma:9b-q5_K_S", expected: "Shieldgemma 9B Q5 K_S"},
		{input: "shieldgemma:9b-q5_K_M", expected: "Shieldgemma 9B Q5 K_M"},
		{input: "shieldgemma:9b-q6_K", expected: "Shieldgemma 9B Q6 K"},
		{input: "shieldgemma:9b-q8_0", expected: "Shieldgemma 9B Q8_0"},
		{input: "shieldgemma:9b-fp16", expected: "Shieldgemma 9B FP16"},
		{input: "shieldgemma:27b-q2_K", expected: "Shieldgemma 27B Q2_K"},
		{input: "shieldgemma:27b-q3_K_S", expected: "Shieldgemma 27B Q3 K_S"},
		{input: "shieldgemma:27b-q3_K_M", expected: "Shieldgemma 27B Q3 K_M"},
		{input: "shieldgemma:27b-q3_K_L", expected: "Shieldgemma 27B Q3 K_L"},
		{input: "shieldgemma:27b-q4_0", expected: "Shieldgemma 27B Q4_0"},
		{input: "shieldgemma:27b-q4_1", expected: "Shieldgemma 27B Q4 1"},
		{input: "shieldgemma:27b-q4_K_S", expected: "Shieldgemma 27B Q4 K_S"},
		{input: "shieldgemma:27b-q4_K_M", expected: "Shieldgemma 27B Q4_K_M"},
		{input: "shieldgemma:27b-q5_0", expected: "Shieldgemma 27B Q5 0"},
		{input: "shieldgemma:27b-q5_1", expected: "Shieldgemma 27B Q5 1"},
		{input: "shieldgemma:27b-q5_K_S", expected: "Shieldgemma 27B Q5 K_S"},
		{input: "shieldgemma:27b-q5_K_M", expected: "Shieldgemma 27B Q5 K_M"},
		{input: "shieldgemma:27b-q6_K", expected: "Shieldgemma 27B Q6 K"},
		{input: "shieldgemma:27b-q8_0", expected: "Shieldgemma 27B Q8_0"},
		{input: "shieldgemma:27b-fp16", expected: "Shieldgemma 27B FP16"},
		{input: "command-a:latest", expected: "Command A (latest)"},
		{input: "command-a", expected: "Command A"},
		{input: "command-a:111b", expected: "Command A 111B"},
		{input: "command-a:111b-03-2025-q4_K_M", expected: "Command A 111B (2025-03) Q4_K_M"},
		{input: "command-a:111b-03-2025-q8_0", expected: "Command A 111B (2025-03) Q8_0"},
		{input: "command-a:111b-03-2025-fp16", expected: "Command A 111B (2025-03) FP16"},
		{input: "llama-pro:latest", expected: "Llama Pro (latest)"},
		{input: "llama-pro", expected: "Llama Pro"},
		{input: "llama-pro:instruct", expected: "Llama Pro Instruct"},
		{input: "llama-pro:text", expected: "Llama Pro Text"},
		{input: "llama-pro:8b-instruct-q2_K", expected: "Llama Pro 8B Instruct Q2_K"},
		{input: "llama-pro:8b-instruct-q3_K_S", expected: "Llama Pro 8B Instruct Q3 K_S"},
		{input: "llama-pro:8b-instruct-q3_K_M", expected: "Llama Pro 8B Instruct Q3 K_M"},
		{input: "llama-pro:8b-instruct-q3_K_L", expected: "Llama Pro 8B Instruct Q3 K_L"},
		{input: "llama-pro:8b-instruct-q4_0", expected: "Llama Pro 8B Instruct Q4_0"},
		{input: "llama-pro:8b-instruct-q4_1", expected: "Llama Pro 8B Instruct Q4 1"},
		{input: "llama-pro:8b-instruct-q4_K_S", expected: "Llama Pro 8B Instruct Q4 K_S"},
		{input: "llama-pro:8b-instruct-q4_K_M", expected: "Llama Pro 8B Instruct Q4_K_M"},
		{input: "llama-pro:8b-instruct-q5_0", expected: "Llama Pro 8B Instruct Q5 0"},
		{input: "llama-pro:8b-instruct-q5_1", expected: "Llama Pro 8B Instruct Q5 1"},
		{input: "llama-pro:8b-instruct-q5_K_S", expected: "Llama Pro 8B Instruct Q5 K_S"},
		{input: "llama-pro:8b-instruct-q5_K_M", expected: "Llama Pro 8B Instruct Q5 K_M"},
		{input: "llama-pro:8b-instruct-q6_K", expected: "Llama Pro 8B Instruct Q6 K"},
		{input: "llama-pro:8b-instruct-q8_0", expected: "Llama Pro 8B Instruct Q8_0"},
		{input: "llama-pro:8b-instruct-fp16", expected: "Llama Pro 8B Instruct FP16"},
		{input: "llama-pro:8b-text-q2_K", expected: "Llama Pro 8B Text Q2_K"},
		{input: "llama-pro:8b-text-q3_K_S", expected: "Llama Pro 8B Text Q3 K_S"},
		{input: "llama-pro:8b-text-q3_K_M", expected: "Llama Pro 8B Text Q3 K_M"},
		{input: "llama-pro:8b-text-q3_K_L", expected: "Llama Pro 8B Text Q3 K_L"},
		{input: "llama-pro:8b-text-q4_0", expected: "Llama Pro 8B Text Q4_0"},
		{input: "llama-pro:8b-text-q4_1", expected: "Llama Pro 8B Text Q4 1"},
		{input: "llama-pro:8b-text-q4_K_S", expected: "Llama Pro 8B Text Q4 K_S"},
		{input: "llama-pro:8b-text-q4_K_M", expected: "Llama Pro 8B Text Q4_K_M"},
		{input: "llama-pro:8b-text-q5_0", expected: "Llama Pro 8B Text Q5 0"},
		{input: "llama-pro:8b-text-q5_1", expected: "Llama Pro 8B Text Q5 1"},
		{input: "llama-pro:8b-text-q5_K_S", expected: "Llama Pro 8B Text Q5 K_S"},
		{input: "llama-pro:8b-text-q5_K_M", expected: "Llama Pro 8B Text Q5 K_M"},
		{input: "llama-pro:8b-text-q6_K", expected: "Llama Pro 8B Text Q6 K"},
		{input: "llama-pro:8b-text-q8_0", expected: "Llama Pro 8B Text Q8_0"},
		{input: "llama-pro:8b-text-fp16", expected: "Llama Pro 8B Text FP16"},
		{input: "mathstral:latest", expected: "Mathstral (latest)"},
		{input: "mathstral", expected: "Mathstral"},
		{input: "mathstral:7b", expected: "Mathstral 7B"},
		{input: "mathstral:7b-v0.1-q2_K", expected: "Mathstral 7B v0.1 Q2_K"},
		{input: "mathstral:7b-v0.1-q3_K_S", expected: "Mathstral 7B v0.1 Q3 K_S"},
		{input: "mathstral:7b-v0.1-q3_K_M", expected: "Mathstral 7B v0.1 Q3 K_M"},
		{input: "mathstral:7b-v0.1-q3_K_L", expected: "Mathstral 7B v0.1 Q3 K_L"},
		{input: "mathstral:7b-v0.1-q4_0", expected: "Mathstral 7B v0.1 Q4_0"},
		{input: "mathstral:7b-v0.1-q4_1", expected: "Mathstral 7B v0.1 Q4 1"},
		{input: "mathstral:7b-v0.1-q4_K_S", expected: "Mathstral 7B v0.1 Q4 K_S"},
		{input: "mathstral:7b-v0.1-q4_K_M", expected: "Mathstral 7B v0.1 Q4_K_M"},
		{input: "mathstral:7b-v0.1-q5_0", expected: "Mathstral 7B v0.1 Q5 0"},
		{input: "mathstral:7b-v0.1-q5_1", expected: "Mathstral 7B v0.1 Q5 1"},
		{input: "mathstral:7b-v0.1-q5_K_S", expected: "Mathstral 7B v0.1 Q5 K_S"},
		{input: "mathstral:7b-v0.1-q5_K_M", expected: "Mathstral 7B v0.1 Q5 K_M"},
		{input: "mathstral:7b-v0.1-q6_K", expected: "Mathstral 7B v0.1 Q6 K"},
		{input: "mathstral:7b-v0.1-q8_0", expected: "Mathstral 7B v0.1 Q8_0"},
		{input: "mathstral:7b-v0.1-fp16", expected: "Mathstral 7B v0.1 FP16"},
		{input: "wizardlm:7b-q2_K", expected: "WizardLM 7B Q2_K"},
		{input: "wizardlm:7b-q3_K_S", expected: "WizardLM 7B Q3 K_S"},
		{input: "wizardlm:7b-q3_K_M", expected: "WizardLM 7B Q3 K_M"},
		{input: "wizardlm:7b-q3_K_L", expected: "WizardLM 7B Q3 K_L"},
		{input: "wizardlm:7b-q4_0", expected: "WizardLM 7B Q4_0"},
		{input: "wizardlm:7b-q4_1", expected: "WizardLM 7B Q4 1"},
		{input: "wizardlm:7b-q4_K_S", expected: "WizardLM 7B Q4 K_S"},
		{input: "wizardlm:7b-q4_K_M", expected: "WizardLM 7B Q4_K_M"},
		{input: "wizardlm:7b-q5_0", expected: "WizardLM 7B Q5 0"},
		{input: "wizardlm:7b-q5_1", expected: "WizardLM 7B Q5 1"},
		{input: "wizardlm:7b-q5_K_S", expected: "WizardLM 7B Q5 K_S"},
		{input: "wizardlm:7b-q5_K_M", expected: "WizardLM 7B Q5 K_M"},
		{input: "wizardlm:7b-q6_K", expected: "WizardLM 7B Q6 K"},
		{input: "wizardlm:7b-q8_0", expected: "WizardLM 7B Q8_0"},
		{input: "wizardlm:7b-fp16", expected: "WizardLM 7B FP16"},
		{input: "wizardlm:13b-llama2-q2_K", expected: "WizardLM 13B Llama2 Q2_K"},
		{input: "wizardlm:13b-llama2-q3_K_S", expected: "WizardLM 13B Llama2 Q3 K_S"},
		{input: "wizardlm:13b-llama2-q3_K_M", expected: "WizardLM 13B Llama2 Q3 K_M"},
		{input: "wizardlm:13b-llama2-q3_K_L", expected: "WizardLM 13B Llama2 Q3 K_L"},
		{input: "wizardlm:13b-llama2-q4_0", expected: "WizardLM 13B Llama2 Q4_0"},
		{input: "wizardlm:13b-llama2-q4_1", expected: "WizardLM 13B Llama2 Q4 1"},
		{input: "wizardlm:13b-llama2-q4_K_S", expected: "WizardLM 13B Llama2 Q4 K_S"},
		{input: "wizardlm:13b-llama2-q4_K_M", expected: "WizardLM 13B Llama2 Q4_K_M"},
		{input: "wizardlm:13b-llama2-q5_0", expected: "WizardLM 13B Llama2 Q5 0"},
		{input: "wizardlm:13b-llama2-q5_1", expected: "WizardLM 13B Llama2 Q5 1"},
		{input: "wizardlm:13b-llama2-q5_K_S", expected: "WizardLM 13B Llama2 Q5 K_S"},
		{input: "wizardlm:13b-llama2-q5_K_M", expected: "WizardLM 13B Llama2 Q5 K_M"},
		{input: "wizardlm:13b-llama2-q6_K", expected: "WizardLM 13B Llama2 Q6 K"},
		{input: "wizardlm:13b-llama2-q8_0", expected: "WizardLM 13B Llama2 Q8_0"},
		{input: "wizardlm:13b-llama2-fp16", expected: "WizardLM 13B Llama2 FP16"},
		{input: "wizardlm:13b-q2_K", expected: "WizardLM 13B Q2_K"},
		{input: "wizardlm:13b-q3_K_S", expected: "WizardLM 13B Q3 K_S"},
		{input: "wizardlm:13b-q3_K_M", expected: "WizardLM 13B Q3 K_M"},
		{input: "wizardlm:13b-q3_K_L", expected: "WizardLM 13B Q3 K_L"},
		{input: "wizardlm:13b-q4_0", expected: "WizardLM 13B Q4_0"},
		{input: "wizardlm:13b-q4_1", expected: "WizardLM 13B Q4 1"},
		{input: "wizardlm:13b-q4_K_S", expected: "WizardLM 13B Q4 K_S"},
		{input: "wizardlm:13b-q4_K_M", expected: "WizardLM 13B Q4_K_M"},
		{input: "wizardlm:13b-q5_0", expected: "WizardLM 13B Q5 0"},
		{input: "wizardlm:13b-q5_1", expected: "WizardLM 13B Q5 1"},
		{input: "wizardlm:13b-q5_K_S", expected: "WizardLM 13B Q5 K_S"},
		{input: "wizardlm:13b-q5_K_M", expected: "WizardLM 13B Q5 K_M"},
		{input: "wizardlm:13b-q6_K", expected: "WizardLM 13B Q6 K"},
		{input: "wizardlm:13b-q8_0", expected: "WizardLM 13B Q8_0"},
		{input: "wizardlm:13b-fp16", expected: "WizardLM 13B FP16"},
		{input: "wizardlm:30b-q2_K", expected: "WizardLM 30B Q2_K"},
		{input: "wizardlm:30b-q3_K_S", expected: "WizardLM 30B Q3 K_S"},
		{input: "wizardlm:30b-q3_K_M", expected: "WizardLM 30B Q3 K_M"},
		{input: "wizardlm:30b-q3_K_L", expected: "WizardLM 30B Q3 K_L"},
		{input: "wizardlm:30b-q4_0", expected: "WizardLM 30B Q4_0"},
		{input: "wizardlm:30b-q4_1", expected: "WizardLM 30B Q4 1"},
		{input: "wizardlm:30b-q4_K_S", expected: "WizardLM 30B Q4 K_S"},
		{input: "wizardlm:30b-q4_K_M", expected: "WizardLM 30B Q4_K_M"},
		{input: "wizardlm:30b-q5_0", expected: "WizardLM 30B Q5 0"},
		{input: "wizardlm:30b-q5_1", expected: "WizardLM 30B Q5 1"},
		{input: "wizardlm:30b-q5_K_S", expected: "WizardLM 30B Q5 K_S"},
		{input: "wizardlm:30b-q5_K_M", expected: "WizardLM 30B Q5 K_M"},
		{input: "wizardlm:30b-q6_K", expected: "WizardLM 30B Q6 K"},
		{input: "wizardlm:30b-q8_0", expected: "WizardLM 30B Q8_0"},
		{input: "wizardlm:30b-fp16", expected: "WizardLM 30B FP16"},
		{input: "wizardlm:70b-llama2-q2_K", expected: "WizardLM 70B Llama2 Q2_K"},
		{input: "wizardlm:70b-llama2-q3_K_S", expected: "WizardLM 70B Llama2 Q3 K_S"},
		{input: "wizardlm:70b-llama2-q3_K_M", expected: "WizardLM 70B Llama2 Q3 K_M"},
		{input: "wizardlm:70b-llama2-q3_K_L", expected: "WizardLM 70B Llama2 Q3 K_L"},
		{input: "wizardlm:70b-llama2-q4_0", expected: "WizardLM 70B Llama2 Q4_0"},
		{input: "wizardlm:70b-llama2-q4_1", expected: "WizardLM 70B Llama2 Q4 1"},
		{input: "wizardlm:70b-llama2-q4_K_S", expected: "WizardLM 70B Llama2 Q4 K_S"},
		{input: "wizardlm:70b-llama2-q4_K_M", expected: "WizardLM 70B Llama2 Q4_K_M"},
		{input: "wizardlm:70b-llama2-q5_0", expected: "WizardLM 70B Llama2 Q5 0"},
		{input: "wizardlm:70b-llama2-q5_K_S", expected: "WizardLM 70B Llama2 Q5 K_S"},
		{input: "wizardlm:70b-llama2-q5_K_M", expected: "WizardLM 70B Llama2 Q5 K_M"},
		{input: "wizardlm:70b-llama2-q6_K", expected: "WizardLM 70B Llama2 Q6 K"},
		{input: "wizardlm:70b-llama2-q8_0", expected: "WizardLM 70B Llama2 Q8_0"},
		{input: "yarn-mistral:latest", expected: "Yarn Mistral (latest)"},
		{input: "yarn-mistral", expected: "Yarn Mistral"},
		{input: "yarn-mistral:7b", expected: "Yarn Mistral 7B"},
		{input: "yarn-mistral:7b-128k", expected: "Yarn Mistral 7B 128K"},
		{input: "yarn-mistral:7b-128k-q2_K", expected: "Yarn Mistral 7B 128K Q2_K"},
		{input: "yarn-mistral:7b-128k-q3_K_S", expected: "Yarn Mistral 7B 128K Q3 K_S"},
		{input: "yarn-mistral:7b-128k-q3_K_M", expected: "Yarn Mistral 7B 128K Q3 K_M"},
		{input: "yarn-mistral:7b-128k-q3_K_L", expected: "Yarn Mistral 7B 128K Q3 K_L"},
		{input: "yarn-mistral:7b-128k-q4_0", expected: "Yarn Mistral 7B 128K Q4_0"},
		{input: "yarn-mistral:7b-128k-q4_1", expected: "Yarn Mistral 7B 128K Q4 1"},
		{input: "yarn-mistral:7b-128k-q4_K_S", expected: "Yarn Mistral 7B 128K Q4 K_S"},
		{input: "yarn-mistral:7b-128k-q4_K_M", expected: "Yarn Mistral 7B 128K Q4_K_M"},
		{input: "yarn-mistral:7b-128k-q5_0", expected: "Yarn Mistral 7B 128K Q5 0"},
		{input: "yarn-mistral:7b-128k-q5_1", expected: "Yarn Mistral 7B 128K Q5 1"},
		{input: "yarn-mistral:7b-128k-q5_K_S", expected: "Yarn Mistral 7B 128K Q5 K_S"},
		{input: "yarn-mistral:7b-128k-q5_K_M", expected: "Yarn Mistral 7B 128K Q5 K_M"},
		{input: "yarn-mistral:7b-128k-q6_K", expected: "Yarn Mistral 7B 128K Q6 K"},
		{input: "yarn-mistral:7b-128k-q8_0", expected: "Yarn Mistral 7B 128K Q8_0"},
		{input: "yarn-mistral:7b-128k-fp16", expected: "Yarn Mistral 7B 128K FP16"},
		{input: "yarn-mistral:7b-64k", expected: "Yarn Mistral 7B 64K"},
		{input: "yarn-mistral:7b-64k-q2_K", expected: "Yarn Mistral 7B 64K Q2_K"},
		{input: "yarn-mistral:7b-64k-q3_K_S", expected: "Yarn Mistral 7B 64K Q3 K_S"},
		{input: "yarn-mistral:7b-64k-q3_K_M", expected: "Yarn Mistral 7B 64K Q3 K_M"},
		{input: "yarn-mistral:7b-64k-q3_K_L", expected: "Yarn Mistral 7B 64K Q3 K_L"},
		{input: "yarn-mistral:7b-64k-q4_0", expected: "Yarn Mistral 7B 64K Q4_0"},
		{input: "yarn-mistral:7b-64k-q4_1", expected: "Yarn Mistral 7B 64K Q4 1"},
		{input: "yarn-mistral:7b-64k-q4_K_S", expected: "Yarn Mistral 7B 64K Q4 K_S"},
		{input: "yarn-mistral:7b-64k-q4_K_M", expected: "Yarn Mistral 7B 64K Q4_K_M"},
		{input: "yarn-mistral:7b-64k-q5_0", expected: "Yarn Mistral 7B 64K Q5 0"},
		{input: "yarn-mistral:7b-64k-q5_1", expected: "Yarn Mistral 7B 64K Q5 1"},
		{input: "yarn-mistral:7b-64k-q5_K_S", expected: "Yarn Mistral 7B 64K Q5 K_S"},
		{input: "yarn-mistral:7b-64k-q5_K_M", expected: "Yarn Mistral 7B 64K Q5 K_M"},
		{input: "yarn-mistral:7b-64k-q6_K", expected: "Yarn Mistral 7B 64K Q6 K"},
		{input: "yarn-mistral:7b-64k-q8_0", expected: "Yarn Mistral 7B 64K Q8_0"},
		{input: "everythinglm:latest", expected: "Everythinglm (latest)"},
		{input: "everythinglm", expected: "Everythinglm"},
		{input: "everythinglm:13b", expected: "Everythinglm 13B"},
		{input: "everythinglm:13b-16k", expected: "Everythinglm 13B 16K"},
		{input: "everythinglm:13b-16k-q2_K", expected: "Everythinglm 13B 16K Q2_K"},
		{input: "everythinglm:13b-16k-q3_K_S", expected: "Everythinglm 13B 16K Q3 K_S"},
		{input: "everythinglm:13b-16k-q3_K_M", expected: "Everythinglm 13B 16K Q3 K_M"},
		{input: "everythinglm:13b-16k-q3_K_L", expected: "Everythinglm 13B 16K Q3 K_L"},
		{input: "everythinglm:13b-16k-q4_0", expected: "Everythinglm 13B 16K Q4_0"},
		{input: "everythinglm:13b-16k-q4_1", expected: "Everythinglm 13B 16K Q4 1"},
		{input: "everythinglm:13b-16k-q4_K_S", expected: "Everythinglm 13B 16K Q4 K_S"},
		{input: "everythinglm:13b-16k-q4_K_M", expected: "Everythinglm 13B 16K Q4_K_M"},
		{input: "everythinglm:13b-16k-q5_0", expected: "Everythinglm 13B 16K Q5 0"},
		{input: "everythinglm:13b-16k-q5_1", expected: "Everythinglm 13B 16K Q5 1"},
		{input: "everythinglm:13b-16k-q5_K_S", expected: "Everythinglm 13B 16K Q5 K_S"},
		{input: "everythinglm:13b-16k-q5_K_M", expected: "Everythinglm 13B 16K Q5 K_M"},
		{input: "everythinglm:13b-16k-q6_K", expected: "Everythinglm 13B 16K Q6 K"},
		{input: "everythinglm:13b-16k-q8_0", expected: "Everythinglm 13B 16K Q8_0"},
		{input: "everythinglm:13b-16k-fp16", expected: "Everythinglm 13B 16K FP16"},
		{input: "nexusraven:latest", expected: "Nexusraven (latest)"},
		{input: "nexusraven", expected: "Nexusraven"},
		{input: "nexusraven:13b", expected: "Nexusraven 13B"},
		{input: "nexusraven:13b-v2-q2_K", expected: "Nexusraven 13B v2 Q2_K"},
		{input: "nexusraven:13b-v2-q3_K_S", expected: "Nexusraven 13B v2 Q3 K_S"},
		{input: "nexusraven:13b-v2-q3_K_M", expected: "Nexusraven 13B v2 Q3 K_M"},
		{input: "nexusraven:13b-v2-q3_K_L", expected: "Nexusraven 13B v2 Q3 K_L"},
		{input: "nexusraven:13b-v2-q4_0", expected: "Nexusraven 13B v2 Q4_0"},
		{input: "nexusraven:13b-v2-q4_1", expected: "Nexusraven 13B v2 Q4 1"},
		{input: "nexusraven:13b-v2-q4_K_S", expected: "Nexusraven 13B v2 Q4 K_S"},
		{input: "nexusraven:13b-v2-q4_K_M", expected: "Nexusraven 13B v2 Q4_K_M"},
		{input: "nexusraven:13b-v2-q5_0", expected: "Nexusraven 13B v2 Q5 0"},
		{input: "nexusraven:13b-v2-q5_1", expected: "Nexusraven 13B v2 Q5 1"},
		{input: "nexusraven:13b-v2-q5_K_S", expected: "Nexusraven 13B v2 Q5 K_S"},
		{input: "nexusraven:13b-v2-q5_K_M", expected: "Nexusraven 13B v2 Q5 K_M"},
		{input: "nexusraven:13b-v2-q6_K", expected: "Nexusraven 13B v2 Q6 K"},
		{input: "nexusraven:13b-v2-q8_0", expected: "Nexusraven 13B v2 Q8_0"},
		{input: "nexusraven:13b-v2-fp16", expected: "Nexusraven 13B v2 FP16"},
		{input: "nexusraven:13b-q2_K", expected: "Nexusraven 13B Q2_K"},
		{input: "nexusraven:13b-q3_K_S", expected: "Nexusraven 13B Q3 K_S"},
		{input: "nexusraven:13b-q3_K_M", expected: "Nexusraven 13B Q3 K_M"},
		{input: "nexusraven:13b-q3_K_L", expected: "Nexusraven 13B Q3 K_L"},
		{input: "nexusraven:13b-q4_0", expected: "Nexusraven 13B Q4_0"},
		{input: "nexusraven:13b-q4_1", expected: "Nexusraven 13B Q4 1"},
		{input: "nexusraven:13b-q4_K_S", expected: "Nexusraven 13B Q4 K_S"},
		{input: "nexusraven:13b-q4_K_M", expected: "Nexusraven 13B Q4_K_M"},
		{input: "nexusraven:13b-q5_0", expected: "Nexusraven 13B Q5 0"},
		{input: "nexusraven:13b-q5_1", expected: "Nexusraven 13B Q5 1"},
		{input: "nexusraven:13b-q5_K_S", expected: "Nexusraven 13B Q5 K_S"},
		{input: "nexusraven:13b-q5_K_M", expected: "Nexusraven 13B Q5 K_M"},
		{input: "nexusraven:13b-q6_K", expected: "Nexusraven 13B Q6 K"},
		{input: "nexusraven:13b-q8_0", expected: "Nexusraven 13B Q8_0"},
		{input: "nexusraven:13b-fp16", expected: "Nexusraven 13B FP16"},
		{input: "codeup:latest", expected: "Codeup (latest)"},
		{input: "codeup", expected: "Codeup"},
		{input: "codeup:13b", expected: "Codeup 13B"},
		{input: "codeup:13b-llama2", expected: "Codeup 13B Llama2"},
		{input: "codeup:13b-llama2-chat", expected: "Codeup 13B Llama2 Chat"},
		{input: "codeup:13b-llama2-chat-q2_K", expected: "Codeup 13B Llama2 Chat Q2_K"},
		{input: "codeup:13b-llama2-chat-q3_K_S", expected: "Codeup 13B Llama2 Chat Q3 K_S"},
		{input: "codeup:13b-llama2-chat-q3_K_M", expected: "Codeup 13B Llama2 Chat Q3 K_M"},
		{input: "codeup:13b-llama2-chat-q3_K_L", expected: "Codeup 13B Llama2 Chat Q3 K_L"},
		{input: "codeup:13b-llama2-chat-q4_0", expected: "Codeup 13B Llama2 Chat Q4_0"},
		{input: "codeup:13b-llama2-chat-q4_1", expected: "Codeup 13B Llama2 Chat Q4 1"},
		{input: "codeup:13b-llama2-chat-q4_K_S", expected: "Codeup 13B Llama2 Chat Q4 K_S"},
		{input: "codeup:13b-llama2-chat-q4_K_M", expected: "Codeup 13B Llama2 Chat Q4_K_M"},
		{input: "codeup:13b-llama2-chat-q5_0", expected: "Codeup 13B Llama2 Chat Q5 0"},
		{input: "codeup:13b-llama2-chat-q5_1", expected: "Codeup 13B Llama2 Chat Q5 1"},
		{input: "codeup:13b-llama2-chat-q5_K_S", expected: "Codeup 13B Llama2 Chat Q5 K_S"},
		{input: "codeup:13b-llama2-chat-q5_K_M", expected: "Codeup 13B Llama2 Chat Q5 K_M"},
		{input: "codeup:13b-llama2-chat-q6_K", expected: "Codeup 13B Llama2 Chat Q6 K"},
		{input: "codeup:13b-llama2-chat-q8_0", expected: "Codeup 13B Llama2 Chat Q8_0"},
		{input: "codeup:13b-llama2-chat-fp16", expected: "Codeup 13B Llama2 Chat FP16"},
		{input: "marco-o1:latest", expected: "Marco o1 (latest)"},
		{input: "marco-o1", expected: "Marco o1"},
		{input: "marco-o1:7b", expected: "Marco o1 7B"},
		{input: "marco-o1:7b-q4_K_M", expected: "Marco o1 7B Q4_K_M"},
		{input: "marco-o1:7b-q8_0", expected: "Marco o1 7B Q8_0"},
		{input: "marco-o1:7b-fp16", expected: "Marco o1 7B FP16"},
		{input: "stablelm-zephyr:latest", expected: "Stablelm Zephyr (latest)"},
		{input: "stablelm-zephyr", expected: "Stablelm Zephyr"},
		{input: "stablelm-zephyr:3b", expected: "Stablelm Zephyr 3B"},
		{input: "stablelm-zephyr:3b-q2_K", expected: "Stablelm Zephyr 3B Q2_K"},
		{input: "stablelm-zephyr:3b-q3_K_S", expected: "Stablelm Zephyr 3B Q3 K_S"},
		{input: "stablelm-zephyr:3b-q3_K_M", expected: "Stablelm Zephyr 3B Q3 K_M"},
		{input: "stablelm-zephyr:3b-q3_K_L", expected: "Stablelm Zephyr 3B Q3 K_L"},
		{input: "stablelm-zephyr:3b-q4_0", expected: "Stablelm Zephyr 3B Q4_0"},
		{input: "stablelm-zephyr:3b-q4_1", expected: "Stablelm Zephyr 3B Q4 1"},
		{input: "stablelm-zephyr:3b-q4_K_S", expected: "Stablelm Zephyr 3B Q4 K_S"},
		{input: "stablelm-zephyr:3b-q4_K_M", expected: "Stablelm Zephyr 3B Q4_K_M"},
		{input: "stablelm-zephyr:3b-q5_0", expected: "Stablelm Zephyr 3B Q5 0"},
		{input: "stablelm-zephyr:3b-q5_1", expected: "Stablelm Zephyr 3B Q5 1"},
		{input: "stablelm-zephyr:3b-q5_K_S", expected: "Stablelm Zephyr 3B Q5 K_S"},
		{input: "stablelm-zephyr:3b-q5_K_M", expected: "Stablelm Zephyr 3B Q5 K_M"},
		{input: "stablelm-zephyr:3b-q6_K", expected: "Stablelm Zephyr 3B Q6 K"},
		{input: "stablelm-zephyr:3b-q8_0", expected: "Stablelm Zephyr 3B Q8_0"},
		{input: "stablelm-zephyr:3b-fp16", expected: "Stablelm Zephyr 3B FP16"},
		{input: "falcon2:latest", expected: "Falcon2 (latest)"},
		{input: "falcon2", expected: "Falcon2"},
		{input: "falcon2:11b", expected: "Falcon2 11B"},
		{input: "falcon2:11b-q2_K", expected: "Falcon2 11B Q2_K"},
		{input: "falcon2:11b-q3_K_S", expected: "Falcon2 11B Q3 K_S"},
		{input: "falcon2:11b-q3_K_M", expected: "Falcon2 11B Q3 K_M"},
		{input: "falcon2:11b-q3_K_L", expected: "Falcon2 11B Q3 K_L"},
		{input: "falcon2:11b-q4_0", expected: "Falcon2 11B Q4_0"},
		{input: "falcon2:11b-q4_1", expected: "Falcon2 11B Q4 1"},
		{input: "falcon2:11b-q4_K_S", expected: "Falcon2 11B Q4 K_S"},
		{input: "falcon2:11b-q4_K_M", expected: "Falcon2 11B Q4_K_M"},
		{input: "falcon2:11b-q5_0", expected: "Falcon2 11B Q5 0"},
		{input: "falcon2:11b-q5_1", expected: "Falcon2 11B Q5 1"},
		{input: "falcon2:11b-q5_K_S", expected: "Falcon2 11B Q5 K_S"},
		{input: "falcon2:11b-q5_K_M", expected: "Falcon2 11B Q5 K_M"},
		{input: "falcon2:11b-q6_K", expected: "Falcon2 11B Q6 K"},
		{input: "falcon2:11b-q8_0", expected: "Falcon2 11B Q8_0"},
		{input: "falcon2:11b-fp16", expected: "Falcon2 11B FP16"},
		{input: "solar-pro:latest", expected: "Solar Pro (latest)"},
		{input: "solar-pro", expected: "Solar Pro"},
		{input: "solar-pro:preview", expected: "Solar Pro Preview"},
		{input: "solar-pro:22b", expected: "Solar Pro 22B"},
		{input: "solar-pro:22b-preview-instruct-q2_K", expected: "Solar Pro 22B Preview Instruct Q2_K"},
		{input: "solar-pro:22b-preview-instruct-q3_K_S", expected: "Solar Pro 22B Preview Instruct Q3 K_S"},
		{input: "solar-pro:22b-preview-instruct-q3_K_M", expected: "Solar Pro 22B Preview Instruct Q3 K_M"},
		{input: "solar-pro:22b-preview-instruct-q3_K_L", expected: "Solar Pro 22B Preview Instruct Q3 K_L"},
		{input: "solar-pro:22b-preview-instruct-q4_0", expected: "Solar Pro 22B Preview Instruct Q4_0"},
		{input: "solar-pro:22b-preview-instruct-q4_1", expected: "Solar Pro 22B Preview Instruct Q4 1"},
		{input: "solar-pro:22b-preview-instruct-q4_K_S", expected: "Solar Pro 22B Preview Instruct Q4 K_S"},
		{input: "solar-pro:22b-preview-instruct-q4_K_M", expected: "Solar Pro 22B Preview Instruct Q4_K_M"},
		{input: "solar-pro:22b-preview-instruct-q5_0", expected: "Solar Pro 22B Preview Instruct Q5 0"},
		{input: "solar-pro:22b-preview-instruct-q5_1", expected: "Solar Pro 22B Preview Instruct Q5 1"},
		{input: "solar-pro:22b-preview-instruct-q5_K_S", expected: "Solar Pro 22B Preview Instruct Q5 K_S"},
		{input: "solar-pro:22b-preview-instruct-q5_K_M", expected: "Solar Pro 22B Preview Instruct Q5 K_M"},
		{input: "solar-pro:22b-preview-instruct-q6_K", expected: "Solar Pro 22B Preview Instruct Q6 K"},
		{input: "solar-pro:22b-preview-instruct-q8_0", expected: "Solar Pro 22B Preview Instruct Q8_0"},
		{input: "solar-pro:22b-preview-instruct-fp16", expected: "Solar Pro 22B Preview Instruct FP16"},
		{input: "duckdb-nsql:latest", expected: "Duckdb Nsql (latest)"},
		{input: "duckdb-nsql", expected: "Duckdb Nsql"},
		{input: "duckdb-nsql:7b", expected: "Duckdb Nsql 7B"},
		{input: "duckdb-nsql:7b-q2_K", expected: "Duckdb Nsql 7B Q2_K"},
		{input: "duckdb-nsql:7b-q3_K_S", expected: "Duckdb Nsql 7B Q3 K_S"},
		{input: "duckdb-nsql:7b-q3_K_M", expected: "Duckdb Nsql 7B Q3 K_M"},
		{input: "duckdb-nsql:7b-q3_K_L", expected: "Duckdb Nsql 7B Q3 K_L"},
		{input: "duckdb-nsql:7b-q4_0", expected: "Duckdb Nsql 7B Q4_0"},
		{input: "duckdb-nsql:7b-q4_1", expected: "Duckdb Nsql 7B Q4 1"},
		{input: "duckdb-nsql:7b-q4_K_S", expected: "Duckdb Nsql 7B Q4 K_S"},
		{input: "duckdb-nsql:7b-q4_K_M", expected: "Duckdb Nsql 7B Q4_K_M"},
		{input: "duckdb-nsql:7b-q5_0", expected: "Duckdb Nsql 7B Q5 0"},
		{input: "duckdb-nsql:7b-q5_1", expected: "Duckdb Nsql 7B Q5 1"},
		{input: "duckdb-nsql:7b-q5_K_S", expected: "Duckdb Nsql 7B Q5 K_S"},
		{input: "duckdb-nsql:7b-q5_K_M", expected: "Duckdb Nsql 7B Q5 K_M"},
		{input: "duckdb-nsql:7b-q6_K", expected: "Duckdb Nsql 7B Q6 K"},
		{input: "duckdb-nsql:7b-q8_0", expected: "Duckdb Nsql 7B Q8_0"},
		{input: "duckdb-nsql:7b-fp16", expected: "Duckdb Nsql 7B FP16"},
		{input: "mistrallite:latest", expected: "Mistrallite (latest)"},
		{input: "mistrallite", expected: "Mistrallite"},
		{input: "mistrallite:7b", expected: "Mistrallite 7B"},
		{input: "mistrallite:7b-v0.1-q2_K", expected: "Mistrallite 7B v0.1 Q2_K"},
		{input: "mistrallite:7b-v0.1-q3_K_S", expected: "Mistrallite 7B v0.1 Q3 K_S"},
		{input: "mistrallite:7b-v0.1-q3_K_M", expected: "Mistrallite 7B v0.1 Q3 K_M"},
		{input: "mistrallite:7b-v0.1-q3_K_L", expected: "Mistrallite 7B v0.1 Q3 K_L"},
		{input: "mistrallite:7b-v0.1-q4_0", expected: "Mistrallite 7B v0.1 Q4_0"},
		{input: "mistrallite:7b-v0.1-q4_1", expected: "Mistrallite 7B v0.1 Q4 1"},
		{input: "mistrallite:7b-v0.1-q4_K_S", expected: "Mistrallite 7B v0.1 Q4 K_S"},
		{input: "mistrallite:7b-v0.1-q4_K_M", expected: "Mistrallite 7B v0.1 Q4_K_M"},
		{input: "mistrallite:7b-v0.1-q5_0", expected: "Mistrallite 7B v0.1 Q5 0"},
		{input: "mistrallite:7b-v0.1-q5_1", expected: "Mistrallite 7B v0.1 Q5 1"},
		{input: "mistrallite:7b-v0.1-q5_K_S", expected: "Mistrallite 7B v0.1 Q5 K_S"},
		{input: "mistrallite:7b-v0.1-q5_K_M", expected: "Mistrallite 7B v0.1 Q5 K_M"},
		{input: "mistrallite:7b-v0.1-q6_K", expected: "Mistrallite 7B v0.1 Q6 K"},
		{input: "mistrallite:7b-v0.1-q8_0", expected: "Mistrallite 7B v0.1 Q8_0"},
		{input: "mistrallite:7b-v0.1-fp16", expected: "Mistrallite 7B v0.1 FP16"},
		{input: "magicoder:latest", expected: "Magicoder (latest)"},
		{input: "magicoder", expected: "Magicoder"},
		{input: "magicoder:7b", expected: "Magicoder 7B"},
		{input: "magicoder:7b-s-cl", expected: "Magicoder 7B S Cl"},
		{input: "magicoder:7b-s-cl-q2_K", expected: "Magicoder 7B S Cl Q2_K"},
		{input: "magicoder:7b-s-cl-q3_K_S", expected: "Magicoder 7B S Cl Q3 K_S"},
		{input: "magicoder:7b-s-cl-q3_K_M", expected: "Magicoder 7B S Cl Q3 K_M"},
		{input: "magicoder:7b-s-cl-q3_K_L", expected: "Magicoder 7B S Cl Q3 K_L"},
		{input: "magicoder:7b-s-cl-q4_0", expected: "Magicoder 7B S Cl Q4_0"},
		{input: "magicoder:7b-s-cl-q4_1", expected: "Magicoder 7B S Cl Q4 1"},
		{input: "magicoder:7b-s-cl-q4_K_S", expected: "Magicoder 7B S Cl Q4 K_S"},
		{input: "magicoder:7b-s-cl-q4_K_M", expected: "Magicoder 7B S Cl Q4_K_M"},
		{input: "magicoder:7b-s-cl-q5_0", expected: "Magicoder 7B S Cl Q5 0"},
		{input: "magicoder:7b-s-cl-q5_1", expected: "Magicoder 7B S Cl Q5 1"},
		{input: "magicoder:7b-s-cl-q5_K_S", expected: "Magicoder 7B S Cl Q5 K_S"},
		{input: "magicoder:7b-s-cl-q5_K_M", expected: "Magicoder 7B S Cl Q5 K_M"},
		{input: "magicoder:7b-s-cl-q6_K", expected: "Magicoder 7B S Cl Q6 K"},
		{input: "magicoder:7b-s-cl-q8_0", expected: "Magicoder 7B S Cl Q8_0"},
		{input: "magicoder:7b-s-cl-fp16", expected: "Magicoder 7B S Cl FP16"},
		{input: "codebooga:latest", expected: "Codebooga (latest)"},
		{input: "codebooga", expected: "Codebooga"},
		{input: "codebooga:34b", expected: "Codebooga 34B"},
		{input: "codebooga:34b-v0.1-q2_K", expected: "Codebooga 34B v0.1 Q2_K"},
		{input: "codebooga:34b-v0.1-q3_K_S", expected: "Codebooga 34B v0.1 Q3 K_S"},
		{input: "codebooga:34b-v0.1-q3_K_M", expected: "Codebooga 34B v0.1 Q3 K_M"},
		{input: "codebooga:34b-v0.1-q3_K_L", expected: "Codebooga 34B v0.1 Q3 K_L"},
		{input: "codebooga:34b-v0.1-q4_0", expected: "Codebooga 34B v0.1 Q4_0"},
		{input: "codebooga:34b-v0.1-q4_1", expected: "Codebooga 34B v0.1 Q4 1"},
		{input: "codebooga:34b-v0.1-q4_K_M", expected: "Codebooga 34B v0.1 Q4_K_M"},
		{input: "codebooga:34b-v0.1-q5_0", expected: "Codebooga 34B v0.1 Q5 0"},
		{input: "codebooga:34b-v0.1-q5_1", expected: "Codebooga 34B v0.1 Q5 1"},
		{input: "codebooga:34b-v0.1-q5_K_S", expected: "Codebooga 34B v0.1 Q5 K_S"},
		{input: "codebooga:34b-v0.1-q5_K_M", expected: "Codebooga 34B v0.1 Q5 K_M"},
		{input: "codebooga:34b-v0.1-q6_K", expected: "Codebooga 34B v0.1 Q6 K"},
		{input: "codebooga:34b-v0.1-q8_0", expected: "Codebooga 34B v0.1 Q8_0"},
		{input: "codebooga:34b-v0.1-fp16", expected: "Codebooga 34B v0.1 FP16"},
		{input: "bespoke-minicheck:latest", expected: "Bespoke Minicheck (latest)"},
		{input: "bespoke-minicheck", expected: "Bespoke Minicheck"},
		{input: "bespoke-minicheck:7b", expected: "Bespoke Minicheck 7B"},
		{input: "bespoke-minicheck:7b-q2_K", expected: "Bespoke Minicheck 7B Q2_K"},
		{input: "bespoke-minicheck:7b-q3_K_S", expected: "Bespoke Minicheck 7B Q3 K_S"},
		{input: "bespoke-minicheck:7b-q3_K_M", expected: "Bespoke Minicheck 7B Q3 K_M"},
		{input: "bespoke-minicheck:7b-q3_K_L", expected: "Bespoke Minicheck 7B Q3 K_L"},
		{input: "bespoke-minicheck:7b-q4_0", expected: "Bespoke Minicheck 7B Q4_0"},
		{input: "bespoke-minicheck:7b-q4_1", expected: "Bespoke Minicheck 7B Q4 1"},
		{input: "bespoke-minicheck:7b-q4_K_S", expected: "Bespoke Minicheck 7B Q4 K_S"},
		{input: "bespoke-minicheck:7b-q4_K_M", expected: "Bespoke Minicheck 7B Q4_K_M"},
		{input: "bespoke-minicheck:7b-q5_0", expected: "Bespoke Minicheck 7B Q5 0"},
		{input: "bespoke-minicheck:7b-q5_1", expected: "Bespoke Minicheck 7B Q5 1"},
		{input: "bespoke-minicheck:7b-q5_K_S", expected: "Bespoke Minicheck 7B Q5 K_S"},
		{input: "bespoke-minicheck:7b-q5_K_M", expected: "Bespoke Minicheck 7B Q5 K_M"},
		{input: "bespoke-minicheck:7b-q6_K", expected: "Bespoke Minicheck 7B Q6 K"},
		{input: "bespoke-minicheck:7b-q8_0", expected: "Bespoke Minicheck 7B Q8_0"},
		{input: "bespoke-minicheck:7b-fp16", expected: "Bespoke Minicheck 7B FP16"},
		{input: "deepseek-ocr:latest", expected: "Deepseek Ocr (latest)"},
		{input: "deepseek-ocr", expected: "Deepseek Ocr"},
		{input: "deepseek-ocr:3b", expected: "Deepseek Ocr 3B"},
		{input: "deepseek-ocr:3b-bf16", expected: "Deepseek Ocr 3B BF16"},
		{input: "nuextract:latest", expected: "Nuextract (latest)"},
		{input: "nuextract", expected: "Nuextract"},
		{input: "nuextract:3.8b", expected: "Nuextract 3.8B"},
		{input: "nuextract:3.8b-q2_K", expected: "Nuextract 3.8B Q2_K"},
		{input: "nuextract:3.8b-q3_K_S", expected: "Nuextract 3.8B Q3 K_S"},
		{input: "nuextract:3.8b-q3_K_M", expected: "Nuextract 3.8B Q3 K_M"},
		{input: "nuextract:3.8b-q3_K_L", expected: "Nuextract 3.8B Q3 K_L"},
		{input: "nuextract:3.8b-q4_0", expected: "Nuextract 3.8B Q4_0"},
		{input: "nuextract:3.8b-q4_1", expected: "Nuextract 3.8B Q4 1"},
		{input: "nuextract:3.8b-q4_K_S", expected: "Nuextract 3.8B Q4 K_S"},
		{input: "nuextract:3.8b-q4_K_M", expected: "Nuextract 3.8B Q4_K_M"},
		{input: "nuextract:3.8b-q5_0", expected: "Nuextract 3.8B Q5 0"},
		{input: "nuextract:3.8b-q5_1", expected: "Nuextract 3.8B Q5 1"},
		{input: "nuextract:3.8b-q5_K_S", expected: "Nuextract 3.8B Q5 K_S"},
		{input: "nuextract:3.8b-q5_K_M", expected: "Nuextract 3.8B Q5 K_M"},
		{input: "nuextract:3.8b-q6_K", expected: "Nuextract 3.8B Q6 K"},
		{input: "nuextract:3.8b-q8_0", expected: "Nuextract 3.8B Q8_0"},
		{input: "nuextract:3.8b-fp16", expected: "Nuextract 3.8B FP16"},
		{input: "wizard-vicuna:latest", expected: "Wizard Vicuna (latest)"},
		{input: "wizard-vicuna", expected: "Wizard Vicuna"},
		{input: "wizard-vicuna:13b", expected: "Wizard Vicuna 13B"},
		{input: "wizard-vicuna:13b-q2_K", expected: "Wizard Vicuna 13B Q2_K"},
		{input: "wizard-vicuna:13b-q3_K_S", expected: "Wizard Vicuna 13B Q3 K_S"},
		{input: "wizard-vicuna:13b-q3_K_M", expected: "Wizard Vicuna 13B Q3 K_M"},
		{input: "wizard-vicuna:13b-q3_K_L", expected: "Wizard Vicuna 13B Q3 K_L"},
		{input: "wizard-vicuna:13b-q4_0", expected: "Wizard Vicuna 13B Q4_0"},
		{input: "wizard-vicuna:13b-q4_1", expected: "Wizard Vicuna 13B Q4 1"},
		{input: "wizard-vicuna:13b-q4_K_S", expected: "Wizard Vicuna 13B Q4 K_S"},
		{input: "wizard-vicuna:13b-q4_K_M", expected: "Wizard Vicuna 13B Q4_K_M"},
		{input: "wizard-vicuna:13b-q5_0", expected: "Wizard Vicuna 13B Q5 0"},
		{input: "wizard-vicuna:13b-q5_1", expected: "Wizard Vicuna 13B Q5 1"},
		{input: "wizard-vicuna:13b-q5_K_S", expected: "Wizard Vicuna 13B Q5 K_S"},
		{input: "wizard-vicuna:13b-q5_K_M", expected: "Wizard Vicuna 13B Q5 K_M"},
		{input: "wizard-vicuna:13b-q6_K", expected: "Wizard Vicuna 13B Q6 K"},
		{input: "wizard-vicuna:13b-q8_0", expected: "Wizard Vicuna 13B Q8_0"},
		{input: "wizard-vicuna:13b-fp16", expected: "Wizard Vicuna 13B FP16"},
		{input: "granite3-guardian:latest", expected: "Granite3 Guardian (latest)"},
		{input: "granite3-guardian", expected: "Granite3 Guardian"},
		{input: "granite3-guardian:2b", expected: "Granite3 Guardian 2B"},
		{input: "granite3-guardian:8b", expected: "Granite3 Guardian 8B"},
		{input: "granite3-guardian:2b-q8_0", expected: "Granite3 Guardian 2B Q8_0"},
		{input: "granite3-guardian:2b-fp16", expected: "Granite3 Guardian 2B FP16"},
		{input: "granite3-guardian:8b-q5_K_S", expected: "Granite3 Guardian 8B Q5 K_S"},
		{input: "granite3-guardian:8b-q5_K_M", expected: "Granite3 Guardian 8B Q5 K_M"},
		{input: "granite3-guardian:8b-q6_K", expected: "Granite3 Guardian 8B Q6 K"},
		{input: "granite3-guardian:8b-q8_0", expected: "Granite3 Guardian 8B Q8_0"},
		{input: "granite3-guardian:8b-fp16", expected: "Granite3 Guardian 8B FP16"},
		{input: "firefunction-v2:latest", expected: "Firefunction v2 (latest)"},
		{input: "firefunction-v2", expected: "Firefunction v2"},
		{input: "firefunction-v2:70b", expected: "Firefunction v2 70B"},
		{input: "firefunction-v2:70b-q2_K", expected: "Firefunction v2 70B Q2_K"},
		{input: "firefunction-v2:70b-q3_K_S", expected: "Firefunction v2 70B Q3 K_S"},
		{input: "firefunction-v2:70b-q3_K_M", expected: "Firefunction v2 70B Q3 K_M"},
		{input: "firefunction-v2:70b-q3_K_L", expected: "Firefunction v2 70B Q3 K_L"},
		{input: "firefunction-v2:70b-q4_0", expected: "Firefunction v2 70B Q4_0"},
		{input: "firefunction-v2:70b-q4_1", expected: "Firefunction v2 70B Q4 1"},
		{input: "firefunction-v2:70b-q4_K_S", expected: "Firefunction v2 70B Q4 K_S"},
		{input: "firefunction-v2:70b-q4_K_M", expected: "Firefunction v2 70B Q4_K_M"},
		{input: "firefunction-v2:70b-q5_0", expected: "Firefunction v2 70B Q5 0"},
		{input: "firefunction-v2:70b-q5_1", expected: "Firefunction v2 70B Q5 1"},
		{input: "firefunction-v2:70b-q5_K_S", expected: "Firefunction v2 70B Q5 K_S"},
		{input: "firefunction-v2:70b-q5_K_M", expected: "Firefunction v2 70B Q5 K_M"},
		{input: "firefunction-v2:70b-q6_K", expected: "Firefunction v2 70B Q6 K"},
		{input: "firefunction-v2:70b-q8_0", expected: "Firefunction v2 70B Q8_0"},
		{input: "firefunction-v2:70b-fp16", expected: "Firefunction v2 70B FP16"},
		{input: "megadolphin:latest", expected: "Megadolphin (latest)"},
		{input: "megadolphin", expected: "Megadolphin"},
		{input: "megadolphin:v2.2", expected: "Megadolphin v2.2"},
		{input: "megadolphin:120b", expected: "Megadolphin 120B"},
		{input: "megadolphin:120b-v2.2", expected: "Megadolphin 120B v2.2"},
		{input: "megadolphin:120b-v2.2-q2_K", expected: "Megadolphin 120B v2.2 Q2_K"},
		{input: "megadolphin:120b-v2.2-q3_K_S", expected: "Megadolphin 120B v2.2 Q3 K_S"},
		{input: "megadolphin:120b-v2.2-q3_K_M", expected: "Megadolphin 120B v2.2 Q3 K_M"},
		{input: "megadolphin:120b-v2.2-q3_K_L", expected: "Megadolphin 120B v2.2 Q3 K_L"},
		{input: "megadolphin:120b-v2.2-q4_0", expected: "Megadolphin 120B v2.2 Q4_0"},
		{input: "megadolphin:120b-v2.2-q4_1", expected: "Megadolphin 120B v2.2 Q4 1"},
		{input: "megadolphin:120b-v2.2-q4_K_S", expected: "Megadolphin 120B v2.2 Q4 K_S"},
		{input: "megadolphin:120b-v2.2-q4_K_M", expected: "Megadolphin 120B v2.2 Q4_K_M"},
		{input: "megadolphin:120b-v2.2-q5_0", expected: "Megadolphin 120B v2.2 Q5 0"},
		{input: "megadolphin:120b-v2.2-q5_1", expected: "Megadolphin 120B v2.2 Q5 1"},
		{input: "megadolphin:120b-v2.2-q5_K_S", expected: "Megadolphin 120B v2.2 Q5 K_S"},
		{input: "megadolphin:120b-v2.2-q5_K_M", expected: "Megadolphin 120B v2.2 Q5 K_M"},
		{input: "megadolphin:120b-v2.2-q6_K", expected: "Megadolphin 120B v2.2 Q6 K"},
		{input: "megadolphin:120b-v2.2-q8_0", expected: "Megadolphin 120B v2.2 Q8_0"},
		{input: "megadolphin:120b-v2.2-fp16", expected: "Megadolphin 120B v2.2 FP16"},
		{input: "notux:latest", expected: "Notux (latest)"},
		{input: "notux", expected: "Notux"},
		{input: "notux:8x7b", expected: "Notux 8x7B"},
		{input: "notux:8x7b-v1", expected: "Notux 8x7B v1"},
		{input: "notux:8x7b-v1-q2_K", expected: "Notux 8x7B v1 Q2_K"},
		{input: "notux:8x7b-v1-q3_K_S", expected: "Notux 8x7B v1 Q3 K_S"},
		{input: "notux:8x7b-v1-q3_K_M", expected: "Notux 8x7B v1 Q3 K_M"},
		{input: "notux:8x7b-v1-q3_K_L", expected: "Notux 8x7B v1 Q3 K_L"},
		{input: "notux:8x7b-v1-q4_0", expected: "Notux 8x7B v1 Q4_0"},
		{input: "notux:8x7b-v1-q4_1", expected: "Notux 8x7B v1 Q4 1"},
		{input: "notux:8x7b-v1-q4_K_S", expected: "Notux 8x7B v1 Q4 K_S"},
		{input: "notux:8x7b-v1-q4_K_M", expected: "Notux 8x7B v1 Q4_K_M"},
		{input: "notux:8x7b-v1-q5_0", expected: "Notux 8x7B v1 Q5 0"},
		{input: "notux:8x7b-v1-q5_1", expected: "Notux 8x7B v1 Q5 1"},
		{input: "notux:8x7b-v1-q5_K_S", expected: "Notux 8x7B v1 Q5 K_S"},
		{input: "notux:8x7b-v1-q5_K_M", expected: "Notux 8x7B v1 Q5 K_M"},
		{input: "notux:8x7b-v1-q6_K", expected: "Notux 8x7B v1 Q6 K"},
		{input: "notux:8x7b-v1-q8_0", expected: "Notux 8x7B v1 Q8_0"},
		{input: "notux:8x7b-v1-fp16", expected: "Notux 8x7B v1 FP16"},
		{input: "open-orca-platypus2:latest", expected: "Open Orca Platypus2 (latest)"},
		{input: "open-orca-platypus2", expected: "Open Orca Platypus2"},
		{input: "open-orca-platypus2:13b", expected: "Open Orca Platypus2 13B"},
		{input: "open-orca-platypus2:13b-q2_K", expected: "Open Orca Platypus2 13B Q2_K"},
		{input: "open-orca-platypus2:13b-q3_K_S", expected: "Open Orca Platypus2 13B Q3 K_S"},
		{input: "open-orca-platypus2:13b-q3_K_M", expected: "Open Orca Platypus2 13B Q3 K_M"},
		{input: "open-orca-platypus2:13b-q3_K_L", expected: "Open Orca Platypus2 13B Q3 K_L"},
		{input: "open-orca-platypus2:13b-q4_0", expected: "Open Orca Platypus2 13B Q4_0"},
		{input: "open-orca-platypus2:13b-q4_1", expected: "Open Orca Platypus2 13B Q4 1"},
		{input: "open-orca-platypus2:13b-q4_K_S", expected: "Open Orca Platypus2 13B Q4 K_S"},
		{input: "open-orca-platypus2:13b-q4_K_M", expected: "Open Orca Platypus2 13B Q4_K_M"},
		{input: "open-orca-platypus2:13b-q5_0", expected: "Open Orca Platypus2 13B Q5 0"},
		{input: "open-orca-platypus2:13b-q5_1", expected: "Open Orca Platypus2 13B Q5 1"},
		{input: "open-orca-platypus2:13b-q5_K_S", expected: "Open Orca Platypus2 13B Q5 K_S"},
		{input: "open-orca-platypus2:13b-q5_K_M", expected: "Open Orca Platypus2 13B Q5 K_M"},
		{input: "open-orca-platypus2:13b-q6_K", expected: "Open Orca Platypus2 13B Q6 K"},
		{input: "open-orca-platypus2:13b-q8_0", expected: "Open Orca Platypus2 13B Q8_0"},
		{input: "open-orca-platypus2:13b-fp16", expected: "Open Orca Platypus2 13B FP16"},
		{input: "sailor2:latest", expected: "Sailor2 (latest)"},
		{input: "sailor2", expected: "Sailor2"},
		{input: "sailor2:1b", expected: "Sailor2 1B"},
		{input: "sailor2:8b", expected: "Sailor2 8B"},
		{input: "sailor2:20b", expected: "Sailor2 20B"},
		{input: "sailor2:1b-chat-q4_K_M", expected: "Sailor2 1B Chat Q4_K_M"},
		{input: "sailor2:1b-chat-q8_0", expected: "Sailor2 1B Chat Q8_0"},
		{input: "sailor2:1b-chat-fp16", expected: "Sailor2 1B Chat FP16"},
		{input: "sailor2:8b-chat-q4_K_M", expected: "Sailor2 8B Chat Q4_K_M"},
		{input: "sailor2:8b-chat-q8_0", expected: "Sailor2 8B Chat Q8_0"},
		{input: "sailor2:8b-chat-fp16", expected: "Sailor2 8B Chat FP16"},
		{input: "sailor2:20b-chat-q4_K_M", expected: "Sailor2 20B Chat Q4_K_M"},
		{input: "sailor2:20b-chat-q8_0", expected: "Sailor2 20B Chat Q8_0"},
		{input: "sailor2:20b-chat-fp16", expected: "Sailor2 20B Chat FP16"},
		{input: "notus:latest", expected: "Notus (latest)"},
		{input: "notus", expected: "Notus"},
		{input: "notus:7b", expected: "Notus 7B"},
		{input: "notus:7b-v1", expected: "Notus 7B v1"},
		{input: "notus:7b-v1-q2_K", expected: "Notus 7B v1 Q2_K"},
		{input: "notus:7b-v1-q3_K_S", expected: "Notus 7B v1 Q3 K_S"},
		{input: "notus:7b-v1-q3_K_M", expected: "Notus 7B v1 Q3 K_M"},
		{input: "notus:7b-v1-q3_K_L", expected: "Notus 7B v1 Q3 K_L"},
		{input: "notus:7b-v1-q4_0", expected: "Notus 7B v1 Q4_0"},
		{input: "notus:7b-v1-q4_1", expected: "Notus 7B v1 Q4 1"},
		{input: "notus:7b-v1-q4_K_S", expected: "Notus 7B v1 Q4 K_S"},
		{input: "notus:7b-v1-q4_K_M", expected: "Notus 7B v1 Q4_K_M"},
		{input: "notus:7b-v1-q5_0", expected: "Notus 7B v1 Q5 0"},
		{input: "notus:7b-v1-q5_1", expected: "Notus 7B v1 Q5 1"},
		{input: "notus:7b-v1-q5_K_S", expected: "Notus 7B v1 Q5 K_S"},
		{input: "notus:7b-v1-q5_K_M", expected: "Notus 7B v1 Q5 K_M"},
		{input: "notus:7b-v1-q6_K", expected: "Notus 7B v1 Q6 K"},
		{input: "notus:7b-v1-q8_0", expected: "Notus 7B v1 Q8_0"},
		{input: "notus:7b-v1-fp16", expected: "Notus 7B v1 FP16"},
		{input: "goliath:latest", expected: "Goliath (latest)"},
		{input: "goliath", expected: "Goliath"},
		{input: "goliath:120b-q2_K", expected: "Goliath 120B Q2_K"},
		{input: "goliath:120b-q3_K_S", expected: "Goliath 120B Q3 K_S"},
		{input: "goliath:120b-q3_K_M", expected: "Goliath 120B Q3 K_M"},
		{input: "goliath:120b-q3_K_L", expected: "Goliath 120B Q3 K_L"},
		{input: "goliath:120b-q4_0", expected: "Goliath 120B Q4_0"},
		{input: "goliath:120b-q4_1", expected: "Goliath 120B Q4 1"},
		{input: "goliath:120b-q4_K_S", expected: "Goliath 120B Q4 K_S"},
		{input: "goliath:120b-q4_K_M", expected: "Goliath 120B Q4_K_M"},
		{input: "goliath:120b-q5_0", expected: "Goliath 120B Q5 0"},
		{input: "goliath:120b-q5_1", expected: "Goliath 120B Q5 1"},
		{input: "goliath:120b-q5_K_S", expected: "Goliath 120B Q5 K_S"},
		{input: "goliath:120b-q5_K_M", expected: "Goliath 120B Q5 K_M"},
		{input: "goliath:120b-q6_K", expected: "Goliath 120B Q6 K"},
		{input: "goliath:120b-q8_0", expected: "Goliath 120B Q8_0"},
		{input: "goliath:120b-fp16", expected: "Goliath 120B FP16"},
		{input: "alfred:latest", expected: "Alfred (latest)"},
		{input: "alfred", expected: "Alfred"},
		{input: "alfred:40b", expected: "Alfred 40B"},
		{input: "alfred:40b-1023-q4_0", expected: "Alfred 40B 1023 Q4_0"},
		{input: "alfred:40b-1023-q4_1", expected: "Alfred 40B 1023 Q4 1"},
		{input: "alfred:40b-1023-q5_0", expected: "Alfred 40B 1023 Q5 0"},
		{input: "alfred:40b-1023-q5_1", expected: "Alfred 40B 1023 Q5 1"},
		{input: "alfred:40b-1023-q8_0", expected: "Alfred 40B 1023 Q8_0"},
		{input: "command-r7b-arabic:latest", expected: "Command R7B Arabic (latest)"},
		{input: "command-r7b-arabic", expected: "Command R7B Arabic"},
		{input: "command-r7b-arabic:7b", expected: "Command R7B Arabic 7B"},
		{input: "command-r7b-arabic:7b-02-2025-q4_K_M", expected: "Command R7B Arabic 7B (2025-02) Q4_K_M"},
		{input: "command-r7b-arabic:7b-02-2025-q8_0", expected: "Command R7B Arabic 7B (2025-02) Q8_0"},
		{input: "command-r7b-arabic:7b-02-2025-fp16", expected: "Command R7B Arabic 7B (2025-02) FP16"},
		{input: "gemini-3-pro-preview:latest", expected: "Gemini 3 Pro Preview (latest)"},
		{input: "gemini-3-pro-preview", expected: "Gemini 3 Pro Preview"},
		{input: "glm-4.6:cloud", expected: "GLM 4.6 Cloud"},
		{input: "gpt-oss-safeguard:latest", expected: "GPT OSS Safeguard (latest)"},
		{input: "gpt-oss-safeguard", expected: "GPT OSS Safeguard"},
		{input: "gpt-oss-safeguard:20b", expected: "GPT OSS Safeguard 20B"},
		{input: "gpt-oss-safeguard:120b", expected: "GPT OSS Safeguard 120B"},
		{input: "minimax-m2:cloud", expected: "MiniMax M2 Cloud"},
		{input: "kimi-k2:1t-cloud", expected: "Kimi K2 1T Cloud"},
		{input: "cogito-2.1:latest", expected: "Cogito 2.1 (latest)"},
		{input: "cogito-2.1", expected: "Cogito 2.1"},
		{input: "cogito-2.1:671b", expected: "Cogito 2.1 671B"},
		{input: "cogito-2.1:671b-cloud", expected: "Cogito 2.1 671B Cloud"},
		{input: "cogito-2.1:671b-q4_K_M", expected: "Cogito 2.1 671B Q4_K_M"},
		{input: "cogito-2.1:671b-q8_0", expected: "Cogito 2.1 671B Q8_0"},
		{input: "cogito-2.1:671b-fp16", expected: "Cogito 2.1 671B FP16"},
		{input: "kimi-k2-thinking:cloud", expected: "Kimi K2 Thinking Cloud"},
		{input: "rnj-1:latest", expected: "Rnj 1 (latest)"},
		{input: "rnj-1", expected: "Rnj 1"},
		{input: "rnj-1:8b", expected: "Rnj 1 8B"},
		{input: "rnj-1:8b-cloud", expected: "Rnj 1 8B Cloud"},
		{input: "rnj-1:8b-instruct-q4_K_M", expected: "Rnj 1 8B Instruct Q4_K_M"},
		{input: "rnj-1:8b-instruct-q8_0", expected: "Rnj 1 8B Instruct Q8_0"},
		{input: "rnj-1:8b-instruct-fp16", expected: "Rnj 1 8B Instruct FP16"},
		{input: "nomic-embed-text-v2-moe:latest", expected: "Nomic Embed Text v2 Moe (latest)"},
		{input: "nomic-embed-text-v2-moe", expected: "Nomic Embed Text v2 Moe"},
		{input: "olmo-3.1:latest", expected: "Olmo 3.1 (latest)"},
		{input: "olmo-3.1", expected: "Olmo 3.1"},
		{input: "olmo-3.1:32b", expected: "Olmo 3.1 32B"},
		{input: "olmo-3.1:32b-instruct", expected: "Olmo 3.1 32B Instruct"},
		{input: "olmo-3.1:32b-instruct-q4_K_M", expected: "Olmo 3.1 32B Instruct Q4_K_M"},
		{input: "olmo-3.1:32b-instruct-q8_0", expected: "Olmo 3.1 32B Instruct Q8_0"},
		{input: "olmo-3.1:32b-instruct-fp16", expected: "Olmo 3.1 32B Instruct FP16"},
		{input: "olmo-3.1:32b-think", expected: "Olmo 3.1 32B Think"},
		{input: "olmo-3.1:32b-think-q4_K_M", expected: "Olmo 3.1 32B Think Q4_K_M"},
		{input: "olmo-3.1:32b-think-q8_0", expected: "Olmo 3.1 32B Think Q8_0"},
		{input: "olmo-3.1:32b-think-fp16", expected: "Olmo 3.1 32B Think FP16"},
		{input: "deepseek-v3.2:cloud", expected: "Deepseek v3.2 Cloud"},
		{input: "mistral-large-3:675b-cloud", expected: "Mistral Large 3 675B Cloud"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, FormatModelTitle(tt.input))
		})
	}
}

func Test_ModelTitle_Print_Ollama_TestCases(t *testing.T) {
	// return

	modelsStr := `nemotron-3-nano:latest
nemotron-3-nano
nemotron-3-nano:30b
nemotron-3-nano:30b-a3b-q4_K_M
nemotron-3-nano:30b-a3b-q8_0
nemotron-3-nano:30b-a3b-fp16
nemotron-3-nano:30b-cloud
functiongemma:latest
functiongemma
functiongemma:270m
functiongemma:270m-it-q8_0
functiongemma:270m-it-fp16
olmo-3:latest
olmo-3
olmo-3:7b
olmo-3:32b
olmo-3:7b-instruct
olmo-3:7b-instruct-q4_K_M
olmo-3:7b-instruct-q8_0
olmo-3:7b-instruct-fp16
olmo-3:7b-think
olmo-3:7b-think-q4_K_M
olmo-3:7b-think-q8_0
olmo-3:7b-think-fp16
olmo-3:32b-think
olmo-3:32b-think-q4_K_M
olmo-3:32b-think-q8_0
olmo-3:32b-think-fp16
gemini-3-flash-preview:latest
gemini-3-flash-preview
gemini-3-flash-preview:cloud
devstral-small-2:latest
devstral-small-2
devstral-small-2:24b
devstral-small-2:24b-cloud
devstral-small-2:24b-instruct-2512-q4_K_M
devstral-small-2:24b-instruct-2512-q8_0
devstral-small-2:24b-instruct-2512-fp16
devstral-2:latest
devstral-2
devstral-2:123b
devstral-2:123b-cloud
devstral-2:123b-instruct-2512-q4_K_M
devstral-2:123b-instruct-2512-q8_0
devstral-2:123b-instruct-2512-fp16
ministral-3:latest
ministral-3
ministral-3:3b
ministral-3:8b
ministral-3:14b
ministral-3:3b-cloud
ministral-3:3b-instruct-2512-q4_K_M
ministral-3:3b-instruct-2512-q8_0
ministral-3:3b-instruct-2512-fp16
ministral-3:8b-cloud
ministral-3:8b-instruct-2512-q4_K_M
ministral-3:8b-instruct-2512-q8_0
ministral-3:8b-instruct-2512-fp16
ministral-3:14b-cloud
ministral-3:14b-instruct-2512-q4_K_M
ministral-3:14b-instruct-2512-q8_0
ministral-3:14b-instruct-2512-fp16
qwen3-vl:latest
qwen3-vl
qwen3-vl:2b
qwen3-vl:4b
qwen3-vl:8b
qwen3-vl:30b
qwen3-vl:32b
qwen3-vl:235b
qwen3-vl:2b-instruct
qwen3-vl:2b-instruct-q4_K_M
qwen3-vl:2b-instruct-q8_0
qwen3-vl:2b-instruct-bf16
qwen3-vl:2b-thinking
qwen3-vl:2b-thinking-q4_K_M
qwen3-vl:2b-thinking-q8_0
qwen3-vl:2b-thinking-bf16
qwen3-vl:4b-instruct
qwen3-vl:4b-instruct-q4_K_M
qwen3-vl:4b-instruct-q8_0
qwen3-vl:4b-instruct-bf16
qwen3-vl:4b-thinking
qwen3-vl:4b-thinking-q4_K_M
qwen3-vl:4b-thinking-q8_0
qwen3-vl:4b-thinking-bf16
qwen3-vl:8b-instruct
qwen3-vl:8b-instruct-q4_K_M
qwen3-vl:8b-instruct-q8_0
qwen3-vl:8b-instruct-bf16
qwen3-vl:8b-thinking
qwen3-vl:8b-thinking-q4_K_M
qwen3-vl:8b-thinking-q8_0
qwen3-vl:8b-thinking-bf16
qwen3-vl:30b-a3b
qwen3-vl:30b-a3b-instruct
qwen3-vl:30b-a3b-instruct-q4_K_M
qwen3-vl:30b-a3b-instruct-q8_0
qwen3-vl:30b-a3b-instruct-bf16
qwen3-vl:30b-a3b-thinking
qwen3-vl:30b-a3b-thinking-q4_K_M
qwen3-vl:30b-a3b-thinking-q8_0
qwen3-vl:30b-a3b-thinking-bf16
qwen3-vl:32b-instruct
qwen3-vl:32b-instruct-q4_K_M
qwen3-vl:32b-instruct-q8_0
qwen3-vl:32b-instruct-bf16
qwen3-vl:32b-thinking
qwen3-vl:32b-thinking-q4_K_M
qwen3-vl:32b-thinking-q8_0
qwen3-vl:32b-thinking-bf16
qwen3-vl:235b-a22b
qwen3-vl:235b-a22b-instruct
qwen3-vl:235b-a22b-instruct-q4_K_M
qwen3-vl:235b-a22b-instruct-q8_0
qwen3-vl:235b-a22b-instruct-bf16
qwen3-vl:235b-a22b-thinking
qwen3-vl:235b-a22b-thinking-q4_K_M
qwen3-vl:235b-a22b-thinking-q8_0
qwen3-vl:235b-a22b-thinking-bf16
qwen3-vl:235b-cloud
qwen3-vl:235b-instruct-cloud
gpt-oss:latest
gpt-oss
gpt-oss:20b
gpt-oss:120b
gpt-oss:20b-cloud
gpt-oss:120b-cloud
deepseek-r1:latest
deepseek-r1
deepseek-r1:1.5b
deepseek-r1:7b
deepseek-r1:8b
deepseek-r1:14b
deepseek-r1:32b
deepseek-r1:70b
deepseek-r1:671b
deepseek-r1:1.5b-qwen-distill-q4_K_M
deepseek-r1:1.5b-qwen-distill-q8_0
deepseek-r1:1.5b-qwen-distill-fp16
deepseek-r1:7b-qwen-distill-q4_K_M
deepseek-r1:7b-qwen-distill-q8_0
deepseek-r1:7b-qwen-distill-fp16
deepseek-r1:8b-0528-qwen3-q4_K_M
deepseek-r1:8b-0528-qwen3-q8_0
deepseek-r1:8b-0528-qwen3-fp16
deepseek-r1:8b-llama-distill-q4_K_M
deepseek-r1:8b-llama-distill-q8_0
deepseek-r1:8b-llama-distill-fp16
deepseek-r1:14b-qwen-distill-q4_K_M
deepseek-r1:14b-qwen-distill-q8_0
deepseek-r1:14b-qwen-distill-fp16
deepseek-r1:32b-qwen-distill-q4_K_M
deepseek-r1:32b-qwen-distill-q8_0
deepseek-r1:32b-qwen-distill-fp16
deepseek-r1:70b-llama-distill-q4_K_M
deepseek-r1:70b-llama-distill-q8_0
deepseek-r1:70b-llama-distill-fp16
deepseek-r1:671b-0528-q4_K_M
deepseek-r1:671b-0528-q8_0
deepseek-r1:671b-0528-fp16
deepseek-r1:671b-q4_K_M
deepseek-r1:671b-q8_0
deepseek-r1:671b-fp16
qwen3-coder:latest
qwen3-coder
qwen3-coder:30b
qwen3-coder:480b
qwen3-coder:30b-a3b-q4_K_M
qwen3-coder:30b-a3b-q8_0
qwen3-coder:30b-a3b-fp16
qwen3-coder:480b-a35b-q4_K_M
qwen3-coder:480b-a35b-q8_0
qwen3-coder:480b-a35b-fp16
qwen3-coder:480b-cloud
gemma3:latest
gemma3
gemma3:270m
gemma3:1b
gemma3:4b
gemma3:12b
gemma3:27b
gemma3:270m-it-qat
gemma3:270m-it-q8_0
gemma3:270m-it-fp16
gemma3:270m-it-bf16
gemma3:1b-it-qat
gemma3:1b-it-q4_K_M
gemma3:1b-it-q8_0
gemma3:1b-it-fp16
gemma3:4b-cloud
gemma3:4b-it-qat
gemma3:4b-it-q4_K_M
gemma3:4b-it-q8_0
gemma3:4b-it-fp16
gemma3:12b-cloud
gemma3:12b-it-qat
gemma3:12b-it-q4_K_M
gemma3:12b-it-q8_0
gemma3:12b-it-fp16
gemma3:27b-cloud
gemma3:27b-it-qat
gemma3:27b-it-q4_K_M
gemma3:27b-it-q8_0
gemma3:27b-it-fp16
llama3.1:latest
llama3.1
llama3.1:8b
llama3.1:70b
llama3.1:405b
llama3.1:8b-instruct-q2_K
llama3.1:8b-instruct-q3_K_S
llama3.1:8b-instruct-q3_K_M
llama3.1:8b-instruct-q3_K_L
llama3.1:8b-instruct-q4_0
llama3.1:8b-instruct-q4_1
llama3.1:8b-instruct-q4_K_S
llama3.1:8b-instruct-q4_K_M
llama3.1:8b-instruct-q5_0
llama3.1:8b-instruct-q5_1
llama3.1:8b-instruct-q5_K_S
llama3.1:8b-instruct-q5_K_M
llama3.1:8b-instruct-q6_K
llama3.1:8b-instruct-q8_0
llama3.1:8b-instruct-fp16
llama3.1:8b-text-q2_K
llama3.1:8b-text-q3_K_S
llama3.1:8b-text-q3_K_M
llama3.1:8b-text-q3_K_L
llama3.1:8b-text-q4_0
llama3.1:8b-text-q4_1
llama3.1:8b-text-q4_K_S
llama3.1:8b-text-q4_K_M
llama3.1:8b-text-q5_0
llama3.1:8b-text-q5_1
llama3.1:8b-text-q5_K_S
llama3.1:8b-text-q5_K_M
llama3.1:8b-text-q6_K
llama3.1:8b-text-q8_0
llama3.1:8b-text-fp16
llama3.1:70b-instruct-q2_K
llama3.1:70b-instruct-q3_K_S
llama3.1:70b-instruct-q3_K_M
llama3.1:70b-instruct-q3_K_L
llama3.1:70b-instruct-q4_0
llama3.1:70b-instruct-q4_K_S
llama3.1:70b-instruct-q4_K_M
llama3.1:70b-instruct-q5_0
llama3.1:70b-instruct-q5_1
llama3.1:70b-instruct-q5_K_S
llama3.1:70b-instruct-q5_K_M
llama3.1:70b-instruct-q6_K
llama3.1:70b-instruct-q8_0
llama3.1:70b-instruct-fp16
llama3.1:70b-text-q2_K
llama3.1:70b-text-q3_K_S
llama3.1:70b-text-q3_K_M
llama3.1:70b-text-q3_K_L
llama3.1:70b-text-q4_0
llama3.1:70b-text-q4_1
llama3.1:70b-text-q4_K_S
llama3.1:70b-text-q4_K_M
llama3.1:70b-text-q5_0
llama3.1:70b-text-q5_1
llama3.1:70b-text-q5_K_S
llama3.1:70b-text-q5_K_M
llama3.1:70b-text-q6_K
llama3.1:70b-text-q8_0
llama3.1:70b-text-fp16
llama3.1:405b-instruct-q2_K
llama3.1:405b-instruct-q3_K_S
llama3.1:405b-instruct-q3_K_M
llama3.1:405b-instruct-q3_K_L
llama3.1:405b-instruct-q4_0
llama3.1:405b-instruct-q4_1
llama3.1:405b-instruct-q4_K_S
llama3.1:405b-instruct-q4_K_M
llama3.1:405b-instruct-q5_0
llama3.1:405b-instruct-q5_1
llama3.1:405b-instruct-q5_K_S
llama3.1:405b-instruct-q5_K_M
llama3.1:405b-instruct-q6_K
llama3.1:405b-instruct-q8_0
llama3.1:405b-instruct-fp16
llama3.1:405b-text-q2_K
llama3.1:405b-text-q3_K_S
llama3.1:405b-text-q3_K_M
llama3.1:405b-text-q3_K_L
llama3.1:405b-text-q4_0
llama3.1:405b-text-q4_1
llama3.1:405b-text-q4_K_S
llama3.1:405b-text-q4_K_M
llama3.1:405b-text-q5_0
llama3.1:405b-text-q5_1
llama3.1:405b-text-q5_K_S
llama3.1:405b-text-q5_K_M
llama3.1:405b-text-q6_K
llama3.1:405b-text-q8_0
llama3.1:405b-text-fp16
llama3.2:latest
llama3.2
llama3.2:1b
llama3.2:3b
llama3.2:1b-instruct-q2_K
llama3.2:1b-instruct-q3_K_S
llama3.2:1b-instruct-q3_K_M
llama3.2:1b-instruct-q3_K_L
llama3.2:1b-instruct-q4_0
llama3.2:1b-instruct-q4_1
llama3.2:1b-instruct-q4_K_S
llama3.2:1b-instruct-q4_K_M
llama3.2:1b-instruct-q5_0
llama3.2:1b-instruct-q5_1
llama3.2:1b-instruct-q5_K_S
llama3.2:1b-instruct-q5_K_M
llama3.2:1b-instruct-q6_K
llama3.2:1b-instruct-q8_0
llama3.2:1b-instruct-fp16
llama3.2:1b-text-q2_K
llama3.2:1b-text-q3_K_S
llama3.2:1b-text-q3_K_M
llama3.2:1b-text-q3_K_L
llama3.2:1b-text-q4_0
llama3.2:1b-text-q4_1
llama3.2:1b-text-q4_K_S
llama3.2:1b-text-q4_K_M
llama3.2:1b-text-q5_0
llama3.2:1b-text-q5_1
llama3.2:1b-text-q5_K_S
llama3.2:1b-text-q5_K_M
llama3.2:1b-text-q6_K
llama3.2:1b-text-q8_0
llama3.2:1b-text-fp16
llama3.2:3b-instruct-q2_K
llama3.2:3b-instruct-q3_K_S
llama3.2:3b-instruct-q3_K_M
llama3.2:3b-instruct-q3_K_L
llama3.2:3b-instruct-q4_0
llama3.2:3b-instruct-q4_1
llama3.2:3b-instruct-q4_K_S
llama3.2:3b-instruct-q4_K_M
llama3.2:3b-instruct-q5_0
llama3.2:3b-instruct-q5_1
llama3.2:3b-instruct-q5_K_S
llama3.2:3b-instruct-q5_K_M
llama3.2:3b-instruct-q6_K
llama3.2:3b-instruct-q8_0
llama3.2:3b-instruct-fp16
llama3.2:3b-text-q2_K
llama3.2:3b-text-q3_K_S
llama3.2:3b-text-q3_K_M
llama3.2:3b-text-q3_K_L
llama3.2:3b-text-q4_0
llama3.2:3b-text-q4_1
llama3.2:3b-text-q4_K_S
llama3.2:3b-text-q4_K_M
llama3.2:3b-text-q5_0
llama3.2:3b-text-q5_1
llama3.2:3b-text-q5_K_S
llama3.2:3b-text-q5_K_M
llama3.2:3b-text-q6_K
llama3.2:3b-text-q8_0
llama3.2:3b-text-fp16
nomic-embed-text:latest
nomic-embed-text
nomic-embed-text:v1.5
nomic-embed-text:137m-v1.5-fp16
mistral:latest
mistral
mistral:instruct
mistral:text
mistral:v0.1
mistral:v0.2
mistral:v0.3
mistral:7b
mistral:7b-instruct
mistral:7b-instruct-q2_K
mistral:7b-instruct-v0.2-q2_K
mistral:7b-instruct-q3_K_S
mistral:7b-instruct-v0.2-q3_K_S
mistral:7b-instruct-q3_K_M
mistral:7b-instruct-v0.2-q3_K_M
mistral:7b-instruct-q3_K_L
mistral:7b-instruct-v0.2-q3_K_L
mistral:7b-instruct-q4_0
mistral:7b-instruct-v0.2-q4_0
mistral:7b-instruct-q4_1
mistral:7b-instruct-v0.2-q4_1
mistral:7b-instruct-q4_K_S
mistral:7b-instruct-v0.2-q4_K_S
mistral:7b-instruct-v0.2-q4_K_M
mistral:7b-instruct-v0.2-q5_0
mistral:7b-instruct-v0.2-q5_1
mistral:7b-instruct-v0.2-q5_K_S
mistral:7b-instruct-v0.2-q5_K_M
mistral:7b-instruct-v0.2-q6_K
mistral:7b-instruct-v0.2-q8_0
mistral:7b-instruct-v0.2-fp16
mistral:7b-instruct-v0.3-q2_K
mistral:7b-instruct-v0.3-q3_K_S
mistral:7b-instruct-v0.3-q3_K_M
mistral:7b-instruct-v0.3-q3_K_L
mistral:7b-instruct-v0.3-q4_0
mistral:7b-instruct-v0.3-q4_1
mistral:7b-instruct-v0.3-q4_K_S
mistral:7b-instruct-q4_K_M
mistral:7b-instruct-v0.3-q4_K_M
mistral:7b-instruct-q5_0
mistral:7b-instruct-v0.3-q5_0
mistral:7b-instruct-q5_1
mistral:7b-instruct-v0.3-q5_1
mistral:7b-instruct-q5_K_S
mistral:7b-instruct-v0.3-q5_K_S
mistral:7b-instruct-q5_K_M
mistral:7b-instruct-v0.3-q5_K_M
mistral:7b-instruct-q6_K
mistral:7b-instruct-v0.3-q6_K
mistral:7b-instruct-q8_0
mistral:7b-instruct-v0.3-q8_0
mistral:7b-instruct-fp16
mistral:7b-instruct-v0.3-fp16
mistral:7b-text
mistral:7b-text-q2_K
mistral:7b-text-v0.2-q2_K
mistral:7b-text-q3_K_S
mistral:7b-text-v0.2-q3_K_S
mistral:7b-text-q3_K_M
mistral:7b-text-v0.2-q3_K_M
mistral:7b-text-q3_K_L
mistral:7b-text-v0.2-q3_K_L
mistral:7b-text-q4_0
mistral:7b-text-v0.2-q4_0
mistral:7b-text-q4_1
mistral:7b-text-v0.2-q4_1
mistral:7b-text-q4_K_S
mistral:7b-text-v0.2-q4_K_S
mistral:7b-text-q4_K_M
mistral:7b-text-v0.2-q4_K_M
mistral:7b-text-q5_0
mistral:7b-text-v0.2-q5_0
mistral:7b-text-q5_1
mistral:7b-text-v0.2-q5_1
mistral:7b-text-q5_K_S
mistral:7b-text-v0.2-q5_K_S
mistral:7b-text-q5_K_M
mistral:7b-text-v0.2-q5_K_M
mistral:7b-text-q6_K
mistral:7b-text-v0.2-q6_K
mistral:7b-text-q8_0
mistral:7b-text-v0.2-q8_0
mistral:7b-text-fp16
mistral:7b-text-v0.2-fp16
qwen2.5:latest
qwen2.5
qwen2.5:0.5b
qwen2.5:1.5b
qwen2.5:3b
qwen2.5:7b
qwen2.5:14b
qwen2.5:32b
qwen2.5:72b
qwen2.5:0.5b-base
qwen2.5:0.5b-base-q2_K
qwen2.5:0.5b-base-q3_K_S
qwen2.5:0.5b-base-q3_K_M
qwen2.5:0.5b-base-q3_K_L
qwen2.5:0.5b-base-q4_0
qwen2.5:0.5b-base-q4_1
qwen2.5:0.5b-base-q4_K_S
qwen2.5:0.5b-base-q4_K_M
qwen2.5:0.5b-base-q5_0
qwen2.5:0.5b-base-q5_1
qwen2.5:0.5b-base-q5_K_S
qwen2.5:0.5b-base-q8_0
qwen2.5:0.5b-instruct
qwen2.5:0.5b-instruct-q2_K
qwen2.5:0.5b-instruct-q3_K_S
qwen2.5:0.5b-instruct-q3_K_M
qwen2.5:0.5b-instruct-q3_K_L
qwen2.5:0.5b-instruct-q4_0
qwen2.5:0.5b-instruct-q4_1
qwen2.5:0.5b-instruct-q4_K_S
qwen2.5:0.5b-instruct-q4_K_M
qwen2.5:0.5b-instruct-q5_0
qwen2.5:0.5b-instruct-q5_1
qwen2.5:0.5b-instruct-q5_K_S
qwen2.5:0.5b-instruct-q5_K_M
qwen2.5:0.5b-instruct-q6_K
qwen2.5:0.5b-instruct-q8_0
qwen2.5:0.5b-instruct-fp16
qwen2.5:1.5b-instruct
qwen2.5:1.5b-instruct-q2_K
qwen2.5:1.5b-instruct-q3_K_S
qwen2.5:1.5b-instruct-q3_K_M
qwen2.5:1.5b-instruct-q3_K_L
qwen2.5:1.5b-instruct-q4_0
qwen2.5:1.5b-instruct-q4_1
qwen2.5:1.5b-instruct-q4_K_S
qwen2.5:1.5b-instruct-q4_K_M
qwen2.5:1.5b-instruct-q5_0
qwen2.5:1.5b-instruct-q5_1
qwen2.5:1.5b-instruct-q5_K_S
qwen2.5:1.5b-instruct-q5_K_M
qwen2.5:1.5b-instruct-q6_K
qwen2.5:1.5b-instruct-q8_0
qwen2.5:1.5b-instruct-fp16
qwen2.5:3b-instruct
qwen2.5:3b-instruct-q2_K
qwen2.5:3b-instruct-q3_K_S
qwen2.5:3b-instruct-q3_K_M
qwen2.5:3b-instruct-q3_K_L
qwen2.5:3b-instruct-q4_0
qwen2.5:3b-instruct-q4_1
qwen2.5:3b-instruct-q4_K_S
qwen2.5:3b-instruct-q4_K_M
qwen2.5:3b-instruct-q5_0
qwen2.5:3b-instruct-q5_1
qwen2.5:3b-instruct-q5_K_S
qwen2.5:3b-instruct-q5_K_M
qwen2.5:3b-instruct-q6_K
qwen2.5:3b-instruct-q8_0
qwen2.5:3b-instruct-fp16
qwen2.5:7b-instruct
qwen2.5:7b-instruct-q2_K
qwen2.5:7b-instruct-q3_K_S
qwen2.5:7b-instruct-q3_K_M
qwen2.5:7b-instruct-q3_K_L
qwen2.5:7b-instruct-q4_0
qwen2.5:7b-instruct-q4_1
qwen2.5:7b-instruct-q4_K_S
qwen2.5:7b-instruct-q4_K_M
qwen2.5:7b-instruct-q5_0
qwen2.5:7b-instruct-q5_1
qwen2.5:7b-instruct-q5_K_S
qwen2.5:7b-instruct-q5_K_M
qwen2.5:7b-instruct-q6_K
qwen2.5:7b-instruct-q8_0
qwen2.5:7b-instruct-fp16
qwen2.5:14b-instruct
qwen2.5:14b-instruct-q2_K
qwen2.5:14b-instruct-q3_K_S
qwen2.5:14b-instruct-q3_K_M
qwen2.5:14b-instruct-q3_K_L
qwen2.5:14b-instruct-q4_0
qwen2.5:14b-instruct-q4_1
qwen2.5:14b-instruct-q4_K_S
qwen2.5:14b-instruct-q4_K_M
qwen2.5:14b-instruct-q5_0
qwen2.5:14b-instruct-q5_1
qwen2.5:14b-instruct-q5_K_S
qwen2.5:14b-instruct-q5_K_M
qwen2.5:14b-instruct-q6_K
qwen2.5:14b-instruct-q8_0
qwen2.5:14b-instruct-fp16
qwen2.5:32b-instruct
qwen2.5:32b-instruct-q2_K
qwen2.5:32b-instruct-q3_K_S
qwen2.5:32b-instruct-q3_K_M
qwen2.5:32b-instruct-q3_K_L
qwen2.5:32b-instruct-q4_0
qwen2.5:32b-instruct-q4_1
qwen2.5:32b-instruct-q4_K_S
qwen2.5:32b-instruct-q4_K_M
qwen2.5:32b-instruct-q5_0
qwen2.5:32b-instruct-q5_1
qwen2.5:32b-instruct-q5_K_S
qwen2.5:32b-instruct-q5_K_M
qwen2.5:32b-instruct-q6_K
qwen2.5:32b-instruct-q8_0
qwen2.5:32b-instruct-fp16
qwen2.5:72b-instruct
qwen2.5:72b-instruct-q2_K
qwen2.5:72b-instruct-q3_K_S
qwen2.5:72b-instruct-q3_K_M
qwen2.5:72b-instruct-q3_K_L
qwen2.5:72b-instruct-q4_0
qwen2.5:72b-instruct-q4_1
qwen2.5:72b-instruct-q4_K_S
qwen2.5:72b-instruct-q4_K_M
qwen2.5:72b-instruct-q5_0
qwen2.5:72b-instruct-q5_1
qwen2.5:72b-instruct-q5_K_S
qwen2.5:72b-instruct-q5_K_M
qwen2.5:72b-instruct-q6_K
qwen2.5:72b-instruct-q8_0
qwen2.5:72b-instruct-fp16
qwen3:latest
qwen3
qwen3:0.6b
qwen3:1.7b
qwen3:4b
qwen3:8b
qwen3:14b
qwen3:30b
qwen3:32b
qwen3:235b
qwen3:0.6b-q4_K_M
qwen3:0.6b-q8_0
qwen3:0.6b-fp16
qwen3:1.7b-q4_K_M
qwen3:1.7b-q8_0
qwen3:1.7b-fp16
qwen3:4b-instruct
qwen3:4b-instruct-2507-q4_K_M
qwen3:4b-instruct-2507-q8_0
qwen3:4b-instruct-2507-fp16
qwen3:4b-thinking
qwen3:4b-thinking-2507-q4_K_M
qwen3:4b-thinking-2507-q8_0
qwen3:4b-thinking-2507-fp16
qwen3:4b-q4_K_M
qwen3:4b-q8_0
qwen3:4b-fp16
qwen3:8b-q4_K_M
qwen3:8b-q8_0
qwen3:8b-fp16
qwen3:14b-q4_K_M
qwen3:14b-q8_0
qwen3:14b-fp16
qwen3:30b-a3b
qwen3:30b-a3b-instruct-2507-q4_K_M
qwen3:30b-a3b-q4_K_M
qwen3:30b-a3b-instruct-2507-q8_0
qwen3:30b-a3b-thinking-2507-q4_K_M
qwen3:30b-a3b-q8_0
qwen3:30b-a3b-thinking-2507-q8_0
qwen3:30b-a3b-fp16
qwen3:30b-a3b-instruct-2507-fp16
qwen3:30b-a3b-thinking-2507-fp16
qwen3:30b-instruct
qwen3:30b-thinking
qwen3:32b-q4_K_M
qwen3:32b-q8_0
qwen3:32b-fp16
qwen3:235b-a22b
qwen3:235b-a22b-instruct-2507-q4_K_M
qwen3:235b-a22b-q4_K_M
qwen3:235b-a22b-instruct-2507-q8_0
qwen3:235b-a22b-thinking-2507-q4_K_M
qwen3:235b-a22b-q8_0
qwen3:235b-a22b-thinking-2507-q8_0
qwen3:235b-a22b-fp16
qwen3:235b-a22b-thinking-2507-fp16
qwen3:235b-instruct
qwen3:235b-thinking
phi3:latest
phi3
phi3:instruct
phi3:medium
phi3:mini
phi3:3.8b
phi3:14b
phi3:3.8b-instruct
phi3:3.8b-mini-128k-instruct-q2_K
phi3:3.8b-mini-128k-instruct-q3_K_S
phi3:3.8b-mini-128k-instruct-q3_K_M
phi3:3.8b-mini-128k-instruct-q3_K_L
phi3:3.8b-mini-128k-instruct-q4_0
phi3:3.8b-mini-128k-instruct-q4_1
phi3:3.8b-mini-128k-instruct-q4_K_S
phi3:3.8b-mini-128k-instruct-q4_K_M
phi3:3.8b-mini-128k-instruct-q5_0
phi3:3.8b-mini-128k-instruct-q5_1
phi3:3.8b-mini-128k-instruct-q5_K_S
phi3:3.8b-mini-128k-instruct-q5_K_M
phi3:3.8b-mini-128k-instruct-q6_K
phi3:3.8b-mini-128k-instruct-q8_0
phi3:3.8b-mini-128k-instruct-fp16
phi3:3.8b-mini-4k-instruct-q2_K
phi3:3.8b-mini-4k-instruct-q3_K_S
phi3:3.8b-mini-4k-instruct-q3_K_M
phi3:3.8b-mini-4k-instruct-q3_K_L
phi3:3.8b-mini-4k-instruct-q4_0
phi3:3.8b-mini-4k-instruct-q4_1
phi3:3.8b-mini-4k-instruct-q4_K_S
phi3:3.8b-mini-4k-instruct-q4_K_M
phi3:3.8b-mini-4k-instruct-q5_0
phi3:3.8b-mini-4k-instruct-q5_1
phi3:3.8b-mini-4k-instruct-q5_K_S
phi3:3.8b-mini-4k-instruct-q5_K_M
phi3:3.8b-mini-4k-instruct-q6_K
phi3:3.8b-mini-4k-instruct-q8_0
phi3:3.8b-mini-4k-instruct-fp16
phi3:14b-instruct
phi3:14b-medium-128k-instruct-q2_K
phi3:14b-medium-128k-instruct-q3_K_S
phi3:14b-medium-128k-instruct-q3_K_M
phi3:14b-medium-128k-instruct-q3_K_L
phi3:14b-medium-128k-instruct-q4_0
phi3:14b-medium-128k-instruct-q4_1
phi3:14b-medium-128k-instruct-q4_K_S
phi3:14b-medium-128k-instruct-q4_K_M
phi3:14b-medium-128k-instruct-q5_0
phi3:14b-medium-128k-instruct-q5_1
phi3:14b-medium-128k-instruct-q5_K_S
phi3:14b-medium-128k-instruct-q5_K_M
phi3:14b-medium-128k-instruct-q6_K
phi3:14b-medium-128k-instruct-q8_0
phi3:14b-medium-128k-instruct-fp16
phi3:14b-medium-4k-instruct-q2_K
phi3:14b-medium-4k-instruct-q3_K_S
phi3:14b-medium-4k-instruct-q3_K_M
phi3:14b-medium-4k-instruct-q3_K_L
phi3:14b-medium-4k-instruct-q4_0
phi3:14b-medium-4k-instruct-q4_1
phi3:14b-medium-4k-instruct-q4_K_S
phi3:14b-medium-4k-instruct-q4_K_M
phi3:14b-medium-4k-instruct-q5_0
phi3:14b-medium-4k-instruct-q5_1
phi3:14b-medium-4k-instruct-q5_K_S
phi3:14b-medium-4k-instruct-q5_K_M
phi3:14b-medium-4k-instruct-q6_K
phi3:14b-medium-4k-instruct-q8_0
phi3:14b-medium-4k-instruct-fp16
phi3:medium-128k
phi3:medium-4k
phi3:mini-128k
phi3:mini-4k
llama3:latest
llama3
llama3:instruct
llama3:text
llama3:8b
llama3:70b
llama3:8b-instruct-q2_K
llama3:8b-instruct-q3_K_S
llama3:8b-instruct-q3_K_M
llama3:8b-instruct-q3_K_L
llama3:8b-instruct-q4_0
llama3:8b-instruct-q4_1
llama3:8b-instruct-q4_K_S
llama3:8b-instruct-q4_K_M
llama3:8b-instruct-q5_0
llama3:8b-instruct-q5_1
llama3:8b-instruct-q5_K_S
llama3:8b-instruct-q5_K_M
llama3:8b-instruct-q6_K
llama3:8b-instruct-q8_0
llama3:8b-instruct-fp16
llama3:8b-text
llama3:8b-text-q2_K
llama3:8b-text-q3_K_S
llama3:8b-text-q3_K_M
llama3:8b-text-q3_K_L
llama3:8b-text-q4_0
llama3:8b-text-q4_1
llama3:8b-text-q4_K_S
llama3:8b-text-q4_K_M
llama3:8b-text-q5_0
llama3:8b-text-q5_1
llama3:8b-text-q5_K_S
llama3:8b-text-q5_K_M
llama3:8b-text-q6_K
llama3:8b-text-q8_0
llama3:8b-text-fp16
llama3:70b-instruct
llama3:70b-instruct-q2_K
llama3:70b-instruct-q3_K_S
llama3:70b-instruct-q3_K_M
llama3:70b-instruct-q3_K_L
llama3:70b-instruct-q4_0
llama3:70b-instruct-q4_1
llama3:70b-instruct-q4_K_S
llama3:70b-instruct-q4_K_M
llama3:70b-instruct-q5_0
llama3:70b-instruct-q5_1
llama3:70b-instruct-q5_K_S
llama3:70b-instruct-q5_K_M
llama3:70b-instruct-q6_K
llama3:70b-instruct-q8_0
llama3:70b-instruct-fp16
llama3:70b-text
llama3:70b-text-q2_K
llama3:70b-text-q3_K_S
llama3:70b-text-q3_K_M
llama3:70b-text-q3_K_L
llama3:70b-text-q4_0
llama3:70b-text-q4_1
llama3:70b-text-q4_K_S
llama3:70b-text-q4_K_M
llama3:70b-text-q5_0
llama3:70b-text-q5_1
llama3:70b-text-q5_K_S
llama3:70b-text-q5_K_M
llama3:70b-text-q6_K
llama3:70b-text-q8_0
llama3:70b-text-fp16
llava:latest
llava
llava:v1.6
llava:7b
llava:13b
llava:34b
llava:7b-v1.5-q2_K
llava:7b-v1.5-q3_K_S
llava:7b-v1.5-q3_K_M
llava:7b-v1.5-q3_K_L
llava:7b-v1.5-q4_0
llava:7b-v1.5-q4_1
llava:7b-v1.5-q4_K_S
llava:7b-v1.5-q4_K_M
llava:7b-v1.5-q5_0
llava:7b-v1.5-q5_1
llava:7b-v1.5-q5_K_S
llava:7b-v1.5-q5_K_M
llava:7b-v1.5-q6_K
llava:7b-v1.5-q8_0
llava:7b-v1.5-fp16
llava:7b-v1.6
llava:7b-v1.6-mistral-q2_K
llava:7b-v1.6-mistral-q3_K_S
llava:7b-v1.6-mistral-q3_K_M
llava:7b-v1.6-mistral-q3_K_L
llava:7b-v1.6-mistral-q4_0
llava:7b-v1.6-mistral-q4_1
llava:7b-v1.6-mistral-q4_K_S
llava:7b-v1.6-mistral-q4_K_M
llava:7b-v1.6-mistral-q5_0
llava:7b-v1.6-mistral-q5_1
llava:7b-v1.6-mistral-q5_K_S
llava:7b-v1.6-mistral-q5_K_M
llava:7b-v1.6-mistral-q6_K
llava:7b-v1.6-mistral-q8_0
llava:7b-v1.6-mistral-fp16
llava:7b-v1.6-vicuna-q2_K
llava:7b-v1.6-vicuna-q3_K_S
llava:7b-v1.6-vicuna-q3_K_M
llava:7b-v1.6-vicuna-q3_K_L
llava:7b-v1.6-vicuna-q4_0
llava:7b-v1.6-vicuna-q4_1
llava:7b-v1.6-vicuna-q4_K_S
llava:7b-v1.6-vicuna-q4_K_M
llava:7b-v1.6-vicuna-q5_0
llava:7b-v1.6-vicuna-q5_1
llava:7b-v1.6-vicuna-q5_K_S
llava:7b-v1.6-vicuna-q5_K_M
llava:7b-v1.6-vicuna-q6_K
llava:7b-v1.6-vicuna-q8_0
llava:7b-v1.6-vicuna-fp16
llava:13b-v1.5-q2_K
llava:13b-v1.5-q3_K_S
llava:13b-v1.5-q3_K_M
llava:13b-v1.5-q3_K_L
llava:13b-v1.5-q4_0
llava:13b-v1.5-q4_1
llava:13b-v1.5-q4_K_S
llava:13b-v1.5-q4_K_M
llava:13b-v1.5-q5_0
llava:13b-v1.5-q5_1
llava:13b-v1.5-q5_K_S
llava:13b-v1.5-q5_K_M
llava:13b-v1.5-q6_K
llava:13b-v1.5-q8_0
llava:13b-v1.5-fp16
llava:13b-v1.6
llava:13b-v1.6-vicuna-q2_K
llava:13b-v1.6-vicuna-q3_K_S
llava:13b-v1.6-vicuna-q3_K_M
llava:13b-v1.6-vicuna-q3_K_L
llava:13b-v1.6-vicuna-q4_0
llava:13b-v1.6-vicuna-q4_1
llava:13b-v1.6-vicuna-q4_K_S
llava:13b-v1.6-vicuna-q4_K_M
llava:13b-v1.6-vicuna-q5_0
llava:13b-v1.6-vicuna-q5_1
llava:13b-v1.6-vicuna-q5_K_S
llava:13b-v1.6-vicuna-q5_K_M
llava:13b-v1.6-vicuna-q6_K
llava:13b-v1.6-vicuna-q8_0
llava:13b-v1.6-vicuna-fp16
llava:34b-v1.6
llava:34b-v1.6-q2_K
llava:34b-v1.6-q3_K_S
llava:34b-v1.6-q3_K_M
llava:34b-v1.6-q3_K_L
llava:34b-v1.6-q4_0
llava:34b-v1.6-q4_1
llava:34b-v1.6-q4_K_S
llava:34b-v1.6-q4_K_M
llava:34b-v1.6-q5_0
llava:34b-v1.6-q5_1
llava:34b-v1.6-q5_K_S
llava:34b-v1.6-q5_K_M
llava:34b-v1.6-q6_K
llava:34b-v1.6-q8_0
llava:34b-v1.6-fp16
gemma2:latest
gemma2
gemma2:2b
gemma2:9b
gemma2:27b
gemma2:2b-instruct-q2_K
gemma2:2b-instruct-q3_K_S
gemma2:2b-instruct-q3_K_M
gemma2:2b-instruct-q3_K_L
gemma2:2b-instruct-q4_0
gemma2:2b-instruct-q4_1
gemma2:2b-instruct-q4_K_S
gemma2:2b-instruct-q4_K_M
gemma2:2b-instruct-q5_0
gemma2:2b-instruct-q5_1
gemma2:2b-instruct-q5_K_S
gemma2:2b-instruct-q5_K_M
gemma2:2b-instruct-q6_K
gemma2:2b-instruct-q8_0
gemma2:2b-instruct-fp16
gemma2:2b-text-q2_K
gemma2:2b-text-q3_K_S
gemma2:2b-text-q3_K_M
gemma2:2b-text-q3_K_L
gemma2:2b-text-q4_0
gemma2:2b-text-q4_1
gemma2:2b-text-q4_K_S
gemma2:2b-text-q4_K_M
gemma2:2b-text-q5_0
gemma2:2b-text-q5_1
gemma2:2b-text-q5_K_S
gemma2:2b-text-q5_K_M
gemma2:2b-text-q6_K
gemma2:2b-text-q8_0
gemma2:2b-text-fp16
gemma2:9b-instruct-q2_K
gemma2:9b-instruct-q3_K_S
gemma2:9b-instruct-q3_K_M
gemma2:9b-instruct-q3_K_L
gemma2:9b-instruct-q4_0
gemma2:9b-instruct-q4_1
gemma2:9b-instruct-q4_K_S
gemma2:9b-instruct-q4_K_M
gemma2:9b-instruct-q5_0
gemma2:9b-instruct-q5_1
gemma2:9b-instruct-q5_K_S
gemma2:9b-instruct-q5_K_M
gemma2:9b-instruct-q6_K
gemma2:9b-instruct-q8_0
gemma2:9b-instruct-fp16
gemma2:9b-text-q2_K
gemma2:9b-text-q3_K_S
gemma2:9b-text-q3_K_M
gemma2:9b-text-q3_K_L
gemma2:9b-text-q4_0
gemma2:9b-text-q4_1
gemma2:9b-text-q4_K_S
gemma2:9b-text-q4_K_M
gemma2:9b-text-q5_0
gemma2:9b-text-q5_1
gemma2:9b-text-q5_K_S
gemma2:9b-text-q5_K_M
gemma2:9b-text-q6_K
gemma2:9b-text-q8_0
gemma2:9b-text-fp16
gemma2:27b-instruct-q2_K
gemma2:27b-instruct-q3_K_S
gemma2:27b-instruct-q3_K_M
gemma2:27b-instruct-q3_K_L
gemma2:27b-instruct-q4_0
gemma2:27b-instruct-q4_1
gemma2:27b-instruct-q4_K_S
gemma2:27b-instruct-q4_K_M
gemma2:27b-instruct-q5_0
gemma2:27b-instruct-q5_1
gemma2:27b-instruct-q5_K_S
gemma2:27b-instruct-q5_K_M
gemma2:27b-instruct-q6_K
gemma2:27b-instruct-q8_0
gemma2:27b-instruct-fp16
gemma2:27b-text-q2_K
gemma2:27b-text-q3_K_S
gemma2:27b-text-q3_K_M
gemma2:27b-text-q3_K_L
gemma2:27b-text-q4_0
gemma2:27b-text-q4_1
gemma2:27b-text-q4_K_S
gemma2:27b-text-q4_K_M
gemma2:27b-text-q5_0
gemma2:27b-text-q5_1
gemma2:27b-text-q5_K_S
gemma2:27b-text-q5_K_M
gemma2:27b-text-q6_K
gemma2:27b-text-q8_0
gemma2:27b-text-fp16
qwen2.5-coder:latest
qwen2.5-coder
qwen2.5-coder:0.5b
qwen2.5-coder:1.5b
qwen2.5-coder:3b
qwen2.5-coder:7b
qwen2.5-coder:14b
qwen2.5-coder:32b
qwen2.5-coder:0.5b-base
qwen2.5-coder:0.5b-base-q2_K
qwen2.5-coder:0.5b-base-q3_K_S
qwen2.5-coder:0.5b-base-q3_K_M
qwen2.5-coder:0.5b-base-q3_K_L
qwen2.5-coder:0.5b-base-q4_0
qwen2.5-coder:0.5b-base-q4_1
qwen2.5-coder:0.5b-base-q4_K_S
qwen2.5-coder:0.5b-base-q4_K_M
qwen2.5-coder:0.5b-base-q5_0
qwen2.5-coder:0.5b-base-q5_1
qwen2.5-coder:0.5b-base-q5_K_S
qwen2.5-coder:0.5b-base-q5_K_M
qwen2.5-coder:0.5b-base-q6_K
qwen2.5-coder:0.5b-base-q8_0
qwen2.5-coder:0.5b-base-fp16
qwen2.5-coder:0.5b-instruct
qwen2.5-coder:0.5b-instruct-q2_K
qwen2.5-coder:0.5b-instruct-q3_K_S
qwen2.5-coder:0.5b-instruct-q3_K_M
qwen2.5-coder:0.5b-instruct-q3_K_L
qwen2.5-coder:0.5b-instruct-q4_0
qwen2.5-coder:0.5b-instruct-q4_1
qwen2.5-coder:0.5b-instruct-q4_K_S
qwen2.5-coder:0.5b-instruct-q4_K_M
qwen2.5-coder:0.5b-instruct-q5_0
qwen2.5-coder:0.5b-instruct-q5_1
qwen2.5-coder:0.5b-instruct-q5_K_S
qwen2.5-coder:0.5b-instruct-q5_K_M
qwen2.5-coder:0.5b-instruct-q6_K
qwen2.5-coder:0.5b-instruct-q8_0
qwen2.5-coder:0.5b-instruct-fp16
qwen2.5-coder:1.5b-base
qwen2.5-coder:1.5b-base-q2_K
qwen2.5-coder:1.5b-base-q3_K_S
qwen2.5-coder:1.5b-base-q3_K_M
qwen2.5-coder:1.5b-base-q3_K_L
qwen2.5-coder:1.5b-base-q4_0
qwen2.5-coder:1.5b-base-q4_1
qwen2.5-coder:1.5b-base-q4_K_S
qwen2.5-coder:1.5b-base-q4_K_M
qwen2.5-coder:1.5b-base-q5_0
qwen2.5-coder:1.5b-base-q5_1
qwen2.5-coder:1.5b-base-q5_K_S
qwen2.5-coder:1.5b-base-q5_K_M
qwen2.5-coder:1.5b-base-q6_K
qwen2.5-coder:1.5b-base-q8_0
qwen2.5-coder:1.5b-base-fp16
qwen2.5-coder:1.5b-instruct
qwen2.5-coder:1.5b-instruct-q2_K
qwen2.5-coder:1.5b-instruct-q3_K_S
qwen2.5-coder:1.5b-instruct-q3_K_M
qwen2.5-coder:1.5b-instruct-q3_K_L
qwen2.5-coder:1.5b-instruct-q4_0
qwen2.5-coder:1.5b-instruct-q4_1
qwen2.5-coder:1.5b-instruct-q4_K_S
qwen2.5-coder:1.5b-instruct-q4_K_M
qwen2.5-coder:1.5b-instruct-q5_0
qwen2.5-coder:1.5b-instruct-q5_1
qwen2.5-coder:1.5b-instruct-q5_K_S
qwen2.5-coder:1.5b-instruct-q5_K_M
qwen2.5-coder:1.5b-instruct-q6_K
qwen2.5-coder:1.5b-instruct-q8_0
qwen2.5-coder:1.5b-instruct-fp16
qwen2.5-coder:3b-base
qwen2.5-coder:3b-base-q2_K
qwen2.5-coder:3b-base-q3_K_S
qwen2.5-coder:3b-base-q3_K_M
qwen2.5-coder:3b-base-q3_K_L
qwen2.5-coder:3b-base-q4_0
qwen2.5-coder:3b-base-q4_1
qwen2.5-coder:3b-base-q4_K_S
qwen2.5-coder:3b-base-q4_K_M
qwen2.5-coder:3b-base-q5_0
qwen2.5-coder:3b-base-q5_1
qwen2.5-coder:3b-base-q5_K_S
qwen2.5-coder:3b-base-q5_K_M
qwen2.5-coder:3b-base-q6_K
qwen2.5-coder:3b-base-q8_0
qwen2.5-coder:3b-base-fp16
qwen2.5-coder:3b-instruct
qwen2.5-coder:3b-instruct-q2_K
qwen2.5-coder:3b-instruct-q3_K_S
qwen2.5-coder:3b-instruct-q3_K_M
qwen2.5-coder:3b-instruct-q3_K_L
qwen2.5-coder:3b-instruct-q4_0
qwen2.5-coder:3b-instruct-q4_1
qwen2.5-coder:3b-instruct-q4_K_S
qwen2.5-coder:3b-instruct-q4_K_M
qwen2.5-coder:3b-instruct-q5_0
qwen2.5-coder:3b-instruct-q5_1
qwen2.5-coder:3b-instruct-q5_K_S
qwen2.5-coder:3b-instruct-q5_K_M
qwen2.5-coder:3b-instruct-q6_K
qwen2.5-coder:3b-instruct-q8_0
qwen2.5-coder:3b-instruct-fp16
qwen2.5-coder:7b-base
qwen2.5-coder:7b-base-q2_K
qwen2.5-coder:7b-base-q3_K_S
qwen2.5-coder:7b-base-q3_K_M
qwen2.5-coder:7b-base-q3_K_L
qwen2.5-coder:7b-base-q4_0
qwen2.5-coder:7b-base-q4_1
qwen2.5-coder:7b-base-q4_K_S
qwen2.5-coder:7b-base-q4_K_M
qwen2.5-coder:7b-base-q5_0
qwen2.5-coder:7b-base-q5_1
qwen2.5-coder:7b-base-q5_K_S
qwen2.5-coder:7b-base-q5_K_M
qwen2.5-coder:7b-base-q6_K
qwen2.5-coder:7b-base-q8_0
qwen2.5-coder:7b-base-fp16
qwen2.5-coder:7b-instruct
qwen2.5-coder:7b-instruct-q2_K
qwen2.5-coder:7b-instruct-q3_K_S
qwen2.5-coder:7b-instruct-q3_K_M
qwen2.5-coder:7b-instruct-q3_K_L
qwen2.5-coder:7b-instruct-q4_0
qwen2.5-coder:7b-instruct-q4_1
qwen2.5-coder:7b-instruct-q4_K_S
qwen2.5-coder:7b-instruct-q4_K_M
qwen2.5-coder:7b-instruct-q5_0
qwen2.5-coder:7b-instruct-q5_1
qwen2.5-coder:7b-instruct-q5_K_S
qwen2.5-coder:7b-instruct-q5_K_M
qwen2.5-coder:7b-instruct-q6_K
qwen2.5-coder:7b-instruct-q8_0
qwen2.5-coder:7b-instruct-fp16
qwen2.5-coder:14b-base
qwen2.5-coder:14b-base-q2_K
qwen2.5-coder:14b-base-q3_K_S
qwen2.5-coder:14b-base-q3_K_M
qwen2.5-coder:14b-base-q3_K_L
qwen2.5-coder:14b-base-q4_0
qwen2.5-coder:14b-base-q4_1
qwen2.5-coder:14b-base-q4_K_S
qwen2.5-coder:14b-base-q4_K_M
qwen2.5-coder:14b-base-q5_0
qwen2.5-coder:14b-base-q5_1
qwen2.5-coder:14b-base-q5_K_S
qwen2.5-coder:14b-base-q5_K_M
qwen2.5-coder:14b-base-q6_K
qwen2.5-coder:14b-base-q8_0
qwen2.5-coder:14b-base-fp16
qwen2.5-coder:14b-instruct
qwen2.5-coder:14b-instruct-q2_K
qwen2.5-coder:14b-instruct-q3_K_S
qwen2.5-coder:14b-instruct-q3_K_M
qwen2.5-coder:14b-instruct-q3_K_L
qwen2.5-coder:14b-instruct-q4_0
qwen2.5-coder:14b-instruct-q4_1
qwen2.5-coder:14b-instruct-q4_K_S
qwen2.5-coder:14b-instruct-q4_K_M
qwen2.5-coder:14b-instruct-q5_0
qwen2.5-coder:14b-instruct-q5_1
qwen2.5-coder:14b-instruct-q5_K_S
qwen2.5-coder:14b-instruct-q5_K_M
qwen2.5-coder:14b-instruct-q6_K
qwen2.5-coder:14b-instruct-q8_0
qwen2.5-coder:14b-instruct-fp16
qwen2.5-coder:32b-base
qwen2.5-coder:32b-base-q2_K
qwen2.5-coder:32b-base-q3_K_S
qwen2.5-coder:32b-base-q3_K_M
qwen2.5-coder:32b-base-q3_K_L
qwen2.5-coder:32b-base-q4_0
qwen2.5-coder:32b-base-q4_1
qwen2.5-coder:32b-base-q4_K_S
qwen2.5-coder:32b-base-q4_K_M
qwen2.5-coder:32b-base-q5_0
qwen2.5-coder:32b-base-q5_1
qwen2.5-coder:32b-base-q5_K_S
qwen2.5-coder:32b-base-q5_K_M
qwen2.5-coder:32b-base-q6_K
qwen2.5-coder:32b-base-q8_0
qwen2.5-coder:32b-base-fp16
qwen2.5-coder:32b-instruct
qwen2.5-coder:32b-instruct-q2_K
qwen2.5-coder:32b-instruct-q3_K_S
qwen2.5-coder:32b-instruct-q3_K_M
qwen2.5-coder:32b-instruct-q3_K_L
qwen2.5-coder:32b-instruct-q4_0
qwen2.5-coder:32b-instruct-q4_1
qwen2.5-coder:32b-instruct-q4_K_S
qwen2.5-coder:32b-instruct-q4_K_M
qwen2.5-coder:32b-instruct-q5_0
qwen2.5-coder:32b-instruct-q5_1
qwen2.5-coder:32b-instruct-q5_K_S
qwen2.5-coder:32b-instruct-q5_K_M
qwen2.5-coder:32b-instruct-q6_K
qwen2.5-coder:32b-instruct-q8_0
qwen2.5-coder:32b-instruct-fp16
phi4:latest
phi4
phi4:14b
phi4:14b-q4_K_M
phi4:14b-q8_0
phi4:14b-fp16
mxbai-embed-large:latest
mxbai-embed-large
mxbai-embed-large:v1
mxbai-embed-large:335m
mxbai-embed-large:335m-v1-fp16
gemma:latest
gemma
gemma:instruct
gemma:text
gemma:v1.1
gemma:2b
gemma:7b
gemma:2b-instruct
gemma:2b-instruct-q2_K
gemma:2b-instruct-v1.1-q2_K
gemma:2b-instruct-q3_K_S
gemma:2b-instruct-v1.1-q3_K_S
gemma:2b-instruct-q3_K_M
gemma:2b-instruct-v1.1-q3_K_M
gemma:2b-instruct-q3_K_L
gemma:2b-instruct-v1.1-q3_K_L
gemma:2b-instruct-q4_0
gemma:2b-instruct-v1.1-q4_0
gemma:2b-instruct-q4_1
gemma:2b-instruct-v1.1-q4_1
gemma:2b-instruct-q4_K_S
gemma:2b-instruct-v1.1-q4_K_S
gemma:2b-instruct-q4_K_M
gemma:2b-instruct-v1.1-q4_K_M
gemma:2b-instruct-q5_0
gemma:2b-instruct-v1.1-q5_0
gemma:2b-instruct-q5_1
gemma:2b-instruct-v1.1-q5_1
gemma:2b-instruct-q5_K_S
gemma:2b-instruct-v1.1-q5_K_S
gemma:2b-instruct-q5_K_M
gemma:2b-instruct-v1.1-q5_K_M
gemma:2b-instruct-q6_K
gemma:2b-instruct-v1.1-q6_K
gemma:2b-instruct-q8_0
gemma:2b-instruct-v1.1-q8_0
gemma:2b-instruct-fp16
gemma:2b-instruct-v1.1-fp16
gemma:2b-text
gemma:2b-text-q2_K
gemma:2b-text-q3_K_S
gemma:2b-text-q3_K_M
gemma:2b-text-q3_K_L
gemma:2b-text-q4_0
gemma:2b-text-q4_1
gemma:2b-text-q4_K_S
gemma:2b-text-q4_K_M
gemma:2b-text-q5_0
gemma:2b-text-q5_1
gemma:2b-text-q5_K_S
gemma:2b-text-q5_K_M
gemma:2b-text-q6_K
gemma:2b-text-q8_0
gemma:2b-text-fp16
gemma:2b-v1.1
gemma:7b-instruct
gemma:7b-instruct-q2_K
gemma:7b-instruct-v1.1-q2_K
gemma:7b-instruct-q3_K_S
gemma:7b-instruct-v1.1-q3_K_S
gemma:7b-instruct-q3_K_M
gemma:7b-instruct-v1.1-q3_K_M
gemma:7b-instruct-q3_K_L
gemma:7b-instruct-v1.1-q3_K_L
gemma:7b-instruct-q4_0
gemma:7b-instruct-v1.1-q4_0
gemma:7b-instruct-q4_1
gemma:7b-instruct-v1.1-q4_1
gemma:7b-instruct-q4_K_S
gemma:7b-instruct-v1.1-q4_K_S
gemma:7b-instruct-q4_K_M
gemma:7b-instruct-v1.1-q4_K_M
gemma:7b-instruct-q5_0
gemma:7b-instruct-v1.1-q5_0
gemma:7b-instruct-q5_1
gemma:7b-instruct-v1.1-q5_1
gemma:7b-instruct-q5_K_S
gemma:7b-instruct-v1.1-q5_K_S
gemma:7b-instruct-q5_K_M
gemma:7b-instruct-v1.1-q5_K_M
gemma:7b-instruct-q6_K
gemma:7b-instruct-v1.1-q6_K
gemma:7b-instruct-q8_0
gemma:7b-instruct-v1.1-q8_0
gemma:7b-instruct-fp16
gemma:7b-instruct-v1.1-fp16
gemma:7b-text
gemma:7b-text-q2_K
gemma:7b-text-q3_K_S
gemma:7b-text-q3_K_M
gemma:7b-text-q3_K_L
gemma:7b-text-q4_0
gemma:7b-text-q4_1
gemma:7b-text-q4_K_S
gemma:7b-text-q4_K_M
gemma:7b-text-q5_0
gemma:7b-text-q5_1
gemma:7b-text-q5_K_S
gemma:7b-text-q5_K_M
gemma:7b-text-q6_K
gemma:7b-text-q8_0
gemma:7b-text-fp16
gemma:7b-v1.1
qwen:latest
qwen
qwen:0.5b
qwen:1.8b
qwen:4b
qwen:7b
qwen:14b
qwen:32b
qwen:72b
qwen:110b
qwen:0.5b-chat
qwen:0.5b-chat-v1.5-q2_K
qwen:0.5b-chat-v1.5-q3_K_S
qwen:0.5b-chat-v1.5-q3_K_M
qwen:0.5b-chat-v1.5-q3_K_L
qwen:0.5b-chat-v1.5-q4_0
qwen:0.5b-chat-v1.5-q4_1
qwen:0.5b-chat-v1.5-q4_K_S
qwen:0.5b-chat-v1.5-q4_K_M
qwen:0.5b-chat-v1.5-q5_0
qwen:0.5b-chat-v1.5-q5_1
qwen:0.5b-chat-v1.5-q5_K_S
qwen:0.5b-chat-v1.5-q5_K_M
qwen:0.5b-chat-v1.5-q6_K
qwen:0.5b-chat-v1.5-q8_0
qwen:0.5b-chat-v1.5-fp16
qwen:0.5b-text
qwen:0.5b-text-v1.5-q2_K
qwen:0.5b-text-v1.5-q3_K_S
qwen:0.5b-text-v1.5-q3_K_M
qwen:0.5b-text-v1.5-q3_K_L
qwen:0.5b-text-v1.5-q4_0
qwen:0.5b-text-v1.5-q4_1
qwen:0.5b-text-v1.5-q4_K_S
qwen:0.5b-text-v1.5-q4_K_M
qwen:0.5b-text-v1.5-q5_0
qwen:0.5b-text-v1.5-q5_1
qwen:0.5b-text-v1.5-q5_K_S
qwen:0.5b-text-v1.5-q5_K_M
qwen:0.5b-text-v1.5-q6_K
qwen:0.5b-text-v1.5-q8_0
qwen:0.5b-text-v1.5-fp16
qwen:1.8b-chat
qwen:1.8b-chat-q2_K
qwen:1.8b-chat-v1.5-q2_K
qwen:1.8b-chat-q3_K_S
qwen:1.8b-chat-v1.5-q3_K_S
qwen:1.8b-chat-q3_K_M
qwen:1.8b-chat-v1.5-q3_K_M
qwen:1.8b-chat-q3_K_L
qwen:1.8b-chat-v1.5-q3_K_L
qwen:1.8b-chat-q4_0
qwen:1.8b-chat-v1.5-q4_0
qwen:1.8b-chat-q4_1
qwen:1.8b-chat-v1.5-q4_1
qwen:1.8b-chat-q4_K_S
qwen:1.8b-chat-v1.5-q4_K_S
qwen:1.8b-chat-q4_K_M
qwen:1.8b-chat-v1.5-q4_K_M
qwen:1.8b-chat-q5_0
qwen:1.8b-chat-v1.5-q5_0
qwen:1.8b-chat-q5_1
qwen:1.8b-chat-v1.5-q5_1
qwen:1.8b-chat-q5_K_S
qwen:1.8b-chat-v1.5-q5_K_S
qwen:1.8b-chat-q5_K_M
qwen:1.8b-chat-v1.5-q5_K_M
qwen:1.8b-chat-q6_K
qwen:1.8b-chat-v1.5-q6_K
qwen:1.8b-chat-q8_0
qwen:1.8b-chat-v1.5-q8_0
qwen:1.8b-chat-fp16
qwen:1.8b-chat-v1.5-fp16
qwen:1.8b-text
qwen:1.8b-text-q2_K
qwen:1.8b-text-v1.5-q2_K
qwen:1.8b-text-q3_K_S
qwen:1.8b-text-v1.5-q3_K_S
qwen:1.8b-text-q3_K_M
qwen:1.8b-text-v1.5-q3_K_M
qwen:1.8b-text-q3_K_L
qwen:1.8b-text-v1.5-q3_K_L
qwen:1.8b-text-q4_0
qwen:1.8b-text-v1.5-q4_0
qwen:1.8b-text-q4_1
qwen:1.8b-text-v1.5-q4_1
qwen:1.8b-text-q4_K_S
qwen:1.8b-text-v1.5-q4_K_S
qwen:1.8b-text-q4_K_M
qwen:1.8b-text-v1.5-q4_K_M
qwen:1.8b-text-q5_0
qwen:1.8b-text-v1.5-q5_0
qwen:1.8b-text-q5_1
qwen:1.8b-text-v1.5-q5_1
qwen:1.8b-text-q5_K_S
qwen:1.8b-text-v1.5-q5_K_S
qwen:1.8b-text-q5_K_M
qwen:1.8b-text-v1.5-q5_K_M
qwen:1.8b-text-q6_K
qwen:1.8b-text-v1.5-q6_K
qwen:1.8b-text-q8_0
qwen:1.8b-text-v1.5-q8_0
qwen:1.8b-text-fp16
qwen:1.8b-text-v1.5-fp16
qwen:4b-chat
qwen:4b-chat-v1.5-q2_K
qwen:4b-chat-v1.5-q3_K_S
qwen:4b-chat-v1.5-q3_K_M
qwen:4b-chat-v1.5-q3_K_L
qwen:4b-chat-v1.5-q4_0
qwen:4b-chat-v1.5-q4_1
qwen:4b-chat-v1.5-q4_K_S
qwen:4b-chat-v1.5-q4_K_M
qwen:4b-chat-v1.5-q5_0
qwen:4b-chat-v1.5-q5_1
qwen:4b-chat-v1.5-q5_K_S
qwen:4b-chat-v1.5-q5_K_M
qwen:4b-chat-v1.5-q6_K
qwen:4b-chat-v1.5-q8_0
qwen:4b-chat-v1.5-fp16
qwen:4b-text
qwen:4b-text-v1.5-q2_K
qwen:4b-text-v1.5-q3_K_S
qwen:4b-text-v1.5-q3_K_M
qwen:4b-text-v1.5-q3_K_L
qwen:4b-text-v1.5-q4_0
qwen:4b-text-v1.5-q4_1
qwen:4b-text-v1.5-q4_K_S
qwen:4b-text-v1.5-q4_K_M
qwen:4b-text-v1.5-q5_0
qwen:4b-text-v1.5-q5_1
qwen:4b-text-v1.5-q5_K_S
qwen:4b-text-v1.5-q5_K_M
qwen:4b-text-v1.5-q6_K
qwen:4b-text-v1.5-q8_0
qwen:4b-text-v1.5-fp16
qwen:7b-chat
qwen:7b-chat-q2_K
qwen:7b-chat-v1.5-q2_K
qwen:7b-chat-q3_K_S
qwen:7b-chat-v1.5-q3_K_S
qwen:7b-chat-q3_K_M
qwen:7b-chat-v1.5-q3_K_M
qwen:7b-chat-q3_K_L
qwen:7b-chat-v1.5-q3_K_L
qwen:7b-chat-q4_0
qwen:7b-chat-v1.5-q4_0
qwen:7b-chat-q4_1
qwen:7b-chat-v1.5-q4_1
qwen:7b-chat-q4_K_S
qwen:7b-chat-v1.5-q4_K_S
qwen:7b-chat-q4_K_M
qwen:7b-chat-v1.5-q4_K_M
qwen:7b-chat-q5_0
qwen:7b-chat-v1.5-q5_0
qwen:7b-chat-q5_1
qwen:7b-chat-v1.5-q5_1
qwen:7b-chat-q5_K_S
qwen:7b-chat-v1.5-q5_K_S
qwen:7b-chat-q5_K_M
qwen:7b-chat-v1.5-q5_K_M
qwen:7b-chat-q6_K
qwen:7b-chat-v1.5-q6_K
qwen:7b-chat-q8_0
qwen:7b-chat-v1.5-q8_0
qwen:7b-chat-fp16
qwen:7b-chat-v1.5-fp16
qwen:7b-text
qwen:7b-text-v1.5-q2_K
qwen:7b-text-v1.5-q3_K_S
qwen:7b-text-v1.5-q3_K_M
qwen:7b-text-v1.5-q3_K_L
qwen:7b-text-v1.5-q4_0
qwen:7b-text-v1.5-q4_1
qwen:7b-text-v1.5-q4_K_S
qwen:7b-text-v1.5-q4_K_M
qwen:7b-text-v1.5-q5_0
qwen:7b-text-v1.5-q5_1
qwen:7b-text-v1.5-q5_K_S
qwen:7b-text-v1.5-q5_K_M
qwen:7b-text-v1.5-q6_K
qwen:7b-text-v1.5-q8_0
qwen:7b-text-v1.5-fp16
qwen:7b-q2_K
qwen:7b-q3_K_S
qwen:7b-q3_K_M
qwen:7b-q3_K_L
qwen:7b-q4_0
qwen:7b-q4_1
qwen:7b-q4_K_S
qwen:7b-q4_K_M
qwen:7b-q5_0
qwen:7b-q5_1
qwen:7b-q5_K_S
qwen:7b-q5_K_M
qwen:7b-q6_K
qwen:7b-q8_0
qwen:7b-fp16
qwen:14b-chat
qwen:14b-chat-q2_K
qwen:14b-chat-v1.5-q2_K
qwen:14b-chat-q3_K_S
qwen:14b-chat-v1.5-q3_K_S
qwen:14b-chat-q3_K_M
qwen:14b-chat-v1.5-q3_K_M
qwen:14b-chat-q3_K_L
qwen:14b-chat-v1.5-q3_K_L
qwen:14b-chat-q4_0
qwen:14b-chat-v1.5-q4_0
qwen:14b-chat-q4_1
qwen:14b-chat-v1.5-q4_1
qwen:14b-chat-q4_K_S
qwen:14b-chat-v1.5-q4_K_S
qwen:14b-chat-q4_K_M
qwen:14b-chat-v1.5-q4_K_M
qwen:14b-chat-q5_0
qwen:14b-chat-v1.5-q5_0
qwen:14b-chat-q5_1
qwen:14b-chat-v1.5-q5_1
qwen:14b-chat-q5_K_S
qwen:14b-chat-v1.5-q5_K_S
qwen:14b-chat-q5_K_M
qwen:14b-chat-v1.5-q5_K_M
qwen:14b-chat-q6_K
qwen:14b-chat-v1.5-q6_K
qwen:14b-chat-q8_0
qwen:14b-chat-v1.5-q8_0
qwen:14b-chat-fp16
qwen:14b-chat-v1.5-fp16
qwen:14b-text
qwen:14b-text-q2_K
qwen:14b-text-v1.5-q2_K
qwen:14b-text-q3_K_S
qwen:14b-text-v1.5-q3_K_S
qwen:14b-text-q3_K_M
qwen:14b-text-v1.5-q3_K_M
qwen:14b-text-q3_K_L
qwen:14b-text-v1.5-q3_K_L
qwen:14b-text-q4_0
qwen:14b-text-v1.5-q4_0
qwen:14b-text-q4_1
qwen:14b-text-v1.5-q4_1
qwen:14b-text-q4_K_S
qwen:14b-text-v1.5-q4_K_S
qwen:14b-text-q4_K_M
qwen:14b-text-v1.5-q4_K_M
qwen:14b-text-q5_0
qwen:14b-text-v1.5-q5_0
qwen:14b-text-q5_1
qwen:14b-text-v1.5-q5_1
qwen:14b-text-q5_K_S
qwen:14b-text-v1.5-q5_K_S
qwen:14b-text-q5_K_M
qwen:14b-text-v1.5-q5_K_M
qwen:14b-text-q6_K
qwen:14b-text-v1.5-q6_K
qwen:14b-text-q8_0
qwen:14b-text-v1.5-q8_0
qwen:14b-text-fp16
qwen:14b-text-v1.5-fp16
qwen:32b-chat
qwen:32b-chat-v1.5-q2_K
qwen:32b-chat-v1.5-q3_K_S
qwen:32b-chat-v1.5-q3_K_M
qwen:32b-chat-v1.5-q3_K_L
qwen:32b-chat-v1.5-q4_0
qwen:32b-chat-v1.5-q4_1
qwen:32b-chat-v1.5-q4_K_S
qwen:32b-chat-v1.5-q4_K_M
qwen:32b-chat-v1.5-q5_0
qwen:32b-chat-v1.5-q5_1
qwen:32b-chat-v1.5-q5_K_S
qwen:32b-chat-v1.5-q5_K_M
qwen:32b-chat-v1.5-q6_K
qwen:32b-chat-v1.5-q8_0
qwen:32b-chat-v1.5-fp16
qwen:32b-text
qwen:32b-text-v1.5-q2_K
qwen:32b-text-v1.5-q3_K_S
qwen:32b-text-v1.5-q3_K_M
qwen:32b-text-v1.5-q3_K_L
qwen:32b-text-v1.5-q4_0
qwen:32b-text-v1.5-q4_1
qwen:32b-text-v1.5-q4_K_S
qwen:32b-text-v1.5-q5_0
qwen:32b-text-v1.5-q5_1
qwen:32b-text-v1.5-q8_0
qwen:72b-chat
qwen:72b-chat-q2_K
qwen:72b-chat-v1.5-q2_K
qwen:72b-chat-q3_K_S
qwen:72b-chat-v1.5-q3_K_S
qwen:72b-chat-q3_K_M
qwen:72b-chat-v1.5-q3_K_M
qwen:72b-chat-q3_K_L
qwen:72b-chat-v1.5-q3_K_L
qwen:72b-chat-q4_0
qwen:72b-chat-v1.5-q4_0
qwen:72b-chat-q4_1
qwen:72b-chat-v1.5-q4_1
qwen:72b-chat-q4_K_S
qwen:72b-chat-v1.5-q4_K_S
qwen:72b-chat-q4_K_M
qwen:72b-chat-v1.5-q4_K_M
qwen:72b-chat-q5_0
qwen:72b-chat-v1.5-q5_0
qwen:72b-chat-q5_1
qwen:72b-chat-v1.5-q5_1
qwen:72b-chat-q5_K_S
qwen:72b-chat-v1.5-q5_K_S
qwen:72b-chat-q5_K_M
qwen:72b-chat-v1.5-q5_K_M
qwen:72b-chat-q6_K
qwen:72b-chat-v1.5-q6_K
qwen:72b-chat-q8_0
qwen:72b-chat-v1.5-q8_0
qwen:72b-chat-fp16
qwen:72b-chat-v1.5-fp16
qwen:72b-text
qwen:72b-text-q2_K
qwen:72b-text-v1.5-q2_K
qwen:72b-text-q3_K_S
qwen:72b-text-v1.5-q3_K_S
qwen:72b-text-q3_K_M
qwen:72b-text-v1.5-q3_K_M
qwen:72b-text-q3_K_L
qwen:72b-text-v1.5-q3_K_L
qwen:72b-text-q4_0
qwen:72b-text-v1.5-q4_0
qwen:72b-text-q4_1
qwen:72b-text-v1.5-q4_1
qwen:72b-text-q4_K_S
qwen:72b-text-v1.5-q4_K_S
qwen:72b-text-q4_K_M
qwen:72b-text-v1.5-q4_K_M
qwen:72b-text-q5_0
qwen:72b-text-v1.5-q5_0
qwen:72b-text-q5_1
qwen:72b-text-v1.5-q5_1
qwen:72b-text-q5_K_S
qwen:72b-text-v1.5-q5_K_S
qwen:72b-text-q5_K_M
qwen:72b-text-v1.5-q5_K_M
qwen:72b-text-q6_K
qwen:72b-text-v1.5-q6_K
qwen:72b-text-q8_0
qwen:72b-text-v1.5-q8_0
qwen:72b-text-fp16
qwen:72b-text-v1.5-fp16
qwen:110b-chat
qwen:110b-chat-v1.5-q2_K
qwen:110b-chat-v1.5-q3_K_S
qwen:110b-chat-v1.5-q3_K_M
qwen:110b-chat-v1.5-q3_K_L
qwen:110b-chat-v1.5-q4_0
qwen:110b-chat-v1.5-q4_1
qwen:110b-chat-v1.5-q4_K_S
qwen:110b-chat-v1.5-q4_K_M
qwen:110b-chat-v1.5-q5_0
qwen:110b-chat-v1.5-q5_1
qwen:110b-chat-v1.5-q5_K_S
qwen:110b-chat-v1.5-q5_K_M
qwen:110b-chat-v1.5-q6_K
qwen:110b-chat-v1.5-q8_0
qwen:110b-chat-v1.5-fp16
qwen:110b-text-v1.5-q2_K
qwen:110b-text-v1.5-q3_K_S
qwen:110b-text-v1.5-q3_K_M
qwen:110b-text-v1.5-q3_K_L
qwen:110b-text-v1.5-q4_0
qwen:110b-text-v1.5-q4_1
qwen:110b-text-v1.5-q4_K_S
qwen:110b-text-v1.5-q4_K_M
qwen:110b-text-v1.5-q5_0
qwen:110b-text-v1.5-q5_1
qwen:110b-text-v1.5-q5_K_S
qwen:110b-text-v1.5-q5_K_M
qwen:110b-text-v1.5-q6_K
qwen:110b-text-v1.5-q8_0
qwen:110b-text-v1.5-fp16
llama2:latest
llama2
llama2:chat
llama2:text
llama2:7b
llama2:13b
llama2:70b
llama2:7b-chat
llama2:7b-chat-q2_K
llama2:7b-chat-q3_K_S
llama2:7b-chat-q3_K_M
llama2:7b-chat-q3_K_L
llama2:7b-chat-q4_0
llama2:7b-chat-q4_1
llama2:7b-chat-q4_K_S
llama2:7b-chat-q4_K_M
llama2:7b-chat-q5_0
llama2:7b-chat-q5_1
llama2:7b-chat-q5_K_S
llama2:7b-chat-q5_K_M
llama2:7b-chat-q6_K
llama2:7b-chat-q8_0
llama2:7b-chat-fp16
llama2:7b-text
llama2:7b-text-q2_K
llama2:7b-text-q3_K_S
llama2:7b-text-q3_K_M
llama2:7b-text-q3_K_L
llama2:7b-text-q4_0
llama2:7b-text-q4_1
llama2:7b-text-q4_K_S
llama2:7b-text-q4_K_M
llama2:7b-text-q5_0
llama2:7b-text-q5_1
llama2:7b-text-q5_K_S
llama2:7b-text-q5_K_M
llama2:7b-text-q6_K
llama2:7b-text-q8_0
llama2:7b-text-fp16
llama2:13b-chat
llama2:13b-chat-q2_K
llama2:13b-chat-q3_K_S
llama2:13b-chat-q3_K_M
llama2:13b-chat-q3_K_L
llama2:13b-chat-q4_0
llama2:13b-chat-q4_1
llama2:13b-chat-q4_K_S
llama2:13b-chat-q4_K_M
llama2:13b-chat-q5_0
llama2:13b-chat-q5_1
llama2:13b-chat-q5_K_S
llama2:13b-chat-q5_K_M
llama2:13b-chat-q6_K
llama2:13b-chat-q8_0
llama2:13b-chat-fp16
llama2:13b-text
llama2:13b-text-q2_K
llama2:13b-text-q3_K_S
llama2:13b-text-q3_K_M
llama2:13b-text-q3_K_L
llama2:13b-text-q4_0
llama2:13b-text-q4_1
llama2:13b-text-q4_K_S
llama2:13b-text-q4_K_M
llama2:13b-text-q5_0
llama2:13b-text-q5_1
llama2:13b-text-q5_K_S
llama2:13b-text-q5_K_M
llama2:13b-text-q6_K
llama2:13b-text-q8_0
llama2:13b-text-fp16
llama2:70b-chat
llama2:70b-chat-q2_K
llama2:70b-chat-q3_K_S
llama2:70b-chat-q3_K_M
llama2:70b-chat-q3_K_L
llama2:70b-chat-q4_0
llama2:70b-chat-q4_1
llama2:70b-chat-q4_K_S
llama2:70b-chat-q4_K_M
llama2:70b-chat-q5_0
llama2:70b-chat-q5_1
llama2:70b-chat-q5_K_S
llama2:70b-chat-q5_K_M
llama2:70b-chat-q6_K
llama2:70b-chat-q8_0
llama2:70b-chat-fp16
llama2:70b-text
llama2:70b-text-q2_K
llama2:70b-text-q3_K_S
llama2:70b-text-q3_K_M
llama2:70b-text-q3_K_L
llama2:70b-text-q4_0
llama2:70b-text-q4_1
llama2:70b-text-q4_K_S
llama2:70b-text-q4_K_M
llama2:70b-text-q5_0
llama2:70b-text-q5_1
llama2:70b-text-q5_K_S
llama2:70b-text-q5_K_M
llama2:70b-text-q6_K
llama2:70b-text-q8_0
llama2:70b-text-fp16
qwen2:latest
qwen2
qwen2:0.5b
qwen2:1.5b
qwen2:7b
qwen2:72b
qwen2:0.5b-instruct
qwen2:0.5b-instruct-q2_K
qwen2:0.5b-instruct-q3_K_S
qwen2:0.5b-instruct-q3_K_M
qwen2:0.5b-instruct-q3_K_L
qwen2:0.5b-instruct-q4_0
qwen2:0.5b-instruct-q4_1
qwen2:0.5b-instruct-q4_K_S
qwen2:0.5b-instruct-q4_K_M
qwen2:0.5b-instruct-q5_0
qwen2:0.5b-instruct-q5_1
qwen2:0.5b-instruct-q5_K_S
qwen2:0.5b-instruct-q5_K_M
qwen2:0.5b-instruct-q6_K
qwen2:0.5b-instruct-q8_0
qwen2:0.5b-instruct-fp16
qwen2:1.5b-instruct
qwen2:1.5b-instruct-q2_K
qwen2:1.5b-instruct-q3_K_S
qwen2:1.5b-instruct-q3_K_M
qwen2:1.5b-instruct-q3_K_L
qwen2:1.5b-instruct-q4_0
qwen2:1.5b-instruct-q4_1
qwen2:1.5b-instruct-q4_K_S
qwen2:1.5b-instruct-q4_K_M
qwen2:1.5b-instruct-q5_0
qwen2:1.5b-instruct-q5_1
qwen2:1.5b-instruct-q5_K_S
qwen2:1.5b-instruct-q5_K_M
qwen2:1.5b-instruct-q6_K
qwen2:1.5b-instruct-q8_0
qwen2:1.5b-instruct-fp16
qwen2:7b-instruct
qwen2:7b-instruct-q2_K
qwen2:7b-instruct-q3_K_S
qwen2:7b-instruct-q3_K_M
qwen2:7b-instruct-q3_K_L
qwen2:7b-instruct-q4_0
qwen2:7b-instruct-q4_1
qwen2:7b-instruct-q4_K_S
qwen2:7b-instruct-q4_K_M
qwen2:7b-instruct-q5_0
qwen2:7b-instruct-q5_1
qwen2:7b-instruct-q5_K_S
qwen2:7b-instruct-q5_K_M
qwen2:7b-instruct-q6_K
qwen2:7b-instruct-q8_0
qwen2:7b-instruct-fp16
qwen2:7b-text
qwen2:7b-text-q2_K
qwen2:7b-text-q3_K_S
qwen2:7b-text-q3_K_M
qwen2:7b-text-q3_K_L
qwen2:7b-text-q4_0
qwen2:7b-text-q4_1
qwen2:7b-text-q4_K_S
qwen2:7b-text-q4_K_M
qwen2:7b-text-q5_0
qwen2:7b-text-q5_1
qwen2:7b-text-q8_0
qwen2:72b-instruct
qwen2:72b-instruct-q2_K
qwen2:72b-instruct-q3_K_S
qwen2:72b-instruct-q3_K_M
qwen2:72b-instruct-q3_K_L
qwen2:72b-instruct-q4_0
qwen2:72b-instruct-q4_1
qwen2:72b-instruct-q4_K_S
qwen2:72b-instruct-q4_K_M
qwen2:72b-instruct-q5_0
qwen2:72b-instruct-q5_1
qwen2:72b-instruct-q5_K_S
qwen2:72b-instruct-q5_K_M
qwen2:72b-instruct-q6_K
qwen2:72b-instruct-q8_0
qwen2:72b-instruct-fp16
qwen2:72b-text
qwen2:72b-text-q2_K
qwen2:72b-text-q3_K_S
qwen2:72b-text-q3_K_M
qwen2:72b-text-q3_K_L
qwen2:72b-text-q4_0
qwen2:72b-text-q4_1
qwen2:72b-text-q4_K_S
qwen2:72b-text-q4_K_M
qwen2:72b-text-q5_0
qwen2:72b-text-q5_1
qwen2:72b-text-q5_K_S
qwen2:72b-text-q5_K_M
qwen2:72b-text-q6_K
qwen2:72b-text-q8_0
qwen2:72b-text-fp16
minicpm-v:latest
minicpm-v
minicpm-v:8b
minicpm-v:8b-2.6-q2_K
minicpm-v:8b-2.6-q3_K_S
minicpm-v:8b-2.6-q3_K_M
minicpm-v:8b-2.6-q3_K_L
minicpm-v:8b-2.6-q4_0
minicpm-v:8b-2.6-q4_1
minicpm-v:8b-2.6-q4_K_S
minicpm-v:8b-2.6-q4_K_M
minicpm-v:8b-2.6-q5_0
minicpm-v:8b-2.6-q5_1
minicpm-v:8b-2.6-q5_K_S
minicpm-v:8b-2.6-q5_K_M
minicpm-v:8b-2.6-q6_K
minicpm-v:8b-2.6-q8_0
minicpm-v:8b-2.6-fp16
codellama:latest
codellama
codellama:code
codellama:instruct
codellama:python
codellama:7b
codellama:13b
codellama:34b
codellama:70b
codellama:7b-code
codellama:7b-code-q2_K
codellama:7b-code-q3_K_S
codellama:7b-code-q3_K_M
codellama:7b-code-q3_K_L
codellama:7b-code-q4_0
codellama:7b-code-q4_1
codellama:7b-code-q4_K_S
codellama:7b-code-q4_K_M
codellama:7b-code-q5_0
codellama:7b-code-q5_1
codellama:7b-code-q5_K_S
codellama:7b-code-q5_K_M
codellama:7b-code-q6_K
codellama:7b-code-q8_0
codellama:7b-code-fp16
codellama:7b-instruct
codellama:7b-instruct-q2_K
codellama:7b-instruct-q3_K_S
codellama:7b-instruct-q3_K_M
codellama:7b-instruct-q3_K_L
codellama:7b-instruct-q4_0
codellama:7b-instruct-q4_1
codellama:7b-instruct-q4_K_S
codellama:7b-instruct-q4_K_M
codellama:7b-instruct-q5_0
codellama:7b-instruct-q5_1
codellama:7b-instruct-q5_K_S
codellama:7b-instruct-q5_K_M
codellama:7b-instruct-q6_K
codellama:7b-instruct-q8_0
codellama:7b-instruct-fp16
codellama:7b-python
codellama:7b-python-q2_K
codellama:7b-python-q3_K_S
codellama:7b-python-q3_K_M
codellama:7b-python-q3_K_L
codellama:7b-python-q4_0
codellama:7b-python-q4_1
codellama:7b-python-q4_K_S
codellama:7b-python-q4_K_M
codellama:7b-python-q5_0
codellama:7b-python-q5_1
codellama:7b-python-q5_K_S
codellama:7b-python-q5_K_M
codellama:7b-python-q6_K
codellama:7b-python-q8_0
codellama:7b-python-fp16
codellama:13b-code
codellama:13b-code-q2_K
codellama:13b-code-q3_K_S
codellama:13b-code-q3_K_M
codellama:13b-code-q3_K_L
codellama:13b-code-q4_0
codellama:13b-code-q4_1
codellama:13b-code-q4_K_S
codellama:13b-code-q4_K_M
codellama:13b-code-q5_0
codellama:13b-code-q5_1
codellama:13b-code-q5_K_S
codellama:13b-code-q5_K_M
codellama:13b-code-q6_K
codellama:13b-code-q8_0
codellama:13b-code-fp16
codellama:13b-instruct
codellama:13b-instruct-q2_K
codellama:13b-instruct-q3_K_S
codellama:13b-instruct-q3_K_M
codellama:13b-instruct-q3_K_L
codellama:13b-instruct-q4_0
codellama:13b-instruct-q4_1
codellama:13b-instruct-q4_K_S
codellama:13b-instruct-q4_K_M
codellama:13b-instruct-q5_0
codellama:13b-instruct-q5_1
codellama:13b-instruct-q5_K_S
codellama:13b-instruct-q5_K_M
codellama:13b-instruct-q6_K
codellama:13b-instruct-q8_0
codellama:13b-instruct-fp16
codellama:13b-python
codellama:13b-python-q2_K
codellama:13b-python-q3_K_S
codellama:13b-python-q3_K_M
codellama:13b-python-q3_K_L
codellama:13b-python-q4_0
codellama:13b-python-q4_1
codellama:13b-python-q4_K_S
codellama:13b-python-q4_K_M
codellama:13b-python-q5_0
codellama:13b-python-q5_1
codellama:13b-python-q5_K_S
codellama:13b-python-q5_K_M
codellama:13b-python-q6_K
codellama:13b-python-q8_0
codellama:13b-python-fp16
codellama:34b-code
codellama:34b-code-q2_K
codellama:34b-code-q3_K_S
codellama:34b-code-q3_K_M
codellama:34b-code-q3_K_L
codellama:34b-code-q4_0
codellama:34b-code-q4_1
codellama:34b-code-q4_K_S
codellama:34b-code-q4_K_M
codellama:34b-code-q5_0
codellama:34b-code-q5_1
codellama:34b-code-q5_K_S
codellama:34b-code-q5_K_M
codellama:34b-code-q6_K
codellama:34b-code-q8_0
codellama:34b-instruct
codellama:34b-instruct-q2_K
codellama:34b-instruct-q3_K_S
codellama:34b-instruct-q3_K_M
codellama:34b-instruct-q3_K_L
codellama:34b-instruct-q4_0
codellama:34b-instruct-q4_1
codellama:34b-instruct-q4_K_S
codellama:34b-instruct-q4_K_M
codellama:34b-instruct-q5_0
codellama:34b-instruct-q5_1
codellama:34b-instruct-q5_K_S
codellama:34b-instruct-q5_K_M
codellama:34b-instruct-q6_K
codellama:34b-instruct-q8_0
codellama:34b-instruct-fp16
codellama:34b-python
codellama:34b-python-q2_K
codellama:34b-python-q3_K_S
codellama:34b-python-q3_K_M
codellama:34b-python-q3_K_L
codellama:34b-python-q4_0
codellama:34b-python-q4_1
codellama:34b-python-q4_K_S
codellama:34b-python-q4_K_M
codellama:34b-python-q5_0
codellama:34b-python-q5_1
codellama:34b-python-q5_K_S
codellama:34b-python-q5_K_M
codellama:34b-python-q6_K
codellama:34b-python-q8_0
codellama:34b-python-fp16
codellama:70b-code
codellama:70b-code-q2_K
codellama:70b-code-q3_K_S
codellama:70b-code-q3_K_M
codellama:70b-code-q3_K_L
codellama:70b-code-q4_0
codellama:70b-code-q4_1
codellama:70b-code-q4_K_S
codellama:70b-code-q4_K_M
codellama:70b-code-q5_0
codellama:70b-code-q5_1
codellama:70b-code-q5_K_S
codellama:70b-code-q5_K_M
codellama:70b-code-q6_K
codellama:70b-code-q8_0
codellama:70b-code-fp16
codellama:70b-instruct
codellama:70b-instruct-q2_K
codellama:70b-instruct-q3_K_S
codellama:70b-instruct-q3_K_M
codellama:70b-instruct-q3_K_L
codellama:70b-instruct-q4_0
codellama:70b-instruct-q4_1
codellama:70b-instruct-q4_K_S
codellama:70b-instruct-q4_K_M
codellama:70b-instruct-q5_0
codellama:70b-instruct-q5_1
codellama:70b-instruct-q5_K_S
codellama:70b-instruct-q5_K_M
codellama:70b-instruct-q6_K
codellama:70b-instruct-q8_0
codellama:70b-instruct-fp16
codellama:70b-python
codellama:70b-python-q2_K
codellama:70b-python-q3_K_S
codellama:70b-python-q3_K_M
codellama:70b-python-q3_K_L
codellama:70b-python-q4_0
codellama:70b-python-q4_1
codellama:70b-python-q4_K_S
codellama:70b-python-q4_K_M
codellama:70b-python-q5_0
codellama:70b-python-q5_1
codellama:70b-python-q5_K_S
codellama:70b-python-q5_K_M
codellama:70b-python-q6_K
codellama:70b-python-q8_0
codellama:70b-python-fp16
dolphin3:latest
dolphin3
dolphin3:8b
dolphin3:8b-llama3.1-q4_K_M
dolphin3:8b-llama3.1-q8_0
dolphin3:8b-llama3.1-fp16
olmo2:latest
olmo2
olmo2:7b
olmo2:13b
olmo2:7b-1124-instruct-q4_K_M
olmo2:7b-1124-instruct-q8_0
olmo2:7b-1124-instruct-fp16
olmo2:13b-1124-instruct-q4_K_M
olmo2:13b-1124-instruct-q8_0
olmo2:13b-1124-instruct-fp16
llama3.2-vision:latest
llama3.2-vision
llama3.2-vision:11b
llama3.2-vision:90b
llama3.2-vision:11b-instruct-q4_K_M
llama3.2-vision:11b-instruct-q8_0
llama3.2-vision:11b-instruct-fp16
llama3.2-vision:90b-instruct-q4_K_M
llama3.2-vision:90b-instruct-q8_0
llama3.2-vision:90b-instruct-fp16
tinyllama:latest
tinyllama
tinyllama:chat
tinyllama:v0.6
tinyllama:v1
tinyllama:1.1b
tinyllama:1.1b-chat
tinyllama:1.1b-chat-v0.6-q2_K
tinyllama:1.1b-chat-v0.6-q3_K_S
tinyllama:1.1b-chat-v0.6-q3_K_M
tinyllama:1.1b-chat-v0.6-q3_K_L
tinyllama:1.1b-chat-v0.6-q4_0
tinyllama:1.1b-chat-v0.6-q4_1
tinyllama:1.1b-chat-v0.6-q4_K_S
tinyllama:1.1b-chat-v0.6-q4_K_M
tinyllama:1.1b-chat-v0.6-q5_0
tinyllama:1.1b-chat-v0.6-q5_1
tinyllama:1.1b-chat-v0.6-q5_K_S
tinyllama:1.1b-chat-v0.6-q5_K_M
tinyllama:1.1b-chat-v0.6-q6_K
tinyllama:1.1b-chat-v0.6-q8_0
tinyllama:1.1b-chat-v0.6-fp16
tinyllama:1.1b-chat-v1-q2_K
tinyllama:1.1b-chat-v1-q3_K_S
tinyllama:1.1b-chat-v1-q3_K_M
tinyllama:1.1b-chat-v1-q3_K_L
tinyllama:1.1b-chat-v1-q4_0
tinyllama:1.1b-chat-v1-q4_1
tinyllama:1.1b-chat-v1-q4_K_S
tinyllama:1.1b-chat-v1-q4_K_M
tinyllama:1.1b-chat-v1-q5_0
tinyllama:1.1b-chat-v1-q5_1
tinyllama:1.1b-chat-v1-q5_K_S
tinyllama:1.1b-chat-v1-q5_K_M
tinyllama:1.1b-chat-v1-q6_K
tinyllama:1.1b-chat-v1-q8_0
tinyllama:1.1b-chat-v1-fp16
mistral-nemo:latest
mistral-nemo
mistral-nemo:12b
mistral-nemo:12b-instruct-2407-q2_K
mistral-nemo:12b-instruct-2407-q3_K_S
mistral-nemo:12b-instruct-2407-q3_K_M
mistral-nemo:12b-instruct-2407-q3_K_L
mistral-nemo:12b-instruct-2407-q4_0
mistral-nemo:12b-instruct-2407-q4_1
mistral-nemo:12b-instruct-2407-q4_K_S
mistral-nemo:12b-instruct-2407-q4_K_M
mistral-nemo:12b-instruct-2407-q5_0
mistral-nemo:12b-instruct-2407-q5_1
mistral-nemo:12b-instruct-2407-q5_K_S
mistral-nemo:12b-instruct-2407-q5_K_M
mistral-nemo:12b-instruct-2407-q6_K
mistral-nemo:12b-instruct-2407-q8_0
mistral-nemo:12b-instruct-2407-fp16
deepseek-v3:latest
deepseek-v3
deepseek-v3:671b
deepseek-v3:671b-q4_K_M
deepseek-v3:671b-q8_0
deepseek-v3:671b-fp16
bge-m3:latest
bge-m3
bge-m3:567m
bge-m3:567m-fp16
llama3.3:latest
llama3.3
llama3.3:70b
llama3.3:70b-instruct-q2_K
llama3.3:70b-instruct-q3_K_S
llama3.3:70b-instruct-q3_K_M
llama3.3:70b-instruct-q4_0
llama3.3:70b-instruct-q4_K_S
llama3.3:70b-instruct-q4_K_M
llama3.3:70b-instruct-q5_0
llama3.3:70b-instruct-q5_1
llama3.3:70b-instruct-q5_K_M
llama3.3:70b-instruct-q6_K
llama3.3:70b-instruct-q8_0
llama3.3:70b-instruct-fp16
deepseek-coder:latest
deepseek-coder
deepseek-coder:base
deepseek-coder:instruct
deepseek-coder:1.3b
deepseek-coder:6.7b
deepseek-coder:33b
deepseek-coder:1.3b-base
deepseek-coder:1.3b-base-q2_K
deepseek-coder:1.3b-base-q3_K_S
deepseek-coder:1.3b-base-q3_K_M
deepseek-coder:1.3b-base-q3_K_L
deepseek-coder:1.3b-base-q4_0
deepseek-coder:1.3b-base-q4_1
deepseek-coder:1.3b-base-q4_K_S
deepseek-coder:1.3b-base-q4_K_M
deepseek-coder:1.3b-base-q5_0
deepseek-coder:1.3b-base-q5_1
deepseek-coder:1.3b-base-q5_K_S
deepseek-coder:1.3b-base-q5_K_M
deepseek-coder:1.3b-base-q6_K
deepseek-coder:1.3b-base-q8_0
deepseek-coder:1.3b-base-fp16
deepseek-coder:1.3b-instruct
deepseek-coder:1.3b-instruct-q2_K
deepseek-coder:1.3b-instruct-q3_K_S
deepseek-coder:1.3b-instruct-q3_K_M
deepseek-coder:1.3b-instruct-q3_K_L
deepseek-coder:1.3b-instruct-q4_0
deepseek-coder:1.3b-instruct-q4_1
deepseek-coder:1.3b-instruct-q4_K_S
deepseek-coder:1.3b-instruct-q4_K_M
deepseek-coder:1.3b-instruct-q5_0
deepseek-coder:1.3b-instruct-q5_1
deepseek-coder:1.3b-instruct-q5_K_S
deepseek-coder:1.3b-instruct-q5_K_M
deepseek-coder:1.3b-instruct-q6_K
deepseek-coder:1.3b-instruct-q8_0
deepseek-coder:1.3b-instruct-fp16
deepseek-coder:6.7b-base
deepseek-coder:6.7b-base-q2_K
deepseek-coder:6.7b-base-q3_K_S
deepseek-coder:6.7b-base-q3_K_M
deepseek-coder:6.7b-base-q3_K_L
deepseek-coder:6.7b-base-q4_0
deepseek-coder:6.7b-base-q4_1
deepseek-coder:6.7b-base-q4_K_S
deepseek-coder:6.7b-base-q4_K_M
deepseek-coder:6.7b-base-q5_0
deepseek-coder:6.7b-base-q5_1
deepseek-coder:6.7b-base-q5_K_S
deepseek-coder:6.7b-base-q5_K_M
deepseek-coder:6.7b-base-q6_K
deepseek-coder:6.7b-base-q8_0
deepseek-coder:6.7b-base-fp16
deepseek-coder:6.7b-instruct
deepseek-coder:6.7b-instruct-q2_K
deepseek-coder:6.7b-instruct-q3_K_S
deepseek-coder:6.7b-instruct-q3_K_M
deepseek-coder:6.7b-instruct-q3_K_L
deepseek-coder:6.7b-instruct-q4_0
deepseek-coder:6.7b-instruct-q4_1
deepseek-coder:6.7b-instruct-q4_K_S
deepseek-coder:6.7b-instruct-q4_K_M
deepseek-coder:6.7b-instruct-q5_0
deepseek-coder:6.7b-instruct-q5_1
deepseek-coder:6.7b-instruct-q5_K_S
deepseek-coder:6.7b-instruct-q5_K_M
deepseek-coder:6.7b-instruct-q6_K
deepseek-coder:6.7b-instruct-q8_0
deepseek-coder:6.7b-instruct-fp16
deepseek-coder:33b-base
deepseek-coder:33b-base-q2_K
deepseek-coder:33b-base-q3_K_S
deepseek-coder:33b-base-q3_K_M
deepseek-coder:33b-base-q3_K_L
deepseek-coder:33b-base-q4_0
deepseek-coder:33b-base-q4_1
deepseek-coder:33b-base-q4_K_S
deepseek-coder:33b-base-q4_K_M
deepseek-coder:33b-base-q5_0
deepseek-coder:33b-base-q5_1
deepseek-coder:33b-base-q5_K_S
deepseek-coder:33b-base-q5_K_M
deepseek-coder:33b-base-q6_K
deepseek-coder:33b-base-q8_0
deepseek-coder:33b-base-fp16
deepseek-coder:33b-instruct
deepseek-coder:33b-instruct-q2_K
deepseek-coder:33b-instruct-q3_K_S
deepseek-coder:33b-instruct-q3_K_M
deepseek-coder:33b-instruct-q3_K_L
deepseek-coder:33b-instruct-q4_0
deepseek-coder:33b-instruct-q4_1
deepseek-coder:33b-instruct-q4_K_S
deepseek-coder:33b-instruct-q4_K_M
deepseek-coder:33b-instruct-q5_0
deepseek-coder:33b-instruct-q5_1
deepseek-coder:33b-instruct-q5_K_S
deepseek-coder:33b-instruct-q5_K_M
deepseek-coder:33b-instruct-q6_K
deepseek-coder:33b-instruct-q8_0
deepseek-coder:33b-instruct-fp16
smollm2:latest
smollm2
smollm2:135m
smollm2:360m
smollm2:1.7b
smollm2:135m-instruct-q2_K
smollm2:135m-instruct-q3_K_S
smollm2:135m-instruct-q3_K_M
smollm2:135m-instruct-q3_K_L
smollm2:135m-instruct-q4_0
smollm2:135m-instruct-q4_1
smollm2:135m-instruct-q4_K_S
smollm2:135m-instruct-q4_K_M
smollm2:135m-instruct-q5_0
smollm2:135m-instruct-q5_1
smollm2:135m-instruct-q5_K_S
smollm2:135m-instruct-q5_K_M
smollm2:135m-instruct-q6_K
smollm2:135m-instruct-q8_0
smollm2:135m-instruct-fp16
smollm2:360m-instruct-q2_K
smollm2:360m-instruct-q3_K_S
smollm2:360m-instruct-q3_K_M
smollm2:360m-instruct-q3_K_L
smollm2:360m-instruct-q4_0
smollm2:360m-instruct-q4_1
smollm2:360m-instruct-q4_K_S
smollm2:360m-instruct-q4_K_M
smollm2:360m-instruct-q5_0
smollm2:360m-instruct-q5_1
smollm2:360m-instruct-q5_K_S
smollm2:360m-instruct-q5_K_M
smollm2:360m-instruct-q6_K
smollm2:360m-instruct-q8_0
smollm2:360m-instruct-fp16
smollm2:1.7b-instruct-q2_K
smollm2:1.7b-instruct-q3_K_S
smollm2:1.7b-instruct-q3_K_M
smollm2:1.7b-instruct-q3_K_L
smollm2:1.7b-instruct-q4_0
smollm2:1.7b-instruct-q4_1
smollm2:1.7b-instruct-q4_K_S
smollm2:1.7b-instruct-q4_K_M
smollm2:1.7b-instruct-q5_0
smollm2:1.7b-instruct-q5_1
smollm2:1.7b-instruct-q5_K_S
smollm2:1.7b-instruct-q5_K_M
smollm2:1.7b-instruct-q6_K
smollm2:1.7b-instruct-q8_0
smollm2:1.7b-instruct-fp16
mistral-small:latest
mistral-small
mistral-small:22b
mistral-small:24b
mistral-small:22b-instruct-2409-q2_K
mistral-small:22b-instruct-2409-q3_K_S
mistral-small:22b-instruct-2409-q3_K_M
mistral-small:22b-instruct-2409-q3_K_L
mistral-small:22b-instruct-2409-q4_0
mistral-small:22b-instruct-2409-q4_1
mistral-small:22b-instruct-2409-q4_K_S
mistral-small:22b-instruct-2409-q4_K_M
mistral-small:22b-instruct-2409-q5_0
mistral-small:22b-instruct-2409-q5_1
mistral-small:22b-instruct-2409-q5_K_S
mistral-small:22b-instruct-2409-q5_K_M
mistral-small:22b-instruct-2409-q6_K
mistral-small:22b-instruct-2409-q8_0
mistral-small:22b-instruct-2409-fp16
mistral-small:24b-instruct-2501-q4_K_M
mistral-small:24b-instruct-2501-q8_0
mistral-small:24b-instruct-2501-fp16
all-minilm:latest
all-minilm
all-minilm:l12
all-minilm:l6
all-minilm:v2
all-minilm:22m
all-minilm:33m
all-minilm:22m-l6-v2-fp16
all-minilm:33m-l12-v2-fp16
all-minilm:l12-v2
all-minilm:l6-v2
llava-llama3:latest
llava-llama3
llava-llama3:8b
llava-llama3:8b-v1.1-q4_0
llava-llama3:8b-v1.1-fp16
qwq:latest
qwq
qwq:32b
qwq:32b-preview-q4_K_M
qwq:32b-preview-q8_0
qwq:32b-preview-fp16
qwq:32b-q4_K_M
qwq:32b-q8_0
qwq:32b-fp16
codegemma:latest
codegemma
codegemma:code
codegemma:instruct
codegemma:2b
codegemma:7b
codegemma:2b-code
codegemma:2b-code-q2_K
codegemma:2b-code-v1.1-q2_K
codegemma:2b-code-q3_K_S
codegemma:2b-code-v1.1-q3_K_S
codegemma:2b-code-q3_K_M
codegemma:2b-code-v1.1-q3_K_M
codegemma:2b-code-q3_K_L
codegemma:2b-code-v1.1-q3_K_L
codegemma:2b-code-q4_0
codegemma:2b-code-v1.1-q4_0
codegemma:2b-code-q4_1
codegemma:2b-code-v1.1-q4_1
codegemma:2b-code-q4_K_S
codegemma:2b-code-v1.1-q4_K_S
codegemma:2b-code-q4_K_M
codegemma:2b-code-v1.1-q4_K_M
codegemma:2b-code-q5_0
codegemma:2b-code-v1.1-q5_0
codegemma:2b-code-q5_1
codegemma:2b-code-v1.1-q5_1
codegemma:2b-code-q5_K_S
codegemma:2b-code-v1.1-q5_K_S
codegemma:2b-code-q5_K_M
codegemma:2b-code-v1.1-q5_K_M
codegemma:2b-code-q6_K
codegemma:2b-code-v1.1-q6_K
codegemma:2b-code-q8_0
codegemma:2b-code-v1.1-q8_0
codegemma:2b-code-fp16
codegemma:2b-code-v1.1-fp16
codegemma:2b-v1.1
codegemma:7b-code
codegemma:7b-code-q2_K
codegemma:7b-code-q3_K_S
codegemma:7b-code-q3_K_M
codegemma:7b-code-q3_K_L
codegemma:7b-code-q4_0
codegemma:7b-code-q4_1
codegemma:7b-code-q4_K_S
codegemma:7b-code-q4_K_M
codegemma:7b-code-q5_0
codegemma:7b-code-q5_1
codegemma:7b-code-q5_K_S
codegemma:7b-code-q5_K_M
codegemma:7b-code-q6_K
codegemma:7b-code-q8_0
codegemma:7b-code-fp16
codegemma:7b-instruct
codegemma:7b-instruct-q2_K
codegemma:7b-instruct-v1.1-q2_K
codegemma:7b-instruct-q3_K_S
codegemma:7b-instruct-v1.1-q3_K_S
codegemma:7b-instruct-q3_K_M
codegemma:7b-instruct-v1.1-q3_K_M
codegemma:7b-instruct-q3_K_L
codegemma:7b-instruct-v1.1-q3_K_L
codegemma:7b-instruct-q4_0
codegemma:7b-instruct-v1.1-q4_0
codegemma:7b-instruct-q4_1
codegemma:7b-instruct-v1.1-q4_1
codegemma:7b-instruct-q4_K_S
codegemma:7b-instruct-v1.1-q4_K_S
codegemma:7b-instruct-q4_K_M
codegemma:7b-instruct-v1.1-q4_K_M
codegemma:7b-instruct-q5_0
codegemma:7b-instruct-v1.1-q5_0
codegemma:7b-instruct-q5_1
codegemma:7b-instruct-v1.1-q5_1
codegemma:7b-instruct-q5_K_S
codegemma:7b-instruct-v1.1-q5_K_S
codegemma:7b-instruct-q5_K_M
codegemma:7b-instruct-v1.1-q5_K_M
codegemma:7b-instruct-q6_K
codegemma:7b-instruct-v1.1-q6_K
codegemma:7b-instruct-q8_0
codegemma:7b-instruct-v1.1-q8_0
codegemma:7b-instruct-fp16
codegemma:7b-instruct-v1.1-fp16
codegemma:7b-v1.1
starcoder2:latest
starcoder2
starcoder2:instruct
starcoder2:3b
starcoder2:7b
starcoder2:15b
starcoder2:3b-q2_K
starcoder2:3b-q3_K_S
starcoder2:3b-q3_K_M
starcoder2:3b-q3_K_L
starcoder2:3b-q4_0
starcoder2:3b-q4_1
starcoder2:3b-q4_K_S
starcoder2:3b-q4_K_M
starcoder2:3b-q5_0
starcoder2:3b-q5_1
starcoder2:3b-q5_K_S
starcoder2:3b-q5_K_M
starcoder2:3b-q6_K
starcoder2:3b-q8_0
starcoder2:3b-fp16
starcoder2:7b-q2_K
starcoder2:7b-q3_K_S
starcoder2:7b-q3_K_M
starcoder2:7b-q3_K_L
starcoder2:7b-q4_0
starcoder2:7b-q4_1
starcoder2:7b-q4_K_S
starcoder2:7b-q4_K_M
starcoder2:7b-q5_0
starcoder2:7b-q5_1
starcoder2:7b-q5_K_S
starcoder2:7b-q5_K_M
starcoder2:7b-q6_K
starcoder2:7b-q8_0
starcoder2:7b-fp16
starcoder2:15b-instruct
starcoder2:15b-instruct-v0.1-q2_K
starcoder2:15b-instruct-v0.1-q3_K_S
starcoder2:15b-instruct-v0.1-q3_K_M
starcoder2:15b-instruct-v0.1-q3_K_L
starcoder2:15b-instruct-q4_0
starcoder2:15b-instruct-v0.1-q4_0
starcoder2:15b-instruct-v0.1-q4_1
starcoder2:15b-instruct-v0.1-q4_K_S
starcoder2:15b-instruct-v0.1-q4_K_M
starcoder2:15b-instruct-v0.1-q5_0
starcoder2:15b-instruct-v0.1-q5_1
starcoder2:15b-instruct-v0.1-q5_K_S
starcoder2:15b-instruct-v0.1-q5_K_M
starcoder2:15b-instruct-v0.1-q6_K
starcoder2:15b-instruct-v0.1-q8_0
starcoder2:15b-instruct-v0.1-fp16
starcoder2:15b-q2_K
starcoder2:15b-q3_K_S
starcoder2:15b-q3_K_M
starcoder2:15b-q3_K_L
starcoder2:15b-q4_0
starcoder2:15b-q4_1
starcoder2:15b-q4_K_S
starcoder2:15b-q4_K_M
starcoder2:15b-q5_0
starcoder2:15b-q5_1
starcoder2:15b-q5_K_S
starcoder2:15b-q5_K_M
starcoder2:15b-q6_K
starcoder2:15b-q8_0
starcoder2:15b-fp16
falcon3:latest
falcon3
falcon3:1b
falcon3:3b
falcon3:7b
falcon3:10b
falcon3:1b-instruct-q4_K_M
falcon3:1b-instruct-q8_0
falcon3:1b-instruct-fp16
falcon3:3b-instruct-q4_K_M
falcon3:3b-instruct-q8_0
falcon3:3b-instruct-fp16
falcon3:7b-instruct-q4_K_M
falcon3:7b-instruct-q8_0
falcon3:7b-instruct-fp16
falcon3:10b-instruct-q4_K_M
falcon3:10b-instruct-q8_0
falcon3:10b-instruct-fp16
granite3.1-moe:latest
granite3.1-moe
granite3.1-moe:1b
granite3.1-moe:3b
granite3.1-moe:1b-instruct-q2_K
granite3.1-moe:1b-instruct-q3_K_S
granite3.1-moe:1b-instruct-q3_K_M
granite3.1-moe:1b-instruct-q3_K_L
granite3.1-moe:1b-instruct-q4_0
granite3.1-moe:1b-instruct-q4_1
granite3.1-moe:1b-instruct-q4_K_S
granite3.1-moe:1b-instruct-q4_K_M
granite3.1-moe:1b-instruct-q5_0
granite3.1-moe:1b-instruct-q5_1
granite3.1-moe:1b-instruct-q5_K_S
granite3.1-moe:1b-instruct-q5_K_M
granite3.1-moe:1b-instruct-q6_K
granite3.1-moe:1b-instruct-q8_0
granite3.1-moe:1b-instruct-fp16
granite3.1-moe:3b-instruct-q2_K
granite3.1-moe:3b-instruct-q3_K_S
granite3.1-moe:3b-instruct-q3_K_M
granite3.1-moe:3b-instruct-q3_K_L
granite3.1-moe:3b-instruct-q4_0
granite3.1-moe:3b-instruct-q4_1
granite3.1-moe:3b-instruct-q4_K_S
granite3.1-moe:3b-instruct-q4_K_M
granite3.1-moe:3b-instruct-q5_0
granite3.1-moe:3b-instruct-q5_1
granite3.1-moe:3b-instruct-q5_K_S
granite3.1-moe:3b-instruct-q5_K_M
granite3.1-moe:3b-instruct-q6_K
granite3.1-moe:3b-instruct-q8_0
granite3.1-moe:3b-instruct-fp16
mixtral:latest
mixtral
mixtral:instruct
mixtral:text
mixtral:v0.1
mixtral:8x7b
mixtral:8x22b
mixtral:8x7b-instruct-v0.1-q2_K
mixtral:8x7b-instruct-v0.1-q3_K_S
mixtral:8x7b-instruct-v0.1-q3_K_M
mixtral:8x7b-instruct-v0.1-q3_K_L
mixtral:8x7b-instruct-v0.1-q4_0
mixtral:8x7b-instruct-v0.1-q4_1
mixtral:8x7b-instruct-v0.1-q4_K_S
mixtral:8x7b-instruct-v0.1-q4_K_M
mixtral:8x7b-instruct-v0.1-q5_0
mixtral:8x7b-instruct-v0.1-q5_1
mixtral:8x7b-instruct-v0.1-q5_K_S
mixtral:8x7b-instruct-v0.1-q5_K_M
mixtral:8x7b-instruct-v0.1-q6_K
mixtral:8x7b-instruct-v0.1-q8_0
mixtral:8x7b-instruct-v0.1-fp16
mixtral:8x7b-text
mixtral:8x7b-text-v0.1-q2_K
mixtral:8x7b-text-v0.1-q3_K_S
mixtral:8x7b-text-v0.1-q3_K_M
mixtral:8x7b-text-v0.1-q3_K_L
mixtral:8x7b-text-v0.1-q4_0
mixtral:8x7b-text-v0.1-q4_1
mixtral:8x7b-text-v0.1-q4_K_S
mixtral:8x7b-text-v0.1-q4_K_M
mixtral:8x7b-text-v0.1-q5_0
mixtral:8x7b-text-v0.1-q5_1
mixtral:8x7b-text-v0.1-q5_K_S
mixtral:8x7b-text-v0.1-q5_K_M
mixtral:8x7b-text-v0.1-q6_K
mixtral:8x7b-text-v0.1-q8_0
mixtral:8x7b-text-v0.1-fp16
mixtral:8x22b-instruct
mixtral:8x22b-instruct-v0.1-q2_K
mixtral:8x22b-instruct-v0.1-q3_K_S
mixtral:8x22b-instruct-v0.1-q3_K_M
mixtral:8x22b-instruct-v0.1-q3_K_L
mixtral:8x22b-instruct-v0.1-q4_0
mixtral:8x22b-instruct-v0.1-q4_1
mixtral:8x22b-instruct-v0.1-q4_K_S
mixtral:8x22b-instruct-v0.1-q4_K_M
mixtral:8x22b-instruct-v0.1-q5_0
mixtral:8x22b-instruct-v0.1-q5_1
mixtral:8x22b-instruct-v0.1-q5_K_S
mixtral:8x22b-instruct-v0.1-q5_K_M
mixtral:8x22b-instruct-v0.1-q6_K
mixtral:8x22b-instruct-v0.1-q8_0
mixtral:8x22b-instruct-v0.1-fp16
mixtral:8x22b-text
mixtral:8x22b-text-v0.1-q2_K
mixtral:8x22b-text-v0.1-q3_K_S
mixtral:8x22b-text-v0.1-q3_K_M
mixtral:8x22b-text-v0.1-q3_K_L
mixtral:8x22b-text-v0.1-q4_0
mixtral:8x22b-text-v0.1-q4_1
mixtral:8x22b-text-v0.1-q4_K_S
mixtral:8x22b-text-v0.1-q4_K_M
mixtral:8x22b-text-v0.1-q5_0
mixtral:8x22b-text-v0.1-q5_1
mixtral:8x22b-text-v0.1-q5_K_S
mixtral:8x22b-text-v0.1-q5_K_M
mixtral:8x22b-text-v0.1-q6_K
mixtral:8x22b-text-v0.1-q8_0
mixtral:8x22b-text-v0.1-fp16
mixtral:v0.1-instruct
llama2-uncensored:latest
llama2-uncensored
llama2-uncensored:7b
llama2-uncensored:70b
llama2-uncensored:7b-chat
llama2-uncensored:7b-chat-q2_K
llama2-uncensored:7b-chat-q3_K_S
llama2-uncensored:7b-chat-q3_K_M
llama2-uncensored:7b-chat-q3_K_L
llama2-uncensored:7b-chat-q4_0
llama2-uncensored:7b-chat-q4_1
llama2-uncensored:7b-chat-q4_K_S
llama2-uncensored:7b-chat-q4_K_M
llama2-uncensored:7b-chat-q5_0
llama2-uncensored:7b-chat-q5_1
llama2-uncensored:7b-chat-q5_K_S
llama2-uncensored:7b-chat-q5_K_M
llama2-uncensored:7b-chat-q6_K
llama2-uncensored:7b-chat-q8_0
llama2-uncensored:7b-chat-fp16
llama2-uncensored:70b-chat
llama2-uncensored:70b-chat-q2_K
llama2-uncensored:70b-chat-q3_K_S
llama2-uncensored:70b-chat-q3_K_M
llama2-uncensored:70b-chat-q3_K_L
llama2-uncensored:70b-chat-q4_0
llama2-uncensored:70b-chat-q4_1
llama2-uncensored:70b-chat-q4_K_S
llama2-uncensored:70b-chat-q4_K_M
llama2-uncensored:70b-chat-q5_0
llama2-uncensored:70b-chat-q5_1
llama2-uncensored:70b-chat-q5_K_S
llama2-uncensored:70b-chat-q5_K_M
llama2-uncensored:70b-chat-q6_K
llama2-uncensored:70b-chat-q8_0
orca-mini:latest
orca-mini
orca-mini:3b
orca-mini:7b
orca-mini:13b
orca-mini:70b
orca-mini:3b-q4_0
orca-mini:3b-q4_1
orca-mini:3b-q5_0
orca-mini:3b-q5_1
orca-mini:3b-q8_0
orca-mini:3b-fp16
orca-mini:7b-v2-q2_K
orca-mini:7b-v2-q3_K_S
orca-mini:7b-v2-q3_K_M
orca-mini:7b-v2-q3_K_L
orca-mini:7b-v2-q4_0
orca-mini:7b-v2-q4_1
orca-mini:7b-v2-q4_K_S
orca-mini:7b-v2-q4_K_M
orca-mini:7b-v2-q5_0
orca-mini:7b-v2-q5_1
orca-mini:7b-v2-q5_K_S
orca-mini:7b-v2-q5_K_M
orca-mini:7b-v2-q6_K
orca-mini:7b-v2-q8_0
orca-mini:7b-v2-fp16
orca-mini:7b-v3
orca-mini:7b-v3-q2_K
orca-mini:7b-v3-q3_K_S
orca-mini:7b-v3-q3_K_M
orca-mini:7b-v3-q3_K_L
orca-mini:7b-v3-q4_0
orca-mini:7b-v3-q4_1
orca-mini:7b-v3-q4_K_S
orca-mini:7b-v3-q4_K_M
orca-mini:7b-v3-q5_0
orca-mini:7b-v3-q5_1
orca-mini:7b-v3-q5_K_S
orca-mini:7b-v3-q5_K_M
orca-mini:7b-v3-q6_K
orca-mini:7b-v3-q8_0
orca-mini:7b-v3-fp16
orca-mini:7b-q2_K
orca-mini:7b-q3_K_S
orca-mini:7b-q3_K_M
orca-mini:7b-q3_K_L
orca-mini:7b-q4_0
orca-mini:7b-q4_1
orca-mini:7b-q4_K_S
orca-mini:7b-q4_K_M
orca-mini:7b-q5_0
orca-mini:7b-q5_1
orca-mini:7b-q5_K_S
orca-mini:7b-q5_K_M
orca-mini:7b-q6_K
orca-mini:7b-q8_0
orca-mini:7b-fp16
orca-mini:13b-v2-q2_K
orca-mini:13b-v2-q3_K_S
orca-mini:13b-v2-q3_K_M
orca-mini:13b-v2-q3_K_L
orca-mini:13b-v2-q4_0
orca-mini:13b-v2-q4_1
orca-mini:13b-v2-q4_K_S
orca-mini:13b-v2-q4_K_M
orca-mini:13b-v2-q5_0
orca-mini:13b-v2-q5_1
orca-mini:13b-v2-q5_K_S
orca-mini:13b-v2-q5_K_M
orca-mini:13b-v2-q6_K
orca-mini:13b-v2-q8_0
orca-mini:13b-v2-fp16
orca-mini:13b-v3
orca-mini:13b-v3-q2_K
orca-mini:13b-v3-q3_K_S
orca-mini:13b-v3-q3_K_M
orca-mini:13b-v3-q3_K_L
orca-mini:13b-v3-q4_0
orca-mini:13b-v3-q4_1
orca-mini:13b-v3-q4_K_S
orca-mini:13b-v3-q4_K_M
orca-mini:13b-v3-q5_0
orca-mini:13b-v3-q5_1
orca-mini:13b-v3-q5_K_S
orca-mini:13b-v3-q5_K_M
orca-mini:13b-v3-q6_K
orca-mini:13b-v3-q8_0
orca-mini:13b-v3-fp16
orca-mini:13b-q2_K
orca-mini:13b-q3_K_S
orca-mini:13b-q3_K_M
orca-mini:13b-q3_K_L
orca-mini:13b-q4_0
orca-mini:13b-q4_1
orca-mini:13b-q4_K_S
orca-mini:13b-q4_K_M
orca-mini:13b-q5_0
orca-mini:13b-q5_1
orca-mini:13b-q5_K_S
orca-mini:13b-q5_K_M
orca-mini:13b-q6_K
orca-mini:13b-q8_0
orca-mini:13b-fp16
orca-mini:70b-v3
orca-mini:70b-v3-q2_K
orca-mini:70b-v3-q3_K_S
orca-mini:70b-v3-q3_K_M
orca-mini:70b-v3-q3_K_L
orca-mini:70b-v3-q4_0
orca-mini:70b-v3-q4_1
orca-mini:70b-v3-q4_K_S
orca-mini:70b-v3-q4_K_M
orca-mini:70b-v3-q5_0
orca-mini:70b-v3-q5_1
orca-mini:70b-v3-q5_K_S
orca-mini:70b-v3-q5_K_M
orca-mini:70b-v3-q6_K
orca-mini:70b-v3-q8_0
orca-mini:70b-v3-fp16
snowflake-arctic-embed:latest
snowflake-arctic-embed
snowflake-arctic-embed:l
snowflake-arctic-embed:m
snowflake-arctic-embed:s
snowflake-arctic-embed:xs
snowflake-arctic-embed:22m
snowflake-arctic-embed:33m
snowflake-arctic-embed:110m
snowflake-arctic-embed:137m
snowflake-arctic-embed:335m
snowflake-arctic-embed:22m-xs-fp16
snowflake-arctic-embed:33m-s-fp16
snowflake-arctic-embed:110m-m-fp16
snowflake-arctic-embed:137m-m-long-fp16
snowflake-arctic-embed:335m-l-fp16
snowflake-arctic-embed:m-long
deepseek-coder-v2:latest
deepseek-coder-v2
deepseek-coder-v2:lite
deepseek-coder-v2:16b
deepseek-coder-v2:236b
deepseek-coder-v2:16b-lite-base-q2_K
deepseek-coder-v2:16b-lite-base-q3_K_S
deepseek-coder-v2:16b-lite-base-q3_K_M
deepseek-coder-v2:16b-lite-base-q3_K_L
deepseek-coder-v2:16b-lite-base-q4_0
deepseek-coder-v2:16b-lite-base-q4_1
deepseek-coder-v2:16b-lite-base-q4_K_S
deepseek-coder-v2:16b-lite-base-q4_K_M
deepseek-coder-v2:16b-lite-base-q5_0
deepseek-coder-v2:16b-lite-base-q5_1
deepseek-coder-v2:16b-lite-base-q5_K_S
deepseek-coder-v2:16b-lite-base-q5_K_M
deepseek-coder-v2:16b-lite-base-q6_K
deepseek-coder-v2:16b-lite-base-q8_0
deepseek-coder-v2:16b-lite-base-fp16
deepseek-coder-v2:16b-lite-instruct-q2_K
deepseek-coder-v2:16b-lite-instruct-q3_K_S
deepseek-coder-v2:16b-lite-instruct-q3_K_M
deepseek-coder-v2:16b-lite-instruct-q3_K_L
deepseek-coder-v2:16b-lite-instruct-q4_0
deepseek-coder-v2:16b-lite-instruct-q4_1
deepseek-coder-v2:16b-lite-instruct-q4_K_S
deepseek-coder-v2:16b-lite-instruct-q4_K_M
deepseek-coder-v2:16b-lite-instruct-q5_0
deepseek-coder-v2:16b-lite-instruct-q5_1
deepseek-coder-v2:16b-lite-instruct-q5_K_S
deepseek-coder-v2:16b-lite-instruct-q5_K_M
deepseek-coder-v2:16b-lite-instruct-q6_K
deepseek-coder-v2:16b-lite-instruct-q8_0
deepseek-coder-v2:16b-lite-instruct-fp16
deepseek-coder-v2:236b-base-q2_K
deepseek-coder-v2:236b-base-q3_K_S
deepseek-coder-v2:236b-base-q3_K_M
deepseek-coder-v2:236b-base-q3_K_L
deepseek-coder-v2:236b-base-q4_0
deepseek-coder-v2:236b-base-q4_1
deepseek-coder-v2:236b-base-q4_K_S
deepseek-coder-v2:236b-base-q4_K_M
deepseek-coder-v2:236b-base-q5_0
deepseek-coder-v2:236b-base-q5_1
deepseek-coder-v2:236b-base-q5_K_S
deepseek-coder-v2:236b-base-q5_K_M
deepseek-coder-v2:236b-base-q6_K
deepseek-coder-v2:236b-base-q8_0
deepseek-coder-v2:236b-base-fp16
deepseek-coder-v2:236b-instruct-q2_K
deepseek-coder-v2:236b-instruct-q3_K_S
deepseek-coder-v2:236b-instruct-q3_K_M
deepseek-coder-v2:236b-instruct-q3_K_L
deepseek-coder-v2:236b-instruct-q4_0
deepseek-coder-v2:236b-instruct-q4_1
deepseek-coder-v2:236b-instruct-q4_K_S
deepseek-coder-v2:236b-instruct-q4_K_M
deepseek-coder-v2:236b-instruct-q5_0
deepseek-coder-v2:236b-instruct-q5_1
deepseek-coder-v2:236b-instruct-q5_K_S
deepseek-coder-v2:236b-instruct-q5_K_M
deepseek-coder-v2:236b-instruct-q6_K
deepseek-coder-v2:236b-instruct-q8_0
deepseek-coder-v2:236b-instruct-fp16
qwen2.5vl:latest
qwen2.5vl
qwen2.5vl:3b
qwen2.5vl:7b
qwen2.5vl:32b
qwen2.5vl:72b
qwen2.5vl:3b-q4_K_M
qwen2.5vl:3b-q8_0
qwen2.5vl:3b-fp16
qwen2.5vl:7b-q4_K_M
qwen2.5vl:7b-q8_0
qwen2.5vl:7b-fp16
qwen2.5vl:32b-q4_K_M
qwen2.5vl:32b-q8_0
qwen2.5vl:32b-fp16
qwen2.5vl:72b-q4_K_M
qwen2.5vl:72b-q8_0
qwen2.5vl:72b-fp16
cogito:latest
cogito
cogito:3b
cogito:8b
cogito:14b
cogito:32b
cogito:70b
cogito:3b-v1-preview-llama-q4_K_M
cogito:3b-v1-preview-llama-q8_0
cogito:3b-v1-preview-llama-fp16
cogito:8b-v1-preview-llama-q4_K_M
cogito:8b-v1-preview-llama-q8_0
cogito:14b-v1-preview-qwen-q4_K_M
cogito:14b-v1-preview-qwen-q8_0
cogito:14b-v1-preview-qwen-fp16
cogito:32b-v1-preview-qwen-q4_K_M
cogito:32b-v1-preview-qwen-q8_0
cogito:32b-v1-preview-qwen-fp16
cogito:70b-v1-preview-llama-q4_K_M
cogito:70b-v1-preview-llama-q8_0
cogito:70b-v1-preview-llama-fp16
gemma3n:latest
gemma3n
gemma3n:e2b
gemma3n:e4b
gemma3n:e2b-it-q4_K_M
gemma3n:e2b-it-q8_0
gemma3n:e2b-it-fp16
gemma3n:e4b-it-q4_K_M
gemma3n:e4b-it-q8_0
gemma3n:e4b-it-fp16
llama4:latest
llama4
llama4:maverick
llama4:scout
llama4:16x17b
llama4:128x17b
llama4:17b-maverick-128e-instruct-q4_K_M
llama4:17b-maverick-128e-instruct-q8_0
llama4:17b-maverick-128e-instruct-fp16
llama4:17b-scout-16e-instruct-q4_K_M
llama4:17b-scout-16e-instruct-q8_0
llama4:17b-scout-16e-instruct-fp16
mistral-small3.2:latest
mistral-small3.2
mistral-small3.2:24b
mistral-small3.2:24b-instruct-2506-q4_K_M
mistral-small3.2:24b-instruct-2506-q8_0
mistral-small3.2:24b-instruct-2506-fp16
deepscaler:latest
deepscaler
deepscaler:1.5b
deepscaler:1.5b-preview-q4_K_M
deepscaler:1.5b-preview-q8_0
deepscaler:1.5b-preview-fp16
phi4-reasoning:latest
phi4-reasoning
phi4-reasoning:plus
phi4-reasoning:14b
phi4-reasoning:14b-plus-q4_K_M
phi4-reasoning:14b-plus-q8_0
phi4-reasoning:14b-plus-fp16
phi4-reasoning:14b-q4_K_M
phi4-reasoning:14b-q8_0
phi4-reasoning:14b-fp16
dolphin-phi:latest
dolphin-phi
dolphin-phi:2.7b
dolphin-phi:2.7b-v2.6
dolphin-phi:2.7b-v2.6-q2_K
dolphin-phi:2.7b-v2.6-q3_K_S
dolphin-phi:2.7b-v2.6-q3_K_M
dolphin-phi:2.7b-v2.6-q3_K_L
dolphin-phi:2.7b-v2.6-q4_0
dolphin-phi:2.7b-v2.6-q4_K_S
dolphin-phi:2.7b-v2.6-q4_K_M
dolphin-phi:2.7b-v2.6-q5_0
dolphin-phi:2.7b-v2.6-q5_K_S
dolphin-phi:2.7b-v2.6-q5_K_M
dolphin-phi:2.7b-v2.6-q6_K
dolphin-phi:2.7b-v2.6-q8_0
magistral:latest
magistral
magistral:24b
magistral:24b-small-2506-q4_K_M
magistral:24b-small-2506-q8_0
magistral:24b-small-2506-fp16
phi:latest
phi
phi:chat
phi:2.7b
phi:2.7b-chat-v2-q2_K
phi:2.7b-chat-v2-q3_K_S
phi:2.7b-chat-v2-q3_K_M
phi:2.7b-chat-v2-q3_K_L
phi:2.7b-chat-v2-q4_0
phi:2.7b-chat-v2-q4_1
phi:2.7b-chat-v2-q4_K_S
phi:2.7b-chat-v2-q4_K_M
phi:2.7b-chat-v2-q5_0
phi:2.7b-chat-v2-q5_1
phi:2.7b-chat-v2-q5_K_S
phi:2.7b-chat-v2-q5_K_M
phi:2.7b-chat-v2-q6_K
phi:2.7b-chat-v2-q8_0
phi:2.7b-chat-v2-fp16
granite3.3:latest
granite3.3
granite3.3:2b
granite3.3:8b
dolphin-mixtral:latest
dolphin-mixtral
dolphin-mixtral:v2.5
dolphin-mixtral:v2.6
dolphin-mixtral:v2.7
dolphin-mixtral:8x7b
dolphin-mixtral:8x22b
dolphin-mixtral:8x7b-v2.5
dolphin-mixtral:8x7b-v2.5-q2_K
dolphin-mixtral:8x7b-v2.5-q3_K_S
dolphin-mixtral:8x7b-v2.5-q3_K_M
dolphin-mixtral:8x7b-v2.5-q3_K_L
dolphin-mixtral:8x7b-v2.5-q4_0
dolphin-mixtral:8x7b-v2.5-q4_1
dolphin-mixtral:8x7b-v2.5-q4_K_S
dolphin-mixtral:8x7b-v2.5-q4_K_M
dolphin-mixtral:8x7b-v2.5-q5_0
dolphin-mixtral:8x7b-v2.5-q5_1
dolphin-mixtral:8x7b-v2.5-q5_K_S
dolphin-mixtral:8x7b-v2.5-q5_K_M
dolphin-mixtral:8x7b-v2.5-q6_K
dolphin-mixtral:8x7b-v2.5-q8_0
dolphin-mixtral:8x7b-v2.5-fp16
dolphin-mixtral:8x7b-v2.6
dolphin-mixtral:8x7b-v2.6-q2_K
dolphin-mixtral:8x7b-v2.6-q3_K_S
dolphin-mixtral:8x7b-v2.6-q3_K_M
dolphin-mixtral:8x7b-v2.6-q3_K_L
dolphin-mixtral:8x7b-v2.6-q4_0
dolphin-mixtral:8x7b-v2.6-q4_1
dolphin-mixtral:8x7b-v2.6-q4_K_S
dolphin-mixtral:8x7b-v2.6-q4_K_M
dolphin-mixtral:8x7b-v2.6-q5_0
dolphin-mixtral:8x7b-v2.6-q5_1
dolphin-mixtral:8x7b-v2.6-q5_K_S
dolphin-mixtral:8x7b-v2.6-q5_K_M
dolphin-mixtral:8x7b-v2.6-q6_K
dolphin-mixtral:8x7b-v2.6-q8_0
dolphin-mixtral:8x7b-v2.6-fp16
dolphin-mixtral:8x7b-v2.7
dolphin-mixtral:8x7b-v2.7-q2_K
dolphin-mixtral:8x7b-v2.7-q3_K_S
dolphin-mixtral:8x7b-v2.7-q3_K_M
dolphin-mixtral:8x7b-v2.7-q3_K_L
dolphin-mixtral:8x7b-v2.7-q4_0
dolphin-mixtral:8x7b-v2.7-q4_1
dolphin-mixtral:8x7b-v2.7-q4_K_S
dolphin-mixtral:8x7b-v2.7-q4_K_M
dolphin-mixtral:8x7b-v2.7-q5_0
dolphin-mixtral:8x7b-v2.7-q5_1
dolphin-mixtral:8x7b-v2.7-q5_K_S
dolphin-mixtral:8x7b-v2.7-q5_K_M
dolphin-mixtral:8x7b-v2.7-q6_K
dolphin-mixtral:8x7b-v2.7-q8_0
dolphin-mixtral:8x7b-v2.7-fp16
dolphin-mixtral:8x22b-v2.9
dolphin-mixtral:8x22b-v2.9-q2_K
dolphin-mixtral:8x22b-v2.9-q3_K_S
dolphin-mixtral:8x22b-v2.9-q3_K_M
dolphin-mixtral:8x22b-v2.9-q3_K_L
dolphin-mixtral:8x22b-v2.9-q4_0
dolphin-mixtral:8x22b-v2.9-q4_1
dolphin-mixtral:8x22b-v2.9-q4_K_S
dolphin-mixtral:8x22b-v2.9-q4_K_M
dolphin-mixtral:8x22b-v2.9-q5_0
dolphin-mixtral:8x22b-v2.9-q5_1
dolphin-mixtral:8x22b-v2.9-q5_K_S
dolphin-mixtral:8x22b-v2.9-q5_K_M
dolphin-mixtral:8x22b-v2.9-q6_K
dolphin-mixtral:8x22b-v2.9-q8_0
dolphin-mixtral:8x22b-v2.9-fp16
phi4-mini:latest
phi4-mini
phi4-mini:3.8b
phi4-mini:3.8b-q4_K_M
phi4-mini:3.8b-q8_0
phi4-mini:3.8b-fp16
dolphin-llama3:latest
dolphin-llama3
dolphin-llama3:v2.9
dolphin-llama3:8b
dolphin-llama3:70b
dolphin-llama3:8b-256k
dolphin-llama3:8b-256k-v2.9
dolphin-llama3:8b-256k-v2.9-q2_K
dolphin-llama3:8b-256k-v2.9-q3_K_S
dolphin-llama3:8b-256k-v2.9-q3_K_M
dolphin-llama3:8b-256k-v2.9-q3_K_L
dolphin-llama3:8b-256k-v2.9-q4_0
dolphin-llama3:8b-256k-v2.9-q4_1
dolphin-llama3:8b-256k-v2.9-q4_K_S
dolphin-llama3:8b-256k-v2.9-q4_K_M
dolphin-llama3:8b-256k-v2.9-q5_0
dolphin-llama3:8b-256k-v2.9-q5_1
dolphin-llama3:8b-256k-v2.9-q5_K_S
dolphin-llama3:8b-256k-v2.9-q5_K_M
dolphin-llama3:8b-256k-v2.9-q6_K
dolphin-llama3:8b-256k-v2.9-q8_0
dolphin-llama3:8b-256k-v2.9-fp16
dolphin-llama3:8b-v2.9
dolphin-llama3:8b-v2.9-q2_K
dolphin-llama3:8b-v2.9-q3_K_S
dolphin-llama3:8b-v2.9-q3_K_M
dolphin-llama3:8b-v2.9-q3_K_L
dolphin-llama3:8b-v2.9-q4_0
dolphin-llama3:8b-v2.9-q4_1
dolphin-llama3:8b-v2.9-q4_K_S
dolphin-llama3:8b-v2.9-q4_K_M
dolphin-llama3:8b-v2.9-q5_0
dolphin-llama3:8b-v2.9-q5_1
dolphin-llama3:8b-v2.9-q5_K_S
dolphin-llama3:8b-v2.9-q5_K_M
dolphin-llama3:8b-v2.9-q6_K
dolphin-llama3:8b-v2.9-q8_0
dolphin-llama3:8b-v2.9-fp16
dolphin-llama3:70b-v2.9
dolphin-llama3:70b-v2.9-q2_K
dolphin-llama3:70b-v2.9-q3_K_S
dolphin-llama3:70b-v2.9-q3_K_M
dolphin-llama3:70b-v2.9-q3_K_L
dolphin-llama3:70b-v2.9-q4_0
dolphin-llama3:70b-v2.9-q4_1
dolphin-llama3:70b-v2.9-q4_K_S
dolphin-llama3:70b-v2.9-q4_K_M
dolphin-llama3:70b-v2.9-q5_0
dolphin-llama3:70b-v2.9-q5_1
dolphin-llama3:70b-v2.9-q5_K_S
dolphin-llama3:70b-v2.9-q5_K_M
dolphin-llama3:70b-v2.9-q6_K
dolphin-llama3:70b-v2.9-q8_0
dolphin-llama3:70b-v2.9-fp16
openthinker:latest
openthinker
openthinker:7b
openthinker:32b
openthinker:7b-v2-q4_K_M
openthinker:7b-v2-q8_0
openthinker:7b-v2-fp16
openthinker:7b-q4_K_M
openthinker:7b-q8_0
openthinker:7b-fp16
openthinker:32b-v2-q4_K_M
openthinker:32b-v2-q8_0
openthinker:32b-v2-fp16
openthinker:32b-q4_K_M
openthinker:32b-q8_0
openthinker:32b-fp16
codestral:latest
codestral
codestral:v0.1
codestral:22b
codestral:22b-v0.1-q2_K
codestral:22b-v0.1-q3_K_S
codestral:22b-v0.1-q3_K_M
codestral:22b-v0.1-q3_K_L
codestral:22b-v0.1-q4_0
codestral:22b-v0.1-q4_1
codestral:22b-v0.1-q4_K_S
codestral:22b-v0.1-q4_K_M
codestral:22b-v0.1-q5_0
codestral:22b-v0.1-q5_1
codestral:22b-v0.1-q5_K_S
codestral:22b-v0.1-q5_K_M
codestral:22b-v0.1-q6_K
codestral:22b-v0.1-q8_0
smollm:latest
smollm
smollm:135m
smollm:360m
smollm:1.7b
smollm:135m-base-v0.2-q2_K
smollm:135m-base-v0.2-q3_K_S
smollm:135m-base-v0.2-q3_K_M
smollm:135m-base-v0.2-q3_K_L
smollm:135m-base-v0.2-q4_0
smollm:135m-base-v0.2-q4_1
smollm:135m-base-v0.2-q4_K_S
smollm:135m-base-v0.2-q4_K_M
smollm:135m-base-v0.2-q5_0
smollm:135m-base-v0.2-q5_1
smollm:135m-base-v0.2-q5_K_S
smollm:135m-base-v0.2-q5_K_M
smollm:135m-base-v0.2-q6_K
smollm:135m-base-v0.2-q8_0
smollm:135m-base-v0.2-fp16
smollm:135m-instruct-v0.2-q2_K
smollm:135m-instruct-v0.2-q3_K_S
smollm:135m-instruct-v0.2-q3_K_M
smollm:135m-instruct-v0.2-q3_K_L
smollm:135m-instruct-v0.2-q4_0
smollm:135m-instruct-v0.2-q4_1
smollm:135m-instruct-v0.2-q4_K_S
smollm:135m-instruct-v0.2-q4_K_M
smollm:135m-instruct-v0.2-q5_0
smollm:135m-instruct-v0.2-q5_1
smollm:135m-instruct-v0.2-q5_K_S
smollm:135m-instruct-v0.2-q5_K_M
smollm:135m-instruct-v0.2-q6_K
smollm:135m-instruct-v0.2-q8_0
smollm:135m-instruct-v0.2-fp16
smollm:360m-base-v0.2-q2_K
smollm:360m-base-v0.2-q3_K_S
smollm:360m-base-v0.2-q3_K_M
smollm:360m-base-v0.2-q3_K_L
smollm:360m-base-v0.2-q4_0
smollm:360m-base-v0.2-q4_1
smollm:360m-base-v0.2-q4_K_S
smollm:360m-base-v0.2-q4_K_M
smollm:360m-base-v0.2-q5_0
smollm:360m-base-v0.2-q5_1
smollm:360m-base-v0.2-q5_K_S
smollm:360m-base-v0.2-q5_K_M
smollm:360m-base-v0.2-q6_K
smollm:360m-base-v0.2-q8_0
smollm:360m-base-v0.2-fp16
smollm:360m-instruct-v0.2-q2_K
smollm:360m-instruct-v0.2-q3_K_S
smollm:360m-instruct-v0.2-q3_K_M
smollm:360m-instruct-v0.2-q3_K_L
smollm:360m-instruct-v0.2-q4_0
smollm:360m-instruct-v0.2-q4_1
smollm:360m-instruct-v0.2-q4_K_S
smollm:360m-instruct-v0.2-q4_K_M
smollm:360m-instruct-v0.2-q5_0
smollm:360m-instruct-v0.2-q5_1
smollm:360m-instruct-v0.2-q5_K_S
smollm:360m-instruct-v0.2-q5_K_M
smollm:360m-instruct-v0.2-q6_K
smollm:360m-instruct-v0.2-q8_0
smollm:360m-instruct-v0.2-fp16
smollm:1.7b-base-v0.2-q2_K
smollm:1.7b-base-v0.2-q3_K_S
smollm:1.7b-base-v0.2-q3_K_M
smollm:1.7b-base-v0.2-q3_K_L
smollm:1.7b-base-v0.2-q4_0
smollm:1.7b-base-v0.2-q4_1
smollm:1.7b-base-v0.2-q4_K_S
smollm:1.7b-base-v0.2-q4_K_M
smollm:1.7b-base-v0.2-q5_0
smollm:1.7b-base-v0.2-q5_1
smollm:1.7b-base-v0.2-q5_K_S
smollm:1.7b-base-v0.2-q5_K_M
smollm:1.7b-base-v0.2-q6_K
smollm:1.7b-base-v0.2-q8_0
smollm:1.7b-base-v0.2-fp16
smollm:1.7b-instruct-v0.2-q2_K
smollm:1.7b-instruct-v0.2-q3_K_S
smollm:1.7b-instruct-v0.2-q3_K_M
smollm:1.7b-instruct-v0.2-q3_K_L
smollm:1.7b-instruct-v0.2-q4_0
smollm:1.7b-instruct-v0.2-q4_1
smollm:1.7b-instruct-v0.2-q4_K_S
smollm:1.7b-instruct-v0.2-q4_K_M
smollm:1.7b-instruct-v0.2-q5_0
smollm:1.7b-instruct-v0.2-q5_1
smollm:1.7b-instruct-v0.2-q5_K_S
smollm:1.7b-instruct-v0.2-q5_K_M
smollm:1.7b-instruct-v0.2-q6_K
smollm:1.7b-instruct-v0.2-q8_0
smollm:1.7b-instruct-v0.2-fp16
granite3.2-vision:latest
granite3.2-vision
granite3.2-vision:2b
granite3.2-vision:2b-q4_K_M
granite3.2-vision:2b-q8_0
granite3.2-vision:2b-fp16
devstral:latest
devstral
devstral:24b
devstral:24b-small-2505-q4_K_M
devstral:24b-small-2505-q8_0
devstral:24b-small-2505-fp16
wizardlm2:latest
wizardlm2
wizardlm2:7b
wizardlm2:8x22b
wizardlm2:7b-q2_K
wizardlm2:7b-q3_K_S
wizardlm2:7b-q3_K_M
wizardlm2:7b-q3_K_L
wizardlm2:7b-q4_0
wizardlm2:7b-q4_1
wizardlm2:7b-q4_K_S
wizardlm2:7b-q4_K_M
wizardlm2:7b-q5_0
wizardlm2:7b-q5_1
wizardlm2:7b-q5_K_S
wizardlm2:7b-q5_K_M
wizardlm2:7b-q6_K
wizardlm2:7b-q8_0
wizardlm2:7b-fp16
wizardlm2:8x22b-q2_K
wizardlm2:8x22b-q4_0
wizardlm2:8x22b-q8_0
wizardlm2:8x22b-fp16
dolphin-mistral:latest
dolphin-mistral
dolphin-mistral:v2
dolphin-mistral:v2.1
dolphin-mistral:v2.2
dolphin-mistral:v2.2.1
dolphin-mistral:v2.6
dolphin-mistral:v2.8
dolphin-mistral:7b
dolphin-mistral:7b-v2
dolphin-mistral:7b-v2-q2_K
dolphin-mistral:7b-v2-q3_K_S
dolphin-mistral:7b-v2-q3_K_M
dolphin-mistral:7b-v2-q3_K_L
dolphin-mistral:7b-v2-q4_0
dolphin-mistral:7b-v2-q4_1
dolphin-mistral:7b-v2-q4_K_S
dolphin-mistral:7b-v2-q4_K_M
dolphin-mistral:7b-v2-q5_0
dolphin-mistral:7b-v2-q5_1
dolphin-mistral:7b-v2-q5_K_S
dolphin-mistral:7b-v2-q5_K_M
dolphin-mistral:7b-v2-q6_K
dolphin-mistral:7b-v2-q8_0
dolphin-mistral:7b-v2-fp16
dolphin-mistral:7b-v2.1
dolphin-mistral:7b-v2.1-q2_K
dolphin-mistral:7b-v2.1-q3_K_S
dolphin-mistral:7b-v2.1-q3_K_M
dolphin-mistral:7b-v2.1-q3_K_L
dolphin-mistral:7b-v2.1-q4_0
dolphin-mistral:7b-v2.1-q4_1
dolphin-mistral:7b-v2.1-q4_K_S
dolphin-mistral:7b-v2.1-q4_K_M
dolphin-mistral:7b-v2.1-q5_0
dolphin-mistral:7b-v2.1-q5_1
dolphin-mistral:7b-v2.1-q5_K_S
dolphin-mistral:7b-v2.1-q5_K_M
dolphin-mistral:7b-v2.1-q6_K
dolphin-mistral:7b-v2.1-q8_0
dolphin-mistral:7b-v2.1-fp16
dolphin-mistral:7b-v2.2
dolphin-mistral:7b-v2.2-q2_K
dolphin-mistral:7b-v2.2-q3_K_S
dolphin-mistral:7b-v2.2-q3_K_M
dolphin-mistral:7b-v2.2-q3_K_L
dolphin-mistral:7b-v2.2-q4_0
dolphin-mistral:7b-v2.2-q4_1
dolphin-mistral:7b-v2.2-q4_K_S
dolphin-mistral:7b-v2.2-q4_K_M
dolphin-mistral:7b-v2.2-q5_0
dolphin-mistral:7b-v2.2-q5_1
dolphin-mistral:7b-v2.2-q5_K_S
dolphin-mistral:7b-v2.2-q5_K_M
dolphin-mistral:7b-v2.2-q6_K
dolphin-mistral:7b-v2.2-q8_0
dolphin-mistral:7b-v2.2-fp16
dolphin-mistral:7b-v2.2.1
dolphin-mistral:7b-v2.2.1-q2_K
dolphin-mistral:7b-v2.2.1-q3_K_S
dolphin-mistral:7b-v2.2.1-q3_K_M
dolphin-mistral:7b-v2.2.1-q3_K_L
dolphin-mistral:7b-v2.2.1-q4_0
dolphin-mistral:7b-v2.2.1-q4_1
dolphin-mistral:7b-v2.2.1-q4_K_S
dolphin-mistral:7b-v2.2.1-q4_K_M
dolphin-mistral:7b-v2.2.1-q5_0
dolphin-mistral:7b-v2.2.1-q5_1
dolphin-mistral:7b-v2.2.1-q5_K_S
dolphin-mistral:7b-v2.2.1-q5_K_M
dolphin-mistral:7b-v2.2.1-q6_K
dolphin-mistral:7b-v2.2.1-q8_0
dolphin-mistral:7b-v2.2.1-fp16
dolphin-mistral:7b-v2.6
dolphin-mistral:7b-v2.6-dpo-laser
dolphin-mistral:7b-v2.6-dpo-laser-q2_K
dolphin-mistral:7b-v2.6-q2_K
dolphin-mistral:7b-v2.6-dpo-laser-q3_K_S
dolphin-mistral:7b-v2.6-q3_K_S
dolphin-mistral:7b-v2.6-dpo-laser-q3_K_M
dolphin-mistral:7b-v2.6-q3_K_M
dolphin-mistral:7b-v2.6-dpo-laser-q3_K_L
dolphin-mistral:7b-v2.6-q3_K_L
dolphin-mistral:7b-v2.6-dpo-laser-q4_0
dolphin-mistral:7b-v2.6-q4_0
dolphin-mistral:7b-v2.6-dpo-laser-q4_1
dolphin-mistral:7b-v2.6-q4_1
dolphin-mistral:7b-v2.6-dpo-laser-q4_K_S
dolphin-mistral:7b-v2.6-q4_K_S
dolphin-mistral:7b-v2.6-dpo-laser-q4_K_M
dolphin-mistral:7b-v2.6-q4_K_M
dolphin-mistral:7b-v2.6-dpo-laser-q5_0
dolphin-mistral:7b-v2.6-q5_0
dolphin-mistral:7b-v2.6-dpo-laser-q5_1
dolphin-mistral:7b-v2.6-q5_1
dolphin-mistral:7b-v2.6-dpo-laser-q5_K_S
dolphin-mistral:7b-v2.6-q5_K_S
dolphin-mistral:7b-v2.6-dpo-laser-q5_K_M
dolphin-mistral:7b-v2.6-q5_K_M
dolphin-mistral:7b-v2.6-dpo-laser-q6_K
dolphin-mistral:7b-v2.6-q6_K
dolphin-mistral:7b-v2.6-dpo-laser-q8_0
dolphin-mistral:7b-v2.6-q8_0
dolphin-mistral:7b-v2.6-dpo-laser-fp16
dolphin-mistral:7b-v2.6-fp16
dolphin-mistral:7b-v2.8
dolphin-mistral:7b-v2.8-q2_K
dolphin-mistral:7b-v2.8-q3_K_S
dolphin-mistral:7b-v2.8-q3_K_M
dolphin-mistral:7b-v2.8-q3_K_L
dolphin-mistral:7b-v2.8-q4_0
dolphin-mistral:7b-v2.8-q4_1
dolphin-mistral:7b-v2.8-q4_K_S
dolphin-mistral:7b-v2.8-q4_K_M
dolphin-mistral:7b-v2.8-q5_0
dolphin-mistral:7b-v2.8-q5_1
dolphin-mistral:7b-v2.8-q5_K_S
dolphin-mistral:7b-v2.8-q5_K_M
dolphin-mistral:7b-v2.8-q6_K
dolphin-mistral:7b-v2.8-q8_0
dolphin-mistral:7b-v2.8-fp16
deepcoder:latest
deepcoder
deepcoder:1.5b
deepcoder:14b
deepcoder:1.5b-preview-q4_K_M
deepcoder:1.5b-preview-q8_0
deepcoder:1.5b-preview-fp16
deepcoder:14b-preview-q4_K_M
deepcoder:14b-preview-q8_0
deepcoder:14b-preview-fp16
moondream:latest
moondream
moondream:v2
moondream:1.8b
moondream:1.8b-v2-q2_K
moondream:1.8b-v2-q3_K_S
moondream:1.8b-v2-q3_K_M
moondream:1.8b-v2-q3_K_L
moondream:1.8b-v2-q4_0
moondream:1.8b-v2-q4_1
moondream:1.8b-v2-q4_K_S
moondream:1.8b-v2-q4_K_M
moondream:1.8b-v2-q5_0
moondream:1.8b-v2-q5_1
moondream:1.8b-v2-q5_K_S
moondream:1.8b-v2-q5_K_M
moondream:1.8b-v2-q6_K
moondream:1.8b-v2-q8_0
moondream:1.8b-v2-fp16
mistral-small3.1:latest
mistral-small3.1
mistral-small3.1:24b
mistral-small3.1:24b-instruct-2503-q4_K_M
mistral-small3.1:24b-instruct-2503-q8_0
mistral-small3.1:24b-instruct-2503-fp16
command-r:latest
command-r
command-r:v0.1
command-r:35b
command-r:35b-08-2024-q2_K
command-r:35b-08-2024-q3_K_S
command-r:35b-08-2024-q3_K_M
command-r:35b-08-2024-q3_K_L
command-r:35b-08-2024-q4_0
command-r:35b-08-2024-q4_1
command-r:35b-08-2024-q4_K_S
command-r:35b-08-2024-q4_K_M
command-r:35b-08-2024-q5_0
command-r:35b-08-2024-q5_1
command-r:35b-08-2024-q5_K_S
command-r:35b-08-2024-q5_K_M
command-r:35b-08-2024-q6_K
command-r:35b-08-2024-q8_0
command-r:35b-08-2024-fp16
command-r:35b-v0.1-q2_K
command-r:35b-v0.1-q3_K_S
command-r:35b-v0.1-q3_K_M
command-r:35b-v0.1-q3_K_L
command-r:35b-v0.1-q4_0
command-r:35b-v0.1-q4_1
command-r:35b-v0.1-q4_K_S
command-r:35b-v0.1-q4_K_M
command-r:35b-v0.1-q5_1
command-r:35b-v0.1-q5_K_S
command-r:35b-v0.1-q5_K_M
command-r:35b-v0.1-q6_K
command-r:35b-v0.1-q8_0
command-r:35b-v0.1-fp16
granite-code:latest
granite-code
granite-code:3b
granite-code:8b
granite-code:20b
granite-code:34b
granite-code:3b-base
granite-code:3b-base-q2_K
granite-code:3b-base-q3_K_S
granite-code:3b-base-q3_K_M
granite-code:3b-base-q3_K_L
granite-code:3b-base-q4_0
granite-code:3b-base-q4_1
granite-code:3b-base-q4_K_S
granite-code:3b-base-q4_K_M
granite-code:3b-base-q5_0
granite-code:3b-base-q5_1
granite-code:3b-base-q5_K_S
granite-code:3b-base-q5_K_M
granite-code:3b-base-q6_K
granite-code:3b-base-q8_0
granite-code:3b-base-fp16
granite-code:3b-instruct
granite-code:3b-instruct-128k-q2_K
granite-code:3b-instruct-q2_K
granite-code:3b-instruct-128k-q3_K_S
granite-code:3b-instruct-q3_K_S
granite-code:3b-instruct-128k-q3_K_M
granite-code:3b-instruct-q3_K_M
granite-code:3b-instruct-128k-q3_K_L
granite-code:3b-instruct-q3_K_L
granite-code:3b-instruct-128k-q4_0
granite-code:3b-instruct-q4_0
granite-code:3b-instruct-128k-q4_1
granite-code:3b-instruct-q4_1
granite-code:3b-instruct-128k-q4_K_S
granite-code:3b-instruct-q4_K_S
granite-code:3b-instruct-128k-q4_K_M
granite-code:3b-instruct-q4_K_M
granite-code:3b-instruct-128k-q5_0
granite-code:3b-instruct-q5_0
granite-code:3b-instruct-128k-q5_1
granite-code:3b-instruct-q5_1
granite-code:3b-instruct-128k-q5_K_S
granite-code:3b-instruct-q5_K_S
granite-code:3b-instruct-128k-q5_K_M
granite-code:3b-instruct-q5_K_M
granite-code:3b-instruct-128k-q6_K
granite-code:3b-instruct-q6_K
granite-code:3b-instruct-128k-q8_0
granite-code:3b-instruct-q8_0
granite-code:3b-instruct-128k-fp16
granite-code:3b-instruct-fp16
granite-code:8b-base
granite-code:8b-base-q2_K
granite-code:8b-base-q3_K_S
granite-code:8b-base-q3_K_M
granite-code:8b-base-q3_K_L
granite-code:8b-base-q4_0
granite-code:8b-base-q4_1
granite-code:8b-base-q4_K_S
granite-code:8b-base-q4_K_M
granite-code:8b-base-q5_0
granite-code:8b-base-q5_1
granite-code:8b-base-q5_K_S
granite-code:8b-base-q5_K_M
granite-code:8b-base-q6_K
granite-code:8b-base-q8_0
granite-code:8b-base-fp16
granite-code:8b-instruct
granite-code:8b-instruct-q2_K
granite-code:8b-instruct-q3_K_S
granite-code:8b-instruct-q3_K_M
granite-code:8b-instruct-q3_K_L
granite-code:8b-instruct-128k-q4_0
granite-code:8b-instruct-q4_0
granite-code:8b-instruct-128k-q4_1
granite-code:8b-instruct-q4_1
granite-code:8b-instruct-q4_K_S
granite-code:8b-instruct-q4_K_M
granite-code:8b-instruct-q5_0
granite-code:8b-instruct-q5_1
granite-code:8b-instruct-q5_K_S
granite-code:8b-instruct-q5_K_M
granite-code:8b-instruct-q6_K
granite-code:8b-instruct-q8_0
granite-code:8b-instruct-fp16
granite-code:20b-base
granite-code:20b-base-q2_K
granite-code:20b-base-q3_K_S
granite-code:20b-base-q3_K_M
granite-code:20b-base-q3_K_L
granite-code:20b-base-q4_0
granite-code:20b-base-q4_1
granite-code:20b-base-q4_K_S
granite-code:20b-base-q4_K_M
granite-code:20b-base-q5_0
granite-code:20b-base-q5_1
granite-code:20b-base-q5_K_S
granite-code:20b-base-q5_K_M
granite-code:20b-base-q6_K
granite-code:20b-base-q8_0
granite-code:20b-base-fp16
granite-code:20b-instruct
granite-code:20b-instruct-8k-q2_K
granite-code:20b-instruct-q2_K
granite-code:20b-instruct-8k-q3_K_S
granite-code:20b-instruct-q3_K_S
granite-code:20b-instruct-8k-q3_K_M
granite-code:20b-instruct-q3_K_M
granite-code:20b-instruct-8k-q3_K_L
granite-code:20b-instruct-q3_K_L
granite-code:20b-instruct-8k-q4_0
granite-code:20b-instruct-q4_0
granite-code:20b-instruct-8k-q4_1
granite-code:20b-instruct-q4_1
granite-code:20b-instruct-8k-q4_K_S
granite-code:20b-instruct-q4_K_S
granite-code:20b-instruct-8k-q4_K_M
granite-code:20b-instruct-q4_K_M
granite-code:20b-instruct-8k-q5_0
granite-code:20b-instruct-q5_0
granite-code:20b-instruct-8k-q5_1
granite-code:20b-instruct-q5_1
granite-code:20b-instruct-8k-q5_K_S
granite-code:20b-instruct-q5_K_S
granite-code:20b-instruct-8k-q5_K_M
granite-code:20b-instruct-q5_K_M
granite-code:20b-instruct-8k-q6_K
granite-code:20b-instruct-q6_K
granite-code:20b-instruct-8k-q8_0
granite-code:20b-instruct-q8_0
granite-code:20b-instruct-8k-fp16
granite-code:34b-base
granite-code:34b-base-q2_K
granite-code:34b-base-q3_K_S
granite-code:34b-base-q3_K_M
granite-code:34b-base-q3_K_L
granite-code:34b-base-q4_0
granite-code:34b-base-q4_1
granite-code:34b-base-q4_K_S
granite-code:34b-base-q4_K_M
granite-code:34b-base-q5_0
granite-code:34b-base-q5_1
granite-code:34b-base-q5_K_S
granite-code:34b-base-q5_K_M
granite-code:34b-base-q6_K
granite-code:34b-base-q8_0
granite-code:34b-instruct
granite-code:34b-instruct-q2_K
granite-code:34b-instruct-q3_K_S
granite-code:34b-instruct-q3_K_M
granite-code:34b-instruct-q3_K_L
granite-code:34b-instruct-q4_0
granite-code:34b-instruct-q4_1
granite-code:34b-instruct-q4_K_S
granite-code:34b-instruct-q4_K_M
granite-code:34b-instruct-q5_0
granite-code:34b-instruct-q5_1
granite-code:34b-instruct-q5_K_S
granite-code:34b-instruct-q5_K_M
granite-code:34b-instruct-q6_K
granite-code:34b-instruct-q8_0
phi3.5:latest
phi3.5
phi3.5:3.8b
phi3.5:3.8b-mini-instruct-q2_K
phi3.5:3.8b-mini-instruct-q3_K_S
phi3.5:3.8b-mini-instruct-q3_K_M
phi3.5:3.8b-mini-instruct-q3_K_L
phi3.5:3.8b-mini-instruct-q4_0
phi3.5:3.8b-mini-instruct-q4_1
phi3.5:3.8b-mini-instruct-q4_K_S
phi3.5:3.8b-mini-instruct-q4_K_M
phi3.5:3.8b-mini-instruct-q5_0
phi3.5:3.8b-mini-instruct-q5_1
phi3.5:3.8b-mini-instruct-q5_K_S
phi3.5:3.8b-mini-instruct-q5_K_M
phi3.5:3.8b-mini-instruct-q6_K
phi3.5:3.8b-mini-instruct-q8_0
phi3.5:3.8b-mini-instruct-fp16
hermes3:latest
hermes3
hermes3:3b
hermes3:8b
hermes3:70b
hermes3:405b
hermes3:3b-llama3.2-q2_K
hermes3:3b-llama3.2-q3_K_S
hermes3:3b-llama3.2-q3_K_M
hermes3:3b-llama3.2-q3_K_L
hermes3:3b-llama3.2-q4_0
hermes3:3b-llama3.2-q4_1
hermes3:3b-llama3.2-q4_K_S
hermes3:3b-llama3.2-q4_K_M
hermes3:3b-llama3.2-q5_0
hermes3:3b-llama3.2-q5_1
hermes3:3b-llama3.2-q5_K_S
hermes3:3b-llama3.2-q5_K_M
hermes3:3b-llama3.2-q6_K
hermes3:3b-llama3.2-q8_0
hermes3:3b-llama3.2-fp16
hermes3:8b-llama3.1-q2_K
hermes3:8b-llama3.1-q3_K_S
hermes3:8b-llama3.1-q3_K_M
hermes3:8b-llama3.1-q3_K_L
hermes3:8b-llama3.1-q4_0
hermes3:8b-llama3.1-q4_1
hermes3:8b-llama3.1-q4_K_S
hermes3:8b-llama3.1-q4_K_M
hermes3:8b-llama3.1-q5_0
hermes3:8b-llama3.1-q5_1
hermes3:8b-llama3.1-q5_K_S
hermes3:8b-llama3.1-q5_K_M
hermes3:8b-llama3.1-q6_K
hermes3:8b-llama3.1-q8_0
hermes3:8b-llama3.1-fp16
hermes3:70b-llama3.1-q2_K
hermes3:70b-llama3.1-q3_K_S
hermes3:70b-llama3.1-q3_K_M
hermes3:70b-llama3.1-q3_K_L
hermes3:70b-llama3.1-q4_0
hermes3:70b-llama3.1-q4_1
hermes3:70b-llama3.1-q4_K_S
hermes3:70b-llama3.1-q4_K_M
hermes3:70b-llama3.1-q5_0
hermes3:70b-llama3.1-q5_1
hermes3:70b-llama3.1-q5_K_S
hermes3:70b-llama3.1-q5_K_M
hermes3:70b-llama3.1-q6_K
hermes3:70b-llama3.1-q8_0
hermes3:70b-llama3.1-fp16
hermes3:405b-llama3.1-q2_K
hermes3:405b-llama3.1-q3_K_S
hermes3:405b-llama3.1-q3_K_M
hermes3:405b-llama3.1-q3_K_L
hermes3:405b-llama3.1-q4_0
hermes3:405b-llama3.1-q4_1
hermes3:405b-llama3.1-q4_K_S
hermes3:405b-llama3.1-q4_K_M
hermes3:405b-llama3.1-q5_0
hermes3:405b-llama3.1-q5_1
hermes3:405b-llama3.1-q5_K_S
hermes3:405b-llama3.1-q5_K_M
hermes3:405b-llama3.1-q6_K
hermes3:405b-llama3.1-q8_0
hermes3:405b-llama3.1-fp16
yi:latest
yi
yi:v1.5
yi:6b
yi:9b
yi:34b
yi:6b-200k
yi:6b-200k-q2_K
yi:6b-200k-q3_K_S
yi:6b-200k-q3_K_M
yi:6b-200k-q3_K_L
yi:6b-200k-q4_0
yi:6b-200k-q4_1
yi:6b-200k-q4_K_S
yi:6b-200k-q4_K_M
yi:6b-200k-q5_0
yi:6b-200k-q5_1
yi:6b-200k-q5_K_S
yi:6b-200k-q5_K_M
yi:6b-200k-q6_K
yi:6b-200k-q8_0
yi:6b-200k-fp16
yi:6b-chat
yi:6b-chat-q2_K
yi:6b-chat-v1.5-q2_K
yi:6b-chat-q3_K_S
yi:6b-chat-v1.5-q3_K_S
yi:6b-chat-q3_K_M
yi:6b-chat-v1.5-q3_K_M
yi:6b-chat-q3_K_L
yi:6b-chat-v1.5-q3_K_L
yi:6b-chat-q4_0
yi:6b-chat-v1.5-q4_0
yi:6b-chat-q4_1
yi:6b-chat-v1.5-q4_1
yi:6b-chat-q4_K_S
yi:6b-chat-v1.5-q4_K_S
yi:6b-chat-q4_K_M
yi:6b-chat-v1.5-q4_K_M
yi:6b-chat-q5_0
yi:6b-chat-v1.5-q5_0
yi:6b-chat-q5_1
yi:6b-chat-v1.5-q5_1
yi:6b-chat-q5_K_S
yi:6b-chat-v1.5-q5_K_S
yi:6b-chat-q5_K_M
yi:6b-chat-v1.5-q5_K_M
yi:6b-chat-q6_K
yi:6b-chat-v1.5-q6_K
yi:6b-chat-q8_0
yi:6b-chat-v1.5-q8_0
yi:6b-chat-fp16
yi:6b-chat-v1.5-fp16
yi:6b-v1.5
yi:6b-v1.5-q2_K
yi:6b-v1.5-q3_K_S
yi:6b-v1.5-q3_K_M
yi:6b-v1.5-q3_K_L
yi:6b-v1.5-q4_0
yi:6b-v1.5-q4_1
yi:6b-v1.5-q4_K_S
yi:6b-v1.5-q4_K_M
yi:6b-v1.5-q5_0
yi:6b-v1.5-q5_1
yi:6b-v1.5-q5_K_S
yi:6b-v1.5-q5_K_M
yi:6b-v1.5-q6_K
yi:6b-v1.5-q8_0
yi:6b-v1.5-fp16
yi:6b-q2_K
yi:6b-q3_K_S
yi:6b-q3_K_M
yi:6b-q3_K_L
yi:6b-q4_0
yi:6b-q4_1
yi:6b-q4_K_S
yi:6b-q4_K_M
yi:6b-q5_0
yi:6b-q5_1
yi:6b-q5_K_S
yi:6b-q5_K_M
yi:6b-q6_K
yi:6b-q8_0
yi:6b-fp16
yi:9b-chat
yi:9b-chat-v1.5-q2_K
yi:9b-chat-v1.5-q3_K_S
yi:9b-chat-v1.5-q3_K_M
yi:9b-chat-v1.5-q3_K_L
yi:9b-chat-v1.5-q4_0
yi:9b-chat-v1.5-q4_1
yi:9b-chat-v1.5-q4_K_S
yi:9b-chat-v1.5-q4_K_M
yi:9b-chat-v1.5-q5_0
yi:9b-chat-v1.5-q5_1
yi:9b-chat-v1.5-q5_K_S
yi:9b-chat-v1.5-q5_K_M
yi:9b-chat-v1.5-q6_K
yi:9b-chat-v1.5-q8_0
yi:9b-chat-v1.5-fp16
yi:9b-v1.5
yi:9b-v1.5-q2_K
yi:9b-v1.5-q3_K_S
yi:9b-v1.5-q3_K_M
yi:9b-v1.5-q3_K_L
yi:9b-v1.5-q4_0
yi:9b-v1.5-q4_1
yi:9b-v1.5-q4_K_S
yi:9b-v1.5-q4_K_M
yi:9b-v1.5-q5_0
yi:9b-v1.5-q5_1
yi:9b-v1.5-q5_K_S
yi:9b-v1.5-q5_K_M
yi:9b-v1.5-q6_K
yi:9b-v1.5-q8_0
yi:9b-v1.5-fp16
yi:34b-chat
yi:34b-chat-q2_K
yi:34b-chat-v1.5-q2_K
yi:34b-chat-q3_K_S
yi:34b-chat-v1.5-q3_K_S
yi:34b-chat-q3_K_M
yi:34b-chat-v1.5-q3_K_M
yi:34b-chat-q3_K_L
yi:34b-chat-v1.5-q3_K_L
yi:34b-chat-q4_0
yi:34b-chat-v1.5-q4_0
yi:34b-chat-q4_1
yi:34b-chat-v1.5-q4_1
yi:34b-chat-q4_K_S
yi:34b-chat-v1.5-q4_K_S
yi:34b-chat-q4_K_M
yi:34b-chat-v1.5-q4_K_M
yi:34b-chat-q5_0
yi:34b-chat-v1.5-q5_0
yi:34b-chat-q5_1
yi:34b-chat-v1.5-q5_1
yi:34b-chat-q5_K_S
yi:34b-chat-v1.5-q5_K_S
yi:34b-chat-q5_K_M
yi:34b-chat-v1.5-q5_K_M
yi:34b-chat-q6_K
yi:34b-chat-v1.5-q6_K
yi:34b-chat-q8_0
yi:34b-chat-v1.5-q8_0
yi:34b-chat-fp16
yi:34b-chat-v1.5-fp16
yi:34b-v1.5
yi:34b-v1.5-q2_K
yi:34b-v1.5-q3_K_S
yi:34b-v1.5-q3_K_M
yi:34b-v1.5-q3_K_L
yi:34b-v1.5-q4_0
yi:34b-v1.5-q4_1
yi:34b-v1.5-q4_K_S
yi:34b-v1.5-q4_K_M
yi:34b-v1.5-q5_0
yi:34b-v1.5-q5_1
yi:34b-v1.5-q5_K_S
yi:34b-v1.5-q5_K_M
yi:34b-v1.5-q6_K
yi:34b-v1.5-q8_0
yi:34b-v1.5-fp16
yi:34b-q2_K
yi:34b-q3_K_S
yi:34b-q3_K_M
yi:34b-q3_K_L
yi:34b-q4_0
yi:34b-q4_1
yi:34b-q4_K_S
yi:34b-q4_K_M
yi:34b-q5_0
yi:34b-q5_1
yi:34b-q5_K_S
yi:34b-q6_K
embeddinggemma:latest
embeddinggemma
embeddinggemma:300m
embeddinggemma:300m-qat-q4_0
embeddinggemma:300m-qat-q8_0
embeddinggemma:300m-bf16
zephyr:latest
zephyr
zephyr:7b
zephyr:141b
zephyr:7b-alpha
zephyr:7b-alpha-q2_K
zephyr:7b-alpha-q3_K_S
zephyr:7b-alpha-q3_K_M
zephyr:7b-alpha-q3_K_L
zephyr:7b-alpha-q4_0
zephyr:7b-alpha-q4_1
zephyr:7b-alpha-q4_K_S
zephyr:7b-alpha-q4_K_M
zephyr:7b-alpha-q5_0
zephyr:7b-alpha-q5_1
zephyr:7b-alpha-q5_K_S
zephyr:7b-alpha-q5_K_M
zephyr:7b-alpha-q6_K
zephyr:7b-alpha-q8_0
zephyr:7b-alpha-fp16
zephyr:7b-beta
zephyr:7b-beta-q2_K
zephyr:7b-beta-q3_K_S
zephyr:7b-beta-q3_K_M
zephyr:7b-beta-q3_K_L
zephyr:7b-beta-q4_0
zephyr:7b-beta-q4_1
zephyr:7b-beta-q4_K_S
zephyr:7b-beta-q4_K_M
zephyr:7b-beta-q5_0
zephyr:7b-beta-q5_1
zephyr:7b-beta-q5_K_S
zephyr:7b-beta-q5_K_M
zephyr:7b-beta-q6_K
zephyr:7b-beta-q8_0
zephyr:7b-beta-fp16
zephyr:141b-v0.1
zephyr:141b-v0.1-q2_K
zephyr:141b-v0.1-q4_0
zephyr:141b-v0.1-q8_0
zephyr:141b-v0.1-fp16
exaone-deep:latest
exaone-deep
exaone-deep:2.4b
exaone-deep:7.8b
exaone-deep:32b
exaone-deep:2.4b-q4_K_M
exaone-deep:2.4b-q8_0
exaone-deep:2.4b-fp16
exaone-deep:7.8b-q4_K_M
exaone-deep:7.8b-q8_0
exaone-deep:7.8b-fp16
exaone-deep:32b-q4_K_M
exaone-deep:32b-q8_0
exaone-deep:32b-fp16
granite4:latest
granite4
granite4:micro
granite4:350m
granite4:1b
granite4:3b
granite4:350m-h
granite4:350m-h-q8_0
granite4:350m-bf16
granite4:1b-h
granite4:1b-h-q8_0
granite4:1b-bf16
granite4:3b-h
granite4:7b-a1b-h
granite4:32b-a9b-h
granite4:micro-h
granite4:small-h
granite4:tiny-h
mistral-large:latest
mistral-large
mistral-large:123b
mistral-large:123b-instruct-2407-q2_K
mistral-large:123b-instruct-2407-q3_K_S
mistral-large:123b-instruct-2407-q3_K_M
mistral-large:123b-instruct-2407-q3_K_L
mistral-large:123b-instruct-2407-q4_0
mistral-large:123b-instruct-2407-q4_1
mistral-large:123b-instruct-2407-q4_K_S
mistral-large:123b-instruct-2407-q4_K_M
mistral-large:123b-instruct-2407-q5_0
mistral-large:123b-instruct-2407-q5_1
mistral-large:123b-instruct-2407-q5_K_S
mistral-large:123b-instruct-2407-q5_K_M
mistral-large:123b-instruct-2407-q6_K
mistral-large:123b-instruct-2407-q8_0
mistral-large:123b-instruct-2407-fp16
mistral-large:123b-instruct-2411-q2_K
mistral-large:123b-instruct-2411-q3_K_S
mistral-large:123b-instruct-2411-q3_K_M
mistral-large:123b-instruct-2411-q3_K_L
mistral-large:123b-instruct-2411-q4_0
mistral-large:123b-instruct-2411-q4_1
mistral-large:123b-instruct-2411-q4_K_S
mistral-large:123b-instruct-2411-q4_K_M
mistral-large:123b-instruct-2411-q5_0
mistral-large:123b-instruct-2411-q5_1
mistral-large:123b-instruct-2411-q5_K_S
mistral-large:123b-instruct-2411-q5_K_M
mistral-large:123b-instruct-2411-q6_K
mistral-large:123b-instruct-2411-q8_0
mistral-large:123b-instruct-2411-fp16
wizard-vicuna-uncensored:latest
wizard-vicuna-uncensored
wizard-vicuna-uncensored:7b
wizard-vicuna-uncensored:13b
wizard-vicuna-uncensored:30b
wizard-vicuna-uncensored:7b-q2_K
wizard-vicuna-uncensored:7b-q3_K_S
wizard-vicuna-uncensored:7b-q3_K_M
wizard-vicuna-uncensored:7b-q3_K_L
wizard-vicuna-uncensored:7b-q4_0
wizard-vicuna-uncensored:7b-q4_1
wizard-vicuna-uncensored:7b-q4_K_S
wizard-vicuna-uncensored:7b-q4_K_M
wizard-vicuna-uncensored:7b-q5_0
wizard-vicuna-uncensored:7b-q5_1
wizard-vicuna-uncensored:7b-q5_K_S
wizard-vicuna-uncensored:7b-q5_K_M
wizard-vicuna-uncensored:7b-q6_K
wizard-vicuna-uncensored:7b-q8_0
wizard-vicuna-uncensored:7b-fp16
wizard-vicuna-uncensored:13b-q2_K
wizard-vicuna-uncensored:13b-q3_K_S
wizard-vicuna-uncensored:13b-q3_K_M
wizard-vicuna-uncensored:13b-q3_K_L
wizard-vicuna-uncensored:13b-q4_0
wizard-vicuna-uncensored:13b-q4_1
wizard-vicuna-uncensored:13b-q4_K_S
wizard-vicuna-uncensored:13b-q4_K_M
wizard-vicuna-uncensored:13b-q5_0
wizard-vicuna-uncensored:13b-q5_1
wizard-vicuna-uncensored:13b-q5_K_S
wizard-vicuna-uncensored:13b-q5_K_M
wizard-vicuna-uncensored:13b-q6_K
wizard-vicuna-uncensored:13b-q8_0
wizard-vicuna-uncensored:13b-fp16
wizard-vicuna-uncensored:30b-q2_K
wizard-vicuna-uncensored:30b-q3_K_S
wizard-vicuna-uncensored:30b-q3_K_M
wizard-vicuna-uncensored:30b-q3_K_L
wizard-vicuna-uncensored:30b-q4_0
wizard-vicuna-uncensored:30b-q4_1
wizard-vicuna-uncensored:30b-q4_K_S
wizard-vicuna-uncensored:30b-q4_K_M
wizard-vicuna-uncensored:30b-q5_0
wizard-vicuna-uncensored:30b-q5_1
wizard-vicuna-uncensored:30b-q5_K_S
wizard-vicuna-uncensored:30b-q5_K_M
wizard-vicuna-uncensored:30b-q6_K
wizard-vicuna-uncensored:30b-q8_0
wizard-vicuna-uncensored:30b-fp16
opencoder:latest
opencoder
opencoder:1.5b
opencoder:8b
opencoder:1.5b-instruct-q4_K_M
opencoder:1.5b-instruct-q8_0
opencoder:1.5b-instruct-fp16
opencoder:8b-instruct-q4_K_M
opencoder:8b-instruct-q8_0
opencoder:8b-instruct-fp16
starcoder:latest
starcoder
starcoder:1b
starcoder:3b
starcoder:7b
starcoder:15b
starcoder:1b-base
starcoder:1b-base-q2_K
starcoder:1b-base-q3_K_S
starcoder:1b-base-q3_K_M
starcoder:1b-base-q3_K_L
starcoder:1b-base-q4_0
starcoder:1b-base-q4_1
starcoder:1b-base-q4_K_S
starcoder:1b-base-q4_K_M
starcoder:1b-base-q5_0
starcoder:1b-base-q5_1
starcoder:1b-base-q5_K_S
starcoder:1b-base-q5_K_M
starcoder:1b-base-q6_K
starcoder:1b-base-q8_0
starcoder:1b-base-fp16
starcoder:3b-base
starcoder:3b-base-q2_K
starcoder:3b-base-q3_K_S
starcoder:3b-base-q3_K_M
starcoder:3b-base-q3_K_L
starcoder:3b-base-q4_0
starcoder:3b-base-q4_1
starcoder:3b-base-q4_K_S
starcoder:3b-base-q4_K_M
starcoder:3b-base-q5_0
starcoder:3b-base-q5_1
starcoder:3b-base-q5_K_S
starcoder:3b-base-q5_K_M
starcoder:3b-base-q6_K
starcoder:3b-base-q8_0
starcoder:3b-base-fp16
starcoder:7b-base
starcoder:7b-base-q2_K
starcoder:7b-base-q3_K_S
starcoder:7b-base-q3_K_M
starcoder:7b-base-q3_K_L
starcoder:7b-base-q4_0
starcoder:7b-base-q4_1
starcoder:7b-base-q4_K_S
starcoder:7b-base-q4_K_M
starcoder:7b-base-q5_0
starcoder:7b-base-q5_1
starcoder:7b-base-q5_K_S
starcoder:7b-base-q5_K_M
starcoder:7b-base-q6_K
starcoder:7b-base-q8_0
starcoder:7b-base-fp16
starcoder:15b-base
starcoder:15b-base-q2_K
starcoder:15b-base-q3_K_S
starcoder:15b-base-q3_K_M
starcoder:15b-base-q3_K_L
starcoder:15b-base-q4_0
starcoder:15b-base-q4_1
starcoder:15b-base-q4_K_S
starcoder:15b-base-q4_K_M
starcoder:15b-base-q5_0
starcoder:15b-base-q5_1
starcoder:15b-base-q5_K_S
starcoder:15b-base-q5_K_M
starcoder:15b-base-q6_K
starcoder:15b-base-q8_0
starcoder:15b-base-fp16
starcoder:15b-plus
starcoder:15b-plus-q2_K
starcoder:15b-plus-q3_K_S
starcoder:15b-plus-q3_K_M
starcoder:15b-plus-q3_K_L
starcoder:15b-plus-q4_0
starcoder:15b-plus-q4_1
starcoder:15b-plus-q4_K_S
starcoder:15b-plus-q4_K_M
starcoder:15b-plus-q5_0
starcoder:15b-plus-q5_1
starcoder:15b-plus-q5_K_S
starcoder:15b-plus-q5_K_M
starcoder:15b-plus-q6_K
starcoder:15b-plus-q8_0
starcoder:15b-plus-fp16
starcoder:15b-q2_K
starcoder:15b-q3_K_S
starcoder:15b-q3_K_M
starcoder:15b-q3_K_L
starcoder:15b-q4_0
starcoder:15b-q4_1
starcoder:15b-q4_K_S
starcoder:15b-q4_K_M
starcoder:15b-q5_0
starcoder:15b-q5_1
starcoder:15b-q5_K_S
starcoder:15b-q5_K_M
starcoder:15b-q6_K
starcoder:15b-q8_0
starcoder:15b-fp16
nous-hermes:latest
nous-hermes
nous-hermes:7b
nous-hermes:13b
nous-hermes:7b-llama2
nous-hermes:7b-llama2-q2_K
nous-hermes:7b-llama2-q3_K_S
nous-hermes:7b-llama2-q3_K_M
nous-hermes:7b-llama2-q3_K_L
nous-hermes:7b-llama2-q4_0
nous-hermes:7b-llama2-q4_1
nous-hermes:7b-llama2-q4_K_S
nous-hermes:7b-llama2-q4_K_M
nous-hermes:7b-llama2-q5_0
nous-hermes:7b-llama2-q5_1
nous-hermes:7b-llama2-q5_K_S
nous-hermes:7b-llama2-q5_K_M
nous-hermes:7b-llama2-q6_K
nous-hermes:7b-llama2-q8_0
nous-hermes:7b-llama2-fp16
nous-hermes:13b-llama2
nous-hermes:13b-llama2-q2_K
nous-hermes:13b-llama2-q3_K_S
nous-hermes:13b-llama2-q3_K_M
nous-hermes:13b-llama2-q3_K_L
nous-hermes:13b-llama2-q4_0
nous-hermes:13b-llama2-q4_1
nous-hermes:13b-llama2-q4_K_S
nous-hermes:13b-llama2-q4_K_M
nous-hermes:13b-llama2-q5_0
nous-hermes:13b-llama2-q5_1
nous-hermes:13b-llama2-q5_K_S
nous-hermes:13b-llama2-q5_K_M
nous-hermes:13b-llama2-q6_K
nous-hermes:13b-llama2-q8_0
nous-hermes:13b-llama2-fp16
nous-hermes:13b-q2_K
nous-hermes:13b-q3_K_S
nous-hermes:13b-q3_K_M
nous-hermes:13b-q3_K_L
nous-hermes:13b-q4_0
nous-hermes:13b-q4_1
nous-hermes:13b-q4_K_S
nous-hermes:13b-q4_K_M
nous-hermes:13b-q5_0
nous-hermes:13b-q5_1
nous-hermes:13b-q5_K_S
nous-hermes:13b-q5_K_M
nous-hermes:13b-q6_K
nous-hermes:13b-q8_0
nous-hermes:13b-fp16
nous-hermes:70b-llama2-q2_K
nous-hermes:70b-llama2-q3_K_S
nous-hermes:70b-llama2-q3_K_M
nous-hermes:70b-llama2-q3_K_L
nous-hermes:70b-llama2-q4_0
nous-hermes:70b-llama2-q4_1
nous-hermes:70b-llama2-q4_K_S
nous-hermes:70b-llama2-q4_K_M
nous-hermes:70b-llama2-q5_0
nous-hermes:70b-llama2-q5_1
nous-hermes:70b-llama2-q5_K_M
nous-hermes:70b-llama2-q6_K
nous-hermes:70b-llama2-fp16
falcon:latest
falcon
falcon:instruct
falcon:text
falcon:7b
falcon:40b
falcon:180b
falcon:7b-instruct
falcon:7b-instruct-q4_0
falcon:7b-instruct-q4_1
falcon:7b-instruct-q5_0
falcon:7b-instruct-q5_1
falcon:7b-instruct-q8_0
falcon:7b-instruct-fp16
falcon:7b-text
falcon:7b-text-q4_0
falcon:7b-text-q4_1
falcon:7b-text-q5_0
falcon:7b-text-q5_1
falcon:7b-text-q8_0
falcon:7b-text-fp16
falcon:40b-instruct
falcon:40b-instruct-q4_0
falcon:40b-instruct-q4_1
falcon:40b-instruct-q5_0
falcon:40b-instruct-q5_1
falcon:40b-instruct-q8_0
falcon:40b-instruct-fp16
falcon:40b-text
falcon:40b-text-q4_0
falcon:40b-text-q4_1
falcon:40b-text-q5_0
falcon:40b-text-q5_1
falcon:40b-text-q8_0
falcon:40b-text-fp16
falcon:180b-chat
falcon:180b-chat-q4_0
falcon:180b-text
falcon:180b-text-q4_0
deepseek-llm:latest
deepseek-llm
deepseek-llm:7b
deepseek-llm:67b
deepseek-llm:7b-base
deepseek-llm:7b-base-q2_K
deepseek-llm:7b-base-q3_K_S
deepseek-llm:7b-base-q3_K_M
deepseek-llm:7b-base-q3_K_L
deepseek-llm:7b-base-q4_0
deepseek-llm:7b-base-q4_1
deepseek-llm:7b-base-q4_K_S
deepseek-llm:7b-base-q4_K_M
deepseek-llm:7b-base-q5_0
deepseek-llm:7b-base-q5_1
deepseek-llm:7b-base-q5_K_S
deepseek-llm:7b-base-q5_K_M
deepseek-llm:7b-base-q6_K
deepseek-llm:7b-base-q8_0
deepseek-llm:7b-base-fp16
deepseek-llm:7b-chat
deepseek-llm:7b-chat-q2_K
deepseek-llm:7b-chat-q3_K_S
deepseek-llm:7b-chat-q3_K_M
deepseek-llm:7b-chat-q3_K_L
deepseek-llm:7b-chat-q4_0
deepseek-llm:7b-chat-q4_1
deepseek-llm:7b-chat-q4_K_S
deepseek-llm:7b-chat-q4_K_M
deepseek-llm:7b-chat-q5_0
deepseek-llm:7b-chat-q5_1
deepseek-llm:7b-chat-q5_K_S
deepseek-llm:7b-chat-q5_K_M
deepseek-llm:7b-chat-q6_K
deepseek-llm:7b-chat-q8_0
deepseek-llm:7b-chat-fp16
deepseek-llm:67b-base
deepseek-llm:67b-base-q2_K
deepseek-llm:67b-base-q3_K_S
deepseek-llm:67b-base-q3_K_M
deepseek-llm:67b-base-q3_K_L
deepseek-llm:67b-base-q4_0
deepseek-llm:67b-base-q4_1
deepseek-llm:67b-base-q4_K_S
deepseek-llm:67b-base-q4_K_M
deepseek-llm:67b-base-q5_0
deepseek-llm:67b-base-q5_1
deepseek-llm:67b-base-q5_K_S
deepseek-llm:67b-base-q5_K_M
deepseek-llm:67b-base-q6_K
deepseek-llm:67b-base-q8_0
deepseek-llm:67b-base-fp16
deepseek-llm:67b-chat
deepseek-llm:67b-chat-q2_K
deepseek-llm:67b-chat-q3_K_S
deepseek-llm:67b-chat-q3_K_M
deepseek-llm:67b-chat-q3_K_L
deepseek-llm:67b-chat-q4_0
deepseek-llm:67b-chat-q4_1
deepseek-llm:67b-chat-q4_K_S
deepseek-llm:67b-chat-q4_K_M
deepseek-llm:67b-chat-q5_0
deepseek-llm:67b-chat-q5_1
deepseek-llm:67b-chat-q5_K_S
deepseek-llm:67b-chat-fp16
openchat:latest
openchat
openchat:7b
openchat:7b-v3.5
openchat:7b-v3.5-0106
openchat:7b-v3.5-0106-q2_K
openchat:7b-v3.5-q2_K
openchat:7b-v3.5-0106-q3_K_S
openchat:7b-v3.5-q3_K_S
openchat:7b-v3.5-0106-q3_K_M
openchat:7b-v3.5-q3_K_M
openchat:7b-v3.5-0106-q3_K_L
openchat:7b-v3.5-q3_K_L
openchat:7b-v3.5-0106-q4_0
openchat:7b-v3.5-q4_0
openchat:7b-v3.5-0106-q4_1
openchat:7b-v3.5-q4_1
openchat:7b-v3.5-0106-q4_K_S
openchat:7b-v3.5-q4_K_S
openchat:7b-v3.5-0106-q4_K_M
openchat:7b-v3.5-q4_K_M
openchat:7b-v3.5-0106-q5_0
openchat:7b-v3.5-q5_0
openchat:7b-v3.5-0106-q5_1
openchat:7b-v3.5-q5_1
openchat:7b-v3.5-0106-q5_K_S
openchat:7b-v3.5-0106-q5_K_M
openchat:7b-v3.5-0106-q6_K
openchat:7b-v3.5-0106-q8_0
openchat:7b-v3.5-0106-fp16
openchat:7b-v3.5-1210
openchat:7b-v3.5-1210-q2_K
openchat:7b-v3.5-1210-q3_K_S
openchat:7b-v3.5-1210-q3_K_M
openchat:7b-v3.5-1210-q3_K_L
openchat:7b-v3.5-1210-q4_0
openchat:7b-v3.5-1210-q4_1
openchat:7b-v3.5-1210-q4_K_S
openchat:7b-v3.5-1210-q4_K_M
openchat:7b-v3.5-1210-q5_0
openchat:7b-v3.5-1210-q5_1
openchat:7b-v3.5-1210-q5_K_S
openchat:7b-v3.5-q5_K_S
openchat:7b-v3.5-1210-q5_K_M
openchat:7b-v3.5-q5_K_M
openchat:7b-v3.5-1210-q6_K
openchat:7b-v3.5-q6_K
openchat:7b-v3.5-1210-q8_0
openchat:7b-v3.5-q8_0
openchat:7b-v3.5-1210-fp16
openchat:7b-v3.5-fp16
vicuna:latest
vicuna
vicuna:7b
vicuna:13b
vicuna:33b
vicuna:7b-16k
vicuna:7b-v1.5-16k-q2_K
vicuna:7b-v1.5-q2_K
vicuna:7b-v1.5-16k-q3_K_S
vicuna:7b-v1.5-q3_K_S
vicuna:7b-v1.5-16k-q3_K_M
vicuna:7b-v1.5-q3_K_M
vicuna:7b-v1.5-16k-q3_K_L
vicuna:7b-v1.5-q3_K_L
vicuna:7b-v1.5-16k-q4_0
vicuna:7b-v1.5-q4_0
vicuna:7b-v1.5-16k-q4_1
vicuna:7b-v1.5-q4_1
vicuna:7b-v1.5-16k-q4_K_S
vicuna:7b-v1.5-q4_K_S
vicuna:7b-v1.5-16k-q4_K_M
vicuna:7b-v1.5-q4_K_M
vicuna:7b-v1.5-16k-q5_0
vicuna:7b-v1.5-q5_0
vicuna:7b-v1.5-16k-q5_1
vicuna:7b-v1.5-q5_1
vicuna:7b-v1.5-16k-q5_K_S
vicuna:7b-v1.5-q5_K_S
vicuna:7b-v1.5-16k-q5_K_M
vicuna:7b-v1.5-q5_K_M
vicuna:7b-v1.5-16k-q6_K
vicuna:7b-v1.5-q6_K
vicuna:7b-v1.5-16k-q8_0
vicuna:7b-v1.5-q8_0
vicuna:7b-v1.5-16k-fp16
vicuna:7b-v1.5-fp16
vicuna:7b-q2_K
vicuna:7b-q3_K_S
vicuna:7b-q3_K_M
vicuna:7b-q3_K_L
vicuna:7b-q4_0
vicuna:7b-q4_1
vicuna:7b-q4_K_S
vicuna:7b-q4_K_M
vicuna:7b-q5_0
vicuna:7b-q5_1
vicuna:7b-q5_K_S
vicuna:7b-q5_K_M
vicuna:7b-q6_K
vicuna:7b-q8_0
vicuna:7b-fp16
vicuna:13b-16k
vicuna:13b-v1.5-16k-q2_K
vicuna:13b-v1.5-q2_K
vicuna:13b-v1.5-16k-q3_K_S
vicuna:13b-v1.5-q3_K_S
vicuna:13b-v1.5-16k-q3_K_M
vicuna:13b-v1.5-q3_K_M
vicuna:13b-v1.5-16k-q3_K_L
vicuna:13b-v1.5-q3_K_L
vicuna:13b-v1.5-16k-q4_0
vicuna:13b-v1.5-q4_0
vicuna:13b-v1.5-16k-q4_1
vicuna:13b-v1.5-q4_1
vicuna:13b-v1.5-16k-q4_K_S
vicuna:13b-v1.5-q4_K_S
vicuna:13b-v1.5-16k-q4_K_M
vicuna:13b-v1.5-q4_K_M
vicuna:13b-v1.5-16k-q5_0
vicuna:13b-v1.5-q5_0
vicuna:13b-v1.5-16k-q5_1
vicuna:13b-v1.5-q5_1
vicuna:13b-v1.5-16k-q5_K_S
vicuna:13b-v1.5-q5_K_S
vicuna:13b-v1.5-16k-q5_K_M
vicuna:13b-v1.5-q5_K_M
vicuna:13b-v1.5-16k-q6_K
vicuna:13b-v1.5-q6_K
vicuna:13b-v1.5-16k-q8_0
vicuna:13b-v1.5-q8_0
vicuna:13b-v1.5-16k-fp16
vicuna:13b-v1.5-fp16
vicuna:13b-q2_K
vicuna:13b-q3_K_S
vicuna:13b-q3_K_M
vicuna:13b-q3_K_L
vicuna:13b-q4_0
vicuna:13b-q4_1
vicuna:13b-q4_K_S
vicuna:13b-q4_K_M
vicuna:13b-q5_0
vicuna:13b-q5_1
vicuna:13b-q5_K_S
vicuna:13b-q5_K_M
vicuna:13b-q6_K
vicuna:13b-q8_0
vicuna:13b-fp16
vicuna:33b-q2_K
vicuna:33b-q3_K_S
vicuna:33b-q3_K_M
vicuna:33b-q3_K_L
vicuna:33b-q4_0
vicuna:33b-q4_1
vicuna:33b-q4_K_S
vicuna:33b-q4_K_M
vicuna:33b-q5_0
vicuna:33b-q5_1
vicuna:33b-q5_K_S
vicuna:33b-q5_K_M
vicuna:33b-q6_K
vicuna:33b-q8_0
vicuna:33b-fp16
deepseek-v2:latest
deepseek-v2
deepseek-v2:lite
deepseek-v2:16b
deepseek-v2:236b
deepseek-v2:16b-lite-chat-q2_K
deepseek-v2:16b-lite-chat-q3_K_S
deepseek-v2:16b-lite-chat-q3_K_M
deepseek-v2:16b-lite-chat-q3_K_L
deepseek-v2:16b-lite-chat-q4_0
deepseek-v2:16b-lite-chat-q4_1
deepseek-v2:16b-lite-chat-q4_K_S
deepseek-v2:16b-lite-chat-q4_K_M
deepseek-v2:16b-lite-chat-q5_0
deepseek-v2:16b-lite-chat-q5_1
deepseek-v2:16b-lite-chat-q5_K_S
deepseek-v2:16b-lite-chat-q5_K_M
deepseek-v2:16b-lite-chat-q6_K
deepseek-v2:16b-lite-chat-q8_0
deepseek-v2:16b-lite-chat-fp16
deepseek-v2:236b-chat-q2_K
deepseek-v2:236b-chat-q3_K_S
deepseek-v2:236b-chat-q3_K_M
deepseek-v2:236b-chat-q3_K_L
deepseek-v2:236b-chat-q4_0
deepseek-v2:236b-chat-q4_1
deepseek-v2:236b-chat-q4_K_S
deepseek-v2:236b-chat-q4_K_M
deepseek-v2:236b-chat-q5_0
deepseek-v2:236b-chat-q5_1
deepseek-v2:236b-chat-q5_K_S
deepseek-v2:236b-chat-q5_K_M
deepseek-v2:236b-chat-q6_K
deepseek-v2:236b-chat-q8_0
deepseek-v2:236b-chat-fp16
openhermes:latest
openhermes
openhermes:v2
openhermes:v2.5
openhermes:7b-mistral-v2-q2_K
openhermes:7b-mistral-v2-q3_K_S
openhermes:7b-mistral-v2-q3_K_M
openhermes:7b-mistral-v2-q3_K_L
openhermes:7b-mistral-v2-q4_0
openhermes:7b-mistral-v2-q4_1
openhermes:7b-mistral-v2-q4_K_S
openhermes:7b-mistral-v2-q4_K_M
openhermes:7b-mistral-v2-q5_0
openhermes:7b-mistral-v2-q5_1
openhermes:7b-mistral-v2-q5_K_S
openhermes:7b-mistral-v2-q5_K_M
openhermes:7b-mistral-v2-q6_K
openhermes:7b-mistral-v2-q8_0
openhermes:7b-mistral-v2-fp16
openhermes:7b-mistral-v2.5-q2_K
openhermes:7b-mistral-v2.5-q3_K_S
openhermes:7b-mistral-v2.5-q3_K_M
openhermes:7b-mistral-v2.5-q3_K_L
openhermes:7b-mistral-v2.5-q4_0
openhermes:7b-mistral-v2.5-q4_1
openhermes:7b-mistral-v2.5-q4_K_S
openhermes:7b-mistral-v2.5-q4_K_M
openhermes:7b-mistral-v2.5-q5_0
openhermes:7b-mistral-v2.5-q5_1
openhermes:7b-mistral-v2.5-q5_K_S
openhermes:7b-mistral-v2.5-q5_K_M
openhermes:7b-mistral-v2.5-q6_K
openhermes:7b-mistral-v2.5-q8_0
openhermes:7b-mistral-v2.5-fp16
openhermes:7b-v2
openhermes:7b-v2.5
codegeex4:latest
codegeex4
codegeex4:9b
codegeex4:9b-all-q2_K
codegeex4:9b-all-q3_K_S
codegeex4:9b-all-q3_K_M
codegeex4:9b-all-q3_K_L
codegeex4:9b-all-q4_0
codegeex4:9b-all-q4_1
codegeex4:9b-all-q4_K_S
codegeex4:9b-all-q4_K_M
codegeex4:9b-all-q5_0
codegeex4:9b-all-q5_1
codegeex4:9b-all-q5_K_S
codegeex4:9b-all-q5_K_M
codegeex4:9b-all-q6_K
codegeex4:9b-all-q8_0
codegeex4:9b-all-fp16
mistral-openorca:latest
mistral-openorca
mistral-openorca:7b
mistral-openorca:7b-q2_K
mistral-openorca:7b-q3_K_S
mistral-openorca:7b-q3_K_M
mistral-openorca:7b-q3_K_L
mistral-openorca:7b-q4_0
mistral-openorca:7b-q4_1
mistral-openorca:7b-q4_K_S
mistral-openorca:7b-q4_K_M
mistral-openorca:7b-q5_0
mistral-openorca:7b-q5_1
mistral-openorca:7b-q5_K_S
mistral-openorca:7b-q5_K_M
mistral-openorca:7b-q6_K
mistral-openorca:7b-q8_0
mistral-openorca:7b-fp16
deepseek-v3.1:latest
deepseek-v3.1
deepseek-v3.1:671b
deepseek-v3.1:671b-cloud
deepseek-v3.1:671b-terminus-q4_K_M
deepseek-v3.1:671b-terminus-q8_0
deepseek-v3.1:671b-terminus-fp16
deepseek-v3.1:671b-q8_0
deepseek-v3.1:671b-fp16
codeqwen:latest
codeqwen
codeqwen:chat
codeqwen:code
codeqwen:v1.5
codeqwen:7b
codeqwen:7b-chat
codeqwen:7b-chat-v1.5-q2_K
codeqwen:7b-chat-v1.5-q3_K_S
codeqwen:7b-chat-v1.5-q3_K_M
codeqwen:7b-chat-v1.5-q3_K_L
codeqwen:7b-chat-v1.5-q4_0
codeqwen:7b-chat-v1.5-q4_1
codeqwen:7b-chat-v1.5-q4_K_S
codeqwen:7b-chat-v1.5-q4_K_M
codeqwen:7b-chat-v1.5-q5_0
codeqwen:7b-chat-v1.5-q5_1
codeqwen:7b-chat-v1.5-q5_K_S
codeqwen:7b-chat-v1.5-q5_K_M
codeqwen:7b-chat-v1.5-q6_K
codeqwen:7b-chat-v1.5-q8_0
codeqwen:7b-chat-v1.5-fp16
codeqwen:7b-code
codeqwen:7b-code-v1.5-q4_0
codeqwen:7b-code-v1.5-q4_1
codeqwen:7b-code-v1.5-q5_0
codeqwen:7b-code-v1.5-q5_1
codeqwen:7b-code-v1.5-q8_0
codeqwen:7b-code-v1.5-fp16
codeqwen:v1.5-chat
codeqwen:v1.5-code
snowflake-arctic-embed2:latest
snowflake-arctic-embed2
snowflake-arctic-embed2:568m
snowflake-arctic-embed2:568m-l-fp16
qwen3-next:latest
qwen3-next
qwen3-next:80b
qwen3-next:80b-a3b-instruct-q4_K_M
qwen3-next:80b-a3b-instruct-q8_0
qwen3-next:80b-a3b-instruct-fp16
qwen3-next:80b-a3b-thinking
qwen3-next:80b-a3b-thinking-q4_K_M
qwen3-next:80b-a3b-thinking-q8_0
qwen3-next:80b-a3b-thinking-fp16
qwen3-next:80b-cloud
command-r-plus:latest
command-r-plus
command-r-plus:104b
command-r-plus:104b-08-2024-q2_K
command-r-plus:104b-08-2024-q3_K_S
command-r-plus:104b-08-2024-q3_K_M
command-r-plus:104b-08-2024-q3_K_L
command-r-plus:104b-08-2024-q4_0
command-r-plus:104b-08-2024-q4_1
command-r-plus:104b-08-2024-q4_K_S
command-r-plus:104b-08-2024-q4_K_M
command-r-plus:104b-08-2024-q5_0
command-r-plus:104b-08-2024-q5_1
command-r-plus:104b-08-2024-q5_K_S
command-r-plus:104b-08-2024-q5_K_M
command-r-plus:104b-08-2024-q6_K
command-r-plus:104b-08-2024-q8_0
command-r-plus:104b-08-2024-fp16
command-r-plus:104b-q2_K
command-r-plus:104b-q4_0
command-r-plus:104b-q8_0
command-r-plus:104b-fp16
qwen2-math:latest
qwen2-math
qwen2-math:1.5b
qwen2-math:7b
qwen2-math:72b
qwen2-math:1.5b-instruct
qwen2-math:1.5b-instruct-q2_K
qwen2-math:1.5b-instruct-q3_K_S
qwen2-math:1.5b-instruct-q3_K_M
qwen2-math:1.5b-instruct-q3_K_L
qwen2-math:1.5b-instruct-q4_0
qwen2-math:1.5b-instruct-q4_1
qwen2-math:1.5b-instruct-q4_K_S
qwen2-math:1.5b-instruct-q4_K_M
qwen2-math:1.5b-instruct-q5_0
qwen2-math:1.5b-instruct-q5_1
qwen2-math:1.5b-instruct-q5_K_S
qwen2-math:1.5b-instruct-q5_K_M
qwen2-math:1.5b-instruct-q6_K
qwen2-math:1.5b-instruct-q8_0
qwen2-math:1.5b-instruct-fp16
qwen2-math:7b-instruct
qwen2-math:7b-instruct-q2_K
qwen2-math:7b-instruct-q3_K_S
qwen2-math:7b-instruct-q3_K_M
qwen2-math:7b-instruct-q3_K_L
qwen2-math:7b-instruct-q4_0
qwen2-math:7b-instruct-q4_1
qwen2-math:7b-instruct-q4_K_S
qwen2-math:7b-instruct-q4_K_M
qwen2-math:7b-instruct-q5_0
qwen2-math:7b-instruct-q5_1
qwen2-math:7b-instruct-q5_K_S
qwen2-math:7b-instruct-q5_K_M
qwen2-math:7b-instruct-q6_K
qwen2-math:7b-instruct-q8_0
qwen2-math:7b-instruct-fp16
qwen2-math:72b-instruct
qwen2-math:72b-instruct-q2_K
qwen2-math:72b-instruct-q3_K_S
qwen2-math:72b-instruct-q3_K_M
qwen2-math:72b-instruct-q3_K_L
qwen2-math:72b-instruct-q4_0
qwen2-math:72b-instruct-q4_1
qwen2-math:72b-instruct-q4_K_S
qwen2-math:72b-instruct-q4_K_M
qwen2-math:72b-instruct-q5_0
qwen2-math:72b-instruct-q5_1
qwen2-math:72b-instruct-q5_K_S
qwen2-math:72b-instruct-q5_K_M
qwen2-math:72b-instruct-q6_K
qwen2-math:72b-instruct-q8_0
qwen2-math:72b-instruct-fp16
qwen3-embedding:latest
qwen3-embedding
qwen3-embedding:0.6b
qwen3-embedding:4b
qwen3-embedding:8b
qwen3-embedding:0.6b-q8_0
qwen3-embedding:0.6b-fp16
qwen3-embedding:4b-q4_K_M
qwen3-embedding:4b-q8_0
qwen3-embedding:4b-fp16
qwen3-embedding:8b-q4_K_M
qwen3-embedding:8b-q8_0
qwen3-embedding:8b-fp16
tinydolphin:latest
tinydolphin
tinydolphin:v2.8
tinydolphin:1.1b
tinydolphin:1.1b-v2.8-q2_K
tinydolphin:1.1b-v2.8-q3_K_S
tinydolphin:1.1b-v2.8-q3_K_M
tinydolphin:1.1b-v2.8-q3_K_L
tinydolphin:1.1b-v2.8-q4_0
tinydolphin:1.1b-v2.8-q4_1
tinydolphin:1.1b-v2.8-q4_K_S
tinydolphin:1.1b-v2.8-q4_K_M
tinydolphin:1.1b-v2.8-q5_0
tinydolphin:1.1b-v2.8-q5_1
tinydolphin:1.1b-v2.8-q5_K_S
tinydolphin:1.1b-v2.8-q5_K_M
tinydolphin:1.1b-v2.8-q6_K
tinydolphin:1.1b-v2.8-q8_0
tinydolphin:1.1b-v2.8-fp16
aya:latest
aya
aya:8b
aya:35b
aya:8b-23
aya:8b-23-q2_K
aya:8b-23-q3_K_S
aya:8b-23-q3_K_M
aya:8b-23-q3_K_L
aya:8b-23-q4_0
aya:8b-23-q4_1
aya:8b-23-q4_K_S
aya:8b-23-q4_K_M
aya:8b-23-q5_0
aya:8b-23-q5_1
aya:8b-23-q5_K_S
aya:8b-23-q5_K_M
aya:8b-23-q6_K
aya:8b-23-q8_0
aya:35b-23
aya:35b-23-q2_K
aya:35b-23-q3_K_S
aya:35b-23-q3_K_M
aya:35b-23-q3_K_L
aya:35b-23-q4_0
aya:35b-23-q4_1
aya:35b-23-q4_K_S
aya:35b-23-q4_K_M
aya:35b-23-q5_0
aya:35b-23-q5_1
aya:35b-23-q5_K_S
aya:35b-23-q5_K_M
aya:35b-23-q6_K
aya:35b-23-q8_0
glm4:latest
glm4
glm4:9b
glm4:9b-chat-q2_K
glm4:9b-chat-q3_K_S
glm4:9b-chat-q3_K_M
glm4:9b-chat-q3_K_L
glm4:9b-chat-q4_0
glm4:9b-chat-q4_1
glm4:9b-chat-q4_K_S
glm4:9b-chat-q4_K_M
glm4:9b-chat-q5_0
glm4:9b-chat-q5_1
glm4:9b-chat-q5_K_S
glm4:9b-chat-q5_K_M
glm4:9b-chat-q6_K
glm4:9b-chat-q8_0
glm4:9b-chat-fp16
glm4:9b-text-q2_K
glm4:9b-text-q3_K_S
glm4:9b-text-q3_K_M
glm4:9b-text-q3_K_L
glm4:9b-text-q4_0
glm4:9b-text-q4_1
glm4:9b-text-q4_K_S
glm4:9b-text-q4_K_M
glm4:9b-text-q5_0
glm4:9b-text-q5_1
glm4:9b-text-q5_K_S
glm4:9b-text-q5_K_M
glm4:9b-text-q6_K
glm4:9b-text-q8_0
glm4:9b-text-fp16
llama2-chinese:latest
llama2-chinese
llama2-chinese:7b
llama2-chinese:13b
llama2-chinese:7b-chat
llama2-chinese:7b-chat-q2_K
llama2-chinese:7b-chat-q3_K_S
llama2-chinese:7b-chat-q3_K_M
llama2-chinese:7b-chat-q3_K_L
llama2-chinese:7b-chat-q4_0
llama2-chinese:7b-chat-q4_1
llama2-chinese:7b-chat-q4_K_S
llama2-chinese:7b-chat-q4_K_M
llama2-chinese:7b-chat-q5_0
llama2-chinese:7b-chat-q5_1
llama2-chinese:7b-chat-q5_K_S
llama2-chinese:7b-chat-q5_K_M
llama2-chinese:7b-chat-q6_K
llama2-chinese:7b-chat-q8_0
llama2-chinese:7b-chat-fp16
llama2-chinese:13b-chat
llama2-chinese:13b-chat-q2_K
llama2-chinese:13b-chat-q3_K_S
llama2-chinese:13b-chat-q3_K_M
llama2-chinese:13b-chat-q3_K_L
llama2-chinese:13b-chat-q4_0
llama2-chinese:13b-chat-q4_1
llama2-chinese:13b-chat-q4_K_S
llama2-chinese:13b-chat-q4_K_M
llama2-chinese:13b-chat-q5_0
llama2-chinese:13b-chat-q5_1
llama2-chinese:13b-chat-q5_K_S
llama2-chinese:13b-chat-q5_K_M
llama2-chinese:13b-chat-q6_K
llama2-chinese:13b-chat-q8_0
llama2-chinese:13b-chat-fp16
granite3.2:latest
granite3.2
granite3.2:2b
granite3.2:8b
granite3.2:2b-instruct-q4_K_M
granite3.2:2b-instruct-q8_0
granite3.2:2b-instruct-fp16
granite3.2:8b-instruct-q4_K_M
granite3.2:8b-instruct-q8_0
granite3.2:8b-instruct-fp16
paraphrase-multilingual:latest
paraphrase-multilingual
paraphrase-multilingual:278m
paraphrase-multilingual:278m-mpnet-base-v2-fp16
stable-code:latest
stable-code
stable-code:code
stable-code:instruct
stable-code:3b
stable-code:3b-code
stable-code:3b-code-q2_K
stable-code:3b-code-q3_K_S
stable-code:3b-code-q3_K_M
stable-code:3b-code-q3_K_L
stable-code:3b-code-q4_0
stable-code:3b-code-q4_1
stable-code:3b-code-q4_K_S
stable-code:3b-code-q4_K_M
stable-code:3b-code-q5_0
stable-code:3b-code-q5_1
stable-code:3b-code-q5_K_S
stable-code:3b-code-q5_K_M
stable-code:3b-code-q6_K
stable-code:3b-code-q8_0
stable-code:3b-code-fp16
stable-code:3b-instruct
stable-code:3b-instruct-q2_K
stable-code:3b-instruct-q3_K_S
stable-code:3b-instruct-q3_K_M
stable-code:3b-instruct-q3_K_L
stable-code:3b-instruct-q4_0
stable-code:3b-instruct-q4_1
stable-code:3b-instruct-q4_K_S
stable-code:3b-instruct-q4_K_M
stable-code:3b-instruct-q5_0
stable-code:3b-instruct-q5_1
stable-code:3b-instruct-q5_K_S
stable-code:3b-instruct-q5_K_M
stable-code:3b-instruct-q6_K
stable-code:3b-instruct-q8_0
stable-code:3b-instruct-fp16
neural-chat:latest
neural-chat
neural-chat:7b
neural-chat:7b-v3.1
neural-chat:7b-v3.1-q2_K
neural-chat:7b-v3.1-q3_K_S
neural-chat:7b-v3.1-q3_K_M
neural-chat:7b-v3.1-q3_K_L
neural-chat:7b-v3.1-q4_0
neural-chat:7b-v3.1-q4_1
neural-chat:7b-v3.1-q4_K_S
neural-chat:7b-v3.1-q4_K_M
neural-chat:7b-v3.1-q5_0
neural-chat:7b-v3.1-q5_1
neural-chat:7b-v3.1-q5_K_S
neural-chat:7b-v3.1-q5_K_M
neural-chat:7b-v3.1-q6_K
neural-chat:7b-v3.1-q8_0
neural-chat:7b-v3.1-fp16
neural-chat:7b-v3.2
neural-chat:7b-v3.2-q2_K
neural-chat:7b-v3.2-q3_K_S
neural-chat:7b-v3.2-q3_K_M
neural-chat:7b-v3.2-q3_K_L
neural-chat:7b-v3.2-q4_0
neural-chat:7b-v3.2-q4_1
neural-chat:7b-v3.2-q4_K_S
neural-chat:7b-v3.2-q4_K_M
neural-chat:7b-v3.2-q5_0
neural-chat:7b-v3.2-q5_1
neural-chat:7b-v3.2-q5_K_S
neural-chat:7b-v3.2-q5_K_M
neural-chat:7b-v3.2-q6_K
neural-chat:7b-v3.2-q8_0
neural-chat:7b-v3.2-fp16
neural-chat:7b-v3.3
neural-chat:7b-v3.3-q2_K
neural-chat:7b-v3.3-q3_K_S
neural-chat:7b-v3.3-q3_K_M
neural-chat:7b-v3.3-q3_K_L
neural-chat:7b-v3.3-q4_0
neural-chat:7b-v3.3-q4_1
neural-chat:7b-v3.3-q4_K_S
neural-chat:7b-v3.3-q4_K_M
neural-chat:7b-v3.3-q5_0
neural-chat:7b-v3.3-q5_1
neural-chat:7b-v3.3-q5_K_S
neural-chat:7b-v3.3-q5_K_M
neural-chat:7b-v3.3-q6_K
neural-chat:7b-v3.3-q8_0
neural-chat:7b-v3.3-fp16
nous-hermes2:latest
nous-hermes2
nous-hermes2:10.7b
nous-hermes2:34b
nous-hermes2:10.7b-solar-q2_K
nous-hermes2:10.7b-solar-q3_K_S
nous-hermes2:10.7b-solar-q3_K_M
nous-hermes2:10.7b-solar-q3_K_L
nous-hermes2:10.7b-solar-q4_0
nous-hermes2:10.7b-solar-q4_1
nous-hermes2:10.7b-solar-q4_K_S
nous-hermes2:10.7b-solar-q4_K_M
nous-hermes2:10.7b-solar-q5_0
nous-hermes2:10.7b-solar-q5_1
nous-hermes2:10.7b-solar-q5_K_S
nous-hermes2:10.7b-solar-q5_K_M
nous-hermes2:10.7b-solar-q6_K
nous-hermes2:10.7b-solar-q8_0
nous-hermes2:10.7b-solar-fp16
nous-hermes2:34b-yi-q2_K
nous-hermes2:34b-yi-q3_K_S
nous-hermes2:34b-yi-q3_K_M
nous-hermes2:34b-yi-q3_K_L
nous-hermes2:34b-yi-q4_0
nous-hermes2:34b-yi-q4_1
nous-hermes2:34b-yi-q4_K_S
nous-hermes2:34b-yi-q4_K_M
nous-hermes2:34b-yi-q5_0
nous-hermes2:34b-yi-q5_1
nous-hermes2:34b-yi-q5_K_S
nous-hermes2:34b-yi-q5_K_M
nous-hermes2:34b-yi-q6_K
nous-hermes2:34b-yi-q8_0
nous-hermes2:34b-yi-fp16
bakllava:latest
bakllava
bakllava:7b
bakllava:7b-v1-q2_K
bakllava:7b-v1-q3_K_S
bakllava:7b-v1-q3_K_M
bakllava:7b-v1-q3_K_L
bakllava:7b-v1-q4_0
bakllava:7b-v1-q4_1
bakllava:7b-v1-q4_K_S
bakllava:7b-v1-q4_K_M
bakllava:7b-v1-q5_0
bakllava:7b-v1-q5_1
bakllava:7b-v1-q5_K_S
bakllava:7b-v1-q5_K_M
bakllava:7b-v1-q6_K
bakllava:7b-v1-q8_0
bakllava:7b-v1-fp16
wizardcoder:latest
wizardcoder
wizardcoder:python
wizardcoder:33b
wizardcoder:7b-python
wizardcoder:7b-python-q2_K
wizardcoder:7b-python-q3_K_S
wizardcoder:7b-python-q3_K_M
wizardcoder:7b-python-q3_K_L
wizardcoder:7b-python-q4_0
wizardcoder:7b-python-q4_1
wizardcoder:7b-python-q4_K_S
wizardcoder:7b-python-q4_K_M
wizardcoder:7b-python-q5_0
wizardcoder:7b-python-q5_1
wizardcoder:7b-python-q5_K_S
wizardcoder:7b-python-q5_K_M
wizardcoder:7b-python-q6_K
wizardcoder:7b-python-q8_0
wizardcoder:7b-python-fp16
wizardcoder:13b-python
wizardcoder:13b-python-q2_K
wizardcoder:13b-python-q3_K_S
wizardcoder:13b-python-q3_K_M
wizardcoder:13b-python-q3_K_L
wizardcoder:13b-python-q4_0
wizardcoder:13b-python-q4_1
wizardcoder:13b-python-q4_K_S
wizardcoder:13b-python-q4_K_M
wizardcoder:13b-python-q5_0
wizardcoder:13b-python-q5_1
wizardcoder:13b-python-q5_K_S
wizardcoder:13b-python-q5_K_M
wizardcoder:13b-python-q6_K
wizardcoder:13b-python-q8_0
wizardcoder:13b-python-fp16
wizardcoder:33b-v1.1
wizardcoder:33b-v1.1-q2_K
wizardcoder:33b-v1.1-q3_K_S
wizardcoder:33b-v1.1-q3_K_M
wizardcoder:33b-v1.1-q3_K_L
wizardcoder:33b-v1.1-q4_0
wizardcoder:33b-v1.1-q4_1
wizardcoder:33b-v1.1-q4_K_S
wizardcoder:33b-v1.1-q4_K_M
wizardcoder:33b-v1.1-q5_0
wizardcoder:33b-v1.1-q5_1
wizardcoder:33b-v1.1-q5_K_S
wizardcoder:33b-v1.1-q5_K_M
wizardcoder:33b-v1.1-q6_K
wizardcoder:33b-v1.1-q8_0
wizardcoder:33b-v1.1-fp16
wizardcoder:34b-python
wizardcoder:34b-python-q2_K
wizardcoder:34b-python-q3_K_S
wizardcoder:34b-python-q3_K_M
wizardcoder:34b-python-q3_K_L
wizardcoder:34b-python-q4_0
wizardcoder:34b-python-q4_1
wizardcoder:34b-python-q4_K_S
wizardcoder:34b-python-q4_K_M
wizardcoder:34b-python-q5_0
wizardcoder:34b-python-q5_1
wizardcoder:34b-python-q5_K_S
wizardcoder:34b-python-q5_K_M
wizardcoder:34b-python-q6_K
wizardcoder:34b-python-q8_0
wizardcoder:34b-python-fp16
sqlcoder:latest
sqlcoder
sqlcoder:7b
sqlcoder:15b
sqlcoder:7b-q2_K
sqlcoder:7b-q3_K_S
sqlcoder:7b-q3_K_M
sqlcoder:7b-q3_K_L
sqlcoder:7b-q4_0
sqlcoder:7b-q4_1
sqlcoder:7b-q4_K_S
sqlcoder:7b-q4_K_M
sqlcoder:7b-q5_0
sqlcoder:7b-q5_1
sqlcoder:7b-q5_K_S
sqlcoder:7b-q5_K_M
sqlcoder:7b-q6_K
sqlcoder:7b-q8_0
sqlcoder:7b-fp16
sqlcoder:15b-q2_K
sqlcoder:15b-q3_K_S
sqlcoder:15b-q3_K_M
sqlcoder:15b-q3_K_L
sqlcoder:15b-q4_0
sqlcoder:15b-q4_1
sqlcoder:15b-q4_K_S
sqlcoder:15b-q4_K_M
sqlcoder:15b-q5_0
sqlcoder:15b-q5_1
sqlcoder:15b-q5_K_S
sqlcoder:15b-q5_K_M
sqlcoder:15b-q6_K
sqlcoder:15b-q8_0
sqlcoder:15b-fp16
sqlcoder:70b-alpha-q2_K
sqlcoder:70b-alpha-q3_K_S
sqlcoder:70b-alpha-q3_K_M
sqlcoder:70b-alpha-q3_K_L
sqlcoder:70b-alpha-q4_0
sqlcoder:70b-alpha-q4_1
sqlcoder:70b-alpha-q4_K_S
sqlcoder:70b-alpha-q4_K_M
sqlcoder:70b-alpha-q5_0
sqlcoder:70b-alpha-q5_1
sqlcoder:70b-alpha-q5_K_S
sqlcoder:70b-alpha-q5_K_M
sqlcoder:70b-alpha-q6_K
sqlcoder:70b-alpha-q8_0
sqlcoder:70b-alpha-fp16
bge-large:latest
bge-large
bge-large:335m
bge-large:335m-en-v1.5-fp16
stablelm2:latest
stablelm2
stablelm2:chat
stablelm2:zephyr
stablelm2:1.6b
stablelm2:12b
stablelm2:1.6b-chat
stablelm2:1.6b-chat-q2_K
stablelm2:1.6b-chat-q3_K_S
stablelm2:1.6b-chat-q3_K_M
stablelm2:1.6b-chat-q3_K_L
stablelm2:1.6b-chat-q4_0
stablelm2:1.6b-chat-q4_1
stablelm2:1.6b-chat-q4_K_S
stablelm2:1.6b-chat-q4_K_M
stablelm2:1.6b-chat-q5_0
stablelm2:1.6b-chat-q5_1
stablelm2:1.6b-chat-q5_K_S
stablelm2:1.6b-chat-q5_K_M
stablelm2:1.6b-chat-q6_K
stablelm2:1.6b-chat-q8_0
stablelm2:1.6b-chat-fp16
stablelm2:1.6b-zephyr
stablelm2:1.6b-zephyr-q2_K
stablelm2:1.6b-zephyr-q3_K_S
stablelm2:1.6b-zephyr-q3_K_M
stablelm2:1.6b-zephyr-q3_K_L
stablelm2:1.6b-zephyr-q4_0
stablelm2:1.6b-zephyr-q4_1
stablelm2:1.6b-zephyr-q4_K_S
stablelm2:1.6b-zephyr-q4_K_M
stablelm2:1.6b-zephyr-q5_0
stablelm2:1.6b-zephyr-q5_1
stablelm2:1.6b-zephyr-q5_K_S
stablelm2:1.6b-zephyr-q5_K_M
stablelm2:1.6b-zephyr-q6_K
stablelm2:1.6b-zephyr-q8_0
stablelm2:1.6b-zephyr-fp16
stablelm2:1.6b-q2_K
stablelm2:1.6b-q3_K_S
stablelm2:1.6b-q3_K_M
stablelm2:1.6b-q3_K_L
stablelm2:1.6b-q4_0
stablelm2:1.6b-q4_1
stablelm2:1.6b-q4_K_S
stablelm2:1.6b-q4_K_M
stablelm2:1.6b-q5_0
stablelm2:1.6b-q5_1
stablelm2:1.6b-q5_K_S
stablelm2:1.6b-q5_K_M
stablelm2:1.6b-q6_K
stablelm2:1.6b-q8_0
stablelm2:1.6b-fp16
stablelm2:12b-chat
stablelm2:12b-chat-q2_K
stablelm2:12b-chat-q3_K_S
stablelm2:12b-chat-q3_K_M
stablelm2:12b-chat-q3_K_L
stablelm2:12b-chat-q4_0
stablelm2:12b-chat-q4_1
stablelm2:12b-chat-q4_K_S
stablelm2:12b-chat-q4_K_M
stablelm2:12b-chat-q5_0
stablelm2:12b-chat-q5_1
stablelm2:12b-chat-q5_K_S
stablelm2:12b-chat-q5_K_M
stablelm2:12b-chat-q6_K
stablelm2:12b-chat-q8_0
stablelm2:12b-chat-fp16
stablelm2:12b-text
stablelm2:12b-q2_K
stablelm2:12b-q3_K_S
stablelm2:12b-q3_K_M
stablelm2:12b-q3_K_L
stablelm2:12b-q4_0
stablelm2:12b-q4_1
stablelm2:12b-q4_K_S
stablelm2:12b-q4_K_M
stablelm2:12b-q5_0
stablelm2:12b-q5_1
stablelm2:12b-q5_K_S
stablelm2:12b-q5_K_M
stablelm2:12b-q6_K
stablelm2:12b-q8_0
stablelm2:12b-fp16
r1-1776:latest
r1-1776
r1-1776:70b
r1-1776:671b
r1-1776:70b-distill-llama-q4_K_M
r1-1776:70b-distill-llama-q8_0
r1-1776:70b-distill-llama-fp16
r1-1776:671b-q4_K_M
r1-1776:671b-q8_0
r1-1776:671b-fp16
yi-coder:latest
yi-coder
yi-coder:1.5b
yi-coder:9b
yi-coder:1.5b-base
yi-coder:1.5b-base-q2_K
yi-coder:1.5b-base-q3_K_S
yi-coder:1.5b-base-q3_K_M
yi-coder:1.5b-base-q3_K_L
yi-coder:1.5b-base-q4_0
yi-coder:1.5b-base-q4_1
yi-coder:1.5b-base-q4_K_S
yi-coder:1.5b-base-q4_K_M
yi-coder:1.5b-base-q5_0
yi-coder:1.5b-base-q5_1
yi-coder:1.5b-base-q5_K_S
yi-coder:1.5b-base-q5_K_M
yi-coder:1.5b-base-q6_K
yi-coder:1.5b-base-q8_0
yi-coder:1.5b-base-fp16
yi-coder:1.5b-chat
yi-coder:1.5b-chat-q2_K
yi-coder:1.5b-chat-q3_K_S
yi-coder:1.5b-chat-q3_K_M
yi-coder:1.5b-chat-q3_K_L
yi-coder:1.5b-chat-q4_0
yi-coder:1.5b-chat-q4_1
yi-coder:1.5b-chat-q4_K_S
yi-coder:1.5b-chat-q4_K_M
yi-coder:1.5b-chat-q5_0
yi-coder:1.5b-chat-q5_1
yi-coder:1.5b-chat-q5_K_S
yi-coder:1.5b-chat-q5_K_M
yi-coder:1.5b-chat-q6_K
yi-coder:1.5b-chat-q8_0
yi-coder:1.5b-chat-fp16
yi-coder:9b-base
yi-coder:9b-base-q2_K
yi-coder:9b-base-q3_K_S
yi-coder:9b-base-q3_K_M
yi-coder:9b-base-q3_K_L
yi-coder:9b-base-q4_0
yi-coder:9b-base-q4_1
yi-coder:9b-base-q4_K_S
yi-coder:9b-base-q4_K_M
yi-coder:9b-base-q5_0
yi-coder:9b-base-q5_1
yi-coder:9b-base-q5_K_S
yi-coder:9b-base-q5_K_M
yi-coder:9b-base-q6_K
yi-coder:9b-base-q8_0
yi-coder:9b-base-fp16
yi-coder:9b-chat
yi-coder:9b-chat-q2_K
yi-coder:9b-chat-q3_K_S
yi-coder:9b-chat-q3_K_M
yi-coder:9b-chat-q3_K_L
yi-coder:9b-chat-q4_0
yi-coder:9b-chat-q4_1
yi-coder:9b-chat-q4_K_S
yi-coder:9b-chat-q4_K_M
yi-coder:9b-chat-q5_0
yi-coder:9b-chat-q5_1
yi-coder:9b-chat-q5_K_S
yi-coder:9b-chat-q5_K_M
yi-coder:9b-chat-q6_K
yi-coder:9b-chat-q8_0
yi-coder:9b-chat-fp16
llava-phi3:latest
llava-phi3
llava-phi3:3.8b
llava-phi3:3.8b-mini-q4_0
llava-phi3:3.8b-mini-fp16
llama3-chatqa:latest
llama3-chatqa
llama3-chatqa:8b
llama3-chatqa:70b
llama3-chatqa:8b-v1.5
llama3-chatqa:8b-v1.5-q2_K
llama3-chatqa:8b-v1.5-q3_K_S
llama3-chatqa:8b-v1.5-q3_K_M
llama3-chatqa:8b-v1.5-q3_K_L
llama3-chatqa:8b-v1.5-q4_0
llama3-chatqa:8b-v1.5-q4_1
llama3-chatqa:8b-v1.5-q4_K_S
llama3-chatqa:8b-v1.5-q4_K_M
llama3-chatqa:8b-v1.5-q5_0
llama3-chatqa:8b-v1.5-q5_1
llama3-chatqa:8b-v1.5-q5_K_S
llama3-chatqa:8b-v1.5-q5_K_M
llama3-chatqa:8b-v1.5-q6_K
llama3-chatqa:8b-v1.5-q8_0
llama3-chatqa:8b-v1.5-fp16
llama3-chatqa:70b-v1.5
llama3-chatqa:70b-v1.5-q2_K
llama3-chatqa:70b-v1.5-q3_K_S
llama3-chatqa:70b-v1.5-q3_K_M
llama3-chatqa:70b-v1.5-q3_K_L
llama3-chatqa:70b-v1.5-q4_0
llama3-chatqa:70b-v1.5-q4_1
llama3-chatqa:70b-v1.5-q4_K_S
llama3-chatqa:70b-v1.5-q4_K_M
llama3-chatqa:70b-v1.5-q5_0
llama3-chatqa:70b-v1.5-q5_1
llama3-chatqa:70b-v1.5-q5_K_S
llama3-chatqa:70b-v1.5-q5_K_M
llama3-chatqa:70b-v1.5-q6_K
llama3-chatqa:70b-v1.5-q8_0
llama3-chatqa:70b-v1.5-fp16
granite3-dense:latest
granite3-dense
granite3-dense:2b
granite3-dense:8b
granite3-dense:2b-instruct-q2_K
granite3-dense:2b-instruct-q3_K_S
granite3-dense:2b-instruct-q3_K_M
granite3-dense:2b-instruct-q3_K_L
granite3-dense:2b-instruct-q4_0
granite3-dense:2b-instruct-q4_1
granite3-dense:2b-instruct-q4_K_S
granite3-dense:2b-instruct-q4_K_M
granite3-dense:2b-instruct-q5_0
granite3-dense:2b-instruct-q5_1
granite3-dense:2b-instruct-q5_K_S
granite3-dense:2b-instruct-q5_K_M
granite3-dense:2b-instruct-q6_K
granite3-dense:2b-instruct-q8_0
granite3-dense:2b-instruct-fp16
granite3-dense:8b-instruct-q2_K
granite3-dense:8b-instruct-q3_K_S
granite3-dense:8b-instruct-q3_K_M
granite3-dense:8b-instruct-q3_K_L
granite3-dense:8b-instruct-q4_0
granite3-dense:8b-instruct-q4_1
granite3-dense:8b-instruct-q4_K_S
granite3-dense:8b-instruct-q4_K_M
granite3-dense:8b-instruct-q5_0
granite3-dense:8b-instruct-q5_1
granite3-dense:8b-instruct-q5_K_S
granite3-dense:8b-instruct-q5_K_M
granite3-dense:8b-instruct-q6_K
granite3-dense:8b-instruct-q8_0
granite3-dense:8b-instruct-fp16
granite3.1-dense:latest
granite3.1-dense
granite3.1-dense:2b
granite3.1-dense:8b
granite3.1-dense:2b-instruct-q2_K
granite3.1-dense:2b-instruct-q3_K_S
granite3.1-dense:2b-instruct-q3_K_M
granite3.1-dense:2b-instruct-q3_K_L
granite3.1-dense:2b-instruct-q4_0
granite3.1-dense:2b-instruct-q4_1
granite3.1-dense:2b-instruct-q4_K_S
granite3.1-dense:2b-instruct-q4_K_M
granite3.1-dense:2b-instruct-q5_0
granite3.1-dense:2b-instruct-q5_1
granite3.1-dense:2b-instruct-q5_K_S
granite3.1-dense:2b-instruct-q5_K_M
granite3.1-dense:2b-instruct-q6_K
granite3.1-dense:2b-instruct-q8_0
granite3.1-dense:2b-instruct-fp16
granite3.1-dense:8b-instruct-q2_K
granite3.1-dense:8b-instruct-q3_K_S
granite3.1-dense:8b-instruct-q3_K_M
granite3.1-dense:8b-instruct-q3_K_L
granite3.1-dense:8b-instruct-q4_0
granite3.1-dense:8b-instruct-q4_1
granite3.1-dense:8b-instruct-q4_K_S
granite3.1-dense:8b-instruct-q4_K_M
granite3.1-dense:8b-instruct-q5_0
granite3.1-dense:8b-instruct-q5_1
granite3.1-dense:8b-instruct-q5_K_S
granite3.1-dense:8b-instruct-q5_K_M
granite3.1-dense:8b-instruct-q6_K
granite3.1-dense:8b-instruct-q8_0
granite3.1-dense:8b-instruct-fp16
exaone3.5:latest
exaone3.5
exaone3.5:2.4b
exaone3.5:7.8b
exaone3.5:32b
exaone3.5:2.4b-instruct-q4_K_M
exaone3.5:2.4b-instruct-q8_0
exaone3.5:2.4b-instruct-fp16
exaone3.5:7.8b-instruct-q4_K_M
exaone3.5:7.8b-instruct-q8_0
exaone3.5:7.8b-instruct-fp16
exaone3.5:32b-instruct-q4_K_M
exaone3.5:32b-instruct-q8_0
exaone3.5:32b-instruct-fp16
granite-embedding:latest
granite-embedding
granite-embedding:30m
granite-embedding:278m
granite-embedding:30m-en
granite-embedding:30m-en-fp16
granite-embedding:278m-fp16
reflection:latest
reflection
reflection:70b
reflection:70b-q2_K
reflection:70b-q3_K_S
reflection:70b-q3_K_M
reflection:70b-q3_K_L
reflection:70b-q4_0
reflection:70b-q4_1
reflection:70b-q4_K_S
reflection:70b-q4_K_M
reflection:70b-q5_0
reflection:70b-q5_1
reflection:70b-q5_K_S
reflection:70b-q5_K_M
reflection:70b-q6_K
reflection:70b-q8_0
reflection:70b-fp16
wizard-math:latest
wizard-math
wizard-math:7b
wizard-math:13b
wizard-math:70b
wizard-math:7b-v1.1-q2_K
wizard-math:7b-v1.1-q3_K_S
wizard-math:7b-v1.1-q3_K_M
wizard-math:7b-v1.1-q3_K_L
wizard-math:7b-v1.1-q4_0
wizard-math:7b-v1.1-q4_1
wizard-math:7b-v1.1-q4_K_S
wizard-math:7b-v1.1-q4_K_M
wizard-math:7b-v1.1-q5_0
wizard-math:7b-v1.1-q5_1
wizard-math:7b-v1.1-q5_K_S
wizard-math:7b-v1.1-q5_K_M
wizard-math:7b-v1.1-q6_K
wizard-math:7b-v1.1-q8_0
wizard-math:7b-v1.1-fp16
wizard-math:7b-q2_K
wizard-math:7b-q3_K_S
wizard-math:7b-q3_K_M
wizard-math:7b-q3_K_L
wizard-math:7b-q4_0
wizard-math:7b-q4_1
wizard-math:7b-q4_K_S
wizard-math:7b-q4_K_M
wizard-math:7b-q5_0
wizard-math:7b-q5_1
wizard-math:7b-q5_K_S
wizard-math:7b-q5_K_M
wizard-math:7b-q6_K
wizard-math:7b-q8_0
wizard-math:7b-fp16
wizard-math:13b-q2_K
wizard-math:13b-q3_K_S
wizard-math:13b-q3_K_M
wizard-math:13b-q3_K_L
wizard-math:13b-q4_0
wizard-math:13b-q4_1
wizard-math:13b-q4_K_S
wizard-math:13b-q4_K_M
wizard-math:13b-q5_0
wizard-math:13b-q5_1
wizard-math:13b-q5_K_S
wizard-math:13b-q5_K_M
wizard-math:13b-q6_K
wizard-math:13b-q8_0
wizard-math:13b-fp16
wizard-math:70b-q2_K
wizard-math:70b-q3_K_S
wizard-math:70b-q3_K_M
wizard-math:70b-q3_K_L
wizard-math:70b-q4_0
wizard-math:70b-q4_1
wizard-math:70b-q4_K_S
wizard-math:70b-q4_K_M
wizard-math:70b-q5_0
wizard-math:70b-q5_1
wizard-math:70b-q5_K_S
wizard-math:70b-q5_K_M
wizard-math:70b-q6_K
wizard-math:70b-q8_0
wizard-math:70b-fp16
llama3-gradient:latest
llama3-gradient
llama3-gradient:instruct
llama3-gradient:1048k
llama3-gradient:8b
llama3-gradient:70b
llama3-gradient:8b-instruct-1048k-q2_K
llama3-gradient:8b-instruct-1048k-q3_K_S
llama3-gradient:8b-instruct-1048k-q3_K_M
llama3-gradient:8b-instruct-1048k-q3_K_L
llama3-gradient:8b-instruct-1048k-q4_0
llama3-gradient:8b-instruct-1048k-q4_1
llama3-gradient:8b-instruct-1048k-q4_K_S
llama3-gradient:8b-instruct-1048k-q4_K_M
llama3-gradient:8b-instruct-1048k-q5_0
llama3-gradient:8b-instruct-1048k-q5_1
llama3-gradient:8b-instruct-1048k-q5_K_S
llama3-gradient:8b-instruct-1048k-q5_K_M
llama3-gradient:8b-instruct-1048k-q6_K
llama3-gradient:8b-instruct-1048k-q8_0
llama3-gradient:8b-instruct-1048k-fp16
llama3-gradient:70b-instruct-1048k-q2_K
llama3-gradient:70b-instruct-1048k-q3_K_S
llama3-gradient:70b-instruct-1048k-q3_K_M
llama3-gradient:70b-instruct-1048k-q3_K_L
llama3-gradient:70b-instruct-1048k-q4_0
llama3-gradient:70b-instruct-1048k-q4_1
llama3-gradient:70b-instruct-1048k-q4_K_S
llama3-gradient:70b-instruct-1048k-q4_K_M
llama3-gradient:70b-instruct-1048k-q5_0
llama3-gradient:70b-instruct-1048k-q5_1
llama3-gradient:70b-instruct-1048k-q5_K_S
llama3-gradient:70b-instruct-1048k-q5_K_M
llama3-gradient:70b-instruct-1048k-q6_K
llama3-gradient:70b-instruct-1048k-q8_0
llama3-gradient:70b-instruct-1048k-fp16
dolphincoder:latest
dolphincoder
dolphincoder:7b
dolphincoder:15b
dolphincoder:7b-starcoder2
dolphincoder:7b-starcoder2-q2_K
dolphincoder:7b-starcoder2-q3_K_S
dolphincoder:7b-starcoder2-q3_K_M
dolphincoder:7b-starcoder2-q3_K_L
dolphincoder:7b-starcoder2-q4_0
dolphincoder:7b-starcoder2-q4_1
dolphincoder:7b-starcoder2-q4_K_S
dolphincoder:7b-starcoder2-q4_K_M
dolphincoder:7b-starcoder2-q5_0
dolphincoder:7b-starcoder2-q5_1
dolphincoder:7b-starcoder2-q5_K_S
dolphincoder:7b-starcoder2-q5_K_M
dolphincoder:7b-starcoder2-q6_K
dolphincoder:7b-starcoder2-q8_0
dolphincoder:7b-starcoder2-fp16
dolphincoder:15b-starcoder2
dolphincoder:15b-starcoder2-q2_K
dolphincoder:15b-starcoder2-q3_K_S
dolphincoder:15b-starcoder2-q3_K_M
dolphincoder:15b-starcoder2-q3_K_L
dolphincoder:15b-starcoder2-q4_0
dolphincoder:15b-starcoder2-q4_1
dolphincoder:15b-starcoder2-q4_K_S
dolphincoder:15b-starcoder2-q4_K_M
dolphincoder:15b-starcoder2-q5_0
dolphincoder:15b-starcoder2-q5_1
dolphincoder:15b-starcoder2-q5_K_S
dolphincoder:15b-starcoder2-q5_K_M
dolphincoder:15b-starcoder2-q6_K
dolphincoder:15b-starcoder2-q8_0
dolphincoder:15b-starcoder2-fp16
samantha-mistral:latest
samantha-mistral
samantha-mistral:7b
samantha-mistral:7b-instruct-q2_K
samantha-mistral:7b-instruct-q3_K_S
samantha-mistral:7b-instruct-q3_K_M
samantha-mistral:7b-instruct-q3_K_L
samantha-mistral:7b-instruct-q4_0
samantha-mistral:7b-instruct-q4_1
samantha-mistral:7b-instruct-q4_K_S
samantha-mistral:7b-instruct-q4_K_M
samantha-mistral:7b-instruct-q5_0
samantha-mistral:7b-instruct-q5_1
samantha-mistral:7b-instruct-q5_K_S
samantha-mistral:7b-instruct-q5_K_M
samantha-mistral:7b-instruct-q6_K
samantha-mistral:7b-instruct-q8_0
samantha-mistral:7b-instruct-fp16
samantha-mistral:7b-text
samantha-mistral:7b-text-q2_K
samantha-mistral:7b-text-q3_K_S
samantha-mistral:7b-text-q3_K_M
samantha-mistral:7b-text-q3_K_L
samantha-mistral:7b-text-q4_0
samantha-mistral:7b-text-q4_1
samantha-mistral:7b-text-q4_K_S
samantha-mistral:7b-text-q4_K_M
samantha-mistral:7b-text-q5_0
samantha-mistral:7b-text-q5_1
samantha-mistral:7b-text-q5_K_S
samantha-mistral:7b-text-q5_K_M
samantha-mistral:7b-text-q6_K
samantha-mistral:7b-text-q8_0
samantha-mistral:7b-text-fp16
samantha-mistral:7b-v1.2-text
samantha-mistral:7b-v1.2-text-q2_K
samantha-mistral:7b-v1.2-text-q3_K_S
samantha-mistral:7b-v1.2-text-q3_K_M
samantha-mistral:7b-v1.2-text-q3_K_L
samantha-mistral:7b-v1.2-text-q4_0
samantha-mistral:7b-v1.2-text-q4_1
samantha-mistral:7b-v1.2-text-q4_K_S
samantha-mistral:7b-v1.2-text-q4_K_M
samantha-mistral:7b-v1.2-text-q5_0
samantha-mistral:7b-v1.2-text-q5_1
samantha-mistral:7b-v1.2-text-q5_K_S
samantha-mistral:7b-v1.2-text-q5_K_M
samantha-mistral:7b-v1.2-text-q6_K
samantha-mistral:7b-v1.2-text-q8_0
samantha-mistral:7b-v1.2-text-fp16
nemotron-mini:latest
nemotron-mini
nemotron-mini:4b
nemotron-mini:4b-instruct-q2_K
nemotron-mini:4b-instruct-q3_K_S
nemotron-mini:4b-instruct-q3_K_M
nemotron-mini:4b-instruct-q3_K_L
nemotron-mini:4b-instruct-q4_0
nemotron-mini:4b-instruct-q4_1
nemotron-mini:4b-instruct-q4_K_S
nemotron-mini:4b-instruct-q4_K_M
nemotron-mini:4b-instruct-q5_0
nemotron-mini:4b-instruct-q5_1
nemotron-mini:4b-instruct-q5_K_S
nemotron-mini:4b-instruct-q5_K_M
nemotron-mini:4b-instruct-q6_K
nemotron-mini:4b-instruct-q8_0
nemotron-mini:4b-instruct-fp16
dbrx:latest
dbrx
dbrx:instruct
dbrx:132b
dbrx:132b-instruct-q2_K
dbrx:132b-instruct-q4_0
dbrx:132b-instruct-q8_0
dbrx:132b-instruct-fp16
internlm2:latest
internlm2
internlm2:1m
internlm2:1.8b
internlm2:7b
internlm2:20b
internlm2:1.8b-chat-v2.5-q2_K
internlm2:1.8b-chat-v2.5-q3_K_S
internlm2:1.8b-chat-v2.5-q3_K_M
internlm2:1.8b-chat-v2.5-q3_K_L
internlm2:1.8b-chat-v2.5-q4_0
internlm2:1.8b-chat-v2.5-q4_1
internlm2:1.8b-chat-v2.5-q4_K_S
internlm2:1.8b-chat-v2.5-q4_K_M
internlm2:1.8b-chat-v2.5-q5_0
internlm2:1.8b-chat-v2.5-q5_1
internlm2:1.8b-chat-v2.5-q5_K_S
internlm2:1.8b-chat-v2.5-q5_K_M
internlm2:1.8b-chat-v2.5-q6_K
internlm2:1.8b-chat-v2.5-q8_0
internlm2:1.8b-chat-v2.5-fp16
internlm2:7b-chat-1m-v2.5-q2_K
internlm2:7b-chat-1m-v2.5-q3_K_S
internlm2:7b-chat-1m-v2.5-q3_K_M
internlm2:7b-chat-1m-v2.5-q3_K_L
internlm2:7b-chat-1m-v2.5-q4_0
internlm2:7b-chat-1m-v2.5-q4_1
internlm2:7b-chat-1m-v2.5-q4_K_S
internlm2:7b-chat-1m-v2.5-q4_K_M
internlm2:7b-chat-1m-v2.5-q5_0
internlm2:7b-chat-1m-v2.5-q5_1
internlm2:7b-chat-1m-v2.5-q5_K_S
internlm2:7b-chat-1m-v2.5-q5_K_M
internlm2:7b-chat-1m-v2.5-q6_K
internlm2:7b-chat-1m-v2.5-q8_0
internlm2:7b-chat-1m-v2.5-fp16
internlm2:7b-chat-v2.5-q2_K
internlm2:7b-chat-v2.5-q3_K_S
internlm2:7b-chat-v2.5-q3_K_M
internlm2:7b-chat-v2.5-q3_K_L
internlm2:7b-chat-v2.5-q4_0
internlm2:7b-chat-v2.5-q4_1
internlm2:7b-chat-v2.5-q4_K_S
internlm2:7b-chat-v2.5-q4_K_M
internlm2:7b-chat-v2.5-q5_0
internlm2:7b-chat-v2.5-q5_1
internlm2:7b-chat-v2.5-q5_K_S
internlm2:7b-chat-v2.5-q5_K_M
internlm2:7b-chat-v2.5-q6_K
internlm2:7b-chat-v2.5-q8_0
internlm2:7b-chat-v2.5-fp16
internlm2:20b-chat-v2.5-q2_K
internlm2:20b-chat-v2.5-q3_K_S
internlm2:20b-chat-v2.5-q3_K_M
internlm2:20b-chat-v2.5-q3_K_L
internlm2:20b-chat-v2.5-q4_0
internlm2:20b-chat-v2.5-q4_1
internlm2:20b-chat-v2.5-q4_K_S
internlm2:20b-chat-v2.5-q4_K_M
internlm2:20b-chat-v2.5-q5_0
internlm2:20b-chat-v2.5-q5_1
internlm2:20b-chat-v2.5-q5_K_S
internlm2:20b-chat-v2.5-q5_K_M
internlm2:20b-chat-v2.5-q6_K
internlm2:20b-chat-v2.5-q8_0
internlm2:20b-chat-v2.5-fp16
tulu3:latest
tulu3
tulu3:8b
tulu3:70b
tulu3:8b-q4_K_M
tulu3:8b-q8_0
tulu3:8b-fp16
tulu3:70b-q4_K_M
tulu3:70b-q8_0
tulu3:70b-fp16
starling-lm:latest
starling-lm
starling-lm:alpha
starling-lm:beta
starling-lm:7b
starling-lm:7b-alpha
starling-lm:7b-alpha-q2_K
starling-lm:7b-alpha-q3_K_S
starling-lm:7b-alpha-q3_K_M
starling-lm:7b-alpha-q3_K_L
starling-lm:7b-alpha-q4_0
starling-lm:7b-alpha-q4_1
starling-lm:7b-alpha-q4_K_S
starling-lm:7b-alpha-q4_K_M
starling-lm:7b-alpha-q5_0
starling-lm:7b-alpha-q5_1
starling-lm:7b-alpha-q5_K_S
starling-lm:7b-alpha-q5_K_M
starling-lm:7b-alpha-q6_K
starling-lm:7b-alpha-q8_0
starling-lm:7b-alpha-fp16
starling-lm:7b-beta
starling-lm:7b-beta-q2_K
starling-lm:7b-beta-q3_K_S
starling-lm:7b-beta-q3_K_M
starling-lm:7b-beta-q3_K_L
starling-lm:7b-beta-q4_0
starling-lm:7b-beta-q4_1
starling-lm:7b-beta-q4_K_S
starling-lm:7b-beta-q4_K_M
starling-lm:7b-beta-q5_0
starling-lm:7b-beta-q5_1
starling-lm:7b-beta-q5_K_S
starling-lm:7b-beta-q5_K_M
starling-lm:7b-beta-q6_K
starling-lm:7b-beta-q8_0
starling-lm:7b-beta-fp16
llama3-groq-tool-use:latest
llama3-groq-tool-use
llama3-groq-tool-use:8b
llama3-groq-tool-use:70b
llama3-groq-tool-use:8b-q2_K
llama3-groq-tool-use:8b-q3_K_S
llama3-groq-tool-use:8b-q3_K_M
llama3-groq-tool-use:8b-q3_K_L
llama3-groq-tool-use:8b-q4_0
llama3-groq-tool-use:8b-q4_1
llama3-groq-tool-use:8b-q4_K_S
llama3-groq-tool-use:8b-q4_K_M
llama3-groq-tool-use:8b-q5_0
llama3-groq-tool-use:8b-q5_1
llama3-groq-tool-use:8b-q5_K_S
llama3-groq-tool-use:8b-q5_K_M
llama3-groq-tool-use:8b-q6_K
llama3-groq-tool-use:8b-q8_0
llama3-groq-tool-use:8b-fp16
llama3-groq-tool-use:70b-q2_K
llama3-groq-tool-use:70b-q3_K_S
llama3-groq-tool-use:70b-q3_K_M
llama3-groq-tool-use:70b-q3_K_L
llama3-groq-tool-use:70b-q4_0
llama3-groq-tool-use:70b-q4_1
llama3-groq-tool-use:70b-q4_K_S
llama3-groq-tool-use:70b-q4_K_M
llama3-groq-tool-use:70b-q5_0
llama3-groq-tool-use:70b-q5_1
llama3-groq-tool-use:70b-q5_K_S
llama3-groq-tool-use:70b-q5_K_M
llama3-groq-tool-use:70b-q6_K
llama3-groq-tool-use:70b-q8_0
llama3-groq-tool-use:70b-fp16
athene-v2:latest
athene-v2
athene-v2:72b
athene-v2:72b-q2_K
athene-v2:72b-q3_K_S
athene-v2:72b-q3_K_M
athene-v2:72b-q3_K_L
athene-v2:72b-q4_0
athene-v2:72b-q4_1
athene-v2:72b-q4_K_S
athene-v2:72b-q4_K_M
athene-v2:72b-q5_0
athene-v2:72b-q5_1
athene-v2:72b-q5_K_S
athene-v2:72b-q5_K_M
athene-v2:72b-q6_K
athene-v2:72b-q8_0
athene-v2:72b-fp16
phind-codellama:latest
phind-codellama
phind-codellama:34b
phind-codellama:34b-python
phind-codellama:34b-python-q2_K
phind-codellama:34b-python-q3_K_S
phind-codellama:34b-python-q3_K_M
phind-codellama:34b-python-q3_K_L
phind-codellama:34b-python-q4_0
phind-codellama:34b-python-q4_1
phind-codellama:34b-python-q4_K_S
phind-codellama:34b-python-q4_K_M
phind-codellama:34b-python-q5_0
phind-codellama:34b-python-q5_1
phind-codellama:34b-python-q5_K_S
phind-codellama:34b-python-q5_K_M
phind-codellama:34b-python-q6_K
phind-codellama:34b-python-q8_0
phind-codellama:34b-python-fp16
phind-codellama:34b-v2
phind-codellama:34b-v2-q2_K
phind-codellama:34b-v2-q3_K_S
phind-codellama:34b-v2-q3_K_M
phind-codellama:34b-v2-q3_K_L
phind-codellama:34b-v2-q4_0
phind-codellama:34b-v2-q4_1
phind-codellama:34b-v2-q4_K_S
phind-codellama:34b-v2-q4_K_M
phind-codellama:34b-v2-q5_0
phind-codellama:34b-v2-q5_1
phind-codellama:34b-v2-q5_K_S
phind-codellama:34b-v2-q5_K_M
phind-codellama:34b-v2-q6_K
phind-codellama:34b-v2-q8_0
phind-codellama:34b-v2-fp16
phind-codellama:34b-q2_K
phind-codellama:34b-q3_K_S
phind-codellama:34b-q3_K_M
phind-codellama:34b-q3_K_L
phind-codellama:34b-q4_0
phind-codellama:34b-q4_1
phind-codellama:34b-q4_K_S
phind-codellama:34b-q4_K_M
phind-codellama:34b-q5_0
phind-codellama:34b-q5_1
phind-codellama:34b-q5_K_S
phind-codellama:34b-q5_K_M
phind-codellama:34b-q6_K
phind-codellama:34b-q8_0
phind-codellama:34b-fp16
solar:latest
solar
solar:10.7b
solar:10.7b-instruct-v1-q2_K
solar:10.7b-instruct-v1-q3_K_S
solar:10.7b-instruct-v1-q3_K_M
solar:10.7b-instruct-v1-q3_K_L
solar:10.7b-instruct-v1-q4_0
solar:10.7b-instruct-v1-q4_1
solar:10.7b-instruct-v1-q4_K_S
solar:10.7b-instruct-v1-q4_K_M
solar:10.7b-instruct-v1-q5_0
solar:10.7b-instruct-v1-q5_1
solar:10.7b-instruct-v1-q5_K_S
solar:10.7b-instruct-v1-q5_K_M
solar:10.7b-instruct-v1-q6_K
solar:10.7b-instruct-v1-q8_0
solar:10.7b-instruct-v1-fp16
solar:10.7b-text-v1-q2_K
solar:10.7b-text-v1-q3_K_S
solar:10.7b-text-v1-q3_K_M
solar:10.7b-text-v1-q3_K_L
solar:10.7b-text-v1-q4_0
solar:10.7b-text-v1-q4_1
solar:10.7b-text-v1-q4_K_S
solar:10.7b-text-v1-q4_K_M
solar:10.7b-text-v1-q5_0
solar:10.7b-text-v1-q5_1
solar:10.7b-text-v1-q5_K_S
solar:10.7b-text-v1-q5_K_M
solar:10.7b-text-v1-q6_K
solar:10.7b-text-v1-q8_0
solar:10.7b-text-v1-fp16
xwinlm:latest
xwinlm
xwinlm:7b
xwinlm:13b
xwinlm:7b-v0.1
xwinlm:7b-v0.1-q2_K
xwinlm:7b-v0.1-q3_K_S
xwinlm:7b-v0.1-q3_K_M
xwinlm:7b-v0.1-q3_K_L
xwinlm:7b-v0.1-q4_0
xwinlm:7b-v0.1-q4_1
xwinlm:7b-v0.1-q4_K_S
xwinlm:7b-v0.1-q4_K_M
xwinlm:7b-v0.1-q5_0
xwinlm:7b-v0.1-q5_1
xwinlm:7b-v0.1-q5_K_S
xwinlm:7b-v0.1-q5_K_M
xwinlm:7b-v0.1-q6_K
xwinlm:7b-v0.1-q8_0
xwinlm:7b-v0.1-fp16
xwinlm:7b-v0.2
xwinlm:7b-v0.2-q2_K
xwinlm:7b-v0.2-q3_K_S
xwinlm:7b-v0.2-q3_K_L
xwinlm:7b-v0.2-q4_0
xwinlm:7b-v0.2-q4_1
xwinlm:7b-v0.2-q4_K_S
xwinlm:7b-v0.2-q4_K_M
xwinlm:7b-v0.2-q5_0
xwinlm:7b-v0.2-q5_K_S
xwinlm:7b-v0.2-q5_K_M
xwinlm:7b-v0.2-q6_K
xwinlm:7b-v0.2-q8_0
xwinlm:7b-v0.2-fp16
xwinlm:13b-v0.1
xwinlm:13b-v0.1-q2_K
xwinlm:13b-v0.1-q3_K_S
xwinlm:13b-v0.1-q3_K_M
xwinlm:13b-v0.1-q3_K_L
xwinlm:13b-v0.1-q4_0
xwinlm:13b-v0.1-q4_1
xwinlm:13b-v0.1-q4_K_S
xwinlm:13b-v0.1-q4_K_M
xwinlm:13b-v0.1-q5_0
xwinlm:13b-v0.1-q5_1
xwinlm:13b-v0.1-q5_K_S
xwinlm:13b-v0.1-q5_K_M
xwinlm:13b-v0.1-q6_K
xwinlm:13b-v0.1-q8_0
xwinlm:13b-v0.1-fp16
xwinlm:13b-v0.2
xwinlm:13b-v0.2-q2_K
xwinlm:13b-v0.2-q3_K_S
xwinlm:13b-v0.2-q3_K_M
xwinlm:13b-v0.2-q3_K_L
xwinlm:13b-v0.2-q4_0
xwinlm:13b-v0.2-q4_1
xwinlm:13b-v0.2-q4_K_S
xwinlm:13b-v0.2-q4_K_M
xwinlm:13b-v0.2-q5_0
xwinlm:13b-v0.2-q5_1
xwinlm:13b-v0.2-q5_K_S
xwinlm:13b-v0.2-q5_K_M
xwinlm:13b-v0.2-q6_K
xwinlm:13b-v0.2-q8_0
xwinlm:13b-v0.2-fp16
xwinlm:70b-v0.1
xwinlm:70b-v0.1-q2_K
xwinlm:70b-v0.1-q3_K_S
xwinlm:70b-v0.1-q3_K_M
xwinlm:70b-v0.1-q3_K_L
xwinlm:70b-v0.1-q4_0
xwinlm:70b-v0.1-q4_1
xwinlm:70b-v0.1-q4_K_S
xwinlm:70b-v0.1-q4_K_M
xwinlm:70b-v0.1-q5_0
xwinlm:70b-v0.1-q5_1
xwinlm:70b-v0.1-q5_K_S
xwinlm:70b-v0.1-q6_K
xwinlm:70b-v0.1-q8_0
xwinlm:70b-v0.1-fp16
nemotron:latest
nemotron
nemotron:70b
nemotron:70b-instruct-q2_K
nemotron:70b-instruct-q3_K_S
nemotron:70b-instruct-q3_K_M
nemotron:70b-instruct-q3_K_L
nemotron:70b-instruct-q4_0
nemotron:70b-instruct-q4_1
nemotron:70b-instruct-q4_K_S
nemotron:70b-instruct-q4_K_M
nemotron:70b-instruct-q5_0
nemotron:70b-instruct-q5_1
nemotron:70b-instruct-q5_K_S
nemotron:70b-instruct-q5_K_M
nemotron:70b-instruct-q6_K
nemotron:70b-instruct-q8_0
nemotron:70b-instruct-fp16
llama-guard3:latest
llama-guard3
llama-guard3:1b
llama-guard3:8b
llama-guard3:1b-q2_K
llama-guard3:1b-q3_K_S
llama-guard3:1b-q3_K_M
llama-guard3:1b-q3_K_L
llama-guard3:1b-q4_0
llama-guard3:1b-q4_1
llama-guard3:1b-q4_K_S
llama-guard3:1b-q4_K_M
llama-guard3:1b-q5_0
llama-guard3:1b-q5_1
llama-guard3:1b-q5_K_S
llama-guard3:1b-q5_K_M
llama-guard3:1b-q6_K
llama-guard3:1b-q8_0
llama-guard3:1b-fp16
llama-guard3:8b-q2_K
llama-guard3:8b-q3_K_S
llama-guard3:8b-q3_K_M
llama-guard3:8b-q3_K_L
llama-guard3:8b-q4_0
llama-guard3:8b-q4_1
llama-guard3:8b-q4_K_S
llama-guard3:8b-q4_K_M
llama-guard3:8b-q5_0
llama-guard3:8b-q5_1
llama-guard3:8b-q5_K_S
llama-guard3:8b-q5_K_M
llama-guard3:8b-q6_K
llama-guard3:8b-q8_0
llama-guard3:8b-fp16
meditron:latest
meditron
meditron:7b
meditron:70b
meditron:7b-q2_K
meditron:7b-q3_K_S
meditron:7b-q3_K_M
meditron:7b-q3_K_L
meditron:7b-q4_0
meditron:7b-q4_1
meditron:7b-q4_K_S
meditron:7b-q4_K_M
meditron:7b-q5_0
meditron:7b-q5_1
meditron:7b-q5_K_S
meditron:7b-q5_K_M
meditron:7b-q6_K
meditron:7b-q8_0
meditron:7b-fp16
meditron:70b-q4_0
meditron:70b-q4_1
meditron:70b-q4_K_S
meditron:70b-q5_1
yarn-llama2:latest
yarn-llama2
yarn-llama2:7b
yarn-llama2:13b
yarn-llama2:7b-128k
yarn-llama2:7b-128k-q2_K
yarn-llama2:7b-128k-q3_K_S
yarn-llama2:7b-128k-q3_K_M
yarn-llama2:7b-128k-q3_K_L
yarn-llama2:7b-128k-q4_0
yarn-llama2:7b-128k-q4_1
yarn-llama2:7b-128k-q4_K_S
yarn-llama2:7b-128k-q4_K_M
yarn-llama2:7b-128k-q5_0
yarn-llama2:7b-128k-q5_1
yarn-llama2:7b-128k-q5_K_S
yarn-llama2:7b-128k-q5_K_M
yarn-llama2:7b-128k-q6_K
yarn-llama2:7b-128k-q8_0
yarn-llama2:7b-128k-fp16
yarn-llama2:7b-64k
yarn-llama2:7b-64k-q2_K
yarn-llama2:7b-64k-q3_K_S
yarn-llama2:7b-64k-q3_K_M
yarn-llama2:7b-64k-q3_K_L
yarn-llama2:7b-64k-q4_0
yarn-llama2:7b-64k-q4_1
yarn-llama2:7b-64k-q4_K_S
yarn-llama2:7b-64k-q4_K_M
yarn-llama2:7b-64k-q5_0
yarn-llama2:7b-64k-q5_1
yarn-llama2:7b-64k-q5_K_S
yarn-llama2:7b-64k-q5_K_M
yarn-llama2:7b-64k-q6_K
yarn-llama2:7b-64k-q8_0
yarn-llama2:7b-64k-fp16
yarn-llama2:13b-128k
yarn-llama2:13b-128k-q2_K
yarn-llama2:13b-128k-q3_K_S
yarn-llama2:13b-128k-q3_K_M
yarn-llama2:13b-128k-q3_K_L
yarn-llama2:13b-128k-q4_0
yarn-llama2:13b-128k-q4_1
yarn-llama2:13b-128k-q4_K_S
yarn-llama2:13b-128k-q4_K_M
yarn-llama2:13b-128k-q5_0
yarn-llama2:13b-128k-q5_1
yarn-llama2:13b-128k-q5_K_S
yarn-llama2:13b-128k-q5_K_M
yarn-llama2:13b-128k-q6_K
yarn-llama2:13b-128k-q8_0
yarn-llama2:13b-128k-fp16
yarn-llama2:13b-64k
yarn-llama2:13b-64k-q2_K
yarn-llama2:13b-64k-q3_K_S
yarn-llama2:13b-64k-q3_K_M
yarn-llama2:13b-64k-q3_K_L
yarn-llama2:13b-64k-q4_0
yarn-llama2:13b-64k-q4_1
yarn-llama2:13b-64k-q4_K_S
yarn-llama2:13b-64k-q4_K_M
yarn-llama2:13b-64k-q5_0
yarn-llama2:13b-64k-q5_1
yarn-llama2:13b-64k-q5_K_S
yarn-llama2:13b-64k-q5_K_M
yarn-llama2:13b-64k-q6_K
yarn-llama2:13b-64k-q8_0
yarn-llama2:13b-64k-fp16
aya-expanse:latest
aya-expanse
aya-expanse:8b
aya-expanse:32b
aya-expanse:8b-q2_K
aya-expanse:8b-q3_K_S
aya-expanse:8b-q3_K_M
aya-expanse:8b-q3_K_L
aya-expanse:8b-q4_0
aya-expanse:8b-q4_1
aya-expanse:8b-q4_K_S
aya-expanse:8b-q4_K_M
aya-expanse:8b-q5_0
aya-expanse:8b-q5_1
aya-expanse:8b-q5_K_S
aya-expanse:8b-q5_K_M
aya-expanse:8b-q6_K
aya-expanse:8b-q8_0
aya-expanse:8b-fp16
aya-expanse:32b-q2_K
aya-expanse:32b-q3_K_S
aya-expanse:32b-q3_K_M
aya-expanse:32b-q3_K_L
aya-expanse:32b-q4_0
aya-expanse:32b-q4_1
aya-expanse:32b-q4_K_S
aya-expanse:32b-q4_K_M
aya-expanse:32b-q5_0
aya-expanse:32b-q5_1
aya-expanse:32b-q5_K_S
aya-expanse:32b-q5_K_M
aya-expanse:32b-q6_K
aya-expanse:32b-q8_0
aya-expanse:32b-fp16
wizardlm-uncensored:latest
wizardlm-uncensored
wizardlm-uncensored:13b
wizardlm-uncensored:13b-llama2
wizardlm-uncensored:13b-llama2-q2_K
wizardlm-uncensored:13b-llama2-q3_K_S
wizardlm-uncensored:13b-llama2-q3_K_M
wizardlm-uncensored:13b-llama2-q3_K_L
wizardlm-uncensored:13b-llama2-q4_0
wizardlm-uncensored:13b-llama2-q4_1
wizardlm-uncensored:13b-llama2-q4_K_S
wizardlm-uncensored:13b-llama2-q4_K_M
wizardlm-uncensored:13b-llama2-q5_0
wizardlm-uncensored:13b-llama2-q5_1
wizardlm-uncensored:13b-llama2-q5_K_S
wizardlm-uncensored:13b-llama2-q5_K_M
wizardlm-uncensored:13b-llama2-q6_K
wizardlm-uncensored:13b-llama2-q8_0
wizardlm-uncensored:13b-llama2-fp16
granite3-moe:latest
granite3-moe
granite3-moe:1b
granite3-moe:3b
granite3-moe:1b-instruct-q2_K
granite3-moe:1b-instruct-q3_K_S
granite3-moe:1b-instruct-q3_K_M
granite3-moe:1b-instruct-q3_K_L
granite3-moe:1b-instruct-q4_0
granite3-moe:1b-instruct-q4_1
granite3-moe:1b-instruct-q4_K_S
granite3-moe:1b-instruct-q4_K_M
granite3-moe:1b-instruct-q5_0
granite3-moe:1b-instruct-q5_1
granite3-moe:1b-instruct-q5_K_S
granite3-moe:1b-instruct-q5_K_M
granite3-moe:1b-instruct-q6_K
granite3-moe:1b-instruct-q8_0
granite3-moe:1b-instruct-fp16
granite3-moe:3b-instruct-q2_K
granite3-moe:3b-instruct-q3_K_S
granite3-moe:3b-instruct-q3_K_M
granite3-moe:3b-instruct-q3_K_L
granite3-moe:3b-instruct-q4_0
granite3-moe:3b-instruct-q4_1
granite3-moe:3b-instruct-q4_K_S
granite3-moe:3b-instruct-q4_K_M
granite3-moe:3b-instruct-q5_0
granite3-moe:3b-instruct-q5_1
granite3-moe:3b-instruct-q5_K_S
granite3-moe:3b-instruct-q5_K_M
granite3-moe:3b-instruct-q6_K
granite3-moe:3b-instruct-q8_0
granite3-moe:3b-instruct-fp16
smallthinker:latest
smallthinker
smallthinker:3b
smallthinker:3b-preview-q4_K_M
smallthinker:3b-preview-q8_0
smallthinker:3b-preview-fp16
orca2:latest
orca2
orca2:7b
orca2:13b
orca2:7b-q2_K
orca2:7b-q3_K_S
orca2:7b-q3_K_M
orca2:7b-q3_K_L
orca2:7b-q4_0
orca2:7b-q4_1
orca2:7b-q4_K_S
orca2:7b-q4_K_M
orca2:7b-q5_0
orca2:7b-q5_1
orca2:7b-q5_K_S
orca2:7b-q5_K_M
orca2:7b-q6_K
orca2:7b-q8_0
orca2:7b-fp16
orca2:13b-q2_K
orca2:13b-q3_K_S
orca2:13b-q3_K_M
orca2:13b-q3_K_L
orca2:13b-q4_0
orca2:13b-q4_1
orca2:13b-q4_K_S
orca2:13b-q4_K_M
orca2:13b-q5_0
orca2:13b-q5_1
orca2:13b-q5_K_S
orca2:13b-q5_K_M
orca2:13b-q6_K
orca2:13b-q8_0
orca2:13b-fp16
medllama2:latest
medllama2
medllama2:7b
medllama2:7b-q2_K
medllama2:7b-q3_K_S
medllama2:7b-q3_K_M
medllama2:7b-q3_K_L
medllama2:7b-q4_0
medllama2:7b-q4_1
medllama2:7b-q4_K_S
medllama2:7b-q4_K_M
medllama2:7b-q5_0
medllama2:7b-q5_1
medllama2:7b-q5_K_S
medllama2:7b-q5_K_M
medllama2:7b-q6_K
medllama2:7b-q8_0
medllama2:7b-fp16
command-r7b:latest
command-r7b
command-r7b:7b
command-r7b:7b-12-2024-q4_K_M
command-r7b:7b-12-2024-q8_0
command-r7b:7b-12-2024-fp16
phi4-mini-reasoning:latest
phi4-mini-reasoning
phi4-mini-reasoning:3.8b
phi4-mini-reasoning:3.8b-q4_K_M
phi4-mini-reasoning:3.8b-q8_0
phi4-mini-reasoning:3.8b-fp16
nous-hermes2-mixtral:latest
nous-hermes2-mixtral
nous-hermes2-mixtral:dpo
nous-hermes2-mixtral:8x7b
nous-hermes2-mixtral:8x7b-dpo-q2_K
nous-hermes2-mixtral:8x7b-dpo-q3_K_S
nous-hermes2-mixtral:8x7b-dpo-q3_K_M
nous-hermes2-mixtral:8x7b-dpo-q3_K_L
nous-hermes2-mixtral:8x7b-dpo-q4_0
nous-hermes2-mixtral:8x7b-dpo-q4_1
nous-hermes2-mixtral:8x7b-dpo-q4_K_S
nous-hermes2-mixtral:8x7b-dpo-q4_K_M
nous-hermes2-mixtral:8x7b-dpo-q5_0
nous-hermes2-mixtral:8x7b-dpo-q5_1
nous-hermes2-mixtral:8x7b-dpo-q5_K_S
nous-hermes2-mixtral:8x7b-dpo-q5_K_M
nous-hermes2-mixtral:8x7b-dpo-q6_K
nous-hermes2-mixtral:8x7b-dpo-q8_0
nous-hermes2-mixtral:8x7b-dpo-fp16
stable-beluga:latest
stable-beluga
stable-beluga:7b
stable-beluga:13b
stable-beluga:70b
stable-beluga:7b-q2_K
stable-beluga:7b-q3_K_S
stable-beluga:7b-q3_K_M
stable-beluga:7b-q3_K_L
stable-beluga:7b-q4_0
stable-beluga:7b-q4_1
stable-beluga:7b-q4_K_S
stable-beluga:7b-q4_K_M
stable-beluga:7b-q5_0
stable-beluga:7b-q5_1
stable-beluga:7b-q5_K_S
stable-beluga:7b-q5_K_M
stable-beluga:7b-q6_K
stable-beluga:7b-q8_0
stable-beluga:7b-fp16
stable-beluga:13b-q2_K
stable-beluga:13b-q3_K_S
stable-beluga:13b-q3_K_M
stable-beluga:13b-q3_K_L
stable-beluga:13b-q4_0
stable-beluga:13b-q4_1
stable-beluga:13b-q4_K_S
stable-beluga:13b-q4_K_M
stable-beluga:13b-q5_0
stable-beluga:13b-q5_1
stable-beluga:13b-q5_K_S
stable-beluga:13b-q5_K_M
stable-beluga:13b-q6_K
stable-beluga:13b-q8_0
stable-beluga:13b-fp16
stable-beluga:70b-q2_K
stable-beluga:70b-q3_K_S
stable-beluga:70b-q3_K_M
stable-beluga:70b-q3_K_L
stable-beluga:70b-q4_0
stable-beluga:70b-q4_1
stable-beluga:70b-q4_K_S
stable-beluga:70b-q4_K_M
stable-beluga:70b-q5_0
stable-beluga:70b-q5_1
stable-beluga:70b-q5_K_S
stable-beluga:70b-q5_K_M
stable-beluga:70b-q6_K
stable-beluga:70b-q8_0
stable-beluga:70b-fp16
deepseek-v2.5:latest
deepseek-v2.5
deepseek-v2.5:236b
deepseek-v2.5:236b-q4_0
deepseek-v2.5:236b-q4_1
deepseek-v2.5:236b-q5_0
deepseek-v2.5:236b-q5_1
deepseek-v2.5:236b-q8_0
reader-lm:latest
reader-lm
reader-lm:0.5b
reader-lm:1.5b
reader-lm:0.5b-q2_K
reader-lm:0.5b-q3_K_S
reader-lm:0.5b-q3_K_M
reader-lm:0.5b-q3_K_L
reader-lm:0.5b-q4_0
reader-lm:0.5b-q4_1
reader-lm:0.5b-q4_K_S
reader-lm:0.5b-q4_K_M
reader-lm:0.5b-q5_0
reader-lm:0.5b-q5_1
reader-lm:0.5b-q5_K_S
reader-lm:0.5b-q5_K_M
reader-lm:0.5b-q6_K
reader-lm:0.5b-q8_0
reader-lm:0.5b-fp16
reader-lm:1.5b-q2_K
reader-lm:1.5b-q3_K_S
reader-lm:1.5b-q3_K_M
reader-lm:1.5b-q3_K_L
reader-lm:1.5b-q4_0
reader-lm:1.5b-q4_1
reader-lm:1.5b-q4_K_S
reader-lm:1.5b-q4_K_M
reader-lm:1.5b-q5_0
reader-lm:1.5b-q5_1
reader-lm:1.5b-q5_K_S
reader-lm:1.5b-q5_K_M
reader-lm:1.5b-q6_K
reader-lm:1.5b-q8_0
reader-lm:1.5b-fp16
shieldgemma:latest
shieldgemma
shieldgemma:2b
shieldgemma:9b
shieldgemma:27b
shieldgemma:2b-q2_K
shieldgemma:2b-q3_K_S
shieldgemma:2b-q3_K_M
shieldgemma:2b-q3_K_L
shieldgemma:2b-q4_0
shieldgemma:2b-q4_1
shieldgemma:2b-q4_K_S
shieldgemma:2b-q4_K_M
shieldgemma:2b-q5_0
shieldgemma:2b-q5_1
shieldgemma:2b-q5_K_S
shieldgemma:2b-q5_K_M
shieldgemma:2b-q6_K
shieldgemma:2b-q8_0
shieldgemma:2b-fp16
shieldgemma:9b-q2_K
shieldgemma:9b-q3_K_S
shieldgemma:9b-q3_K_M
shieldgemma:9b-q3_K_L
shieldgemma:9b-q4_0
shieldgemma:9b-q4_1
shieldgemma:9b-q4_K_S
shieldgemma:9b-q4_K_M
shieldgemma:9b-q5_0
shieldgemma:9b-q5_1
shieldgemma:9b-q5_K_S
shieldgemma:9b-q5_K_M
shieldgemma:9b-q6_K
shieldgemma:9b-q8_0
shieldgemma:9b-fp16
shieldgemma:27b-q2_K
shieldgemma:27b-q3_K_S
shieldgemma:27b-q3_K_M
shieldgemma:27b-q3_K_L
shieldgemma:27b-q4_0
shieldgemma:27b-q4_1
shieldgemma:27b-q4_K_S
shieldgemma:27b-q4_K_M
shieldgemma:27b-q5_0
shieldgemma:27b-q5_1
shieldgemma:27b-q5_K_S
shieldgemma:27b-q5_K_M
shieldgemma:27b-q6_K
shieldgemma:27b-q8_0
shieldgemma:27b-fp16
command-a:latest
command-a
command-a:111b
command-a:111b-03-2025-q4_K_M
command-a:111b-03-2025-q8_0
command-a:111b-03-2025-fp16
llama-pro:latest
llama-pro
llama-pro:instruct
llama-pro:text
llama-pro:8b-instruct-q2_K
llama-pro:8b-instruct-q3_K_S
llama-pro:8b-instruct-q3_K_M
llama-pro:8b-instruct-q3_K_L
llama-pro:8b-instruct-q4_0
llama-pro:8b-instruct-q4_1
llama-pro:8b-instruct-q4_K_S
llama-pro:8b-instruct-q4_K_M
llama-pro:8b-instruct-q5_0
llama-pro:8b-instruct-q5_1
llama-pro:8b-instruct-q5_K_S
llama-pro:8b-instruct-q5_K_M
llama-pro:8b-instruct-q6_K
llama-pro:8b-instruct-q8_0
llama-pro:8b-instruct-fp16
llama-pro:8b-text-q2_K
llama-pro:8b-text-q3_K_S
llama-pro:8b-text-q3_K_M
llama-pro:8b-text-q3_K_L
llama-pro:8b-text-q4_0
llama-pro:8b-text-q4_1
llama-pro:8b-text-q4_K_S
llama-pro:8b-text-q4_K_M
llama-pro:8b-text-q5_0
llama-pro:8b-text-q5_1
llama-pro:8b-text-q5_K_S
llama-pro:8b-text-q5_K_M
llama-pro:8b-text-q6_K
llama-pro:8b-text-q8_0
llama-pro:8b-text-fp16
mathstral:latest
mathstral
mathstral:7b
mathstral:7b-v0.1-q2_K
mathstral:7b-v0.1-q3_K_S
mathstral:7b-v0.1-q3_K_M
mathstral:7b-v0.1-q3_K_L
mathstral:7b-v0.1-q4_0
mathstral:7b-v0.1-q4_1
mathstral:7b-v0.1-q4_K_S
mathstral:7b-v0.1-q4_K_M
mathstral:7b-v0.1-q5_0
mathstral:7b-v0.1-q5_1
mathstral:7b-v0.1-q5_K_S
mathstral:7b-v0.1-q5_K_M
mathstral:7b-v0.1-q6_K
mathstral:7b-v0.1-q8_0
mathstral:7b-v0.1-fp16
wizardlm:7b-q2_K
wizardlm:7b-q3_K_S
wizardlm:7b-q3_K_M
wizardlm:7b-q3_K_L
wizardlm:7b-q4_0
wizardlm:7b-q4_1
wizardlm:7b-q4_K_S
wizardlm:7b-q4_K_M
wizardlm:7b-q5_0
wizardlm:7b-q5_1
wizardlm:7b-q5_K_S
wizardlm:7b-q5_K_M
wizardlm:7b-q6_K
wizardlm:7b-q8_0
wizardlm:7b-fp16
wizardlm:13b-llama2-q2_K
wizardlm:13b-llama2-q3_K_S
wizardlm:13b-llama2-q3_K_M
wizardlm:13b-llama2-q3_K_L
wizardlm:13b-llama2-q4_0
wizardlm:13b-llama2-q4_1
wizardlm:13b-llama2-q4_K_S
wizardlm:13b-llama2-q4_K_M
wizardlm:13b-llama2-q5_0
wizardlm:13b-llama2-q5_1
wizardlm:13b-llama2-q5_K_S
wizardlm:13b-llama2-q5_K_M
wizardlm:13b-llama2-q6_K
wizardlm:13b-llama2-q8_0
wizardlm:13b-llama2-fp16
wizardlm:13b-q2_K
wizardlm:13b-q3_K_S
wizardlm:13b-q3_K_M
wizardlm:13b-q3_K_L
wizardlm:13b-q4_0
wizardlm:13b-q4_1
wizardlm:13b-q4_K_S
wizardlm:13b-q4_K_M
wizardlm:13b-q5_0
wizardlm:13b-q5_1
wizardlm:13b-q5_K_S
wizardlm:13b-q5_K_M
wizardlm:13b-q6_K
wizardlm:13b-q8_0
wizardlm:13b-fp16
wizardlm:30b-q2_K
wizardlm:30b-q3_K_S
wizardlm:30b-q3_K_M
wizardlm:30b-q3_K_L
wizardlm:30b-q4_0
wizardlm:30b-q4_1
wizardlm:30b-q4_K_S
wizardlm:30b-q4_K_M
wizardlm:30b-q5_0
wizardlm:30b-q5_1
wizardlm:30b-q5_K_S
wizardlm:30b-q5_K_M
wizardlm:30b-q6_K
wizardlm:30b-q8_0
wizardlm:30b-fp16
wizardlm:70b-llama2-q2_K
wizardlm:70b-llama2-q3_K_S
wizardlm:70b-llama2-q3_K_M
wizardlm:70b-llama2-q3_K_L
wizardlm:70b-llama2-q4_0
wizardlm:70b-llama2-q4_1
wizardlm:70b-llama2-q4_K_S
wizardlm:70b-llama2-q4_K_M
wizardlm:70b-llama2-q5_0
wizardlm:70b-llama2-q5_K_S
wizardlm:70b-llama2-q5_K_M
wizardlm:70b-llama2-q6_K
wizardlm:70b-llama2-q8_0
yarn-mistral:latest
yarn-mistral
yarn-mistral:7b
yarn-mistral:7b-128k
yarn-mistral:7b-128k-q2_K
yarn-mistral:7b-128k-q3_K_S
yarn-mistral:7b-128k-q3_K_M
yarn-mistral:7b-128k-q3_K_L
yarn-mistral:7b-128k-q4_0
yarn-mistral:7b-128k-q4_1
yarn-mistral:7b-128k-q4_K_S
yarn-mistral:7b-128k-q4_K_M
yarn-mistral:7b-128k-q5_0
yarn-mistral:7b-128k-q5_1
yarn-mistral:7b-128k-q5_K_S
yarn-mistral:7b-128k-q5_K_M
yarn-mistral:7b-128k-q6_K
yarn-mistral:7b-128k-q8_0
yarn-mistral:7b-128k-fp16
yarn-mistral:7b-64k
yarn-mistral:7b-64k-q2_K
yarn-mistral:7b-64k-q3_K_S
yarn-mistral:7b-64k-q3_K_M
yarn-mistral:7b-64k-q3_K_L
yarn-mistral:7b-64k-q4_0
yarn-mistral:7b-64k-q4_1
yarn-mistral:7b-64k-q4_K_S
yarn-mistral:7b-64k-q4_K_M
yarn-mistral:7b-64k-q5_0
yarn-mistral:7b-64k-q5_1
yarn-mistral:7b-64k-q5_K_S
yarn-mistral:7b-64k-q5_K_M
yarn-mistral:7b-64k-q6_K
yarn-mistral:7b-64k-q8_0
everythinglm:latest
everythinglm
everythinglm:13b
everythinglm:13b-16k
everythinglm:13b-16k-q2_K
everythinglm:13b-16k-q3_K_S
everythinglm:13b-16k-q3_K_M
everythinglm:13b-16k-q3_K_L
everythinglm:13b-16k-q4_0
everythinglm:13b-16k-q4_1
everythinglm:13b-16k-q4_K_S
everythinglm:13b-16k-q4_K_M
everythinglm:13b-16k-q5_0
everythinglm:13b-16k-q5_1
everythinglm:13b-16k-q5_K_S
everythinglm:13b-16k-q5_K_M
everythinglm:13b-16k-q6_K
everythinglm:13b-16k-q8_0
everythinglm:13b-16k-fp16
nexusraven:latest
nexusraven
nexusraven:13b
nexusraven:13b-v2-q2_K
nexusraven:13b-v2-q3_K_S
nexusraven:13b-v2-q3_K_M
nexusraven:13b-v2-q3_K_L
nexusraven:13b-v2-q4_0
nexusraven:13b-v2-q4_1
nexusraven:13b-v2-q4_K_S
nexusraven:13b-v2-q4_K_M
nexusraven:13b-v2-q5_0
nexusraven:13b-v2-q5_1
nexusraven:13b-v2-q5_K_S
nexusraven:13b-v2-q5_K_M
nexusraven:13b-v2-q6_K
nexusraven:13b-v2-q8_0
nexusraven:13b-v2-fp16
nexusraven:13b-q2_K
nexusraven:13b-q3_K_S
nexusraven:13b-q3_K_M
nexusraven:13b-q3_K_L
nexusraven:13b-q4_0
nexusraven:13b-q4_1
nexusraven:13b-q4_K_S
nexusraven:13b-q4_K_M
nexusraven:13b-q5_0
nexusraven:13b-q5_1
nexusraven:13b-q5_K_S
nexusraven:13b-q5_K_M
nexusraven:13b-q6_K
nexusraven:13b-q8_0
nexusraven:13b-fp16
codeup:latest
codeup
codeup:13b
codeup:13b-llama2
codeup:13b-llama2-chat
codeup:13b-llama2-chat-q2_K
codeup:13b-llama2-chat-q3_K_S
codeup:13b-llama2-chat-q3_K_M
codeup:13b-llama2-chat-q3_K_L
codeup:13b-llama2-chat-q4_0
codeup:13b-llama2-chat-q4_1
codeup:13b-llama2-chat-q4_K_S
codeup:13b-llama2-chat-q4_K_M
codeup:13b-llama2-chat-q5_0
codeup:13b-llama2-chat-q5_1
codeup:13b-llama2-chat-q5_K_S
codeup:13b-llama2-chat-q5_K_M
codeup:13b-llama2-chat-q6_K
codeup:13b-llama2-chat-q8_0
codeup:13b-llama2-chat-fp16
marco-o1:latest
marco-o1
marco-o1:7b
marco-o1:7b-q4_K_M
marco-o1:7b-q8_0
marco-o1:7b-fp16
stablelm-zephyr:latest
stablelm-zephyr
stablelm-zephyr:3b
stablelm-zephyr:3b-q2_K
stablelm-zephyr:3b-q3_K_S
stablelm-zephyr:3b-q3_K_M
stablelm-zephyr:3b-q3_K_L
stablelm-zephyr:3b-q4_0
stablelm-zephyr:3b-q4_1
stablelm-zephyr:3b-q4_K_S
stablelm-zephyr:3b-q4_K_M
stablelm-zephyr:3b-q5_0
stablelm-zephyr:3b-q5_1
stablelm-zephyr:3b-q5_K_S
stablelm-zephyr:3b-q5_K_M
stablelm-zephyr:3b-q6_K
stablelm-zephyr:3b-q8_0
stablelm-zephyr:3b-fp16
falcon2:latest
falcon2
falcon2:11b
falcon2:11b-q2_K
falcon2:11b-q3_K_S
falcon2:11b-q3_K_M
falcon2:11b-q3_K_L
falcon2:11b-q4_0
falcon2:11b-q4_1
falcon2:11b-q4_K_S
falcon2:11b-q4_K_M
falcon2:11b-q5_0
falcon2:11b-q5_1
falcon2:11b-q5_K_S
falcon2:11b-q5_K_M
falcon2:11b-q6_K
falcon2:11b-q8_0
falcon2:11b-fp16
solar-pro:latest
solar-pro
solar-pro:preview
solar-pro:22b
solar-pro:22b-preview-instruct-q2_K
solar-pro:22b-preview-instruct-q3_K_S
solar-pro:22b-preview-instruct-q3_K_M
solar-pro:22b-preview-instruct-q3_K_L
solar-pro:22b-preview-instruct-q4_0
solar-pro:22b-preview-instruct-q4_1
solar-pro:22b-preview-instruct-q4_K_S
solar-pro:22b-preview-instruct-q4_K_M
solar-pro:22b-preview-instruct-q5_0
solar-pro:22b-preview-instruct-q5_1
solar-pro:22b-preview-instruct-q5_K_S
solar-pro:22b-preview-instruct-q5_K_M
solar-pro:22b-preview-instruct-q6_K
solar-pro:22b-preview-instruct-q8_0
solar-pro:22b-preview-instruct-fp16
duckdb-nsql:latest
duckdb-nsql
duckdb-nsql:7b
duckdb-nsql:7b-q2_K
duckdb-nsql:7b-q3_K_S
duckdb-nsql:7b-q3_K_M
duckdb-nsql:7b-q3_K_L
duckdb-nsql:7b-q4_0
duckdb-nsql:7b-q4_1
duckdb-nsql:7b-q4_K_S
duckdb-nsql:7b-q4_K_M
duckdb-nsql:7b-q5_0
duckdb-nsql:7b-q5_1
duckdb-nsql:7b-q5_K_S
duckdb-nsql:7b-q5_K_M
duckdb-nsql:7b-q6_K
duckdb-nsql:7b-q8_0
duckdb-nsql:7b-fp16
mistrallite:latest
mistrallite
mistrallite:7b
mistrallite:7b-v0.1-q2_K
mistrallite:7b-v0.1-q3_K_S
mistrallite:7b-v0.1-q3_K_M
mistrallite:7b-v0.1-q3_K_L
mistrallite:7b-v0.1-q4_0
mistrallite:7b-v0.1-q4_1
mistrallite:7b-v0.1-q4_K_S
mistrallite:7b-v0.1-q4_K_M
mistrallite:7b-v0.1-q5_0
mistrallite:7b-v0.1-q5_1
mistrallite:7b-v0.1-q5_K_S
mistrallite:7b-v0.1-q5_K_M
mistrallite:7b-v0.1-q6_K
mistrallite:7b-v0.1-q8_0
mistrallite:7b-v0.1-fp16
magicoder:latest
magicoder
magicoder:7b
magicoder:7b-s-cl
magicoder:7b-s-cl-q2_K
magicoder:7b-s-cl-q3_K_S
magicoder:7b-s-cl-q3_K_M
magicoder:7b-s-cl-q3_K_L
magicoder:7b-s-cl-q4_0
magicoder:7b-s-cl-q4_1
magicoder:7b-s-cl-q4_K_S
magicoder:7b-s-cl-q4_K_M
magicoder:7b-s-cl-q5_0
magicoder:7b-s-cl-q5_1
magicoder:7b-s-cl-q5_K_S
magicoder:7b-s-cl-q5_K_M
magicoder:7b-s-cl-q6_K
magicoder:7b-s-cl-q8_0
magicoder:7b-s-cl-fp16
codebooga:latest
codebooga
codebooga:34b
codebooga:34b-v0.1-q2_K
codebooga:34b-v0.1-q3_K_S
codebooga:34b-v0.1-q3_K_M
codebooga:34b-v0.1-q3_K_L
codebooga:34b-v0.1-q4_0
codebooga:34b-v0.1-q4_1
codebooga:34b-v0.1-q4_K_M
codebooga:34b-v0.1-q5_0
codebooga:34b-v0.1-q5_1
codebooga:34b-v0.1-q5_K_S
codebooga:34b-v0.1-q5_K_M
codebooga:34b-v0.1-q6_K
codebooga:34b-v0.1-q8_0
codebooga:34b-v0.1-fp16
bespoke-minicheck:latest
bespoke-minicheck
bespoke-minicheck:7b
bespoke-minicheck:7b-q2_K
bespoke-minicheck:7b-q3_K_S
bespoke-minicheck:7b-q3_K_M
bespoke-minicheck:7b-q3_K_L
bespoke-minicheck:7b-q4_0
bespoke-minicheck:7b-q4_1
bespoke-minicheck:7b-q4_K_S
bespoke-minicheck:7b-q4_K_M
bespoke-minicheck:7b-q5_0
bespoke-minicheck:7b-q5_1
bespoke-minicheck:7b-q5_K_S
bespoke-minicheck:7b-q5_K_M
bespoke-minicheck:7b-q6_K
bespoke-minicheck:7b-q8_0
bespoke-minicheck:7b-fp16
deepseek-ocr:latest
deepseek-ocr
deepseek-ocr:3b
deepseek-ocr:3b-bf16
nuextract:latest
nuextract
nuextract:3.8b
nuextract:3.8b-q2_K
nuextract:3.8b-q3_K_S
nuextract:3.8b-q3_K_M
nuextract:3.8b-q3_K_L
nuextract:3.8b-q4_0
nuextract:3.8b-q4_1
nuextract:3.8b-q4_K_S
nuextract:3.8b-q4_K_M
nuextract:3.8b-q5_0
nuextract:3.8b-q5_1
nuextract:3.8b-q5_K_S
nuextract:3.8b-q5_K_M
nuextract:3.8b-q6_K
nuextract:3.8b-q8_0
nuextract:3.8b-fp16
wizard-vicuna:latest
wizard-vicuna
wizard-vicuna:13b
wizard-vicuna:13b-q2_K
wizard-vicuna:13b-q3_K_S
wizard-vicuna:13b-q3_K_M
wizard-vicuna:13b-q3_K_L
wizard-vicuna:13b-q4_0
wizard-vicuna:13b-q4_1
wizard-vicuna:13b-q4_K_S
wizard-vicuna:13b-q4_K_M
wizard-vicuna:13b-q5_0
wizard-vicuna:13b-q5_1
wizard-vicuna:13b-q5_K_S
wizard-vicuna:13b-q5_K_M
wizard-vicuna:13b-q6_K
wizard-vicuna:13b-q8_0
wizard-vicuna:13b-fp16
granite3-guardian:latest
granite3-guardian
granite3-guardian:2b
granite3-guardian:8b
granite3-guardian:2b-q8_0
granite3-guardian:2b-fp16
granite3-guardian:8b-q5_K_S
granite3-guardian:8b-q5_K_M
granite3-guardian:8b-q6_K
granite3-guardian:8b-q8_0
granite3-guardian:8b-fp16
firefunction-v2:latest
firefunction-v2
firefunction-v2:70b
firefunction-v2:70b-q2_K
firefunction-v2:70b-q3_K_S
firefunction-v2:70b-q3_K_M
firefunction-v2:70b-q3_K_L
firefunction-v2:70b-q4_0
firefunction-v2:70b-q4_1
firefunction-v2:70b-q4_K_S
firefunction-v2:70b-q4_K_M
firefunction-v2:70b-q5_0
firefunction-v2:70b-q5_1
firefunction-v2:70b-q5_K_S
firefunction-v2:70b-q5_K_M
firefunction-v2:70b-q6_K
firefunction-v2:70b-q8_0
firefunction-v2:70b-fp16
megadolphin:latest
megadolphin
megadolphin:v2.2
megadolphin:120b
megadolphin:120b-v2.2
megadolphin:120b-v2.2-q2_K
megadolphin:120b-v2.2-q3_K_S
megadolphin:120b-v2.2-q3_K_M
megadolphin:120b-v2.2-q3_K_L
megadolphin:120b-v2.2-q4_0
megadolphin:120b-v2.2-q4_1
megadolphin:120b-v2.2-q4_K_S
megadolphin:120b-v2.2-q4_K_M
megadolphin:120b-v2.2-q5_0
megadolphin:120b-v2.2-q5_1
megadolphin:120b-v2.2-q5_K_S
megadolphin:120b-v2.2-q5_K_M
megadolphin:120b-v2.2-q6_K
megadolphin:120b-v2.2-q8_0
megadolphin:120b-v2.2-fp16
notux:latest
notux
notux:8x7b
notux:8x7b-v1
notux:8x7b-v1-q2_K
notux:8x7b-v1-q3_K_S
notux:8x7b-v1-q3_K_M
notux:8x7b-v1-q3_K_L
notux:8x7b-v1-q4_0
notux:8x7b-v1-q4_1
notux:8x7b-v1-q4_K_S
notux:8x7b-v1-q4_K_M
notux:8x7b-v1-q5_0
notux:8x7b-v1-q5_1
notux:8x7b-v1-q5_K_S
notux:8x7b-v1-q5_K_M
notux:8x7b-v1-q6_K
notux:8x7b-v1-q8_0
notux:8x7b-v1-fp16
open-orca-platypus2:latest
open-orca-platypus2
open-orca-platypus2:13b
open-orca-platypus2:13b-q2_K
open-orca-platypus2:13b-q3_K_S
open-orca-platypus2:13b-q3_K_M
open-orca-platypus2:13b-q3_K_L
open-orca-platypus2:13b-q4_0
open-orca-platypus2:13b-q4_1
open-orca-platypus2:13b-q4_K_S
open-orca-platypus2:13b-q4_K_M
open-orca-platypus2:13b-q5_0
open-orca-platypus2:13b-q5_1
open-orca-platypus2:13b-q5_K_S
open-orca-platypus2:13b-q5_K_M
open-orca-platypus2:13b-q6_K
open-orca-platypus2:13b-q8_0
open-orca-platypus2:13b-fp16
sailor2:latest
sailor2
sailor2:1b
sailor2:8b
sailor2:20b
sailor2:1b-chat-q4_K_M
sailor2:1b-chat-q8_0
sailor2:1b-chat-fp16
sailor2:8b-chat-q4_K_M
sailor2:8b-chat-q8_0
sailor2:8b-chat-fp16
sailor2:20b-chat-q4_K_M
sailor2:20b-chat-q8_0
sailor2:20b-chat-fp16
notus:latest
notus
notus:7b
notus:7b-v1
notus:7b-v1-q2_K
notus:7b-v1-q3_K_S
notus:7b-v1-q3_K_M
notus:7b-v1-q3_K_L
notus:7b-v1-q4_0
notus:7b-v1-q4_1
notus:7b-v1-q4_K_S
notus:7b-v1-q4_K_M
notus:7b-v1-q5_0
notus:7b-v1-q5_1
notus:7b-v1-q5_K_S
notus:7b-v1-q5_K_M
notus:7b-v1-q6_K
notus:7b-v1-q8_0
notus:7b-v1-fp16
goliath:latest
goliath
goliath:120b-q2_K
goliath:120b-q3_K_S
goliath:120b-q3_K_M
goliath:120b-q3_K_L
goliath:120b-q4_0
goliath:120b-q4_1
goliath:120b-q4_K_S
goliath:120b-q4_K_M
goliath:120b-q5_0
goliath:120b-q5_1
goliath:120b-q5_K_S
goliath:120b-q5_K_M
goliath:120b-q6_K
goliath:120b-q8_0
goliath:120b-fp16
alfred:latest
alfred
alfred:40b
alfred:40b-1023-q4_0
alfred:40b-1023-q4_1
alfred:40b-1023-q5_0
alfred:40b-1023-q5_1
alfred:40b-1023-q8_0
command-r7b-arabic:latest
command-r7b-arabic
command-r7b-arabic:7b
command-r7b-arabic:7b-02-2025-q4_K_M
command-r7b-arabic:7b-02-2025-q8_0
command-r7b-arabic:7b-02-2025-fp16
gemini-3-pro-preview:latest
gemini-3-pro-preview
glm-4.6:cloud
gpt-oss-safeguard:latest
gpt-oss-safeguard
gpt-oss-safeguard:20b
gpt-oss-safeguard:120b
minimax-m2:cloud
kimi-k2:1t-cloud
cogito-2.1:latest
cogito-2.1
cogito-2.1:671b
cogito-2.1:671b-cloud
cogito-2.1:671b-q4_K_M
cogito-2.1:671b-q8_0
cogito-2.1:671b-fp16
kimi-k2-thinking:cloud
rnj-1:latest
rnj-1
rnj-1:8b
rnj-1:8b-cloud
rnj-1:8b-instruct-q4_K_M
rnj-1:8b-instruct-q8_0
rnj-1:8b-instruct-fp16
nomic-embed-text-v2-moe:latest
nomic-embed-text-v2-moe
olmo-3.1:latest
olmo-3.1
olmo-3.1:32b
olmo-3.1:32b-instruct
olmo-3.1:32b-instruct-q4_K_M
olmo-3.1:32b-instruct-q8_0
olmo-3.1:32b-instruct-fp16
olmo-3.1:32b-think
olmo-3.1:32b-think-q4_K_M
olmo-3.1:32b-think-q8_0
olmo-3.1:32b-think-fp16
deepseek-v3.2:cloud
mistral-large-3:675b-cloud
`
	_ = modelsStr
	// models := strings.Fields(modelsStr)
	//
	// for _, field := range models {
	// 	fmt.Printf(`{input: "%s", expected: "%s"},`+"\n", field, FormatModelTitle(field))
	// }
}
